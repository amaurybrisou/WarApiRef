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
| Addons seen in | AdvancedPetAssist, AdvancedRenownTrainer, AggroMeter, Aura, BankArkel, BuffHead, CM_ClosetGoblin, DAoCBuff |
| Files seen in | `/workspace/data/raw/AdvancedPetAssist/APAGui.xml:0`, `/workspace/data/raw/AggroMeter/AggroMeter.xml:0`, `/workspace/data/raw/Aura/Source/AuraConfig.xml:0`, `/workspace/data/raw/Aura/Source/AuraSettings.xml:0`, `/workspace/data/raw/Aura/Source/AuraShares.xml:0`, `/workspace/data/raw/Aura/Source/AuraTexture.xml:0`, `/workspace/data/raw/BankArkel/BankArkel.xml:0`, `/workspace/data/raw/BuffHead/Setup/SetupAdvancedCompression.xml:0` |
| Namespaces detected | EA_Button_DefaultWindowClose |
| Source kinds | xml_attributes |
| Example locations | APAOptionsClose, AdvancedRenownTrainingLinkWindowClose, AdvancedRenownTrainingWindowClose, AggroMeterGrayWindowClose, AuraConfigClose, AuraSettingsClose |
| XML usage count | 84 |
| XML attribute usage count | 84 |
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

- AdvancedPetAssist
- AdvancedRenownTrainer
- AggroMeter
- Aura
- BankArkel
- BuffHead
- CM_ClosetGoblin
- DAoCBuff
- Enemy
- Killer
- LibGroup
- Pocket Palette
- PotionBar
- RoR_SoR
- Shinies
- TexturedButtons
- TidyChat
- TidyRoll
- TurretRange
- WSCT
- WarBoard
- WhoHealedMe
- bigger_MacroWindow

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
- ClosetGoblinCharacterWindowClose
- ClosetGoblinZoneWindowClose
- DAoCBuffMessageWindowClose
- DAoCBuff_SettingsClose
- DAoCBuff_Settings_FilterClose
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
- KillerScoreDetailsWindowClose
- KillerSettingsWindowClose
- LibGroupSetupWindowClose
- MacroIconSelectionWindowClose
- PPMainClose
- PackWinClose
- PotionBarAboutClose
- PotionBarTypeTemplateClose
- RoR_SoR_OffsetClose
- RoR_SoR_OpacityClose
- RoR_SoR_ScaleClose
- ShiniesWindowClose
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
- WSCTOptionsClose
- WSCTOptionsColorPickerWindowClose
- WarBoardOptionsClose
- WhoHealedMeDetailsClose
- WhoHealedMeOptionsClose

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
