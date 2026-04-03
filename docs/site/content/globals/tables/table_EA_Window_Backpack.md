# EA_Window_Backpack

- Category: Global Table
- Confidence level: HIGH
- Confidence score: 100/100

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 130

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, seen in 4 or more addons, matches a known engine namespace.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +25 Referenced by generated docs or reference files: The symbol is reinforced outside a single call page.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | BagOMatic, BankArkel, Killer, Shinies |
| Files seen in | `/workspace/data/raw/BankArkel/BankArkel.lua:371`, `/workspace/data/raw/Killer/KillerRenown.lua:108`, `/workspace/data/raw/Killer/KillerRenown.lua:127`, `/workspace/data/raw/Shinies/Modules/UI/Shinies-UI-Browse.lua:298` |
| Namespaces detected | EA_Window_Backpack |
| Source kinds | globals, lua_calls |
| Example locations | BankArkel: BankArkel.UpdateData, Killer: Killer.CaptureInitialWarCrests, Killer: Killer.OnCurrencyUpdated, Shinies: _G.Shinies:NewModule.EA_Window_Backpack_EquipmentLButtonDown |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 4 |
| Global usage count | 1 |
| Local definition count | 4 |
| Documentation references | 1 |
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

- EA_Window_Backpack.GetItemsFromBackpack

## Observed Members

- none

## Seen In

- BagOMatic
- BankArkel
- Killer
- Shinies

## Examples

- BankArkel: BankArkel.UpdateData -> EA_Window_Backpack.GetItemsFromBackpack(Read[i])
- Killer: Killer.CaptureInitialWarCrests -> EA_Window_Backpack.GetItemsFromBackpack(EA_Window_Backpack.TYPE_CURRENCY)
- Killer: Killer.OnCurrencyUpdated -> EA_Window_Backpack.GetItemsFromBackpack(EA_Window_Backpack.TYPE_CURRENCY)
- Shinies: _G.Shinies:NewModule.EA_Window_Backpack_EquipmentLButtonDown -> EA_Window_Backpack.GetItemsFromBackpack(EA_Window_Backpack.currentMode)

## Related APIs

- [DialogManager.MakeOneButtonDialog](../functions/global_DialogManager.MakeOneButtonDialog.md) (HIGH 100/100) - Global Function
- [EA_ChatWindow.Print](../functions/global_EA_ChatWindow.Print.md) (HIGH 100/100) - Global Function
- [LibSlash.IsSlashCmdRegistered](../functions/global_LibSlash.IsSlashCmdRegistered.md) (HIGH 100/100) - Global Function
- [LibSlash.RegisterSlashCmd](../functions/global_LibSlash.RegisterSlashCmd.md) (HIGH 100/100) - Global Function
- [WindowSetLayer](../../window_api/functions/window_WindowSetLayer.md) (HIGH 100/100) - Window Function
- [WindowSetParent](../../window_api/functions/window_WindowSetParent.md) (HIGH 100/100) - Window Function

## Used With

- [CreateWindow](../functions/global_CreateWindow.md) (HIGH 75/100) - Global Function

## Triggered By

- none

## Affects

- none

## Notes

- none
