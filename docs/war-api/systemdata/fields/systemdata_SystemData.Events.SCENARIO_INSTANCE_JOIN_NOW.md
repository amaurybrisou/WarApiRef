# SystemData.Events.SCENARIO_INSTANCE_JOIN_NOW

- Category: SystemData Field
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 4 addons

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
| Addons seen in | Enemy, Enemy, Queue Queuer, TimeInQueue, Queue Queuer, TimeInQueue |
| Files seen in | `/workspace/Enemy/Code/ScenarioAlerter/ScenarioAlerter.lua:27`, `/workspace/QueueQueuer/QueueQueuer.lua:436`, `/workspace/QueueQueuer/QueueQueuer.lua:466`, `/workspace/TimeInQueue/TimeInQueue.lua:19` |
| Namespaces detected | SystemData |
| Source kinds | event_page, event_registration, flow, lua_call |
| Example locations | Enemy.ScenarioAlerter_OnScenarioInstanceJoinNow, Enemy._ScenarioAlerterEnabledChanged, QueueQueuer.OnInitialize, QueueQueuer.OnJoinNow, QueueQueuer.OnShutdown, SystemData.Events.SCENARIO_INSTANCE_JOIN_NOW |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 13 |
| Global usage count | 13 |
| Local definition count | 0 |
| Documentation references | 4 |
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

Observed SystemData field used by 4 addons through generated function calls, event pages, or lifecycle evidence.

## Seen In

- Enemy
- Enemy, Queue Queuer, TimeInQueue
- Queue Queuer
- TimeInQueue

## Related APIs

- none

## Used With

- none

## Triggered By

- none

## Affects

- none

## Notes

- Observed in contexts: Enemy.ScenarioAlerter_OnScenarioInstanceJoinNow, Enemy._ScenarioAlerterEnabledChanged, QueueQueuer.OnInitialize, QueueQueuer.OnJoinNow, QueueQueuer.OnShutdown, SystemData.Events.SCENARIO_INSTANCE_JOIN_NOW
