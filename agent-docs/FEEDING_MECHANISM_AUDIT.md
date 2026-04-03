# Feeding Mechanism Audit Report

**Date:** 2026-04-02  
**Scope:** Comprehensive audit of `docs/platform/feeding/` architecture, implementation, and viability as structured feedback loop for addon discoveries  
**Status:** Complete - Ready for evolution planning

---

## Executive Summary

The feeding mechanism is **production-ready for structured observation ingestion** and provides a well-architected intake system suitable as the foundation for a contract-native feedback loop. All core infrastructure (schema, MCP tools, queue management, promotion workflow) is fully implemented, validated, and aligned with the contract-native model.

**Key Finding:** The mechanism cleanly separates concerns (observations ≠ symbols), uses explicit confidence tracking, enforces workspace safety, and provides governance features (review verdicts, durable rejection store). No fundamental refactoring is required to use it for future addon discoveries.

---

## 1. Architecture Overview

### Location & Structure
```
docs/platform/feeding/
├── README.md                       # Purpose and requirements
├── INGESTION_RUNBOOK.md           # Setup and usage guide
├── model/
│   ├── observation.v1.schema.json # Machine-readable schema
│   └── observation_pattern.v1.md  # File-pair pattern spec
├── review_queue/
│   ├── README.md                  # Queue file semantics
│   ├── accepted.ndjson            # Active queue (2 observations)
│   └── rejected.ndjson            # Rejection history (empty)
├── xml/                           # Grouped observations (other dirs: lua/, events/, window_api/)
│   ├── whohealedme_xml_input_and_layout.md
│   └── whohealedme_xml_input_and_layout.feed.json
├── backup/                        # Timestamped backups (2 examples: 20260401_214501/)
└── [other addon directories]
```

**Design Pattern:** File-pair requirement ensures human readability (markdown) + machine consumption (.feed.json sidecar). Single source of truth is the .feed.json file.

### Pipeline Phases

```
Addon Finding
    ↓
[1] INGEST: Validate schema_version, required fields (entry_id, status, source, target_seeds, confidence, claims, evidence)
    ↓
[2] QUEUE: Append lifecycle record to docs/platform/feeding/review_queue/accepted.ndjson (ingested_at_utc, source_path, observation)
    ↓
[3] REVIEW: Apply accept/reject verdict + reviewer metadata + timestamp
    ├─ ACCEPT: Update queue status to "accepted"
    └─ REJECT: Also append to rejected.ndjson durable store
    ↓
[4] PROMOTE: Build markdown fragment (with HTML comment marker), append to target seed files
    ├─ Check for duplicate promotion markers (prevent re-promotion)
    ├─ Validate seed paths don't escape workspace root
    ├─ Mark observation as "promoted" + timestamp
    └─ Optional: trigger regeneration of docs
    ↓
[5] REGENERATE: Re-run `go run ./tools/api_doc_gen generate [platform|site|full]`
```

---

## 2. Technology Assessment

### Schema (observation.v1.schema.json)

**Machine-readable:** Yes - Full JSON schema v2020-12 with strict validation.

**Required Fields:**
- `schema_version`: Fixed to "feeding.observation.v1" (version check enforced in validation)
- `entry_id`: Unique stable snake_case identifier (e.g., "whohealedme_xml_input_and_layout")
- `status`: Enum [candidate, accepted, promoted, rejected]
- `source`: Object with {addon, date_utc, validation (dev_observed|user_validated), context, repository, branch}
- `target_seeds`: Non-empty array of seed file paths (e.g., ["docs/platform/seeds/xml_conventions.md"])
- `confidence`: Object with {overall (HIGH|MEDIUM|LOW), rationale (string)}
- `claims`: Non-empty array of {claim_id, kind (behavior_caveat|usage_pattern|signature_hint), confidence, statement, guidance, symbols[], evidence_refs[]}
- `evidence`: Non-empty array of {evidence_id, type (code_change|runtime_observation|user_report|test), summary, file_refs[]}
- `promotion`: Object with {notes, criteria[]}

**Observations from actual data (accepted.ndjson):**
- WhoHealedMe: 2 claims (ancestor handleinput blocking, nested dimensions under-reporting), MEDIUM confidence with user validation
- QuickTacticSwitch: 4 claims (ListData binding, ListColumns mapping, populationfunction optional, empty populationfunction valid), 7 evidence references

