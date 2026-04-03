# EA_Window_Backpack

- Category: Global Table
- Confidence level: HIGH
- Confidence score: 100/100

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 125

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, seen in 4 or more addons, matches a known engine namespace.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | AuctionStats, AutoSalvage, BankArkel, BankWindowFix, EZCraftX, Info_Loot, ItemRack, JunkDump |
| Files seen in | Modules/UI/Shinies-UI-Browse.lua, Source/BankWindowFix.lua, Source/EZCraftX.lua, Source/Nisp.lua |
| Namespaces detected | EA_Window_Backpack |
| Source kinds | lua_calls |
| Example locations | AuctionStats: EA_Window_Backpack_EquipmentLButtonDown_Override, AuctionStats: PutUpForAuction, AutoSalvage: IsItemSalvagable, AutoSalvage: OnUpdate, BankArkel: UpdateData, BankWindowFix: BagEquipmentRButtonUp |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 34 |
| Global usage count | 13 |
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

Shared function table with 13 member functions; the primary API surface for 15 addons.

## Functions

- EA_Window_Backpack.AutoAddCraftingItemIfPossible
- EA_Window_Backpack.ConfirmThenRefine
- EA_Window_Backpack.GetCursorForBackpack
- EA_Window_Backpack.GetItemsFromBackpack
- EA_Window_Backpack.GetPocketName
- EA_Window_Backpack.GetPocketNumberForSlot
- EA_Window_Backpack.GetSlotFromActionButtonGroup
- EA_Window_Backpack.IsBackpackFull
- EA_Window_Backpack.IsRefinable
- EA_Window_Backpack.IsSlotLocked
- EA_Window_Backpack.ReleaseLockForSlot
- EA_Window_Backpack.RequestLockForSlot
- EA_Window_Backpack.UpdateBackpackSlots

## Observed Members

- none

## Seen In

- AuctionStats
- AutoSalvage
- BankArkel
- BankWindowFix
- EZCraftX
- Info_Loot
- ItemRack
- JunkDump
- Killer
- LootAlert
- Mass Refine
- Miracle Grow Remix
- NPC Item Sale Price
- Shinies
- zMailMod

## Examples

- AuctionStats: EA_Window_Backpack_EquipmentLButtonDown_Override -> EA_Window_Backpack.GetItemsFromBackpack(EA_Window_Backpack.currentMode)
- AuctionStats: PutUpForAuction -> EA_Window_Backpack.RequestLockForSlot(inventorySlot.slot, inventorySlot.backpack, "AuctionHouseWindowCreateAuction")
- AuctionStats: PutUpForAuction -> EA_Window_Backpack.ReleaseLockForSlot(CreateAuctionWindow.itemInventorySlot.slot, CreateAuctionWindow.itemInventorySlot.backpack, "AuctionHouseWindowCreateAuction")
- AutoSalvage: IsItemSalvagable -> EA_Window_Backpack.IsSlotLocked(selectedSlot, selectedBackpack)
- AutoSalvage: OnUpdate -> EA_Window_Backpack.GetItemsFromBackpack(selectedBackpack)
- BankArkel: UpdateData -> EA_Window_Backpack.GetItemsFromBackpack(Read[i])

## Notes

- none
