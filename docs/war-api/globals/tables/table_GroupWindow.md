# GroupWindow

- Category: Global Table
- Confidence level: HIGH
- Confidence score: 100/100

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 128

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, matches a known engine namespace, referenced by generated docs or reference files.

## Evidence Signals

- +18 Seen in 2 to 3 addons: Cross-addon spread is present but limited.
- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +10 Referenced from initialization flow: Lifecycle reconstruction references this symbol.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +25 Referenced by generated docs or reference files: The symbol is reinforced outside a single call page.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | DeepSleep, Effigy |
| Files seen in | `/workspace/Effigy/Effigy.lua:501` |
| Namespaces detected | GroupWindow |
| Source kinds | globals, lua_calls |
| Example locations | Effigy: Effigy.RButtonUp |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 1 |
| Global usage count | 1 |
| Local definition count | 1 |
| Documentation references | 1 |
| Initialization flow references | 1 |
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

Observed shared global table or namespace surfaced in 2 addons.

## Functions

- GroupWindow.ShowMenu

## Observed Members

- none

## Seen In

- DeepSleep
- Effigy

## Examples

- Effigy: Effigy.RButtonUp -> GroupWindow.ShowMenu(member_name)

## Related APIs

- [PartyUtils.GetPartyData](../functions/global_PartyUtils.GetPartyData.md) (HIGH 100/100) - Global Function

## Used With

- [PartyUtils.GetPartyData](../functions/global_PartyUtils.GetPartyData.md) (HIGH 100/100) - Global Function

## Triggered By

- none

## Affects

- none

## Notes

- none
