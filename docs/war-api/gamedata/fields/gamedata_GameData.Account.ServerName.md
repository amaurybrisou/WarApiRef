# GameData.Account.ServerName

- Category: GameData Field
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 8 addons

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
| Addons seen in | Arsenal Rank, Asshat, CDown, GuardRange, HealGrid, MyReasons, SessionRPs, TacticSetNames |
| Files seen in | ArsenalRank.lua, Asshat.lua, CDown.lua, CDownSettings.lua, GuardRange.lua, HealGrid.lua, MyReasons.lua, source/TacticSetNames.lua |
| Namespaces detected | GameData |
| Source kinds | lua_call |
| Example locations | CheckSettings, FetchStats, Init, Initialize, OnInitialize, SettingsInitialize |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 10 |
| Global usage count | 10 |
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

GameData.GameData.Account.ServerName field accessed by 8 addons; commonly found in CheckSettings and FetchStats, Init, Initialize, OnInitialize, SettingsInitialize, ShowTab, StartTracker, lua_call, update contexts.

## Seen In

- Arsenal Rank
- Asshat
- CDown
- GuardRange
- HealGrid
- MyReasons
- SessionRPs
- TacticSetNames

## Related APIs

- [GameData.GetScenarioPlayers](../../globals/functions/global_GameData.GetScenarioPlayers.md) (HIGH 100/100) - Global Function
- [LayoutEditor.RegisterEditCallback](../../window_api/functions/window_LayoutEditor.RegisterEditCallback.md) (HIGH 100/100) - Window Function
- [WindowSetLayer](../../window_api/functions/window_WindowSetLayer.md) (HIGH 100/100) - Window Function
- [OnInitialize](../../xml/handlers/handler_OnInitialize.md) (HIGH 88/100) - XML Event

## Used With

- [SystemData.Events.INTERFACE_RELOADED](../../systemdata/fields/systemdata_SystemData.Events.INTERFACE_RELOADED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.LOADING_END](../../systemdata/fields/systemdata_SystemData.Events.LOADING_END.md) (HIGH 100/100) - SystemData Field

## Notes

- Observed in contexts: CheckSettings, FetchStats, Init, Initialize, OnInitialize, SettingsInitialize
