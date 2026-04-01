# SystemData.Events.PLAYER_CAREER_RANK_UPDATED

- Category: SystemData Field
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 2 addons

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 176

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, matches a known engine namespace, referenced by generated docs or reference files.

## Evidence Signals

- +18 Seen in 2 to 3 addons: Cross-addon spread is present but limited.
- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +18 Used in event registration or dispatch: Observed in event-driven engine hooks.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +25 Referenced by generated docs or reference files: The symbol is reinforced outside a single call page.
- +20 Reinforced across multiple generated source types: Evidence comes from several independent API_Ref source types.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | Effigy, Enemy |
| Files seen in | `/workspace/Effigy/States/EffigyStatePlayer.lua:45`, `/workspace/Enemy/Code/Core/Groups/Groups.lua:22` |
| Namespaces detected | SystemData |
| Source kinds | event_page, event_registration, lua_call |
| Example locations | Effigy.RegisterStateInfoForPlayer, Enemy.GroupsInitialize, Enemy.Groups_OnCurrentPlayerUpdated, SystemData.Events.PLAYER_CAREER_RANK_UPDATED, event_page, event_registration |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 7 |
| Global usage count | 7 |
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

Observed SystemData field used by 2 addons through generated function calls, event pages, or lifecycle evidence.

## Seen In

- Effigy
- Enemy

## Related APIs

- [WindowRegisterCoreEventHandler](../../window_api/functions/window_WindowRegisterCoreEventHandler.md) (HIGH 100/100) - Window Function

## Used With

- none

## Triggered By

- none

## Affects

- none

## Notes

- Observed in contexts: Effigy.RegisterStateInfoForPlayer, Enemy.GroupsInitialize, Enemy.Groups_OnCurrentPlayerUpdated, SystemData.Events.PLAYER_CAREER_RANK_UPDATED, event_page, event_registration
