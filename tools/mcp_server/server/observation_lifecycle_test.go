package server

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"roraddons/tools/mcp_server/schema"
)

// buildTestRecord creates a minimal lifecycle record for tests.
func buildTestRecord(entryID, status string) lifecycleRecord {
	obs := map[string]any{
		"schema_version": "feeding.observation.v1",
		"entry_id":       entryID,
		"status":         "candidate",
		"source": map[string]any{
			"addon":      "TestAddon",
			"date_utc":   "2026-04-01",
			"validation": "dev_observed",
			"context":    "test",
		},
		"target_seeds": []any{"docs/platform/seeds/xml_conventions.md"},
		"confidence": map[string]any{
			"overall":   "MEDIUM",
			"rationale": "test rationale",
		},
		"claims": []any{
			map[string]any{
				"claim_id":   "c1",
				"kind":       "usage_pattern",
				"confidence": "MEDIUM",
				"statement":  "Test claim statement for " + entryID,
				"guidance":   "Test guidance",
			},
		},
		"evidence": []any{
			map[string]any{
				"evidence_id": "e1",
				"type":        "code_change",
				"summary":     "test evidence",
			},
		},
		"promotion": map[string]any{"notes": "test promotion notes"},
	}
	return lifecycleRecord{
		IngestedAtUTC:   time.Now().UTC().Format(time.RFC3339),
		SourcePath:      "test",
		Observation:     obs,
		LifecycleStatus: status,
	}
}

// setupTestFeedingRoot creates a temp workspace with the same directory structure
// as the real project: tempDir/docs/platform/feeding.
// The returned feedingRoot path is at tempDir/docs/platform/feeding so that
// workspaceRoot() resolves correctly to tempDir.
func setupTestFeedingRoot(t *testing.T, records []lifecycleRecord) (feedingRoot string) {
	t.Helper()
	wsRoot := t.TempDir()
	feedingRoot = filepath.Join(wsRoot, "docs", "platform", "feeding")
	queueDir := filepath.Join(feedingRoot, "review_queue")
	if err := os.MkdirAll(queueDir, 0o755); err != nil {
		t.Fatalf("mkdir queue dir: %v", err)
	}
	queuePath := filepath.Join(queueDir, defaultQueueFileName)
	if err := writeQueueRecords(queuePath, records); err != nil {
		t.Fatalf("write test queue: %v", err)
	}
	return feedingRoot
}

// wsRootFromFeeding returns the workspace root given a feedingRoot path.
func wsRootFromFeeding(feedingRoot string) string {
	return filepath.Clean(filepath.Join(feedingRoot, "..", "..", ".."))
}

// newTestApp creates an App pointing at the real docs/war-api but with
// feedingRoot overridden to the supplied path.
func newTestApp(t *testing.T, feedingRoot string) *App {
	t.Helper()
	app, err := NewApp(docsRoot(t))
	if err != nil {
		t.Fatalf("new app: %v", err)
	}
	app.SetFeedingRoot(feedingRoot)
	return app
}

// --- list_pending_observations ---

func TestListPendingObservationsReturnsCandidates(t *testing.T) {
	records := []lifecycleRecord{
		buildTestRecord("obs_candidate_1", lifecycleStatusCandidate),
		buildTestRecord("obs_accepted_1", lifecycleStatusAccepted),
		buildTestRecord("obs_promoted_1", lifecycleStatusPromoted),
		buildTestRecord("obs_rejected_1", lifecycleStatusRejected),
	}
	app := newTestApp(t, setupTestFeedingRoot(t, records))

	resp := app.listPendingObservations(schema.ListPendingObservationsRequest{})
	if resp.TotalCount != 2 {
		t.Fatalf("expected 2 pending (candidate+accepted), got %d: %+v", resp.TotalCount, resp.Observations)
	}
	for _, obs := range resp.Observations {
		if obs.Status != lifecycleStatusCandidate && obs.Status != lifecycleStatusAccepted {
			t.Fatalf("unexpected status in pending list: %s", obs.Status)
		}
	}
}

