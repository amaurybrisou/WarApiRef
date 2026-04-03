# SystemData.Events.PLAYER_POSITION_UPDATED

- Category: SystemData Field
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 17 addons

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 168

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, seen in 4 or more addons, matches a known engine namespace.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +18 Used in event registration or dispatch: Observed in event-driven engine hooks.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +25 Referenced by generated docs or reference files: The symbol is reinforced outside a single call page.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | Atlas, BuffHead, CCTV, Compass3D, Ding, LibSurveyor, MapMonster, MiniMapMonster |
| Files seen in | CCTV.lua, Core.lua, CurseProfilerCompiled.lua, Ding.lua, LibSurveyor.lua, Map/Main.lua, MiniMapMonster.lua, NerfedButtons.lua |
| Namespaces detected | SystemData |
| Source kinds | lua_call |
| Example locations | Init, Initialize, InitializePlayer, OnInitialize, OnShutdown, RegisterEventHandlers |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 13 |
| Global usage count | 13 |
| Local definition count | 0 |
| Documentation references | 1 |
| Initialization flow references | 0 |
| Known engine namespace | yes |
| Default UI presence | yes |
| Event binding presence | yes |
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

SystemData.SystemData.Events.PLAYER_POSITION_UPDATED field accessed by 17 addons; commonly found in Init and Initialize, InitializePlayer, OnInitialize, OnShutdown, RegisterEventHandlers, Shutdown, ShutdownPlayer, Unload, UnregisterEventHandlers, load, lua_call, unload contexts.

## Seen In

- Atlas
- BuffHead
- CCTV
- Compass3D
- Ding
- LibSurveyor
- MapMonster
- MiniMapMonster
- NerfedButtons
- RVAPI_Range
- Rangechecker
- SimpleXY
- TomeTracker
- TurretRange
- WarBoard_Loc
- ZCurse_Profiler
- compass

## Related APIs

- [OnInitialize](../../xml/handlers/handler_OnInitialize.md) (HIGH 88/100) - XML Event
- [OnShutdown](../../xml/handlers/handler_OnShutdown.md) (HIGH 88/100) - XML Event

## Notes

- Observed in contexts: Init, Initialize, InitializePlayer, OnInitialize, OnShutdown, RegisterEventHandlers
