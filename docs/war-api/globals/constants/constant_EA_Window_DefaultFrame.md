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
| Addons seen in | AdvancedPetAssist, AdvancedRenownTrainer, CM_ClosetGoblin, DAoCBuff, Enemy, Killer, Pocket Palette, PotionBar |
| Files seen in | `/workspace/data/raw/AdvancedPetAssist/APAGui.xml:0`, `/workspace/data/raw/ClosetGoblin/ClosetGoblin.xml:0`, `/workspace/data/raw/DAoCBuff/Source/DAoCBuffMsgWindow.xml:0`, `/workspace/data/raw/DAoCBuff/Source/DAoCBuffSettings.xml:0`, `/workspace/data/raw/DAoCBuff/Source/DAoCBuffSettingsTabs.xml:968`, `/workspace/data/raw/Enemy/Code/CombatLog/CombatLogSnapshotWindow.xml:0`, `/workspace/data/raw/Enemy/Code/CombatLog/CombatLogStatsWindow.xml:0`, `/workspace/data/raw/Enemy/Code/Core/ChooseIconDialog.xml:0` |
| Namespaces detected | EA_Window_DefaultFrame |
| Source kinds | xml_attributes |
| Example locations | APAOptionsBackground, AdvancedRenownTrainingExportWindowBackground, AdvancedRenownTrainingImportNameInputWindowBackground, AdvancedRenownTrainingImportWindowBackground, AdvancedRenownTrainingLinkWindowBackground, AdvancedRenownTrainingPresetsWindowBackground |
| XML usage count | 45 |
| XML attribute usage count | 45 |
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

- AdvancedPetAssist
- AdvancedRenownTrainer
- CM_ClosetGoblin
- DAoCBuff
- Enemy
- Killer
- Pocket Palette
- PotionBar
- RoR_SoR
- TidyChat
- WSCT
- WhoHealedMe

## Used By

- APAOptionsBackground
- AdvancedRenownTrainingExportWindowBackground
- AdvancedRenownTrainingImportNameInputWindowBackground
- AdvancedRenownTrainingImportWindowBackground
- AdvancedRenownTrainingLinkWindowBackground
- AdvancedRenownTrainingPresetsWindowBackground
- AdvancedRenownTrainingWindowBackground
- ClosetGoblinCharacterWindowBackground
- ClosetGoblinCharacterWindowContentsActionBarSettingsBG
- ClosetGoblinCharacterWindowContentsSetListBackground
- ClosetGoblinZoneWindowBackground
- ClosetGoblinZoneWindowContentsSetListBackground
- DAoCBuffMessageWindowBackground
- DAoCBuff_SettingsBackground
- DAoCBuff_Settings_FilterBackground
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
- PPMainBG
- PotionBarAboutBackground
- PotionBarButtonsBackground
- PotionBarTypeTemplateBackground
- RoR_SoR_OffsetBackground
- RoR_SoR_OpacityBackground
- RoR_SoR_ScaleBackground
- TidyChatCopyBackground
- WSCTOptionsBackground
- WSCTOptionsColorPickerWindowBackground
- WhoHealedMeDetailsBackground
- WhoHealedMeOptionsBackground
- WhoHealedMeWindowBackground

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
