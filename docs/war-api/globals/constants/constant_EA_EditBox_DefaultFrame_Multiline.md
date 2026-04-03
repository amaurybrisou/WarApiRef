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
| Addons seen in | BuffHead, Enemy, TidyChat, bigger_MacroWindow |
| Files seen in | `/workspace/data/raw/BuffHead/Setup/SetupEffectCacheTable.xml:0`, `/workspace/data/raw/BuffHead/Setup/SetupLayoutManager.xml:0`, `/workspace/data/raw/Enemy/Code/Core/Common.xml:0`, `/workspace/data/raw/TidyChat/TidyChatCopy.xml:0`, `/workspace/data/raw/bigger_macrowindow/Source/MacroWindow.xml:0` |
| Namespaces detected | EA_EditBox_DefaultFrame_Multiline |
| Source kinds | xml_attributes |
| Example locations | BuffHeadSetupEffectCacheTableWindowTableEditBox, BuffHeadSetupLayoutManagerWindowElementExportEditBox, BuffHeadSetupLayoutManagerWindowElementImportEditBox, EA_Window_MacroDetailsText, EnemyTextEntryDialogValue, TidyChatCopyLog |
| XML usage count | 6 |
| XML attribute usage count | 6 |
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

Observed engine XML template or inherited constant referenced by 4 addons.

## Seen In

- BuffHead
- Enemy
- TidyChat
- bigger_MacroWindow

## Used By

- BuffHeadSetupEffectCacheTableWindowTableEditBox
- BuffHeadSetupLayoutManagerWindowElementExportEditBox
- BuffHeadSetupLayoutManagerWindowElementImportEditBox
- EA_Window_MacroDetailsText
- EnemyTextEntryDialogValue
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
