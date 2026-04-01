# EA_Window_DefaultSeparator

- Category: Constant
- Confidence level: HIGH
- Confidence score: 100/100

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 135

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, seen in 4 or more addons, used directly in xml handler attributes.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +30 Used directly in XML handler attributes: XML exposure suggests an engine-level contract.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | AdvancedRenownTrainer, AutoBand, EA_UiDebugTools, EA_UiModWindow, Enemy, Killer, ObjectInspector, PotionBar |
| Files seen in | `/workspace/Autoband/AutoBandWindowConfig.xml:107`, `/workspace/Autoband/AutoBandWindowConfig.xml:233`, `/workspace/Autoband/AutoBandWindowConfig.xml:73`, `/workspace/Autoband/AutoBandWindowTemplate.xml:81`, `/workspace/Enemy/Code/Core/Groups/EffectFilterDialog.xml:167`, `/workspace/Enemy/Code/Core/Groups/EffectFilterDialog.xml:57`, `/workspace/Enemy/Code/Intercom/IntercomDialog.xml:46`, `/workspace/Enemy/Code/UnitFrames/ClickCastingDialog.xml:203` |
| Namespaces detected | EA_Window_DefaultSeparator |
| Source kinds | xml_attributes |
| Example locations | AdvancedRenownTrainingPresetsWindowFooterSeparator, AdvancedRenownTrainingWindowFooterSeparator1, AdvancedRenownTrainingWindowFooterSeparator2, AutoBandWindowConfigDisplaySeperator1, AutoBandWindowConfigDisplaySeperator2, AutoBandWindowConfigDisplaySeperator3 |
| XML usage count | 36 |
| XML attribute usage count | 36 |
| Lua usage count | 0 |
| Global usage count | 0 |
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

Observed engine XML template or inherited constant referenced by 12 addons.

## Seen In

- AdvancedRenownTrainer
- AutoBand
- EA_UiDebugTools
- EA_UiModWindow
- Enemy
- Killer
- ObjectInspector
- PotionBar
- Queue Queuer
- WarBoard
- WhoHealedMe
- wbLeadHelper

## Used By

- AdvancedRenownTrainingPresetsWindowFooterSeparator
- AdvancedRenownTrainingWindowFooterSeparator1
- AdvancedRenownTrainingWindowFooterSeparator2
- AutoBandWindowConfigDisplaySeperator1
- AutoBandWindowConfigDisplaySeperator2
- AutoBandWindowConfigDisplaySeperator3
- AutoBandWindowTemplateDisplaySeperator
- DebugWindowLogDisplaySeperator
- EnemyClickCastingDialogContentScrollChildDivider1
- EnemyClickCastingDialogContentScrollChildDivider2
- EnemyClickCastingDialogContentScrollChildDivider3
- EnemyEffectFilterDialogDivider1
- EnemyEffectFilterDialogDivider2
- EnemyEffectsIndicatorDialogContentScrollChildDivider1
- EnemyEffectsIndicatorDialogContentScrollChildDivider2
- EnemyEffectsIndicatorDialogContentScrollChildDivider3
- EnemyIntercomDialogDivider1
- EnemyUnitFramePartDialogContentScrollChildDivider1
- EnemyUnitFramePartDialogContentScrollChildDivider2
- KillerSettingsWindowContentSeparator
- KillerSettingsWindowContentStorageSeparator
- ObjectInspectorHorizDivide
- ObjectInspectorHorizDivide2
- PotionBarTypeTemplateTitleSeperator
- QueueQueuer_GUI_DisplaySeperator
- UiModVersionMismatchWindowModListSeperator
- UiModWindowListBottomSeperator
- UiModWindowTopSeperator
- WarBoardOptionsTabMainSocketHorizontalDivider
- WarBoardOptionsTabMainSocketHorizontalDivider2
- WarBoardOptionsTabMainSocketHorizontalDivider3
- WarBoardOptionsTabMainSocketHorizontalDividerTop
- WbLeadHelperMessageContentScrollChildDivider1
- WbLeadHelperMessageContentScrollChildDivider2
- WhoHealedMeOptionsContentSeparator1
- wbLeadHelperConfigTabDisplaySeperator1

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
