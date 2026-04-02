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
| Addons seen in | AdvancedPetAssist, AdvancedRenownTrainer, AggroMeter, AnywhereTrainer, AnywhereTrainerAdditions, Aura, BagOMatic, BankArkel |
| Files seen in | `/workspace_addons/AdvancedPetAssist/APAGui.xml:116`, `/workspace_addons/AdvancedPetAssist/APAGui.xml:1386`, `/workspace_addons/AdvancedPetAssist/APAGui.xml:141`, `/workspace_addons/AdvancedPetAssist/APAGui.xml:1427`, `/workspace_addons/AdvancedPetAssist/APAGui.xml:1468`, `/workspace_addons/AdvancedPetAssist/APAGui.xml:155`, `/workspace_addons/AdvancedPetAssist/APAGui.xml:169`, `/workspace_addons/AdvancedPetAssist/APAGui.xml:183` |
| Namespaces detected | OnLButtonUp |
| Source kinds | bindings, xml_handlers |
| Example locations | AdvancedPetAssist: APAFollowTargetHUD.OnLButtonUp, AdvancedPetAssist: APAInstantOnlyHUD.OnLButtonUp, AdvancedPetAssist: APAKitingHUD.OnLButtonUp, AdvancedPetAssist: APAOptionsClose.OnLButtonUp, AdvancedPetAssist: APAOptionsTabsAutoRecall.OnLButtonUp, AdvancedPetAssist: APAOptionsTabsControls.OnLButtonUp |
| XML usage count | 824 |
| XML attribute usage count | 824 |
| Lua usage count | 823 |
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

Observed as an XML handler hook bound by 51 addons through frame event handlers.

## Expected Lua Binding

```lua
function(...)
```

## Element Types

- Button
- ColorPicker
- DynamicImage
- FullResizeImage
- Label
- ListBox
- MapDisplay
- Window

## Seen In

- AdvancedPetAssist
- AdvancedRenownTrainer
- AggroMeter
- AnywhereTrainer
- AnywhereTrainerAdditions
- Aura
- BagOMatic
- BankArkel
- BuffHead
- Busted
- CM_ClosetGoblin
- CMap
- Cheeseboard
- DAoCBuff
- EA_UiDebugTools
- EA_UiModWindow
- Effigy
- Enemy
- FastInteract
- GetStats
- JunkDump
- Killer
- LibGroup
- LoyalPet
- MapMonster
- MapPin
- Miracle Grow Remix
- MiracleGrow
- MiracleGrowLight
- MoraleCircle
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
- wbLeadHelper

## Examples

- AdvancedPetAssist: APAFollowTargetHUD -> APAFollowTargetHUD.OnLButtonUp -> APAGui.OnFollowTargetHUDClicked
- AdvancedPetAssist: APAInstantOnlyHUD -> APAInstantOnlyHUD.OnLButtonUp -> APAGui.OnInstantOnlyHUDClicked
- AdvancedPetAssist: APAKitingHUD -> APAKitingHUD.OnLButtonUp -> APAGui.OnKitingHUDClicked
- AdvancedPetAssist: APAOptionsClose -> APAOptionsClose.OnLButtonUp -> APAGui.Hide
- AdvancedPetAssist: APAOptionsTabsAutoRecall -> APAOptionsTabsAutoRecall.OnLButtonUp -> APAGui.OnTabButtonUp
- AdvancedPetAssist: APAOptionsTabsControls -> APAOptionsTabsControls.OnLButtonUp -> APAGui.OnTabButtonUp

## Related APIs

