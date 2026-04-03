# Pipeline Debug Fast Path

## Goal

Debug generator and platform behavior quickly without running the full pipeline across all game addons.

## Why this exists

Running the full source set is expensive for iteration. This workflow runs a targeted addon subset and writes outputs to debug folders, leaving canonical generated docs untouched.

## Script

Use:

- `scripts/debug_pipeline_subset.ps1`

This script:

1. accepts addon filters (`-Addons` or `-AddonListFile`)
2. generates addon artifacts only for the selected addons
3. generates platform docs from those subset artifacts
4. optionally generates site output
5. optionally runs the contract-family audit over the subset artifacts

Before generation, it materializes a temporary source root containing only the
selected addons (default: `.debug/source-subset`). Platform generation uses
that subset source root, avoiding full-source scans.

Outputs default to:

- `docs/addon-api-debug`
- `docs/war-api-debug`
- `docs/site-debug/content` (only when `-IncludeSite`)
- `.debug/source-subset`

## Quick Examples

Single addon debug:

```powershell
powershell -ExecutionPolicy Bypass -File .\scripts\debug_pipeline_subset.ps1 -Addons Clock
```

Two-addon regression pass:

```powershell
powershell -ExecutionPolicy Bypass -File .\scripts\debug_pipeline_subset.ps1 -Addons Clock,ClosetGoblin
```

Subset + site + audit:

```powershell
powershell -ExecutionPolicy Bypass -File .\scripts\debug_pipeline_subset.ps1 -Addons Clock,ClosetGoblin -IncludeSite -RunAudit
```

Addon list file (`#` comments and blank lines allowed):

```powershell
powershell -ExecutionPolicy Bypass -File .\scripts\debug_pipeline_subset.ps1 -AddonListFile .\agent-docs\debug_addons.txt
```

## Notes

- The script runs local `go run ... generate ...` commands directly, so it can target `./addons` precisely.
- Use `-SubsetSourceRoot` to change where the temporary subset source is staged.
- Use this workflow for fast iteration; use the canonical full workflow only when validating final behavior.
- Existing `make` targets are still useful, but they are not optimized for frequent subset debugging.
