# DialogManager

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
| Addons seen in | AdvancedRenownTrainer, Aura, BankArkel, CM_ClosetGoblin, Enemy, RoR_SoR, Shinies, TidyChat |
| Files seen in | `/workspace/data/raw/Aura/Source/AuraShares.lua:419`, `/workspace/data/raw/BankArkel/BankArkel.lua:85`, `/workspace/data/raw/BankArkel/BankArkel.lua:95`, `/workspace/data/raw/ClosetGoblin/ClosetGoblinCharacterWindow.lua:147`, `/workspace/data/raw/ClosetGoblin/ClosetGoblinCharacterWindow.lua:154`, `/workspace/data/raw/ClosetGoblin/ClosetGoblinCharacterWindow.lua:212`, `/workspace/data/raw/ClosetGoblin/ClosetGoblinCharacterWindow.lua:216`, `/workspace/data/raw/ClosetGoblin/ClosetGoblinCharacterWindow.lua:234` |
| Namespaces detected | DialogManager |
| Source kinds | lua_calls |
| Example locations | AdvancedRenownTrainer: AdvancedRenownTraining.Respecialize, Aura: AuraShares.OnImportExportOkButton, BankArkel: BankArkel.ConvertDB, BankArkel: BankArkel.Init, CM_ClosetGoblin: ClosetGoblinCharacterWindow.OnClickDeleteSetButton, CM_ClosetGoblin: ClosetGoblinCharacterWindow.OnClickNewSetButton |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 41 |
| Global usage count | 3 |
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

Observed shared global table or namespace surfaced in 8 addons.

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
- CM_ClosetGoblin
- Enemy
- RoR_SoR
- Shinies
- TidyChat

## Examples

- AdvancedRenownTrainer: AdvancedRenownTraining.Respecialize -> DialogManager.MakeOneButtonDialog(GetStringFormatFromTable("TrainingStrings",StringTables.Training.TEXT_RESPEC_NOT_ENOUGH_MONEY,{MoneyFrame.FormatMoneyString(respecCost)}), GetString(StringTables.Default.LABEL_OKAY), nil)
- AdvancedRenownTrainer: AdvancedRenownTraining.Respecialize -> DialogManager.MakeTwoButtonDialog(GetStringFormatFromTable("TrainingStrings",StringTables.Training.TEXT_RESPEC_CONFIRMATION,{MoneyFrame.FormatMoneyString(respecCost)}), GetString(StringTables.Default.LABEL_YES), AdvancedRenownTraining.RefundRenownPoints, GetString(StringTables.Default.LABEL_NO))
- Aura: AuraShares.OnImportExportOkButton -> DialogManager.MakeOneButtonDialog(T["An error occurred attempting to load the Aura.  Verify you have entered the correct data."], L "Ok", nil)
- BankArkel: BankArkel.ConvertDB -> DialogManager.MakeOneButtonDialog(StringToWString(convTxt.."v.1 >>> v.2"), StringToWString(errBtn1), nil)
- BankArkel: BankArkel.Init -> DialogManager.MakeOneButtonDialog(StringToWString(errTxt1), StringToWString(errBtn1), BankArkel.ResetDB)
- CM_ClosetGoblin: ClosetGoblinCharacterWindow.OnClickDeleteSetButton -> DialogManager.MakeTwoButtonDialog(cgL["confirm_delete_set"]:gsub(L "#1#",set.name), confirmYes, ClosetGoblinCharacterWindow.OnConfirmDeleteSet, confirmNo, nil)

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
