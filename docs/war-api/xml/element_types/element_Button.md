# Button

- Category: XML Element Type
- Confidence level: HIGH
- Confidence score: 100/100

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 198

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, seen in 4 or more addons, used directly in xml handler attributes.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +30 Used directly in XML handler attributes: XML exposure suggests an engine-level contract.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +18 Used in event registration or dispatch: Observed in event-driven engine hooks.
- +20 Observed in both XML and Lua paths: Cross-source linkage reinforces platform-level usage.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +25 Referenced by generated docs or reference files: The symbol is reinforced outside a single call page.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | AdvancedPetAssist, AdvancedRenownTrainer, AggroMeter, AnywhereTrainer, Aura, BankArkel, BuffHead, Busted |
| Files seen in | `/workspace_addons/AdvancedPetAssist/APAGui.xml:114`, `/workspace_addons/AdvancedPetAssist/APAGui.xml:131`, `/workspace_addons/AdvancedPetAssist/APAGui.xml:145`, `/workspace_addons/AdvancedPetAssist/APAGui.xml:159`, `/workspace_addons/AdvancedPetAssist/APAGui.xml:173`, `/workspace_addons/AdvancedPetAssist/APAGui.xml:187`, `/workspace_addons/AdvancedPetAssist/APAGui.xml:201`, `/workspace_addons/AdvancedPetAssist/APAGui.xml:215` |
| Namespaces detected | Button |
| Source kinds | xml_frames, xml_handlers |
| Example locations | AdvancedPetAssist: APAOptionsClose, AdvancedPetAssist: APAOptionsTabsAutoRecall, AdvancedPetAssist: APAOptionsTabsControls, AdvancedPetAssist: APAOptionsTabsFollowTarget, AdvancedPetAssist: APAOptionsTabsGeneral, AdvancedPetAssist: APAOptionsTabsHUD |
| XML usage count | 917 |
| XML attribute usage count | 917 |
| Lua usage count | 10 |
| Global usage count | 0 |
| Local definition count | 0 |
| Documentation references | 1 |
| Initialization flow references | 0 |
| Known engine namespace | yes |
| Default UI presence | yes |
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

Observed XML element type instantiated by 51 addons.

## Common Attributes

- name
- inherits
- id
- textalign
- font
- layer
- handleinput
- backgroundtexture
- highlighttexture
- drawchildrenfirst
- texturescale
- alpha

## Common Handlers

- OnLButtonUp
- OnMouseOver
- OnRButtonUp
- OnLButtonDown
- OnMouseOverEnd
- OnRButtonDown
- OnMouseDrag
- OnInitialize
- OnMButtonUp
- OnMouseWheel

## Common Inherits

- EA_Button_DefaultResizeable
- EA_Button_DefaultWindowClose
- EA_Button_DefaultText
- EA_Button_DefaultCheckBox
- EA_Window_MacroIconButton
- MacroIconSelectionWindowIconButton
- EA_Button_Default
- EA_Button_Tab
- ClosetGoblinEquipmentButton
- DefaultButton
- EA_Button_DefaultListSort
- CG_ItemRackEquipmentButton

## Seen In

- AdvancedPetAssist
- AdvancedRenownTrainer
- AggroMeter
- AnywhereTrainer
- Aura
- BankArkel
- BuffHead
- Busted
- CM_ClosetGoblin
- CMap
- Cheeseboard
- DAoCBuff
- DaemonAssist
- EA_UiDebugTools
- EA_UiModWindow
- Enemy
- FastInteract
- GetStats
- GuardList
- GuardRange
- JunkDump
- Killer
- LibGroup
- LoyalPet
- MapMonster
- MapPin
- Miracle Grow Remix
- MiracleGrow
- MiracleGrowLight
- ObjectInspector
- Pocket Palette
- PotionBar
- QuickTacticSwitch
- QuickWarReport
- RVAPI_ColorDialog
- RVMOD_Manager
- RandomMount
- RoR_SoR
- Shinies
- TexturedButtons
- TidyChat
- TidyRoll
- Tortall_DPS
- TurretRange
- WSCT
- WarBoard
- WhoHealedMe
- bigger_MacroWindow
- followTheLeader
- nRarity
- wbLeadHelper

