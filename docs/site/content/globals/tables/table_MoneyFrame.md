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
| Addons seen in | AdvancedRenownTrainer, BankArkel, HealAll, JunkDump |
| Files seen in | `/workspace_addons/BankArkel/BankArkel.lua:242`, `/workspace_addons/HealAll/HealAll.lua:8`, `/workspace_addons/JunkDump/JunkDump.lua:525`, `/workspace_addons/advancedrenowntrainer/AdvancedRenownTraining.lua:763` |
| Namespaces detected | MoneyFrame |
| Source kinds | lua_calls |
| Example locations | AdvancedRenownTrainer: AdvancedRenownTraining.Respecialize, BankArkel: BankArkel.PackShow, HealAll: HealAll.Heal, JunkDump: JunkDump.showMeDaMoney |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 4 |
| Global usage count | 2 |
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

Observed shared global table or namespace surfaced in 4 addons.

## Functions

- MoneyFrame.FormatMoney
- MoneyFrame.FormatMoneyString

## Observed Members

- none

## Seen In

- AdvancedRenownTrainer
- BankArkel
- HealAll
- JunkDump

## Examples

- AdvancedRenownTrainer: AdvancedRenownTraining.Respecialize -> MoneyFrame.FormatMoneyString(respecCost)
- BankArkel: BankArkel.PackShow -> MoneyFrame.FormatMoney("PackWinMoney", BankArkel.db.Entry[index].Money, MoneyFrame.SHOW_EMPTY_WINDOWS)
- HealAll: HealAll.Heal -> MoneyFrame.FormatMoneyString(EA_Window_InteractionHealer.costToRemoveSinglePenalty*EA_Window_InteractionHealer.penaltyCount)
- JunkDump: JunkDump.showMeDaMoney -> MoneyFrame.FormatMoneyString(totalMoney)

## Related APIs

- none

## Used With

- none

## Triggered By

- none

## Affects

- none

## Notes

- none
