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
| Addons seen in | ActionFraction, AdvancedRenownTrainer, Aura, BankArkel, BankWindowFix, CM_ClosetGoblin, Crusher, Enemy |
| Files seen in | Code/CombatLog/CombatLog.lua, Code/CombatLog/CombatLogStatsWindow.lua, Code/Core/Groups/EnemyEffectFilter.lua, Code/Core/Main.lua, Code/Marks/Marks.lua, Code/UnitFrames/ClickCasting.lua, Code/UnitFrames/EffectsIndicator.lua, Code/UnitFrames/UnitFramePart.lua |
| Namespaces detected | DialogManager |
| Source kinds | lua_calls |
| Example locations | ActionFraction: SetLocationActionPointBar, AdvancedRenownTrainer: Respecialize, Aura: OnImportExportOkButton, BankArkel: ConvertDB, BankArkel: Init, BankWindowFix: BagEquipmentRButtonUp |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 90 |
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

Shared function table with 4 member functions; the primary API surface for 25 addons.

## Functions

- DialogManager.MakeOneButtonDialog
- DialogManager.MakeTextEntryDialog
- DialogManager.MakeThreeButtonDialog
- DialogManager.MakeTwoButtonDialog

## Observed Members

- none

## Seen In

- ActionFraction
- AdvancedRenownTrainer
- Aura
- BankArkel
- BankWindowFix
- CM_ClosetGoblin
- Crusher
- Enemy
- Hopper
- Lib RuString
- MapMonster
- MapPin
- Minmap
- Miracle Grow Remix
- Motion
- Pure
- RoR_SoR
- Shinies
- Squared
- TastyButtons
- TidyChat
- Vectors
- Wikki's Cooldown Bar
- Wikki's Cooldown Pulse
- wbLeadHelper

## Examples

- ActionFraction: SetLocationActionPointBar -> DialogManager.MakeOneButtonDialog(LOC_TEXT.PLAYERWINDOW_UNAVAILABLE, GetPregameString(StringTables.Pregame.LABEL_CONTINUE))
- AdvancedRenownTrainer: Respecialize -> DialogManager.MakeOneButtonDialog(GetStringFormatFromTable("TrainingStrings",StringTables.Training.TEXT_RESPEC_NOT_ENOUGH_MONEY,{MoneyFrame.FormatMoneyString(respecCost)}), GetString(StringTables.Default.LABEL_OKAY), nil)
- AdvancedRenownTrainer: Respecialize -> DialogManager.MakeTwoButtonDialog(GetStringFormatFromTable("TrainingStrings",StringTables.Training.TEXT_RESPEC_CONFIRMATION,{MoneyFrame.FormatMoneyString(respecCost)}), GetString(StringTables.Default.LABEL_YES), AdvancedRenownTraining.RefundRenownPoints, GetString(StringTables.Default.LABEL_NO))
- Aura: OnImportExportOkButton -> DialogManager.MakeOneButtonDialog(T["An error occurred attempting to load the Aura.  Verify you have entered the correct data."], L "Ok", nil)
- BankArkel: ConvertDB -> DialogManager.MakeOneButtonDialog(StringToWString(convTxt.."v.1 >>> v.2"), StringToWString(errBtn1), nil)
- BankArkel: Init -> DialogManager.MakeOneButtonDialog(StringToWString(errTxt1), StringToWString(errBtn1), BankArkel.ResetDB)

## Notes

- none
