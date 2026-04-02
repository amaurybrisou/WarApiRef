# EA_Window_InteractionStore

- Category: Global Table
- Confidence level: HIGH
- Confidence score: 100/100

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 118

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, matches a known engine namespace, referenced by generated docs or reference files.

## Evidence Signals

- +18 Seen in 2 to 3 addons: Cross-addon spread is present but limited.
- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +25 Referenced by generated docs or reference files: The symbol is reinforced outside a single call page.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | BagOMatic, BankWindowFix, nRarity |
| Files seen in | `/workspace_addons/bagomatic/BagOMatic.lua:771`, `/workspace_addons/bagomatic/BagOMatic.lua:776`, `/workspace_addons/bagomatic/BagOMatic.lua:782`, `/workspace_addons/nRarity/source/containers/InteractionStore.lua:41` |
| Namespaces detected | EA_Window_InteractionStore |
| Source kinds | globals, lua_calls |
| Example locations | BagOMatic: BagOMatic.buyItem, BagOMatic: BagOMatic.buyLastItem, BagOMatic: BagOMatic.sellItemAt, nRarity: InteractionStore_EA_Window_InteractionStore_OnItemLButtonDown |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 4 |
| Global usage count | 3 |
| Local definition count | 1 |
| Documentation references | 1 |
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
| Conflicting signatures | no |
| Conflicting roles | no |
| Wrapper likely | no |
| Never outside local graph | no |
| Local helper only | no |

## Description

Observed shared global table or namespace surfaced in 3 addons.

## Functions

- EA_Window_InteractionStore.ConfirmThenBuyItem
- EA_Window_InteractionStore.OnItemLButtonDown
- EA_Window_InteractionStore.SellItem

## Observed Members

- none

## Seen In

- BagOMatic
- BankWindowFix
- nRarity

## Examples

- BagOMatic: BagOMatic.buyItem -> EA_Window_InteractionStore.ConfirmThenBuyItem(item_slot_num, 1)
- BagOMatic: BagOMatic.buyLastItem -> EA_Window_InteractionStore.ConfirmThenBuyItem(13, 1)
- BagOMatic: BagOMatic.sellItemAt -> EA_Window_InteractionStore.SellItem(slot_id, 1)
- nRarity: InteractionStore_EA_Window_InteractionStore_OnItemLButtonDown -> EA_Window_InteractionStore.OnItemLButtonDown(...)

## Related APIs

- [DialogManager.MakeTwoButtonDialog](../functions/global_DialogManager.MakeTwoButtonDialog.md) (HIGH 100/100) - Global Function
- [EA_Window_Backpack.AutoAddCraftingItemIfPossible](../functions/global_EA_Window_Backpack.AutoAddCraftingItemIfPossible.md) (HIGH 100/100) - Global Function
- [EA_Window_Backpack.ConfirmThenRefine](../functions/global_EA_Window_Backpack.ConfirmThenRefine.md) (HIGH 100/100) - Global Function
- [EA_Window_Backpack.GetCursorForBackpack](../functions/global_EA_Window_Backpack.GetCursorForBackpack.md) (HIGH 100/100) - Global Function
- [EA_Window_Backpack.GetItemsFromBackpack](../functions/global_EA_Window_Backpack.GetItemsFromBackpack.md) (HIGH 100/100) - Global Function
- [EA_Window_Backpack.IsRefinable](../functions/global_EA_Window_Backpack.IsRefinable.md) (HIGH 100/100) - Global Function
- [EA_Window_Backpack.IsSlotLocked](../functions/global_EA_Window_Backpack.IsSlotLocked.md) (HIGH 100/100) - Global Function
- [BankWindow.IsShowing](../functions/global_BankWindow.IsShowing.md) (HIGH 96/100) - Global Function
- [EA_Window_InteractionLibrarianStore.ConfirmThenRepairItem](../functions/global_EA_Window_InteractionLibrarianStore.ConfirmThenRepairItem.md) (HIGH 80/100) - Global Function
- [EA_Window_InteractionLibrarianStore.ConfirmThenSellItem](../functions/global_EA_Window_InteractionLibrarianStore.ConfirmThenSellItem.md) (HIGH 80/100) - Global Function
- [EA_Window_InteractionLibrarianStore.InteractingWithLibrarianStore](../functions/global_EA_Window_InteractionLibrarianStore.InteractingWithLibrarianStore.md) (HIGH 80/100) - Global Function
- [EA_Window_InteractionLibrarianStore.InteractingWithRepairMan](../functions/global_EA_Window_InteractionLibrarianStore.InteractingWithRepairMan.md) (HIGH 80/100) - Global Function

## Used With

- [DialogManager.MakeTwoButtonDialog](../functions/global_DialogManager.MakeTwoButtonDialog.md) (HIGH 100/100) - Global Function
- [EA_Window_Backpack.AutoAddCraftingItemIfPossible](../functions/global_EA_Window_Backpack.AutoAddCraftingItemIfPossible.md) (HIGH 100/100) - Global Function
- [EA_Window_Backpack.ConfirmThenRefine](../functions/global_EA_Window_Backpack.ConfirmThenRefine.md) (HIGH 100/100) - Global Function
- [GameData.ItemLocs.INVENTORY](../../gamedata/fields/gamedata_GameData.ItemLocs.INVENTORY.md) (HIGH 100/100) - GameData Field
- [BankWindow.IsShowing](../functions/global_BankWindow.IsShowing.md) (HIGH 96/100) - Global Function

## Triggered By

- none

## Affects

- none

## Notes

- none