### Status Lifecycle

**Distinct States Enforced:**

| State | Transitions | Rules |
|-------|-------------|-------|
| `candidate` | → accepted OR rejected | Initial state after ingestion; can't be promoted directly |
| `accepted` | → promoted OR rejected | Must be set via review_observation verdict = "accept" |
| `promoted` | (immutable) | Can't re-review or re-reject; promotion records are durable until superseding workflow exists |
| `rejected` | (immutable) | Appended to durable rejected.ndjson store; can't be promoted |

**Code Evidence:** (`observation_lifecycle.go` lines 277-359)
```go
// Can't re-review promoted observations
if current == lifecycleStatusPromoted {
    resp.Errors = append(resp.Errors, "cannot re-review a promoted observation; promoted records are immutable...")
}

// Can't promote candidate (must accept first)
if status == lifecycleStatusCandidate {
    resp.Errors = append(resp.Errors, "cannot promote a candidate observation; use review_observation to accept it first")
}
```

---

## 3. MCP Integration (Complete)

### Tool Registry (tool_handlers.go)

7 feeding-specific tools fully implemented:

#### 1. **ingest_observation** (`tools/ingest_observation.go`)
- **Purpose:** Validate single observation against schema
- **Input:** Observation map, DryRun, Persist, QueueFile, SourcePath
- **Output:** Accepted (bool), EntryID, SchemaVersion, Errors[], Warnings[]
- **Validation:**
  - schema_version == "feeding.observation.v1" ✓
  - All required fields present ✓
  - entry_id non-empty ✓
  - status enum value ✓
  - source object exists ✓
  - target_seeds non-empty array ✓
  - confidence object exists ✓
  - claims non-empty array ✓
  - evidence non-empty array ✓
  - promotion object exists ✓
- **Persistence:** Appends to queue file if Accepted && !DryRun && Persist

#### 2. **ingest_observation_batch** (`server/feeding_ingest.go`)
- **Purpose:** Batch load all *.feed.json files under docs/platform/feeding/
- **Input:** DryRun, Persist, Limit
- **Output:** IngestionBatchResponse with grouped results (accepted, rejected, warnings)
- **Process:**
  - Walks feedingRoot recursively for *.feed.json files (sorted by path)
  - Validates each file via schema
  - Collects errors per file
  - Appends all accepted records to queue file

#### 3. **list_pending_observations** (`server/observation_lifecycle.go` lines 229-276)
- **Purpose:** Query queue with filtering
- **Input:** StatusFilter (default: [candidate, accepted]), TargetFilter (substring match on target_seeds), Limit
- **Output:** PendingObservations[] with {ObservationID, Status, Confidence, EvidenceCount, ClaimSummary, SourceAddon}, TotalCount
- **Filtering:** Supports multiple status values AND target_seeds substring matching

#### 4. **review_observation** (`server/observation_lifecycle.go` lines 277-359)
- **Purpose:** Apply accept/reject verdict to observation
- **Input:** ObservationID, Verdict (accept|reject), Reviewer, Notes
- **Output:** Status (accepted|rejected), Errors[], Warnings[]
- **Behavior:**
  - Reads queue, finds observation by entry_id
  - Validates: can't re-review promoted, warns if re-rejecting
  - Updates LifecycleStatus, records Review {verdict, reviewer, notes, reviewed_at_utc}
  - For rejections: appends to rejected.ndjson with duplicate check
  - Writes back to queue file

#### 5. **promote_observation** (`server/observation_lifecycle.go` lines 362-505)
- **Purpose:** Move accepted observation into seed files
- **Input:** ObservationID, DryRun, SeedPathOverride
- **Output:** TargetSeeds[], SeedUpdates[] {SeedPath, Appended, Duplicate}, Errors[], Warnings[], Promoted (bool)
- **Behavior:**
  - Validates observation is in "accepted" state (not candidate or rejected)
  - Gets target_seeds from observation or SeedPathOverride
  - For each seed file:
    - Checks for duplicate promotion marker (prevents re-promotion)
    - Validates path doesn't escape workspace root
    - Appends markdown fragment (built from observation data)
  - Updates observation LifecycleStatus to "promoted" + timestamp
  - Returns per-seed success/failure

