# SystemData.Events.PLAYER_START_INTERACT_TIMER

- Category: SystemData Field
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 3 addons

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 186

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, matches a known engine namespace, referenced by generated docs or reference files.

## Evidence Signals

- +18 Seen in 2 to 3 addons: Cross-addon spread is present but limited.
- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +18 Used in event registration or dispatch: Observed in event-driven engine hooks.
- +10 Referenced from initialization flow: Lifecycle reconstruction references this symbol.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +25 Referenced by generated docs or reference files: The symbol is reinforced outside a single call page.
- +20 Reinforced across multiple generated source types: Evidence comes from several independent addon-api source types.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | PartyCast, PartyCast, ZCurse_Profiler, ZCurse_Profiler |
| Files seen in | `/workspace/data/raw/CurseProfiler/CurseProfilerCompiled.lua:2134`, `/workspace/data/raw/PartyCast/PartyCast.lua:153`, `/workspace/data/raw/PartyCast/PartyCast.lua:51` |
| Namespaces detected | SystemData |
| Source kinds | event_page, event_registration, flow, lua_call |
| Example locations | CurseProfiler.Initialize, CurseProfiler.untilin, PartyCast.Init, PartyCast.OnShutdown, PartyCast.StartInteract, SystemData.Events.PLAYER_START_INTERACT_TIMER |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 11 |
| Global usage count | 11 |
| Local definition count | 0 |
| Documentation references | 4 |
| Initialization flow references | 1 |
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

Observed SystemData field used by 3 addons through generated function calls, event pages, or lifecycle evidence.

## Seen In

- PartyCast
- PartyCast, ZCurse_Profiler
- ZCurse_Profiler

## Related APIs

- [LibSlash.RegisterSlashCmd](../../globals/functions/global_LibSlash.RegisterSlashCmd.md) (HIGH 100/100) - Global Function
- [RegisterEventHandler](../../globals/functions/global_RegisterEventHandler.md) (HIGH 93/100) - Global Function
- [UnregisterEventHandler](../../globals/functions/global_UnregisterEventHandler.md) (HIGH 81/100) - Global Function

## Used With

- none

## Triggered By

- none

## Affects

- none

## Notes

- Observed in contexts: CurseProfiler.Initialize, CurseProfiler.untilin, PartyCast.Init, PartyCast.OnShutdown, PartyCast.StartInteract, SystemData.Events.PLAYER_START_INTERACT_TIMER
