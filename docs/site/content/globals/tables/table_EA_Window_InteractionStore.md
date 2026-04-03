# EA_Window_InteractionStore

- Category: Global Table
- Confidence level: HIGH
- Confidence score: 100/100

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 113

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, matches a known engine namespace, called globally with no local definition.

## Evidence Signals

- +18 Seen in 2 to 3 addons: Cross-addon spread is present but limited.
- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | BagOMatic, BankWindowFix, nRarity |
| Files seen in | Source/BankWindowFix.lua, source/containers/InteractionStore.lua |
| Namespaces detected | EA_Window_InteractionStore |
| Source kinds | lua_calls |
| Example locations | BagOMatic: buyItem, BagOMatic: buyLastItem, BagOMatic: sellItemAt, BankWindowFix: BagEquipmentRButtonUp, nRarity: InteractionStore_EA_Window_InteractionStore_OnItemLButtonDown |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 8 |
| Global usage count | 7 |
| Local definition count | 0 |
| Documentation references | 0 |
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

Shared function table with 7 member functions; the primary API surface for 3 addons.

## Functions

- EA_Window_InteractionStore.ConfirmThenBuyItem
- EA_Window_InteractionStore.ConfirmThenRepairItem
- EA_Window_InteractionStore.ConfirmThenSellItem
- EA_Window_InteractionStore.InteractingWithRepairMan
- EA_Window_InteractionStore.InteractingWithStore
- EA_Window_InteractionStore.OnItemLButtonDown
- EA_Window_InteractionStore.SellItem

## Observed Members

- none

## Seen In

- BagOMatic
- BankWindowFix
- nRarity

## Examples

- BagOMatic: buyItem -> EA_Window_InteractionStore.ConfirmThenBuyItem(item_slot_num, 1)
- BagOMatic: buyLastItem -> EA_Window_InteractionStore.ConfirmThenBuyItem(13, 1)
- BagOMatic: sellItemAt -> EA_Window_InteractionStore.SellItem(slot_id, 1)
- BankWindowFix: BagEquipmentRButtonUp -> EA_Window_InteractionStore.InteractingWithStore()
- BankWindowFix: BagEquipmentRButtonUp -> EA_Window_InteractionStore.InteractingWithRepairMan()
- BankWindowFix: BagEquipmentRButtonUp -> EA_Window_InteractionStore.ConfirmThenRepairItem(slot)

## Notes

- none
