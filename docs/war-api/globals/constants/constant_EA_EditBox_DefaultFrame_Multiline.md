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
| Addons seen in | BuffHead, Busted, EA_UiDebugTools, Enemy, MapMonster, MapPin, ObjectInspector, TidyChat |
| Files seen in | `/workspace_addons/BuffHead/Setup/SetupEffectCacheTable.xml:58`, `/workspace_addons/BuffHead/Setup/SetupLayoutManager.xml:142`, `/workspace_addons/BuffHead/Setup/SetupLayoutManager.xml:203`, `/workspace_addons/Busted/Busted.xml:96`, `/workspace_addons/Enemy/Code/Core/Common.xml:128`, `/workspace_addons/MapMonster/Source/MapMonster_EditorWindow.xml:386`, `/workspace_addons/MapPin/source/MapPin.xml:588`, `/workspace_addons/MapPin/source/MapPin.xml:599` |
| Namespaces detected | EA_EditBox_DefaultFrame_Multiline |
| Source kinds | xml_attributes |
| Example locations | BuffHeadSetupEffectCacheTableWindowTableEditBox, BuffHeadSetupLayoutManagerWindowElementExportEditBox, BuffHeadSetupLayoutManagerWindowElementImportEditBox, BustedGUIErrorMessage, DevPadWindowDevPadCode, EA_Window_MacroDetailsText |
| XML usage count | 12 |
| XML attribute usage count | 12 |
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
- Busted
- EA_UiDebugTools
- Enemy
- MapMonster
- MapPin
- ObjectInspector
- TidyChat
- bigger_MacroWindow

## Used By

- BuffHeadSetupEffectCacheTableWindowTableEditBox
- BuffHeadSetupLayoutManagerWindowElementExportEditBox
- BuffHeadSetupLayoutManagerWindowElementImportEditBox
- BustedGUIErrorMessage
- DevPadWindowDevPadCode
- EA_Window_MacroDetailsText
- EnemyTextEntryDialogValue
- MapMonster_EditorWindowNoteEditBox
- MapPin_SetupDescBox
- MapPin_SetupTextBox
- ObjectInspectorObjectEditBox
- TidyChatCopyLog

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
