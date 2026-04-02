package server

import (
	"encoding/json"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"roraddons/tools/mcp_server/model"
	"roraddons/tools/mcp_server/schema"
	toolops "roraddons/tools/mcp_server/tools"
)

const defaultIngestQueueFile = "accepted.ndjson"

func inferFeedingRoot(docsRoot string) string {
	abs, err := filepath.Abs(docsRoot)
	if err != nil {
		return filepath.Join("docs", "platform", "feeding")
	}
	docsDir := filepath.Dir(abs)
	return filepath.Join(docsDir, "platform", "feeding")
}

func (a *App) ingestObservation(req schema.IngestObservationRequest) schema.IngestObservationResponse {
	resp := toolops.IngestObservation(req)
	resp.SourcePath = strings.TrimSpace(req.SourcePath)
	if !resp.Accepted || req.DryRun || !req.Persist {
		return resp
	}

	queueFile := strings.TrimSpace(req.QueueFile)
	if queueFile == "" {
		queueFile = filepath.Join(a.feedingRoot, "review_queue", defaultIngestQueueFile)
	} else if !filepath.IsAbs(queueFile) {
		queueFile = filepath.Join(a.feedingRoot, queueFile)
	}

	record := map[string]any{
		"ingested_at_utc": time.Now().UTC().Format(time.RFC3339),
		"source_path":     resp.SourcePath,
		"observation":     req.Observation,
	}
	if err := appendNDJSON(queueFile, record); err != nil {
		resp.Accepted = false
		resp.Errors = append(resp.Errors, "persist failed: "+err.Error())
	}
	return resp
}

func (a *App) ingestObservationBatch(req schema.IngestObservationBatchRequest) schema.IngestObservationBatchResponse {
	files, scanWarnings := listFeedJSONFiles(a.feedingRoot)
	if req.Limit > 0 && len(files) > req.Limit {
		files = files[:req.Limit]
	}

	resp := schema.IngestObservationBatchResponse{
		TotalFiles: len(files),
		Results:    make([]schema.IngestObservationBatchItem, 0, len(files)),
	}
	resp.Warnings = append(resp.Warnings, scanWarnings...)

	for _, p := range files {
		item := schema.IngestObservationBatchItem{SourcePath: p}
		raw, err := os.ReadFile(p)
		if err != nil {
			item.Errors = []string{"read failed: " + err.Error()}
			resp.RejectedCount++
			resp.Results = append(resp.Results, item)
			if !req.ContinueOnError {
				break
			}
			continue
		}
		var observation map[string]any
		if err := json.Unmarshal(raw, &observation); err != nil {
			item.Errors = []string{"invalid JSON: " + err.Error()}
			resp.RejectedCount++
			resp.Results = append(resp.Results, item)
			if !req.ContinueOnError {
				break
			}
			continue
		}

		single := a.ingestObservation(schema.IngestObservationRequest{
			Observation: observation,
			DryRun:      req.DryRun,
			Persist:     req.Persist,
			QueueFile:   req.QueueFile,
			SourcePath:  p,
		})
		item.Accepted = single.Accepted
		item.EntryID = single.EntryID
		item.Errors = append(item.Errors, single.Errors...)
		if item.Accepted {
			resp.AcceptedCount++
		} else {
			resp.RejectedCount++
			if !req.ContinueOnError {
				resp.Results = append(resp.Results, item)
				break
			}
		}
		resp.Results = append(resp.Results, item)
	}

	resp.ProcessedFiles = len(resp.Results)
	return resp
}

func listFeedJSONFiles(root string) ([]string, []model.Warning) {
	files := []string{}
	warnings := []model.Warning{}
	if strings.TrimSpace(root) == "" {
		warnings = append(warnings, model.Warning{Code: "feeding_root_empty", Message: "feeding root is empty"})
		return files, warnings
	}
	_, err := os.Stat(root)
	if err != nil {
		warnings = append(warnings, model.Warning{Code: "feeding_root_unavailable", Message: "feeding root unavailable: " + err.Error()})
		return files, warnings
	}
	_ = filepath.WalkDir(root, func(path string, d fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			warnings = append(warnings, model.Warning{Code: "walk_error", Message: walkErr.Error()})
			return nil
		}
		if d.IsDir() {
			return nil
		}
		if strings.HasSuffix(strings.ToLower(path), ".feed.json") {
			files = append(files, path)
		}
		return nil
	})
	sort.Strings(files)
	return files, warnings
}

func appendNDJSON(path string, payload map[string]any) error {
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return err
	}
	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
	if err != nil {
		return err
	}
	defer f.Close()
	b, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	if _, err := f.Write(append(b, '\n')); err != nil {
		return err
	}
	return nil
}