# SystemData.Events.GROUP_PLAYER_ADDED

- Category: SystemData Field
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 5 addons

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
| Addons seen in | Enemy, LibGuard, MegaphonePlus, Pure, followTheLeader |
| Files seen in | Code/Core/Groups/Groups.lua, MegaphonePlus.lua, Source/LibGuard.lua, Source/PureGroup.lua, Source/PureGroupPet.lua, followTheLeader.lua |
| Namespaces detected | SystemData |
| Source kinds | lua_call |
| Example locations | GroupsInitialize, Init, Initialize, LoadUnitFrame, OnShutdown, UnloadUnitFrame |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 7 |
| Global usage count | 7 |
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

SystemData.SystemData.Events.GROUP_PLAYER_ADDED field accessed by 5 addons; commonly found in GroupsInitialize and Init, Initialize, LoadUnitFrame, OnShutdown, UnloadUnitFrame, lua_call contexts.

## Seen In

- Enemy
- LibGuard
- MegaphonePlus
- Pure
- followTheLeader

## Related APIs

- [WindowSetDrawWhenInterfaceHidden](../../window_api/functions/window_WindowSetDrawWhenInterfaceHidden.md) (HIGH 100/100) - Window Function
- [OnShutdown](../../xml/handlers/handler_OnShutdown.md) (HIGH 88/100) - XML Event

## Used With

- [SystemData.Events.BATTLEGROUP_UPDATED](systemdata_SystemData.Events.BATTLEGROUP_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.GROUP_STATUS_UPDATED](systemdata_SystemData.Events.GROUP_STATUS_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.GROUP_UPDATED](systemdata_SystemData.Events.GROUP_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.LOADING_END](systemdata_SystemData.Events.LOADING_END.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_PET_HEALTH_UPDATED](systemdata_SystemData.Events.PLAYER_PET_HEALTH_UPDATED.md) (HIGH 100/100) - SystemData Field

## Notes

- Observed in contexts: GroupsInitialize, Init, Initialize, LoadUnitFrame, OnShutdown, UnloadUnitFrame
