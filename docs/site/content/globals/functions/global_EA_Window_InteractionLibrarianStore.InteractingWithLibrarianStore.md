# EA_Window_InteractionLibrarianStore.InteractingWithLibrarianStore

- Category: Global Function
- Confidence level: HIGH
- Confidence score: 80/100
- Seen in: 1 addons

## Confidence Assessment

- Level: HIGH

- Score: 80/100

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, matches a known engine namespace, called globally with no local definition.

## Evidence Signals

- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | BankWindowFix |
| Files seen in | `/workspace_addons/BankWindowFix/Source/BankWindowFix.lua:108` |
| Namespaces detected | EA_Window_InteractionLibrarianStore |
| Source kinds | lua_calls |
| Example locations | BankWindowFix: BankWindowFix.BagEquipmentRButtonUp |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 1 |
| Global usage count | 1 |
| Local definition count | 0 |
| Documentation references | 0 |
| Initialization flow references | 0 |
| Known engine namespace | yes |
| Default UI presence | yes |
| Event binding presence | no |
| Observed in XML and Lua | no |
| Consistent role | no |
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

## Signature (inferred)

```lua
EA_Window_InteractionLibrarianStore.InteractingWithLibrarianStore()
```

## Description

Observed as a global function across 1 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |

## Returns

- Not confidently inferable from addon-api docs alone.

## Side Effects

- No side effect is confidently inferable from addon-api docs alone.

## Seen In

- BankWindowFix

## Examples

- BankWindowFix: BankWindowFix.BagEquipmentRButtonUp -> EA_Window_InteractionLibrarianStore.InteractingWithLibrarianStore()

## Related APIs

- none

## Used With

- [DialogManager.MakeTwoButtonDialog](global_DialogManager.MakeTwoButtonDialog.md) (HIGH 100/100) - Global Function
- [EA_Window_Backpack.AutoAddCraftingItemIfPossible](global_EA_Window_Backpack.AutoAddCraftingItemIfPossible.md) (HIGH 100/100) - Global Function
- [EA_Window_Backpack.ConfirmThenRefine](global_EA_Window_Backpack.ConfirmThenRefine.md) (HIGH 100/100) - Global Function
- [EA_Window_Backpack.GetCursorForBackpack](global_EA_Window_Backpack.GetCursorForBackpack.md) (HIGH 100/100) - Global Function
- [EA_Window_Backpack.GetItemsFromBackpack](global_EA_Window_Backpack.GetItemsFromBackpack.md) (HIGH 100/100) - Global Function
- [EA_Window_Backpack.IsRefinable](global_EA_Window_Backpack.IsRefinable.md) (HIGH 100/100) - Global Function
- [EA_Window_Backpack.IsSlotLocked](global_EA_Window_Backpack.IsSlotLocked.md) (HIGH 100/100) - Global Function
- [EA_Window_InteractionStore](../tables/table_EA_Window_InteractionStore.md) (HIGH 100/100) - Global Table
- [GameData.ItemLocs.INVENTORY](../../gamedata/fields/gamedata_GameData.ItemLocs.INVENTORY.md) (HIGH 100/100) - GameData Field
- [WindowGetShowing](../../window_api/functions/window_WindowGetShowing.md) (HIGH 100/100) - Window Function
- [BankWindow.IsShowing](global_BankWindow.IsShowing.md) (HIGH 96/100) - Global Function
- [EA_Window_InteractionLibrarianStore.ConfirmThenRepairItem](global_EA_Window_InteractionLibrarianStore.ConfirmThenRepairItem.md) (HIGH 80/100) - Global Function
- [EA_Window_InteractionLibrarianStore.ConfirmThenSellItem](global_EA_Window_InteractionLibrarianStore.ConfirmThenSellItem.md) (HIGH 80/100) - Global Function
- [EA_Window_InteractionLibrarianStore.InteractingWithRepairMan](global_EA_Window_InteractionLibrarianStore.InteractingWithRepairMan.md) (HIGH 80/100) - Global Function

## Triggered By

- none

## Affects

- [EA_Window_Backpack.GetCursorForBackpack](global_EA_Window_Backpack.GetCursorForBackpack.md) (HIGH 100/100) - Global Function
- [EA_Window_Backpack.GetItemsFromBackpack](global_EA_Window_Backpack.GetItemsFromBackpack.md) (HIGH 100/100) - Global Function
- [EA_Window_InteractionStore](../tables/table_EA_Window_InteractionStore.md) (HIGH 100/100) - Global Table
- [GameData.ItemLocs.INVENTORY](../../gamedata/fields/gamedata_GameData.ItemLocs.INVENTORY.md) (HIGH 100/100) - GameData Field
- [BankWindow.IsShowing](global_BankWindow.IsShowing.md) (HIGH 96/100) - Global Function

## Notes

- Only one addon surfaced this symbol in the current corpus.
- Canonical entry built from observed call sites, not from engine source or decompiled definitions.
