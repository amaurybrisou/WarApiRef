# OnLButtonUp

- Type: XML Handler
- Confidence level: HIGH
- Confidence score: 100/100

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 138

- Rationale: Promoted as HIGH confidence because seen in 4 or more addons, used directly in xml handler attributes, referenced by generated docs or reference files.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +30 Used directly in XML handler attributes: XML exposure suggests an engine-level contract.
- +18 Used in event registration or dispatch: Observed in event-driven engine hooks.
- +20 Observed in both XML and Lua paths: Cross-source linkage reinforces platform-level usage.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +25 Referenced by generated docs or reference files: The symbol is reinforced outside a single call page.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | AdvancedPetAssist, AdvancedRenownTrainer, AggroMeter, AnywhereTrainer, Atlas, AuctionStats, Aura, AutoBand |
| Files seen in | Source/DebugWindow.xml |
| Namespaces detected | OnLButtonUp |
| Source kinds | bindings, xml_handlers |
| Example locations | AdvancedPetAssist: .OnLButtonUp, AdvancedRenownTrainer: .OnLButtonUp, AggroMeter: .OnLButtonUp, AnywhereTrainer: .OnLButtonUp, Atlas: .OnLButtonUp, AuctionStats: .OnLButtonUp |
| XML usage count | 1 |
| XML attribute usage count | 1 |
| Lua usage count | 1653 |
| Global usage count | 0 |
| Local definition count | 0 |
| Documentation references | 1 |
| Initialization flow references | 0 |
| Known engine namespace | no |
| Default UI presence | no |
| Event binding presence | yes |
| Observed in XML and Lua | yes |
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

XML handler event commonly bound to Button elements in 165 addons.

## Expected Lua Binding

```lua
function(...)
```

## Element Types

- Button

## Seen In

- AdvancedPetAssist
- AdvancedRenownTrainer
- AggroMeter
- AnywhereTrainer
- Atlas
- AuctionStats
- Aura
- AutoBand
- AutoSalvage
- BagOMatic
- BankArkel
- BlackBook
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
- CleanUnitFrames
- Crusher
- DAoCBuff
- DPSMeter
- DammazKron
- DetauntHelper
- DuffTimer
- Dye Preview
- EA_OpenPartyWindow
- EA_ScenarioGroupWindow
- EA_UiDebugTools
- EA_UiModWindow
- EZCraft
- EZCraftX
- Effigy
- Emojii
- Enemy
- FastFriends
- FastInteract
- FozAuction
- GetStats
- Group Icons SG
- GroupRange
- GroupSpotter
- GuildWarden
- HealGrid
- Hopper
- JunkDump
- KeyBar
- Keyset
- KeysetMonsterPlay
- KillTracker
- Killer
- Kwestor
- LibAddonButton
- LibGroup
- LootAlert
- LoyalPet
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
- MoraleCircle
- MoraleSet
- Motion
- NerfedButtons
- NoOverheal
- ObjectInspector
- Obsidian
- Phantom
- PieTracker
- Pocket Palette
- PotionBar
- Pure
- QuickTacticSwitch
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
- RealmStatus
- Refer
- ReliquaryHunter
- RoR_SoR
- RoR_debolster
- Rolodex
- Rotation
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
- SocialWindow 2.0
- Soloq
- Squared
- SquaredClick
- Statdoll
- Statdoll Remix
- TacticSetNames
- TalismanGenie
- TastyButtons
- TaxPayer
- TexturedButtons
- ThankTheTank
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
- Twister
- TwisterSet
- WARCommander
- WBStutterLess
- WSCT
- WTes
- WarBoard
- WarBoard_AAOTracker
- WarBoard_FPS
- WarBoard_Loc
- WarBoard_Session
- WarBoard_TDPS
- WarBoard_TaliMon
- WhoHealedMe
- Wikki's Cooldown Bar
- Wikki's Cooldown Pulse
- XpStatus+G
- ZonePOP
- alertMod
- bigger_MacroWindow
- emotes
- followTheLeader
- minesweep
- nLootLink
- talisman-monitor
- wbLeadHelper

## Examples

- EA_UiDebugTools: DevPadNewWindowCancel -> DevPadNewWindowCancel.OnLButtonUp -> DevPadWindow.HideNewWindow
- AdvancedPetAssist: .OnLButtonUp -> APAGui.OnTabButtonUp
- AdvancedPetAssist: .OnLButtonUp -> APAGui.Hide
- AdvancedPetAssist: .OnLButtonUp -> APAGui.OnFollowTargetHUDClicked
- AdvancedPetAssist: .OnLButtonUp -> APAGui.OnKitingHUDClicked
- AdvancedPetAssist: .OnLButtonUp -> APAGui.OnInstantOnlyHUDClicked

## Related APIs

