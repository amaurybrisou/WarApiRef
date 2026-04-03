# CreateWindow

- Category: Global Function
- Confidence level: HIGH
- Confidence score: 75/100
- Seen in: 153 addons

## Confidence Assessment

- Level: HIGH

- Score: 75/100

- Rationale: Promoted as HIGH confidence because seen in 4 or more addons, called globally with no local definition, role is consistent across addons.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +10 Argument pattern is consistent: Observed argument positions remain stable.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | AbilityAlert, AbilityNotifier, ActionFraction, ActionPoints, AdvancedPetAssist, AdvancedRenownTrainer, AggroMeter, Asshat |
| Files seen in | AAOTracker.lua, APAGui.lua, AbilityAlert.lua, AbilityNotifier.lua, ActionPoints.lua, AdvancedRenownTraining.lua, AdvancedRenownTrainingImportExport.lua, AggroMeter.lua |
| Namespaces detected | CreateWindow |
| Source kinds | lua_calls |
| Example locations | AbilityAlert: Initialize, AbilityNotifier: Initialize, ActionFraction: Initialize, ActionPoints: Initialize, AdvancedPetAssist: Show, AdvancedRenownTrainer: Initialize |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 342 |
| Global usage count | 342 |
| Local definition count | 0 |
| Documentation references | 0 |
| Initialization flow references | 0 |
| Known engine namespace | no |
| Default UI presence | no |
| Event binding presence | no |
| Observed in XML and Lua | no |
| Consistent role | yes |
| Consistent arguments | yes |
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

## Signature (inferred)

```lua
CreateWindow(windowName, showOnCreate)
```

## Description

Observed creating a top-level XML window from a loaded definition.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| windowName | Observed as a top-level window name. | Observed values: "AAWindow", "APAOptions", "AbHelpWindow" |
| showOnCreate | Observed as a boolean visibility flag. | Observed values: KillTracker_Options.showToggle, MiracleGrow.Settings.showing, false |

## Returns

- Not confidently inferable from contract artifacts alone.

## Side Effects

- Creates or instantiates UI objects from XML-backed definitions.

## Seen In

- AbilityAlert
- AbilityNotifier
- ActionFraction
- ActionPoints
- AdvancedPetAssist
- AdvancedRenownTrainer
- AggroMeter
- Asshat
- Aura
- AutoMark
- BagOMatic
- BankArkel
- BlackBox
- Brizio's Crappy Computer Medic
- BuddyBind
- Busted
- CCTV
- CDown
- CM_ClosetGoblin
- Calling
- ChattyCathy
- Cheeseboard
- CleanUnitFrames
- CleansingBuddy
- Clock
- CoolDownLine
- Countdown
- DAoCBuff
- DPSMeter
- DeepSleep
- Dye Preview
- EA_OpenPartyWindow
- EA_UiDebugTools
- Emojii
- Enemy
- EveryBodyGuard
- FastFriends
- FastInteract
- FlagCap
- GCDTracker
- GetStats
- Group Icons SG
- GroupSpotter
- GuardLine
- GuardList
- GuardRange
- GuildWarden
- HealGrid
- InfoScroller
- JunkDump
- KeyBar
- KillTracker
- Killer
- Kwestor
- LibAddonButton
- LootAlert
- LoyalPet
- MacroIcons
- ManualScenarioRefresh
- Map
- MapMonster
- Mass Refine
- Mech
- MegaphonePlusPlus
- Minmap
- Miracle Grow Remix
- MiracleGrow
- MiracleGrowLight
- MoraleCircle
- MoraleSet
- Moth
- MouseHint
- MyReasons
- NerfedButtons
- NoOverheal
- ObjectInspector
- Paint the leader
- PeaceOut
- Phantom
- PieTracker
- PlayEffectsOn
- PotionBar
- Pure
- QuickWarReport
- RVAPI_ColorDialog
- RVAPI_Range
- RVMOD_3DPortrait
- RVMOD_Manager
- RVMOD_PlayerStatus
- RVMOD_SquaredDistances
- RVMOD_Targets
- RaidMeter
- RandomMount
- Rangechecker
- Refer
- ReliquaryHunter
- ResHelp
- RetAlert
- RoR_SoR
- RoR_debolster
- Rolodex
- Rotation
- RvRContribution
- RvRStats
- RvRStatsTab
- SNT_CASTBAR
- SNT_INFO
- SNT_PANEL
- SOR
- ScenarioStats
- Sequencer
- Shinies
- SimpleXY
- SocialWindow 2.0
- Soloq
- Squared
- Statdoll Remix
- Swift Assist
- Swinger
- TalismanGenie
- TastyButtons
- TaxPayer
- ThankTheTank
- TheSeeker
- ThinkOutLoud
- TidyChat
- TidyRoll
- TokenMachine
- TomeTracker
- Trakario
- TurretScrap
- Twister
- TwisterSet
- Vectors
- WARCommander
- WARRatingBuster
- WSCT
- WTes
- WarBoard_AAOTracker
- WarBoard_FPS
- WarBoard_WarWhisperer
- WhatsCooking
- XpStatus+G
- ZonePOP
- bigger_MacroWindow
- emotes
- followTheLeader
- fpsbox
- minesweep
- nLootLink
- wbLeadHelper
- yAssistHelper
- zMailMod

## Examples

- AbilityAlert: Initialize -> CreateWindow("AAWindow", true)
- AbilityNotifier: Initialize -> CreateWindow("AbHelpWindow", true)
- ActionFraction: Initialize -> CreateWindow(windowName, true)
- ActionPoints: Initialize -> CreateWindow("ActionPointsWindow", true)
- AdvancedPetAssist: Show -> CreateWindow("APAOptions", true)
- AdvancedRenownTrainer: Initialize -> CreateWindow("AdvancedRenownTrainingPresetsWindow", false)

