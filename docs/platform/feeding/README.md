# Platform Feeding Documents

This folder is the intake area for implementation-validated findings that are not yet promoted into reusable seed docs.

## Purpose

- Capture addon-driven platform observations with explicit provenance.
- Keep candidate knowledge separate from normalized seed guidance until reviewed.
- Make later promotion into `docs/platform/seeds/` straightforward and auditable.

## Structure

- Group findings by topic folder when practical:
  - `xml/`
  - `lua/`
  - `events/`
  - `window_api/`
- Use one file per coherent observation set.
- Prefer filenames that include source addon and topic, for example:
  - `whohealedme_xml_input_and_layout.md`

## Required Sections

Each feeding document should include:

1. `Status`
2. `Source`
3. `Target Seed`
4. `Confidence`
5. `Observed Behavior`
6. `Evidence`
7. `Promotion Notes`

## Promotion Rule
- Only move content from this folder into `docs/platform/seeds/` after review confirms the note is reusable, narrowly phrased, and confidence-tagged correctly.

---

## Key Principles (Strengthened Mechanism)

✅ **Observations are reviewed knowledge input only**, never semantic sources of truth

✅ **Candidate observations are isolated** — they don't trigger regeneration or affect MCP queries until explicitly promoted

✅ **Symbol validation during promotion is advisory** — observations reference symbols informally; validation warns but doesn't block

✅ **Seed files are traceability-enabled** — HTML comments track which observations were promoted

✅ **Clear regeneration policy** — promote observation → immediately regenerate (or scheduled batch)

For detailed explanation, see:
- [`FEEDING_AND_SEMANTIC_CONTRACTS.md`](../../agent-docs/FEEDING_AND_SEMANTIC_CONTRACTS.md) — Architecture and principles
- [`FEEDING_MECHANISM_AUDIT.md`](../../agent-docs/FEEDING_MECHANISM_AUDIT.md) — Implementation details and validation

---

## Semantic Contracts Are Immutable

This feeding system does NOT affect semantic sources of truth:
- **xml-tree** — XML structural analysis (parser-driven)
- **lua-analysis** — Lua semantic analysis (parser-driven)  
- **xml-lua-links** — Cross-validation rules (generator-driven)

Architecture:
```
Observation (reviewed input)
  ↓
Seed File (formalized guidance)
  ↓
Generated Docs (canonical, authoritative)
```

Never expect observations to bypass or override semantic contracts.

- Only move content from this folder into `docs/platform/seeds/` after review confirms the note is reusable, narrowly phrased, and confidence-tagged correctly.