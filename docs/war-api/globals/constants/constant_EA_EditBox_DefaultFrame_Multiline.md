# EA_EditBox_DefaultFrame_Multiline

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
| Addons seen in | AutoBand, BuffHead, Busted, EA_UiDebugTools, Enemy, LibAddonButton, MapMonster, MapPin |
| Files seen in | AutoBandWindow.xml, Busted.xml, Code/Core/Common.xml, Manager/CustomItem.xml, ObjectInspector.xml, Setup/SetupEffectCacheTable.xml, Setup/SetupLayoutManager.xml, Source/DebugWindow.xml |
| Namespaces detected | EA_EditBox_DefaultFrame_Multiline |
| Source kinds | xml_attributes |
| Example locations | AutoBandCopyLinkWindowUrlInput, BuffHeadSetupEffectCacheTableWindowTableEditBox, BuffHeadSetupLayoutManagerWindowElementExportEditBox, BuffHeadSetupLayoutManagerWindowElementImportEditBox, BustedGUIErrorMessage, DevPadWindowDevPadCode |
| XML usage count | 19 |
| XML attribute usage count | 19 |
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

Engine-supplied XML constant or template class referenced by 14 addons.

## Seen In

- AutoBand
- BuffHead
- Busted
- EA_UiDebugTools
- Enemy
- LibAddonButton
- MapMonster
- MapPin
- ObjectInspector
- Squared
- TastyButtons
- TidyChat
- bigger_MacroWindow
- zMailMod

## Used By

- AutoBandCopyLinkWindowUrlInput
- BuffHeadSetupEffectCacheTableWindowTableEditBox
- BuffHeadSetupLayoutManagerWindowElementExportEditBox
- BuffHeadSetupLayoutManagerWindowElementImportEditBox
- BustedGUIErrorMessage
- DevPadWindowDevPadCode
- EA_Window_MacroDetailsText
- EnemyTextEntryDialogValue
- LibAddonButtonManagerCustomItemWindowMacroEditBox
- LibAddonButtonManagerCustomItemWindowMacroTooltipEditBox
- MapMonster_EditorWindowNoteEditBox
- MapPin_SetupDescBox
- MapPin_SetupTextBox
- ObjectInspectorObjectEditBox
- SquaredImportExportFrameData
- TastyButtonsButtonSelectWindowButtonRangeViewEditComplex
- TidyChatCopyLog
- zMailModMassMailMessageBodyEditBox
- zMailModSendMessageEditBoxDefault

## Related APIs

- [EditBox](../../xml/element_types/element_EditBox.md) (HIGH 100/100) - XML Element Type

## Notes

- none
