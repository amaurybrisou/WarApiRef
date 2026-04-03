# Feeding Mechanism Quick Reference

**For:** Addon developers, maintainers, and API teams  
**Purpose:** Quick checklist and command reference  
**See Also:** [`INGESTION_RUNBOOK.md`](../docs/platform/feeding/INGESTION_RUNBOOK.md) for detailed examples

---

## Key Points (Review Before Starting)

✅ **Observations are advisory input, not truth** — They inform seed guidance, which informs docs  
✅ **Candidates don't affect docs** — Only after review+promotion+regeneration  
✅ **Symbol warnings don't block** — Symbols may be forward-references or future  
✅ **Clear governance gates** — Ingest → Review → Promote → Regenerate  
✅ **Seed files are traceable** — HTML comments show which observations were promoted  

---

## Observation File Template

### File Pair Structure

| File | Purpose | Example |
|------|---------|---------|
| `.md` | Human narrative (for reviewers) | `docs/platform/feeding/xml/myAddon_pattern.md` |
| `.feed.json` | Machine-readable JSON (source of truth) | `docs/platform/feeding/xml/myAddon_pattern.feed.json` |

### Minimal `.feed.json` Template

```json
{
  "schema_version": "feeding.observation.v1",
  "entry_id": "myAddon_xml_pattern",
  "status": "candidate",
  "source": {
    "addon": "MyAddon",
    "date_utc": "2026-04-03T10:00:00Z",
    "validation": "dev_observed",
    "context": "Brief context of discovery"
  },
  "target_seeds": ["docs/platform/seeds/xml_conventions.md"],
  "confidence": {
    "overall": "MEDIUM",
    "rationale": "Observed in own addon implementation"
  },
  "claims": [
    {
      "claim_id": "claim_1",
      "kind": "behavior_caveat",
      "confidence": "MEDIUM",
      "statement": "Clear claim statement here",
      "guidance": "How addon developers should apply this",
      "symbols": ["SymbolName"],
      "evidence_refs": ["evidence_1"]
    }
  ],
  "evidence": [
    {
      "evidence_id": "evidence_1",
      "type": "code_change",
      "summary": "Description of the evidence",
      "file_refs": ["addons/MyAddon/file.lua:42"]
    }
  ],
  "promotion": {
    "notes": "Notes on when/why to promote",
    "criteria": ["Meets criteria for seed inclusion"]
  }
}
```

### Confidence Levels

| Level | Criteria |
|-------|----------|
| `HIGH` | Works reliably; cross-validated; documented in canonical docs |
| `MEDIUM` | Observed in own addon; reasonable theory; user-confirmed |
| `LOW` | Theory only; needs more testing; provisional |

### Claim Kinds

| Kind | Usage |
|------|-------|
| `behavior_caveat` | Unexpected behavior, quirk, or workaround needed |
| `usage_pattern` | Idiom that works well; recommended approach |
| `signature_hint` | Symbol exists but undocumented; usage examples |

### Evidence Types

| Type | Usage |
|------|-------|
| `code_change` | File+line showing workaround or pattern |
| `runtime_observation` | Runtime behavior (e.g., memory, timing) |
| `user_report` | End-user testing or feedback |
| `test` | Test case that demonstrates behavior |

---

## Developer Workflow

### Step 1: Create Observation

```powershell
# Create directory if needed
New-Item -ItemType Directory -Path "docs\platform\feeding\xml" -Force

# Create narrative file (human-readable)
@"
## Status
Candidate (awaiting review)

## Source
- Addon: MyAddon
- Discovered: 2026-04-03
- Validation: dev_observed

## Target Seed
docs/platform/seeds/xml_conventions.md

## Confidence
MEDIUM — Observed in own addon implementation

## Observed Behavior
1. XML form input loses focus after parent window update

## Evidence
MyAddon/forms.lua:142 — State preservation pattern

## Promotion Notes
Review and validate against current codebase
"@ | Out-File "docs\platform\feeding\xml\myAddon_form_input.md" -Encoding UTF8

# Create JSON file (machine-readable, see template above)
# Copy template and customize...
```

### Step 2: Validate Schema (Dry-Run)

```powershell
# Start MCP server
docker compose up --build -d mcp-server
# OR: go run .\tools\mcp_server\main.go --transport http --listen :8091 --docs-root docs\war-api --feeding-root docs\platform\feeding

# Validate batch (dry-run, no persistence)
$payload = @{
  jsonrpc = "2.0"
  method = "feeding/ingest_batch"
  params = @{ dry_run = $true; continue_on_error = $true }
} | ConvertTo-Json -Depth 8

$result = Invoke-RestMethod -Method Post -Uri "http://127.0.0.1:8091/mcp" \
  -ContentType "application/json" -Body $payload

$result.result | ConvertTo-Json -Depth 5
```

### Step 3: Fix Issues and Persist

