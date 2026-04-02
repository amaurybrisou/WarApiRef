# EA_Window_Backpack.GetItemsFromBackpack

- Category: Global Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 4 addons

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 168

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, seen in 4 or more addons, matches a known engine namespace.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
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
| Addons seen in | BankArkel, BankWindowFix, Killer, Shinies |
| Files seen in | `/workspace_addons/BankArkel/BankArkel.lua:371`, `/workspace_addons/BankWindowFix/Source/BankWindowFix.lua:108`, `/workspace_addons/Killer/KillerRenown.lua:108`, `/workspace_addons/Killer/KillerRenown.lua:127`, `/workspace_addons/Shinies/Modules/UI/Shinies-UI-Browse.lua:298` |
| Namespaces detected | EA_Window_Backpack |
| Source kinds | globals, lua_calls |
| Example locations | BankArkel: BankArkel.UpdateData, BankWindowFix: BankWindowFix.BagEquipmentRButtonUp, Killer: Killer.CaptureInitialWarCrests, Killer: Killer.OnCurrencyUpdated, Shinies: _G.Shinies:NewModule.EA_Window_Backpack_EquipmentLButtonDown |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 5 |
| Global usage count | 5 |
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
EA_Window_Backpack.GetItemsFromBackpack(arg1)
```

## Description

Observed as a global function across 4 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a function or method reference. | Observed values: EA_Window_Backpack.TYPE_CURRENCY, EA_Window_Backpack.currentMode, Read[i] |

## Returns

- Observed as a query-style API. The concrete return shape is not inferable from addon-api docs alone.

## Side Effects

- No side effect is confidently inferable from addon-api docs alone.

## Seen In

- BankArkel
- BankWindowFix
- Killer
- Shinies

## Examples

- BankArkel: BankArkel.UpdateData -> EA_Window_Backpack.GetItemsFromBackpack(Read[i])
- BankWindowFix: BankWindowFix.BagEquipmentRButtonUp -> EA_Window_Backpack.GetItemsFromBackpack(EA_Window_Backpack.currentMode)
- Killer: Killer.CaptureInitialWarCrests -> EA_Window_Backpack.GetItemsFromBackpack(EA_Window_Backpack.TYPE_CURRENCY)
- Killer: Killer.OnCurrencyUpdated -> EA_Window_Backpack.GetItemsFromBackpack(EA_Window_Backpack.TYPE_CURRENCY)
- Shinies: _G.Shinies:NewModule.EA_Window_Backpack_EquipmentLButtonDown -> EA_Window_Backpack.GetItemsFromBackpack(EA_Window_Backpack.currentMode)

## Related APIs

- [DialogManager.MakeTwoButtonDialog](global_DialogManager.MakeTwoButtonDialog.md) (HIGH 100/100) - Global Function
- [EA_Window_Backpack.AutoAddCraftingItemIfPossible](global_EA_Window_Backpack.AutoAddCraftingItemIfPossible.md) (HIGH 100/100) - Global Function
- [EA_Window_Backpack.ConfirmThenRefine](global_EA_Window_Backpack.ConfirmThenRefine.md) (HIGH 100/100) - Global Function
- [EA_Window_Backpack.GetCursorForBackpack](global_EA_Window_Backpack.GetCursorForBackpack.md) (HIGH 100/100) - Global Function
- [EA_Window_Backpack.IsRefinable](global_EA_Window_Backpack.IsRefinable.md) (HIGH 100/100) - Global Function
- [EA_Window_Backpack.IsSlotLocked](global_EA_Window_Backpack.IsSlotLocked.md) (HIGH 100/100) - Global Function
- [BankWindow.IsShowing](global_BankWindow.IsShowing.md) (HIGH 96/100) - Global Function
- [EA_Window_InteractionLibrarianStore.ConfirmThenRepairItem](global_EA_Window_InteractionLibrarianStore.ConfirmThenRepairItem.md) (HIGH 80/100) - Global Function
- [EA_Window_InteractionLibrarianStore.ConfirmThenSellItem](global_EA_Window_InteractionLibrarianStore.ConfirmThenSellItem.md) (HIGH 80/100) - Global Function
- [EA_Window_InteractionLibrarianStore.InteractingWithLibrarianStore](global_EA_Window_InteractionLibrarianStore.InteractingWithLibrarianStore.md) (HIGH 80/100) - Global Function
- [EA_Window_InteractionLibrarianStore.InteractingWithRepairMan](global_EA_Window_InteractionLibrarianStore.InteractingWithRepairMan.md) (HIGH 80/100) - Global Function

## Used With

- [DialogManager.MakeTwoButtonDialog](global_DialogManager.MakeTwoButtonDialog.md) (HIGH 100/100) - Global Function
- [EA_Window_Backpack.AutoAddCraftingItemIfPossible](global_EA_Window_Backpack.AutoAddCraftingItemIfPossible.md) (HIGH 100/100) - Global Function
- [EA_Window_Backpack.ConfirmThenRefine](global_EA_Window_Backpack.ConfirmThenRefine.md) (HIGH 100/100) - Global Function
- [EA_Window_Backpack.GetCursorForBackpack](global_EA_Window_Backpack.GetCursorForBackpack.md) (HIGH 100/100) - Global Function
- [EA_Window_Backpack.IsRefinable](global_EA_Window_Backpack.IsRefinable.md) (HIGH 100/100) - Global Function
- [EA_Window_Backpack.IsSlotLocked](global_EA_Window_Backpack.IsSlotLocked.md) (HIGH 100/100) - Global Function
- [EA_Window_InteractionStore](../tables/table_EA_Window_InteractionStore.md) (HIGH 100/100) - Global Table
- [GameData.ItemLocs.INVENTORY](../../gamedata/fields/gamedata_GameData.ItemLocs.INVENTORY.md) (HIGH 100/100) - GameData Field
- [GameData.Player](../../gamedata/fields/gamedata_GameData.Player.md) (HIGH 100/100) - GameData Field
- [GameData.Player.zone](../../gamedata/fields/gamedata_GameData.Player.zone.md) (HIGH 100/100) - GameData Field
- [WindowGetShowing](../../window_api/functions/window_WindowGetShowing.md) (HIGH 100/100) - Window Function
- [BankWindow.IsShowing](global_BankWindow.IsShowing.md) (HIGH 96/100) - Global Function
- [EA_Window_InteractionLibrarianStore.ConfirmThenRepairItem](global_EA_Window_InteractionLibrarianStore.ConfirmThenRepairItem.md) (HIGH 80/100) - Global Function
- [EA_Window_InteractionLibrarianStore.ConfirmThenSellItem](global_EA_Window_InteractionLibrarianStore.ConfirmThenSellItem.md) (HIGH 80/100) - Global Function
- [EA_Window_InteractionLibrarianStore.InteractingWithLibrarianStore](global_EA_Window_InteractionLibrarianStore.InteractingWithLibrarianStore.md) (HIGH 80/100) - Global Function
- [EA_Window_InteractionLibrarianStore.InteractingWithRepairMan](global_EA_Window_InteractionLibrarianStore.InteractingWithRepairMan.md) (HIGH 80/100) - Global Function

## Triggered By

- [SystemData.Events.PLAYER_CURRENCY_SLOT_UPDATED](../../events/game_events/game_event_SystemData.Events.PLAYER_CURRENCY_SLOT_UPDATED.md) (HIGH 100/100) - Game Event

## Affects

- [EA_Window_Backpack.GetCursorForBackpack](global_EA_Window_Backpack.GetCursorForBackpack.md) (HIGH 100/100) - Global Function
- [EA_Window_InteractionStore](../tables/table_EA_Window_InteractionStore.md) (HIGH 100/100) - Global Table
- [GameData.ItemLocs.INVENTORY](../../gamedata/fields/gamedata_GameData.ItemLocs.INVENTORY.md) (HIGH 100/100) - GameData Field
- [GameData.Player](../../gamedata/fields/gamedata_GameData.Player.md) (HIGH 100/100) - GameData Field
- [GameData.Player.zone](../../gamedata/fields/gamedata_GameData.Player.zone.md) (HIGH 100/100) - GameData Field
- [BankWindow.IsShowing](global_BankWindow.IsShowing.md) (HIGH 96/100) - Global Function

## Notes

- Canonical entry built from observed call sites, not from engine source or decompiled definitions.
