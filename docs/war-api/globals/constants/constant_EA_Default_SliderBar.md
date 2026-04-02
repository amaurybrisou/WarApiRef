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
| Addons seen in | Aura, BuffHead, Enemy, LibGroup, MapMonster, MapPin, MoraleCircle, PotionBar |
| Files seen in | `/workspace_addons/Aura/Source/AuraColorPicker.xml:26`, `/workspace_addons/Aura/Source/AuraColorPicker.xml:58`, `/workspace_addons/Aura/Source/AuraColorPicker.xml:78`, `/workspace_addons/Aura/Source/AuraColorPicker.xml:98`, `/workspace_addons/BuffHead/Setup/General.xml:132`, `/workspace_addons/BuffHead/Setup/General.xml:171`, `/workspace_addons/BuffHead/Setup/General.xml:211`, `/workspace_addons/BuffHead/Setup/SetupAdvancedContainersItemProperties.xml:793` |
| Namespaces detected | EA_Default_SliderBar |
| Source kinds | xml_attributes |
| Example locations | AuraColorPickerAlpha, AuraColorPickerBlue, AuraColorPickerGreen, AuraColorPickerRed, BuffHeadSetupAdvancedContainersItemPropertiesWindowElementScaleSlider, BuffHeadSetupContainerWindowContainerColumnsSlider |
| XML usage count | 82 |
| XML attribute usage count | 82 |
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

Observed engine XML template or inherited constant referenced by 17 addons.

## Seen In

- Aura
- BuffHead
- Enemy
- LibGroup
- MapMonster
- MapPin
- MoraleCircle
- PotionBar
- RVAPI_ColorDialog
- RVMOD_Manager
- RoR_SoR
- Shinies
- TexturedButtons
- TidyChat
- TurretRange
- WSCT
- WarBoard

## Used By

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
- EnemyEffectsIndicatorDialogContentScrollChildAlpha
- EnemyEffectsIndicatorDialogContentScrollChildColorB
- EnemyEffectsIndicatorDialogContentScrollChildColorG
- EnemyEffectsIndicatorDialogContentScrollChildColorR
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
- PotionBarTypeTemplateOpacitySlider
- PotionBarTypeTemplateScaleSlider
- RVAPI_ColorDialogSliderTemplate
- RVMOD_ManagerSliderTemplate
- RoR_SoR_OffsetSlider
- RoR_SoR_OpacitySlider
- RoR_SoR_ScaleSlider
- ShiniesConfigGeneralUIScaleSlider
- SliderWindowTemplateSlider
- TChatTabTextEntryTemplateBackgroundAlphaSlider
- TChatTabTextEntryTemplateChannelAlphaSlider
- TChatTabWindowsGroupTemplateScrollbarFadeinAlphaSlider
- TChatTabWindowsGroupTemplateScrollbarFadeoutAlphaSlider
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
