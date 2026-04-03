# SystemData.Events.PLAYER_TARGET_UPDATED

- Category: SystemData Field
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 42 addons

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
| Addons seen in | AutoMark, BlackBook, BuffHead, Calling, Calling, DammazKron, LibRange, Squared, CastSequence, CleanUnitFrames, CleansingBuddy |
| Files seen in | BlackBook.lua, Calling.lua, CallingEvents.lua, CallingTargetMarker.lua, CleanTargetWindow.lua, CleansingBuddy.lua, Code/Core/Groups/Groups.lua, Code/Core/Main.lua |
| Namespaces detected | SystemData |
| Source kinds | event_page, event_registration, lua_call |
| Example locations | Calling.OnTargetChanged, GroupsInitialize, Init, Initialize, LibRange.OnPlayerTargetUpdated, LoadUnitFrame |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 32 |
| Global usage count | 32 |
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

SystemData.SystemData.Events.PLAYER_TARGET_UPDATED field accessed by 42 addons; commonly found in Calling.OnTargetChanged and GroupsInitialize, Init, Initialize, LibRange.OnPlayerTargetUpdated, LoadUnitFrame, OnInitialize, OnLoadingEnd, OnShutdown, OnSlashCommand, RegisterEventHandler, RegisterEventHandlers, RegisterEvents, SetShowMarker, Set_autoMode, Shutdown, Squared.TargetUpdate, SquaredRangeFading.TargetUpdated, Start, StateTrackingSet, Stop, SystemData.Events.PLAYER_TARGET_UPDATED, TargetInfoFix.SET_TARGETINFO_FIX_UPDATE_FLAG_DONOTTOUCH, UnloadUnitFrame, UnregisterEventHandler, UnregisterEventHandlers, UnregisterEvents, _Initialize, event_page, event_registration, init, lua_call contexts.

## Seen In

- AutoMark
- BlackBook
- BuffHead
- Calling
- Calling, DammazKron, LibRange, Squared
- CastSequence
- CleanUnitFrames
- CleansingBuddy
- DAoCBuff
- DammazKron
- DetauntHelper
- DuffTimer
- EZGuard
- Enemy
- EveryBodyGuard
- GCDsaver
- GuardBot
- HealGrid
- HealHoverAssist
- I HATE YOU THIS MUCH
- LibRange
- LoyalPet
- MarkBuff
- Moth
- NerfedButtons
- Obsidian
- OilTimer
- Pure
- RVAPI_Range
- Rangechecker
- ResHelp
- ShowHealthPercent
- Squared
- TargetInfoRing
- TargetRing
- Targets
- TastyButtons
- TheSeeker
- WoH-Reticle
- ZCurse_Profiler
- followTheLeader
- xHUD

## Related APIs

- [LayoutEditor.RegisterEditCallback](../../window_api/functions/window_LayoutEditor.RegisterEditCallback.md) (HIGH 100/100) - Window Function
- [OnInitialize](../../xml/handlers/handler_OnInitialize.md) (HIGH 88/100) - XML Event
- [OnShutdown](../../xml/handlers/handler_OnShutdown.md) (HIGH 88/100) - XML Event
- [RegisterEventHandler](../../globals/functions/global_RegisterEventHandler.md) (MEDIUM 68/100) - Global Function
- [UnregisterEventHandler](../../globals/functions/global_UnregisterEventHandler.md) (MEDIUM 68/100) - Global Function
- [Start](../../events/game_events/game_event_Start.md) (MEDIUM 43/100) - Game Event
- [Stop](../../events/game_events/game_event_Stop.md) (MEDIUM 43/100) - Game Event

## Used With

- [SystemData.Events.PLAYER_TARGET_EFFECTS_UPDATED](systemdata_SystemData.Events.PLAYER_TARGET_EFFECTS_UPDATED.md) (HIGH 100/100) - SystemData Field
- [RegisterEventHandler](../../globals/functions/global_RegisterEventHandler.md) (MEDIUM 68/100) - Global Function
- [UnregisterEventHandler](../../globals/functions/global_UnregisterEventHandler.md) (MEDIUM 68/100) - Global Function

## Notes

- Observed in contexts: Calling.OnTargetChanged, GroupsInitialize, Init, Initialize, LibRange.OnPlayerTargetUpdated, LoadUnitFrame
