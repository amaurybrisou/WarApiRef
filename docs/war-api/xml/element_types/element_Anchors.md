# Anchors

- Category: XML Element Type
- Confidence level: HIGH
- Confidence score: 100/100

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 110

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, seen in 4 or more addons, used directly in xml handler attributes.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +30 Used directly in XML handler attributes: XML exposure suggests an engine-level contract.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | AbilityAlert, AbilityNotifier, ActionPoints, AdjustTheTip, AdvancedPetAssist, AdvancedRenownTrainer, AggroMeter, AnywhereTrainer |
| Files seen in | Bars/HealGridProgressBar.xml, Code/Assist/Assist.xml, Code/Assist/AssistConfiguration.xml, Code/CombatLog/CombatLogConfiguration.xml, Code/CombatLog/CombatLogEpsWindow.xml, Code/CombatLog/CombatLogIDS.xml, Code/CombatLog/CombatLogSnapshotWindow.xml, Code/CombatLog/CombatLogStatsWindow.xml |
| Namespaces detected | Anchors |
| Source kinds | xml_frames |
| Example locations | AbilityAlert: AAWindow, AbilityNotifier: AbHelpWindow, ActionPoints: ActionPointsWindow, ActionPoints: ActionPointsWindowText, AdjustTheTip: AdjustTheTipMenuItemSliderSliderBar, AdjustTheTip: AdjustTheTipMenuItemSliderSliderLabel |
| XML usage count | 12561 |
| XML attribute usage count | 12561 |
| Lua usage count | 0 |
| Global usage count | 0 |
| Local definition count | 0 |
| Documentation references | 0 |
| Initialization flow references | 0 |
| Known engine namespace | no |
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
| Conflicting signatures | no |
| Conflicting roles | no |
| Wrapper likely | no |
| Never outside local graph | no |
| Local helper only | no |

## Description

Anchors is a structural XML sub-element. It commonly appears under ActionButtonGroup and AnimatedImage. It is typically used to organize structural children such as AbsPoint, Anchor.

## Common Attributes

- none

## Common Inherits

- none

## Common Parent Elements

- [Label](element_Label.md) — 4631× (HIGH)
- [Window](element_Window.md) — 2735× (HIGH)
- [Button](element_Button.md) — 1954× (HIGH)
- [DynamicImage](element_DynamicImage.md) — 1187× (HIGH)
- [ComboBox](element_ComboBox.md) — 643× (HIGH)
- [FullResizeImage](element_FullResizeImage.md) — 409× (HIGH)
- [EditBox](element_EditBox.md) — 380× (HIGH)
- [SliderBar](element_SliderBar.md) — 219× (HIGH)
- [ListBox](element_ListBox.md) — 110× (HIGH)
- [VerticalScrollbar](element_VerticalScrollbar.md) — 56× (HIGH)
- [CircleImage](element_CircleImage.md) — 54× (HIGH)
- [HorizontalResizeImage](element_HorizontalResizeImage.md) — 48× (HIGH)
- [ScrollWindow](element_ScrollWindow.md) — 43× (HIGH)
- [AnimatedImage](element_AnimatedImage.md) — 38× (HIGH)
- [StatusBar](element_StatusBar.md) — 32× (HIGH)
- [VerticalResizeImage](element_VerticalResizeImage.md) — 14× (HIGH)
- [CooldownDisplay](element_CooldownDisplay.md) — 10× (HIGH)
- [MapDisplay](element_MapDisplay.md) — 8× (MEDIUM)
- [ActionButtonGroup](element_ActionButtonGroup.md) — 3× (MEDIUM)
- [LogDisplay](element_LogDisplay.md) — 3× (MEDIUM)
- [ColorPicker](element_ColorPicker.md) — 1× (LOW)
- [NifDisplay](element_NifDisplay.md) — 1× (LOW)
- [PageWindow](element_PageWindow.md) — 1× (LOW)

## Common Structural Child Elements

- [Anchor](element_Anchor.md) — 14161× (HIGH)
- [AbsPoint](element_AbsPoint.md) — 6× (MEDIUM)

## Structural Sub-Elements

### [Anchor](element_Anchor.md)

