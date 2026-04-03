# EA_Window_Default

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
| Addons seen in | AdvancedRenownTrainer, Aura, BuffHead, CM_ClosetGoblin, DAoCBuff, Enemy, LibGroup, LibWBToggler |
| Files seen in | `/workspace/data/raw/Aura/Source/AuraColorPicker.xml:0`, `/workspace/data/raw/BuffHead/Setup/SetupAdvancedCompression.xml:0`, `/workspace/data/raw/BuffHead/Setup/SetupAdvancedCompressionItem.xml:0`, `/workspace/data/raw/BuffHead/Setup/SetupAdvancedCompressionItemEffect.xml:0`, `/workspace/data/raw/BuffHead/Setup/SetupAdvancedContainers.xml:0`, `/workspace/data/raw/BuffHead/Setup/SetupAdvancedContainersItem.xml:0`, `/workspace/data/raw/BuffHead/Setup/SetupAdvancedContainersItemProperties.xml:0`, `/workspace/data/raw/BuffHead/Setup/SetupContainer.xml:0` |
| Namespaces detected | EA_Window_Default |
| Source kinds | xml_attributes |
| Example locations | AdvancedRenownTrainingExportWindow, AdvancedRenownTrainingImportNameInputWindow, AdvancedRenownTrainingImportWindow, AdvancedRenownTrainingLinkWindow, AdvancedRenownTrainingPresetsWindow, AdvancedRenownTrainingWindow |
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

Observed engine XML template or inherited constant referenced by 13 addons.

## Seen In

- AdvancedRenownTrainer
- Aura
- BuffHead
- CM_ClosetGoblin
- DAoCBuff
- Enemy
- LibGroup
- LibWBToggler
- Shinies
- TexturedButtons
- TurretRange
- WarBoard
- bigger_MacroWindow

## Used By

- AdvancedRenownTrainingExportWindow
- AdvancedRenownTrainingImportNameInputWindow
- AdvancedRenownTrainingImportWindow
- AdvancedRenownTrainingLinkWindow
- AdvancedRenownTrainingPresetsWindow
- AdvancedRenownTrainingWindow
- AuraColorPicker
- BuffHeadSetupAdvancedCompressionItemEffectWindow
- BuffHeadSetupAdvancedCompressionItemWindow
- BuffHeadSetupAdvancedCompressionWindow
- BuffHeadSetupAdvancedContainersItemPropertiesWindow
- BuffHeadSetupAdvancedContainersItemWindow
- BuffHeadSetupAdvancedContainersWindow
- BuffHeadSetupContainerWindow
- BuffHeadSetupDisplayWindow
- BuffHeadSetupEffectCacheTableWindow
- BuffHeadSetupEffectCacheWindow
- BuffHeadSetupFilterWindow
- BuffHeadSetupGeneralWindow
- BuffHeadSetupLayoutManagerWindow
- BuffHeadSetupLayoutPropertiesWindow
- BuffHeadSetupLayoutWindow
- BuffHeadSetupMenuWindow
- BuffHeadSetupPerformanceWindow
- BuffHeadSetupPriorityEffectsItemWindow
- BuffHeadSetupPriorityEffectsWindow
- BuffHeadSetupTrackersWindow
- CG_ItemRackEQShow1
- ClosetGoblinCharacterWindow
- ClosetGoblinOptionWindow
- ClosetGoblinZoneWindow
- DAoCBuff_Settings
- DAoCBuff_Settings_Filter
- EA_Window_Macro
- EnemyChooseChannelDialog
- EnemyChooseIconDialog
- EnemyClickCastingDialog
- EnemyCombatLogEpsWindow
- EnemyCombatLogIDSAnchor
- EnemyCombatLogIDSRow
- EnemyCombatLogSnapshotWindow
- EnemyCombatLogStatsWindow
- EnemyCombatLogTargetDefeseTotalWindow
- EnemyCombatLogTargetDefeseWindow
- EnemyConfigDialog
- EnemyDebug
- EnemyEffectFilterDialog
- EnemyEffectsIndicatorDialog
- EnemyIntercomDialog
- EnemyIntercomJoinDialog
- EnemyKillSpamAreaStatsDialog
- EnemyKillSpamDialog
- EnemyKilledBy
- EnemyMarkConfigDialog
- EnemyPlayerKDR
- EnemyScenarioInfoDialog
- EnemyTextEntryDialog
- EnemyTimer
- EnemyUnitFramePartDialog
- EnemyUnitFramesAnchor1
- EnemyUnitFramesAnchor2
- LayoutTemplate
- LibGroupSetupWindow
- ShiniesAuctionsUI
- ShiniesAutoUI
- ShiniesBrowseUI_Searches
- ShiniesConfigUI
- ShiniesPostUI
- ShiniesWindow
- TexturedButtonsSetupActionbarWindow
- TexturedButtonsSetupAdvancedTexturesWindow
- TexturedButtonsSetupCooldownWindow
- TexturedButtonsSetupFontsWindow
- TexturedButtonsSetupMenuWindow
- TexturedButtonsSetupMiscWindow
- TexturedButtonsSetupTexturesWindow
- TexturedButtonsSetupTintWindow
- TurretRangeSetupDisplayWindow
- TurretRangeSetupDistanceWindow
- TurretRangeSetupDistancesWindow
- TurretRangeSetupGeneralWindow
- TurretRangeSetupMenuWindow
- WBTogglerTemplate
- WarBoard
- WarBoardAlignOptions
- WarBoardAlignOptionsWindow
- WarBoardBottomLayoutModeButtonWindow
- WarBoardBottomOptionsButtonWindow
- WarBoardButtonCheck
- WarBoardHorizontalPadding
- WarBoardLayoutModeButtonWindow
- WarBoardOptions
- WarBoardOptionsButtonWindow
- WarBoardPadOptionsWindow
- WarBoardVerticalPadding

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
