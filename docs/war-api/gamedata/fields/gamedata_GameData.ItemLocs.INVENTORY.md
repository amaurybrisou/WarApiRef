# GameData.ItemLocs.INVENTORY

- Category: GameData Field
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 4 addons

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
| Addons seen in | BankWindowFix, PotionBar, RandomMount, Shinies |
| Files seen in | `/workspace_addons/BankWindowFix/Source/BankWindowFix.lua:108`, `/workspace_addons/PotionBar/source/Floating.lua:33`, `/workspace_addons/RandomMount/RandomMountCore.lua:203`, `/workspace_addons/Shinies/Modules/API/Shinies-API-Inventory.lua:114`, `/workspace_addons/Shinies/Modules/API/Shinies-API-Inventory.lua:135`, `/workspace_addons/Shinies/Modules/API/Shinies-API-Inventory.lua:167`, `/workspace_addons/Shinies/Modules/API/Shinies-API-Inventory.lua:193`, `/workspace_addons/Shinies/Modules/API/Shinies-API-Inventory.lua:215` |
| Namespaces detected | GameData |
| Source kinds | lua_call |
| Example locations | BankWindowFix.BagEquipmentRButtonUp, PotionBarFloating.LButtonUp, RandomMount.Mount, ShiniesDataInventory.OnPlayerInventorySlotUpdated, _G.Shinies:GetModule:Inventory_GetAvailableSlot, _G.Shinies:GetModule:Inventory_GetItemSlotByStackCount |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 10 |
| Global usage count | 10 |
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

Observed GameData field used by 4 addons through generated function calls, event pages, or lifecycle evidence.

## Seen In

- BankWindowFix
- PotionBar
- RandomMount
- Shinies

## Related APIs

- [DialogManager.MakeTwoButtonDialog](../../globals/functions/global_DialogManager.MakeTwoButtonDialog.md) (HIGH 100/100) - Global Function
- [EA_Window_Backpack.AutoAddCraftingItemIfPossible](../../globals/functions/global_EA_Window_Backpack.AutoAddCraftingItemIfPossible.md) (HIGH 100/100) - Global Function
- [EA_Window_Backpack.ConfirmThenRefine](../../globals/functions/global_EA_Window_Backpack.ConfirmThenRefine.md) (HIGH 100/100) - Global Function
- [EA_Window_Backpack.GetCursorForBackpack](../../globals/functions/global_EA_Window_Backpack.GetCursorForBackpack.md) (HIGH 100/100) - Global Function
- [EA_Window_Backpack.GetItemsFromBackpack](../../globals/functions/global_EA_Window_Backpack.GetItemsFromBackpack.md) (HIGH 100/100) - Global Function
- [EA_Window_Backpack.IsRefinable](../../globals/functions/global_EA_Window_Backpack.IsRefinable.md) (HIGH 100/100) - Global Function
- [EA_Window_Backpack.IsSlotLocked](../../globals/functions/global_EA_Window_Backpack.IsSlotLocked.md) (HIGH 100/100) - Global Function
- [BankWindow.IsShowing](../../globals/functions/global_BankWindow.IsShowing.md) (HIGH 96/100) - Global Function
- [EA_Window_InteractionLibrarianStore.ConfirmThenRepairItem](../../globals/functions/global_EA_Window_InteractionLibrarianStore.ConfirmThenRepairItem.md) (HIGH 80/100) - Global Function
- [EA_Window_InteractionLibrarianStore.ConfirmThenSellItem](../../globals/functions/global_EA_Window_InteractionLibrarianStore.ConfirmThenSellItem.md) (HIGH 80/100) - Global Function
- [EA_Window_InteractionLibrarianStore.InteractingWithLibrarianStore](../../globals/functions/global_EA_Window_InteractionLibrarianStore.InteractingWithLibrarianStore.md) (HIGH 80/100) - Global Function
- [EA_Window_InteractionLibrarianStore.InteractingWithRepairMan](../../globals/functions/global_EA_Window_InteractionLibrarianStore.InteractingWithRepairMan.md) (HIGH 80/100) - Global Function

## Used With

- [DialogManager.MakeTwoButtonDialog](../../globals/functions/global_DialogManager.MakeTwoButtonDialog.md) (HIGH 100/100) - Global Function
- [EA_Window_Backpack.AutoAddCraftingItemIfPossible](../../globals/functions/global_EA_Window_Backpack.AutoAddCraftingItemIfPossible.md) (HIGH 100/100) - Global Function
- [EA_Window_Backpack.ConfirmThenRefine](../../globals/functions/global_EA_Window_Backpack.ConfirmThenRefine.md) (HIGH 100/100) - Global Function
- [EA_Window_Backpack.GetCursorForBackpack](../../globals/functions/global_EA_Window_Backpack.GetCursorForBackpack.md) (HIGH 100/100) - Global Function
- [EA_Window_Backpack.GetItemsFromBackpack](../../globals/functions/global_EA_Window_Backpack.GetItemsFromBackpack.md) (HIGH 100/100) - Global Function
- [EA_Window_Backpack.IsRefinable](../../globals/functions/global_EA_Window_Backpack.IsRefinable.md) (HIGH 100/100) - Global Function
- [EA_Window_Backpack.IsSlotLocked](../../globals/functions/global_EA_Window_Backpack.IsSlotLocked.md) (HIGH 100/100) - Global Function
- [EA_Window_InteractionStore](../../globals/tables/table_EA_Window_InteractionStore.md) (HIGH 100/100) - Global Table
- [WindowGetShowing](../../window_api/functions/window_WindowGetShowing.md) (HIGH 100/100) - Window Function
- [BankWindow.IsShowing](../../globals/functions/global_BankWindow.IsShowing.md) (HIGH 96/100) - Global Function
- [EA_Window_InteractionLibrarianStore.ConfirmThenRepairItem](../../globals/functions/global_EA_Window_InteractionLibrarianStore.ConfirmThenRepairItem.md) (HIGH 80/100) - Global Function
- [EA_Window_InteractionLibrarianStore.ConfirmThenSellItem](../../globals/functions/global_EA_Window_InteractionLibrarianStore.ConfirmThenSellItem.md) (HIGH 80/100) - Global Function
- [EA_Window_InteractionLibrarianStore.InteractingWithLibrarianStore](../../globals/functions/global_EA_Window_InteractionLibrarianStore.InteractingWithLibrarianStore.md) (HIGH 80/100) - Global Function
- [EA_Window_InteractionLibrarianStore.InteractingWithRepairMan](../../globals/functions/global_EA_Window_InteractionLibrarianStore.InteractingWithRepairMan.md) (HIGH 80/100) - Global Function

## Triggered By

- none

## Affects

- none

## Notes

- Observed in contexts: BankWindowFix.BagEquipmentRButtonUp, PotionBarFloating.LButtonUp, RandomMount.Mount, ShiniesDataInventory.OnPlayerInventorySlotUpdated, _G.Shinies:GetModule:Inventory_GetAvailableSlot, _G.Shinies:GetModule:Inventory_GetItemSlotByStackCount
