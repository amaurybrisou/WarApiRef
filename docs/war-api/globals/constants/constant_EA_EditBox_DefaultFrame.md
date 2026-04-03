# EA_EditBox_DefaultFrame

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
| Addons seen in | AdvancedRenownTrainer, Aura, BuffHead, DAoCBuff, Enemy, Killer, Pocket Palette, Shinies |
| Files seen in | `/workspace/data/raw/Aura/Source/AuraShares.xml:0`, `/workspace/data/raw/BuffHead/Setup/General.xml:0`, `/workspace/data/raw/BuffHead/Setup/SetupAdvancedCompressionItem.xml:0`, `/workspace/data/raw/BuffHead/Setup/SetupAdvancedCompressionItemEffect.xml:0`, `/workspace/data/raw/BuffHead/Setup/SetupAdvancedContainersItemProperties.xml:0`, `/workspace/data/raw/BuffHead/Setup/SetupDisplay.xml:0`, `/workspace/data/raw/BuffHead/Setup/SetupEffectCache.xml:0`, `/workspace/data/raw/BuffHead/Setup/SetupFilter.xml:0` |
| Namespaces detected | EA_EditBox_DefaultFrame |
| Source kinds | xml_attributes |
| Example locations | AdvancedRenownTrainingImportNameInputWindowNameInputBox, AdvancedRenownTrainingImportWindowLinkInputBox, AdvancedRenownTrainingImportWindowNameInputBox, AdvancedRenownTrainingLinkWindowLinkBox, AdvancedRenownTrainingPresetsWindowSaveNameInput, AuraSharesImportExportAuraText |
| XML usage count | 110 |
| XML attribute usage count | 110 |
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
- Aura
- BuffHead
- DAoCBuff
- Enemy
- Killer
- Pocket Palette
- Shinies
- TexturedButtons
- TidyRoll
- TurretRange
- bigger_MacroWindow

## Used By

