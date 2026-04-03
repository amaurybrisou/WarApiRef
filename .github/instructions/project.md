# CLEAN PROJECT — STRICT INSTRUCTIONS

## CONTEXT

This repository implements a multi-phase analysis pipeline:
- XML parsing
- Lua analysis
- Semantic enrichment
- Catalog generation

The current codebase contains inconsistencies, dead code, and partial structures.
This task is a **strict cleanup phase** before any further feature work.

---

## GLOBAL RULES

- DO NOT add new features
- DO NOT change business logic
- DO NOT guess missing intent
- KEEP behavior strictly identical
- PRIORITIZE consistency and determinism

---

## TASKS

### 1. REMOVE DEAD CODE

- Delete:
  - unused functions
  - unused structs/types
  - unused files
- Remove:
  - commented-out code blocks
  - legacy parsing logic not used in `RunPipeline`

---

### 2. NORMALIZE NAMING

Apply a **single naming strategy** across the entire project:

- XML entities
- Lua functions
- EnrichedElement
- Catalog structures

Rules:
- consistent casing (choose one: camelCase OR PascalCase OR snake_case)
- no duplicated semantic variants
- no implicit naming differences

---

### 3. REMOVE DUPLICATES

Ensure **one canonical representation** for:

- functions
- events
- XML elements

Tasks:
- merge duplicates
- remove redundant entries
- align references

---

### 4. FIX IMPORTS

- remove unused imports
- group imports (std / external / internal)
- ensure deterministic builds

---

### 5. PACKAGE STRUCTURE

Enforce clear separation:

- `/xml` → XML parsing only
- `/lua` → Lua analysis only
- `/semantic` → enrichment logic
- `/catalog` → final structures

Tasks:
- move misplaced files
- remove circular dependencies
- ensure clean boundaries

---

### 6. LOGGING

- remove noisy logs
- keep only:
  - structured debug logs
  - pipeline validation logs

---

### 7. PIPELINE VALIDATION

Validate:

- `RunPipeline` executes end-to-end
- no nil / empty unexpected states
- no broken references
- output remains stable

---

## OUTPUT REQUIREMENTS

Provide:

1. List of removed files / functions / structs
2. List of renamed elements
3. List of deduplicated entities
4. Summary of package restructuring
5. Confirmation that pipeline runs successfully

---

## STRICT WARNING

If something is unclear:
→ DO NOT GUESS
→ KEEP EXISTING BEHAVIOR

Cleanup must improve clarity without altering functionality.