## Related APIs

- [ButtonSetPressedFlag](../../window_api/functions/window_ButtonSetPressedFlag.md) (HIGH 100/100) - Window Function
- [ButtonSetText](../../window_api/functions/window_ButtonSetText.md) (HIGH 100/100) - Window Function
- [ComboBoxSetSelectedMenuItem](../../window_api/functions/window_ComboBoxSetSelectedMenuItem.md) (HIGH 100/100) - Window Function
- [DynamicImageSetTextureDimensions](../../window_api/functions/window_DynamicImageSetTextureDimensions.md) (HIGH 100/100) - Window Function
- [InterfaceCore.GetScale](global_InterfaceCore.GetScale.md) (HIGH 100/100) - Global Function
- [LabelSetText](../../window_api/functions/window_LabelSetText.md) (HIGH 100/100) - Window Function
- [LayoutEditor.RegisterWindow](../../window_api/functions/window_LayoutEditor.RegisterWindow.md) (HIGH 100/100) - Window Function
- [WindowAddAnchor](../../window_api/functions/window_WindowAddAnchor.md) (HIGH 100/100) - Window Function
- [WindowClearAnchors](../../window_api/functions/window_WindowClearAnchors.md) (HIGH 100/100) - Window Function
- [WindowGetAnchor](../../window_api/functions/window_WindowGetAnchor.md) (HIGH 100/100) - Window Function
- [WindowGetDimensions](../../window_api/functions/window_WindowGetDimensions.md) (HIGH 100/100) - Window Function
- [WindowRegisterCoreEventHandler](../../window_api/functions/window_WindowRegisterCoreEventHandler.md) (HIGH 100/100) - Window Function
- [WindowSetDimensions](../../window_api/functions/window_WindowSetDimensions.md) (HIGH 100/100) - Window Function
- [WindowSetId](../../window_api/functions/window_WindowSetId.md) (HIGH 100/100) - Window Function
- [WindowSetLayer](../../window_api/functions/window_WindowSetLayer.md) (HIGH 100/100) - Window Function
- [WindowSetScale](../../window_api/functions/window_WindowSetScale.md) (HIGH 100/100) - Window Function
- [WindowSetShowing](../../window_api/functions/window_WindowSetShowing.md) (HIGH 100/100) - Window Function
- [OnInitialize](../../xml/handlers/handler_OnInitialize.md) (HIGH 88/100) - XML Event
- [DoesWindowExist](global_DoesWindowExist.md) (HIGH 83/100) - Global Function
- [CreateWindowFromTemplate](global_CreateWindowFromTemplate.md) (HIGH 75/100) - Global Function
- [RegisterEventHandler](global_RegisterEventHandler.md) (MEDIUM 68/100) - Global Function

## Used With

- [ButtonSetText](../../window_api/functions/window_ButtonSetText.md) (HIGH 100/100) - Window Function
- [LabelSetText](../../window_api/functions/window_LabelSetText.md) (HIGH 100/100) - Window Function
- [LayoutEditor.RegisterWindow](../../window_api/functions/window_LayoutEditor.RegisterWindow.md) (HIGH 100/100) - Window Function
- [SystemData.Events.INTERFACE_RELOADED](../../systemdata/fields/systemdata_SystemData.Events.INTERFACE_RELOADED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.LOADING_END](../../systemdata/fields/systemdata_SystemData.Events.LOADING_END.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_COMBAT_FLAG_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_COMBAT_FLAG_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.RELOAD_INTERFACE](../../systemdata/fields/systemdata_SystemData.Events.RELOAD_INTERFACE.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.WORLD_OBJ_COMBAT_EVENT](../../systemdata/fields/systemdata_SystemData.Events.WORLD_OBJ_COMBAT_EVENT.md) (HIGH 100/100) - SystemData Field
- [WindowAddAnchor](../../window_api/functions/window_WindowAddAnchor.md) (HIGH 100/100) - Window Function
- [WindowSetShowing](../../window_api/functions/window_WindowSetShowing.md) (HIGH 100/100) - Window Function
- [OnInitialize](../../xml/handlers/handler_OnInitialize.md) (HIGH 88/100) - XML Event
- [LibSlash.RegisterSlashCmd](global_LibSlash.RegisterSlashCmd.md) (HIGH 85/100) - Global Function
- [RegisterEventHandler](global_RegisterEventHandler.md) (MEDIUM 68/100) - Global Function

## Affects

- [SystemData.ChatLogFilters.SHOUT](../../systemdata/fields/systemdata_SystemData.ChatLogFilters.SHOUT.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.ENTER_WORLD](../../systemdata/fields/systemdata_SystemData.Events.ENTER_WORLD.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.INTERFACE_RELOADED](../../systemdata/fields/systemdata_SystemData.Events.INTERFACE_RELOADED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.LOADING_END](../../systemdata/fields/systemdata_SystemData.Events.LOADING_END.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_CAREER_CATEGORY_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_CAREER_CATEGORY_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_COMBAT_FLAG_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_COMBAT_FLAG_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_CUR_ACTION_POINTS_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_CUR_ACTION_POINTS_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.RELOAD_INTERFACE](../../systemdata/fields/systemdata_SystemData.Events.RELOAD_INTERFACE.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.WORLD_OBJ_COMBAT_EVENT](../../systemdata/fields/systemdata_SystemData.Events.WORLD_OBJ_COMBAT_EVENT.md) (HIGH 100/100) - SystemData Field

## Notes

- Canonical entry built from observed call sites, not from engine source or decompiled definitions.
- Advanced return analysis: No strong return evidence observed
