# GameData.TradeSkills.CULTIVATION

- Category: GameData Field
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 4 addons

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 150

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, seen in 4 or more addons, matches a known engine namespace.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +25 Referenced by generated docs or reference files: The symbol is reinforced outside a single call page.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | Crafting Info Tooltip, Miracle Grow Remix, MiracleGrow, MiracleGrowLight |
| Files seen in | `/workspace/CraftValueTip/CraftInfoAPI.lua:138`, `/workspace/MGRemix/harvrepeat.lua:215`, `/workspace/MGRemix/harvrepeat.lua:343`, `/workspace/MiracleGrow/MiracleGrow.lua:60`, `/workspace/MiracleGrowLight/MiracleGrowLight.lua:301` |
| Namespaces detected | GameData |
| Source kinds | lua_call |
| Example locations | CraftItemInfo.GetItemFamily, MiracleGrow.Initialize, MiracleGrow2.HRInit, MiracleGrow2.HRUpdatePlot, MiracleGrowLight.Initialize, lua_call |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 6 |
| Global usage count | 6 |
| Local definition count | 0 |
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

Observed GameData field used by 4 addons through generated function calls, event pages, or lifecycle evidence.

## Seen In

- Crafting Info Tooltip
- Miracle Grow Remix
- MiracleGrow
- MiracleGrowLight

## Related APIs

- none

## Used With

- none

## Triggered By

- none

## Affects

- none

## Notes

- Observed in contexts: CraftItemInfo.GetItemFamily, MiracleGrow.Initialize, MiracleGrow2.HRInit, MiracleGrow2.HRUpdatePlot, MiracleGrowLight.Initialize, lua_call
