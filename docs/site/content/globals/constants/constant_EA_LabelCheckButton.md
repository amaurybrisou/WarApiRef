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
| Addons seen in | BuffHead, EA_UiDebugTools, EA_UiModWindow, FastInteract, LibGroup, Queue Queuer, TexturedButtons, TurretRange |
| Files seen in | `/workspace/BuffHead/Setup/SetupAdvancedContainersItemProperties.xml:1015`, `/workspace/BuffHead/Setup/SetupAdvancedContainersItemProperties.xml:319`, `/workspace/BuffHead/Setup/SetupAdvancedContainersItemProperties.xml:344`, `/workspace/BuffHead/Setup/SetupAdvancedContainersItemProperties.xml:990`, `/workspace/BuffHead/Setup/SetupDisplay.xml:199`, `/workspace/BuffHead/Setup/SetupEffectCache.xml:100`, `/workspace/BuffHead/Setup/SetupFilter.xml:20`, `/workspace/BuffHead/Setup/SetupLayout.xml:341` |
| Namespaces detected | EA_LabelCheckButton |
| Source kinds | xml_attributes |
| Example locations | BlacklistCheckBox, BuffHeadSetupAdvancedContainersItemPropertiesWindowElementEffectsAlwaysShowEnable, BuffHeadSetupAdvancedContainersItemPropertiesWindowElementEffectsPermanent, BuffHeadSetupAdvancedContainersItemPropertiesWindowElementHandleInputEnableRemovable, BuffHeadSetupAdvancedContainersItemPropertiesWindowElementHandleInputShowTooltips, BuffHeadSetupDisplayWindowEnableSort |
| XML usage count | 73 |
| XML attribute usage count | 73 |
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
- EA_UiDebugTools
- EA_UiModWindow
- FastInteract
- LibGroup
- Queue Queuer
- TexturedButtons
- TurretRange
- WSCT
- WarBoard

## Used By

- BlacklistCheckBox
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
- EA_LabelCheckButtonSmallCopy
- FastInteractWindowCheckHealer
- FastInteractWindowCheckQuestAccept
- FastInteractWindowCheckQuestComplete
- FastInteractWindowCheckRalleyMaster
- FastInteractWindowCheckVendorJD
- FastInteractWindowCheckVendorSkip
- LibGroupSetupWindowEnableGroupDistance
- OptionCheckBox
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
- UiModAdvancedWindowShowInAddOnsListCheck
- WSCTCheckBox
- WarBoardButtonCheckBox
- WarBoardButtonCheckBoxDefault
- WarBoardButtonCheckBoxEnabled

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
