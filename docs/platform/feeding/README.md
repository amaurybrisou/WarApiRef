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