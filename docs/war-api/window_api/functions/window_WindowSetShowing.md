# WindowSetShowing

- Category: Window Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 236 addons

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 110

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, seen in 4 or more addons, matches a known engine namespace.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- -15 Conflicting signatures across usages: Observed arity or argument shape conflicts across usages.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | Ace, ActionBarHide, ActionFraction, AdjustTheTip, AdvancedPetAssist, AdvancedRenownTrainer, AggroMeter, Amethyst |
| Files seen in | AAOTracker.lua, APAGui.lua, APAGuiHUD.lua, ActionBarHide.lua, AdjustTheTip.lua, AdvancedPetAssist.lua, AdvancedRenownTraining.lua, AdvancedRenownTrainingImportExport.lua |
| Namespaces detected | WindowSetShowing |
| Source kinds | lua_calls |
| Example locations | Ace: Hide, Ace: Show, ActionBarHide: Hide, ActionBarHide: OnLoad, ActionBarHide: Show, ActionFraction: Initialize |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 3234 |
| Global usage count | 3234 |
| Local definition count | 0 |
| Documentation references | 0 |
| Initialization flow references | 0 |
| Known engine namespace | yes |
| Default UI presence | yes |
| Event binding presence | no |
| Observed in XML and Lua | no |
| Consistent role | yes |
| Consistent arguments | no |
| Consistent returns | no |
| Slash command presence | no |
| Weak usage only | no |
| Project-specific name | no |
| Placeholder or computed name | no |
| Conflicting signatures | yes |
| Conflicting roles | no |
| Wrapper likely | no |
| Never outside local graph | no |
| Local helper only | no |

## Signature (inferred)

```lua
WindowSetShowing(windowName, arg2)
```

## Description

Observed mutating runtime window state or presentation.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| windowName | Observed as a target window name. | Observed values: "AGTSettingsWindow", "APAInstantOnlyHUD", "APAKitingHUD" |
| arg2 | Observed as a boolean toggle. | Observed values: "ST"..SNTCAST_DB.PRESET, #ResHelp.Table>0, (GameData.Player.rvrPermaFlagged or GameData.Player.rvrZoneFlagged)and CMapWindow.VisSettings[CFG_RVRINDI].value |

## Returns

- Not confidently inferable from contract artifacts alone.

## Side Effects

- Mutates engine-visible UI state or presentation.

## Seen In

