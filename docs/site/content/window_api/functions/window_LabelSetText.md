# LabelSetText

- Category: Window Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 184 addons

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 135

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, seen in 4 or more addons, matches a known engine namespace.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +10 Argument pattern is consistent: Observed argument positions remain stable.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | AbilityAlert, Ace, ActionBarHide, ActionFraction, AdjustTheTip, AdvancedPetAssist, AdvancedRenownTrainer, AggroMeter |
| Files seen in | APAGui.lua, APAGuiHUD.lua, AbilityAlert.lua, AdjustTheTip.lua, AdvancedRenownTraining.lua, AdvancedRenownTrainingImportExport.lua, AggroMeter.lua, AuctionAssist.lua |
| Namespaces detected | LabelSetText |
| Source kinds | lua_calls |
| Example locations | AbilityAlert: Display, Ace: Clear, Ace: SetText, ActionBarHide: Clear, ActionBarHide: SetText, ActionFraction: Initialize |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 3778 |
| Global usage count | 3778 |
| Local definition count | 0 |
| Documentation references | 0 |
| Initialization flow references | 0 |
| Known engine namespace | yes |
| Default UI presence | yes |
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
LabelSetText(windowName, text)
```

## Description

Observed updating label text or label styling on existing controls.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| windowName | Observed as a target control name. | Observed values: "10minTimeValue", "10secTimeValue", "1minTimeValue" |
| text | Observed as a text or wstring payload. | Observed values: #MailWindowTabAuction.listData..L " "..zL["MESSAGE"], #MailWindowTabAuction.listData..L " "..zL["MESSAGES"], #MailWindowTabInbox.listData..L " "..zL["MESSAGE"] |

## Returns

- Not confidently inferable from contract artifacts alone.

## Side Effects

- Mutates engine-visible UI state or presentation.

## Seen In

- AbilityAlert
- Ace
- ActionBarHide
- ActionFraction
- AdjustTheTip
- AdvancedPetAssist
- AdvancedRenownTrainer
- AggroMeter
- Amethyst
- Atlas
- AuctionStats
- Aura
- AutoBand
- AutoSalvage
- BankArkel
- BlackBook
- BlackBox
- Bloody Mess
- Brizio's Crappy Computer Medic
- BuddyBind
- BuffHead
- Busted
- CDown
- CM_ClosetGoblin
- CaVES
- Calling
- CastSequence
- ChattyCathy
- CleanUnitFrames
- CombatTextNames
- CoolDownLine
- Countdown
- Crafting Info Tooltip
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
- EA_ScenarioGroupWindow
- EA_UiDebugTools
- EZCraftX
- EZGuard
- Effigy
- Emojii
- Enemy
- EveryBodyGuard
- FastFriends
- FastInteract
- FlagCap
- GCDTracker
- GCDsaver
- GetStats
- Group Icons SG
- GroupRange
- GroupSpotter
- GuildWarden
- HealGrid
- Hopper
- InfoScroller
- JunkDump
- KeyBar
- Keyset
- KeysetMonsterPlay
- Killer
- Kwestor
- LibAddonButton
- LibGroup
- LibWBToggler
- LootAlert
- Map
- MapMonster
- MarkBuff
- Mass Refine
- Mech
- MegaphonePlusPlus
- Minmap
- Miracle Grow Remix
- MiracleGrow
- MiracleGrowLight
- MoraleCircle
- Moth
- Motion
- NaturalLog
- NerfedButtons
- NoOverheal
- Obsidian
- OverheadFonts
- PartyCast
- PeaceOut
- Phantom
- Pocket Palette
- PotionBar
- Pure
- Pure Careerbar
- QuickTacticSwitch
- QuickWarReport
- RO-Style Combat Text
- RVAPI_ColorDialog
- RVAPI_Range
- RVMOD_3DPortrait
- RVMOD_Manager
- RVMOD_PlayerStatus
- RVMOD_SquaredDistances
- RVMOD_Targets
- RandomMount
- Rangechecker
- RealmStatus
- Refer
- ReliquaryHunter
- ResHelp
- RoR_SoR
- RoR_debolster
- RvRContribution
- RvRStats
- RvRStatsTab
- SNT_BUTTONS
- SNT_CASTBAR
- SNT_PANEL
- SOR
- ScenarioStats
- Sequencer
- SessionRPs
- Shinies
- ShowHealthPercent
- Soloq
- Squared
- SquaredClick
- Statdoll
- Statdoll Light
- Statdoll Remix
- Swift Assist
- Swinger
- TacticSetNames
- TalismanGenie
- TargetInfoRing
- TargetRing
- Targets
- TastyButtons
- TexturedButtons
- ThankTheTank
- ThinkOutLoud
- TidyChat
- TidyQueue
- TidyRoll
- TokenMachine
- Tokens
- Tome Titan
- TomeTracker
- Trakario
- TurretRange
- Twister
- VPBreakdown
- Vectors
- WARCommander
- WARRatingBuster
- WSCT
- WarBoard_Loc
- WarBoard_Session
- WarBoard_TDPS
- WarTriage
- WhatsCooking
- WhoHealedMe
- Wikki's Cooldown Bar
- Wikki's Cooldown Pulse
- WoH-Reticle
- XpStatus+G
- ZonePOP
- alertMod
- emotes
- minesweep
- nLootLink
- scenarioInfo
- talisman-monitor
- warboard_TogglerWarBuilder
- wbLeadHelper
- xHUD
- xPanels
- yAssistHelper
- zMailMod

## Examples

- AbilityAlert: Display -> LabelSetText("AAText", msg)
- Ace: Clear -> LabelSetText(self.name, L "")
- Ace: SetText -> LabelSetText(self.name, towstring(text))
- ActionBarHide: Clear -> LabelSetText(self.name, L "")
- ActionBarHide: SetText -> LabelSetText(self.name, towstring(text))
- ActionFraction: Initialize -> LabelSetText("ActionFractionWindowContextColorCodeCurrentAPCheckBoxLabel", LOC_TEXT.CURRENTAP_COLOR_CODING)

## Related APIs

- [OnLButtonUp](../../xml/handlers/handler_OnLButtonUp.md) (HIGH 100/100) - XML Event
- [OnInitialize](../../xml/handlers/handler_OnInitialize.md) (HIGH 88/100) - XML Event
- [OnMouseOver](../../xml/handlers/handler_OnMouseOver.md) (HIGH 88/100) - XML Event
- [OnMouseOverEnd](../../xml/handlers/handler_OnMouseOverEnd.md) (HIGH 88/100) - XML Event
- [OnShown](../../xml/handlers/handler_OnShown.md) (HIGH 88/100) - XML Event
- [OnSlide](../../xml/handlers/handler_OnSlide.md) (HIGH 88/100) - XML Event
- [OnUpdate](../../xml/handlers/handler_OnUpdate.md) (HIGH 88/100) - XML Event
- [DoesWindowExist](../../globals/functions/global_DoesWindowExist.md) (HIGH 83/100) - Global Function
- [CreateWindow](../../globals/functions/global_CreateWindow.md) (HIGH 75/100) - Global Function

## Used With

- [ButtonSetPressedFlag](window_ButtonSetPressedFlag.md) (HIGH 100/100) - Window Function
- [ButtonSetText](window_ButtonSetText.md) (HIGH 100/100) - Window Function
- [ComboBoxAddMenuItem](window_ComboBoxAddMenuItem.md) (HIGH 100/100) - Window Function
- [ComboBoxSetSelectedMenuItem](window_ComboBoxSetSelectedMenuItem.md) (HIGH 100/100) - Window Function
- [DynamicImageSetTexture](window_DynamicImageSetTexture.md) (HIGH 100/100) - Window Function
- [DynamicImageSetTextureScale](window_DynamicImageSetTextureScale.md) (HIGH 100/100) - Window Function
- [GameData.PlayerActions.SET_TARGET](../../gamedata/fields/gamedata_GameData.PlayerActions.SET_TARGET.md) (HIGH 100/100) - GameData Field
- [Icons.GetCareerIconIDFromCareerLine](../../globals/functions/global_Icons.GetCareerIconIDFromCareerLine.md) (HIGH 100/100) - Global Function
- [LabelGetTextDimensions](window_LabelGetTextDimensions.md) (HIGH 100/100) - Window Function
- [LabelSetTextColor](window_LabelSetTextColor.md) (HIGH 100/100) - Window Function
- [TextEditBoxSetText](window_TextEditBoxSetText.md) (HIGH 100/100) - Window Function
- [WindowAddAnchor](window_WindowAddAnchor.md) (HIGH 100/100) - Window Function
- [WindowGetScreenPosition](window_WindowGetScreenPosition.md) (HIGH 100/100) - Window Function
- [WindowSetAlpha](window_WindowSetAlpha.md) (HIGH 100/100) - Window Function
- [WindowSetShowing](window_WindowSetShowing.md) (HIGH 100/100) - Window Function
- [WindowSetTintColor](window_WindowSetTintColor.md) (HIGH 100/100) - Window Function
- [EA_Window_InteractionRenownTraining.GetPointsAvailable](../../globals/functions/global_EA_Window_InteractionRenownTraining.GetPointsAvailable.md) (HIGH 98/100) - Global Function
- [EA_Window_InteractionRenownTraining.GetPointsSpent](../../globals/functions/global_EA_Window_InteractionRenownTraining.GetPointsSpent.md) (HIGH 98/100) - Global Function
- [OnUpdate](../../xml/handlers/handler_OnUpdate.md) (HIGH 88/100) - XML Event
- [DoesWindowExist](../../globals/functions/global_DoesWindowExist.md) (HIGH 83/100) - Global Function
- [GetIconData](../../globals/functions/global_GetIconData.md) (HIGH 83/100) - Global Function
- [CreateWindow](../../globals/functions/global_CreateWindow.md) (HIGH 75/100) - Global Function
- [CreateWindowFromTemplate](../../globals/functions/global_CreateWindowFromTemplate.md) (HIGH 75/100) - Global Function
- [wstring.sub](../../globals/functions/global_wstring.sub.md) (HIGH 75/100) - Global Function
- [RegisterEventHandler](../../globals/functions/global_RegisterEventHandler.md) (MEDIUM 68/100) - Global Function

## Affects

- [SystemData.Events.LOADING_END](../../systemdata/fields/systemdata_SystemData.Events.LOADING_END.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_AGRO_MODE_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_AGRO_MODE_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_CUR_ACTION_POINTS_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_CUR_ACTION_POINTS_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_CUR_HIT_POINTS_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_CUR_HIT_POINTS_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_HEALTH_FADE_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_HEALTH_FADE_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_MAX_ACTION_POINTS_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_MAX_ACTION_POINTS_UPDATED.md) (HIGH 100/100) - SystemData Field

## Notes

- Advanced return analysis: No strong return evidence observed
