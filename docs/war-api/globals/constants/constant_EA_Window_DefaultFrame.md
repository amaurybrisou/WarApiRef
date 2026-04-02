# EA_Window_DefaultFrame

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
| Addons seen in | AdvancedPetAssist, AdvancedRenownTrainer, Busted, CM_ClosetGoblin, DAoCBuff, EA_UiDebugTools, EA_UiModWindow, Enemy |
| Files seen in | `/workspace_addons/AdvancedPetAssist/APAGui.xml:101`, `/workspace_addons/Busted/Busted.xml:19`, `/workspace_addons/ClosetGoblin/ClosetGoblin.xml:1232`, `/workspace_addons/ClosetGoblin/ClosetGoblin.xml:1310`, `/workspace_addons/ClosetGoblin/ClosetGoblin.xml:268`, `/workspace_addons/ClosetGoblin/ClosetGoblin.xml:384`, `/workspace_addons/ClosetGoblin/ClosetGoblin.xml:478`, `/workspace_addons/DAoCBuff/Source/DAoCBuffMsgWindow.xml:30` |
| Namespaces detected | EA_Window_DefaultFrame |
| Source kinds | xml_attributes |
| Example locations | APAOptionsBackground, AdvancedRenownTrainingExportWindowBackground, AdvancedRenownTrainingImportNameInputWindowBackground, AdvancedRenownTrainingImportWindowBackground, AdvancedRenownTrainingLinkWindowBackground, AdvancedRenownTrainingPresetsWindowBackground |
| XML usage count | 85 |
| XML attribute usage count | 85 |
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

Observed engine XML template or inherited constant referenced by 22 addons.

## Seen In

- AdvancedPetAssist
- AdvancedRenownTrainer
- Busted
- CM_ClosetGoblin
- DAoCBuff
- EA_UiDebugTools
- EA_UiModWindow
- Enemy
- Killer
- LoyalPet
- MapPin
- Miracle Grow Remix
- Pocket Palette
- PotionBar
- QuickWarReport
- RandomMount
- RoR_SoR
- TidyChat
- Tortall_DPS
- WSCT
- WhoHealedMe
- wbLeadHelper

## Used By

- APAOptionsBackground
- AdvancedRenownTrainingExportWindowBackground
- AdvancedRenownTrainingImportNameInputWindowBackground
- AdvancedRenownTrainingImportWindowBackground
- AdvancedRenownTrainingLinkWindowBackground
- AdvancedRenownTrainingPresetsWindowBackground
- AdvancedRenownTrainingWindowBackground
- BustedGUIBack
- ClosetGoblinCharacterWindowBackground
- ClosetGoblinCharacterWindowContentsActionBarSettingsBG
- ClosetGoblinCharacterWindowContentsSetListBackground
- ClosetGoblinZoneWindowBackground
- ClosetGoblinZoneWindowContentsSetListBackground
- DAoCBuffMessageWindowBackground
- DAoCBuff_SettingsBackground
- DebugWindowBackground
- DebugWindowLogDisplaySeperator2
- DebugWindowOptionsBackground
- DevPadConfirmLoadBackground
- DevPadConfirmLoadNameSeperator
- DevPadDeleteWindowBackground
- DevPadDeleteWindowNameSeperator
- DevPadNewWindowBackground
- DevPadNewWindowNameSeperator
- DevPadProjectLoadBackground
- DevPadProjectLoadNameSeperator
- DevPadRenameWindowBackground
- DevPadRenameWindowNameSeperator
- DevPadSaveWindowBackground
- DevPadSaveWindowNameSeperator
- DevPadWindowBackground
- DevPadWindowNameSeperator
- EnemyChooseChannelDialogBackground
- EnemyChooseIconDialogBackground
- EnemyClickCastingDialogBackground
- EnemyCombatLogSnapshotWindowBackground
- EnemyCombatLogStatsWindowBackground
- EnemyConfigDialogBackground
- EnemyEffectFilterDialogBackground
- EnemyEffectsIndicatorDialogBackground
- EnemyIntercomDialogBackground
- EnemyIntercomJoinDialogBackground
- EnemyKillSpamAreaStatsDialogBackground
- EnemyMarkConfigDialogBackground
- EnemyTextEntryDialogBackground
- EnemyUnitFramePartDialogBackground
- KillerScoreDetailsWindowBackground
- KillerSettingsWindowBackground
- KillerWindowBackground
- LPETOptionsBackground
- MapPinChooseIconDialogBackground
- MiracleGrow2Background
- ObjectInspectorNameSeperator
- PPMainBG
- PotionBarAboutBackground
- PotionBarButtonsBackground
- PotionBarTypeTemplateBackground
- QuickWarReportConfirmBackground
- RandomMountWindowBackground
- RoR_SoR_OffsetBackground
- RoR_SoR_OpacityBackground
- RoR_SoR_ScaleBackground
- TidyChatCopyBackground
- TortallDPSDetailTemplateBG
- TortallDPSGroupTemplateBG
- TortallDPSMeterAbilitiesBG
- TortallDPSMeterBG
- TortallDPSMeterDetailWindowBG
- TortallDPSMeterGroupBG
- TortallDPSMeterGroupWindowBG
- TortallDPSMeterResetBG
- TortallDPSMeterScaleBG
- TortallDPSMeterTargetsBG
- TortallDPSMeterWhatHMBG
- TortallDPSMeterWhoHMBG
- TortallDPSToggleBG
- UiModAdvancedWindowCustomUIBackground
- UiModAdvancedWindowModsBackground
- WSCTOptionsBackground
- WSCTOptionsColorPickerWindowBackground
- WbLeadHelperMessageBackground
- WhoHealedMeDetailsBackground
- WhoHealedMeOptionsBackground
- WhoHealedMeWindowBackground
- wbLeadHelperChooseIconDialogBackground

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
