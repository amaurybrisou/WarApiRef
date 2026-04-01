# DialogManager

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
| Addons seen in | AdvancedRenownTrainer, Aura, BankArkel, BankWindowFix, CM_ClosetGoblin, Effigy, Enemy, MapMonster |
| Files seen in | `/workspace/Aura/Source/AuraShares.lua:419`, `/workspace/BankArkel/BankArkel.lua:85`, `/workspace/BankArkel/BankArkel.lua:95`, `/workspace/BankWindowFix/Source/BankWindowFix.lua:108`, `/workspace/ClosetGoblin/ClosetGoblinCharacterWindow.lua:147`, `/workspace/ClosetGoblin/ClosetGoblinCharacterWindow.lua:154`, `/workspace/ClosetGoblin/ClosetGoblinCharacterWindow.lua:212`, `/workspace/ClosetGoblin/ClosetGoblinCharacterWindow.lua:216` |
| Namespaces detected | DialogManager |
| Source kinds | globals, lua_calls |
| Example locations | AdvancedRenownTrainer: AdvancedRenownTraining.Respecialize, Aura: AuraShares.OnImportExportOkButton, BankArkel: BankArkel.ConvertDB, BankArkel: BankArkel.Init, BankWindowFix: BankWindowFix.BagEquipmentRButtonUp, CM_ClosetGoblin: ClosetGoblinCharacterWindow.OnClickDeleteSetButton |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 63 |
| Global usage count | 3 |
| Local definition count | 1 |
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

Observed shared global table or namespace surfaced in 14 addons.

## Functions

- DialogManager.MakeOneButtonDialog
- DialogManager.MakeTextEntryDialog
- DialogManager.MakeTwoButtonDialog

## Observed Members

- none

## Seen In

- AdvancedRenownTrainer
- Aura
- BankArkel
- BankWindowFix
- CM_ClosetGoblin
- Effigy
- Enemy
- MapMonster
- MapPin
- Miracle Grow Remix
- RoR_SoR
- Shinies
- TidyChat
- wbLeadHelper

## Examples

- AdvancedRenownTrainer: AdvancedRenownTraining.Respecialize -> DialogManager.MakeOneButtonDialog(GetStringFormatFromTable("TrainingStrings",StringTables.Training.TEXT_RESPEC_NOT_ENOUGH_MONEY,{MoneyFrame.FormatMoneyString(respecCost)}), GetString(StringTables.Default.LABEL_OKAY), nil)
- AdvancedRenownTrainer: AdvancedRenownTraining.Respecialize -> DialogManager.MakeTwoButtonDialog(GetStringFormatFromTable("TrainingStrings",StringTables.Training.TEXT_RESPEC_CONFIRMATION,{MoneyFrame.FormatMoneyString(respecCost)}), GetString(StringTables.Default.LABEL_YES), AdvancedRenownTraining.RefundRenownPoints, GetString(StringTables.Default.LABEL_NO))
- Aura: AuraShares.OnImportExportOkButton -> DialogManager.MakeOneButtonDialog(T["An error occurred attempting to load the Aura.  Verify you have entered the correct data."], L "Ok", nil)
- BankArkel: BankArkel.ConvertDB -> DialogManager.MakeOneButtonDialog(StringToWString(convTxt.."v.1 >>> v.2"), StringToWString(errBtn1), nil)
- BankArkel: BankArkel.Init -> DialogManager.MakeOneButtonDialog(StringToWString(errTxt1), StringToWString(errBtn1), BankArkel.ResetDB)
- BankWindowFix: BankWindowFix.BagEquipmentRButtonUp -> DialogManager.MakeTwoButtonDialog(text, GetString(StringTables.Default.LABEL_YES), function()DestroyItem(cursorType,slot)end, GetString(StringTables.Default.LABEL_NO), nil, nil, nil, nil, nil, nil, DIALOGID_DESTROY_ITEM)

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
