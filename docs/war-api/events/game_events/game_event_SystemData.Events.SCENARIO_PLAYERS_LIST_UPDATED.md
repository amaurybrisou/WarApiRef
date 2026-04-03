# SystemData.Events.SCENARIO_PLAYERS_LIST_UPDATED

- Category: Game Event
- Confidence level: HIGH
- Confidence score: 100/100

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 156

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, matches a known engine namespace, referenced by generated docs or reference files.

## Evidence Signals

- +18 Seen in 2 to 3 addons: Cross-addon spread is present but limited.
- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +18 Used in event registration or dispatch: Observed in event-driven engine hooks.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +25 Referenced by generated docs or reference files: The symbol is reinforced outside a single call page.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | Enemy, followTheLeader |
| Files seen in | `/workspace/data/raw/Enemy/Code/ScenarioInfo/ScenarioInfo.lua:303`, `/workspace/data/raw/followTheLeader/followTheLeader.lua:87` |
| Namespaces detected | SystemData |
| Source kinds | event_page, lua_event_registration |
| Example locations | Enemy: Enemy._ScenarioInfoEnabledChanged, followTheLeader: followTheLeader.Initialize |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 2 |
| Global usage count | 2 |
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

Observed as a shared SystemData runtime event used by 2 addons.

## Handler Pattern

Observed as a runtime event ID routed through RegisterEventHandler-style APIs.

## Payload

- Payload shape is not inferable from addon-api docs alone; treat this as an engine event identifier.

## Seen In

- Enemy
- followTheLeader

## Registrars And Handlers

- Enemy.ScenarioInfoUpdate
- RegisterEventHandler
- followTheLeader.RequestUpdate
- global

## Examples

- Enemy: Enemy._ScenarioInfoEnabledChanged -> SystemData.Events.SCENARIO_PLAYERS_LIST_UPDATED -> Enemy.ScenarioInfoUpdate
- followTheLeader: followTheLeader.Initialize -> SystemData.Events.SCENARIO_PLAYERS_LIST_UPDATED -> followTheLeader.RequestUpdate
- Enemy: Enemy.ScenarioInfoUpdate -> RegisterEventHandler(SystemData.Events.SCENARIO_PLAYERS_LIST_UPDATED, Enemy.ScenarioInfoUpdate)
- followTheLeader: followTheLeader.RequestUpdate -> RegisterEventHandler(SystemData.Events.SCENARIO_PLAYERS_LIST_UPDATED, followTheLeader.RequestUpdate)

## Related APIs

- [RegisterEventHandler](../../globals/functions/global_RegisterEventHandler.md) (HIGH 93/100) - Global Function

## Used With

- none

## Triggered By

- none

## Affects

- none

## Notes

- none