func TestListPendingObservationsStatusFilter(t *testing.T) {
	records := []lifecycleRecord{
		buildTestRecord("obs_cand", lifecycleStatusCandidate),
		buildTestRecord("obs_acc", lifecycleStatusAccepted),
	}
	app := newTestApp(t, setupTestFeedingRoot(t, records))

	resp := app.listPendingObservations(schema.ListPendingObservationsRequest{StatusFilter: "accepted"})
	if resp.TotalCount != 1 || resp.Observations[0].ObservationID != "obs_acc" {
		t.Fatalf("expected 1 accepted observation, got %+v", resp)
	}
}

func TestListPendingObservationsLimitRespected(t *testing.T) {
	var records []lifecycleRecord
	for i := 0; i < 5; i++ {
		records = append(records, buildTestRecord(fmt.Sprintf("obs_limit_%d", i), lifecycleStatusCandidate))
	}
	app := newTestApp(t, setupTestFeedingRoot(t, records))

	resp := app.listPendingObservations(schema.ListPendingObservationsRequest{Limit: 2})
	if resp.TotalCount != 2 {
		t.Fatalf("expected 2 results with limit=2, got %d", resp.TotalCount)
	}
}

func TestListPendingObservationsEmptyQueue(t *testing.T) {
	app := newTestApp(t, setupTestFeedingRoot(t, nil))
	resp := app.listPendingObservations(schema.ListPendingObservationsRequest{})
	if resp.TotalCount != 0 {
		t.Fatalf("expected 0 from empty queue, got %d", resp.TotalCount)
	}
}

// --- review_observation ---

func TestReviewObservationAccept(t *testing.T) {
	rec := buildTestRecord("obs_to_accept", lifecycleStatusCandidate)
	feedRoot := setupTestFeedingRoot(t, []lifecycleRecord{rec})
	app := newTestApp(t, feedRoot)

	resp := app.reviewObservation(schema.ReviewObservationRequest{
		ObservationID: "obs_to_accept",
		Verdict:       "accept",
		Reviewer:      "alice",
		Notes:         "looks good",
	})
	if len(resp.Errors) > 0 {
		t.Fatalf("unexpected errors: %v", resp.Errors)
	}
	if resp.Status != lifecycleStatusAccepted {
		t.Fatalf("expected status accepted, got %s", resp.Status)
	}

	// Verify the change was persisted.
	records, _, err := readQueueRecords(filepath.Join(feedRoot, "review_queue", defaultQueueFileName))
	if err != nil {
		t.Fatalf("read queue: %v", err)
	}
	if records[0].effectiveStatus() != lifecycleStatusAccepted {
		t.Fatalf("persisted status should be accepted")
	}
	if records[0].Review == nil || records[0].Review.Reviewer != "alice" {
		t.Fatalf("review metadata not persisted correctly: %+v", records[0].Review)
	}
}

func TestReviewObservationReject(t *testing.T) {
	rec := buildTestRecord("obs_to_reject", lifecycleStatusCandidate)
	feedRoot := setupTestFeedingRoot(t, []lifecycleRecord{rec})
	app := newTestApp(t, feedRoot)

	resp := app.reviewObservation(schema.ReviewObservationRequest{
		ObservationID: "obs_to_reject",
		Verdict:       "reject",
		Reviewer:      "bob",
		Notes:         "insufficient evidence",
	})
	if len(resp.Errors) > 0 {
		t.Fatalf("unexpected errors: %v", resp.Errors)
	}
	if resp.Status != lifecycleStatusRejected {
		t.Fatalf("expected status rejected, got %s", resp.Status)
	}

	// The rejected.ndjson should have been populated.
	rejectedPath := filepath.Join(feedRoot, "review_queue", defaultRejectedFileName)
	rejRecords, _, err := readQueueRecords(rejectedPath)
	if err != nil {
		t.Fatalf("read rejected file: %v", err)
	}
	if len(rejRecords) != 1 || rejRecords[0].entryID() != "obs_to_reject" {
		t.Fatalf("rejected record not found in rejected store: %+v", rejRecords)
	}
}

