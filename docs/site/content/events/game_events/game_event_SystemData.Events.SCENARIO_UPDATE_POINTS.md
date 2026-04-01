# SystemData.Events.SCENARIO_UPDATE_POINTS

- Category: Game Event
- Confidence level: HIGH
- Confidence score: 100/100

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 123

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, matches a known engine namespace, referenced by generated docs or reference files.

## Evidence Signals

- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +18 Used in event registration or dispatch: Observed in event-driven engine hooks.
- +25 Referenced by generated docs or reference files: The symbol is reinforced outside a single call page.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | Enemy |
| Files seen in | `/workspace/Enemy/Code/ScenarioInfo/ScenarioInfo.lua:303` |
| Namespaces detected | SystemData |
| Source kinds | event_page, lua_event_registration |
| Example locations | Enemy: Enemy._ScenarioInfoEnabledChanged |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 1 |
| Global usage count | 1 |
| Local definition count | 0 |
| Documentation references | 1 |
| Initialization flow references | 0 |
| Known engine namespace | yes |
| Default UI presence | yes |
| Event binding presence | yes |
| Observed in XML and Lua | no |
| Consistent role | no |
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

Observed as a shared SystemData runtime event used by 1 addons.

## Handler Pattern

Observed as a runtime event ID routed through RegisterEventHandler-style APIs.

## Payload

- Payload shape is not inferable from API_Ref alone; treat this as an engine event identifier.

## Seen In

- Enemy

## Registrars And Handlers

- Enemy.ScenarioInfoUpdate
- RegisterEventHandler
- global

## Examples

- Enemy: Enemy._ScenarioInfoEnabledChanged -> SystemData.Events.SCENARIO_UPDATE_POINTS -> Enemy.ScenarioInfoUpdate
- Enemy: Enemy.ScenarioInfoUpdate -> RegisterEventHandler(SystemData.Events.SCENARIO_UPDATE_POINTS, Enemy.ScenarioInfoUpdate)

## Related APIs

- none

## Used With

- none

## Triggered By

- none

## Affects

- none

## Notes

- Only one addon surfaced this event in the current API_Ref corpus.
