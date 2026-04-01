---
applyTo: '**/*.{lua,xml,mod}'
---

# WAR API Policy (MCP-Oriented)

This file is policy only. It is not an API catalog.

## Source of Truth

- Treat API facts as authoritative only when they come from:
  - `docs/platform/`
  - `docs/platform/seeds/`
  - MCP server responses backed by generated WAR docs
- Prefer MCP/resource lookups before making assumptions.

## Strict Non-Invention Rule

- Do not invent or guess:
  - API names
  - signatures
  - argument meanings
  - return shapes/types
  - XML behavior/event semantics
- If a fact is not confirmed by sources of truth, mark it as uncertain.

## Uncertainty and Confidence

- Preserve uncertainty explicitly in code comments or explanation text when needed.
- Carry confidence context forward:
  - high confidence: canonical and repeated evidence
  - medium confidence: partial or indirect evidence
  - low confidence: weak or single-source evidence
- Never present inferred behavior as guaranteed behavior.

## Knowledge Improvement Loop

- If implementation work reveals better or corrected API knowledge:
  - propose updates to `docs/platform/seeds/`
  - propose updates to canonical docs inputs/transformers when applicable
  - keep instruction files thin and unchanged unless policy itself changes