func TestReviewObservationNotFound(t *testing.T) {
	app := newTestApp(t, setupTestFeedingRoot(t, nil))
	resp := app.reviewObservation(schema.ReviewObservationRequest{
		ObservationID: "nonexistent",
		Verdict:       "accept",
	})
	if len(resp.Errors) == 0 {
		t.Fatalf("expected error for unknown observation_id")
	}
}

// --- promote_observation ---

func TestPromoteObservationSuccess(t *testing.T) {
	rec := buildTestRecord("obs_to_promote", lifecycleStatusAccepted)
	feedRoot := setupTestFeedingRoot(t, []lifecycleRecord{rec})
	wsRoot := wsRootFromFeeding(feedRoot)

	// Create the seed file the test observation targets.
	seedRelPath := "docs/platform/seeds/xml_conventions.md"
	seedAbsPath := filepath.Join(wsRoot, filepath.FromSlash(seedRelPath))
	if err := os.MkdirAll(filepath.Dir(seedAbsPath), 0o755); err != nil {
		t.Fatalf("mkdir seed: %v", err)
	}
	if err := os.WriteFile(seedAbsPath, []byte("# XML Conventions\n"), 0o644); err != nil {
		t.Fatalf("write seed: %v", err)
	}

	app := newTestApp(t, feedRoot)
	resp := app.promoteObservation(schema.PromoteObservationRequest{
		ObservationID: "obs_to_promote",
	})
	if len(resp.Errors) > 0 {
		t.Fatalf("unexpected errors: %v", resp.Errors)
	}
	if !resp.Promoted {
		t.Fatalf("expected promoted=true")
	}
	if len(resp.SeedUpdates) == 0 || !resp.SeedUpdates[0].Appended {
		t.Fatalf("expected at least one appended seed update: %+v", resp.SeedUpdates)
	}

	// Verify seed file was updated.
	content, err := os.ReadFile(seedAbsPath)
	if err != nil {
		t.Fatalf("read seed file: %v", err)
	}
	if !strings.Contains(string(content), "<!-- observation:obs_to_promote") {
		t.Fatalf("promotion marker not found in seed file; content:\n%s", string(content))
	}
	if !strings.Contains(string(content), "Test claim statement") {
		t.Fatalf("claim statement not found in seed file; content:\n%s", string(content))
	}

	// Queue record should be marked promoted.
	records, _, err := readQueueRecords(filepath.Join(feedRoot, "review_queue", defaultQueueFileName))
	if err != nil {
		t.Fatalf("read queue: %v", err)
	}
	if records[0].effectiveStatus() != lifecycleStatusPromoted {
		t.Fatalf("queue record should be marked promoted, got %s", records[0].effectiveStatus())
	}
}

func TestPromoteObservationDryRun(t *testing.T) {
	rec := buildTestRecord("obs_dry_run", lifecycleStatusAccepted)
	feedRoot := setupTestFeedingRoot(t, []lifecycleRecord{rec})
	wsRoot := wsRootFromFeeding(feedRoot)

	seedRelPath := "docs/platform/seeds/xml_conventions.md"
	seedAbsPath := filepath.Join(wsRoot, filepath.FromSlash(seedRelPath))
	if err := os.MkdirAll(filepath.Dir(seedAbsPath), 0o755); err != nil {
		t.Fatalf("mkdir seed: %v", err)
	}
	original := "# XML Conventions\n"
	if err := os.WriteFile(seedAbsPath, []byte(original), 0o644); err != nil {
		t.Fatalf("write seed: %v", err)
	}

	app := newTestApp(t, feedRoot)
	resp := app.promoteObservation(schema.PromoteObservationRequest{
		ObservationID: "obs_dry_run",
		DryRun:        true,
	})
	if len(resp.Errors) > 0 {
		t.Fatalf("unexpected errors: %v", resp.Errors)
	}
	if resp.Promoted {
		t.Fatalf("dry_run should not set promoted=true")
	}

	// Seed file must be unchanged.
	content, err := os.ReadFile(seedAbsPath)
	if err != nil {
		t.Fatalf("read seed: %v", err)
	}
	if string(content) != original {
		t.Fatalf("seed file should not be modified in dry_run mode; got:\n%s", string(content))
	}
}

