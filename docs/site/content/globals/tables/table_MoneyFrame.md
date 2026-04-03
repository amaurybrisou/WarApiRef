# MoneyFrame

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
| Addons seen in | AdvancedRenownTrainer, AuctionStats, BankArkel, FozAuction, HealAll, JunkDump, zMailMod |
| Files seen in | Source/auctionwindow.lua, Source/auctionwindowsellcontrols.lua |
| Namespaces detected | MoneyFrame |
| Source kinds | lua_calls |
| Example locations | AdvancedRenownTrainer: Respecialize, AuctionStats: AddExtraWindow, AuctionStats: OnRButtonUpItem, AuctionStats: UpdateItemOptions, BankArkel: PackShow, FozAuction: Clear |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 25 |
| Global usage count | 4 |
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

Shared function table with 4 member functions; the primary API surface for 7 addons.

## Functions

- MoneyFrame.ConvertBrassToCurrency
- MoneyFrame.ConvertCurrencyToBrass
- MoneyFrame.FormatMoney
- MoneyFrame.FormatMoneyString

## Observed Members

- none

## Seen In

- AdvancedRenownTrainer
- AuctionStats
- BankArkel
- FozAuction
- HealAll
- JunkDump
- zMailMod

## Examples

- AdvancedRenownTrainer: Respecialize -> MoneyFrame.FormatMoneyString(respecCost)
- AuctionStats: AddExtraWindow -> MoneyFrame.ConvertBrassToCurrency(AuctionDB.Data[id].minBuyOutPrice*stackCount)
- AuctionStats: AddExtraWindow -> MoneyFrame.ConvertBrassToCurrency(AuctionDB.Data[id].maxBuyOutPrice*stackCount)
- AuctionStats: AddExtraWindow -> MoneyFrame.ConvertBrassToCurrency(AuctionDB.Data[id].minBidPrice*stackCount)
- AuctionStats: AddExtraWindow -> MoneyFrame.ConvertBrassToCurrency(AuctionDB.Data[id].maxBidPrice*stackCount)
- AuctionStats: AddExtraWindow -> MoneyFrame.ConvertBrassToCurrency(median)

## Notes

- none
