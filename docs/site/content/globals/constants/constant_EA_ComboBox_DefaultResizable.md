# EA_ComboBox_DefaultResizable

- Category: Constant
- Confidence level: HIGH
- Confidence score: 100/100

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 170

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, seen in 4 or more addons, used directly in xml handler attributes.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +30 Used directly in XML handler attributes: XML exposure suggests an engine-level contract.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +10 Referenced from initialization flow: Lifecycle reconstruction references this symbol.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +25 Referenced by generated docs or reference files: The symbol is reinforced outside a single call page.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | AdvancedRenownTrainer, AutoBand, BuffHead, Busted, Cheeseboard, Enemy, MapMonster, Pocket Palette |
| Files seen in | `/workspace/Autoband/AutoBandWindowConfig.xml:95`, `/workspace/Autoband/AutoBandWindowTemplate.xml:46`, `/workspace/BuffHead/Setup/SetupAdvancedContainersItem.xml:116`, `/workspace/BuffHead/Setup/SetupAdvancedContainersItem.xml:138`, `/workspace/BuffHead/Setup/SetupAdvancedContainersItemProperties.xml:182`, `/workspace/BuffHead/Setup/SetupAdvancedContainersItemProperties.xml:215`, `/workspace/BuffHead/Setup/SetupAdvancedContainersItemProperties.xml:380`, `/workspace/BuffHead/Setup/SetupAdvancedContainersItemProperties.xml:402` |
| Namespaces detected | EA_ComboBox_DefaultResizable |
| Source kinds | flows, xml_attributes |
| Example locations | AdvancedRenownTrainingPresetsWindowLoadComboBox, AutoBandWindowConfigComboBox, AutoBandWindowTemplateComboBox, BuffHeadSetupAdvancedContainersItemPropertiesWindowElementComboBox, BuffHeadSetupAdvancedContainersItemPropertiesWindowElementEffectsBuffsComboBox, BuffHeadSetupAdvancedContainersItemPropertiesWindowElementEffectsDebuffsComboBox |
| XML usage count | 96 |
| XML attribute usage count | 96 |
| Lua usage count | 0 |
| Global usage count | 0 |
| Local definition count | 0 |
| Documentation references | 1 |
| Initialization flow references | 3 |
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

Observed engine XML template or inherited constant referenced by 14 addons.

## Seen In

- AdvancedRenownTrainer
- AutoBand
- BuffHead
- Busted
- Cheeseboard
- Enemy
- MapMonster
- Pocket Palette
- PotionBar
- TexturedButtons
- ThinkOutLoud
- TidyChat
- TurretRange
- wbLeadHelper

## Used By