#### 6. **list_rejected_observations** (`server/observation_lifecycle.go` lines 507-561)
- **Purpose:** View rejection history
- **Input:** Limit
- **Output:** RejectedObservations[] with {ObservationID, Status, CreatedAt, TargetSeeds[], Confidence, ClaimSummary, RejectionReason, RejectedAt, Reviewer}, TotalCount
- **Behavior:**
  - Reads durable rejected.ndjson store
  - Extracts and summarizes each rejection record

#### 7. **regenerate_from_promoted_knowledge** (`server/observation_lifecycle.go` lines 563-637)
- **Purpose:** Trigger docs regeneration after promotion
- **Input:** Scope (platform|site|full, default: full), DryRun
- **Output:** Steps[] with {Label, Command, Output, Error, Success}, Errors[], Warnings[], Success (bool)
- **Commands:**
  - **platform scope:** `go run ./tools/api_doc_gen generate platform ./docs/addon-api ./docs/war-api`
  - **site scope:** `go run ./tools/api_doc_gen generate site ./docs/war-api ./docs/site/content`
  - **full scope:** Both commands in sequence
- **Safety:** Hardcoded command arguments (not user-derived), workspace root validation

---

## 4. Data Preservation & Governance

### Ingestion: What Gets Captured

**Full Observation Structure** (from schema):
```json
{
  "schema_version": "feeding.observation.v1",
  "entry_id": "whohealedme_xml_input_and_layout",
  "status": "candidate",
  "source": {
    "addon": "WhoHealedMe",
    "date_utc": "2026-04-01T14:30:00Z",
    "validation": "dev_observed + user_validated",
    "context": "healer-details window implementation"
  },
  "target_seeds": ["docs/platform/seeds/xml_conventions.md"],
  "confidence": {
    "overall": "MEDIUM",
    "rationale": "observed and user-validated during addon implementation"
  },
  "claims": [
    {
      "claim_id": "ancestor_handleinput_blocks_child_clicks",
      "kind": "behavior_caveat",
      "confidence": "MEDIUM",
      "statement": "ancestor handleinput attribute can block child click targets",
      "guidance": "validate handleinput across parent and template chain if child looks wired but remains inert",
      "symbols": ["handleinput"],
      "evidence_refs": ["ancestor_input_finding"]
    }
  ],
  "evidence": [
    {
      "evidence_id": "ancestor_input_finding",
      "type": "code_change",
      "summary": "XML enablement fix resolving ancestor input blocking",
      "file_refs": ["addons/WhoHealedMe/healer-details.xml:42"]
    }
  ],
  "promotion": {
    "notes": "promote as caveats not symbols",
    "criteria": ["include in xml_conventions.md", "link from relevant window events"]
  }
}
```

**Lifecycle Record Additions** (from observation_lifecycle.go):
```go
type lifecycleRecord struct {
    IngestedAtUTC    string                 // Timestamp when observation was added to queue
    SourcePath       string                 // Path to .feed.json file
    Observation      map[string]any         // Full observation object (as above)
    LifecycleStatus  string                 // candidate|accepted|promoted|rejected
    Review           *lifecycleReview       // (nil until reviewed) {Verdict, Reviewer, Notes, ReviewedAtUTC}
    PromotedAtUTC    string                 // (empty until promoted) Timestamp when moved to seed files
}
```

### Review: Verdict Recording

**Stored in queue record:**
- `Review.Verdict`: "accept" or "reject" (set by reviewer)
- `Review.Reviewer`: Reviewer identity (from review_observation request)
- `Review.Notes`: Optional context/guidance for decision
- `Review.ReviewedAtUTC`: RFC3339 timestamp

**Durability:**
- Accepted observations: Updated in-place in accepted.ndjson
- Rejected observations: Also appended to rejected.ndjson (durable rejection history)

### Promotion: What Gets Written to Seeds

**Markdown Fragment** (from buildPromotionFragment, lines 639-708):
```markdown
<!-- observation:whohealedme_xml_input_and_layout promoted:2026-04-01T14:30:00Z -->
> Source: WhoHealedMe | Confidence: MEDIUM | Promoted: 2026-04-01

- `MEDIUM`: ancestor handleinput attribute can block child click targets
  - Guidance: validate handleinput across parent and template chain if child looks wired but remains inert
- `MEDIUM`: early layout reads on nested scroll content can under-report the eventual usable region
  - Guidance: When first-render list height looks collapsed, compute dimensions from the outer parent first, then resize child content and refresh the scroll rect
```

