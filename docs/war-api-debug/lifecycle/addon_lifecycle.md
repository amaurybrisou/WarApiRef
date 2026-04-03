# Addon Lifecycle Semantics

> Structured lifecycle facts derived from `.mod` manifest analysis.

## What This List Represents

- This page is not an inventory of all addon directories under the source root.
- It lists addons that reached `.mod` semantic lifecycle extraction in the source-first pipeline.
- Addons can be excluded for explicit reasons (no manifest, no resolved source files, or `.toc`-only).
- Use the diagnostics section below to identify exclusions and counts.

## Diagnostics

| Metric | Value |
| --- | --- |
| source_root | C:/Return of Reckoning/Interface/AddOns/WAR_API_Ref/.debug/source-subset |
| source_directories | 2 |
| manifest_discovered | 2 |
| source_scanned_addons | 2 |
| mod_semantic_addons | 2 |
| lifecycle_catalog_addons | 2 |
| excluded_entries | 0 |

## Source Addon Inventory

| Directory | Lifecycle Entry | Status | Reason |
| --- | --- | --- | --- |
| Clock | [Clock](addons/clock) | included | .mod semantic lifecycle entry emitted |
| ClosetGoblin | ClosetGoblin | excluded | not present in .mod lifecycle semantic output |

## Lifecycle Addons (.mod Semantic Output)

- [CM_ClosetGoblin](addons/cm_closetgoblin)
- [Clock](addons/clock)
