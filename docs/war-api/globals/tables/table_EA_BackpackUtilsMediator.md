# EA_BackpackUtilsMediator

- Category: Global Table
- Confidence level: HIGH
- Confidence score: 100/100

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 118

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, matches a known engine namespace, referenced by generated docs or reference files.

## Evidence Signals

- +18 Seen in 2 to 3 addons: Cross-addon spread is present but limited.
- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +25 Referenced by generated docs or reference files: The symbol is reinforced outside a single call page.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | BagOMatic, TidyRoll |
| Files seen in | `/workspace/data/raw/TidyRoll/CustomAutoRoll.lua:400`, `/workspace/data/raw/bagomatic/BagOMatic.lua:818`, `/workspace/data/raw/bagomatic/BagOMatic.lua:888` |
| Namespaces detected | EA_BackpackUtilsMediator |
| Source kinds | globals, lua_calls |
| Example locations | BagOMatic: BagOMatic.SalvageHook, BagOMatic: BagOMatic.findItemInBagPack, TidyRoll: TidyRoll.CustomAutoRoll.OnListLbuttonUp |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 7 |
| Global usage count | 4 |
| Local definition count | 2 |
| Documentation references | 1 |
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

Observed shared global table or namespace surfaced in 2 addons.

## Functions

- EA_BackpackUtilsMediator.GetBackpack
- EA_BackpackUtilsMediator.GetCurrentBackpackType
- EA_BackpackUtilsMediator.GetCursorForBackpack
- EA_BackpackUtilsMediator.GetItemsFromBackpack

## Observed Members

- none

## Seen In

- BagOMatic
- TidyRoll

## Examples

- BagOMatic: BagOMatic.SalvageHook -> EA_BackpackUtilsMediator.GetBackpack()
- BagOMatic: BagOMatic.SalvageHook -> EA_BackpackUtilsMediator.GetItemsFromBackpack(backpackType)
- BagOMatic: BagOMatic.findItemInBagPack -> EA_BackpackUtilsMediator.GetCurrentBackpackType()
- BagOMatic: BagOMatic.findItemInBagPack -> EA_BackpackUtilsMediator.GetItemsFromBackpack(backpackType)
- TidyRoll: TidyRoll.CustomAutoRoll.OnListLbuttonUp -> EA_BackpackUtilsMediator.GetCurrentBackpackType()
- TidyRoll: TidyRoll.CustomAutoRoll.OnListLbuttonUp -> EA_BackpackUtilsMediator.GetCursorForBackpack(backpackType)

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
