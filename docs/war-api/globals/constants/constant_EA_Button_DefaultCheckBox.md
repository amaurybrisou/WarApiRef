# EA_Button_DefaultCheckBox

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
| Addons seen in | AdvancedRenownTrainer, Aura, AutoBand, CM_ClosetGoblin, CMap, Cheeseboard, EA_UiModWindow, Enemy |
| Files seen in | `/workspace/Aura/Source/Templates.xml:362`, `/workspace/Aura/Source/Templates.xml:393`, `/workspace/Autoband/AutoBandWindowConfig.xml:138`, `/workspace/Autoband/AutoBandWindowConfig.xml:163`, `/workspace/Autoband/AutoBandWindowConfig.xml:320`, `/workspace/Autoband/AutoBandWindowTemplate.xml:70`, `/workspace/Autoband/AutoBandWindowTools.xml:50`, `/workspace/Autoband/AutoBandWindowTools.xml:74` |
| Namespaces detected | EA_Button_DefaultCheckBox |
| Source kinds | xml_attributes |
| Example locations | AdvancedRenownTrainingPresetsWindowSaveSelectedCheckBox, Aura_LabelCheckButtonButton, Aura_LargeLabelCheckButtonButton, AutoBandWindowConfigKickCheckBox, AutoBandWindowConfigKickToofarCheckBox, AutoBandWindowConfigMapIconCheckBox |
| XML usage count | 76 |
| XML attribute usage count | 76 |
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

Observed engine XML template or inherited constant referenced by 20 addons.

## Seen In

- AdvancedRenownTrainer
- Aura
- AutoBand
- CM_ClosetGoblin
- CMap
- Cheeseboard
- EA_UiModWindow
- Enemy
- Killer
- MapMonster
- MapPin
- MegaphonePlusPlus
- Miracle Grow Remix
- Pocket Palette
- PotionBar
- Shinies
- TidyChat
- TidyQueue
- TidyRoll
- wbLeadHelper

## Used By

- AdvancedRenownTrainingPresetsWindowSaveSelectedCheckBox
- Aura_LabelCheckButtonButton
- Aura_LargeLabelCheckButtonButton
- AutoBandWindowConfigKickCheckBox
- AutoBandWindowConfigKickToofarCheckBox
- AutoBandWindowConfigMapIconCheckBox
- AutoBandWindowTemplateDefaultCheckBox
- AutoBandWindowToolsKickOfflineCheckBox
- AutoBandWindowToolsKickRankCheckBox
- AutoBandWindowToolsKickZoneCheckBox
- CheeseboardWindowDefaultCheck
- ClosetGoblinCharacterWindowContentsActionBarSettingsCheckboxAB1
- ClosetGoblinCharacterWindowContentsActionBarSettingsCheckboxAB2
- ClosetGoblinCharacterWindowContentsActionBarSettingsCheckboxAB3
- ClosetGoblinCharacterWindowContentsActionBarSettingsCheckboxAB4
- ClosetGoblinCharacterWindowContentsActionBarSettingsCheckboxAB5
- ClosetGoblinCharacterWindowContentsCheckboxUseItemRack
- ClosetGoblinCharacterWindowContentsEquipmentShowCloak
- ClosetGoblinCharacterWindowContentsEquipmentShowCloakHeraldry
- ClosetGoblinCharacterWindowContentsEquipmentShowHelm
- Cmap_MapPinFilterEntryButton1
- Cmap_MapPinFilterEntryButton2
- Cmap_MapPinGutterEntryButton1
- Cmap_MapPinGutterEntryButton2
- EA_Button_UiModuleListRowTemplateEnabled
- EnemyAssistConfigurationMarkNewTarget
- EnemyAssistConfigurationNewTargetSound
- EnemyAssistConfigurationTargetOnNotifyClick
- EnemyClickCastingDialogContentScrollChildArchetype1
- EnemyClickCastingDialogContentScrollChildArchetype2
- EnemyClickCastingDialogContentScrollChildArchetype3
- EnemyClickCastingDialogContentScrollChildExceptMe
- EnemyClickCastingDialogContentScrollChildKeyModifier1
- EnemyClickCastingDialogContentScrollChildKeyModifier2
- EnemyClickCastingDialogContentScrollChildKeyModifier3
- EnemyConfigurationWindow_PropertyBoolTemplateValue
- EnemyEffectsIndicatorDialogContentScrollChildArchetype1
- EnemyEffectsIndicatorDialogContentScrollChildArchetype2
- EnemyEffectsIndicatorDialogContentScrollChildArchetype3
- EnemyEffectsIndicatorDialogContentScrollChildCircleIcon
- EnemyEffectsIndicatorDialogContentScrollChildExceptMe
- EnemyEffectsIndicatorDialogContentScrollChildLScaleCheckBox
- EnemyScenarioAlerterConfigurationEnabled
- EnemyUnitFramePartDialogContentScrollChildArchetype1
- EnemyUnitFramePartDialogContentScrollChildArchetype2
- EnemyUnitFramePartDialogContentScrollChildArchetype3
- EnemyUnitFramePartDialogContentScrollChildExceptMe
- EnemyUnitFramesConfigurationContentScrollChildEnabled
- EnemyUnitFramesConfigurationContentScrollChildSortingEnabled
- KillerSettingsWindowContentPersonalEnabled
- MapFilterContextMenuChoiceCheckBox
- MapMonster_EditorWindowIsLockedCheckBox
- MapMonster_EditorWindowIsPrivateCheckBox
- MapMonster_PinTypeEditorWindowIsLockedCheckBox
- MapMonster_PinTypeEditorWindowIsPrivateCheckBox
- MapPin_SetupBGBox
- MegaphoneMainHighlightLeaderCheckbox
- MegaphoneMainShowLeaderCheckbox
- MiracleGrow2LayoutSettingsTemplateStage0
- MiracleGrow2LayoutSettingsTemplateStage1
- MiracleGrow2LayoutSettingsTemplateStage2
- MiracleGrow2LayoutSettingsTemplateStage3
- MiracleGrow2LayoutSettingsTemplateStage4
- PPMainSaveSettingsButton
- PotionBarTypeTemplateAutohideCheck
- PotionBarTypeTemplateCheck
- PotionBarTypeTemplateDontSplitNameCheck
- PotionBarTypeTemplateShow1Check
- PotionBarTypeTemplateShowLastCheck
- ShiniesConfigPrice_CheckBox
- ShiniesConfigPrice_PriorityRowEnable
- Shinies_CheckButtonWithLabelButton
- TChatCheckboxTemplateButton
- TQueueCheckboxTemplateButton
- TRollCheckboxTemplateButton
- wbLeadHelperConfigTabLfgIconsCheckBox

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