```powershell
# After fixing any validation errors, persist to queue
$payload = @{
  jsonrpc = "2.0"
  method = "feeding/ingest_batch"
  params = @{ 
    dry_run = $false
    persist = $true
    continue_on_error = $true
  }
} | ConvertTo-Json -Depth 8

$result = Invoke-RestMethod -Method Post -Uri "http://127.0.0.1:8091/mcp" \
  -ContentType "application/json" -Body $payload

if ($result.result.accepted_count -gt 0) {
  Write-Host "✅ Ingested $($result.result.accepted_count) observations"
}
```

### Step 4: Wait for Review

```
Your observation is now in the queue as 'candidate' status.
Maintainers will review and decide to accept or reject.
```

---

## Maintainer Workflow

### Step 1: List Pending Observations

```powershell
$payload = @{
  jsonrpc = "2.0"
  method = "feeding/list_pending_observations"
  params = @{
    status_filter = @("candidate")
    target_filter = "xml_conventions"
    limit = 10
  }
} | ConvertTo-Json -Depth 8

$result = Invoke-RestMethod -Method Post -Uri "http://127.0.0.1:8091/mcp" \
  -ContentType "application/json" -Body $payload

$result.result.observations | Format-Table -AutoSize
```

**Response includes:**
- ObservationID, Status, Confidence
- SourceAddon, EvidenceCount, ClaimSummary
- TargetSeeds

### Step 2: Review and Decide

```powershell
# ACCEPT the observation
$payload = @{
  jsonrpc = "2.0"
  method = "feeding/review_observation"
  params = @{
    observation_id = "myAddon_form_input"
    verdict = "accept"  # or "reject"
    reviewer = "you@example.com"
    notes = "Validated against current addon code; valuable pattern"
  }
} | ConvertTo-Json -Depth 8

$result = Invoke-RestMethod -Method Post -Uri "http://127.0.0.1:8091/mcp" \
  -ContentType "application/json" -Body $payload

Write-Host "Status: $($result.result.status)"
```

**Verdict options:**
- `verdict = "accept"` → Status becomes `accepted` (eligible for promotion)
- `verdict = "reject"` → Status becomes `rejected` (archived, immutable)

### Step 3: Promote to Seed File

```powershell
# DRY-RUN FIRST
$payload = @{
  jsonrpc = "2.0"
  method = "feeding/promote"
  params = @{
    observation_id = "myAddon_form_input"
    dry_run = $true  # Preview first
  }
} | ConvertTo-Json -Depth 8

$result = Invoke-RestMethod -Method Post -Uri "http://127.0.0.1:8091/mcp" \
  -ContentType "application/json" -Body $payload

# Check warnings
$result.result.warnings | ForEach-Object { Write-Host "⚠️  $($_.message)" }

# Then proceed (dry_run=false)
$payload.params.dry_run = $false
$result = Invoke-RestMethod -Method Post -Uri "http://127.0.0.1:8091/mcp" \
  -ContentType "application/json" -Body $payload

if ($result.result.promoted) {
  Write-Host "✅ Promoted to seed file(s)"
}
```

**During promotion:**
- ✅ Symbol validation runs (issues non-blocking warnings)
- ✅ Duplicate prevention check (prevents re-promotion)
- ✅ Seed file updated with metadata comment
- ✅ Observation marked as `promoted` (immutable)

### Step 4: Regenerate Docs Immediately

```powershell
# DRY-RUN FIRST
$payload = @{
  jsonrpc = "2.0"
  method = "feeding/regenerate"
  params = @{
    scope = "full"  # or "platform" or "site"
    dry_run = $true
  }
} | ConvertTo-Json -Depth 8

$result = Invoke-RestMethod -Method Post -Uri "http://127.0.0.1:8091/mcp" \
  -ContentType "application/json" -Body $payload

$result.result.steps | Format-Table -AutoSize

# Then proceed (dry_run=false)
$payload.params.dry_run = $false
$result = Invoke-RestMethod -Method Post -Uri "http://127.0.0.1:8091/mcp" \
  -ContentType "application/json" -Body $payload

if ($result.result.success) {
  Write-Host "✅ Docs regenerated successfully"
}
```

**Regeneration scopes:**
- `scope = "full"` — Rebuild platform + site (recommended)
- `scope = "platform"` — Rebuild docs/war-api/ only
- `scope = "site"` — Rebuild docs/site/content/ only

### Step 5: Verify in Generated Docs

```bash
# Check that observation content is in generated docs
grep -r "Form input loses focus" docs/war-api/

# Verify seed file shows observation metadata
grep "<!-- OBSERVATION:" docs/platform/seeds/xml_conventions.md

# Check file timestamps (should match regeneration time)
ls -la docs/war-api/conventions*.md
```

---

## Common Issues and Fixes

### "Symbol 'X' could not be resolved"

**What it means:**
- Symbol referenced in observation not found in generated docs
- This is a **warning**, not an error
- Promotion proceeds anyway

**Action:**
1. Verify symbol name spelling
2. Check if symbol should exist: `grep -r "SymbolName" docs/war-api/`
3. If symbol missing, consider updating canonical docs first
4. If intentional forward-reference, promotion is OK

