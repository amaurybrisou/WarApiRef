# SystemData.Events.INTERACT_LEAVE_SCENARIO_QUEUE

- Category: Game Event
- Confidence level: HIGH
- Confidence score: 100/100

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
| Addons seen in | Enemy, Queue Queuer, TimeInQueue |
| Files seen in | `/workspace/Enemy/Code/ScenarioAlerter/ScenarioAlerter.lua:27`, `/workspace/QueueQueuer/QueueQueuer.lua:436`, `/workspace/TimeInQueue/TimeInQueue.lua:19` |
| Namespaces detected | SystemData |
| Source kinds | event_page, flows, lua_event_registration |
| Example locations | Enemy: Enemy._ScenarioAlerterEnabledChanged, Queue Queuer: QueueQueuer.OnInitialize, TimeInQueue: TimeInQueue.onInit |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 3 |
| Global usage count | 3 |
| Local definition count | 0 |
| Documentation references | 2 |
| Initialization flow references | 2 |
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

Observed as a shared SystemData runtime event used by 3 addons.

## Handler Pattern

Observed as a runtime event ID routed through RegisterEventHandler-style APIs.

## Payload

- Payload shape is not inferable from addon-api docs alone; treat this as an engine event identifier.

## Seen In

- Enemy
- Queue Queuer
- TimeInQueue

## Registrars And Handlers

- Enemy.ScenarioAlerter_OnInteractLeaveScenarioQueue
- QueueQueuer.OnLeave
- RegisterEventHandler
- TimeInQueue.ResetTimer
- global

## Examples

- Enemy: Enemy._ScenarioAlerterEnabledChanged -> SystemData.Events.INTERACT_LEAVE_SCENARIO_QUEUE -> Enemy.ScenarioAlerter_OnInteractLeaveScenarioQueue
- Queue Queuer: QueueQueuer.OnInitialize -> SystemData.Events.INTERACT_LEAVE_SCENARIO_QUEUE -> QueueQueuer.OnLeave
- TimeInQueue: TimeInQueue.onInit -> SystemData.Events.INTERACT_LEAVE_SCENARIO_QUEUE -> TimeInQueue.ResetTimer
- Enemy: Enemy.ScenarioAlerter_OnInteractLeaveScenarioQueue -> RegisterEventHandler(SystemData.Events.INTERACT_LEAVE_SCENARIO_QUEUE, Enemy.ScenarioAlerter_OnInteractLeaveScenarioQueue)
- Queue Queuer: QueueQueuer.OnLeave -> RegisterEventHandler(SystemData.Events.INTERACT_LEAVE_SCENARIO_QUEUE, QueueQueuer.OnLeave)
- TimeInQueue: TimeInQueue.ResetTimer -> RegisterEventHandler(SystemData.Events.INTERACT_LEAVE_SCENARIO_QUEUE, TimeInQueue.ResetTimer)

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
