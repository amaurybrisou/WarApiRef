# Size

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
| Files seen in | Bars/HealGridActionPointBar.xml, Bars/HealGridCareerBar.xml, Bars/HealGridCastBar.xml, Bars/HealGridMoraleBar.xml, Bars/HealGridProgressBar.xml, Code/Assist/Assist.xml, Code/Assist/AssistConfiguration.xml, Code/CombatLog/CombatLogConfiguration.xml |
| Namespaces detected | Size |
| Source kinds | xml_frames |
| Example locations | AbilityAlert: AAText, AbilityAlert: AAWindow, AbilityNotifier: AbHelpText, AbilityNotifier: AbHelpWindow, ActionPoints: ActionPointsWindow, ActionPoints: ActionPointsWindowText |
| XML usage count | 9065 |
| XML attribute usage count | 9065 |
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

Size is a structural XML sub-element. It commonly appears under AnimatedImage and Button. It is typically used to organize structural children such as AbsPoint.

## Common Attributes

- none

## Common Inherits

- none

## Common Parent Elements

- [Label](element_Label.md) — 4173× (HIGH)
- [Window](element_Window.md) — 1752× (HIGH)
- [Button](element_Button.md) — 1023× (HIGH)
- [DynamicImage](element_DynamicImage.md) — 965× (HIGH)
- [EditBox](element_EditBox.md) — 355× (HIGH)
- [ComboBox](element_ComboBox.md) — 186× (HIGH)
- [SliderBar](element_SliderBar.md) — 186× (HIGH)
- [FullResizeImage](element_FullResizeImage.md) — 113× (HIGH)
- [ListBox](element_ListBox.md) — 76× (HIGH)
- [CircleImage](element_CircleImage.md) — 51× (HIGH)
- [HorizontalResizeImage](element_HorizontalResizeImage.md) — 49× (HIGH)
- [VerticalScrollbar](element_VerticalScrollbar.md) — 47× (HIGH)
- [AnimatedImage](element_AnimatedImage.md) — 32× (HIGH)
- [ScrollWindow](element_ScrollWindow.md) — 21× (HIGH)
- [StatusBar](element_StatusBar.md) — 19× (HIGH)
- [CooldownDisplay](element_CooldownDisplay.md) — 10× (HIGH)
- [VerticalResizeImage](element_VerticalResizeImage.md) — 10× (HIGH)
- [MapDisplay](element_MapDisplay.md) — 3× (MEDIUM)
- [ColorPicker](element_ColorPicker.md) — 1× (LOW)
- [LogDisplay](element_LogDisplay.md) — 1× (LOW)
- [NifDisplay](element_NifDisplay.md) — 1× (LOW)
- [PageWindow](element_PageWindow.md) — 1× (LOW)

## Common Structural Child Elements

- [AbsPoint](element_AbsPoint.md) — 9073× (HIGH)

## Structural Sub-Elements

### [AbsPoint](element_AbsPoint.md)

Observed 9073 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `x` | **required** | 800, 200, 0, 100 |
| `y` | **required** | 50, -50, 100, -200 |
## Recursive Hierarchy

- Root: [Size](element_Size.md)
- [AbsPoint](element_AbsPoint.md) (structural, 9073×, HIGH)

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
- BagOMatic
- BankArkel
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
- Pure
- Pure Careerbar
- Queue Queuer
- QuickTacticSwitch
- QuickWarReport
- RO-Style Combat Text
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
- ShowHealthPercent
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
- VerticalMorale
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
- nRarity
- talisman-monitor
- wbLeadHelper
- yAssistHelper
- zMailMod

## Examples

- AbilityAlert: AAText -> Size in Label AAText
- AbilityAlert: AAWindow -> Size in Window AAWindow
- AbilityNotifier: AbHelpText -> Size in Label AbHelpText
- AbilityNotifier: AbHelpWindow -> Size in Window AbHelpWindow
- ActionPoints: ActionPointsWindow -> Size in Window ActionPointsWindow
- ActionPoints: ActionPointsWindowText -> Size in Label ActionPointsWindowText

## Related APIs

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
- [WindowSetDimensions](../../window_api/functions/window_WindowSetDimensions.md) (HIGH 100/100) - Window Function
- [WindowSetScale](../../window_api/functions/window_WindowSetScale.md) (HIGH 100/100) - Window Function
- [LogDisplay](element_LogDisplay.md) (HIGH 98/100) - XML Element Type
- [OnUpdate](../handlers/handler_OnUpdate.md) (HIGH 88/100) - XML Event
- [AbsPoint](element_AbsPoint.md) (MEDIUM 45/100) - XML Element Type
- [ColorPicker](element_ColorPicker.md) (MEDIUM 45/100) - XML Element Type
- [NifDisplay](element_NifDisplay.md) (MEDIUM 45/100) - XML Element Type
- [PageWindow](element_PageWindow.md) (MEDIUM 45/100) - XML Element Type

## Used With

- [Anchors](element_Anchors.md) (HIGH 100/100) - XML Element Type
- [EventHandlers](element_EventHandlers.md) (HIGH 100/100) - XML Element Type
- [Windows](element_Windows.md) (HIGH 100/100) - XML Element Type
