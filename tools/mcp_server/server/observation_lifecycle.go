package server

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"roraddons/tools/mcp_server/model"
	"roraddons/tools/mcp_server/schema"
)

const (
	lifecycleStatusCandidate = "candidate"
	lifecycleStatusAccepted  = "accepted"
	lifecycleStatusRejected  = "rejected"
	lifecycleStatusPromoted  = "promoted"

	defaultQueueFileName    = "accepted.ndjson"
	defaultRejectedFileName = "rejected.ndjson"
)

// lifecycleRecord is the on-disk representation of one observation entry.
type lifecycleRecord struct {
	IngestedAtUTC   string         `json:"ingested_at_utc"`
	SourcePath      string         `json:"source_path,omitempty"`
	Observation     map[string]any `json:"observation"`
	LifecycleStatus string         `json:"lifecycle_status,omitempty"`
	Review          *lifecycleReview `json:"review,omitempty"`
	PromotedAtUTC   string         `json:"promoted_at_utc,omitempty"`
}

type lifecycleReview struct {
	Verdict       string `json:"verdict"`
	Reviewer      string `json:"reviewer,omitempty"`
	Notes         string `json:"notes,omitempty"`
	ReviewedAtUTC string `json:"reviewed_at_utc"`
}

// effectiveStatus returns the record's lifecycle status, defaulting to "candidate".
func (r lifecycleRecord) effectiveStatus() string {
	if r.LifecycleStatus != "" {
		return r.LifecycleStatus
	}
	return lifecycleStatusCandidate
}

// entryID extracts the observation's entry_id.
func (r lifecycleRecord) entryID() string {
	if r.Observation != nil {
		if id, ok := r.Observation["entry_id"].(string); ok {
			return id
		}
	}
	return ""
}

// queueFilePath returns the absolute path to the main accepted.ndjson queue.
func (a *App) queueFilePath() string {
	return filepath.Join(a.feedingRoot, "review_queue", defaultQueueFileName)
}

// rejectedFilePath returns the absolute path to the rejected.ndjson file.
func (a *App) rejectedFilePath() string {
	return filepath.Join(a.feedingRoot, "review_queue", defaultRejectedFileName)
}

// workspaceRoot derives the project root from the feeding root (feeding root is 3 dirs below project root).
func (a *App) workspaceRoot() string {
	abs, err := filepath.Abs(a.feedingRoot)
	if err != nil {
		return "."
	}
	return filepath.Clean(filepath.Join(abs, "..", "..", ".."))
}

// readQueueRecords reads all lifecycle records from an ndjson file.
// Missing or empty files return an empty slice (not an error).
func readQueueRecords(path string) ([]lifecycleRecord, []model.Warning, error) {
	f, err := os.Open(path)
	if os.IsNotExist(err) {
		return nil, nil, nil
	}
	if err != nil {
		return nil, nil, fmt.Errorf("open queue file: %w", err)
	}
	defer f.Close()

	var records []lifecycleRecord
	var warnings []model.Warning
	scanner := bufio.NewScanner(f)
	// Allow large lines (observations can be substantial JSON).
	const maxLine = 1 << 20 // 1 MiB
	buf := make([]byte, maxLine)
	scanner.Buffer(buf, maxLine)
	lineNum := 0
	for scanner.Scan() {
		lineNum++
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		var rec lifecycleRecord
		if err := json.Unmarshal([]byte(line), &rec); err != nil {
			warnings = append(warnings, model.Warning{
				Code:    "parse_error",
				Message: fmt.Sprintf("line %d: %v", lineNum, err),
			})
			continue
		}
		records = append(records, rec)
	}
	if err := scanner.Err(); err != nil {
		return records, warnings, fmt.Errorf("scan queue file: %w", err)
	}
	return records, warnings, nil
}

// writeQueueRecords atomically rewrites the ndjson queue file.
func writeQueueRecords(path string, records []lifecycleRecord) error {
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return fmt.Errorf("mkdir for queue: %w", err)
	}
	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	for _, rec := range records {
		if err := enc.Encode(rec); err != nil {
			return fmt.Errorf("marshal record %s: %w", rec.entryID(), err)
		}
	}
	tmp := path + ".tmp"
	if err := os.WriteFile(tmp, buf.Bytes(), 0o644); err != nil {
		return fmt.Errorf("write temp queue: %w", err)
	}
	if err := os.Rename(tmp, path); err != nil {
		_ = os.Remove(tmp)
		return fmt.Errorf("rename queue: %w", err)
	}
	return nil
}

