# EditBox

- Category: XML Element Type
- Confidence level: HIGH
- Confidence score: 100/100

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 123

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, used directly in xml handler attributes, matches a known engine namespace.

## Evidence Signals

- +18 Seen in 2 to 3 addons: Cross-addon spread is present but limited.
- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +30 Used directly in XML handler attributes: XML exposure suggests an engine-level contract.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | TidyChat, TidyRoll |
| Files seen in | `/workspace/data/raw/TidyChat/TidyChatCopy.xml:38`, `/workspace/data/raw/TidyRoll/CustomAutoRoll.xml:181`, `/workspace/data/raw/TidyRoll/CustomAutoRoll.xml:204` |
| Namespaces detected | EditBox |
| Source kinds | xml_frames |
| Example locations | TidyChat: TidyChatCopyLog, TidyRoll: TRollAutoRollAddByIdIdEditBox, TidyRoll: TRollAutoRollAddByIdNameEditBox |
| XML usage count | 3 |
| XML attribute usage count | 3 |
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

Observed XML element type instantiated by 2 addons.

## Common Attributes

- inherits
- name
- font
- maxchars

## Common Inherits

- EA_EditBox_DefaultFrame
- EA_EditBox_DefaultFrame_Multiline

## Common Parent Elements

- [Window](element_Window.md)

## Common Template Bases

- EA_EditBox_DefaultFrame
- EA_EditBox_DefaultFrame_Multiline

## Attribute Reference

| Attribute | Required | Usage % | Sample Values |
| --- | --- | --- | --- |
| `inherits` | **required** | 100% | EA_EditBox_DefaultFrame_Multiline, EA_EditBox_DefaultFrame |
| `font` | optional | 33% | font_clear_small_bold |
| `maxchars` | optional | 33% | 10240 |
## Seen In

- TidyChat
- TidyRoll

## Examples

- TidyChat: TidyChatCopyLog -> EditBox TidyChatCopyLog
- TidyRoll: TRollAutoRollAddByIdIdEditBox -> EditBox TRollAutoRollAddByIdIdEditBox
- TidyRoll: TRollAutoRollAddByIdNameEditBox -> EditBox TRollAutoRollAddByIdNameEditBox

## Related APIs

- none

## Used With

- none

## Triggered By

- none

## Affects

- none