**Don't:**
- ❌ Don't assume observation is invalid
- ❌ Don't hold up promotion if symbol is intentional forward-ref

### "Observation already promoted"

**What it means:**
- Trying to promote an observation marked as `promoted`
- Promoted observations are immutable

**Action:**
1. Create new observation with updated information
2. Use different entry_id (e.g., add `_v2` suffix)
3. Promote new observation
4. Optionally mark old one as deprecated in comments

### "Candidate observation not showing in queries"

**What it means:**
- Candidate observations are isolated by design
- They don't show in `search_symbols`, `lookup_symbol`, etc. until promoted

**Action:**
1. Use `feeding/list_pending_observations` to see candidates
2. Review and promote (status→accepted→promoted)
3. After regeneration, observation appears in standard queries

---

## Batch Promotion Workflow

For efficiency when promoting multiple observations:

```powershell
# 1. List multiple pending observations
$pending = Invoke-RestMethod -Method Post -Uri "http://127.0.0.1:8091/mcp" \
  -Body (@{ method = "feeding/list_pending_observations"; params = @{ limit = 10 } } | ConvertTo-Json) \
  -ContentType "application/json"

# 2. Accept all (or select subset)
$pending.result.observations | ForEach-Object {
  $payload = @{
    method = "feeding/review"
    params = @{
      observation_id = $_.observation_id
      verdict = "accept"
      reviewer = "you@example.com"
      notes = "Batch validation"
    }
  } | ConvertTo-Json
  
  Invoke-RestMethod -Method Post -Uri "http://127.0.0.1:8091/mcp" \
    -Body $payload -ContentType "application/json" | Out-Null
}

# 3. Promote all to same seed file
foreach ($id in @("obs1", "obs2", "obs3")) {
  $payload = @{
    method = "feeding/promote"
    params = @{ observation_id = $id; dry_run = $false }
  } | ConvertTo-Json
  
  Invoke-RestMethod -Method Post -Uri "http://127.0.0.1:8091/mcp" \
    -Body $payload -ContentType "application/json" | Out-Null
}

# 4. Regenerate ONCE (much faster than per-observation)
$payload = @{
  method = "feeding/regenerate"
  params = @{ scope = "full"; dry_run = $false }
} | ConvertTo-Json

Invoke-RestMethod -Method Post -Uri "http://127.0.0.1:8091/mcp" \
  -Body $payload -ContentType "application/json"

Write-Host "✅ Batch promotion and regeneration complete"
```

---

## Seed File Anatomy

After promotion, seed files contain:

```markdown
# Existing seed content...

<!-- OBSERVATION:myAddon_form_input (promoted:2026-04-03T14:30:00Z) -->
> Source: MyAddon | Confidence: MEDIUM | Promoted: 2026-04-03

- `MEDIUM`: Form input loses focus after parent window update
  - Guidance: Save form state before parent updates; restore after render
- `MEDIUM`: Another claim from same observation
  - Guidance: How to apply this claim
```

**Parts:**
1. **HTML Comment:** `<!-- OBSERVATION:entry_id (promoted:timestamp) -->`
   - For audit trail and duplicate prevention

2. **Source Attribution:**
   - Addon name, confidence, promotion date
   - Helps maintainers track origins

3. **Claims with Guidance:**
   - Extracted from observation
   - Confidence level labeled
   - Actionable guidance provided

---

## Troubleshooting Commands

```powershell
# View queue contents
Get-Content docs\platform\feeding\review_queue\accepted.ndjson | ConvertFrom-Json | Format-List

# Count observations by status
$queue = Get-Content docs\platform\feeding\review_queue\accepted.ndjson | ConvertFrom-Json
$queue | Group-Object { $_.lifecycle_status } | Select-Object Name, Count

# View rejected observations
Get-Content docs\platform\feeding\review_queue\rejected.ndjson | ConvertFrom-Json | Format-List

# Find observations in seed files
Select-String -Path docs\platform\seeds\*.md -Pattern "<!-- OBSERVATION:"

# Verify seed file syntax (before regeneration)
Test-Path docs\platform\seeds\xml_conventions.md
Get-Content docs\platform\seeds\xml_conventions.md | Measure-Object -Line
```

---

## Quick Links

| Document | Purpose |
|----------|---------|
| [INGESTION_RUNBOOK.md](../docs/platform/feeding/INGESTION_RUNBOOK.md) | Detailed workflow with examples |
| [FEEDING_AND_SEMANTIC_CONTRACTS.md](../agent-docs/FEEDING_AND_SEMANTIC_CONTRACTS.md) | Architecture and governance deep-dive |
| [observation.v1.schema.json](../docs/platform/feeding/model/observation.v1.schema.json) | Complete JSON schema reference |
| [FEEDING_MECHANISM_AUDIT.md](../agent-docs/FEEDING_MECHANISM_AUDIT.md) | Audit and implementation details |

---

**Last Updated:** April 3, 2026  
**Maintained By:** API Team  
**Status:** Ready for Production
