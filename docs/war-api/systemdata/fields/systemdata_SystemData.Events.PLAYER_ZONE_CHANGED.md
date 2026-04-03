# SystemData.Events.PLAYER_ZONE_CHANGED

- Category: SystemData Field
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 27 addons

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
| Addons seen in | Atlas, Auto Cape/Helm Show, BarText (Influence), CM_ClosetGoblin, CMap, Ding, EA_OpenPartyWindow, Enemy |
| Files seen in | BarText_Influence.lua, CMap.lua, ClosetGoblin.lua, Code/Core/Groups/Groups.lua, Code/GroupIcons/GroupIcons.lua, Code/TalismanAlerter/TalismanAlerter.lua, Core.lua, CurseProfilerCompiled.lua |
| Namespaces detected | SystemData |
| Source kinds | lua_call |
| Example locations | GroupsInitialize, Init, Initialize, InitializePlayer, OnDisable, OnEnable |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 20 |
| Global usage count | 20 |
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

SystemData.SystemData.Events.PLAYER_ZONE_CHANGED field accessed by 27 addons; commonly found in GroupsInitialize and Init, Initialize, InitializePlayer, OnDisable, OnEnable, OnInitialize, OnShutdown, RegisterEventHandlers, RegisterEvents, Shutdown, ShutdownPlayer, UnregisterEventHandlers, UnregisterEvents, ZoneChangeInit, ZoneChangeShutdown, _GroupIconsEnabledChanged, _TalismanAlerterEnabledChanged, init, lua_call contexts.

## Seen In

- Atlas
- Auto Cape/Helm Show
- BarText (Influence)
- CM_ClosetGoblin
- CMap
- Ding
- EA_OpenPartyWindow
- Enemy
- GCDsaver
- Killer
- Kwestor
- LibWBToggler
- Map
- MapMonster
- MarkBuff
- MegaphonePlusPlus
- MiniMapMonster
- Minmap
- RVAPI_Range
- RoR_debolster
- Shinies
- VPBreakdown
- WarBoard_TaliMon
- ZCurse_Profiler
- ZonePOP
- hideInf
- talisman-monitor

## Related APIs

- [ButtonSetStayDownFlag](../../window_api/functions/window_ButtonSetStayDownFlag.md) (HIGH 100/100) - Window Function
- [WindowRegisterEventHandler](../../window_api/functions/window_WindowRegisterEventHandler.md) (HIGH 100/100) - Window Function
- [OnInitialize](../../xml/handlers/handler_OnInitialize.md) (HIGH 88/100) - XML Event
- [OnShutdown](../../xml/handlers/handler_OnShutdown.md) (HIGH 88/100) - XML Event
- [DefaultColor.LabelSetTextColor](../../globals/functions/global_DefaultColor.LabelSetTextColor.md) (MEDIUM 55/100) - Global Function

## Used With

- [SystemData.Events.BATTLEGROUP_MEMBER_UPDATED](systemdata_SystemData.Events.BATTLEGROUP_MEMBER_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.BATTLEGROUP_UPDATED](systemdata_SystemData.Events.BATTLEGROUP_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.GROUP_SETTINGS_UPDATED](systemdata_SystemData.Events.GROUP_SETTINGS_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.GROUP_UPDATED](systemdata_SystemData.Events.GROUP_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.L_BUTTON_UP_PROCESSED](systemdata_SystemData.Events.L_BUTTON_UP_PROCESSED.md) (HIGH 100/100) - SystemData Field

## Notes

- Observed in contexts: GroupsInitialize, Init, Initialize, InitializePlayer, OnDisable, OnEnable
