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
| Addons seen in | AdvancedPetAssist, AdvancedRenownTrainer, Aura, BankArkel, CM_ClosetGoblin, DAoCBuff, Enemy, Pocket Palette |
| Files seen in | `/workspace/data/raw/AdvancedPetAssist/APAGui.xml:0`, `/workspace/data/raw/Aura/Source/AuraConfig.xml:0`, `/workspace/data/raw/Aura/Source/AuraSettings.xml:0`, `/workspace/data/raw/Aura/Source/AuraShares.xml:0`, `/workspace/data/raw/Aura/Source/AuraTexture.xml:0`, `/workspace/data/raw/BankArkel/BankArkel.xml:0`, `/workspace/data/raw/ClosetGoblin/ClosetGoblin.xml:0`, `/workspace/data/raw/DAoCBuff/Source/DAoCBuffMsgWindow.xml:0` |
| Namespaces detected | EA_TitleBar_Default |
| Source kinds | xml_attributes |
| Example locations | APAOptionsTitle, AdvancedRenownTrainingExportWindowTitleBar, AdvancedRenownTrainingImportNameInputWindowTitleBar, AdvancedRenownTrainingImportWindowTitleBar, AdvancedRenownTrainingLinkWindowTitleBar, AdvancedRenownTrainingWindowTitleBar |
| XML usage count | 46 |
| XML attribute usage count | 46 |
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

Observed engine XML template or inherited constant referenced by 14 addons.

## Seen In

- AdvancedPetAssist
- AdvancedRenownTrainer
- Aura
- BankArkel
- CM_ClosetGoblin
- DAoCBuff
- Enemy
- Pocket Palette
- PotionBar
- RoR_SoR
- TidyChat
- WSCT
- WarBoard
- bigger_MacroWindow

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
- DAoCBuff_Settings_FilterTitleBar
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
- ItemWindowTitleBar
- MacroIconSelectionWindowTitleBar
- PPMainTitleBar
- PackWinTitleBar
- PotionBarAboutTitleBar
- PotionBarTypeTemplateTitleBar
- RoR_SoR_OffsetTitleBar
- RoR_SoR_OpacityTitleBar
- RoR_SoR_ScaleTitleBar
- TidyChatCopyTitleBar
- TidyChatLootRollTitleBar
- WSCTOptionsColorPickerWindowTitle
- WSCTOptionsTitle
- WarBoardOptionsTitleBar

## Related APIs

- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Used With

- none

## Triggered By

- none

## Affects

- none

## Notes

- none
