# CircleImage

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
| Files seen in | `/workspace/data/raw/PartyCast/PartyCast.xml:481`, `/workspace/data/raw/PartyCast/PartyCast.xml:91` |
| Namespaces detected | CircleImage |
| Source kinds | xml_frames |
| Example locations | PartyCast: PartyCastWindow_TemplateButtonIcon, PartyCast: PartyCastWindow_Template_SpecialButtonIcon |
| XML usage count | 2 |
| XML attribute usage count | 2 |
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

- handleinput
- layer
- name
- numsegments
- popable
- sticky
- textureScale

## Common Inherits

- none

## Common Parent Elements

- [Button](element_Button.md)


> **Note**: This element type commonly acts as a template base.

## Attribute Reference

| Attribute | Required | Usage % | Sample Values |
| --- | --- | --- | --- |
| `handleinput` | **required** | 100% | false |
| `layer` | **required** | 100% | overlay, default |
| `numsegments` | **required** | 100% | 32, 16 |
| `popable` | **required** | 100% | false |
| `sticky` | **required** | 100% | true |
| `textureScale` | **required** | 100% | 0.36, 0.53 |
## Seen In

- PartyCast

## Examples

- PartyCast: PartyCastWindow_TemplateButtonIcon -> CircleImage PartyCastWindow_TemplateButtonIcon
- PartyCast: PartyCastWindow_Template_SpecialButtonIcon -> CircleImage PartyCastWindow_Template_SpecialButtonIcon

## Related APIs

- [Button](element_Button.md) (HIGH 100/100) - XML Element Type

## Used With

- none

## Triggered By

- none

## Affects

- none
