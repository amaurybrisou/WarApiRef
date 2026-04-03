# EA_Default_SliderBar

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
| Addons seen in | AdjustTheTip, Atlas, Aura, BuffHead, CDown, ChattyCathy, DAoCBuff, DuffTimer |
| Files seen in | AdjustTheTip.xml, CDownSettingsTabs.xml, ChattyCathy.xml, Code/UnitFrames/EffectsIndicatorDialog.xml, DuffTimerOptionsDefn.xml, GroupIconsSG.xml, Gui/HealGridGuiTemplates.xml, Modules/Config/Shinies-Config-General.xml |
| Namespaces detected | EA_Default_SliderBar |
| Source kinds | xml_attributes |
| Example locations | AdjustTheTipMenuItemSliderSliderBar, AuraColorPickerAlpha, AuraColorPickerBlue, AuraColorPickerGreen, AuraColorPickerRed, BuffHeadSetupAdvancedContainersItemPropertiesWindowElementScaleSlider |
| XML usage count | 179 |
| XML attribute usage count | 179 |
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

Engine-supplied XML constant or template class referenced by 41 addons.

## Seen In

- AdjustTheTip
- Atlas
- Aura
- BuffHead
- CDown
- ChattyCathy
- DAoCBuff
- DuffTimer
- EA_ScenarioGroupWindow
- Enemy
- Group Icons SG
- GroupRange
- GroupSpotter
- HealGrid
- LibGroup
- MapMonster
- MapPin
- MoraleCircle
- Obsidian
- PotionBar
- RVAPI_ColorDialog
- RVAPI_Range
- RVMOD_Manager
- RVMOD_SquaredDistances
- ReliquaryHunter
- RoR_SoR
- SNT_INFO
- SNT_PANEL
- SOR
- Shinies
- Statdoll Remix
- TacticSetNames
- TastyButtons
- TexturedButtons
- TidyChat
- TurretRange
- WSCT
- WarBoard
- WarBoard_FPS
- XpStatus+G
- alertMod

## Used By