## Examples

- AdvancedPetAssist: APAOptionsClose -> Button APAOptionsClose
- AdvancedPetAssist: APAOptionsTabsAutoRecall -> Button APAOptionsTabsAutoRecall
- AdvancedPetAssist: APAOptionsTabsControls -> Button APAOptionsTabsControls
- AdvancedPetAssist: APAOptionsTabsFollowTarget -> Button APAOptionsTabsFollowTarget
- AdvancedPetAssist: APAOptionsTabsGeneral -> Button APAOptionsTabsGeneral
- AdvancedPetAssist: APAOptionsTabsHUD -> Button APAOptionsTabsHUD

## Related APIs

- [AlertTextWindow.AddLine](../../globals/functions/global_AlertTextWindow.AddLine.md) (HIGH 100/100) - Global Function
- [ButtonGetDisabledFlag](../../window_api/functions/window_ButtonGetDisabledFlag.md) (HIGH 100/100) - Window Function
- [ButtonGetText](../../window_api/functions/window_ButtonGetText.md) (HIGH 100/100) - Window Function
- [ComboBoxClearMenuItems](../../window_api/functions/window_ComboBoxClearMenuItems.md) (HIGH 100/100) - Window Function
- [ComboBoxGetSelectedMenuItem](../../window_api/functions/window_ComboBoxGetSelectedMenuItem.md) (HIGH 100/100) - Window Function
- [ComboBoxGetSelectedText](../../window_api/functions/window_ComboBoxGetSelectedText.md) (HIGH 100/100) - Window Function
- [ComboBoxSetSelectedMenuItem](../../window_api/functions/window_ComboBoxSetSelectedMenuItem.md) (HIGH 100/100) - Window Function
- [Cursor.IconOnCursor](../../globals/functions/global_Cursor.IconOnCursor.md) (HIGH 100/100) - Global Function
- [Cursor.PickUp](../../globals/functions/global_Cursor.PickUp.md) (HIGH 100/100) - Global Function
- [DebugWindow.ClearTextLog](../../globals/functions/global_DebugWindow.ClearTextLog.md) (HIGH 100/100) - Global Function
- [DebugWindow.OnShowFocus](../../globals/functions/global_DebugWindow.OnShowFocus.md) (HIGH 100/100) - Global Function
- [DialogManager.MakeOneButtonDialog](../../globals/functions/global_DialogManager.MakeOneButtonDialog.md) (HIGH 100/100) - Global Function
- [DialogManager.MakeTextEntryDialog](../../globals/functions/global_DialogManager.MakeTextEntryDialog.md) (HIGH 100/100) - Global Function
- [DialogManager.MakeTwoButtonDialog](../../globals/functions/global_DialogManager.MakeTwoButtonDialog.md) (HIGH 100/100) - Global Function
- [EA_ChatWindow.InsertText](../../globals/functions/global_EA_ChatWindow.InsertText.md) (HIGH 100/100) - Global Function
- [EA_Window_ContextMenu.AddMenuDivider](../../globals/functions/global_EA_Window_ContextMenu.AddMenuDivider.md) (HIGH 100/100) - Global Function
- [EA_Window_ContextMenu.Hide](../../globals/functions/global_EA_Window_ContextMenu.Hide.md) (HIGH 100/100) - Global Function
- [GameData.Player.GetRenownRefundCost](../../globals/functions/global_GameData.Player.GetRenownRefundCost.md) (HIGH 100/100) - Global Function
- [InterfaceCore.ReloadUI](../../globals/functions/global_InterfaceCore.ReloadUI.md) (HIGH 100/100) - Global Function
- [LabelGetText](../../window_api/functions/window_LabelGetText.md) (HIGH 100/100) - Window Function
- [TextEditBoxGetText](../../window_api/functions/window_TextEditBoxGetText.md) (HIGH 100/100) - Window Function
- [TextEditBoxSelectAll](../../window_api/functions/window_TextEditBoxSelectAll.md) (HIGH 100/100) - Window Function
- [WindowAssignFocus](../../window_api/functions/window_WindowAssignFocus.md) (HIGH 100/100) - Window Function
- [WindowGetAnchorCount](../../window_api/functions/window_WindowGetAnchorCount.md) (HIGH 100/100) - Window Function
- [WindowGetHandleInput](../../window_api/functions/window_WindowGetHandleInput.md) (HIGH 100/100) - Window Function
- [WindowGetId](../../window_api/functions/window_WindowGetId.md) (HIGH 100/100) - Window Function
- [WindowGetParent](../../window_api/functions/window_WindowGetParent.md) (HIGH 100/100) - Window Function
- [WindowGetShowing](../../window_api/functions/window_WindowGetShowing.md) (HIGH 100/100) - Window Function
- [WindowSetShowing](../../window_api/functions/window_WindowSetShowing.md) (HIGH 100/100) - Window Function
- [wstring.match](../../globals/functions/global_wstring.match.md) (HIGH 100/100) - Global Function

