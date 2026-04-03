# SystemData.Events.BATTLEGROUP_UPDATED

- Category: SystemData Field
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 12 addons

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
| Addons seen in | Aura, AutoBand, Enemy, EveryBodyGuard, Hopper, LibGroup, MegaphonePlus, MegaphonePlusPlus |
| Files seen in | AutoBand.lua, Code/Core/Groups/Groups.lua, EveryBodyGuard.lua, LibGroup.lua, MegaphonePlus.lua, MegaphonePlusPlus.lua, RVAPI_Range.lua, Refer.lua |
| Namespaces detected | SystemData |
| Source kinds | event_page, event_registration, lua_call |
| Example locations | GroupsInitialize, Initialize, JumpStartEventBasedAuras, LoadUnitFrame, OnLoad, OnShutdown |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 20 |
| Global usage count | 20 |
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

SystemData.SystemData.Events.BATTLEGROUP_UPDATED field accessed by 12 addons; commonly found in GroupsInitialize and Initialize, JumpStartEventBasedAuras, LoadUnitFrame, OnLoad, OnShutdown, OnUnload, RegisterEventHandlers, Shutdown, Squared.ChangeMode, SquaredWarband.UpdateGroupProxy, Start, Stop, SystemData.Events.BATTLEGROUP_UPDATED, UnloadUnitFrame, UnregisterEventHandlers, event_page, event_registration, init, lua_call contexts.

## Seen In

- Aura
- AutoBand
- Enemy
- EveryBodyGuard
- Hopper
- LibGroup
- MegaphonePlus
- MegaphonePlusPlus
- Pure
- RVAPI_Range
- Refer
- Squared

## Related APIs

- [LayoutEditor.RegisterWindow](../../window_api/functions/window_LayoutEditor.RegisterWindow.md) (HIGH 100/100) - Window Function
- [LibSlash.IsSlashCmdRegistered](../../globals/functions/global_LibSlash.IsSlashCmdRegistered.md) (HIGH 100/100) - Global Function
- [WindowSetDrawWhenInterfaceHidden](../../window_api/functions/window_WindowSetDrawWhenInterfaceHidden.md) (HIGH 100/100) - Window Function
- [OnShutdown](../../xml/handlers/handler_OnShutdown.md) (HIGH 88/100) - XML Event
- [RegisterEventHandler](../../globals/functions/global_RegisterEventHandler.md) (MEDIUM 68/100) - Global Function
- [Start](../../events/game_events/game_event_Start.md) (MEDIUM 43/100) - Game Event
- [Stop](../../events/game_events/game_event_Stop.md) (MEDIUM 43/100) - Game Event

## Used With

- [SystemData.Events.GROUP_PLAYER_ADDED](systemdata_SystemData.Events.GROUP_PLAYER_ADDED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.GROUP_STATUS_UPDATED](systemdata_SystemData.Events.GROUP_STATUS_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.GROUP_UPDATED](systemdata_SystemData.Events.GROUP_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.LOADING_END](systemdata_SystemData.Events.LOADING_END.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_PET_HEALTH_UPDATED](systemdata_SystemData.Events.PLAYER_PET_HEALTH_UPDATED.md) (HIGH 100/100) - SystemData Field

## Notes

- Observed in contexts: GroupsInitialize, Initialize, JumpStartEventBasedAuras, LoadUnitFrame, OnLoad, OnShutdown