// appendRejectedRecord appends a record to the rejected.ndjson file.
func appendRejectedRecord(path string, rec lifecycleRecord) error {
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return fmt.Errorf("mkdir for rejected: %w", err)
	}
	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
	if err != nil {
		return fmt.Errorf("open rejected file: %w", err)
	}
	defer f.Close()
	b, err := json.Marshal(rec)
	if err != nil {
		return fmt.Errorf("marshal rejected record: %w", err)
	}
	_, err = f.Write(append(b, '\n'))
	return err
}

// summarizeRecord builds a compact ObservationSummary from a lifecycle record.
func summarizeRecord(rec lifecycleRecord) schema.ObservationSummary {
	obs := rec.Observation
	s := schema.ObservationSummary{
		ObservationID: rec.entryID(),
		Status:        rec.effectiveStatus(),
		CreatedAt:     rec.IngestedAtUTC,
	}
	if obs == nil {
		return s
	}
	if seeds, ok := obs["target_seeds"].([]any); ok {
		for _, seed := range seeds {
			if sv, ok := seed.(string); ok {
				s.TargetSeeds = append(s.TargetSeeds, sv)
			}
		}
	}
	if conf, ok := obs["confidence"].(map[string]any); ok {
		if overall, ok := conf["overall"].(string); ok {
			s.Confidence = overall
		}
	}
	if claims, ok := obs["claims"].([]any); ok {
		s.EvidenceCount = len(claims)
		if len(claims) > 0 {
			if claim, ok := claims[0].(map[string]any); ok {
				if stmt, ok := claim["statement"].(string); ok {
					s.ClaimSummary = stmt
				}
			}
		}
	}
	if ev, ok := obs["evidence"].([]any); ok && s.EvidenceCount == 0 {
		s.EvidenceCount = len(ev)
	}
	if src, ok := obs["source"].(map[string]any); ok {
		if addon, ok := src["addon"].(string); ok {
			s.SourceAddon = addon
		}
	}
	return s
}

// listPendingObservations returns observations with status candidate or accepted.
func (a *App) listPendingObservations(req schema.ListPendingObservationsRequest) schema.ListPendingObservationsResponse {
	records, warnings, err := readQueueRecords(a.queueFilePath())
	resp := schema.ListPendingObservationsResponse{Warnings: warnings}
	if err != nil {
		resp.Warnings = append(resp.Warnings, model.Warning{Code: "read_error", Message: err.Error()})
		return resp
	}

	statusFilter := strings.ToLower(strings.TrimSpace(req.StatusFilter))
	targetFilter := strings.TrimSpace(req.TargetFilter)

	for _, rec := range records {
		status := rec.effectiveStatus()
		// Default: show candidate and accepted; if filter given, apply it.
		if statusFilter != "" {
			if status != statusFilter {
				continue
			}
		} else {
			if status != lifecycleStatusCandidate && status != lifecycleStatusAccepted {
				continue
			}
		}
		// Target filter: match against any target_seeds entry.
		if targetFilter != "" {
			matched := false
			if seeds, ok := rec.Observation["target_seeds"].([]any); ok {
				for _, seed := range seeds {
					if sv, ok := seed.(string); ok && strings.Contains(sv, targetFilter) {
						matched = true
						break
					}
				}
			}
			if !matched {
				continue
			}
		}
		resp.Observations = append(resp.Observations, summarizeRecord(rec))
		if req.Limit > 0 && len(resp.Observations) >= req.Limit {
			break
		}
	}
	resp.TotalCount = len(resp.Observations)
	return resp
}

