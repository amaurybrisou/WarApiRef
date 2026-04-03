# SystemData.Events.INTERACT_LEAVE_SCENARIO_QUEUE

- Category: SystemData Field
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 4 addons

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
| Addons seen in | Enemy, Minmap, Queue Queuer, TidyQueue |
| Files seen in | Code/ScenarioAlerter/ScenarioAlerter.lua, QueueQueuer.lua, TidyQueue.lua, core.lua |
| Namespaces detected | SystemData |
| Source kinds | lua_call |
| Example locations | DoTask, LeaveNextScenarios, LeaveQueues, LeaveScenario, OnInitialize, OnShutdown |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 8 |
| Global usage count | 8 |
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

SystemData.SystemData.Events.INTERACT_LEAVE_SCENARIO_QUEUE field accessed by 4 addons; commonly found in DoTask and LeaveNextScenarios, LeaveQueues, LeaveScenario, OnInitialize, OnShutdown, _ScenarioAlerterEnabledChanged, lua_call contexts.

## Seen In

- Enemy
- Minmap
- Queue Queuer
- TidyQueue

## Related APIs

- [EA_Window_ScenarioLobby.OnLeaveActiveQueueFromLobby](../../globals/functions/global_EA_Window_ScenarioLobby.OnLeaveActiveQueueFromLobby.md) (HIGH 90/100) - Global Function
- [OnInitialize](../../xml/handlers/handler_OnInitialize.md) (HIGH 88/100) - XML Event
- [OnShutdown](../../xml/handlers/handler_OnShutdown.md) (HIGH 88/100) - XML Event

## Notes

- Observed in contexts: DoTask, LeaveNextScenarios, LeaveQueues, LeaveScenario, OnInitialize, OnShutdown