- [AlertTextWindow.AddLine](../../globals/functions/global_AlertTextWindow.AddLine.md) (HIGH 100/100) - Global Function
- [ButtonGetDisabledFlag](../../window_api/functions/window_ButtonGetDisabledFlag.md) (HIGH 100/100) - Window Function
- [ButtonGetPressedFlag](../../window_api/functions/window_ButtonGetPressedFlag.md) (HIGH 100/100) - Window Function
- [ButtonGetText](../../window_api/functions/window_ButtonGetText.md) (HIGH 100/100) - Window Function
- [ComboBoxClearMenuItems](../../window_api/functions/window_ComboBoxClearMenuItems.md) (HIGH 100/100) - Window Function
- [ComboBoxGetSelectedMenuItem](../../window_api/functions/window_ComboBoxGetSelectedMenuItem.md) (HIGH 100/100) - Window Function
- [ComboBoxGetSelectedText](../../window_api/functions/window_ComboBoxGetSelectedText.md) (HIGH 100/100) - Window Function
- [ComboBoxSetSelectedMenuItem](../../window_api/functions/window_ComboBoxSetSelectedMenuItem.md) (HIGH 100/100) - Window Function
- [Cursor.Clear](../../globals/functions/global_Cursor.Clear.md) (HIGH 100/100) - Global Function
- [Cursor.IconOnCursor](../../globals/functions/global_Cursor.IconOnCursor.md) (HIGH 100/100) - Global Function
- [DebugWindow.ClearTextLog](../../globals/functions/global_DebugWindow.ClearTextLog.md) (HIGH 100/100) - Global Function
- [DebugWindow.OnShowFocus](../../globals/functions/global_DebugWindow.OnShowFocus.md) (HIGH 100/100) - Global Function
- [DialogManager.MakeOneButtonDialog](../../globals/functions/global_DialogManager.MakeOneButtonDialog.md) (HIGH 100/100) - Global Function
- [DialogManager.MakeTextEntryDialog](../../globals/functions/global_DialogManager.MakeTextEntryDialog.md) (HIGH 100/100) - Global Function
- [DialogManager.MakeTwoButtonDialog](../../globals/functions/global_DialogManager.MakeTwoButtonDialog.md) (HIGH 100/100) - Global Function
- [EA_BackpackUtilsMediator.GetCurrentBackpackType](../../globals/functions/global_EA_BackpackUtilsMediator.GetCurrentBackpackType.md) (HIGH 100/100) - Global Function
- [EA_BackpackUtilsMediator.GetCursorForBackpack](../../globals/functions/global_EA_BackpackUtilsMediator.GetCursorForBackpack.md) (HIGH 100/100) - Global Function
- [EA_BackpackUtilsMediator.GetItemsFromBackpack](../../globals/functions/global_EA_BackpackUtilsMediator.GetItemsFromBackpack.md) (HIGH 100/100) - Global Function
- [EA_ChatWindow.InsertText](../../globals/functions/global_EA_ChatWindow.InsertText.md) (HIGH 100/100) - Global Function
- [EA_Window_ContextMenu.AddMenuDivider](../../globals/functions/global_EA_Window_ContextMenu.AddMenuDivider.md) (HIGH 100/100) - Global Function
- [EA_Window_ContextMenu.Hide](../../globals/functions/global_EA_Window_ContextMenu.Hide.md) (HIGH 100/100) - Global Function
- [EA_Window_WorldMap.ShowZone](../../globals/functions/global_EA_Window_WorldMap.ShowZone.md) (HIGH 100/100) - Global Function
- [GameData.Player.GetRenownRefundCost](../../globals/functions/global_GameData.Player.GetRenownRefundCost.md) (HIGH 100/100) - Global Function
- [InterfaceCore.GetResolutionScale](../../globals/functions/global_InterfaceCore.GetResolutionScale.md) (HIGH 100/100) - Global Function
- [InterfaceCore.ReloadUI](../../globals/functions/global_InterfaceCore.ReloadUI.md) (HIGH 100/100) - Global Function
- [LabelGetText](../../window_api/functions/window_LabelGetText.md) (HIGH 100/100) - Window Function
- [PartyUtils.GetWarbandLeader](../../globals/functions/global_PartyUtils.GetWarbandLeader.md) (HIGH 100/100) - Global Function
- [PartyUtils.IsPartyActive](../../globals/functions/global_PartyUtils.IsPartyActive.md) (HIGH 100/100) - Global Function
- [TextEditBoxGetText](../../window_api/functions/window_TextEditBoxGetText.md) (HIGH 100/100) - Window Function
- [TextEditBoxSelectAll](../../window_api/functions/window_TextEditBoxSelectAll.md) (HIGH 100/100) - Window Function
- [WindowAssignFocus](../../window_api/functions/window_WindowAssignFocus.md) (HIGH 100/100) - Window Function
- [WindowGetAnchorCount](../../window_api/functions/window_WindowGetAnchorCount.md) (HIGH 100/100) - Window Function
- [WindowGetHandleInput](../../window_api/functions/window_WindowGetHandleInput.md) (HIGH 100/100) - Window Function
- [WindowGetId](../../window_api/functions/window_WindowGetId.md) (HIGH 100/100) - Window Function
- [WindowGetParent](../../window_api/functions/window_WindowGetParent.md) (HIGH 100/100) - Window Function
- [WindowGetShowing](../../window_api/functions/window_WindowGetShowing.md) (HIGH 100/100) - Window Function
- [WindowSetGameActionData](../../window_api/functions/window_WindowSetGameActionData.md) (HIGH 100/100) - Window Function
- [WindowSetShowing](../../window_api/functions/window_WindowSetShowing.md) (HIGH 100/100) - Window Function
- [wstring.match](../../globals/functions/global_wstring.match.md) (HIGH 100/100) - Global Function

## Used With

- [Button](../element_types/element_Button.md) (HIGH 100/100) - XML Element Type
- [ButtonGetDisabledFlag](../../window_api/functions/window_ButtonGetDisabledFlag.md) (HIGH 100/100) - Window Function
- [ButtonGetPressedFlag](../../window_api/functions/window_ButtonGetPressedFlag.md) (HIGH 100/100) - Window Function
- [ButtonSetPressedFlag](../../window_api/functions/window_ButtonSetPressedFlag.md) (HIGH 100/100) - Window Function
- [ComboBoxGetSelectedMenuItem](../../window_api/functions/window_ComboBoxGetSelectedMenuItem.md) (HIGH 100/100) - Window Function
- [ComboBoxGetSelectedText](../../window_api/functions/window_ComboBoxGetSelectedText.md) (HIGH 100/100) - Window Function
- [DynamicImage](../element_types/element_DynamicImage.md) (HIGH 100/100) - XML Element Type
- [FullResizeImage](../element_types/element_FullResizeImage.md) (HIGH 100/100) - XML Element Type
- [Label](../element_types/element_Label.md) (HIGH 100/100) - XML Element Type
- [ListBox](../element_types/element_ListBox.md) (HIGH 100/100) - XML Element Type
- [MapDisplay](../element_types/element_MapDisplay.md) (HIGH 100/100) - XML Element Type
- [TextEditBoxGetText](../../window_api/functions/window_TextEditBoxGetText.md) (HIGH 100/100) - Window Function
- [Window](../element_types/element_Window.md) (HIGH 100/100) - XML Element Type
- [WindowGetId](../../window_api/functions/window_WindowGetId.md) (HIGH 100/100) - Window Function
- [WindowGetParent](../../window_api/functions/window_WindowGetParent.md) (HIGH 100/100) - Window Function
- [WindowSetShowing](../../window_api/functions/window_WindowSetShowing.md) (HIGH 100/100) - Window Function
- [wstring.match](../../globals/functions/global_wstring.match.md) (HIGH 100/100) - Global Function

## Triggered By

- none

## Affects

- none

## Notes

- Expected binding arguments remain uncertain because addon-api docs capture symbol linkage, not full handler signatures.
