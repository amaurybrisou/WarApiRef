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
| Addons seen in | AdvancedRenownTrainer, Aura, BuffHead, CM_ClosetGoblin, DAoCBuff, DaemonAssist, EA_UiModWindow, Enemy |
| Files seen in | `/workspace_addons/Aura/Source/Templates.xml:18`, `/workspace_addons/BuffHead/Setup/SetupAdvancedCompression.xml:124`, `/workspace_addons/BuffHead/Setup/SetupAdvancedCompressionItem.xml:200`, `/workspace_addons/BuffHead/Setup/SetupAdvancedCompressionItem.xml:216`, `/workspace_addons/BuffHead/Setup/SetupAdvancedCompressionItem.xml:230`, `/workspace_addons/BuffHead/Setup/SetupAdvancedCompressionItemEffect.xml:102`, `/workspace_addons/BuffHead/Setup/SetupAdvancedCompressionItemEffect.xml:88`, `/workspace_addons/BuffHead/Setup/SetupAdvancedContainers.xml:124` |
| Namespaces detected | EA_Button_DefaultResizeable |
| Source kinds | xml_attributes |
| Example locations | AdvancedRenownTrainingExportWindowCancelButton, AdvancedRenownTrainingExportWindowHyperLinkButton, AdvancedRenownTrainingExportWindowWardrobeButton, AdvancedRenownTrainingImportNameInputWindowCancelButton, AdvancedRenownTrainingImportNameInputWindowOkButton, AdvancedRenownTrainingImportWindowCancelButton |
| XML usage count | 150 |
| XML attribute usage count | 150 |
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

Observed engine XML template or inherited constant referenced by 29 addons.

## Seen In

- AdvancedRenownTrainer
- Aura
- BuffHead
- CM_ClosetGoblin
- DAoCBuff
- DaemonAssist
- EA_UiModWindow
- Enemy
- JunkDump
- LoyalPet
- MapMonster
- MapPin
- Miracle Grow Remix
- ObjectInspector
- Pocket Palette
- QuickWarReport
- RVAPI_ColorDialog
- RVMOD_Manager
- RandomMount
- Shinies
- TidyChat
- TidyRoll
- TurretRange
- WSCT
- WarBoard
- WhoHealedMe
- bigger_MacroWindow
- followTheLeader
- wbLeadHelper

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
- DAoCBuffMessageWindowScrollWindowScrollChildOpenOptions
- DAoCBuff_SettingsAddFrameButton
- DAoCBuff_SettingsReloadTrackers
- DaemonAssistWindowToggleButton
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
- JunkDumpButtonWin
- LPETOptionsAddProfileButton
- LPETOptionsLoadProfileButton
- LPETOptionsRemoveProfileButton
- LPETOptionsRenameProfileButton
- LPETOptionsSaveProfileButton
- MapMonsterEditorWindowButtonDefault
- MapMonsterPinTypeEditorWindowButtonDefault
- MapMonster_FilterButton
- MapPin_SetupAcceptButton
- MapPin_SetupCancelButton
- MiracleGrow2PresetTemplateLayoutButton
- MiracleGrow2PresetTemplateSettingsButton
- ObjectInspectorClearButton
- ObjectInspectorCloseButton
- ObjectInspectorInspectButton
- OpenUiModWindowButton
- PPMainCharacterWindowBtn
- PPMainTogglePickerBtn
- QuickWarReportConfirmAccept
- QuickWarReportConfirmDecline
- RVAPI_ColorDialogWindowButtonCancel
- RVAPI_ColorDialogWindowButtonOK
- RVMOD_ManagerWindowContentButtonReloadUI
- RandomMountWindowSettingsContentAddButton
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
- UiModAdvancedWindowCancelButton
- UiModAdvancedWindowDebugWindowButton
- UiModAdvancedWindowOkayButton
- UiModVersionMismatchWindowCancelButton
- UiModVersionMismatchWindowReEnableButton
- UiModWindowAdvancedButton
- UiModWindowCancelButton
- UiModWindowDisableAllButton
- UiModWindowEnableAllButton
- UiModWindowOkayButton
- UiModWindowReEnableAllAutoDisabledButton
- WBTogglerManagerButton
- WSCTButtonTemplate
- WSCTEventColor
- WSCTOptionsColorPickerWindowAcceptButton
- WbLeadHelperMessageCancelButton
- WbLeadHelperMessageOkButton
- WhoHealedMeOptionsContentCloseButton
- WhoHealedMeOptionsContentResetButton
- WhoHealedMeWindowOptionsButton
- followTheLeaderWindow
- wbLeadHelperConfigTabApplyButton
- wbLeadHelperConfigTabResetButton
- wbLeadHelperConfigTabSaveButton
- wbLeadHelperMessagesTabApplyButton
- wbLeadHelperMessagesTabResetButton
- wbLeadHelperMessagesTabSaveButton

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