- [AuctionWindow.Hide](../../globals/functions/global_AuctionWindow.Hide.md) (HIGH 100/100) - Global Function
- [Button](../element_types/element_Button.md) (HIGH 100/100) - XML Element Type
- [ButtonSetPressedFlag](../../window_api/functions/window_ButtonSetPressedFlag.md) (HIGH 100/100) - Window Function
- [ComboBox](../element_types/element_ComboBox.md) (HIGH 100/100) - XML Element Type
- [Cursor.Clear](../../globals/functions/global_Cursor.Clear.md) (HIGH 100/100) - Global Function
- [Cursor.IconOnCursor](../../globals/functions/global_Cursor.IconOnCursor.md) (HIGH 100/100) - Global Function
- [DialogManager.MakeTextEntryDialog](../../globals/functions/global_DialogManager.MakeTextEntryDialog.md) (HIGH 100/100) - Global Function
- [DialogManager.MakeTwoButtonDialog](../../globals/functions/global_DialogManager.MakeTwoButtonDialog.md) (HIGH 100/100) - Global Function
- [DynamicImage](../element_types/element_DynamicImage.md) (HIGH 100/100) - XML Element Type
- [EA_ChatWindow.Print](../../globals/functions/global_EA_ChatWindow.Print.md) (HIGH 100/100) - Global Function
- [EA_Window_ContextMenu.AddMenuItem](../../globals/functions/global_EA_Window_ContextMenu.AddMenuItem.md) (HIGH 100/100) - Global Function
- [EA_Window_ContextMenu.CreateContextMenu](../../globals/functions/global_EA_Window_ContextMenu.CreateContextMenu.md) (HIGH 100/100) - Global Function
- [EA_Window_ContextMenu.Finalize](../../globals/functions/global_EA_Window_ContextMenu.Finalize.md) (HIGH 100/100) - Global Function
- [EA_Window_ContextMenu.HideAll](../../globals/functions/global_EA_Window_ContextMenu.HideAll.md) (HIGH 100/100) - Global Function
- [EA_Window_OpenParty.ToggleFullWindow](../../globals/functions/global_EA_Window_OpenParty.ToggleFullWindow.md) (HIGH 100/100) - Global Function
- [EditBox](../element_types/element_EditBox.md) (HIGH 100/100) - XML Element Type
- [FullResizeImage](../element_types/element_FullResizeImage.md) (HIGH 100/100) - XML Element Type
- [Label](../element_types/element_Label.md) (HIGH 100/100) - XML Element Type
- [LabelSetText](../../window_api/functions/window_LabelSetText.md) (HIGH 100/100) - Window Function
- [LayoutEditor.UnregisterWindow](../../window_api/functions/window_LayoutEditor.UnregisterWindow.md) (HIGH 100/100) - Window Function
- [ListBox](../element_types/element_ListBox.md) (HIGH 100/100) - XML Element Type
- [MapDisplay](../element_types/element_MapDisplay.md) (HIGH 100/100) - XML Element Type
- [ScenarioGroupWindow.LeaveGroup](../../globals/functions/global_ScenarioGroupWindow.LeaveGroup.md) (HIGH 100/100) - Global Function
- [SliderBar](../element_types/element_SliderBar.md) (HIGH 100/100) - XML Element Type
- [VerticalScrollbar](../element_types/element_VerticalScrollbar.md) (HIGH 100/100) - XML Element Type
- [Window](../element_types/element_Window.md) (HIGH 100/100) - XML Element Type
- [WindowGetAnchor](../../window_api/functions/window_WindowGetAnchor.md) (HIGH 100/100) - Window Function
- [WindowGetId](../../window_api/functions/window_WindowGetId.md) (HIGH 100/100) - Window Function
- [WindowGetMovable](../../window_api/functions/window_WindowGetMovable.md) (HIGH 100/100) - Window Function
- [WindowGetParent](../../window_api/functions/window_WindowGetParent.md) (HIGH 100/100) - Window Function
- [WindowGetTintColor](../../window_api/functions/window_WindowGetTintColor.md) (HIGH 100/100) - Window Function
- [WindowSetAlpha](../../window_api/functions/window_WindowSetAlpha.md) (HIGH 100/100) - Window Function
- [WindowSetShowing](../../window_api/functions/window_WindowSetShowing.md) (HIGH 100/100) - Window Function
- [WindowSetTintColor](../../window_api/functions/window_WindowSetTintColor.md) (HIGH 100/100) - Window Function
- [OnUpdate](handler_OnUpdate.md) (HIGH 88/100) - XML Event
- [DoesWindowExist](../../globals/functions/global_DoesWindowExist.md) (HIGH 83/100) - Global Function
- [DestroyWindow](../../globals/functions/global_DestroyWindow.md) (HIGH 75/100) - Global Function
- [SendUseItem](../../globals/functions/global_SendUseItem.md) (HIGH 75/100) - Global Function
- [ColorPicker](../element_types/element_ColorPicker.md) (MEDIUM 45/100) - XML Element Type
- [NifDisplay](../element_types/element_NifDisplay.md) (MEDIUM 45/100) - XML Element Type

## Used With

- [DialogManager.MakeTwoButtonDialog](../../globals/functions/global_DialogManager.MakeTwoButtonDialog.md) (HIGH 100/100) - Global Function
- [EA_ChatWindow.Print](../../globals/functions/global_EA_ChatWindow.Print.md) (HIGH 100/100) - Global Function
- [WindowGetTintColor](../../window_api/functions/window_WindowGetTintColor.md) (HIGH 100/100) - Window Function
- [OnUpdate](handler_OnUpdate.md) (HIGH 88/100) - XML Event
- [DestroyWindow](../../globals/functions/global_DestroyWindow.md) (HIGH 75/100) - Global Function

## Affects

- [SystemData.ActiveWindow.name](../../systemdata/fields/systemdata_SystemData.ActiveWindow.name.md) (HIGH 100/100) - SystemData Field
- [SystemData.MouseOverWindow.name](../../systemdata/fields/systemdata_SystemData.MouseOverWindow.name.md) (HIGH 100/100) - SystemData Field

## Notes

- Expected binding arguments remain uncertain because contract artifacts capture symbol linkage, not full handler signatures.
