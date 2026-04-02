# SliderBar

- Category: XML Element Type
- Confidence level: HIGH
- Confidence score: 90/100

## Confidence Assessment

- Level: HIGH

- Score: 90/100

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, used directly in xml handler attributes, matches a known engine namespace.

## Evidence Signals

- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +30 Used directly in XML handler attributes: XML exposure suggests an engine-level contract.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | TidyChat |
| Files seen in | `/workspace/data/raw/TidyChat/TidyChatTabTextEntry.xml:128`, `/workspace/data/raw/TidyChat/TidyChatTabTextEntry.xml:87`, `/workspace/data/raw/TidyChat/TidyChatTabWindows.xml:144`, `/workspace/data/raw/TidyChat/TidyChatTabWindows.xml:174` |
| Namespaces detected | SliderBar |
| Source kinds | xml_frames |
| Example locations | TidyChat: TChatTabTextEntryTemplateBackgroundAlphaSlider, TidyChat: TChatTabTextEntryTemplateChannelAlphaSlider, TidyChat: TChatTabWindowsGroupTemplateScrollbarFadeinAlphaSlider, TidyChat: TChatTabWindowsGroupTemplateScrollbarFadeoutAlphaSlider |
| XML usage count | 4 |
| XML attribute usage count | 4 |
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

Observed XML element type instantiated by 1 addons.

## Common Attributes

- inherits
- name

## Common Inherits

- EA_Default_SliderBar

## Common Parent Elements

- [Window](element_Window.md)

## Common Template Bases

- EA_Default_SliderBar


> **Note**: This element type commonly acts as a template base.

## Attribute Reference

| Attribute | Required | Usage % | Sample Values |
| --- | --- | --- | --- |
| `inherits` | **required** | 100% | EA_Default_SliderBar |
## Seen In

- TidyChat

## Examples

- TidyChat: TChatTabTextEntryTemplateBackgroundAlphaSlider -> SliderBar TChatTabTextEntryTemplateBackgroundAlphaSlider
- TidyChat: TChatTabTextEntryTemplateChannelAlphaSlider -> SliderBar TChatTabTextEntryTemplateChannelAlphaSlider
- TidyChat: TChatTabWindowsGroupTemplateScrollbarFadeinAlphaSlider -> SliderBar TChatTabWindowsGroupTemplateScrollbarFadeinAlphaSlider
- TidyChat: TChatTabWindowsGroupTemplateScrollbarFadeoutAlphaSlider -> SliderBar TChatTabWindowsGroupTemplateScrollbarFadeoutAlphaSlider

## Related APIs

- none

## Used With

- none

## Triggered By

- none

## Affects

- none