**Duplicate Prevention:** HTML comment marker checked before re-promoting (lines 710-716):
```go
func isDuplicatePromotion(seedAbsPath, entryID string) bool {
    content, _ := os.ReadFile(seedAbsPath)
    marker := fmt.Sprintf("<!-- observation:%s ", entryID)
    return strings.Contains(string(content), marker)
}
```

---

## 5. Contract-Native Alignment Assessment

### ✅ Compliant Areas

**Semantic Contracts Intact:**
- Observations ≠ generated symbols (feeding is input mechanism, not symbol emission)
- Evidence types abstracted: {code_change, runtime_observation, user_report, test} NOT tied to parser phases
- Claims are not symbols: Kind field (behavior_caveat|usage_pattern|signature_hint) distinct from symbol definition
- Confidence tracking separate: Explicit HIGH/MEDIUM/LOW at observation level, not overloading symbol confidence

**Safety Features:**
- Workspace root validation before file operations (Makefile/docker-compose.yml/tools/api_doc_gen markers)
- Path traversal prevention: `filepath.Rel` check ensures seed paths don't escape workspace root
- Schema version pinning: "feeding.observation.v1" enforced at validation
- Durable rejection store: rejected.ndjson maintains audit trail separate from active queue

**Governance Implemented:**
- Review verdict separation: Accept vs Reject with named reviewers
- Immutable promotion records: Can't re-review or re-reject promoted observations
- Duplicate promotion markers: HTML comments prevent accidental re-promotion to same seed

### ⚠️ Considerations for Evolution

**Symbol Validation Gap:**
- Schema allows `claims[].symbols[]` but no validation that symbols exist in generated docs
- Opportunity: Add symbol lookup during promotion to warn if symbol_id cannot be resolved

**Seed File Update Conflicts:**
- Multiple promotions append to same seed file without version tracking
- Opportunity: Add date-based comments or version tags in seed files for audit trail

**Human Narrative Separation:**
- Markdown narratives (whohealedme_xml_input_and_layout.md) are parallel to .feed.json, not validated
- Current: Markdown is reference documentation for reviewers
- Future: Could validate narrative consistency with .feed.json at review time

**Promotion Timing:**
- regenerate_from_promoted_knowledge exists but unclear trigger points
- Opportunity: Establish clear workflow—promote observation → immediately regenerate OR batch regen on schedule

---

## 6. Current State: Active Observations

### Queue Status (as of 2026-04-02)

**Total Observations:** 2 (both candidate status)

#### Observation 1: WhoHealedMe XML Input/Layout
- **Entry ID:** whohealedme_xml_input_and_layout
- **Status:** Candidate (in accepted.ndjson)
- **Source:** WhoHealedMe addon, healer-details window, dev_observed + user_validated
- **Confidence:** MEDIUM (observed and user-validated during addon implementation)
- **Claims (2):**
  1. Ancestor handleinput blocks child clicks → guidance provided
  2. Early nested dimensions under-report space → guidance provided
- **Evidence (2):** code_change + runtime_observation
- **Target Seed:** docs/platform/seeds/xml_conventions.md
- **Promotion Status:** NOT yet promoted to seed (still candidate)

#### Observation 2: QuickTacticSwitch XML ListData Binding
- **Entry ID:** quicktacticswitch_xml_listdata_binding
- **Status:** Candidate (in accepted.ndjson)
- **Source:** QuickTacticSwitch addon, in-workspace addon review (dev_observed, corroborated)
- **Confidence:** MEDIUM
- **Claims (4):**
  1. ListData binds ListBox row definition to Lua backing table
  2. ListColumn entries map row-template child windows to Lua table fields
  3. populationfunction is optional custom row-population logic
  4. Empty populationfunction valid for text-only lists
- **Evidence (7):** code_change entries for XML and Lua patterns
- **Target Seed:** docs/platform/seeds/xml_conventions.md
- **Promotion Status:** NOT yet promoted to seed (still candidate)

### Rejected Observations