Observed 14161 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `point` | **required** | center, topleft, topright, bottomright |
| `relativePoint` | **required** | center, topleft, topright, bottomright |
| `relativeTo` | **required** | Root, $parent, $parentSliderLabel, $parentSliderValueLabel |
| `relativeto` | optional | $parentTitle, $parentPrint, $parentAlert, $parentHyperlink |
| `layer` | optional | secondary, overlay |
| `parent` | optional | Root, $parent |
| `relateiveTo` | optional | $parentDevPadCode, $parentObjectEditBox |
| `textalign` | optional | center |
| `handleinput` | optional | false |
| `relative` | optional | $parent |
| `xOffset` | optional | 0 |
| `yOffset` | optional | 0 |
| `input` | optional | numbers |
| `realtivePoint` | optional | center |
| `realtiveTo` | optional | $parentBackground |
### [AbsPoint](element_AbsPoint.md)

Observed 6 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `x` | **required** | 800, 200, 0, 100 |
| `y` | **required** | 50, -50, 100, -200 |
## Recursive Hierarchy

- Root: [Anchors](element_Anchors.md)
- [AbsPoint](element_AbsPoint.md) (structural, 6×, MEDIUM)
- [Anchor](element_Anchor.md) (structural, 14161×, HIGH)
  - [AbsPoint](element_AbsPoint.md) (structural, 12549×, HIGH)
  - [Anchor](element_Anchor.md) (structural, 29×, HIGH)
    - (cycle)

## Seen In

- AbilityAlert
- AbilityNotifier
- ActionPoints
- AdjustTheTip
- AdvancedPetAssist
- AdvancedRenownTrainer
- AggroMeter
- AnywhereTrainer
- ArmorGraphicToggle
- Asshat
- Assist
- Atlas
- AuctionStats
- Aura
- AutoBand
- AutoMark
- AutoSalvage
- BBars - Mechanic Only
- BankArkel
- BankWindowFix
- BarText (Influence)
- BlackBook
- BlackBox
- Bloody Mess
- Brizio's Crappy Computer Medic
- BuddyBind
- BuffHead
- Busted
- CCTV
- CDown
- CM_ClosetGoblin
- CMap
- CaVES
- Calling
- CastSequence
- ChattyCathy
- Cheeseboard
- CleanCastbar
- CleanUnitFrames
- CleansingBuddy
- Clock
- CombatTextNames
- Compass3D
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
- Deathblow
- Deathblow2
- DeepSleep
- DetauntHelper
- Duel
- DuffTimer
- Dye Preview
- EA_LoadingScreen
- EA_OpenPartyWindow
- EA_ScenarioGroupWindow
- EA_ThreePartBar
- EA_UiDebugTools
- EA_UiModWindow
- EZCraftX
- Effigy
- Emojii
- Enemy
- EveryBodyGuard
- FastFriends
- FastInteract
- FlagCap
- FozAuction
- GCDTracker
- GDes
- Ges
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
- HealGrid
- Hopper
- I HATE YOU THIS MUCH
- InfoScroller
- ItemRack
- JunkDump
- KeyBar
- Keyset
- KeysetMonsterPlay
- KillStreak
- KillTracker
- Killer
- Kwestor
- LibAddonButton
- LibGroup
- LibSurveyor
- LibWBToggler
- LootAlert
- LoyalPet
- MacroIcons
- Map
- MapMonster
- MapPin
- MarkBuff
- Mass Refine
- Mech
- MegaphonePlusPlus
- MiniMapMonster
- Minmap
- Miracle Grow Remix
- MiracleGrow
- MiracleGrowLight
- MoraleCircle
- MoraleSet
- Moth
- Motion
- MouseHint
- MyReasons
- NaturalLog
- NerfedButtons
- NoOverheal
- ObjectInspector
- Obsidian
- OverheadFonts
- Paint the leader
- PartyAd
- PartyCast
- PeaceOut
- Phantom
- PieTracker
- PlayEffectsOn
- Pocket Palette
- PotionBar
- Preciousss
- Pure
- Pure Careerbar
- Queue Queuer
- QuickTacticSwitch
- QuickWarReport
- RPs
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
- RealmStatus
- Refer
- ReliquaryHunter
- RememberIt
- Res
- ResHelp
- RetAlert
- RoR_SoR
- RoR_debolster
- Rolodex
- Rotation
- RvRContribution
- RvRStats
- RvRStatsTab
- SNT_BUTTONS
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
- SpamBayes
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
- Targets
- TastyButtons
- TaxPayer
- TexturedButtons
- ThankTheTank
- TheSeeker
- ThinkOutLoud
- TidyChat
- TidyQueue
- TidyRoll
- TokenMachine
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
- WARCommander
- WARRatingBuster
- WBStutterLess
- WSCT
- WTes
- WaaaghBar
- WarBoard
- WarBoard_AAOTracker
- WarBoard_FPS
- WarBoard_Loc
- WarBoard_Menu
- WarBoard_Session
- WarBoard_TDPS
- WarBoard_TaliMon
- Warbuilder
- Wargames
- WhatsCooking
- WhoHealedMe
- Wikki's Cooldown Bar
- Wikki's Cooldown Pulse
- XpStatus+G
- ZonePOP
- alertMod
- bigger_MacroWindow
- compass
- emotes
- fpsbox
- minesweep
- nLootLink
- talisman-monitor
- wbLeadHelper
- yAssistHelper
- zMailMod