- AdvancedRenownTrainingImportNameInputWindowNameInputBox
- AdvancedRenownTrainingImportWindowLinkInputBox
- AdvancedRenownTrainingImportWindowNameInputBox
- AdvancedRenownTrainingLinkWindowLinkBox
- AdvancedRenownTrainingPresetsWindowSaveNameInput
- AuraSharesImportExportAuraText
- BuffHeadSetupAdvancedCompressionItemEffectWindowAbilityEditBox
- BuffHeadSetupAdvancedCompressionItemWindowNameEditBox
- BuffHeadSetupAdvancedCompressionItemWindowOverrideIconEditBox
- BuffHeadSetupAdvancedCompressionItemWindowOverrideNameEditBox
- BuffHeadSetupAdvancedContainersItemPropertiesWindowElementContainerSizeColumnsEditBox
- BuffHeadSetupAdvancedContainersItemPropertiesWindowElementContainerSizeRowsEditBox
- BuffHeadSetupAdvancedContainersItemPropertiesWindowElementMaximumThresholdEditBox
- BuffHeadSetupAdvancedContainersItemPropertiesWindowElementMinimumThresholdEditBox
- BuffHeadSetupAdvancedContainersItemPropertiesWindowElementOffsetOffsetXEditBox
- BuffHeadSetupAdvancedContainersItemPropertiesWindowElementOffsetOffsetYEditBox
- BuffHeadSetupAdvancedContainersItemPropertiesWindowElementPaddingContainerPaddingXEditBox
- BuffHeadSetupAdvancedContainersItemPropertiesWindowElementPaddingContainerPaddingYEditBox
- BuffHeadSetupAdvancedContainersItemPropertiesWindowElementPaddingIndicatorPaddingXEditBox
- BuffHeadSetupAdvancedContainersItemPropertiesWindowElementPaddingIndicatorPaddingYEditBox
- BuffHeadSetupAdvancedContainersItemPropertiesWindowElementPositionPlacementFixedOffsetXEditBox
- BuffHeadSetupAdvancedContainersItemPropertiesWindowElementPositionPlacementFixedOffsetYEditBox
- BuffHeadSetupDisplayWindowOffsetXEditBox
- BuffHeadSetupDisplayWindowOffsetYEditBox
- BuffHeadSetupEffectCacheWindowSearchEditBox
- BuffHeadSetupFilterWindowAbilityNameEditBox
- BuffHeadSetupLayoutManagerWindowElementLayoutsCurrentLayoutSaveEditBox
- BuffHeadSetupLayoutPropertiesWindowElementCoreSizeSizeHeightEditBox
- BuffHeadSetupLayoutPropertiesWindowElementCoreSizeSizeWidthEditBox
- BuffHeadSetupLayoutPropertiesWindowElementOffsetOffsetXEditBox
- BuffHeadSetupLayoutPropertiesWindowElementOffsetOffsetYEditBox
- BuffHeadSetupLayoutPropertiesWindowElementSizeScaleEditBox
- BuffHeadSetupLayoutPropertiesWindowElementSizeSizeHeightEditBox
- BuffHeadSetupLayoutPropertiesWindowElementSizeSizeWidthEditBox
- BuffHeadSetupPriorityEffectsItemWindowAbilityIdEditBox
- BuffHeadSetupSelectColorWindowTintBlueEditBox
- BuffHeadSetupSelectColorWindowTintGreenEditBox
- BuffHeadSetupSelectColorWindowTintRedEditBox
- DAoCBuffFrameSettingsTab_ScrollChild_FilterNameEditBox
- DAoCBuffFrameSettingsTab_ScrollChild_FrameNameEditBox
- DAoCBuffListManagerTab_ScrollChild_AddListEditBox
- DAoCBuffListManagerTab_ScrollChild_IDEditBox
- DAoCBuffListManagerTab_ScrollChild_NameEditBox
- DyeWindowFilterEditBox
- EA_Window_MacroDetailsName
- EnemyChooseChannelDialogTellDetailsName
- EnemyClickCastingDialogContentScrollChildActionConfig2Command
- EnemyClickCastingDialogContentScrollChildName
- EnemyConfigurationWindow_PropertyColorTemplateValue1
- EnemyConfigurationWindow_PropertyColorTemplateValue2
- EnemyConfigurationWindow_PropertyColorTemplateValue3
- EnemyConfigurationWindow_PropertyNumberArray2TemplateValue1
- EnemyConfigurationWindow_PropertyNumberArray2TemplateValue2
- EnemyConfigurationWindow_PropertyNumberArray3TemplateValue1
- EnemyConfigurationWindow_PropertyNumberArray3TemplateValue2
- EnemyConfigurationWindow_PropertyNumberArray3TemplateValue3
- EnemyConfigurationWindow_PropertyNumberTemplateValue
- EnemyConfigurationWindow_PropertyStringTemplateValue
- EnemyEffectFilterDialogAbilityIds
- EnemyEffectFilterDialogDescription
- EnemyEffectFilterDialogDurationMax
- EnemyEffectFilterDialogDurationMin
- EnemyEffectFilterDialogFilterName
- EnemyEffectFilterDialogName
- EnemyEffectFilterDialogStackCount
- EnemyEffectsIndicatorDialogContentScrollChildLScale
- EnemyEffectsIndicatorDialogContentScrollChildLogic
- EnemyEffectsIndicatorDialogContentScrollChildName
- EnemyEffectsIndicatorDialogContentScrollChildOffsetX
- EnemyEffectsIndicatorDialogContentScrollChildOffsetY
- EnemyEffectsIndicatorDialogContentScrollChildScale
- EnemyEffectsIndicatorDialogContentScrollChildSizeX
- EnemyEffectsIndicatorDialogContentScrollChildSizeY
- EnemyUnitFramePartDialogContentScrollChildName
- KillerSettingsWindowContentFeedHistoryLimit
- KillerSettingsWindowContentHistoryMaxZones
- KillerSettingsWindowContentMainHeight
- KillerSettingsWindowContentMainWidth
- KillerSettingsWindowContentPersonalDeaths
- KillerSettingsWindowContentPersonalKills
- KillerSettingsWindowContentWindowMinutes
- Shinies_Default_Editbox_ClearLargeBold
- Shinies_Default_Editbox_ClearMedium
- Shinies_Default_Editbox_ClearMediumBold
- Shinies_Default_Editbox_ClearSmallBold
- Shinies_Default_Editbox_DefaultGiant
- Shinies_Default_Editbox_DefaultHuge
- Shinies_Default_Editbox_DefaultLarge
- TRollAutoRollAddByIdIdEditBox
- TRollAutoRollAddByIdNameEditBox
- TexturedButtonsSetupActionbarWindowPaddingXEditBox
- TexturedButtonsSetupActionbarWindowPaddingYEditBox
- TexturedButtonsSetupActionbarWindowSpacingXEditBox
- TexturedButtonsSetupActionbarWindowSpacingYEditBox
- TexturedButtonsSetupSelectColorWindowTintBlueEditBox
- TexturedButtonsSetupSelectColorWindowTintGreenEditBox
- TexturedButtonsSetupSelectColorWindowTintRedEditBox
- TexturedButtonsSetupTintWindowTintBlueEditBox
- TexturedButtonsSetupTintWindowTintGreenEditBox
- TexturedButtonsSetupTintWindowTintRedEditBox
- TurretRangeSetupDisplayWindowDistanceOffsetXEditBox
- TurretRangeSetupDisplayWindowDistanceOffsetYEditBox
- TurretRangeSetupDisplayWindowGraphicLimitEditBox
- TurretRangeSetupDisplayWindowSelectColorTintBlueEditBox
- TurretRangeSetupDisplayWindowSelectColorTintGreenEditBox
- TurretRangeSetupDisplayWindowSelectColorTintRedEditBox
- TurretRangeSetupDistanceWindowColorTintBlueEditBox
- TurretRangeSetupDistanceWindowColorTintGreenEditBox
- TurretRangeSetupDistanceWindowColorTintRedEditBox
- TurretRangeSetupDistanceWindowDistanceEditBox

## Related APIs

- [EditBox](../../xml/element_types/element_EditBox.md) (HIGH 100/100) - XML Element Type

## Used With

- none

## Triggered By

- none

## Affects

- none

## Notes

- none
