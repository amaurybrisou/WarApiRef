package tools

import (
	"strings"

	"roraddons/tools/mcp_server/model"
	"roraddons/tools/mcp_server/schema"
)

const feedingObservationV1 = "feeding.observation.v1"

func IngestObservation(req schema.IngestObservationRequest) schema.IngestObservationResponse {
	resp := schema.IngestObservationResponse{}
	obs := req.Observation
	if len(obs) == 0 {
		resp.Errors = append(resp.Errors, "observation is required")
		resp.Accepted = false
		return resp
	}

	schemaVersion := getString(obs, "schema_version")
	entryID := getString(obs, "entry_id")
	resp.SchemaVersion = schemaVersion
	resp.EntryID = entryID

	if schemaVersion != feedingObservationV1 {
		resp.Errors = append(resp.Errors, "schema_version must be feeding.observation.v1")
	}
	if entryID == "" {
		resp.Errors = append(resp.Errors, "entry_id is required")
	}
	if getString(obs, "status") == "" {
		resp.Errors = append(resp.Errors, "status is required")
	}
	if !isObject(obs, "source") {
		resp.Errors = append(resp.Errors, "source object is required")
	}
	if !hasNonEmptyArray(obs, "target_seeds") {
		resp.Errors = append(resp.Errors, "target_seeds must be a non-empty array")
	}
	if !isObject(obs, "confidence") {
		resp.Errors = append(resp.Errors, "confidence object is required")
	}
	if !hasNonEmptyArray(obs, "claims") {
		resp.Errors = append(resp.Errors, "claims must be a non-empty array")
	}
	if !hasNonEmptyArray(obs, "evidence") {
		resp.Errors = append(resp.Errors, "evidence must be a non-empty array")
	}
	if !isObject(obs, "promotion") {
		resp.Errors = append(resp.Errors, "promotion object is required")
	}

	if !req.DryRun && !req.Persist {
		resp.Warnings = append(resp.Warnings, model.Warning{
			Code:    "dry_run_recommended",
			Message: "ingestion currently validates payload only; set dry_run=true to make this explicit",
		})
	}
	if req.DryRun && req.Persist {
		resp.Warnings = append(resp.Warnings, model.Warning{
			Code:    "persist_ignored_in_dry_run",
			Message: "persist=true is ignored when dry_run=true",
		})
	}

	resp.Accepted = len(resp.Errors) == 0
	if resp.Accepted && strings.TrimSpace(req.SourcePath) != "" {
		resp.Warnings = append(resp.Warnings, model.Warning{
			Code:    "source_path_unverified",
			Message: "source_path is metadata only and is not resolved by the server",
		})
	}
	return resp
}

func getString(obj map[string]any, key string) string {
	v, ok := obj[key]
	if !ok {
		return ""
	}
	s, _ := v.(string)
	return strings.TrimSpace(s)
}

func isObject(obj map[string]any, key string) bool {
	v, ok := obj[key]
	if !ok {
		return false
	}
	_, ok = v.(map[string]any)
	return ok
}

func hasNonEmptyArray(obj map[string]any, key string) bool {
	v, ok := obj[key]
	if !ok {
		return false
	}
	a, ok := v.([]any)
	if !ok {
		return false
	}
	return len(a) > 0
}