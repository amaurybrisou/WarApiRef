# PROJECT INSTRUCTIONS

## Goal

Maintain a reliable multi-phase analysis pipeline for:
- XML parsing
- Lua analysis
- semantic enrichment
- catalog / graph generation

Priorities:
1. correctness
2. determinism
3. traceability
4. minimal safe changes

---

## Core Rules

- Do not guess missing intent or semantics
- Do not invent relationships, edges, or correlations
- Prefer explicit evidence over heuristics
- Keep behavior unchanged unless the task explicitly requires otherwise
- Make the smallest safe change possible
- Do not refactor unrelated parts while fixing a targeted issue

---

## Pipeline Discipline

When working on analysis or graph issues, think in phases:

1. Phase 1 contract input: `xml-tree` (`phase1-tree/v1`)
2. Phase 2 contract input: `lua-analysis` (`phase2-lua/v1`)
3. Phase 3 contract input: `xml-lua-links` (`phase3-xml-lua-link/v1`)
4. semantic enrichment
5. catalog / graph emission

Contract-only rule:

- Semantic input is contract artifacts only (`xml-tree`, `lua-analysis`, `xml-lua-links`).
- Missing/invalid contracts are hard failures.
- Degraded/fallback semantic modes are not allowed.
- Markdown under `docs/addon-api`, `docs/war-api`, and site docs is presentation output only and must never be used as semantic evidence.

For any bug, identify:
- what evidence exists
- in which phase it appears
- in which phase it is lost, ignored, or mis-modeled

Do not patch blindly at the graph layer if the evidence is missing upstream.

---

## Cleanup Rules

Allowed:
- remove confirmed dead code
- remove unused imports
- normalize names only through existing canonicalization paths
- deduplicate only when equivalence is proven
- reduce noisy logs

Not allowed:
- speculative cleanup
- architecture rewrites during a bugfix
- changing emitted semantics without proof
- deleting code that is not confirmed unused

---

## Naming and Canonicalization

- Use one canonical lookup strategy for keys
- Normalize at ingestion / accumulator boundaries
- Do not change display labels unless explicitly required
- Preserve public/output naming when possible

Canonicalization must prevent duplicate internal symbols, not silently alter published meaning.

---

## Graph Rules

- Structural nodes and behavioral nodes must stay distinct
- XML element types are not XML events
- XML events are not Lua functions
- Only emit edges backed by existing pipeline evidence
- Prefer no edge over a weak or overstated edge

When a node is dangling, classify it before patching:
- no current evidence
- missing emission rule
- missing XML ↔ Lua correlation
- model / vocabulary limitation

---

## Validation

After meaningful changes:
- validate edited files
- run the relevant generator / pipeline when possible
- regenerate artifacts when graph/catalog logic changed
- compare before/after results
- report unresolved cases explicitly

If full runtime validation is unavailable, say so clearly.

---

## Output Expectations

For non-trivial changes, report briefly:
1. files changed
2. what was removed / normalized / fixed
3. what was validated
4. what remains unresolved