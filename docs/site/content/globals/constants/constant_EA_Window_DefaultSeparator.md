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
| Addons seen in | AdvancedRenownTrainer, EA_UiDebugTools, EA_UiModWindow, Enemy, Killer, ObjectInspector, PotionBar, WarBoard |
| Files seen in | `/workspace_addons/Enemy/Code/Core/Groups/EffectFilterDialog.xml:167`, `/workspace_addons/Enemy/Code/Core/Groups/EffectFilterDialog.xml:57`, `/workspace_addons/Enemy/Code/Intercom/IntercomDialog.xml:46`, `/workspace_addons/Enemy/Code/UnitFrames/ClickCastingDialog.xml:203`, `/workspace_addons/Enemy/Code/UnitFrames/ClickCastingDialog.xml:305`, `/workspace_addons/Enemy/Code/UnitFrames/ClickCastingDialog.xml:87`, `/workspace_addons/Enemy/Code/UnitFrames/EffectsIndicatorDialog.xml:241`, `/workspace_addons/Enemy/Code/UnitFrames/EffectsIndicatorDialog.xml:376` |
| Namespaces detected | EA_Window_DefaultSeparator |
| Source kinds | xml_attributes |
| Example locations | AdvancedRenownTrainingPresetsWindowFooterSeparator, AdvancedRenownTrainingWindowFooterSeparator1, AdvancedRenownTrainingWindowFooterSeparator2, DebugWindowLogDisplaySeperator, EnemyClickCastingDialogContentScrollChildDivider1, EnemyClickCastingDialogContentScrollChildDivider2 |
| XML usage count | 31 |
| XML attribute usage count | 31 |
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

Observed engine XML template or inherited constant referenced by 10 addons.

## Seen In

- AdvancedRenownTrainer
- EA_UiDebugTools
- EA_UiModWindow
- Enemy
- Killer
- ObjectInspector
- PotionBar
- WarBoard
- WhoHealedMe
- wbLeadHelper

## Used By

- AdvancedRenownTrainingPresetsWindowFooterSeparator
- AdvancedRenownTrainingWindowFooterSeparator1
- AdvancedRenownTrainingWindowFooterSeparator2
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