func TestPromoteObservationDuplicatePrevented(t *testing.T) {
	rec := buildTestRecord("obs_dedup", lifecycleStatusAccepted)
	feedRoot := setupTestFeedingRoot(t, []lifecycleRecord{rec})
	wsRoot := wsRootFromFeeding(feedRoot)

	seedRelPath := "docs/platform/seeds/xml_conventions.md"
	seedAbsPath := filepath.Join(wsRoot, filepath.FromSlash(seedRelPath))
	if err := os.MkdirAll(filepath.Dir(seedAbsPath), 0o755); err != nil {
		t.Fatalf("mkdir seed: %v", err)
	}
	// Pre-populate with an existing promotion marker for the same entry.
	existingContent := "# XML Conventions\n\n<!-- observation:obs_dedup promoted:2026-01-01T00:00:00Z -->\n- Some claim\n"
	if err := os.WriteFile(seedAbsPath, []byte(existingContent), 0o644); err != nil {
		t.Fatalf("write seed: %v", err)
	}

	app := newTestApp(t, feedRoot)
	resp := app.promoteObservation(schema.PromoteObservationRequest{
		ObservationID: "obs_dedup",
	})
	if len(resp.Errors) > 0 {
		t.Fatalf("unexpected errors: %v", resp.Errors)
	}
	for _, u := range resp.SeedUpdates {
		if !u.Duplicate {
			t.Fatalf("expected duplicate=true for all seed updates, got: %+v", u)
		}
	}
	// Seed file should be unchanged.
	content, err := os.ReadFile(seedAbsPath)
	if err != nil {
		t.Fatalf("read seed: %v", err)
	}
	if string(content) != existingContent {
		t.Fatalf("seed file should not have been modified for duplicate promotion")
	}
}

func TestPromoteObservationRejectedBlocked(t *testing.T) {
	rec := buildTestRecord("obs_rejected", lifecycleStatusRejected)
	app := newTestApp(t, setupTestFeedingRoot(t, []lifecycleRecord{rec}))
	resp := app.promoteObservation(schema.PromoteObservationRequest{
		ObservationID: "obs_rejected",
	})
	if len(resp.Errors) == 0 {
		t.Fatalf("expected error when promoting a rejected observation")
	}
}

// --- list_rejected_observations ---

func TestListRejectedObservationsReflectsRejectedStore(t *testing.T) {
	feedRoot := setupTestFeedingRoot(t, nil)
	rejectedPath := filepath.Join(feedRoot, "review_queue", defaultRejectedFileName)

	rec := buildTestRecord("obs_rej_store", lifecycleStatusRejected)
	rec.Review = &lifecycleReview{
		Verdict:       "reject",
		Reviewer:      "carol",
		Notes:         "not enough evidence",
		ReviewedAtUTC: time.Now().UTC().Format(time.RFC3339),
	}
	if err := appendRejectedRecord(rejectedPath, rec); err != nil {
		t.Fatalf("append rejected: %v", err)
	}

	app := newTestApp(t, feedRoot)
	resp := app.listRejectedObservations(schema.ListRejectedObservationsRequest{})
	if resp.TotalCount != 1 {
		t.Fatalf("expected 1 rejected observation, got %d", resp.TotalCount)
	}
	if resp.Observations[0].ObservationID != "obs_rej_store" {
		t.Fatalf("unexpected observation_id: %s", resp.Observations[0].ObservationID)
	}
	if resp.Observations[0].RejectionReason != "not enough evidence" {
		t.Fatalf("unexpected rejection reason: %q", resp.Observations[0].RejectionReason)
	}
}

