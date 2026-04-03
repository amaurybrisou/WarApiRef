# EA_Button_DefaultResizeable

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
| Addons seen in | AdvancedRenownTrainer, Aura, BuffHead, CM_ClosetGoblin, DAoCBuff, Enemy, Pocket Palette, Shinies |
| Files seen in | `/workspace/data/raw/Aura/Source/Templates.xml:0`, `/workspace/data/raw/BuffHead/Setup/SetupAdvancedCompression.xml:0`, `/workspace/data/raw/BuffHead/Setup/SetupAdvancedCompressionItem.xml:0`, `/workspace/data/raw/BuffHead/Setup/SetupAdvancedCompressionItemEffect.xml:0`, `/workspace/data/raw/BuffHead/Setup/SetupAdvancedContainers.xml:0`, `/workspace/data/raw/BuffHead/Setup/SetupAdvancedContainersItem.xml:0`, `/workspace/data/raw/BuffHead/Setup/SetupAdvancedContainersItemProperties.xml:0`, `/workspace/data/raw/BuffHead/Setup/SetupEffectCache.xml:0` |
| Namespaces detected | EA_Button_DefaultResizeable |
| Source kinds | xml_attributes |
| Example locations | AdvancedRenownTrainingExportWindowCancelButton, AdvancedRenownTrainingExportWindowHyperLinkButton, AdvancedRenownTrainingExportWindowWardrobeButton, AdvancedRenownTrainingImportNameInputWindowCancelButton, AdvancedRenownTrainingImportNameInputWindowOkButton, AdvancedRenownTrainingImportWindowCancelButton |
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

Observed engine XML template or inherited constant referenced by 16 addons.

## Seen In

- AdvancedRenownTrainer
- Aura
- BuffHead
- CM_ClosetGoblin
- DAoCBuff
- Enemy
- Pocket Palette
- Shinies
- TidyChat
- TidyRoll
- TurretRange
- WSCT
- WarBoard
- WhoHealedMe
- bigger_MacroWindow
- followTheLeader

## Used By

