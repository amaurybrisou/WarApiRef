# EA_BackpackUtilsMediator

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
| Addons seen in | AuctionStats, BagOMatic, Crusher, FozAuction, Motion, TidyRoll, zMailMod |
| Files seen in | Source/Crusher.lua, Source/Motion.lua, Source/auctionwindowsellcontrols.lua |
| Namespaces detected | EA_BackpackUtilsMediator |
| Source kinds | lua_calls |
| Example locations | AuctionStats: OnSearchResultsReceived, AuctionStats: PutUpForAuction, BagOMatic: SalvageHook, BagOMatic: findItemInBagPack, Crusher: StopCrushing, Crusher: continueCrushing |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 32 |
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

Shared function table with 7 member functions; the primary API surface for 7 addons.

## Functions

- EA_BackpackUtilsMediator.GetBackpack
- EA_BackpackUtilsMediator.GetCurrentBackpackType
- EA_BackpackUtilsMediator.GetCursorForBackpack
- EA_BackpackUtilsMediator.GetItemsFromBackpack
- EA_BackpackUtilsMediator.ReleaseAllLocksForWindow
- EA_BackpackUtilsMediator.ReleaseLockForSlot
- EA_BackpackUtilsMediator.RequestLockForSlot

## Observed Members

- none

## Seen In

- AuctionStats
- BagOMatic
- Crusher
- FozAuction
- Motion
- TidyRoll
- zMailMod

## Examples

- AuctionStats: OnSearchResultsReceived -> EA_BackpackUtilsMediator.GetItemsFromBackpack(CreateAuctionWindow.itemInventorySlot.backpack)
- AuctionStats: PutUpForAuction -> EA_BackpackUtilsMediator.GetItemsFromBackpack(CreateAuctionWindow.itemInventorySlot.backpack)
- BagOMatic: SalvageHook -> EA_BackpackUtilsMediator.GetBackpack()
- BagOMatic: SalvageHook -> EA_BackpackUtilsMediator.GetItemsFromBackpack(backpackType)
- BagOMatic: findItemInBagPack -> EA_BackpackUtilsMediator.GetCurrentBackpackType()
- BagOMatic: findItemInBagPack -> EA_BackpackUtilsMediator.GetItemsFromBackpack(backpackType)

## Notes

- none