## Examples

- AbilityAlert: AAWindow -> Anchors in Window AAWindow
- AbilityNotifier: AbHelpWindow -> Anchors in Window AbHelpWindow
- ActionPoints: ActionPointsWindow -> Anchors in Window ActionPointsWindow
- ActionPoints: ActionPointsWindowText -> Anchors in Label ActionPointsWindowText
- AdjustTheTip: AdjustTheTipMenuItemSliderSliderBar -> Anchors in SliderBar AdjustTheTipMenuItemSliderSliderBar
- AdjustTheTip: AdjustTheTipMenuItemSliderSliderLabel -> Anchors in Label AdjustTheTipMenuItemSliderSliderLabel

## Related APIs

- [Anchor](element_Anchor.md) (HIGH 100/100) - XML Element Type
- [AnimatedImage](element_AnimatedImage.md) (HIGH 100/100) - XML Element Type
- [Button](element_Button.md) (HIGH 100/100) - XML Element Type
- [CircleImage](element_CircleImage.md) (HIGH 100/100) - XML Element Type
- [ComboBox](element_ComboBox.md) (HIGH 100/100) - XML Element Type
- [CooldownDisplay](element_CooldownDisplay.md) (HIGH 100/100) - XML Element Type
- [DynamicImage](element_DynamicImage.md) (HIGH 100/100) - XML Element Type
- [EditBox](element_EditBox.md) (HIGH 100/100) - XML Element Type
- [FullResizeImage](element_FullResizeImage.md) (HIGH 100/100) - XML Element Type
- [HorizontalResizeImage](element_HorizontalResizeImage.md) (HIGH 100/100) - XML Element Type
- [Label](element_Label.md) (HIGH 100/100) - XML Element Type
- [ListBox](element_ListBox.md) (HIGH 100/100) - XML Element Type
- [MapDisplay](element_MapDisplay.md) (HIGH 100/100) - XML Element Type
- [ScrollWindow](element_ScrollWindow.md) (HIGH 100/100) - XML Element Type
- [SliderBar](element_SliderBar.md) (HIGH 100/100) - XML Element Type
- [StatusBar](element_StatusBar.md) (HIGH 100/100) - XML Element Type
- [VerticalResizeImage](element_VerticalResizeImage.md) (HIGH 100/100) - XML Element Type
- [VerticalScrollbar](element_VerticalScrollbar.md) (HIGH 100/100) - XML Element Type
- [Window](element_Window.md) (HIGH 100/100) - XML Element Type
- [ActionButtonGroup](element_ActionButtonGroup.md) (HIGH 98/100) - XML Element Type
- [LogDisplay](element_LogDisplay.md) (HIGH 98/100) - XML Element Type
- [AbsPoint](element_AbsPoint.md) (MEDIUM 45/100) - XML Element Type
- [ColorPicker](element_ColorPicker.md) (MEDIUM 45/100) - XML Element Type
- [NifDisplay](element_NifDisplay.md) (MEDIUM 45/100) - XML Element Type
- [PageWindow](element_PageWindow.md) (MEDIUM 45/100) - XML Element Type

## Used With

- [EventHandlers](element_EventHandlers.md) (HIGH 100/100) - XML Element Type
- [Size](element_Size.md) (HIGH 100/100) - XML Element Type
- [Windows](element_Windows.md) (HIGH 100/100) - XML Element Type