- AdvancedRenownTrainingExportWindowCancelButton
- AdvancedRenownTrainingExportWindowHyperLinkButton
- AdvancedRenownTrainingExportWindowWardrobeButton
- AdvancedRenownTrainingImportNameInputWindowCancelButton
- AdvancedRenownTrainingImportNameInputWindowOkButton
- AdvancedRenownTrainingImportWindowCancelButton
- AdvancedRenownTrainingImportWindowOkButton
- AdvancedRenownTrainingLinkWindowCloseButton
- AdvancedRenownTrainingPresetsWindowDeleteButton
- AdvancedRenownTrainingPresetsWindowExportButton
- AdvancedRenownTrainingPresetsWindowLoadButton
- AdvancedRenownTrainingPresetsWindowSaveButton
- AdvancedRenownTrainingPresetsWindowWardrobeButton
- AdvancedRenownTrainingWindowCancelButton
- AdvancedRenownTrainingWindowPresetButton
- AdvancedRenownTrainingWindowPurchaseButton
- AdvancedRenownTrainingWindowRespecializeButton
- ApplyRowAlignButton
- AuraWindowButton
- BuffHeadSetupAdvancedCompressionItemEffectWindowApplyButton
- BuffHeadSetupAdvancedCompressionItemEffectWindowCreateButton
- BuffHeadSetupAdvancedCompressionItemWindowAddButton
- BuffHeadSetupAdvancedCompressionItemWindowApplyButton
- BuffHeadSetupAdvancedCompressionItemWindowCreateButton
- BuffHeadSetupAdvancedCompressionWindowNewButton
- BuffHeadSetupAdvancedContainersItemPropertiesWindowElementEffectsAlwaysIgnoreButton
- BuffHeadSetupAdvancedContainersItemPropertiesWindowElementEffectsAlwaysShowButton
- BuffHeadSetupAdvancedContainersItemWindowApplyButton
- BuffHeadSetupAdvancedContainersItemWindowCreateButton
- BuffHeadSetupAdvancedContainersWindowNewButton
- BuffHeadSetupEffectCacheWindowRefreshButton
- BuffHeadSetupEffectCacheWindowResetButton
- BuffHeadSetupFilterWindowAddButton
- BuffHeadSetupGeneralWindowAlwaysIgnoreButton
- BuffHeadSetupGeneralWindowAlwaysShowButton
- BuffHeadSetupLayoutManagerWindowElementImportButton
- BuffHeadSetupLayoutManagerWindowElementLayoutsCurrentLayoutSaveButton
- BuffHeadSetupLayoutManagerWindowElementLayoutsSavedLayoutActivateButton
- BuffHeadSetupLayoutManagerWindowElementLayoutsSavedLayoutDeleteButton
- BuffHeadSetupLayoutPropertiesWindowElementCoreSizeAutoSizeButton
- BuffHeadSetupLayoutWindowApplyButton
- BuffHeadSetupLayoutWindowManagerButton
- BuffHeadSetupPriorityEffectsItemWindowApplyButton
- BuffHeadSetupPriorityEffectsItemWindowCreateButton
- BuffHeadSetupPriorityEffectsWindowNewButton
- BuffHeadSetupTrackersWindowAlwaysIgnoreButton
- BuffHeadSetupTrackersWindowAlwaysShowButton
- ClosetGoblinCharacterWindowContentsDeleteSet
- ClosetGoblinCharacterWindowContentsNewSet
- ClosetGoblinCharacterWindowContentsZoneConfig
- ClosetGoblinZoneWindowContentsChangeZoneButton
- DAoCBuffFrameSettingsTab_ScrollChild_AddFilterButton
- DAoCBuffFrameSettingsTab_ScrollChild_FilterEditButton
- DAoCBuffFrameSettingsTab_ScrollChild_FilterNameEditBoxButton
- DAoCBuffFrameSettingsTab_ScrollChild_FrameNameEditBoxButton
- DAoCBuffFrameSettingsTab_ScrollChild_RemoveFilterButton
- DAoCBuffFrameSettingsTab_ScrollChild_RemoveFrame
- DAoCBuffFrameSettings_G4Filter_G4RecursiveConditionsButton
- DAoCBuffGeneralSettingsTab_ScrollChild_ChangeGlobalBorderButton
- DAoCBuffGeneralSettingsTab_ScrollChild_ChangeGlobalFontButton
- DAoCBuffGeneralSettingsTab_ScrollChild_ChangeGlobalGlassButton
- DAoCBuffGeneralSettingsTab_ScrollChild_ChangeGlobalRefreshButton
- DAoCBuffGeneralSettingsTab_ScrollChild_ChangeGlobalSizeButton
- DAoCBuffGeneralSettingsTab_ScrollChild_ImExportButton
- DAoCBuffListManagerTab_ScrollChild_Add
- DAoCBuffListManagerTab_ScrollChild_AddListEditBoxButton
- DAoCBuffListManagerTab_ScrollChild_ClearLeft
- DAoCBuffListManagerTab_ScrollChild_ClearRight
- DAoCBuffListManagerTab_ScrollChild_CopyLeftToRight
- DAoCBuffListManagerTab_ScrollChild_CopyRightToLeft
- DAoCBuffListManagerTab_ScrollChild_MoveLeftToRight
- DAoCBuffListManagerTab_ScrollChild_MoveRightToLeft
- DAoCBuffListManagerTab_ScrollChild_RemoveLeft
- DAoCBuffListManagerTab_ScrollChild_RemoveListEditBoxButton
- DAoCBuffListManagerTab_ScrollChild_RemoveRight
- DAoCBuffMessageWindowScrollWindowScrollChildOpenOptions
- DAoCBuff_SettingsAddFrameButton
- DAoCBuff_SettingsReloadTrackers
- DAoCBuff_Settings_FilterFrame_ScrollChild_AddClassTableButton
- DAoCBuff_Settings_FilterFrame_ScrollChild_AddConditionButton
- DAoCBuff_Settings_FilterFrame_ScrollChild_RemoveConditionButton
- EA_Window_MacroDetailsSave
- EnemyChooseChannelDialogOkButton
- EnemyClickCastingDialogCancelButton
- EnemyClickCastingDialogOkButton
- EnemyConfigDialogCancelButton
- EnemyConfigDialogOkButton
- EnemyConfigurationWindow_ButtonTemplateValue
- EnemyEffectFilterDialogCancelButton
- EnemyEffectFilterDialogOkButton
- EnemyEffectsIndicatorDialogCancelButton
- EnemyEffectsIndicatorDialogOkButton
- EnemyIntercomDialogCancelButton
- EnemyIntercomDialogCreateButton
- EnemyIntercomDialogCreateButton2
- EnemyIntercomDialogCreateButton3
- EnemyIntercomDialogCreateButton4
- EnemyIntercomDialogCreateButton5
- EnemyIntercomDialogCreateButton6
- EnemyIntercomDialogCreateButton7
- EnemyIntercomDialogInviteButton
- EnemyIntercomJoinDialogOkButton
- EnemyMarkConfigDialogCancelButton
- EnemyMarkConfigDialogOkButton
- EnemyScenarioInfoDialogLeaveButton
- EnemyTextEntryDialogCancelButton
- EnemyTextEntryDialogOkButton
- EnemyUnitFramePartDialogCancelButton
- EnemyUnitFramePartDialogOkButton
- OpenUiModWindowButton
- PPMainCharacterWindowBtn
- PPMainTogglePickerBtn
- Shinies_Default_Button_ClearLargeBoldFont
- Shinies_Default_Button_ClearMediumBoldFont
- Shinies_Default_Button_ClearMediumFont
- Shinies_Default_Button_ClearSmallFont
- TChatButton
- TRollButton
- TidyChatCopyNext
- TidyChatCopyPrev
- ToggleLayoutButton
- TurretRangeSetupDistanceWindowApplyButton
- TurretRangeSetupDistanceWindowCreateButton
- TurretRangeSetupDistancesWindowAddButton
- TurretRangeSetupMenuWindowDisplaySetupButton
- TurretRangeSetupMenuWindowDistancesSetupButton
- TurretRangeSetupMenuWindowGeneralSetupButton
- WBTogglerManagerButton
- WSCTButtonTemplate
- WSCTEventColor
- WSCTOptionsColorPickerWindowAcceptButton
- WhoHealedMeOptionsContentCloseButton
- WhoHealedMeOptionsContentResetButton
- WhoHealedMeWindowOptionsButton
- followTheLeaderWindow

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
