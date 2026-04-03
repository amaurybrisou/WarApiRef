# SystemData.Events.LOADING_END

- Category: SystemData Field
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 106 addons

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
| Addons seen in | ActionBarHide, ActionFraction, AdvancedPetAssist, AdvancedRenownTrainer, ArmorGraphicToggle, Asshat, Atlas, Aura |
| Files seen in | ActionBarHide.lua, AdvancedPetAssist.lua, AdvancedRenownTraining.lua, ArmorGraphicToggle.lua, Asshat.lua, BagOMatic.lua, CDown.lua, CMap.lua |
| Namespaces detected | SystemData |
| Source kinds | event_page, event_registration, lua_call |
| Example locations | ArmorGraphicToggle_Disable, ArmorGraphicToggle_Enable, DelayInit, DisableBars, EnableBars, Init |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 49 |
| Global usage count | 49 |
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

SystemData.SystemData.Events.LOADING_END field accessed by 106 addons; commonly found in ArmorGraphicToggle_Disable and ArmorGraphicToggle_Enable, DelayInit, DisableBars, EnableBars, Init, InitComplete, Init_Warden, Initialise, Initialize, InitializePlayer, LOADING_END, LibRange.OnLoad, LoadUnitFrame, NAMBLA_Initialize, OnInit, OnInitialize, OnLoad, OnLoadCastbar, OnLoadingEnd, OnShutdown, RegisterCoreEvents, RegisterEventHandlerDK, RegisterEventHandlers, RegisterLoadingEnd, RegisterStateInfoForCastbar, RegisterStateInfoForPlayer, RegisterStateInfoForTargets, Shutdown, ShutdownPlayer, Squared.OnLoad, SquaredRangeFading.LoadingEnd, SystemData.Events.LOADING_END, Unload, UnloadUnitFrame, UnregisterCoreEvents, UnregisterEventHandlerDK, UnregisterEventHandlers, UnregisterLoadingEnd, ZoneChangeInit, ZoneChangeShutdown, _ScenarioInfoEnabledChanged, entry_point, event_page, event_registration, init, lua_call, onInit, onInitialize contexts.

## Seen In

- ActionBarHide
- ActionFraction
- AdvancedPetAssist
- AdvancedRenownTrainer
- ArmorGraphicToggle
- Asshat
- Atlas
- Aura
- BagOMatic
- BuffHead
- CDown
- CM_ClosetGoblin
- CMap
- CastSequence
- CleanUnitFrames
- CleansingBuddy
- ClickSoundSuppressor
- CoolDownLine
- Crusher
- DammazKron
- Default Tactic Set
- Ding
- EA_LoadingScreen
- Effigy
- Enemy
- FastFriends
- GroupRange
- GuardBot
- GuildWarden
- HammerTime
- HealGrid
- HealHoverAssist
- Hopper
- Keyset
- KeysetMonsterPlay
- KillTracker
- Killer
- Kwestor
- Lib RuString
- LibGroup
- LibRange
- LibRange, Squared
- LibWBToggler
- Map
- MapMonster
- MarkBuff
- MegaphonePlusPlus
- Minmap
- Miracle Grow Remix
- MiracleGrow
- MiracleGrowLight
- MoraleBG
- Motion
- MyReasons
- NAMBLA
- NaturalLog
- NerfedButtons
- Obsidian
- Paint the leader
- PeaceOut
- PetFixWindow
- Phantom
- PlanB
- Pocket Palette
- Pure
- Pure Careerbar
- QuickNameActions+
- RVAPI_Range
- RVMOD_3DPortrait
- RVMOD_PlayerStatus
- RealmStatus
- RoR_SoR
- RvRContribution
- SNT_PANEL
- SOR
- SessionRPs
- Soloq
- Squared
- SquaredHDIndicator
- SquaredHotIndicators
- Statdoll Remix
- TexturedButtons
- TidyChat
- TidyQueue
- TidyRoll
- TimeToDie
- TokenMachine
- Tome Titan
- TomeTracker
- TurretRange
- Twister
- TwisterSet
- VPBreakdown
- Vectors
- WTes
- WarBoard
- WarBoard_TaliMon
- WarBoard_WarWhisperer
- WhoHealedMe
- Wikki's Cooldown Bar
- Wikki's Cooldown Pulse
- fpsbox
- nLootLink
- scnoload
- wbLeadHelper
- xHUD