- AdvancedRenownTrainingPresetsWindowLoadComboBox
- AutoBandWindowConfigComboBox
- AutoBandWindowTemplateComboBox
- BuffHeadSetupAdvancedContainersItemPropertiesWindowElementComboBox
- BuffHeadSetupAdvancedContainersItemPropertiesWindowElementEffectsBuffsComboBox
- BuffHeadSetupAdvancedContainersItemPropertiesWindowElementEffectsDebuffsComboBox
- BuffHeadSetupAdvancedContainersItemPropertiesWindowElementLayerLayerComboBox
- BuffHeadSetupAdvancedContainersItemPropertiesWindowElementLayoutLayoutComboBox
- BuffHeadSetupAdvancedContainersItemPropertiesWindowElementMaximumThresholdComboBox
- BuffHeadSetupAdvancedContainersItemPropertiesWindowElementMinimumThresholdComboBox
- BuffHeadSetupAdvancedContainersItemPropertiesWindowElementPositionPlacementAnchoredAnchorComboBox
- BuffHeadSetupAdvancedContainersItemPropertiesWindowElementPositionPlacementComboBox
- BuffHeadSetupAdvancedContainersItemPropertiesWindowPropertiesComboBox
- BuffHeadSetupAdvancedContainersItemWindowPositionComboBox
- BuffHeadSetupAdvancedContainersItemWindowTargetComboBox
- BuffHeadSetupContainerWindowContainerAlwaysShowAnchorCombo
- BuffHeadSetupContainerWindowContainerBuffsAnchorCombo
- BuffHeadSetupContainerWindowContainerDebuffsAnchorCombo
- BuffHeadSetupContainerWindowContainerTypeList
- BuffHeadSetupDisplayWindowIndicatorLayer
- BuffHeadSetupDisplayWindowSortByList
- BuffHeadSetupDisplayWindowSortDirectionList
- BuffHeadSetupGeneralWindowCompressList
- BuffHeadSetupLayoutManagerWindowElementExportComboBox
- BuffHeadSetupLayoutManagerWindowElementLayoutsComboBox
- BuffHeadSetupLayoutPropertiesWindowElementComboBox
- BuffHeadSetupLayoutPropertiesWindowElementDurationFormatFormatComboBox
- BuffHeadSetupLayoutPropertiesWindowElementFontAlignmentComboBox
- BuffHeadSetupLayoutPropertiesWindowElementIconBorderColorTypeComboBox
- BuffHeadSetupLayoutPropertiesWindowElementLayerLayerComboBox
- BuffHeadSetupLayoutPropertiesWindowElementStatusBarForegroundColorTypeComboBox
- BuffHeadSetupLayoutPropertiesWindowElementStatusBarOrientationOrientationComboBox
- BuffHeadSetupLayoutPropertiesWindowPropertiesComboBox
- BuffHeadSetupPerformanceWindowEffectAnchoringList
- BuffHeadSetupPriorityEffectsWindowAnimationComboBox
- BuffHeadSetupTrackersWindowTrackerBuffsList
- BuffHeadSetupTrackersWindowTrackerDebuffsList
- BustedGUIAddonSelect
- CheeseboardWindowClassComboBox
- CheeseboardWindowTotalsComboBox
- DyeWindowDyeOrderCombo
- DyeWindowDyeTypeCombo
- EnemyChooseChannelDialogChannelList
- EnemyClickCastingDialogContentScrollChildAction
- EnemyClickCastingDialogContentScrollChildMouseButton
- EnemyClickCastingDialogContentScrollChildPlayerType
- EnemyCombatLogStatsWindowType
- EnemyEffectFilterDialogCastedByMe
- EnemyEffectFilterDialogDurationType
- EnemyEffectsIndicatorDialogContentScrollChildCanDispell
- EnemyEffectsIndicatorDialogContentScrollChildPlayerType
- EnemyIntercomJoinDialogGroupList
- EnemyUnitFramePartDialogContentScrollChildPlayerType
- MapMonster_EditorWindowPinTypeComboBox
- MapMonster_EditorWindowSubTypeComboBox
- MapMonster_EditorWindowZoneNameComboBox
- PotionBarTypeTemplateActivatorCombo
- PotionBarTypeTemplateBuildCombo
- PotionBarTypeTemplateCombo
- PotionBarTypeTemplateInfoTextBRCombo
- PotionBarTypeTemplateInfoTextTRCombo
- TChatTabTextEntryTemplateAnchorPointCombo
- TChatTabTextEntryTemplateRelativeToCombo
- TChatTabWindowsGroupTemplateScrollbarPositionCombo
- TChatTabWindowsTemplateSelectWindowCombo
- TOLSettingsWindowEventComboBox
- TOLSettingsWindowPhraseComboBox
- TOLSettingsWindowPhraseEditWindowPhraseChannelComboBox
- TOLSettingsWindowPhraseEditWindowPhraseSayselfComboBox
- TOLSettingsWindowSkillComboBox
- TOLSettingsWindowSkillEditWindowSkillIsUrgentComboBox
- TOLSettingsWindowSkillEditWindowSkillSpeakChanceComboBox
- TexturedButtonsSetupActionbarWindowBarComboBox
- TexturedButtonsSetupAdvancedTexturesWindowCustomButtonComboBox
- TexturedButtonsSetupAdvancedTexturesWindowCustomTextureComboBox
- TexturedButtonsSetupAdvancedTexturesWindowCustomTextureSliceComboBox
- TexturedButtonsSetupAdvancedTexturesWindowPresetComboBox
- TexturedButtonsSetupAdvancedTexturesWindowSlotTypeComboBox
- TexturedButtonsSetupMiscWindowActionButtonPickUpModifierComboBox
- TexturedButtonsSetupMiscWindowCustomGlowComboBox
- TexturedButtonsSetupTexturesWindowCustomButtonComboBox
- TexturedButtonsSetupTexturesWindowCustomTextureComboBox
- TexturedButtonsSetupTexturesWindowCustomTextureSliceComboBox
- TexturedButtonsSetupTexturesWindowPresetComboBox
- TexturedButtonsSetupTintWindowTintTypeComboBox
- TurretRangeSetupDisplayWindowCircleModeComboBox
- TurretRangeSetupDisplayWindowCircleTypeComboBox
- TurretRangeSetupDisplayWindowDistanceLayoutComboBox
- TurretRangeSetupDisplayWindowDistanceTypeComboBox
- TurretRangeSetupDisplayWindowElementComboBox
- TurretRangeSetupDisplayWindowGraphicTypeComboBox
- WbLeadHelperMessageContentScrollChildLabelColorCombobox
- WbLeadHelperMessageContentScrollChildMessageColorCombobox
- WbLeadHelperMessageContentScrollChildSubmenuCombobox
- WbLeadHelperMessageContentScrollChildTypeCombobox
- wbLeadHelperConfigTabMessageTextColorCombobox

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
