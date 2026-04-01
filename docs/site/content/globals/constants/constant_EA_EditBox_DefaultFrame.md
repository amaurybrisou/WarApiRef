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
| Addons seen in | AdvancedRenownTrainer, Aura, AutoBand, BuffHead, EA_UiDebugTools, EA_UiModWindow, Enemy, Killer |
| Files seen in | `/workspace/Aura/Source/AuraShares.xml:41`, `/workspace/Autoband/AutoBandWindowTemplate.xml:198`, `/workspace/BuffHead/Setup/General.xml:146`, `/workspace/BuffHead/Setup/General.xml:186`, `/workspace/BuffHead/Setup/General.xml:225`, `/workspace/BuffHead/Setup/SetupAdvancedCompressionItem.xml:105`, `/workspace/BuffHead/Setup/SetupAdvancedCompressionItem.xml:138`, `/workspace/BuffHead/Setup/SetupAdvancedCompressionItem.xml:160` |
| Namespaces detected | EA_EditBox_DefaultFrame |
| Source kinds | xml_attributes |
| Example locations | AdvancedRenownTrainingImportNameInputWindowNameInputBox, AdvancedRenownTrainingImportWindowLinkInputBox, AdvancedRenownTrainingImportWindowNameInputBox, AdvancedRenownTrainingLinkWindowLinkBox, AdvancedRenownTrainingPresetsWindowSaveNameInput, AuraSharesImportExportAuraText |
| XML usage count | 135 |
| XML attribute usage count | 135 |
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

Observed engine XML template or inherited constant referenced by 23 addons.

## Seen In

- AdvancedRenownTrainer
- Aura
- AutoBand
- BuffHead
- EA_UiDebugTools
- EA_UiModWindow
- Enemy
- Killer
- LoyalPet
- MapMonster
- MapPin
- ObjectInspector
- Pocket Palette
- RVAPI_ColorDialog
- RVMOD_Manager
- RandomMount
- Shinies
- TexturedButtons
- ThinkOutLoud
- TidyRoll
- TurretRange
- bigger_MacroWindow
- wbLeadHelper

## Used By

- AdvancedRenownTrainingImportNameInputWindowNameInputBox
- AdvancedRenownTrainingImportWindowLinkInputBox
- AdvancedRenownTrainingImportWindowNameInputBox
- AdvancedRenownTrainingLinkWindowLinkBox
- AdvancedRenownTrainingPresetsWindowSaveNameInput
- AuraSharesImportExportAuraText
- AutoBandWindowTemplateSaveNameEditBox
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
- DebugWindowTextBox
- DevPadNewWindowTextBox
- DevPadRenameWindowTextBox
- DevPadSaveWindowTextBox
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
- LPETOptionsProfilesEdit
- MapMonsterEditorWindowEditBoxCoord
- MapMonsterEditorWindowEditBoxDefault
- MapMonsterPinTypeEditorWindowEditBoxDefault
- MapMonster_PinTypeEditorWindowRadiusEditBox
- MapPin_SetupIDBox
- MapPin_SetupTimersHour
- MapPin_SetupTimersMin
- MapPin_SetupTimersSec
- MapPin_SetupTitleBox
- MapPin_SetupXBox
- MapPin_SetupYBox
- ObjectInspectorCommandEditBox
- ObjectInspectorDepthEditBox
- RVAPI_ColorDialogEditBoxTemplate
- RVMOD_ManagerModInfoTemplateScrollChildProjectURLText
- RandomMountWindowSettingsContentMinLevelEdit
- Shinies_Default_Editbox_ClearLargeBold
- Shinies_Default_Editbox_ClearMedium
- Shinies_Default_Editbox_ClearMediumBold
- Shinies_Default_Editbox_ClearSmallBold
- Shinies_Default_Editbox_DefaultGiant
- Shinies_Default_Editbox_DefaultHuge
- Shinies_Default_Editbox_DefaultLarge
- TOLSettingsWindowPhraseEditWindowPhraseTextEdit
- TOLSettingsWindowSkillEditWindowSkillNameEdit
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
- UiModAdvancedWindowAddOnsDirectory
- UiModAdvancedWindowCustomUiDirectory
- WbLeadHelperMessageContentScrollChildLabelText
- WbLeadHelperMessageContentScrollChildMessageText
- wbLeadHelperConfigTabMessageEndText
- wbLeadHelperConfigTabMessageStartText

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
