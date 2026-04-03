# EA_LabelCheckButton

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
| Addons seen in | BuffHead, DAoCBuff, LibGroup, TexturedButtons, TurretRange, WSCT, WarBoard |
| Files seen in | `/workspace/data/raw/BuffHead/Setup/SetupAdvancedContainersItemProperties.xml:0`, `/workspace/data/raw/BuffHead/Setup/SetupDisplay.xml:0`, `/workspace/data/raw/BuffHead/Setup/SetupEffectCache.xml:0`, `/workspace/data/raw/BuffHead/Setup/SetupFilter.xml:0`, `/workspace/data/raw/BuffHead/Setup/SetupLayout.xml:0`, `/workspace/data/raw/BuffHead/Setup/SetupLayoutManager.xml:0`, `/workspace/data/raw/BuffHead/Setup/SetupLayoutProperties.xml:0`, `/workspace/data/raw/BuffHead/Setup/SetupPerformance.xml:0` |
| Namespaces detected | EA_LabelCheckButton |
| Source kinds | xml_attributes |
| Example locations | BuffHeadSetupAdvancedContainersItemPropertiesWindowElementEffectsAlwaysShowEnable, BuffHeadSetupAdvancedContainersItemPropertiesWindowElementEffectsPermanent, BuffHeadSetupAdvancedContainersItemPropertiesWindowElementHandleInputEnableRemovable, BuffHeadSetupAdvancedContainersItemPropertiesWindowElementHandleInputShowTooltips, BuffHeadSetupDisplayWindowEnableSort, BuffHeadSetupEffectCacheWindowEnableCaching |
| XML usage count | 86 |
| XML attribute usage count | 86 |
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

Observed engine XML template or inherited constant referenced by 7 addons.

## Seen In

- BuffHead
- DAoCBuff
- LibGroup
- TexturedButtons
- TurretRange
- WSCT
- WarBoard

## Used By

