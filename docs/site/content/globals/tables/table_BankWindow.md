# BankWindow

- Category: Global Table
- Confidence level: HIGH
- Confidence score: 100/100

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 140

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, seen in 4 or more addons, matches a known engine namespace.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +10 Referenced from initialization flow: Lifecycle reconstruction references this symbol.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +25 Referenced by generated docs or reference files: The symbol is reinforced outside a single call page.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | AnywhereTrainer, AnywhereTrainerAdditions, BagOMatic, BankArkel, BankWindowFix, nRarity |
| Files seen in | `/workspace/AnywhereTrainer/source/AnywhereTrainer.lua:271`, `/workspace/AnywhereTrainerAdditions/AnywhereTrainerAdditions.lua:131`, `/workspace/AnywhereTrainerAdditions/AnywhereTrainerAdditions.lua:218`, `/workspace/BankArkel/BankArkel.lua:563`, `/workspace/BankArkel/BankArkel.lua:573`, `/workspace/BankWindowFix/Source/BankWindowFix.lua:108`, `/workspace/BankWindowFix/Source/BankWindowFix.lua:30`, `/workspace/bagomatic/BagOMatic.lua:202` |
| Namespaces detected | BankWindow |
| Source kinds | globals, lua_calls |
| Example locations | AnywhereTrainer: AnywhereTrainer.OnLeftClickBank, AnywhereTrainerAdditions: AnywhereTrainerAdditions.EquipmentLButtonDown, AnywhereTrainerAdditions: AnywhereTrainerAdditions.OnLeftClickBank, BagOMatic: BagOMatic.cmd_showbank, BagOMatic: BagOMatic.cmd_sortall, BankArkel: BankArkel.CloseBank |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 14 |
| Global usage count | 7 |
| Local definition count | 4 |
| Documentation references | 1 |
| Initialization flow references | 4 |
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
| Conflicting signatures | no |
| Conflicting roles | no |
| Wrapper likely | no |
| Never outside local graph | no |
| Local helper only | no |

## Description

Observed shared global table or namespace surfaced in 6 addons.

## Functions

- BankWindow.GetItem
- BankWindow.GetSlotNumberForButtonIndex
- BankWindow.Hide
- BankWindow.IsShowing
- BankWindow.OpenBank
- BankWindow.Show
- BankWindow.UpdateAllBankSlots

## Observed Members

- none

## Seen In

- AnywhereTrainer
- AnywhereTrainerAdditions
- BagOMatic
- BankArkel
- BankWindowFix
- nRarity

## Examples

- AnywhereTrainer: AnywhereTrainer.OnLeftClickBank -> BankWindow.Show()
- AnywhereTrainer: AnywhereTrainer.OnLeftClickBank -> BankWindow.Hide()
- AnywhereTrainerAdditions: AnywhereTrainerAdditions.EquipmentLButtonDown -> BankWindow.GetSlotNumberForButtonIndex(buttonIndex)
- AnywhereTrainerAdditions: AnywhereTrainerAdditions.EquipmentLButtonDown -> BankWindow.GetItem(slot)
- AnywhereTrainerAdditions: AnywhereTrainerAdditions.OnLeftClickBank -> BankWindow.OpenBank()
- AnywhereTrainerAdditions: AnywhereTrainerAdditions.OnLeftClickBank -> BankWindow.Hide()

## Related APIs

- [Cursor.Clear](../functions/global_Cursor.Clear.md) (HIGH 100/100) - Global Function
- [Cursor.IconOnCursor](../functions/global_Cursor.IconOnCursor.md) (HIGH 100/100) - Global Function
- [Cursor.PickUp](../functions/global_Cursor.PickUp.md) (HIGH 100/100) - Global Function
- [DialogManager.MakeOneButtonDialog](../functions/global_DialogManager.MakeOneButtonDialog.md) (HIGH 100/100) - Global Function
- [EA_ChatWindow.InsertItemLink](../functions/global_EA_ChatWindow.InsertItemLink.md) (HIGH 100/100) - Global Function
- [EA_ChatWindow.Print](../functions/global_EA_ChatWindow.Print.md) (HIGH 100/100) - Global Function
- [WindowSetLayer](../../window_api/functions/window_WindowSetLayer.md) (HIGH 100/100) - Window Function
- [BankWindow.GetItem](../functions/global_BankWindow.GetItem.md) (HIGH 96/100) - Global Function
- [BankWindow.GetSlotNumberForButtonIndex](../functions/global_BankWindow.GetSlotNumberForButtonIndex.md) (HIGH 96/100) - Global Function
- [DestroyWindow](../functions/global_DestroyWindow.md) (HIGH 75/100) - Global Function

## Used With

- [Cursor.Clear](../functions/global_Cursor.Clear.md) (HIGH 100/100) - Global Function
- [Cursor.IconOnCursor](../functions/global_Cursor.IconOnCursor.md) (HIGH 100/100) - Global Function
- [Cursor.PickUp](../functions/global_Cursor.PickUp.md) (HIGH 100/100) - Global Function
- [DialogManager.MakeOneButtonDialog](../functions/global_DialogManager.MakeOneButtonDialog.md) (HIGH 100/100) - Global Function
- [EA_ChatWindow.InsertItemLink](../functions/global_EA_ChatWindow.InsertItemLink.md) (HIGH 100/100) - Global Function
- [EA_ChatWindow.OnKeyEnter](../functions/global_EA_ChatWindow.OnKeyEnter.md) (HIGH 100/100) - Global Function
- [EA_ChatWindow.Print](../functions/global_EA_ChatWindow.Print.md) (HIGH 100/100) - Global Function
- [EA_Window_Backpack](table_EA_Window_Backpack.md) (HIGH 100/100) - Global Table
- [SystemData.Events.INTERACT_OPEN_BANK](../../systemdata/fields/systemdata_SystemData.Events.INTERACT_OPEN_BANK.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.LOG_OUT](../../systemdata/fields/systemdata_SystemData.Events.LOG_OUT.md) (HIGH 100/100) - SystemData Field
- [WindowSetLayer](../../window_api/functions/window_WindowSetLayer.md) (HIGH 100/100) - Window Function
- [WindowSetShowing](../../window_api/functions/window_WindowSetShowing.md) (HIGH 100/100) - Window Function
- [BankWindow.GetItem](../functions/global_BankWindow.GetItem.md) (HIGH 96/100) - Global Function
- [BankWindow.GetSlotNumberForButtonIndex](../functions/global_BankWindow.GetSlotNumberForButtonIndex.md) (HIGH 96/100) - Global Function
- [BankWindow.Hide](../functions/global_BankWindow.Hide.md) (HIGH 88/100) - Global Function
- [CreateWindow](../functions/global_CreateWindow.md) (HIGH 75/100) - Global Function
- [RegisterEventHandler](../functions/global_RegisterEventHandler.md) (MEDIUM 68/100) - Global Function

## Triggered By

- none

## Affects

- none

## Notes

- none
