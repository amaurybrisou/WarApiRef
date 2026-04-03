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
| Addons seen in | BuffHead, CDown, ChattyCathy, DAoCBuff, DuffTimer, EA_OpenPartyWindow, Enemy, GDes |
| Files seen in | CDownSettingsTabs.xml, ChattyCathy.xml, Code/Core/Groups/EffectFilterDialog.xml, Code/UnitFrames/ClickCastingDialog.xml, Code/UnitFrames/EffectsIndicatorDialog.xml, Code/UnitFrames/UnitFramePartDialog.xml, Code/UnitFrames/UnitFramesConfiguration.xml, CustomAutoRoll.xml |
| Namespaces detected | EA_ComboBox_DefaultResizableSmall |
| Source kinds | xml_attributes |
| Example locations | AutoRollRowTemplateChoice, BuffHeadSetupAdvancedCompressionItemEffectWindowCastByComboBox, BuffHeadSetupAdvancedContainersItemPropertiesWindowElementGrowthHorizontalComboBox, BuffHeadSetupAdvancedContainersItemPropertiesWindowElementGrowthVerticalComboBox, BuffHeadSetupContainerWindowContainerAlwaysShowPlacementCombo, BuffHeadSetupContainerWindowContainerBuffsPlacementCombo |
| XML usage count | 117 |
| XML attribute usage count | 117 |
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

Engine-supplied XML constant or template class referenced by 28 addons.

## Seen In

- BuffHead
- CDown
- ChattyCathy
- DAoCBuff
- DuffTimer
- EA_OpenPartyWindow
- Enemy
- GDes
- Ges
- GroupRange
- ItemRack
- Killer
- LibAddonButton
- MapPin
- MarkBuff
- Obsidian
- RVMOD_Manager
- RVMOD_Targets
- SOR
- TastyButtons
- TaxPayer
- TexturedButtons
- TidyRoll
- TokenMachine
- WSCT
- WarBoard
- WarBoard_FPS
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
- CCBoxTemplateCombo
- CCBoxTemplateTabCombo
- CDownGeneralSettingsTab_ScrollChild_LayoutComboBox
- CDownGeneralSettingsTab_ScrollChild_MaxCDComboBox
- CDownGeneralSettingsTab_ScrollChild_MinCDComboBox
- CDownGeneralSettingsTab_ScrollChild_RefreshComboBox
- CDownGeneralSettingsTab_ScrollChild_SortOrderComboBox
- CDownNLayoutSettingsTab_ScrollChild_NCountComboBox
- CDownNLayoutSettingsTab_ScrollChild_NFontComboBox
- CDownNLayoutSettingsTab_ScrollChild_NGrowHorizontalComboBox
- CDownNLayoutSettingsTab_ScrollChild_NGrowLeftComboBox
- CDownNLayoutSettingsTab_ScrollChild_NGrowUpComboBox
- CDownNLayoutSettingsTab_ScrollChild_NRowComboBox
- CDownSLayoutSettingsTab_ScrollChild_SFontComboBox
- CDownSLayoutSettingsTab_ScrollChild_SGrowHorizontalComboBox
- CDownSLayoutSettingsTab_ScrollChild_SGrowLeftComboBox
- CDownSLayoutSettingsTab_ScrollChild_SGrowUpComboBox
- CDownSLayoutSettingsTab_ScrollChild_SMaxCountComboBox
- CDownSLayoutSettingsTab_ScrollChild_SNameFontComboBox
- ChattyCathyOptToCombo
- ChattyCathyOptWindowCombo
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
- DuffTimerOptions_ComboBox_cmb
- EA_Template_AutoRollComboBox
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
- GDesOptionsColourDeathColourComboBox
- GDesOptionsColourDefaultColourComboBox
- GDesOptionsColourFullHealthColourComboBox
- GDesOptionsColourHighHealthColourComboBox
- GDesOptionsColourMedHealthColourComboBox
- GDesOptionsColourMinHealthColourComboBox
- GDesOptionsColourOfflineColourComboBox
- GesOptionsColourDeathColourComboBox
- GesOptionsColourDefaultColourComboBox
- GesOptionsColourFullHealthColourComboBox
- GesOptionsColourHighHealthColourComboBox
- GesOptionsColourMedHealthColourComboBox
- GesOptionsColourMinHealthColourComboBox
- GesOptionsColourOfflineColourComboBox
- GroupRangeSetupStyleGroupBoxWindowGrowthCombo
- ItemRackSetsCombobox
- KillerSettingsWindowContentFeedFont
- KillerSettingsWindowContentPersonalKdMode
- LibAddonButtonManagerCustomItemWindowTypeComboBox
- MapPin_SetupComboRealm
- MarkBuffSetupSmartBuffWindowClassTypeList
- ObsidianSetupCastbarWindowElementFillUpdatePriorityComboBox
- ObsidianSetupCastbarWindowElementGlobalCooldownPositionComboBox
- ObsidianSetupCastbarWindowElementIconPositionComboBox
- ObsidianSetupCastbarWindowElementNameAlignmentComboBox
- ObsidianSetupCastbarWindowElementTimerAlignmentComboBox
- ObsidianSetupEffectTrackerWindowElementBarsElementFillUpdatePriorityComboBox
- ObsidianSetupEffectTrackerWindowElementBarsElementNameAlignmentComboBox
- ObsidianSetupEffectTrackerWindowElementBarsElementTimerAlignmentComboBox
- ObsidianSetupEffectTrackerWindowElementTrackerElementIconPositionComboBox
- RVMOD_ManagerWindowSortBy
- RVMOD_TargetsFrameRowTemplateTemplates
- SOROptions.DropList.DL
- TastyButtonsButtonSelectWindowButtonRangeViewComboFirstButton
- TastyButtonsButtonSelectWindowButtonRangeViewComboLastButton
- TastyButtonsGeneralOptionsProfileWindowComboProfiles
- TastyButtonsOptionsButtonInfoWindowComboButton
- TastyButtonsOptionsWindowGroupViewComboColumnDirection
- TastyButtonsOptionsWindowGroupViewComboRowDirection
- TastyButtonsOptionsWindowStateViewComboShowing
- TastyButtonsOptionsWindowStateViewComboStateType
- TaxPayerOptionsPayerComboBox
- TexturedButtonsSetupActionbarWindowSelectorComboBox
- TokenOptionTemplateChoice
- WSCTEventFrameCombo
- WarBoardComboAlignRow1
- WarBoardComboAlignRow2
- WarBoardComboAlignRow3
- WarBoardComboAlignRow4
- WarBoard_FPSOptions_cmbLName
- WhoHealedMeOptionsContentGroupOnlyCombo
- WhoHealedMeOptionsContentHudVisibleCombo
- WhoHealedMeOptionsContentMyHealsCombo
- WhoHealedMeOptionsContentSpellStorageCombo

## Related APIs

- [ComboBox](../../xml/element_types/element_ComboBox.md) (HIGH 100/100) - XML Element Type

## Notes

- none
