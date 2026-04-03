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
| Addons seen in | AdvancedRenownTrainer, BuffHead, DAoCBuff, Enemy, Pocket Palette, PotionBar, TexturedButtons, TidyChat |
| Files seen in | `/workspace/data/raw/BuffHead/Setup/SetupAdvancedContainersItem.xml:0`, `/workspace/data/raw/BuffHead/Setup/SetupAdvancedContainersItemProperties.xml:0`, `/workspace/data/raw/BuffHead/Setup/SetupContainer.xml:0`, `/workspace/data/raw/BuffHead/Setup/SetupDisplay.xml:0`, `/workspace/data/raw/BuffHead/Setup/SetupGeneral.xml:0`, `/workspace/data/raw/BuffHead/Setup/SetupLayoutManager.xml:0`, `/workspace/data/raw/BuffHead/Setup/SetupLayoutProperties.xml:0`, `/workspace/data/raw/BuffHead/Setup/SetupPerformance.xml:0` |
| Namespaces detected | EA_ComboBox_DefaultResizable |
| Source kinds | flows, xml_attributes |
| Example locations | AdvancedRenownTrainingPresetsWindowLoadComboBox, BuffHeadSetupAdvancedContainersItemPropertiesWindowElementComboBox, BuffHeadSetupAdvancedContainersItemPropertiesWindowElementEffectsBuffsComboBox, BuffHeadSetupAdvancedContainersItemPropertiesWindowElementEffectsDebuffsComboBox, BuffHeadSetupAdvancedContainersItemPropertiesWindowElementLayerLayerComboBox, BuffHeadSetupAdvancedContainersItemPropertiesWindowElementLayoutLayoutComboBox |
| XML usage count | 94 |
| XML attribute usage count | 94 |
| Lua usage count | 0 |
| Global usage count | 0 |
| Local definition count | 0 |
| Documentation references | 1 |
| Initialization flow references | 1 |
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

Observed engine XML template or inherited constant referenced by 9 addons.

## Seen In

- AdvancedRenownTrainer
- BuffHead
- DAoCBuff
- Enemy
- Pocket Palette
- PotionBar
- TexturedButtons
- TidyChat
- TurretRange

## Used By

- AdvancedRenownTrainingPresetsWindowLoadComboBox
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
- DAoCBuffFrameSettings_G1Filter_G1ComboBox
- DAoCBuffFrameSettings_G2Filter_G2ComboBox
- DAoCBuffFrameSettings_G5Filter_G5ComboBox
- DAoCBuffFrameSettings_G5Filter_G5DurationComboBox
- DAoCBuffGeneralSettingsTab_ScrollChild_ChangeGlobalBorderComboBox
- DAoCBuffGeneralSettingsTab_ScrollChild_ChangeGlobalFontComboBox
- DAoCBuffGeneralSettingsTab_ScrollChild_ChangeGlobalGlassComboBox
- DAoCBuffGeneralSettingsTab_ScrollChild_ChangeGlobalRefreshComboBox
- DAoCBuffGeneralSettingsTab_ScrollChild_ChangeGlobalSizeComboBox
- DAoCBuffListManagerTab_ScrollChild_LeftListComboBox
- DAoCBuffListManagerTab_ScrollChild_ManagerModeComboBox
- DAoCBuffListManagerTab_ScrollChild_RemoveListComboBox
- DAoCBuffListManagerTab_ScrollChild_RightListComboBox
- DAoCBuff_Settings_FilterFrame_ScrollChild_ClassTableComboBox
- DAoCBuff_Settings_FilterFrame_ScrollChild_ConditionComboBox
- DAoCBuff_Settings_FilterFrame_ScrollChild_ConditionTypeComboBox
- DAoCBuff_Settings_FilterFrame_ScrollChild_G4HistoryBrowserComboBox
- DAoCBuff_Settings_FilterFrame_ScrollChild_TextureComboBox
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
- PotionBarTypeTemplateActivatorCombo
- PotionBarTypeTemplateBuildCombo
- PotionBarTypeTemplateCombo
- PotionBarTypeTemplateInfoTextBRCombo
- PotionBarTypeTemplateInfoTextTRCombo
- TChatTabTextEntryTemplateAnchorPointCombo
- TChatTabTextEntryTemplateRelativeToCombo
- TChatTabWindowsGroupTemplateScrollbarPositionCombo
- TChatTabWindowsTemplateSelectWindowCombo
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

## Related APIs

- [ComboBox](../../xml/element_types/element_ComboBox.md) (HIGH 100/100) - XML Element Type

## Used With

- none

## Triggered By

- none

## Affects

- none

## Notes

- none