- Ace
- ActionBarHide
- ActionFraction
- AdjustTheTip
- AdvancedPetAssist
- AdvancedRenownTrainer
- AggroMeter
- Amethyst
- AnywhereTrainerAdditions
- ArmorGraphicToggle
- Atlas
- AuctionStats
- Aura
- AutoBand
- AutoMark
- AutoSalvage
- BBars - Mechanic Only
- BankArkel
- BarText (Influence)
- BlackBook
- BlackBox
- Bloody Mess
- Brizio's Crappy Computer Medic
- BuffHead
- Busted
- CCTV
- CDown
- CM_ClosetGoblin
- CMap
- CaVES
- Calling
- CastSequence
- CharacterScreenTabFix
- ChattyCathy
- Cheeseboard
- CleanCastbar
- CleanUnitFrames
- CombatTextNames
- CoolDownLine
- Countdown
- Crafting Info Tooltip
- CraftingWillard
- Crusher
- DAoCBuff
- DPSMeter
- DaemonAssist
- DammazKron
- Dascore
- DeepSleep
- DetauntHelper
- DuffTimer
- Dye Preview
- EA_LoadingScreen
- EA_OpenPartyWindow
- EA_ScenarioGroupWindow
- EA_ThreePartBar
- EA_UiDebugTools
- EA_UiModWindow
- EZGuard
- Effigy
- Emojii
- Enemy
- EveryBodyGuard
- FastFriends
- FastInteract
- FixGit
- FlagCap
- FozAuction
- GCDTracker
- GCDsaver
- GetStats
- Group Icons
- Group Icons SG
- GroupRange
- GroupSpotter
- GuardBot
- GuardLine
- GuardList
- GuardRange
- GuildWarden
- Hopper
- I HATE YOU THIS MUCH
- InfoScroller
- JunkDump
- KeyBar
- Keyset
- KeysetMonsterPlay
- KillTracker
- Killer
- Kwestor
- LibAddonButton
- LibGroup
- LibRange
- LibSlash
- LibWBToggler
- LootAlert
- LoyalPet
- MacroIcons
- ManualScenarioRefresh
- Map
- MapMonster
- MapPin
- MarkBuff
- Mass Refine
- MegaphonePlusPlus
- MiniMapMonster
- Minmap
- Miracle Grow Remix
- MiracleGrow
- MiracleGrowLight
- MoraleBG
- MoraleSet
- Moth
- Motion
- MouseHint
- NAMBLA
- NaturalLog
- NerfedButtons
- NoOverheal
- NoUselessMods-Assist
- ObjectInspector
- Obsidian
- OverheadFonts
- Paint the leader
- PartyCast
- PeaceOut
- PetFixWindow
- PieTracker
- Pocket Palette
- PotionBar
- Pure
- Pure Careerbar
- Queue Queuer
- QuickTacticSwitch
- QuickWarReport
- RO-Style Combat Text
- RVAPI_ColorDialog
- RVMOD_3DPortrait
- RaidMeter
- RandomMount
- Rangechecker
- RealmStatus
- Refer
- ReliquaryHunter
- ResHelp
- RoR_SoR
- RoR_debolster
- Rolodex
- Rotation
- RvRStats
- RvRStatsTab
- SNT Info [TDPS]
- SNT_CASTBAR
- SNT_INFO
- SNT_PANEL
- SOR
- ScenarioStats
- Sequencer
- SessionRPs
- Shinies
- SimpleCombatText
- SimpleXY
- SocialWindow 2.0
- Soloq
- Squared
- SquaredClick
- SquaredHDIndicator
- SquaredHotIndicators
- Statdoll
- Statdoll Light
- Statdoll Remix
- Swift Assist
- Swinger
- TacticSetNames
- TalismanGenie
- TargetInfoRing
- TargetRing
- TastyButtons
- TaxPayer
- TexturedButtons
- ThankTheTank
- ThinkOutLoud
- TidyChat
- TidyQueue
- TidyRoll
- Tokens
- Tome Titan
- TomeTracker
- Tortall_DPS
- Trakario
- TurretRange
- TurretScrap
- Twister
- TwisterSet
- VPBreakdown
- Vectors
- VerticalMorale
- VerticalTactics
- WARCommander
- WARRatingBuster
- WBStutterLess
- WSCT
- WTes
- WarBoard
- WarBoard_AAOTracker
- WarBoard_Loc
- WarBoard_Menu
- WarBoard_TogglerRankedLeaderboard
- WarBoard_TogglerWARCommander
- WarBoard_WarWhisperer
- WarTriage
- Wargames
- WhatsCooking
- WhoHealedMe
- Wikki's Cooldown Bar
- Wikki's Cooldown Pulse
- WoH-Reticle
- XpStatus+G
- ZCurse_Profiler
- ZonePOP
- alertMod
- bigger_MacroWindow
- compass
- emotes
- followTheLeader
- minesweep
- nLootLink
- nRarity
- scenarioInfo
- scnoload
- talisman-monitor
- wbLeadHelper
- xHUD
- xPanels
- yAssistHelper
- zMailMod

## Examples

- Ace: Hide -> WindowSetShowing(self.name, false)
- Ace: Show -> WindowSetShowing(self.name, true)
- ActionBarHide: Hide -> WindowSetShowing(self.name, false)
- ActionBarHide: OnLoad -> WindowSetShowing("ActionBarLockToggler", false)
- ActionBarHide: Show -> WindowSetShowing(self.name, true)
- ActionFraction: Initialize -> WindowSetShowing("ActionFractionWindowContextColorCodeCurrentAP", false)

## Related APIs

- [OnLButtonUp](../../xml/handlers/handler_OnLButtonUp.md) (HIGH 100/100) - XML Event
- [OnHidden](../../xml/handlers/handler_OnHidden.md) (HIGH 88/100) - XML Event
- [OnInitialize](../../xml/handlers/handler_OnInitialize.md) (HIGH 88/100) - XML Event
- [OnKeyEscape](../../xml/handlers/handler_OnKeyEscape.md) (HIGH 88/100) - XML Event
- [OnMouseOver](../../xml/handlers/handler_OnMouseOver.md) (HIGH 88/100) - XML Event
- [OnMouseOverEnd](../../xml/handlers/handler_OnMouseOverEnd.md) (HIGH 88/100) - XML Event
- [OnRButtonUp](../../xml/handlers/handler_OnRButtonUp.md) (HIGH 88/100) - XML Event
- [OnShown](../../xml/handlers/handler_OnShown.md) (HIGH 88/100) - XML Event
- [OnUpdate](../../xml/handlers/handler_OnUpdate.md) (HIGH 88/100) - XML Event
- [DoesWindowExist](../../globals/functions/global_DoesWindowExist.md) (HIGH 83/100) - Global Function
- [CreateWindow](../../globals/functions/global_CreateWindow.md) (HIGH 75/100) - Global Function
- [Stop](../../events/game_events/game_event_Stop.md) (MEDIUM 43/100) - Game Event

## Used With

