# SystemData.Events.SCENARIO_SHOW_JOIN_PROMPT

- Category: Game Event
- Confidence level: HIGH
- Confidence score: 100/100

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 198

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, seen in 4 or more addons, matches a known engine namespace.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
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
| Addons seen in | Enemy, ScenarioAlert, TidyQueue, TimeInQueue |
| Files seen in | `/workspace/Enemy/Code/ScenarioAlerter/ScenarioAlerter.lua:27`, `/workspace/ScenarioAlert/ScenarioAlert.lua:4`, `/workspace/TidyQueue/TidyQueue.lua:256`, `/workspace/TimeInQueue/TimeInQueue.lua:19` |
| Namespaces detected | SystemData |
| Source kinds | event_page, flows, lua_event_registration |
| Example locations | Enemy: Enemy._ScenarioAlerterEnabledChanged, ScenarioAlert: ScenarioAlert.OnInitialize, TidyQueue: TidyQueue.Initialize, TimeInQueue: TimeInQueue.onInit |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 5 |
| Global usage count | 5 |
| Local definition count | 0 |
| Documentation references | 2 |
| Initialization flow references | 3 |
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

Observed as a shared SystemData runtime event used by 4 addons.

## Handler Pattern

Observed as a runtime event ID routed through RegisterEventHandler-style APIs.

## Payload

- Payload shape is not inferable from addon-api docs alone; treat this as an engine event identifier.

## Seen In

- Enemy
- ScenarioAlert
- TidyQueue
- TimeInQueue

## Registrars And Handlers

- Enemy.ScenarioAlerter_OnScenarioShowJoinPrompt
- RegisterEventHandler
- ScenarioAlert.recordScPop
- TidyQueue.JoinPromptShowed
- TimeInQueue.ResetTimer
- TimeInQueue.scPop
- global

## Examples

- Enemy: Enemy._ScenarioAlerterEnabledChanged -> SystemData.Events.SCENARIO_SHOW_JOIN_PROMPT -> Enemy.ScenarioAlerter_OnScenarioShowJoinPrompt
- ScenarioAlert: ScenarioAlert.OnInitialize -> SystemData.Events.SCENARIO_SHOW_JOIN_PROMPT -> ScenarioAlert.recordScPop
- TidyQueue: TidyQueue.Initialize -> SystemData.Events.SCENARIO_SHOW_JOIN_PROMPT -> TidyQueue.JoinPromptShowed
- TimeInQueue: TimeInQueue.onInit -> SystemData.Events.SCENARIO_SHOW_JOIN_PROMPT -> TimeInQueue.ResetTimer
- TimeInQueue: TimeInQueue.onInit -> SystemData.Events.SCENARIO_SHOW_JOIN_PROMPT -> TimeInQueue.scPop
- Enemy: Enemy.ScenarioAlerter_OnScenarioShowJoinPrompt -> RegisterEventHandler(SystemData.Events.SCENARIO_SHOW_JOIN_PROMPT, Enemy.ScenarioAlerter_OnScenarioShowJoinPrompt)

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