- AdjustTheTipMenuItemSliderSliderBar
- AuraColorPickerAlpha
- AuraColorPickerBlue
- AuraColorPickerGreen
- AuraColorPickerRed
- BuffHeadSetupAdvancedContainersItemPropertiesWindowElementScaleSlider
- BuffHeadSetupContainerWindowContainerColumnsSlider
- BuffHeadSetupContainerWindowContainerPaddingHorizontalSlider
- BuffHeadSetupContainerWindowContainerPaddingVerticalSlider
- BuffHeadSetupContainerWindowContainerRowsSlider
- BuffHeadSetupDisplayWindowIndicatorPaddingHorizontalSlider
- BuffHeadSetupDisplayWindowIndicatorPaddingVerticalSlider
- BuffHeadSetupDisplayWindowIndicatorScaleSlider
- BuffHeadSetupGeneralWindowOptionsMaximumThresholdSlider
- BuffHeadSetupLayoutPropertiesWindowElementAlphaAlphaSlider
- BuffHeadSetupLayoutPropertiesWindowElementIconBorderAlphaSlider
- BuffHeadSetupLayoutPropertiesWindowElementSizeScaleSlider
- BuffHeadSetupLayoutPropertiesWindowElementStatusBarBackgroundAlphaSlider
- BuffHeadSetupLayoutPropertiesWindowElementStatusBarForegroundAlphaSlider
- BuffHeadSetupPerformanceWindowGeneralUpdateDelaySlider
- BuffHeadSetupPerformanceWindowMaximumUpdatesSlider
- BuffHeadSetupPerformanceWindowPriorityUpdateDelaySlider
- BuffHeadSetupPerformanceWindowPriorityUpdateStartSlider
- BuffHeadSetupPerformanceWindowResyncTargetDelaySlider
- BuffHeadSetupSelectColorWindowTintBlueSlider
- BuffHeadSetupSelectColorWindowTintGreenSlider
- BuffHeadSetupSelectColorWindowTintRedSlider
- CDownColorSettingsTab_ScrollChild_AbilityColor_BlueSlider
- CDownColorSettingsTab_ScrollChild_AbilityColor_GreenSlider
- CDownColorSettingsTab_ScrollChild_AbilityColor_RedSlider
- CDownColorSettingsTab_ScrollChild_SingleColor_BlueSlider
- CDownColorSettingsTab_ScrollChild_SingleColor_GreenSlider
- CDownColorSettingsTab_ScrollChild_SingleColor_RedSlider
- CDownSLayoutSettingsTab_ScrollChild_SLengthSlider
- ChattyCathyOptAlphaSlider
- DAoCBuff_Settings_FilterFrame_ScrollChild_SliderBB
- DAoCBuff_Settings_FilterFrame_ScrollChild_SliderBG
- DAoCBuff_Settings_FilterFrame_ScrollChild_SliderBR
- DAoCBuff_Settings_FilterFrame_ScrollChild_SliderCB
- DAoCBuff_Settings_FilterFrame_ScrollChild_SliderCG
- DAoCBuff_Settings_FilterFrame_ScrollChild_SliderCR
- DuffTimerOptions_Slider_sld
- EnemyEffectsIndicatorDialogContentScrollChildAlpha
- EnemyEffectsIndicatorDialogContentScrollChildColorB
- EnemyEffectsIndicatorDialogContentScrollChildColorG
- EnemyEffectsIndicatorDialogContentScrollChildColorR
- GroupIconsSGOptionsSlider1
- GroupIconsSGOptionsSlider2
- GroupIconsSGOptionsSlider3
- GroupIconsSGOptionsSlider4
- GroupRangeSetupGeneralWindowUpdateDelaySlider
- GroupSpotterSettingsWindow_SlidergrpAlpha
- GroupSpotterSettingsWindow_SlidergrpB
- GroupSpotterSettingsWindow_SlidergrpG
- GroupSpotterSettingsWindow_SlidergrpR
- GroupSpotterSettingsWindow_SlidersingleAlpha
- GroupSpotterSettingsWindow_SlidersingleB
- GroupSpotterSettingsWindow_SlidersingleG
- GroupSpotterSettingsWindow_SlidersingleR
- HGG_SliderFullLabels
- LibGroupSetupWindowGroupDistanceCacheUpdateSlider
- LibGroupSetupWindowGroupDistanceSearchUpdateSlider
- LibGroupSetupWindowGroupUpdateDelaySlider
- MapMonster_PinTypeEditorWindowCustomColorBlueSlider
- MapMonster_PinTypeEditorWindowCustomColorGreenSlider
- MapMonster_PinTypeEditorWindowCustomColorRedSlider
- MapPin_SetupRotateFactorSlider
- MapPin_SetupScaleFactorSlider
- MoraleSlidersCustomColorBlueSlider
- MoraleSlidersCustomColorGreenSlider
- MoraleSlidersCustomColorRedSlider
- MoraleSlidersEmptyCustomColorBlueSlider
- MoraleSlidersEmptyCustomColorGreenSlider
- MoraleSlidersEmptyCustomColorRedSlider
- MoraleSlidersFillCustomColorBlueSlider
- MoraleSlidersFillCustomColorGreenSlider
- MoraleSlidersFillCustomColorRedSlider
- MoraleSlidersFullCustomColorBlueSlider
- MoraleSlidersFullCustomColorGreenSlider
- MoraleSlidersFullCustomColorRedSlider
- ObsidianSetupCastbarWindowElementBackgroundAlphaSlider
- ObsidianSetupCastbarWindowElementFillAlphaSlider
- ObsidianSetupCastbarWindowElementGlobalCooldownBackgroundAlphaSlider
- ObsidianSetupCastbarWindowElementGlobalCooldownDividerAlphaSlider
- ObsidianSetupCastbarWindowElementGlobalCooldownFillAlphaSlider
- ObsidianSetupCastbarWindowElementGlobalCooldownFillReadyAlphaSlider
- ObsidianSetupCastbarWindowElementGlobalCooldownSparkAlphaSlider
- ObsidianSetupCastbarWindowElementIconAlphaSlider
- ObsidianSetupCastbarWindowElementIconScaleSlider
- ObsidianSetupCastbarWindowElementLatencyAlphaSlider
- ObsidianSetupCastbarWindowElementLatencyTextAlphaSlider
- ObsidianSetupCastbarWindowElementLatencyTextScaleSlider
- ObsidianSetupCastbarWindowElementNameAlphaSlider
- ObsidianSetupCastbarWindowElementNameScaleSlider
- ObsidianSetupCastbarWindowElementNameWidthSlider
- ObsidianSetupCastbarWindowElementSparkAlphaSlider
- ObsidianSetupCastbarWindowElementTimerAlphaSlider
- ObsidianSetupCastbarWindowElementTimerDecimalsSlider
- ObsidianSetupCastbarWindowElementTimerScaleSlider
- ObsidianSetupEffectTrackerWindowElementBarsElementBackgroundAlphaSlider
- ObsidianSetupEffectTrackerWindowElementBarsElementFillAlphaSlider
- ObsidianSetupEffectTrackerWindowElementBarsElementNameAlphaSlider
- ObsidianSetupEffectTrackerWindowElementBarsElementNameScaleSlider
- ObsidianSetupEffectTrackerWindowElementBarsElementNameWidthSlider
- ObsidianSetupEffectTrackerWindowElementBarsElementTimerAlphaSlider
- ObsidianSetupEffectTrackerWindowElementBarsElementTimerScaleSlider
- ObsidianSetupEffectTrackerWindowElementTrackerElementGeneralMaximumEffectsSlider
- ObsidianSetupEffectTrackerWindowElementTrackerElementIconAlphaSlider
- ObsidianSetupEffectTrackerWindowElementTrackerElementIconScaleSlider
- ObsidianSetupSelectColorWindowTintBlueSlider
- ObsidianSetupSelectColorWindowTintGreenSlider
- ObsidianSetupSelectColorWindowTintRedSlider
- PotionBarTypeTemplateOpacitySlider
- PotionBarTypeTemplateScaleSlider
- RVAPI_ColorDialogSliderTemplate
- RVAPI_RangeSliderTemplate
- RVMOD_ManagerSliderTemplate
- RVMOD_SquaredDistancesSliderTemplate
- ReliquaryHunterOptionsWindowEnableWorldMapWindowBackgroundOpacitySlider
- ReliquaryHunterOptionsWindowEnableWorldMapWindowOpacitySlider
- RoR_SoR_OffsetSlider
- RoR_SoR_OpacitySlider
- RoR_SoR_ScaleSlider
- SNT_INFO_SETUP_WINDOW_SA
- SNT_INFO_SETUP_WINDOW_SB
- SNT_INFO_SETUP_WINDOW_SG
- SNT_INFO_SETUP_WINDOW_SR
- SNT_INFO_SETUP_WINDOW_SS
- SNT_PANEL_SETUP_WINDOW_A1
- SNT_PANEL_SETUP_WINDOW_A2
- SNT_PANEL_SETUP_WINDOW_A3
- SNT_PANEL_SETUP_WINDOW_B1
- SNT_PANEL_SETUP_WINDOW_B2
- SNT_PANEL_SETUP_WINDOW_G1
- SNT_PANEL_SETUP_WINDOW_G2
- SNT_PANEL_SETUP_WINDOW_R1
- SNT_PANEL_SETUP_WINDOW_R2
- SOROptions.Slider.slider
- SOROptions.SliderST.slider
- ScenarioGroupSetOpacityWindowSlider
- ShiniesConfigGeneralUIScaleSlider
- SliderWindowTemplateSlider
- StatdollOptionsBackgroundAlphaSlider
- StatdollOptionsScaleFactorSlider
- TAtlasSlider
- TChatTabTextEntryTemplateBackgroundAlphaSlider
- TChatTabTextEntryTemplateChannelAlphaSlider
- TChatTabWindowsGroupTemplateScrollbarFadeinAlphaSlider
- TChatTabWindowsGroupTemplateScrollbarFadeoutAlphaSlider
- TSN_SliderRowSlider
- TastyButtonsOptionsWindowCreateViewScaleSlider
- TastyButtonsOptionsWindowEditViewScaleSlider
- TastyButtonsOptionsWindowEditViewSliderColorB
- TastyButtonsOptionsWindowEditViewSliderColorG
- TastyButtonsOptionsWindowEditViewSliderColorR
- TexturedButtonsSetupCooldownWindowCooldownAnimationAlphaSlider
- TexturedButtonsSetupSelectColorWindowTintBlueSlider
- TexturedButtonsSetupSelectColorWindowTintGreenSlider
- TexturedButtonsSetupSelectColorWindowTintRedSlider
- TexturedButtonsSetupTintWindowTintBlueSlider
- TexturedButtonsSetupTintWindowTintGreenSlider
- TexturedButtonsSetupTintWindowTintRedSlider
- TurretRangeSetupDisplayWindowDistanceScaleSlider
- TurretRangeSetupDisplayWindowSelectColorTintBlueSlider
- TurretRangeSetupDisplayWindowSelectColorTintGreenSlider
- TurretRangeSetupDisplayWindowSelectColorTintRedSlider
- TurretRangeSetupDistanceWindowColorTintBlueSlider
- TurretRangeSetupDistanceWindowColorTintGreenSlider
- TurretRangeSetupDistanceWindowColorTintRedSlider
- TurretRangeSetupGeneralWindowUpdateDelaySlider
- WSCTOptionsColorPickerWindowCustomColorBlueSlider
- WSCTOptionsColorPickerWindowCustomColorGreenSlider
- WSCTOptionsColorPickerWindowCustomColorRedSlider
- WSCTSliderTemplateSlider
- WarBoard_FPSOptions_SliderB
- WarBoard_FPSOptions_SliderG
- WarBoard_FPSOptions_SliderR
- XpStatusSetOpacityWindowSlider
- alertModTemplateSlider

## Related APIs

- [SliderBar](../../xml/element_types/element_SliderBar.md) (HIGH 100/100) - XML Element Type

## Notes

- none
