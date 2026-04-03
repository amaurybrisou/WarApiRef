# BankWindow

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
| Addons seen in | AnywhereTrainer, AnywhereTrainerAdditions, BagOMatic, BankArkel, BankWindowFix |
| Files seen in | Source/BankWindowFix.lua, source/AnywhereTrainer.lua |
| Namespaces detected | BankWindow |
| Source kinds | lua_calls |
| Example locations | AnywhereTrainer: OnLeftClickBank, AnywhereTrainerAdditions: EquipmentLButtonDown, AnywhereTrainerAdditions: OnLeftClickBank, BagOMatic: cmd_showbank, BagOMatic: cmd_sortall, BankArkel: CloseBank |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 14 |
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

Shared function table with 7 member functions; the primary API surface for 5 addons.

## Functions

- BankWindow.GetItem
- BankWindow.GetSlotNumberForButtonIndex
- BankWindow.Hide
- BankWindow.IsShowing
- BankWindow.OpenBank
- BankWindow.Show
- BankWindow.UpdateAllBankSlots

## Observed Members

- none

## Seen In

- AnywhereTrainer
- AnywhereTrainerAdditions
- BagOMatic
- BankArkel
- BankWindowFix

## Examples

- AnywhereTrainer: OnLeftClickBank -> BankWindow.Show()
- AnywhereTrainer: OnLeftClickBank -> BankWindow.Hide()
- AnywhereTrainerAdditions: EquipmentLButtonDown -> BankWindow.GetSlotNumberForButtonIndex(buttonIndex)
- AnywhereTrainerAdditions: EquipmentLButtonDown -> BankWindow.GetItem(slot)
- AnywhereTrainerAdditions: OnLeftClickBank -> BankWindow.OpenBank()
- AnywhereTrainerAdditions: OnLeftClickBank -> BankWindow.Hide()

## Notes

- none
