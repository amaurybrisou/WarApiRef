# EA_ComboBox_DefaultResizableSmall

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
| Addons seen in | BuffHead, Enemy, Killer, MapPin, RVMOD_Manager, TexturedButtons, TidyRoll, WSCT |
| Files seen in | `/workspace/BuffHead/Setup/SetupAdvancedCompressionItemEffect.xml:80`, `/workspace/BuffHead/Setup/SetupAdvancedContainersItemProperties.xml:843`, `/workspace/BuffHead/Setup/SetupAdvancedContainersItemProperties.xml:865`, `/workspace/BuffHead/Setup/SetupContainer.xml:200`, `/workspace/BuffHead/Setup/SetupContainer.xml:231`, `/workspace/BuffHead/Setup/SetupContainer.xml:262`, `/workspace/BuffHead/Setup/SetupTrackers.xml:58`, `/workspace/Enemy/Code/Core/Groups/EffectFilterDialog.xml:118` |
| Namespaces detected | EA_ComboBox_DefaultResizableSmall |
| Source kinds | xml_attributes |
| Example locations | AutoRollRowTemplateChoice, BuffHeadSetupAdvancedCompressionItemEffectWindowCastByComboBox, BuffHeadSetupAdvancedContainersItemPropertiesWindowElementGrowthHorizontalComboBox, BuffHeadSetupAdvancedContainersItemPropertiesWindowElementGrowthVerticalComboBox, BuffHeadSetupContainerWindowContainerAlwaysShowPlacementCombo, BuffHeadSetupContainerWindowContainerBuffsPlacementCombo |
| XML usage count | 37 |
| XML attribute usage count | 37 |
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

Observed engine XML template or inherited constant referenced by 10 addons.

## Seen In

- BuffHead
- Enemy
- Killer
- MapPin
- RVMOD_Manager
- TexturedButtons
- TidyRoll
- WSCT
- WarBoard
- WhoHealedMe

## Used By

- AutoRollRowTemplateChoice
- BuffHeadSetupAdvancedCompressionItemEffectWindowCastByComboBox
- BuffHeadSetupAdvancedContainersItemPropertiesWindowElementGrowthHorizontalComboBox
- BuffHeadSetupAdvancedContainersItemPropertiesWindowElementGrowthVerticalComboBox
- BuffHeadSetupContainerWindowContainerAlwaysShowPlacementCombo
- BuffHeadSetupContainerWindowContainerBuffsPlacementCombo
- BuffHeadSetupContainerWindowContainerDebuffsPlacementCombo
- BuffHeadSetupTrackersWindowTrackerTypeList
- EnemyClickCastingDialogContentScrollChildArchetypeMatch
- EnemyClickCastingDialogContentScrollChildPlayerTypeMatch
- EnemyEffectFilterDialogDescriptionMatch
- EnemyEffectFilterDialogNameMatch
- EnemyEffectFilterDialogTypeMatch
- EnemyEffectsIndicatorDialogContentScrollChildAnchorFrom
- EnemyEffectsIndicatorDialogContentScrollChildAnchorTo
- EnemyEffectsIndicatorDialogContentScrollChildArchetypeMatch
- EnemyEffectsIndicatorDialogContentScrollChildIcon
- EnemyEffectsIndicatorDialogContentScrollChildPlayerTypeMatch
- EnemyUnitFramePartDialogContentScrollChildArchetypeMatch
- EnemyUnitFramePartDialogContentScrollChildPlayerTypeMatch
- EnemyUnitFramesConfigurationContentScrollChildSorting1
- EnemyUnitFramesConfigurationContentScrollChildSorting2
- EnemyUnitFramesConfigurationContentScrollChildSorting3
- KillerSettingsWindowContentFeedFont
- KillerSettingsWindowContentPersonalKdMode
- MapPin_SetupComboRealm
- RVMOD_ManagerWindowSortBy
- TexturedButtonsSetupActionbarWindowSelectorComboBox
- WSCTEventFrameCombo
- WarBoardComboAlignRow1
- WarBoardComboAlignRow2
- WarBoardComboAlignRow3
- WarBoardComboAlignRow4
- WhoHealedMeOptionsContentGroupOnlyCombo
- WhoHealedMeOptionsContentHudVisibleCombo
- WhoHealedMeOptionsContentMyHealsCombo
- WhoHealedMeOptionsContentSpellStorageCombo

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