func TestListRejectedObservationsEmptyReturnsEmpty(t *testing.T) {
	app := newTestApp(t, setupTestFeedingRoot(t, nil))
	resp := app.listRejectedObservations(schema.ListRejectedObservationsRequest{})
	if resp.TotalCount != 0 || len(resp.Observations) != 0 {
		t.Fatalf("expected empty rejected list, got %+v", resp)
	}
}

// --- regenerate_from_promoted_knowledge ---

func TestRegenerateDryRunReturnsSteps(t *testing.T) {
	app := newTestApp(t, setupTestFeedingRoot(t, nil))
	resp := app.regenerateFromPromotedKnowledge(schema.RegenerateRequest{Scope: "full", DryRun: true})
	if !resp.DryRun {
		t.Fatalf("expected dry_run=true in response")
	}
	if resp.Scope != "full" {
		t.Fatalf("expected scope=full, got %s", resp.Scope)
	}
	if len(resp.Steps) != 2 {
		t.Fatalf("expected 2 steps for scope=full, got %d: %+v", len(resp.Steps), resp.Steps)
	}
	if !resp.Success {
		t.Fatalf("dry_run should report success=true")
	}
	for _, step := range resp.Steps {
		if step.Label == "" || step.Command == "" {
			t.Fatalf("step missing label or command: %+v", step)
		}
	}
}

func TestRegeneratePlatformScopeOneStep(t *testing.T) {
	app := newTestApp(t, setupTestFeedingRoot(t, nil))
	resp := app.regenerateFromPromotedKnowledge(schema.RegenerateRequest{Scope: "platform", DryRun: true})
	if len(resp.Steps) != 1 {
		t.Fatalf("expected 1 step for scope=platform, got %d", len(resp.Steps))
	}
}

func TestRegenerateSiteScopeOneStep(t *testing.T) {
	app := newTestApp(t, setupTestFeedingRoot(t, nil))
	resp := app.regenerateFromPromotedKnowledge(schema.RegenerateRequest{Scope: "site", DryRun: true})
	if len(resp.Steps) != 1 {
		t.Fatalf("expected 1 step for scope=site, got %d", len(resp.Steps))
	}
}

// --- MCP dispatch path tests ---

func TestFeedingListPendingViaDispatch(t *testing.T) {
	app, err := NewApp(docsRoot(t))
	if err != nil {
		t.Fatalf("new app: %v", err)
	}
	resp := app.handlePayload([]byte(`{"jsonrpc":"2.0","id":10,"method":"feeding/list_pending","params":{}}`))
	if resp.Error != nil {
		t.Fatalf("feeding/list_pending returned error: %#v", resp.Error)
	}
}

func TestFeedingListRejectedViaDispatch(t *testing.T) {
	app, err := NewApp(docsRoot(t))
	if err != nil {
		t.Fatalf("new app: %v", err)
	}
	resp := app.handlePayload([]byte(`{"jsonrpc":"2.0","id":11,"method":"feeding/list_rejected","params":{}}`))
	if resp.Error != nil {
		t.Fatalf("feeding/list_rejected returned error: %#v", resp.Error)
	}
}

func TestFeedingReviewMissingIDReturnsError(t *testing.T) {
	app, err := NewApp(docsRoot(t))
	if err != nil {
		t.Fatalf("new app: %v", err)
	}
	resp := app.handlePayload([]byte(`{"jsonrpc":"2.0","id":12,"method":"feeding/review","params":{"verdict":"accept"}}`))
	if resp.Error == nil {
		t.Fatalf("expected error for missing observation_id")
	}
}

func TestFeedingRegenerateDryRunViaDispatch(t *testing.T) {
	app, err := NewApp(docsRoot(t))
	if err != nil {
		t.Fatalf("new app: %v", err)
	}
	resp := app.handlePayload([]byte(`{"jsonrpc":"2.0","id":13,"method":"feeding/regenerate","params":{"dry_run":true}}`))
	if resp.Error != nil {
		t.Fatalf("feeding/regenerate returned error: %#v", resp.Error)
	}
}

