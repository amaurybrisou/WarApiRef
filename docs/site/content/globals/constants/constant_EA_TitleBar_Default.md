# EA_TitleBar_Default

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
| Addons seen in | AdvancedPetAssist, AdvancedRenownTrainer, Aura, BankArkel, CM_ClosetGoblin, DAoCBuff, DaemonAssist, EA_UiModWindow |
| Files seen in | `/workspace_addons/AdvancedPetAssist/APAGui.xml:111`, `/workspace_addons/Aura/Source/AuraConfig.xml:35`, `/workspace_addons/Aura/Source/AuraSettings.xml:85`, `/workspace_addons/Aura/Source/AuraShares.xml:160`, `/workspace_addons/Aura/Source/AuraShares.xml:28`, `/workspace_addons/Aura/Source/AuraTexture.xml:105`, `/workspace_addons/BankArkel/BankArkel.xml:32`, `/workspace_addons/ClosetGoblin/ClosetGoblin.xml:1242` |
| Namespaces detected | EA_TitleBar_Default |
| Source kinds | xml_attributes |
| Example locations | APAOptionsTitle, AdvancedRenownTrainingExportWindowTitleBar, AdvancedRenownTrainingImportNameInputWindowTitleBar, AdvancedRenownTrainingImportWindowTitleBar, AdvancedRenownTrainingLinkWindowTitleBar, AdvancedRenownTrainingWindowTitleBar |
| XML usage count | 66 |
| XML attribute usage count | 66 |
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

Observed engine XML template or inherited constant referenced by 27 addons.

## Seen In

- AdvancedPetAssist
- AdvancedRenownTrainer
- Aura
- BankArkel
- CM_ClosetGoblin
- DAoCBuff
- DaemonAssist
- EA_UiModWindow
- Enemy
- FastInteract
- JunkDump
- LoyalPet
- MapMonster
- MapPin
- ObjectInspector
- Pocket Palette
- PotionBar
- QuickTacticSwitch
- QuickWarReport
- RVMOD_Manager
- RandomMount
- RoR_SoR
- TidyChat
- WSCT
- WarBoard
- bigger_MacroWindow
- wbLeadHelper

## Used By

- APAOptionsTitle
- AdvancedRenownTrainingExportWindowTitleBar
- AdvancedRenownTrainingImportNameInputWindowTitleBar
- AdvancedRenownTrainingImportWindowTitleBar
- AdvancedRenownTrainingLinkWindowTitleBar
- AdvancedRenownTrainingWindowTitleBar
- AuraConfigTitleBar
- AuraSettingsTitleBar
- AuraSharesImportExportTitleBar
- AuraSharesTitleBar
- AuraTextureTitleBar
- ClosetGoblinCharacterWindowTitleBar
- ClosetGoblinZoneWindowTitleBar
- DAoCBuffMessageWindowTitleBar
- DAoCBuff_SettingsTitleBar
- DaemonAssistWindowTitleBar
- DyeWindowTitleBar
- EA_Window_MacroTitleBar
- EnemyChooseChannelDialogTitleBar
- EnemyChooseIconDialogTitleBar
- EnemyClickCastingDialogTitleBar
- EnemyCombatLogSnapshotWindowTitleBar
- EnemyCombatLogStatsWindowTitleBar
- EnemyConfigDialogTitleBar
- EnemyEffectFilterDialogTitleBar
- EnemyEffectsIndicatorDialogTitleBar
- EnemyIntercomDialogTitleBar
- EnemyIntercomJoinDialogTitleBar
- EnemyKillSpamAreaStatsDialogTitleBar
- EnemyMarkConfigDialogTitleBar
- EnemyTextEntryDialogTitleBar
- EnemyUnitFramePartDialogTitleBar
- FastInteractWindowTitleBar
- ItemWindowTitleBar
- JunkDumpOptionsWinTitleBar
- LPETOptionsTitle
- MacroIconSelectionWindowTitleBar
- MapMonster_CalibrateWindowTitleBar
- MapMonster_EditorWindowTitleBar
- MapMonster_IconChooserWindowTitleBar
- MapMonster_PinTypeEditorWindowTitleBar
- MapPinChooseIconDialogTitleBar
- MapPin_SetupTitleBar
- ObjectInspectorTitleBar
- PPMainTitleBar
- PackWinTitleBar
- PotionBarAboutTitleBar
- PotionBarTypeTemplateTitleBar
- QuickTacticSwitchWindowTitleBar
- QuickWarReportConfirmTitleBar
- RVMOD_ManagerWindowTitleBar
- RandomMountWindowTitleBar
- RoR_SoR_OffsetTitleBar
- RoR_SoR_OpacityTitleBar
- RoR_SoR_ScaleTitleBar
- TidyChatCopyTitleBar
- TidyChatLootRollTitleBar
- UiModAdvancedWindowTitleBar
- UiModVersionMismatchWindowTitleBar
- UiModWindowTitleBar
- WSCTOptionsColorPickerWindowTitle
- WSCTOptionsTitle
- WarBoardOptionsTitleBar
- WbLeadHelperMessageTitleBar
- wbLeadHelperChooseIconDialogTitleBar
- wbLeadHelperConfigWindowTitleBar

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