- [ButtonSetDisabledFlag](window_ButtonSetDisabledFlag.md) (HIGH 100/100) - Window Function
- [ButtonSetPressedFlag](window_ButtonSetPressedFlag.md) (HIGH 100/100) - Window Function
- [ButtonSetText](window_ButtonSetText.md) (HIGH 100/100) - Window Function
- [ComboBoxGetSelectedText](window_ComboBoxGetSelectedText.md) (HIGH 100/100) - Window Function
- [Cursor.Clear](../../globals/functions/global_Cursor.Clear.md) (HIGH 100/100) - Global Function
- [DynamicImageSetTexture](window_DynamicImageSetTexture.md) (HIGH 100/100) - Window Function
- [EA_ChatWindow.UpdateSocialWindowButton](../../globals/functions/global_EA_ChatWindow.UpdateSocialWindowButton.md) (HIGH 100/100) - Global Function
- [Icons.GetCareerIconIDFromCareerLine](../../globals/functions/global_Icons.GetCareerIconIDFromCareerLine.md) (HIGH 100/100) - Global Function
- [LabelSetText](window_LabelSetText.md) (HIGH 100/100) - Window Function
- [LayoutEditor.RegisterWindow](window_LayoutEditor.RegisterWindow.md) (HIGH 100/100) - Window Function
- [TextEditBoxGetText](window_TextEditBoxGetText.md) (HIGH 100/100) - Window Function
- [TextEditBoxSelectAll](window_TextEditBoxSelectAll.md) (HIGH 100/100) - Window Function
- [TextEditBoxSetMaxChars](window_TextEditBoxSetMaxChars.md) (HIGH 100/100) - Window Function
- [TextEditBoxSetText](window_TextEditBoxSetText.md) (HIGH 100/100) - Window Function
- [WindowAddAnchor](window_WindowAddAnchor.md) (HIGH 100/100) - Window Function
- [WindowAssignFocus](window_WindowAssignFocus.md) (HIGH 100/100) - Window Function
- [WindowClearAnchors](window_WindowClearAnchors.md) (HIGH 100/100) - Window Function
- [WindowGetAlpha](window_WindowGetAlpha.md) (HIGH 100/100) - Window Function
- [WindowGetDimensions](window_WindowGetDimensions.md) (HIGH 100/100) - Window Function
- [WindowGetScale](window_WindowGetScale.md) (HIGH 100/100) - Window Function
- [WindowGetScreenPosition](window_WindowGetScreenPosition.md) (HIGH 100/100) - Window Function
- [WindowGetShowing](window_WindowGetShowing.md) (HIGH 100/100) - Window Function
- [WindowSetDimensions](window_WindowSetDimensions.md) (HIGH 100/100) - Window Function
- [WindowSetHandleInput](window_WindowSetHandleInput.md) (HIGH 100/100) - Window Function
- [WindowSetScale](window_WindowSetScale.md) (HIGH 100/100) - Window Function
- [WindowSetTintColor](window_WindowSetTintColor.md) (HIGH 100/100) - Window Function
- [WindowStartAlphaAnimation](window_WindowStartAlphaAnimation.md) (HIGH 100/100) - Window Function
- [WindowStopAlphaAnimation](window_WindowStopAlphaAnimation.md) (HIGH 100/100) - Window Function
- [wstring.match](../../globals/functions/global_wstring.match.md) (HIGH 100/100) - Global Function
- [OnHidden](../../xml/handlers/handler_OnHidden.md) (HIGH 88/100) - XML Event
- [OnInitialize](../../xml/handlers/handler_OnInitialize.md) (HIGH 88/100) - XML Event
- [OnShown](../../xml/handlers/handler_OnShown.md) (HIGH 88/100) - XML Event
- [OnUpdate](../../xml/handlers/handler_OnUpdate.md) (HIGH 88/100) - XML Event
- [LibSlash.RegisterSlashCmd](../../globals/functions/global_LibSlash.RegisterSlashCmd.md) (HIGH 85/100) - Global Function
- [DoesWindowExist](../../globals/functions/global_DoesWindowExist.md) (HIGH 83/100) - Global Function
- [GetIconData](../../globals/functions/global_GetIconData.md) (HIGH 83/100) - Global Function
- [AlertTextWindow.AddLine](../../globals/functions/global_AlertTextWindow.AddLine.md) (HIGH 75/100) - Global Function
- [CreateWindow](../../globals/functions/global_CreateWindow.md) (HIGH 75/100) - Global Function
- [RegisterEventHandler](../../globals/functions/global_RegisterEventHandler.md) (MEDIUM 68/100) - Global Function

## Affects

- [SystemData.ChatLogFilters.SHOUT](../../systemdata/fields/systemdata_SystemData.ChatLogFilters.SHOUT.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.ENTER_WORLD](../../systemdata/fields/systemdata_SystemData.Events.ENTER_WORLD.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.INTERFACE_RELOADED](../../systemdata/fields/systemdata_SystemData.Events.INTERFACE_RELOADED.md) (HIGH 100/100) - SystemData Field

## Notes

- Advanced return analysis: No strong return evidence observed
