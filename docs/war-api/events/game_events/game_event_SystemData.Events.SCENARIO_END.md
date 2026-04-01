# SystemData.Events.SCENARIO_END

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
| Addons seen in | Enemy, LibGuard, RoR_SoR, followTheLeader |
| Files seen in | `/workspace/Enemy/Code/Core/Groups/Groups.lua:22`, `/workspace/LibGuard/Source/LibGuard.lua:47`, `/workspace/RoR_SoR/RoR_SoR.lua:178`, `/workspace/followTheLeader/followTheLeader.lua:87` |
| Namespaces detected | SystemData |
| Source kinds | event_page, flows, lua_event_registration |
| Example locations | Enemy: Enemy.GroupsInitialize, LibGuard: LibGuard.Init, RoR_SoR: RoR_SoR:RegisterSelfEvents, followTheLeader: followTheLeader.Initialize |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 4 |
| Global usage count | 4 |
| Local definition count | 0 |
| Documentation references | 2 |
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

Observed as a shared SystemData runtime event used by 4 addons.

## Handler Pattern

Observed as a runtime event ID routed through RegisterEventHandler-style APIs.

## Payload

- Payload shape is not inferable from addon-api docs alone; treat this as an engine event identifier.

## Seen In

- Enemy
- LibGuard
- RoR_SoR
- followTheLeader

## Registrars And Handlers

- Enemy.GroupsUpdateType
- LibGuard.GROUP_UPDATED
- RegisterEventHandler
- RoR_SoR.OnScenario
- followTheLeader.OnScenarioEnd
- global

## Examples

- Enemy: Enemy.GroupsInitialize -> SystemData.Events.SCENARIO_END -> Enemy.GroupsUpdateType
- LibGuard: LibGuard.Init -> SystemData.Events.SCENARIO_END -> LibGuard.GROUP_UPDATED
- RoR_SoR: RoR_SoR:RegisterSelfEvents -> SystemData.Events.SCENARIO_END -> RoR_SoR.OnScenario
- followTheLeader: followTheLeader.Initialize -> SystemData.Events.SCENARIO_END -> followTheLeader.OnScenarioEnd
- Enemy: Enemy.GroupsUpdateType -> RegisterEventHandler(SystemData.Events.SCENARIO_END, Enemy.GroupsUpdateType)
- LibGuard: LibGuard.GROUP_UPDATED -> RegisterEventHandler(SystemData.Events.SCENARIO_END, LibGuard.GROUP_UPDATED)

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
