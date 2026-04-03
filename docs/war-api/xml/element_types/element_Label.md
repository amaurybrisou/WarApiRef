# Label

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
| Addons seen in | Moth, TidyChat, TidyRoll |
| Files seen in | `/workspace/data/raw/Moth/Moth.xml:173`, `/workspace/data/raw/TidyChat/TidyChat.xml:41`, `/workspace/data/raw/TidyChat/TidyChat.xml:52`, `/workspace/data/raw/TidyChat/TidyChatCopy.xml:65`, `/workspace/data/raw/TidyChat/TidyChatLootRoll.xml:16`, `/workspace/data/raw/TidyChat/TidyChatLootRoll.xml:27`, `/workspace/data/raw/TidyChat/TidyChatLootRoll.xml:88`, `/workspace/data/raw/TidyChat/TidyChatTabTextEntry.xml:117` |
| Namespaces detected | Label |
| Source kinds | xml_frames |
| Example locations | Moth: MothCellTemplateText, TidyChat: TChatCheckboxTemplateLabel, TidyChat: TChatLabel, TidyChat: TChatTabTextEntryTemplateAnchorPointLabel, TidyChat: TChatTabTextEntryTemplateBackgroundAlphaLabel, TidyChat: TChatTabTextEntryTemplateChannelAlphaLabel |
| XML usage count | 32 |
| XML attribute usage count | 32 |
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

Observed XML element type instantiated by 3 addons.

## Common Attributes

- name
- inherits
- font
- handleinput
- textalign
- autoresize
- wordwrap
- maxchars
- autoresizewidth
- sticky
- text

## Common Inherits

- TChatLabel
- TRollLabel
- EA_Label_DefaultText

## Common Parent Elements

- [Window](element_Window.md)

## Common Template Bases

- EA_Label_DefaultText
- TChatLabel
- TRollLabel


> **Note**: This element type commonly acts as a template base.

## Attribute Reference

| Attribute | Required | Usage % | Sample Values |
| --- | --- | --- | --- |
| `inherits` | optional | 75% | EA_Label_DefaultText, TChatLabel, TRollLabel |
| `font` | optional | 43% | font_clear_small_bold, font_clear_large_bold, font_clear_medium, font_clear_default |
| `handleinput` | optional | 31% | false |
| `textalign` | optional | 28% | left, center, right |
| `autoresize` | optional | 21% | true |
| `wordwrap` | optional | 18% | true |
| `maxchars` | optional | 6% | 1024, 256 |
| `autoresizewidth` | optional | 3% | true |
| `sticky` | optional | 3% | false |
| `text` | optional | 3% |  |
## Lua Functions Manipulating This Type

- TidyRollOptions.Initialize

## Seen In

- Moth
- TidyChat
- TidyRoll

## Examples

- Moth: MothCellTemplateText -> Label MothCellTemplateText
- TidyChat: TChatCheckboxTemplateLabel -> Label TChatCheckboxTemplateLabel
- TidyChat: TChatLabel -> Label TChatLabel
- TidyChat: TChatTabTextEntryTemplateAnchorPointLabel -> Label TChatTabTextEntryTemplateAnchorPointLabel
- TidyChat: TChatTabTextEntryTemplateBackgroundAlphaLabel -> Label TChatTabTextEntryTemplateBackgroundAlphaLabel
- TidyChat: TChatTabTextEntryTemplateChannelAlphaLabel -> Label TChatTabTextEntryTemplateChannelAlphaLabel

## Related APIs

- [Window](element_Window.md) (HIGH 100/100) - XML Element Type

## Used With

- none

## Triggered By

- none

## Affects

- none
