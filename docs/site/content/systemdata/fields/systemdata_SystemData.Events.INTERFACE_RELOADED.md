# SystemData.Events.INTERFACE_RELOADED

- Category: SystemData Field
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 43 addons

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
| Addons seen in | ActionFraction, ArmorGraphicToggle, Asshat, AutoBand, BBars - Mechanic Only, CDown, CMap, CaVES |
| Files seen in | ArmorGraphicToggle.lua, Asshat.lua, AutoBand.lua, BBars.lua, CDown.lua, CMap.lua, ClickSoundSuppressor.lua, Code/Core/Main.lua |
| Namespaces detected | SystemData |
| Source kinds | lua_call |
| Example locations | ArmorGraphicToggle_Disable, ArmorGraphicToggle_Enable, DelayInit, Init, Init_Warden, Initialize |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 23 |
| Global usage count | 23 |
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

SystemData.SystemData.Events.INTERFACE_RELOADED field accessed by 43 addons; commonly found in ArmorGraphicToggle_Disable and ArmorGraphicToggle_Enable, DelayInit, Init, Init_Warden, Initialize, InitializePlayer, OnInit, OnInitialize, OnInterfaceLoaded, OnLoad, OnShutdown, RegisterEventHandlerDK, RegisterEvents, RegisterSelfEvents, Shutdown, ShutdownPlayer, Start, UnregisterEventHandlerDK, UnregisterEvents, init, lua_call, onInit contexts.

## Seen In

- ActionFraction
- ArmorGraphicToggle
- Asshat
- AutoBand
- BBars - Mechanic Only
- CDown
- CMap
- CaVES
- CastSequence
- ClickSoundSuppressor
- CoolDownLine
- DAoCBuff
- DammazKron
- DetauntHelper
- Ding
- Enemy
- FixGit
- GCDsaver
- GroupRange
- GuardBot
- GuildWarden
- I HATE YOU THIS MUCH
- KillTracker
- Map
- MapMonster
- MegaphonePlus
- MegaphonePlusPlus
- MyReasons
- Paint the leader
- PartyCast
- PeaceOut
- PetFixWindow
- ReliquaryHunter
- ResHelp
- RoR_SoR
- RoR_debolster
- Sequencer
- TacticSetNames
- TomeTracker
- VPBreakdown
- Vectors
- ZonePOP
- fpsbox

## Related APIs

- [ButtonSetPressedFlag](../../window_api/functions/window_ButtonSetPressedFlag.md) (HIGH 100/100) - Window Function
- [ButtonSetStayDownFlag](../../window_api/functions/window_ButtonSetStayDownFlag.md) (HIGH 100/100) - Window Function
- [LabelSetFont](../../window_api/functions/window_LabelSetFont.md) (HIGH 100/100) - Window Function
- [LayoutEditor.RegisterEditCallback](../../window_api/functions/window_LayoutEditor.RegisterEditCallback.md) (HIGH 100/100) - Window Function
- [LayoutEditor.RegisterWindow](../../window_api/functions/window_LayoutEditor.RegisterWindow.md) (HIGH 100/100) - Window Function
- [LibSlash.IsSlashCmdRegistered](../../globals/functions/global_LibSlash.IsSlashCmdRegistered.md) (HIGH 100/100) - Global Function
- [LibSlash.RegisterWSlashCmd](../../globals/functions/global_LibSlash.RegisterWSlashCmd.md) (HIGH 100/100) - Global Function
- [LibSlash.UnregisterSlashCmd](../../globals/functions/global_LibSlash.UnregisterSlashCmd.md) (HIGH 100/100) - Global Function
- [WindowRegisterCoreEventHandler](../../window_api/functions/window_WindowRegisterCoreEventHandler.md) (HIGH 100/100) - Window Function
- [WindowRegisterEventHandler](../../window_api/functions/window_WindowRegisterEventHandler.md) (HIGH 100/100) - Window Function
- [WindowSetLayer](../../window_api/functions/window_WindowSetLayer.md) (HIGH 100/100) - Window Function
- [WindowSetShowing](../../window_api/functions/window_WindowSetShowing.md) (HIGH 100/100) - Window Function
- [OnInitialize](../../xml/handlers/handler_OnInitialize.md) (HIGH 88/100) - XML Event
- [OnShutdown](../../xml/handlers/handler_OnShutdown.md) (HIGH 88/100) - XML Event
- [CreateWindow](../../globals/functions/global_CreateWindow.md) (HIGH 75/100) - Global Function
- [wstring.format](../../globals/functions/global_wstring.format.md) (HIGH 75/100) - Global Function
- [UnregisterEventHandler](../../globals/functions/global_UnregisterEventHandler.md) (MEDIUM 68/100) - Global Function
- [Start](../../events/game_events/game_event_Start.md) (MEDIUM 43/100) - Game Event

## Used With

- [GameData.Account.ServerName](../../gamedata/fields/gamedata_GameData.Account.ServerName.md) (HIGH 100/100) - GameData Field
- [LibSlash.IsSlashCmdRegistered](../../globals/functions/global_LibSlash.IsSlashCmdRegistered.md) (HIGH 100/100) - Global Function
- [LibSlash.RegisterWSlashCmd](../../globals/functions/global_LibSlash.RegisterWSlashCmd.md) (HIGH 100/100) - Global Function
- [SystemData.ChatLogFilters.SHOUT](systemdata_SystemData.ChatLogFilters.SHOUT.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.BATTLEGROUP_UPDATED](systemdata_SystemData.Events.BATTLEGROUP_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.CHAT_TEXT_ARRIVED](systemdata_SystemData.Events.CHAT_TEXT_ARRIVED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.ENTER_WORLD](systemdata_SystemData.Events.ENTER_WORLD.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.LOADING_END](systemdata_SystemData.Events.LOADING_END.md) (HIGH 100/100) - SystemData Field
- [CreateWindow](../../globals/functions/global_CreateWindow.md) (HIGH 75/100) - Global Function
- [RegisterEventHandler](../../globals/functions/global_RegisterEventHandler.md) (MEDIUM 68/100) - Global Function
- [UnregisterEventHandler](../../globals/functions/global_UnregisterEventHandler.md) (MEDIUM 68/100) - Global Function
- [Start](../../events/game_events/game_event_Start.md) (MEDIUM 43/100) - Game Event

## Notes

- Observed in contexts: ArmorGraphicToggle_Disable, ArmorGraphicToggle_Enable, DelayInit, Init, Init_Warden, Initialize