- BuffHeadSetupAdvancedContainersItemPropertiesWindowElementEffectsAlwaysShowEnable
- BuffHeadSetupAdvancedContainersItemPropertiesWindowElementEffectsPermanent
- BuffHeadSetupAdvancedContainersItemPropertiesWindowElementHandleInputEnableRemovable
- BuffHeadSetupAdvancedContainersItemPropertiesWindowElementHandleInputShowTooltips
- BuffHeadSetupDisplayWindowEnableSort
- BuffHeadSetupEffectCacheWindowEnableCaching
- BuffHeadSetupFilterRowTemplateEnable
- BuffHeadSetupLayoutManagerWindowExport
- BuffHeadSetupLayoutManagerWindowImport
- BuffHeadSetupLayoutManagerWindowLayouts
- BuffHeadSetupLayoutPropertiesWindowElementStatusBarEnable
- BuffHeadSetupLayoutPropertiesWindowElementStatusBarForegroundStretch
- BuffHeadSetupLayoutPropertiesWindowElementStatusBarOrientationReverse
- BuffHeadSetupLayoutWindowLayerWindowLayer0Checkbox
- BuffHeadSetupLayoutWindowLayerWindowLayer1Checkbox
- BuffHeadSetupLayoutWindowLayerWindowLayer2Checkbox
- BuffHeadSetupLayoutWindowLayerWindowLayer3Checkbox
- BuffHeadSetupLayoutWindowLayerWindowLayer4Checkbox
- BuffHeadSetupPerformanceWindowEnableFading
- BuffHeadSetupPerformanceWindowEnablePriorityUpdates
- BuffHeadSetupPerformanceWindowEnableSync
- BuffHeadSetupPriorityEffectsItemWindowTargetTypeFriendly
- BuffHeadSetupPriorityEffectsItemWindowTargetTypeGroup
- BuffHeadSetupPriorityEffectsItemWindowTargetTypeHostile
- BuffHeadSetupPriorityEffectsItemWindowTargetTypeSelf
- BuffHeadSetupPriorityEffectsWindowSortFirst
- BuffHeadSetupTrackersWindowOnTargetChangeClearAlwaysShow
- BuffHeadSetupTrackersWindowOnTargetChangeClearBuffs
- BuffHeadSetupTrackersWindowOnTargetChangeClearDebuffs
- BuffHeadSetupTrackersWindowTrackerEnable
- BuffHeadSetupTrackersWindowTrackerPermanentEffects
- DAoCBuffFrameSettingsTab_ScrollChild_ActiveCheckBox
- DAoCBuffFrameSettingsTab_ScrollChild_AdvancedFiltersCheckBox
- DAoCBuffFrameSettingsTab_ScrollChild_BuffsBelowCheckBox
- DAoCBuffFrameSettingsTab_ScrollChild_GlassCheckBox
- DAoCBuffFrameSettingsTab_ScrollChild_HPTECheckBox
- DAoCBuffFrameSettingsTab_ScrollChild_HideLongtimeCheckBox
- DAoCBuffFrameSettingsTab_ScrollChild_LongtoPermaCheckBox
- DAoCBuffFrameSettingsTab_ScrollChild_SelfEffectsCheckBox
- DAoCBuffFrameSettingsTab_ScrollChild_ShowBorderCheckBox
- DAoCBuffFrameSettingsTab_ScrollChild_StaticCondenseCheckBox
- DAoCBuffFrameSettings_G4Filter_G4NotResultCheckBox
- DAoCBuffGeneralSettingsTab_ScrollChild_KillBuffsCheckBox
- DAoCBuffGeneralSettingsTab_ScrollChild_ToggleTestModeCheckBox
- DAoCBuffListManagerTab_ScrollChild_RightClickCopyCheckBox
- DAoCBuff_Settings_FilterFrame_ScrollChild_ColorBCheckBox
- DAoCBuff_Settings_FilterFrame_ScrollChild_ColorCCheckBox
- DAoCBuff_Settings_FilterFrame_ScrollChild_CombatHideCheckBox
- DAoCBuff_Settings_FilterFrame_ScrollChild_CondenseCheckBox
- DAoCBuff_Settings_FilterFrame_ScrollChild_DeleteCheckBox
- DAoCBuff_Settings_FilterFrame_ScrollChild_EnableClasstableCheckBox
- DAoCBuff_Settings_FilterFrame_ScrollChild_EnableFilterCheckBox
- DAoCBuff_Settings_FilterFrame_ScrollChild_EndaCheckBox
- DAoCBuff_Settings_FilterFrame_ScrollChild_InvertResultCheckBox
- LibGroupSetupWindowEnableGroupDistance
- TexturedButtonsSetupActionbarWindowEnable
- TexturedButtonsSetupActionbarWindowHideBackground
- TexturedButtonsSetupActionbarWindowHideEmpty
- TexturedButtonsSetupAdvancedTexturesWindowDisabled
- TexturedButtonsSetupAdvancedTexturesWindowUseCustom
- TexturedButtonsSetupAdvancedTexturesWindowUsePreset
- TexturedButtonsSetupCooldownWindowEnable
- TexturedButtonsSetupCooldownWindowEnableButtonTinting
- TexturedButtonsSetupCooldownWindowHideFlash
- TexturedButtonsSetupCooldownWindowRemoveCooldownS
- TexturedButtonsSetupCooldownWindowShowGlobalCooldownText
- TexturedButtonsSetupFontsWindowEnable
- TexturedButtonsSetupFontsWindowHideCooldown
- TexturedButtonsSetupFontsWindowHideKeybind
- TexturedButtonsSetupMiscWindowEnableCooldownPulse
- TexturedButtonsSetupMiscWindowHideActive
- TexturedButtonsSetupMiscWindowHideGlow
- TexturedButtonsSetupMiscWindowHideQuicklock
- TexturedButtonsSetupMiscWindowMovableQuicklock
- TexturedButtonsSetupMiscWindowSaveQuicklock
- TexturedButtonsSetupTexturesWindowDisabled
- TexturedButtonsSetupTexturesWindowUseCustom
- TexturedButtonsSetupTexturesWindowUsePreset
- TexturedButtonsSetupTintWindowEnable
- TurretRangeSetupDisplayWindowCircleInvert
- TurretRangeSetupDistanceWindowHide
- TurretRangeSetupGeneralWindowEnable
- WSCTCheckBox
- WarBoardButtonCheckBox
- WarBoardButtonCheckBoxDefault
- WarBoardButtonCheckBoxEnabled

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