## Related APIs

- [ButtonSetStayDownFlag](../../window_api/functions/window_ButtonSetStayDownFlag.md) (HIGH 100/100) - Window Function
- [LayoutEditor.RegisterEditCallback](../../window_api/functions/window_LayoutEditor.RegisterEditCallback.md) (HIGH 100/100) - Window Function
- [LibSlash.IsSlashCmdRegistered](../../globals/functions/global_LibSlash.IsSlashCmdRegistered.md) (HIGH 100/100) - Global Function
- [LibSlash.UnregisterSlashCmd](../../globals/functions/global_LibSlash.UnregisterSlashCmd.md) (HIGH 100/100) - Global Function
- [WindowRegisterEventHandler](../../window_api/functions/window_WindowRegisterEventHandler.md) (HIGH 100/100) - Window Function
- [WindowSetLayer](../../window_api/functions/window_WindowSetLayer.md) (HIGH 100/100) - Window Function
- [WindowUnregisterEventHandler](../../window_api/functions/window_WindowUnregisterEventHandler.md) (HIGH 100/100) - Window Function
- [EA_Window_InteractionRenownTraining.GetPointsAvailable](../../globals/functions/global_EA_Window_InteractionRenownTraining.GetPointsAvailable.md) (HIGH 98/100) - Global Function
- [EA_Window_InteractionRenownTraining.GetPointsSpent](../../globals/functions/global_EA_Window_InteractionRenownTraining.GetPointsSpent.md) (HIGH 98/100) - Global Function
- [WindowUnRegisterEventHandler](../../window_api/functions/window_WindowUnRegisterEventHandler.md) (HIGH 98/100) - Window Function
- [OnInitialize](../../xml/handlers/handler_OnInitialize.md) (HIGH 88/100) - XML Event
- [OnShutdown](../../xml/handlers/handler_OnShutdown.md) (HIGH 88/100) - XML Event
- [CreateWindow](../../globals/functions/global_CreateWindow.md) (HIGH 75/100) - Global Function
- [RegisterEventHandler](../../globals/functions/global_RegisterEventHandler.md) (MEDIUM 68/100) - Global Function
- [UnregisterEventHandler](../../globals/functions/global_UnregisterEventHandler.md) (MEDIUM 68/100) - Global Function

## Used With

- [GameData.Account.ServerName](../../gamedata/fields/gamedata_GameData.Account.ServerName.md) (HIGH 100/100) - GameData Field
- [SystemData.Events.ALL_MODULES_INITIALIZED](systemdata_SystemData.Events.ALL_MODULES_INITIALIZED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.INTERFACE_RELOADED](systemdata_SystemData.Events.INTERFACE_RELOADED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_AGRO_MODE_UPDATED](systemdata_SystemData.Events.PLAYER_AGRO_MODE_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_CUR_ACTION_POINTS_UPDATED](systemdata_SystemData.Events.PLAYER_CUR_ACTION_POINTS_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_CUR_HIT_POINTS_UPDATED](systemdata_SystemData.Events.PLAYER_CUR_HIT_POINTS_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_MAX_ACTION_POINTS_UPDATED](systemdata_SystemData.Events.PLAYER_MAX_ACTION_POINTS_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.RELOAD_INTERFACE](systemdata_SystemData.Events.RELOAD_INTERFACE.md) (HIGH 100/100) - SystemData Field
- [OnInitialize](../../xml/handlers/handler_OnInitialize.md) (HIGH 88/100) - XML Event
- [CreateWindow](../../globals/functions/global_CreateWindow.md) (HIGH 75/100) - Global Function
- [RegisterEventHandler](../../globals/functions/global_RegisterEventHandler.md) (MEDIUM 68/100) - Global Function

## Notes

- Observed in contexts: ArmorGraphicToggle_Disable, ArmorGraphicToggle_Enable, DelayInit, DisableBars, EnableBars, Init
