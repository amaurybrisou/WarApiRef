# EA_Button_DefaultWindowClose

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
| Addons seen in | AdvancedPetAssist, AdvancedRenownTrainer, AggroMeter, Aura, AutoBand, BankArkel, BuffHead, CM_ClosetGoblin |
| Files seen in | `/workspace/AdvancedPetAssist/APAGui.xml:114`, `/workspace/AggroMeter/AggroMeter.xml:84`, `/workspace/Aura/Source/AuraConfig.xml:37`, `/workspace/Aura/Source/AuraSettings.xml:87`, `/workspace/Aura/Source/AuraShares.xml:162`, `/workspace/Aura/Source/AuraShares.xml:29`, `/workspace/Aura/Source/AuraTexture.xml:106`, `/workspace/Autoband/AutoBandWindow.xml:29` |
| Namespaces detected | EA_Button_DefaultWindowClose |
| Source kinds | xml_attributes |
| Example locations | APAOptionsClose, AdvancedRenownTrainingLinkWindowClose, AdvancedRenownTrainingWindowClose, AggroMeterGrayWindowClose, AuraConfigClose, AuraSettingsClose |
| XML usage count | 111 |
| XML attribute usage count | 111 |
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

Observed engine XML template or inherited constant referenced by 43 addons.

## Seen In

- AdvancedPetAssist
- AdvancedRenownTrainer
- AggroMeter
- Aura
- AutoBand
- BankArkel
- BuffHead
- CM_ClosetGoblin
- Cheeseboard
- DAoCBuff
- DaemonAssist
- DeepSleep
- EA_UiModWindow
- Enemy
- FastInteract
- GetStats
- JunkDump
- Killer
- LibGroup
- LoyalPet
- MapMonster
- MapPin
- MegaphonePlusPlus
- ObjectInspector
- Pocket Palette
- PotionBar
- Queue Queuer
- QuickTacticSwitch
- QuickWarReport
- RVMOD_Manager
- RandomMount
- RoR_SoR
- Shinies
- TexturedButtons
- ThinkOutLoud
- TidyChat
- TidyRoll
- TurretRange
- WSCT
- WarBoard
- WhoHealedMe
- bigger_MacroWindow
- wbLeadHelper

## Used By

- APAOptionsClose
- AdvancedRenownTrainingLinkWindowClose
- AdvancedRenownTrainingWindowClose
- AggroMeterGrayWindowClose
- AuraConfigClose
- AuraSettingsClose
- AuraSharesClose
- AuraSharesImportExportClose
- AuraTextureClose
- AutoBandWindowClose
- AutoBandWindowTemplateSaveClose
- BuffHeadSetupAdvancedCompressionItemEffectWindowClose
- BuffHeadSetupAdvancedCompressionItemWindowClose
- BuffHeadSetupAdvancedCompressionWindowClose
- BuffHeadSetupAdvancedContainersItemPropertiesWindowClose
- BuffHeadSetupAdvancedContainersItemWindowClose
- BuffHeadSetupAdvancedContainersWindowClose
- BuffHeadSetupContainerWindowClose
- BuffHeadSetupDisplayWindowClose
- BuffHeadSetupEffectCacheTableWindowClose
- BuffHeadSetupEffectCacheWindowClose
- BuffHeadSetupFilterWindowClose
- BuffHeadSetupGeneralWindowClose
- BuffHeadSetupLayoutManagerWindowClose
- BuffHeadSetupLayoutPropertiesWindowClose
- BuffHeadSetupLayoutWindowClose
- BuffHeadSetupMenuWindowClose
- BuffHeadSetupPerformanceWindowClose
- BuffHeadSetupPriorityEffectsItemWindowClose
- BuffHeadSetupPriorityEffectsWindowClose
- BuffHeadSetupTrackersWindowClose
- CheeseboardWindowClose
- ClosetGoblinCharacterWindowClose
- ClosetGoblinZoneWindowClose
- DAoCBuffMessageWindowClose
- DAoCBuff_SettingsClose
- DaemonAssistWindowClose
- DeepSleep_SettingsClose
- EA_Window_MacroClose
- EnemyChooseChannelDialogClose
- EnemyChooseIconDialogClose
- EnemyClickCastingDialogClose
- EnemyCombatLogSnapshotWindowClose
- EnemyCombatLogStatsWindowClose
- EnemyConfigDialogClose
- EnemyEffectFilterDialogClose
- EnemyEffectsIndicatorDialogClose
- EnemyIntercomDialogClose
- EnemyIntercomJoinDialogClose
- EnemyKillSpamAreaStatsDialogClose
- EnemyMarkConfigDialogClose
- EnemyScenarioInfoDialogCancelButton
- EnemyTextEntryDialogClose
- EnemyUnitFramePartDialogClose
- FastInteractWindowClose
- GetStatsWindowClose
- JunkDumpOptionsWinClose
- KillerScoreDetailsWindowClose
- KillerSettingsWindowClose
- LPETOptionsClose
- LibGroupSetupWindowClose
- MacroIconSelectionWindowClose
- MapMonster_CalibrateWindowClose
- MapMonster_EditorWindowClose
- MapMonster_IconChooserWindowClose
- MapMonster_PinTypeEditorWindowClose
- MapPinChooseIconDialogClose
- MapPin_SetupClose
- MegaphoneMainClose
- ObjectInspectorClose
- PPMainClose
- PackWinClose
- PotionBarAboutClose
- PotionBarTypeTemplateClose
- QueueQueuer_GUIClose
- QuickTacticSwitchWindowClose
- QuickWarReportConfirmClose
- RVMOD_ManagerWindowClose
- RandomMountWindowClose
- RoR_SoR_OffsetClose
- RoR_SoR_OpacityClose
- RoR_SoR_ScaleClose
- ShiniesWindowClose
- TOLSettingsMainWindowClose
- TRollAutoRollCloseButton
- TexturedButtonsSetupActionbarWindowClose
- TexturedButtonsSetupAdvancedTexturesWindowClose
- TexturedButtonsSetupCooldownWindowClose
- TexturedButtonsSetupFontsWindowClose
- TexturedButtonsSetupMenuWindowClose
- TexturedButtonsSetupMiscWindowClose
- TexturedButtonsSetupTexturesWindowClose
- TexturedButtonsSetupTintWindowClose
- TidyChatCopyClose
- TidyChatLootRollClose
- TidyChatOptionsCloseButton
- TurretRangeSetupDisplayWindowClose
- TurretRangeSetupDistanceWindowClose
- TurretRangeSetupDistancesWindowClose
- TurretRangeSetupGeneralWindowClose
- TurretRangeSetupMenuWindowClose
- UiModAdvancedWindowClose
- UiModVersionMismatchWindowClose
- UiModWindowClose
- WSCTOptionsClose
- WSCTOptionsColorPickerWindowClose
- WarBoardOptionsClose
- WbLeadHelperMessageClose
- WhoHealedMeOptionsClose
- wbLeadHelperChooseIconDialogClose
- wbLeadHelperConfigWindowClose

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
