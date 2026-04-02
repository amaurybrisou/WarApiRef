# EA_Window_Backpack.GetCursorForBackpack

- Category: Global Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 2 addons

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 156

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, matches a known engine namespace, referenced by generated docs or reference files.

## Evidence Signals

- +18 Seen in 2 to 3 addons: Cross-addon spread is present but limited.
- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +10 Argument pattern is consistent: Observed argument positions remain stable.
- +8 Return usage is consistent: Observed as a stable query-style API.
- +25 Referenced by generated docs or reference files: The symbol is reinforced outside a single call page.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | BankWindowFix, Miracle Grow Remix |
| Files seen in | `/workspace_addons/BankWindowFix/Source/BankWindowFix.lua:108`, `/workspace_addons/MGRemix/MGRemix.lua:996` |
| Namespaces detected | EA_Window_Backpack |
| Source kinds | globals, lua_calls |
| Example locations | BankWindowFix: BankWindowFix.BagEquipmentRButtonUp, Miracle Grow Remix: MiracleGrow2.GetCursorForBackpack |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 2 |
| Global usage count | 2 |
| Local definition count | 0 |
| Documentation references | 1 |
| Initialization flow references | 0 |
| Known engine namespace | yes |
| Default UI presence | yes |
| Event binding presence | no |
| Observed in XML and Lua | no |
| Consistent role | yes |
| Consistent arguments | yes |
| Consistent returns | yes |
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
EA_Window_Backpack.GetCursorForBackpack(arg1)
```

## Description

Observed as a global function across 2 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a function or method reference. | Observed values: EA_Window_Backpack.currentMode, nBackpackType |

## Returns

- Observed as a query-style API. The concrete return shape is not inferable from addon-api docs alone.

## Side Effects

- No side effect is confidently inferable from addon-api docs alone.

## Seen In

- BankWindowFix
- Miracle Grow Remix

## Examples

- BankWindowFix: BankWindowFix.BagEquipmentRButtonUp -> EA_Window_Backpack.GetCursorForBackpack(EA_Window_Backpack.currentMode)
- Miracle Grow Remix: MiracleGrow2.GetCursorForBackpack -> EA_Window_Backpack.GetCursorForBackpack(nBackpackType)

## Related APIs

- [DialogManager.MakeTwoButtonDialog](global_DialogManager.MakeTwoButtonDialog.md) (HIGH 100/100) - Global Function
- [EA_Window_Backpack.AutoAddCraftingItemIfPossible](global_EA_Window_Backpack.AutoAddCraftingItemIfPossible.md) (HIGH 100/100) - Global Function
- [EA_Window_Backpack.ConfirmThenRefine](global_EA_Window_Backpack.ConfirmThenRefine.md) (HIGH 100/100) - Global Function
- [EA_Window_Backpack.GetItemsFromBackpack](global_EA_Window_Backpack.GetItemsFromBackpack.md) (HIGH 100/100) - Global Function
- [EA_Window_Backpack.IsRefinable](global_EA_Window_Backpack.IsRefinable.md) (HIGH 100/100) - Global Function
- [EA_Window_Backpack.IsSlotLocked](global_EA_Window_Backpack.IsSlotLocked.md) (HIGH 100/100) - Global Function
- [BankWindow.IsShowing](global_BankWindow.IsShowing.md) (HIGH 96/100) - Global Function
- [EA_Window_InteractionLibrarianStore.ConfirmThenRepairItem](global_EA_Window_InteractionLibrarianStore.ConfirmThenRepairItem.md) (HIGH 80/100) - Global Function
- [EA_Window_InteractionLibrarianStore.ConfirmThenSellItem](global_EA_Window_InteractionLibrarianStore.ConfirmThenSellItem.md) (HIGH 80/100) - Global Function
- [EA_Window_InteractionLibrarianStore.InteractingWithLibrarianStore](global_EA_Window_InteractionLibrarianStore.InteractingWithLibrarianStore.md) (HIGH 80/100) - Global Function
- [EA_Window_InteractionLibrarianStore.InteractingWithRepairMan](global_EA_Window_InteractionLibrarianStore.InteractingWithRepairMan.md) (HIGH 80/100) - Global Function

## Used With

- none

## Triggered By

- none

## Affects

- [EA_Window_Backpack.GetItemsFromBackpack](global_EA_Window_Backpack.GetItemsFromBackpack.md) (HIGH 100/100) - Global Function
- [EA_Window_InteractionStore](../tables/table_EA_Window_InteractionStore.md) (HIGH 100/100) - Global Table
- [GameData.ItemLocs.INVENTORY](../../gamedata/fields/gamedata_GameData.ItemLocs.INVENTORY.md) (HIGH 100/100) - GameData Field
- [BankWindow.IsShowing](global_BankWindow.IsShowing.md) (HIGH 96/100) - Global Function

## Notes

- Canonical entry built from observed call sites, not from engine source or decompiled definitions.
- Advanced return analysis: No strong return evidence observed
