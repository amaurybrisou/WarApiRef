# EA_Window_CurrentEvents

- Category: Global Table
- Confidence level: HIGH
- Confidence score: 80/100

## Confidence Assessment

- Level: HIGH

- Score: 80/100

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, matches a known engine namespace, called globally with no local definition.

## Evidence Signals

- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | QuickWarReport |
| Files seen in | `/workspace_addons/QuickWarReport/QWRCore.lua:62`, `/workspace_addons/QuickWarReport/QWRCore.lua:89` |
| Namespaces detected | EA_Window_CurrentEvents |
| Source kinds | lua_calls |
| Example locations | QuickWarReport: QuickWarReport.JumpToTier, QuickWarReport: QuickWarReport.JumpToTierCurrentZone |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 2 |
| Global usage count | 1 |
| Local definition count | 0 |
| Documentation references | 0 |
| Initialization flow references | 0 |
| Known engine namespace | yes |
| Default UI presence | yes |
| Event binding presence | no |
| Observed in XML and Lua | no |
| Consistent role | yes |
| Consistent arguments | no |
| Consistent returns | no |
| Slash command presence | no |
| Weak usage only | no |
| Project-specific name | no |
| Placeholder or computed name | no |
| Conflicting signatures | no |
| Conflicting roles | no |
| Wrapper likely | no |
| Never outside local graph | no |
| Local helper only | no |

## Description

Observed shared global table or namespace surfaced in 1 addons.

## Functions

- EA_Window_CurrentEvents.SetCurrentTier

## Observed Members

- none

## Seen In

- QuickWarReport

## Examples

- QuickWarReport: QuickWarReport.JumpToTier -> EA_Window_CurrentEvents.SetCurrentTier(tier)
- QuickWarReport: QuickWarReport.JumpToTierCurrentZone -> EA_Window_CurrentEvents.SetCurrentTier(tier)

## Related APIs

- none

## Used With

- none

## Triggered By

- none

## Affects

- none

## Notes

- none