// reviewObservation applies an accept or reject verdict to a queued observation.
func (a *App) reviewObservation(req schema.ReviewObservationRequest) schema.ReviewObservationResponse {
	resp := schema.ReviewObservationResponse{
		ObservationID: req.ObservationID,
		Verdict:       req.Verdict,
	}

	queuePath := a.queueFilePath()
	records, warnings, err := readQueueRecords(queuePath)
	resp.Warnings = warnings
	if err != nil {
		resp.Errors = append(resp.Errors, "read queue: "+err.Error())
		return resp
	}

	idx := -1
	for i, rec := range records {
		if rec.entryID() == req.ObservationID {
			idx = i
			break
		}
	}
	if idx < 0 {
		resp.Errors = append(resp.Errors, "observation not found: "+req.ObservationID)
		return resp
	}

	current := records[idx].effectiveStatus()
	if current == lifecycleStatusPromoted {
		resp.Warnings = append(resp.Warnings, model.Warning{
			Code:    "already_promoted",
			Message: "observation is already promoted; review has no effect on promotion status",
		})
	}
	if current == lifecycleStatusRejected && req.Verdict == "reject" {
		resp.Warnings = append(resp.Warnings, model.Warning{
			Code:    "already_rejected",
			Message: "observation is already rejected",
		})
	}

	now := time.Now().UTC().Format(time.RFC3339)
	newStatus := lifecycleStatusAccepted
	if req.Verdict == "reject" {
		newStatus = lifecycleStatusRejected
	}

	records[idx].LifecycleStatus = newStatus
	records[idx].Review = &lifecycleReview{
		Verdict:       req.Verdict,
		Reviewer:      strings.TrimSpace(req.Reviewer),
		Notes:         strings.TrimSpace(req.Notes),
		ReviewedAtUTC: now,
	}

	if err := writeQueueRecords(queuePath, records); err != nil {
		resp.Errors = append(resp.Errors, "write queue: "+err.Error())
		return resp
	}

	// For rejections: also append to the durable rejected store.
	if newStatus == lifecycleStatusRejected {
		if err := appendRejectedRecord(a.rejectedFilePath(), records[idx]); err != nil {
			resp.Warnings = append(resp.Warnings, model.Warning{
				Code:    "rejected_store_write_failed",
				Message: "rejection written to queue but rejected store append failed: " + err.Error(),
			})
		}
	}

	resp.Status = newStatus
	return resp
}

// promoteObservation patches seed files and marks the observation as promoted.
func (a *App) promoteObservation(req schema.PromoteObservationRequest) schema.PromoteObservationResponse {
	resp := schema.PromoteObservationResponse{
		ObservationID: req.ObservationID,
		DryRun:        req.DryRun,
	}

	queuePath := a.queueFilePath()
	records, warnings, err := readQueueRecords(queuePath)
	resp.Warnings = warnings
	if err != nil {
		resp.Errors = append(resp.Errors, "read queue: "+err.Error())
		return resp
	}

	idx := -1
	for i, rec := range records {
		if rec.entryID() == req.ObservationID {
			idx = i
			break
		}
	}
	if idx < 0 {
		resp.Errors = append(resp.Errors, "observation not found: "+req.ObservationID)
		return resp
	}

	rec := records[idx]
	status := rec.effectiveStatus()
	if status == lifecycleStatusRejected {
		resp.Errors = append(resp.Errors, "cannot promote a rejected observation")
		return resp
	}
	if status == lifecycleStatusPromoted {
		resp.Warnings = append(resp.Warnings, model.Warning{
			Code:    "already_promoted",
			Message: "observation was already promoted",
		})
		resp.Promoted = true
		return resp
	}
	if status == lifecycleStatusCandidate {
		resp.Warnings = append(resp.Warnings, model.Warning{
			Code:    "not_yet_reviewed",
			Message: "promoting a candidate observation that has not been explicitly accepted; consider using review_observation first",
		})
	}

	// Determine seed targets.
	var seedPaths []string
	if strings.TrimSpace(req.SeedPathOverride) != "" {
		seedPaths = []string{strings.TrimSpace(req.SeedPathOverride)}
	} else if rec.Observation != nil {
		if seeds, ok := rec.Observation["target_seeds"].([]any); ok {
			for _, s := range seeds {
				if sv, ok := s.(string); ok && strings.TrimSpace(sv) != "" {
					seedPaths = append(seedPaths, strings.TrimSpace(sv))
				}
			}
		}
	}
	if len(seedPaths) == 0 {
		resp.Errors = append(resp.Errors, "no target seed paths found; set target_seeds in observation or use seed_path_override")
		return resp
	}
	resp.TargetSeeds = seedPaths

	wsRoot := a.workspaceRoot()
	entryID := rec.entryID()
	promotedAt := time.Now().UTC().Format(time.RFC3339)

	// Build the markdown fragment to append for each claim.
	fragment := buildPromotionFragment(rec, entryID, promotedAt)

	for _, seedRelPath := range seedPaths {
		seedAbsPath := filepath.Join(wsRoot, filepath.FromSlash(seedRelPath))

		// Ensure the resolved path stays within the workspace root (prevent traversal).
		rel, err := filepath.Rel(wsRoot, seedAbsPath)
		if err != nil || strings.HasPrefix(rel, "..") {
			resp.Errors = append(resp.Errors, fmt.Sprintf("seed path %q escapes workspace root; skipped", seedRelPath))
			continue
		}

		update := schema.SeedUpdate{SeedPath: seedRelPath}

		// Check for duplicate promotion marker.
		if isDuplicatePromotion(seedAbsPath, entryID) {
			update.Duplicate = true
			resp.SeedUpdates = append(resp.SeedUpdates, update)
			resp.Warnings = append(resp.Warnings, model.Warning{
				Code:    "duplicate_promotion",
				Message: fmt.Sprintf("entry %s is already promoted in %s", entryID, seedRelPath),
			})
			continue
		}

		if req.DryRun {
			update.Appended = false
			resp.SeedUpdates = append(resp.SeedUpdates, update)
			continue
		}

		if err := appendToSeedFile(seedAbsPath, fragment); err != nil {
			resp.Errors = append(resp.Errors, fmt.Sprintf("append to %s: %v", seedRelPath, err))
			resp.SeedUpdates = append(resp.SeedUpdates, update)
			continue
		}
		update.Appended = true
		resp.SeedUpdates = append(resp.SeedUpdates, update)
	}

	if len(resp.Errors) > 0 {
		return resp
	}

	if !req.DryRun {
		records[idx].LifecycleStatus = lifecycleStatusPromoted
		records[idx].PromotedAtUTC = promotedAt
		if err := writeQueueRecords(queuePath, records); err != nil {
			resp.Warnings = append(resp.Warnings, model.Warning{
				Code:    "queue_update_failed",
				Message: "seed files were updated but queue status update failed: " + err.Error(),
			})
		} else {
			resp.Promoted = true
		}
	}

	return resp
}

