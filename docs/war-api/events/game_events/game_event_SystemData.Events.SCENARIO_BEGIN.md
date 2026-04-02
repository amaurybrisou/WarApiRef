# SystemData.Events.SCENARIO_BEGIN

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
| Addons seen in | Enemy, RoR_SoR, followTheLeader |
| Files seen in | `/workspace_addons/Enemy/Code/Core/Groups/Groups.lua:22`, `/workspace_addons/Enemy/Code/ScenarioInfo/ScenarioInfo.lua:303`, `/workspace_addons/RoR_SoR/RoR_SoR.lua:178`, `/workspace_addons/followTheLeader/followTheLeader.lua:87` |
| Namespaces detected | SystemData |
| Source kinds | event_page, lua_event_registration |
| Example locations | Enemy: Enemy.GroupsInitialize, Enemy: Enemy._ScenarioInfoEnabledChanged, RoR_SoR: RoR_SoR:RegisterSelfEvents, followTheLeader: followTheLeader.Initialize |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 4 |
| Global usage count | 4 |
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

Observed as a shared SystemData runtime event used by 3 addons.

## Handler Pattern

Observed as a runtime event ID routed through RegisterEventHandler-style APIs.

## Payload

- Payload shape is not inferable from addon-api docs alone; treat this as an engine event identifier.

## Seen In

- Enemy
- RoR_SoR
- followTheLeader

## Registrars And Handlers

- Enemy.GroupsUpdateType
- Enemy.ScenarioInfoUpdate
- RegisterEventHandler
- RoR_SoR.OnScenario
- followTheLeader.OnScenarioBegin
- global

## Examples

- Enemy: Enemy.GroupsInitialize -> SystemData.Events.SCENARIO_BEGIN -> Enemy.GroupsUpdateType
- Enemy: Enemy._ScenarioInfoEnabledChanged -> SystemData.Events.SCENARIO_BEGIN -> Enemy.ScenarioInfoUpdate
- RoR_SoR: RoR_SoR:RegisterSelfEvents -> SystemData.Events.SCENARIO_BEGIN -> RoR_SoR.OnScenario
- followTheLeader: followTheLeader.Initialize -> SystemData.Events.SCENARIO_BEGIN -> followTheLeader.OnScenarioBegin
- Enemy: Enemy.GroupsUpdateType -> RegisterEventHandler(SystemData.Events.SCENARIO_BEGIN, Enemy.GroupsUpdateType)
- Enemy: Enemy.ScenarioInfoUpdate -> RegisterEventHandler(SystemData.Events.SCENARIO_BEGIN, Enemy.ScenarioInfoUpdate)

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