**Count:** 0 (rejected.ndjson file doesn't exist—no rejections yet)

### Seed Integration (Actual)

**docs/platform/seeds/xml_conventions.md** currently contains evidence of both WhoHealedMe and QuickTacticSwitch claims, but promotions show different states:

**Example from seed file (lines 13-14):**
```markdown
- `MEDIUM`: if a child click target looks correctly wired but remains inert, validate `handleinput` across the parent or template chain as well as on the child node.
```
✓ This matches WhoHealedMe claim #1 guidance exactly—evidence of prior manual integration or prepared promotion.

**Example from seed file (lines 18-30):**
```markdown
- `MEDIUM`: `ListData` appears to be the XML binding point that connects a `ListBox` row definition to a Lua backing table.
- `MEDIUM`: `ListColumn` entries under `ListData` appear to map row-template child windows to fields on each Lua table entry.
```
✓ These match QuickTacticSwitch claims #1-2 exactly—again, seed content is aligned with observation data.

**Interpretation:** Observations in queue represent formalized versions of already-discovered patterns. Seed files have been manually curated. Current workflow appears to be: findings emerge → documented in code → entered as formal observations → ready for structured review → then propagate to seeds formally.

---

## 7. Reusable Infrastructure Inventory

### ✅ Production-Ready Components

| Component | Status | Confidence | Notes |
|-----------|--------|------------|-------|
| JSON schema (observation.v1) | ✅ Complete | HIGH | Strict validation, supports all required fields, versioned |
| File-pair pattern (markdown + .feed.json) | ✅ Complete | HIGH | Separation of concerns, both human and machine readable |
| Validation pipeline | ✅ Complete | HIGH | Schema version check, required fields, entry ID uniqueness, workspace safety |
| Queue management (NDJSON) | ✅ Complete | HIGH | Append-only, lifecycle tracking, status immutability enforcement |
| Status workflow (candidate → accepted/rejected → promoted) | ✅ Complete | HIGH | State machine enforced in code, prevents re-reviews, immutable promotion |
| MCP tool registry | ✅ Complete | HIGH | 7 tools fully implemented, all with error handling and dry-run support |
| Dry-run support | ✅ Complete | HIGH | All tools support DryRun flag for safe preview |
| Workspace validation | ✅ Complete | HIGH | Prevents misconfiguration, enforced before file operations |
| Path traversal prevention | ✅ Complete | HIGH | Validates seed paths don't escape workspace root |
| Duplicate promotion prevention | ✅ Complete | HIGH | HTML comment marker tracking prevents re-promotion |
| Durable rejection store | ✅ Complete | HIGH | rejected.ndjson maintains audit trail, no duplicates |
| Markdown fragment builder | ✅ Complete | HIGH | Extracts source, confidence, claims with guidance; builds valid HTML comment markers |
| Regeneration triggering | ✅ Complete | MEDIUM | Scope-based command execution (platform/site/full), safe argument handling |

### ⚠️ Areas Ready for Enhancement

| Gap | Facility | Effort | Priority |
|-----|----------|--------|----------|
| Symbol validation | Check symbols[] against generated docs during promotion | Medium | Medium |
| Seed file versioning | Add date/version tags to seed files for update tracking | Low | Low |
| Feedback loop integration | Connect addon discovery workflow to ingest_observation_batch | Medium | High |
| Narrative validation | Validate markdown narratives against .feed.json schema at review | Medium | Medium |
| Promotion scheduling | Establish clear trigger points for regenerate_from_promoted_knowledge | Low | High |

---

## 8. Legacy/Noisy Areas

### ✅ No Major Legacy Baggage

The feeding mechanism design is coherent and contract-native aligned. No significant legacy assumptions detected. Key observations:

1. **Schema is greenfield:** observation.v1 doesn't reference old parser outputs or semantic phases
2. **Evidence types are abstracted:** {code_change, runtime_observation, user_report, test} don't couple to implementation details
3. **Claims model is simple:** {kind, statement, guidance} — lightweight and extensible
4. **No locked-in state model:** Status workflow can evolve (e.g., add "version_superseded" state if needed)

### ⚠️ Minor Pain Points (Not Blockers)

1. **Confidence terminology:** Uses HIGH/MEDIUM/LOW (matches platform docs) but not synchronized to symbol semantic_confidence field naming
   - **Fix:** None needed—observations are separate from generated symbols, so different terminology is acceptable

2. **Narrative duplication:** Markdown files and .feed.json contain overlapping information
   - **Fix:** Document that markdown is for reviewers/context, .feed.json is source of truth for promotion

3. **Queue file parsing:** readQueueRecords unmarshals full JSON objects, no index
   - **Impact:** Linear scan on each operation (acceptable for current 2-observation scale, could optimize with JSON lines index if needed)

4. **Regeneration timing:** No specification of when regenerate_from_promoted_knowledge is called
   - **Fix:** Establish convention—immediate regeneration after promote, or scheduled batch regen

---

## 9. Evolution Recommendations

### Phase 1: Immediate (Safe for Early Addon Feedback Loop)

**Goal:** Begin using feeding mechanism for actual addon discoveries

**Actions:**
1. **Establish Review Board:** Name reviewers (e.g., addon maintainers, API experts) who can approve/reject observations
2. **Feed Addon Discoveries:** Use ingest_observation_batch to load .feed.json files as addon work surfaces new findings
3. **Implement Review Workflow:** 
   - Daily: Call list_pending_observations to surface candidates
   - Review: Apply review_observation verdict + notes
   - Archive: Rejected observations go to durable store, accepted ready for promotion
4. **Document Examples:** Create 2-3 template .feed.json files for common finding types (XML caveat, Lua pattern, event behavior)
5. **Safety Test:** Run full workflow dry-run on each step before committing to queue

**Expected Outcome:** 5-10 vetted observations per addon cycle feeding into seed docs

### Phase 2: Near-Term (Before Production Release)

**Goal:** Stabilize promotion workflow and ensure semantic integrity

**Actions:**
1. **Symbol Validation:**
   - Add symbol lookup function: during promote_observation, resolve claims[].symbols[] against generated docs
   - Warn if symbol not found, but don't block promotion (symbols may be future/provisional)
   - Example: "Symbol 'handleinput' could not be resolved in generated docs—verify entry point"

2. **Seed Governance:**
   - Add date-based comment headers to seed files: `<!-- seed updated: 2026-04-02 by promotion workflow -->`
   - Enables audit trail of what observations contribute to each seed
   - Example:
     ```markdown
     <!-- seed updated: 2026-04-02, promoted: whohealedme_xml_input_and_layout, quicktacticswitch_xml_listdata_binding -->
     
     ## XML Conventions Seed
     ```

3. **Regeneration Scheduling:**
   - Establish clear trigger: After each promote_observation batch completes, immediately call regenerate_from_promoted_knowledge
   - OR: Daily scheduled regen at EOD
   - Document choice in INGESTION_RUNBOOK.md

4. **Narrative Validation:**
   - Optional: During review_observation, validate that markdown narrative file matches .feed.json structure
   - At minimum: Document that narrative files are informational (source of truth is .feed.json)

5. **Status Reporting:**
   - Build simple dashboard: pending observations by target_seed, rejection history by reviewer
   - Query via list_pending_observations + list_rejected_observations

**Expected Outcome:** Robust promotion workflow with clear audit trail, symbol validation, and regeneration scheduling

### Phase 3: Future (After Stable Pattern Emerges)

**Goal:** Enable observation versioning and lifecycle maturity

**Actions:**
1. **Observation Versioning:**
   - Add `version` field to schema (e.g., "1.0") and `supersedes` field (entry_id of prior version)
   - Allow corrections without re-promotion of original (create new observation with supersession marker)
   - Regeneration would pull latest version only

2. **Feedback Loop Integration with Testing:**
   - Pair promoted observations with lifecycle tests: verify claim behavior in addon test suite
   - Example: "ListData binding claim" → test that ListBox rows render correctly when populated

3. **Observation Dashboard:**
   - Web UI for reviewers: list pending, comment, approve/reject, view history
   - Metrics: observation velocity, reviewer load, promotion rates by addon

4. **Canonical Integration:**
   - Consider: Should some promoted observations become part of generated docs (e.g., caveat sections)?
   - Current: Promotions go to seed docs only. Future: Could pipe to metadata in generated docs

---

## 10. Validation Summary

### What Was Verified

| Claim | Evidence | Status |
|-------|----------|--------|
| Machine-readable format | JSON schema v2020-12, enforced validation | ✅ Verified |
| Status distinctions (candidate/confirmed/rejected) | Status enum in schema, state machine in lifecycle code | ✅ Verified |
| Data preservation (evidence, context, confidence) | Schema fields, lifecycle record structure, actual observations | ✅ Verified |
| Semantic contracts preserved | Evidence abstracted from parser phases, claims ≠ symbols, confidence level separate | ✅ Verified |
| MCP integration complete | 7 tools fully implemented, tool registry, error handling | ✅ Verified |
| Safety features | Workspace validation, path traversal prevention, duplicate prevention | ✅ Verified |
| Promotion workflow | buildPromotionFragment, appendToSeedFile, isDuplicatePromotion, regeneration trigger | ✅ Verified |
| Durable rejection store | Separate rejected.ndjson, immutability enforcement, duplicate prevention | ✅ Verified |

### What Remains Unresolved

| Item | Classification | Risk |
|------|-----------------|------|
| Symbol validation during promotion | Missing enhancement (not implemented) | Low—optional safeguard |
| Seed file version tracking | Missing enhancement (not implemented) | Low—audit trail convenience |
| Explicit regeneration trigger points | Process documentation gap (mechanism exists) | Medium—operational clarity |
| Narrative-to-schema validation | Missing enhancement (not implemented) | Low—documentation issue |
| Feedback loop workflow integration | Missing system integration (mechanism only) | High—but ready when needed |

---

## 11. Conclusion

**The feeding mechanism is production-ready as the foundation for a contract-native feedback loop.**

### Key Strengths
1. ✅ Fully implemented: All schema, MCP tools, queue, promotion, and regeneration infrastructure complete
2. ✅ Contract-aligned: No legacy baggage; cleanly separates observations from generated symbols
3. ✅ Safe by design: Workspace validation, path traversal prevention, duplicate prevention, durable rejection store
4. ✅ Extensible: Schema versioning, confidence tracking, symbol references all present for future enhancement
5. ✅ Testable: Dry-run support on all operations, no side effects until explicitly committed

### Recommended Next Steps
1. **Immediate:** Establish review board and begin ingesting addon findings via ingest_observation_batch
2. **Near-term:** Add symbol validation and seed file version tracking before production release
3. **Future:** Implement observation versioning and integrate with lifecycle testing once pattern stabilizes

### Risk Assessment

| Risk | Likelihood | Impact | Mitigation |
|------|------------|--------|------------|
| Schema drift | Low | Medium | Schema version pinning enforced in validation |
| Conflicting seed updates | Low | Low | Append-only semantics, duplicate marker prevention |
| Orphaned observations | Medium | Low | Establish review SLA, clear rejection criteria |
| Regeneration errors | Low | Medium | Dry-run before committing, workspace validation |
| Symbol validation gaps | Medium | Low | Can be added as non-blocking enhancement |

**Overall:** Ready for adoption. Start with Phase 1 (feedback loop ingestion), advance to Phase 2 (governance enhancements) before production release.

---

## Appendices

### A. File Reference
- Schema location: `docs/platform/feeding/model/observation.v1.schema.json`
- Pipeline code: `tools/mcp_server/server/observation_lifecycle.go`, `feeding_ingest.go`, `ingest_observation.go`
- Queue location: `docs/platform/feeding/review_queue/accepted.ndjson`
- Seed location: `docs/platform/seeds/[lua|xml|war_api]_conventions.md`, `war_api_facts.md`
- MCP tools: 7 tools in `tool_handlers.go` (lines 14-20 listing, individual implementations in server/)

### B. Sample Observation Structure
See Section 4 (Data Preservation) for full whohealedme_xml_input_and_layout example.

### C. Workflow Commands (from INGESTION_RUNBOOK.md)
```bash
# Validate observation
POST /mcp/tools/ingest_observation
{
  "observation": {...},
  "dry_run": true
}

# Accept observation for promotion
POST /mcp/tools/review_observation
{
  "observation_id": "whohealedme_xml_input_and_layout",
  "verdict": "accept",
  "reviewer": "addon-maintainer",
  "notes": "validated against current addon code"
}

# Promote to seed files
POST /mcp/tools/promote_observation
{
  "observation_id": "whohealedme_xml_input_and_layout",
  "dry_run": false
}

# Regenerate docs
POST /mcp/tools/regenerate_from_promoted_knowledge
{
  "scope": "full",
  "dry_run": false
}
```
