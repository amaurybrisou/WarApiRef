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
| Addons seen in | BuffHead, DAoCBuff, Enemy, Killer, TexturedButtons, TidyRoll, WSCT, WarBoard |
| Files seen in | `/workspace/data/raw/BuffHead/Setup/SetupAdvancedCompressionItemEffect.xml:0`, `/workspace/data/raw/BuffHead/Setup/SetupAdvancedContainersItemProperties.xml:0`, `/workspace/data/raw/BuffHead/Setup/SetupContainer.xml:0`, `/workspace/data/raw/BuffHead/Setup/SetupTrackers.xml:0`, `/workspace/data/raw/DAoCBuff/Source/DAoCBuffSettingsTabs.xml:145`, `/workspace/data/raw/DAoCBuff/Source/DAoCBuffSettingsTabs.xml:1668`, `/workspace/data/raw/DAoCBuff/Source/DAoCBuffSettingsTabs.xml:170`, `/workspace/data/raw/DAoCBuff/Source/DAoCBuffSettingsTabs.xml:1726` |
| Namespaces detected | EA_ComboBox_DefaultResizableSmall |
| Source kinds | xml_attributes |
| Example locations | AutoRollRowTemplateChoice, BuffHeadSetupAdvancedCompressionItemEffectWindowCastByComboBox, BuffHeadSetupAdvancedContainersItemPropertiesWindowElementGrowthHorizontalComboBox, BuffHeadSetupAdvancedContainersItemPropertiesWindowElementGrowthVerticalComboBox, BuffHeadSetupContainerWindowContainerAlwaysShowPlacementCombo, BuffHeadSetupContainerWindowContainerBuffsPlacementCombo |
| XML usage count | 52 |
| XML attribute usage count | 52 |
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

Observed engine XML template or inherited constant referenced by 9 addons.

## Seen In

- BuffHead
- DAoCBuff
- Enemy
- Killer
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
- DAoCBuffFrameSettingsTab_ScrollChild_BufforderComboBox
- DAoCBuffFrameSettingsTab_ScrollChild_CountComboBox
- DAoCBuffFrameSettingsTab_ScrollChild_DivideComboBox
- DAoCBuffFrameSettingsTab_ScrollChild_FontComboBox
- DAoCBuffFrameSettingsTab_ScrollChild_GrowHorizontalComboBox
- DAoCBuffFrameSettingsTab_ScrollChild_GrowLeftComboBox
- DAoCBuffFrameSettingsTab_ScrollChild_GrowUpComboBox
- DAoCBuffFrameSettingsTab_ScrollChild_PermabuffsComboBox
- DAoCBuffFrameSettingsTab_ScrollChild_RefreshComboBox
- DAoCBuffFrameSettingsTab_ScrollChild_RowComboBox
- DAoCBuffFrameSettingsTab_ScrollChild_StickComboBox
- DAoCBuffFrameSettingsTab_ScrollChild_TargetComboBox
- DAoCBuffFrameSettingsTab_ScrollChild_TypeComboBox
- DAoCBuffFrameSettings_G4Filter_G4UseandComboBox
- DAoCBuff_Settings_FilterFrame_ScrollChild_FilterPropertyComboBox
- DAoCBuff_Settings_FilterFrame_ScrollChild_TextureTypeComboBox
- DAoCBuff_Settings_FilterFrame_ScrollChild_UseandComboBox
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
