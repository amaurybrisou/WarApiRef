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
| Addons seen in | AdvancedRenownTrainer, Aura, AutoBand, BuffHead, CM_ClosetGoblin, DAoCBuff, DaemonAssist, EA_UiModWindow |
| Files seen in | `/workspace/Aura/Source/Templates.xml:18`, `/workspace/Autoband/AutoBandWindowConfig.xml:379`, `/workspace/Autoband/AutoBandWindowConfig.xml:61`, `/workspace/Autoband/AutoBandWindowTemplate.xml:100`, `/workspace/Autoband/AutoBandWindowTemplate.xml:112`, `/workspace/Autoband/AutoBandWindowTemplate.xml:124`, `/workspace/Autoband/AutoBandWindowTemplate.xml:217`, `/workspace/Autoband/AutoBandWindowTemplate.xml:229` |
| Namespaces detected | EA_Button_DefaultResizeable |
| Source kinds | xml_attributes |
| Example locations | AdvancedRenownTrainingExportWindowCancelButton, AdvancedRenownTrainingExportWindowHyperLinkButton, AdvancedRenownTrainingExportWindowWardrobeButton, AdvancedRenownTrainingImportNameInputWindowCancelButton, AdvancedRenownTrainingImportNameInputWindowOkButton, AdvancedRenownTrainingImportWindowCancelButton |
| XML usage count | 175 |
| XML attribute usage count | 175 |
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

Observed engine XML template or inherited constant referenced by 34 addons.

## Seen In

- AdvancedRenownTrainer
- Aura
- AutoBand
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
- MegaphonePlusPlus
- Miracle Grow Remix
- ObjectInspector
- Pocket Palette
- Queue Queuer
- QuickWarReport
- RVAPI_ColorDialog
- RVMOD_Manager
- RandomMount
- Shinies
- ThinkOutLoud
- TidyChat
- TidyQueue
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
- AutoBandWindowConfigClearTemplateButton
- AutoBandWindowConfigResetButton
- AutoBandWindowTemplateApplyButton
- AutoBandWindowTemplateDeleteButton
- AutoBandWindowTemplateSaveButton
- AutoBandWindowTemplateSaveCancelButton
- AutoBandWindowTemplateSaveOKButton
- AutoBandWindowToolsKickButton
- AutoBandWindowToolsPrintButton
- AutoBandWindowToolsResetButton
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
- MegaphoneMainCloseButton
- MegaphoneMainTestButton
- MiracleGrow2PresetTemplateLayoutButton
- MiracleGrow2PresetTemplateSettingsButton
- ObjectInspectorClearButton
- ObjectInspectorCloseButton
- ObjectInspectorInspectButton
- OpenUiModWindowButton
- PPMainCharacterWindowBtn
- PPMainTogglePickerBtn
- QueueQueuer_GUI_BlacklistAllButton
- QueueQueuer_GUI_BlacklistNoneButton
- QueueQueuer_GUI_JoinButton
- QueueQueuer_GUI_LeaveButton
- QueueQueuer_GUI_QueuerCheckButton
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
- TOLSettingsWindowPhraseEditWindowDeleteButton
- TOLSettingsWindowPhraseEditWindowSaveButton
- TOLSettingsWindowSkillEditWindowDeleteButton
- TOLSettingsWindowSkillEditWindowEditPhrasesButton
- TOLSettingsWindowSkillEditWindowSaveButton
- TRollButton
- TidyChatCopyNext
- TidyChatCopyPrev
- TidyQueueJoinSelected
- TidyQueueSelectAll
- TidyQueueSelectNone
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
