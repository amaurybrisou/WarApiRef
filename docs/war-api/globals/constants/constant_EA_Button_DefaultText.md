# EA_Button_DefaultText

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
| Addons seen in | Cheeseboard, DuffTimer, EA_ScenarioGroupWindow, EA_UiDebugTools, Enemy, GetStats, KillTracker, TastyButtons |
| Files seen in | Cheeseboard.xml, Code/Assist/AssistConfiguration.xml, Code/CombatLog/CombatLogStatsWindow.xml, Code/Core/ConfigDialog.xml, Code/ScenarioInfo/ScenarioInfo.xml, Code/UnitFrames/EffectsIndicatorDialog.xml, Code/UnitFrames/UnitFramesConfiguration.xml, DuffTimerOptionsDefn.xml |
| Namespaces detected | EA_Button_DefaultText |
| Source kinds | xml_attributes |
| Example locations | CheeseboardWindowButton, DebugWindowClose, DebugWindowOptionsClearLogText, DebugWindowOptionsClose, DebugWindowReloadUi, DebugWindowToggleCopy |
| XML usage count | 95 |
| XML attribute usage count | 95 |
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

Engine-supplied XML constant or template class referenced by 9 addons.

## Seen In

- Cheeseboard
- DuffTimer
- EA_ScenarioGroupWindow
- EA_UiDebugTools
- Enemy
- GetStats
- KillTracker
- TastyButtons
- wbLeadHelper

## Used By

- CheeseboardWindowButton
- DebugWindowClose
- DebugWindowOptionsClearLogText
- DebugWindowOptionsClose
- DebugWindowReloadUi
- DebugWindowToggleCopy
- DebugWindowToggleDevPad
- DebugWindowToggleLogging
- DebugWindowToggleObject
- DebugWindowToggleOptions
- DevPadConfirmLoadCancel
- DevPadConfirmLoadLoad
- DevPadConfirmLoadSaveLoad
- DevPadDeleteWindowCancel
- DevPadDeleteWindowDeleteFile
- DevPadNewWindowCancel
- DevPadNewWindowNewFile
- DevPadProjectLoadCancel
- DevPadProjectLoadLoadFile
- DevPadRenameWindowCancel
- DevPadRenameWindowRenameFile
- DevPadSaveWindowCancel
- DevPadSaveWindowSaveFile
- DevPadWindowClose
- DevPadWindowExecute
- DevPadWindowFile
- DevPadWindowSave
- DevPadWindowUndo
- DuffTimerOptionsAdvHeaderText
- EnemyAssistConfigurationMarkNewTargetEditTemplateButton
- EnemyCombatLogStatsWindowSessionAddButton
- EnemyCombatLogStatsWindowSessionDeleteButton
- EnemyCombatLogStatsWindowSessionRenameButton
- EnemyConfigDialogResetAllButton
- EnemyConfigDialogResetButton
- EnemyEffectsIndicatorDialogContentScrollChildChooseIconButton
- EnemyEffectsIndicatorDialogContentScrollChildEffectFiltersAddButton
- EnemyEffectsIndicatorDialogContentScrollChildEffectFiltersDeleteButton
- EnemyEffectsIndicatorDialogContentScrollChildEffectFiltersDownButton
- EnemyEffectsIndicatorDialogContentScrollChildEffectFiltersEditButton
- EnemyEffectsIndicatorDialogContentScrollChildEffectFiltersUpButton
- EnemyScenarioInfoDialogCancel2Button
- EnemyScenarioInfoDialogSwitchToStandardButton
- EnemyScenarioInfoToggleButtonTemplate
- EnemyUnitFramesConfigurationContentScrollChildClickCastingsAddButton
- EnemyUnitFramesConfigurationContentScrollChildClickCastingsDeleteButton
- EnemyUnitFramesConfigurationContentScrollChildClickCastingsDownButton
- EnemyUnitFramesConfigurationContentScrollChildClickCastingsEditButton
- EnemyUnitFramesConfigurationContentScrollChildClickCastingsEnableButton
- EnemyUnitFramesConfigurationContentScrollChildClickCastingsExportButton
- EnemyUnitFramesConfigurationContentScrollChildClickCastingsImportButton
- EnemyUnitFramesConfigurationContentScrollChildClickCastingsResetButton
- EnemyUnitFramesConfigurationContentScrollChildClickCastingsUpButton
- EnemyUnitFramesConfigurationContentScrollChildEffectsIndicatorsAddButton
- EnemyUnitFramesConfigurationContentScrollChildEffectsIndicatorsDeleteButton
- EnemyUnitFramesConfigurationContentScrollChildEffectsIndicatorsDownButton
- EnemyUnitFramesConfigurationContentScrollChildEffectsIndicatorsEditButton
- EnemyUnitFramesConfigurationContentScrollChildEffectsIndicatorsEnableButton
- EnemyUnitFramesConfigurationContentScrollChildEffectsIndicatorsExportButton
- EnemyUnitFramesConfigurationContentScrollChildEffectsIndicatorsImportButton
- EnemyUnitFramesConfigurationContentScrollChildEffectsIndicatorsResetButton
- EnemyUnitFramesConfigurationContentScrollChildEffectsIndicatorsUpButton
- EnemyUnitFramesConfigurationContentScrollChildUnitFramePartsAddButton
- EnemyUnitFramesConfigurationContentScrollChildUnitFramePartsDeleteButton
- EnemyUnitFramesConfigurationContentScrollChildUnitFramePartsDownButton
- EnemyUnitFramesConfigurationContentScrollChildUnitFramePartsEditButton
- EnemyUnitFramesConfigurationContentScrollChildUnitFramePartsEnableButton
- EnemyUnitFramesConfigurationContentScrollChildUnitFramePartsExportButton
- EnemyUnitFramesConfigurationContentScrollChildUnitFramePartsImportButton
- EnemyUnitFramesConfigurationContentScrollChildUnitFramePartsResetButton
- EnemyUnitFramesConfigurationContentScrollChildUnitFramePartsUpButton
- GetStatsWindowRefresh
- JoinScenarioGroupButton
- KillTracker_KillCountEntryLeftTemplateName
- KillTracker_KillCountEntryRightTemplateName
- ObjectInspectorClearButton
- ObjectInspectorCloseButton
- ObjectInspectorInspectButton
- ScenarioGroupWindowLeaveButton
- TastyButtonsOptionsWindowStateViewStateModeButton
- TastyButtonsOptions_ButtonSelectTemplateButton
- TastyButtonsOptions_ButtonSelectTemplateButtonGroup
- TastyButtonsOptions_InfoWindowStanceTemplateButton
- TastyButtonsOptions_InfoWindowStateTemplateButton
- wbLeadHelperConfigTabChooseIconButtonMessageEnd
- wbLeadHelperConfigTabChooseIconButtonMessageStart
- wbLeadHelperMessagesTabAddButton
- wbLeadHelperMessagesTabCloneButton
- wbLeadHelperMessagesTabDeleteButton
- wbLeadHelperMessagesTabDownButton
- wbLeadHelperMessagesTabEditButton
- wbLeadHelperMessagesTabEnableButton
- wbLeadHelperMessagesTabUpButton
- wbLeadHelperWindowChatButton
- wbLeadHelperWindowCloseButton

## Related APIs

- [Button](../../xml/element_types/element_Button.md) (HIGH 100/100) - XML Element Type

## Notes

- none
