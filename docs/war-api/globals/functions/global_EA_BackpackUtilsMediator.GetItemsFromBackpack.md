# EA_BackpackUtilsMediator.GetItemsFromBackpack

- Category: Global Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 4 addons

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 143

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, seen in 4 or more addons, matches a known engine namespace.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +10 Argument pattern is consistent: Observed argument positions remain stable.
- +8 Return usage is consistent: Observed as a stable query-style API.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | AuctionStats, BagOMatic, FozAuction, TidyRoll |
| Files seen in | AuctionAssist.lua, BagOMatic.lua, CustomAutoRoll.lua, Source/auctionwindowsellcontrols.lua |
| Namespaces detected | EA_BackpackUtilsMediator |
| Source kinds | lua_calls |
| Example locations | AuctionStats: OnSearchResultsReceived, AuctionStats: PutUpForAuction, BagOMatic: SalvageHook, BagOMatic: findItemInBagPack, FozAuction: DropItemIfPossible, FozAuction: ItemSlotMouseOver |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 8 |
| Global usage count | 8 |
| Local definition count | 0 |
| Documentation references | 0 |
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
EA_BackpackUtilsMediator.GetItemsFromBackpack(arg1)
```

## Description

Observed as a global function across 4 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a function or method reference. | Observed values: AuctionWindowSellControls.itemInventorySlot.backpack, CreateAuctionWindow.itemInventorySlot.backpack, backpackType |

## Returns

- Observed as a query-style API. The concrete return shape is not inferable from contract artifacts alone.

## Side Effects

- No side effect is confidently inferable from contract artifacts alone.

## Seen In

- AuctionStats
- BagOMatic
- FozAuction
- TidyRoll

## Examples

- AuctionStats: OnSearchResultsReceived -> EA_BackpackUtilsMediator.GetItemsFromBackpack(CreateAuctionWindow.itemInventorySlot.backpack)
- AuctionStats: PutUpForAuction -> EA_BackpackUtilsMediator.GetItemsFromBackpack(CreateAuctionWindow.itemInventorySlot.backpack)
- BagOMatic: SalvageHook -> EA_BackpackUtilsMediator.GetItemsFromBackpack(backpackType)
- BagOMatic: findItemInBagPack -> EA_BackpackUtilsMediator.GetItemsFromBackpack(backpackType)
- FozAuction: DropItemIfPossible -> EA_BackpackUtilsMediator.GetItemsFromBackpack(backpackType)
- FozAuction: ItemSlotMouseOver -> EA_BackpackUtilsMediator.GetItemsFromBackpack(AuctionWindowSellControls.itemInventorySlot.backpack)

## Used With

- [Cursor.Clear](global_Cursor.Clear.md) (HIGH 100/100) - Global Function
- [Cursor.IconOnCursor](global_Cursor.IconOnCursor.md) (HIGH 100/100) - Global Function
- [Cursor.PickUp](global_Cursor.PickUp.md) (HIGH 100/100) - Global Function
- [EA_BackpackUtilsMediator.GetCurrentBackpackType](global_EA_BackpackUtilsMediator.GetCurrentBackpackType.md) (HIGH 100/100) - Global Function
- [EA_BackpackUtilsMediator.GetCursorForBackpack](global_EA_BackpackUtilsMediator.GetCursorForBackpack.md) (HIGH 100/100) - Global Function
- [SystemData.TrialAlert.ALERT_AUCTION](../../systemdata/fields/systemdata_SystemData.TrialAlert.ALERT_AUCTION.md) (HIGH 100/100) - SystemData Field
- [EA_BackpackUtilsMediator.ReleaseLockForSlot](global_EA_BackpackUtilsMediator.ReleaseLockForSlot.md) (HIGH 90/100) - Global Function

## Affects

- [SystemData.TrialAlert.ALERT_AUCTION](../../systemdata/fields/systemdata_SystemData.TrialAlert.ALERT_AUCTION.md) (HIGH 100/100) - SystemData Field

## Notes

- Canonical entry built from observed call sites, not from engine source or decompiled definitions.
