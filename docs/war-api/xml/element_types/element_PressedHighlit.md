# PressedHighlit

- Category: XML Element Type
- Confidence level: MEDIUM
- Confidence score: 45/100

## Confidence Assessment

- Level: MEDIUM

- Score: 45/100

- Rationale: Promoted as MEDIUM confidence because matches default ui or extracted base ui surface, used directly in xml handler attributes.

## Evidence Signals

- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +30 Used directly in XML handler attributes: XML exposure suggests an engine-level contract.
- -20 Only one weak usage site: Evidence is too shallow to trust as platform API.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | PartyCast |
| Files seen in | `/workspace/data/raw/PartyCast/PartyCast.xml:186`, `/workspace/data/raw/PartyCast/PartyCast.xml:327`, `/workspace/data/raw/PartyCast/PartyCast.xml:469`, `/workspace/data/raw/PartyCast/PartyCast.xml:631`, `/workspace/data/raw/PartyCast/PartyCast.xml:80` |
| Namespaces detected | PressedHighlit |
| Source kinds | xml_frames |
| Example locations | PartyCast: PartyCastWindow_TemplateButton, PartyCast: PartyCastWindow_Template_LargeButton, PartyCast: PartyCastWindow_Template_PlainButton, PartyCast: PartyCastWindow_Template_SmallButton, PartyCast: PartyCastWindow_Template_SpecialButton |
| XML usage count | 5 |
| XML attribute usage count | 5 |
| Lua usage count | 0 |
| Global usage count | 0 |
| Local definition count | 0 |
| Documentation references | 0 |
| Initialization flow references | 0 |
| Known engine namespace | no |
| Default UI presence | yes |
| Event binding presence | no |
| Observed in XML and Lua | no |
| Consistent role | no |
| Consistent arguments | no |
| Consistent returns | no |
| Slash command presence | no |
| Weak usage only | yes |
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

- id

## Common Inherits

- none

## Common Parent Elements

- [TexSlices](element_TexSlices.md) — 5× (HIGH)

## Attribute Reference

| Attribute | Required | Usage % | Sample Values |
| --- | --- | --- | --- |
| `id` | **required** | 100% | morale-white, ability-white, tactic-black |
## Seen In

- PartyCast

## Examples

- PartyCast: PartyCastWindow_TemplateButton -> PressedHighlit in Button PartyCastWindow_TemplateButton
- PartyCast: PartyCastWindow_Template_LargeButton -> PressedHighlit in Button PartyCastWindow_Template_LargeButton
- PartyCast: PartyCastWindow_Template_PlainButton -> PressedHighlit in Button PartyCastWindow_Template_PlainButton
- PartyCast: PartyCastWindow_Template_SmallButton -> PressedHighlit in Button PartyCastWindow_Template_SmallButton
- PartyCast: PartyCastWindow_Template_SpecialButton -> PressedHighlit in Button PartyCastWindow_Template_SpecialButton

## Related APIs

- [TexSlices](element_TexSlices.md) (MEDIUM 45/100) - XML Element Type

## Used With

- none

## Triggered By

- none

## Affects

- none
