# EA_Window_DefaultBackgroundFrame

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
| Addons seen in | AggroMeter, Aura, BankArkel, BuffHead, DaemonAssist, EA_UiDebugTools, EA_UiModWindow, FastInteract |
| Files seen in | `/workspace_addons/AggroMeter/AggroMeter.xml:94`, `/workspace_addons/Aura/Source/AuraColorPicker.xml:25`, `/workspace_addons/Aura/Source/AuraConfig.xml:33`, `/workspace_addons/Aura/Source/AuraSettings.xml:83`, `/workspace_addons/Aura/Source/AuraShares.xml:158`, `/workspace_addons/Aura/Source/AuraShares.xml:27`, `/workspace_addons/Aura/Source/AuraTexture.xml:104`, `/workspace_addons/BankArkel/BankArkel.xml:302` |
| Namespaces detected | EA_Window_DefaultBackgroundFrame |
| Source kinds | xml_attributes |
| Example locations | AggroMeterGrayWindowBackground, AuraColorPickerBackground, AuraConfigBackground, AuraSettingsBackground, AuraSharesBackground, AuraSharesImportExportBackground |
| XML usage count | 67 |
| XML attribute usage count | 67 |
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

Observed engine XML template or inherited constant referenced by 22 addons.

## Seen In

- AggroMeter
- Aura
- BankArkel
- BuffHead
- DaemonAssist
- EA_UiDebugTools
- EA_UiModWindow
- FastInteract
- LibGroup
- MapMonster
- MapPin
- ObjectInspector
- Pocket Palette
- QuickTacticSwitch
- RVMOD_Manager
- Shinies
- TexturedButtons
- TidyChat
- TurretRange
- WarBoard
- bigger_MacroWindow
- wbLeadHelper

## Used By

- AggroMeterGrayWindowBackground
- AuraColorPickerBackground
- AuraConfigBackground
- AuraSettingsBackground
- AuraSharesBackground
- AuraSharesImportExportBackground
- AuraTextureBackground
- BankArkelBackpackComboFrame
- BuffHeadSetupAdvancedCompressionItemEffectWindowBackground
- BuffHeadSetupAdvancedCompressionItemWindowBackground
- BuffHeadSetupAdvancedCompressionWindowBackground
- BuffHeadSetupAdvancedContainersItemPropertiesWindowBackground
- BuffHeadSetupAdvancedContainersItemWindowBackground
- BuffHeadSetupAdvancedContainersWindowBackground
- BuffHeadSetupContainerWindowBackground
- BuffHeadSetupDisplayWindowBackground
- BuffHeadSetupEffectCacheTableWindowBackground
- BuffHeadSetupEffectCacheWindowBackground
- BuffHeadSetupFilterWindowBackground
- BuffHeadSetupGeneralWindowBackground
- BuffHeadSetupLayoutManagerWindowBackground
- BuffHeadSetupLayoutPropertiesWindowBackground
- BuffHeadSetupLayoutWindowBackground
- BuffHeadSetupMenuWindowBackground
- BuffHeadSetupPerformanceWindowBackground
- BuffHeadSetupPriorityEffectsItemWindowBackground
- BuffHeadSetupPriorityEffectsWindowBackground
- BuffHeadSetupTrackersWindowBackground
- DaemonAssistWindowBackground
- DyeWindowBackground
- EA_Window_MacroBackground
- FastInteractWindowBackground
- ItemWindowBackground
- LibGroupSetupWindowBackground
- MacroIconSelectionWindowBackground
- MapMonster_CalibrateWindowMain
- MapMonster_EditorWindowBackground
- MapMonster_IconChooserWindowBackground
- MapMonster_PinTypeEditorWindowBackground
- MapPin_SetupBackground
- ObjectInspectorBackground
- PackWinBackground
- QuickTacticSwitchWindowBackground
- RVMOD_ManagerWindowBackground
- ShiniesBrowseUI_SearchesBackground
- ShiniesWindowBackground
- TexturedButtonsSetupActionbarWindowBackground
- TexturedButtonsSetupAdvancedTexturesWindowBackground
- TexturedButtonsSetupCooldownWindowBackground
- TexturedButtonsSetupFontsWindowBackground
- TexturedButtonsSetupMenuWindowBackground
- TexturedButtonsSetupMiscWindowBackground
- TexturedButtonsSetupTexturesWindowBackground
- TexturedButtonsSetupTintWindowBackground
- TidyChatLootRollBackground
- TurretRangeBoxBackground
- TurretRangeSetupDisplayWindowBackground
- TurretRangeSetupDistanceWindowBackground
- TurretRangeSetupDistancesWindowBackground
- TurretRangeSetupGeneralWindowBackground
- TurretRangeSetupMenuWindowBackground
- UiModAdvancedWindowBackground
- UiModVersionMismatchWindowBackground
- UiModWindowBackground
- WarBoardAlignOptionsWindowBackground
- WarBoardOptionsBackground
- wbLeadHelperConfigWindowBackground

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
