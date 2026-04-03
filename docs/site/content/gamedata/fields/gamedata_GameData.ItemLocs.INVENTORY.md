# GameData.ItemLocs.INVENTORY

- Category: GameData Field
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 5 addons

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 150

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, seen in 4 or more addons, matches a known engine namespace.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +25 Referenced by generated docs or reference files: The symbol is reinforced outside a single call page.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | BankWindowFix, Crusher, PotionBar, RandomMount, Shinies |
| Files seen in | Modules/API/Shinies-API-Inventory.lua, Modules/Data/Shinies-Data-Inventory.lua, RandomMountCore.lua, Source/BankWindowFix.lua, Source/Crusher.lua, source/Floating.lua |
| Namespaces detected | GameData |
| Source kinds | lua_call |
| Example locations | BagEquipmentRButtonUp, Inventory_GetAvailableSlot, Inventory_GetItemSlotByStackCount, Inventory_GetItemSlots, Inventory_TotalAvailableSlots, Inventory_TotalItemCount |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 12 |
| Global usage count | 12 |
| Local definition count | 0 |
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

GameData.GameData.ItemLocs.INVENTORY field accessed by 5 addons; commonly found in BagEquipmentRButtonUp and Inventory_GetAvailableSlot, Inventory_GetItemSlotByStackCount, Inventory_GetItemSlots, Inventory_TotalAvailableSlots, Inventory_TotalItemCount, LButtonUp, Mount, OnPlayerInventorySlotUpdated, getAllAvailableSlotsForItem, getTotalItemCount, lua_call contexts.

## Seen In

- BankWindowFix
- Crusher
- PotionBar
- RandomMount
- Shinies

## Related APIs

- [DialogManager.MakeTwoButtonDialog](../../globals/functions/global_DialogManager.MakeTwoButtonDialog.md) (HIGH 100/100) - Global Function
- [EA_Window_Backpack.GetCursorForBackpack](../../globals/functions/global_EA_Window_Backpack.GetCursorForBackpack.md) (HIGH 100/100) - Global Function
- [EA_Window_Backpack.GetItemsFromBackpack](../../globals/functions/global_EA_Window_Backpack.GetItemsFromBackpack.md) (HIGH 100/100) - Global Function
- [EA_Window_Backpack.IsSlotLocked](../../globals/functions/global_EA_Window_Backpack.IsSlotLocked.md) (HIGH 100/100) - Global Function
- [EA_Window_Backpack.AutoAddCraftingItemIfPossible](../../globals/functions/global_EA_Window_Backpack.AutoAddCraftingItemIfPossible.md) (HIGH 90/100) - Global Function
- [EA_Window_Backpack.ConfirmThenRefine](../../globals/functions/global_EA_Window_Backpack.ConfirmThenRefine.md) (HIGH 80/100) - Global Function
- [EA_Window_Backpack.IsRefinable](../../globals/functions/global_EA_Window_Backpack.IsRefinable.md) (HIGH 80/100) - Global Function
- [EA_Window_InteractionLibrarianStore.ConfirmThenRepairItem](../../globals/functions/global_EA_Window_InteractionLibrarianStore.ConfirmThenRepairItem.md) (HIGH 80/100) - Global Function
- [EA_Window_InteractionLibrarianStore.ConfirmThenSellItem](../../globals/functions/global_EA_Window_InteractionLibrarianStore.ConfirmThenSellItem.md) (HIGH 80/100) - Global Function
- [EA_Window_InteractionLibrarianStore.InteractingWithLibrarianStore](../../globals/functions/global_EA_Window_InteractionLibrarianStore.InteractingWithLibrarianStore.md) (HIGH 80/100) - Global Function
- [EA_Window_InteractionLibrarianStore.InteractingWithRepairMan](../../globals/functions/global_EA_Window_InteractionLibrarianStore.InteractingWithRepairMan.md) (HIGH 80/100) - Global Function
- [EA_Window_InteractionStore.ConfirmThenRepairItem](../../globals/functions/global_EA_Window_InteractionStore.ConfirmThenRepairItem.md) (HIGH 80/100) - Global Function
- [EA_Window_InteractionStore.ConfirmThenSellItem](../../globals/functions/global_EA_Window_InteractionStore.ConfirmThenSellItem.md) (HIGH 80/100) - Global Function
- [EA_Window_InteractionStore.InteractingWithRepairMan](../../globals/functions/global_EA_Window_InteractionStore.InteractingWithRepairMan.md) (HIGH 80/100) - Global Function
- [EA_Window_InteractionStore.InteractingWithStore](../../globals/functions/global_EA_Window_InteractionStore.InteractingWithStore.md) (HIGH 80/100) - Global Function
- [EA_Window_Trade.AddInventoryItem](../../globals/functions/global_EA_Window_Trade.AddInventoryItem.md) (HIGH 80/100) - Global Function
- [EA_Window_Trade.ClearInventoryItem](../../globals/functions/global_EA_Window_Trade.ClearInventoryItem.md) (HIGH 80/100) - Global Function
- [EA_Window_Trade.TradeOpen](../../globals/functions/global_EA_Window_Trade.TradeOpen.md) (HIGH 80/100) - Global Function
- [SendUseItem](../../globals/functions/global_SendUseItem.md) (HIGH 75/100) - Global Function
- [BankWindow.IsShowing](../../globals/functions/global_BankWindow.IsShowing.md) (HIGH 71/100) - Global Function

## Used With

- [DialogManager.MakeTwoButtonDialog](../../globals/functions/global_DialogManager.MakeTwoButtonDialog.md) (HIGH 100/100) - Global Function
- [EA_Window_Backpack.GetCursorForBackpack](../../globals/functions/global_EA_Window_Backpack.GetCursorForBackpack.md) (HIGH 100/100) - Global Function
- [EA_Window_Backpack.AutoAddCraftingItemIfPossible](../../globals/functions/global_EA_Window_Backpack.AutoAddCraftingItemIfPossible.md) (HIGH 90/100) - Global Function
- [EA_Window_Backpack.ConfirmThenRefine](../../globals/functions/global_EA_Window_Backpack.ConfirmThenRefine.md) (HIGH 80/100) - Global Function
- [BankWindow.IsShowing](../../globals/functions/global_BankWindow.IsShowing.md) (HIGH 71/100) - Global Function

## Notes

- Observed in contexts: BagEquipmentRButtonUp, Inventory_GetAvailableSlot, Inventory_GetItemSlotByStackCount, Inventory_GetItemSlots, Inventory_TotalAvailableSlots, Inventory_TotalItemCount
