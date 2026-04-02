# SystemData.Events.SCENARIO_PLAYERS_LIST_GROUPS_UPDATED

- Category: SystemData Field
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 4 addons

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 188

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, seen in 4 or more addons, matches a known engine namespace.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +18 Used in event registration or dispatch: Observed in event-driven engine hooks.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +25 Referenced by generated docs or reference files: The symbol is reinforced outside a single call page.
- +20 Reinforced across multiple generated source types: Evidence comes from several independent addon-api source types.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | Enemy, Enemy, LibGuard, LibGroup, LibGuard |
| Files seen in | `/workspace_addons/Enemy/Code/Core/Groups/Groups.lua:22`, `/workspace_addons/LibGroup/LibGroup.lua:343`, `/workspace_addons/LibGuard/Source/LibGuard.lua:47`, `/workspace_addons/LibGuard/Source/LibGuard.lua:91` |
| Namespaces detected | SystemData |
| Source kinds | event_page, event_registration, lua_call |
| Example locations | Enemy.GroupsInitialize, Enemy.GroupsUpdate, LibGroup.Initialize, LibGuard.GROUP_UPDATED, LibGuard.Init, LibGuard.OnShutdown |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 10 |
| Global usage count | 10 |
| Local definition count | 0 |
| Documentation references | 3 |
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

Observed SystemData field used by 4 addons through generated function calls, event pages, or lifecycle evidence.

## Seen In

- Enemy
- Enemy, LibGuard
- LibGroup
- LibGuard

## Related APIs

- none

## Used With

- none

## Triggered By

- none

## Affects

- none

## Notes

- Observed in contexts: Enemy.GroupsInitialize, Enemy.GroupsUpdate, LibGroup.Initialize, LibGuard.GROUP_UPDATED, LibGuard.Init, LibGuard.OnShutdown