// listRejectedObservations returns records from the durable rejected store.
func (a *App) listRejectedObservations(req schema.ListRejectedObservationsRequest) schema.ListRejectedObservationsResponse {
	records, warnings, err := readQueueRecords(a.rejectedFilePath())
	resp := schema.ListRejectedObservationsResponse{Warnings: warnings}
	if err != nil {
		resp.Warnings = append(resp.Warnings, model.Warning{Code: "read_error", Message: err.Error()})
		return resp
	}

	for _, rec := range records {
		s := schema.RejectedObservationSummary{
			ObservationID: rec.entryID(),
			Status:        rec.effectiveStatus(),
			CreatedAt:     rec.IngestedAtUTC,
		}
		if rec.Observation != nil {
			if seeds, ok := rec.Observation["target_seeds"].([]any); ok {
				for _, seed := range seeds {
					if sv, ok := seed.(string); ok {
						s.TargetSeeds = append(s.TargetSeeds, sv)
					}
				}
			}
			if conf, ok := rec.Observation["confidence"].(map[string]any); ok {
				if overall, ok := conf["overall"].(string); ok {
					s.Confidence = overall
				}
			}
			if claims, ok := rec.Observation["claims"].([]any); ok && len(claims) > 0 {
				if claim, ok := claims[0].(map[string]any); ok {
					if stmt, ok := claim["statement"].(string); ok {
						s.ClaimSummary = stmt
					}
				}
			}
		}
		if rec.Review != nil {
			s.RejectionReason = rec.Review.Notes
			s.RejectedAt = rec.Review.ReviewedAtUTC
			s.Reviewer = rec.Review.Reviewer
		}
		resp.Observations = append(resp.Observations, s)
		if req.Limit > 0 && len(resp.Observations) >= req.Limit {
			break
		}
	}
	resp.TotalCount = len(resp.Observations)
	return resp
}