## Used With

- [ButtonGetDisabledFlag](../../window_api/functions/window_ButtonGetDisabledFlag.md) (HIGH 100/100) - Window Function
- [ButtonGetText](../../window_api/functions/window_ButtonGetText.md) (HIGH 100/100) - Window Function
- [ComboBoxClearMenuItems](../../window_api/functions/window_ComboBoxClearMenuItems.md) (HIGH 100/100) - Window Function
- [ComboBoxGetSelectedMenuItem](../../window_api/functions/window_ComboBoxGetSelectedMenuItem.md) (HIGH 100/100) - Window Function
- [ComboBoxGetSelectedText](../../window_api/functions/window_ComboBoxGetSelectedText.md) (HIGH 100/100) - Window Function
- [ComboBoxSetSelectedMenuItem](../../window_api/functions/window_ComboBoxSetSelectedMenuItem.md) (HIGH 100/100) - Window Function
- [Cursor.IconOnCursor](../../globals/functions/global_Cursor.IconOnCursor.md) (HIGH 100/100) - Global Function
- [Cursor.PickUp](../../globals/functions/global_Cursor.PickUp.md) (HIGH 100/100) - Global Function
- [DebugWindow.ClearTextLog](../../globals/functions/global_DebugWindow.ClearTextLog.md) (HIGH 100/100) - Global Function
- [DialogManager.MakeOneButtonDialog](../../globals/functions/global_DialogManager.MakeOneButtonDialog.md) (HIGH 100/100) - Global Function
- [DialogManager.MakeTextEntryDialog](../../globals/functions/global_DialogManager.MakeTextEntryDialog.md) (HIGH 100/100) - Global Function
- [DialogManager.MakeTwoButtonDialog](../../globals/functions/global_DialogManager.MakeTwoButtonDialog.md) (HIGH 100/100) - Global Function
- [EA_ChatWindow.InsertText](../../globals/functions/global_EA_ChatWindow.InsertText.md) (HIGH 100/100) - Global Function
- [EA_Window_ContextMenu.Hide](../../globals/functions/global_EA_Window_ContextMenu.Hide.md) (HIGH 100/100) - Global Function
- [EA_Window_InteractionRenownTraining](../../globals/tables/table_EA_Window_InteractionRenownTraining.md) (HIGH 100/100) - Global Table
- [GameData.Player.GetRenownRefundCost](../../globals/functions/global_GameData.Player.GetRenownRefundCost.md) (HIGH 100/100) - Global Function
- [InterfaceCore.ReloadUI](../../globals/functions/global_InterfaceCore.ReloadUI.md) (HIGH 100/100) - Global Function
- [LabelGetText](../../window_api/functions/window_LabelGetText.md) (HIGH 100/100) - Window Function
- [OnInitialize](../handlers/handler_OnInitialize.md) (HIGH 100/100) - XML Handler
- [OnInitialize](../../events/window_events/window_event_OnInitialize.md) (HIGH 100/100) - Window Event
- [OnLButtonDown](../handlers/handler_OnLButtonDown.md) (HIGH 100/100) - XML Handler
- [OnLButtonDown](../../events/window_events/window_event_OnLButtonDown.md) (HIGH 100/100) - Window Event
- [OnLButtonUp](../handlers/handler_OnLButtonUp.md) (HIGH 100/100) - XML Handler
- [OnLButtonUp](../../events/window_events/window_event_OnLButtonUp.md) (HIGH 100/100) - Window Event
- [OnMButtonUp](../handlers/handler_OnMButtonUp.md) (HIGH 100/100) - XML Handler
- [OnMButtonUp](../../events/window_events/window_event_OnMButtonUp.md) (HIGH 100/100) - Window Event
- [OnMouseDrag](../handlers/handler_OnMouseDrag.md) (HIGH 100/100) - XML Handler
- [OnMouseDrag](../../events/window_events/window_event_OnMouseDrag.md) (HIGH 100/100) - Window Event
- [OnMouseOver](../handlers/handler_OnMouseOver.md) (HIGH 100/100) - XML Handler
- [OnMouseOver](../../events/window_events/window_event_OnMouseOver.md) (HIGH 100/100) - Window Event
- [OnMouseOverEnd](../handlers/handler_OnMouseOverEnd.md) (HIGH 100/100) - XML Handler
- [OnMouseOverEnd](../../events/window_events/window_event_OnMouseOverEnd.md) (HIGH 100/100) - Window Event
- [OnMouseWheel](../handlers/handler_OnMouseWheel.md) (HIGH 100/100) - XML Handler
- [OnMouseWheel](../../events/window_events/window_event_OnMouseWheel.md) (HIGH 100/100) - Window Event
- [OnRButtonDown](../handlers/handler_OnRButtonDown.md) (HIGH 100/100) - XML Handler
- [OnRButtonDown](../../events/window_events/window_event_OnRButtonDown.md) (HIGH 100/100) - Window Event
- [OnRButtonUp](../../events/window_events/window_event_OnRButtonUp.md) (HIGH 100/100) - Window Event
- [OnRButtonUp](../handlers/handler_OnRButtonUp.md) (HIGH 100/100) - XML Handler
- [SystemData.ActiveWindow.name](../../systemdata/fields/systemdata_SystemData.ActiveWindow.name.md) (HIGH 100/100) - SystemData Field
- [SystemData.MouseOverWindow.name](../../systemdata/fields/systemdata_SystemData.MouseOverWindow.name.md) (HIGH 100/100) - SystemData Field
- [TextEditBoxGetText](../../window_api/functions/window_TextEditBoxGetText.md) (HIGH 100/100) - Window Function
- [TextEditBoxSelectAll](../../window_api/functions/window_TextEditBoxSelectAll.md) (HIGH 100/100) - Window Function
- [TextEditBoxSetText](../../window_api/functions/window_TextEditBoxSetText.md) (HIGH 100/100) - Window Function
- [WindowAssignFocus](../../window_api/functions/window_WindowAssignFocus.md) (HIGH 100/100) - Window Function
- [WindowGetId](../../window_api/functions/window_WindowGetId.md) (HIGH 100/100) - Window Function
- [WindowGetParent](../../window_api/functions/window_WindowGetParent.md) (HIGH 100/100) - Window Function
- [WindowGetShowing](../../window_api/functions/window_WindowGetShowing.md) (HIGH 100/100) - Window Function
- [WindowSetShowing](../../window_api/functions/window_WindowSetShowing.md) (HIGH 100/100) - Window Function
- [wstring.match](../../globals/functions/global_wstring.match.md) (HIGH 100/100) - Global Function

## Triggered By

- none

## Affects

- none