// regenerateFromPromotedKnowledge builds and optionally executes regeneration commands.
func (a *App) regenerateFromPromotedKnowledge(req schema.RegenerateRequest) schema.RegenerateResponse {
	scope := strings.ToLower(strings.TrimSpace(req.Scope))
	if scope == "" {
		scope = "full"
	}
	resp := schema.RegenerateResponse{
		Scope:  scope,
		DryRun: req.DryRun,
	}

	wsRoot := a.workspaceRoot()

	type stepDef struct {
		label string
		args  []string // arguments to "go run"
	}

	platformArgs := []string{"run", "./tools/api_doc_gen", "generate", "platform", "./docs/addon-api", "./docs/war-api"}
	siteArgs := []string{"run", "./tools/api_doc_gen", "generate", "site", "./docs/war-api", "./docs/site/content"}

	var steps []stepDef
	switch scope {
	case "platform":
		steps = []stepDef{{label: "generate platform docs", args: platformArgs}}
	case "site":
		steps = []stepDef{{label: "generate site content", args: siteArgs}}
	default: // "full"
		steps = []stepDef{
			{label: "generate platform docs", args: platformArgs},
			{label: "generate site content", args: siteArgs},
		}
	}

	allSuccess := true
	for _, s := range steps {
		step := schema.RegenerationStep{
			Label:   s.label,
			Command: "go " + strings.Join(s.args, " "),
		}
		if req.DryRun {
			resp.Steps = append(resp.Steps, step)
			continue
		}
		// s.args are hardcoded constants selected by the validated scope parameter;
		// they are not derived from user input, so passing them to exec.Command is safe.
		cmd := exec.Command("go", s.args...)
		cmd.Dir = wsRoot
		out, execErr := cmd.CombinedOutput()
		step.Output = strings.TrimSpace(string(out))
		if execErr != nil {
			step.Error = execErr.Error()
			step.Success = false
			allSuccess = false
			resp.Warnings = append(resp.Warnings, model.Warning{
				Code:    "step_failed",
				Message: fmt.Sprintf("%s: %v", s.label, execErr),
			})
		} else {
			step.Success = true
		}
		resp.Steps = append(resp.Steps, step)
	}

	if req.DryRun {
		resp.Success = true
	} else {
		resp.Success = allSuccess
	}
	return resp
}

// buildPromotionFragment creates the markdown block appended to a seed file.
func buildPromotionFragment(rec lifecycleRecord, entryID, promotedAt string) string {
	var sb strings.Builder
	sb.WriteString("\n")
	sb.WriteString(fmt.Sprintf("<!-- observation:%s promoted:%s -->\n", entryID, promotedAt))

	obs := rec.Observation
	if obs == nil {
		return sb.String()
	}

	sourceAddon := ""
	if src, ok := obs["source"].(map[string]any); ok {
		if addon, ok := src["addon"].(string); ok {
			sourceAddon = addon
		}
	}
	overallConf := ""
	if conf, ok := obs["confidence"].(map[string]any); ok {
		if overall, ok := conf["overall"].(string); ok {
			overallConf = overall
		}
	}
	if sourceAddon != "" || overallConf != "" {
		sb.WriteString(fmt.Sprintf("> Source: %s | Confidence: %s | Promoted: %s\n\n",
			sourceAddon, overallConf, promotedAt[:10]))
	}

	if claims, ok := obs["claims"].([]any); ok {
		for _, rawClaim := range claims {
			claim, ok := rawClaim.(map[string]any)
			if !ok {
				continue
			}
			conf := ""
			if c, ok := claim["confidence"].(string); ok {
				conf = c
			}
			stmt := ""
			if s, ok := claim["statement"].(string); ok {
				stmt = s
			}
			guidance := ""
			if g, ok := claim["guidance"].(string); ok {
				guidance = g
			}
			if stmt == "" {
				continue
			}
			if conf != "" {
				sb.WriteString(fmt.Sprintf("- `%s`: %s\n", conf, stmt))
			} else {
				sb.WriteString(fmt.Sprintf("- %s\n", stmt))
			}
			if guidance != "" {
				sb.WriteString(fmt.Sprintf("  - Guidance: %s\n", guidance))
			}
		}
	}
	return sb.String()
}

// isDuplicatePromotion checks whether the seed file already contains a promotion marker for entryID.
func isDuplicatePromotion(seedAbsPath, entryID string) bool {
	content, err := os.ReadFile(seedAbsPath)
	if err != nil {
		return false
	}
	marker := fmt.Sprintf("<!-- observation:%s ", entryID)
	return strings.Contains(string(content), marker)
}

// appendToSeedFile creates or appends the fragment to the seed file.
func appendToSeedFile(seedAbsPath, fragment string) error {
	if err := os.MkdirAll(filepath.Dir(seedAbsPath), 0o755); err != nil {
		return fmt.Errorf("mkdir for seed: %w", err)
	}
	f, err := os.OpenFile(seedAbsPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
	if err != nil {
		return fmt.Errorf("open seed file: %w", err)
	}
	defer f.Close()
	_, err = f.WriteString(fragment)
	return err
}
