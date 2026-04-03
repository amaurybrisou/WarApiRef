# HorizontalResizeImage

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
| Addons seen in | PartyCast |
| Files seen in | `/workspace/data/raw/PartyCast/PartyCast.xml:9` |
| Namespaces detected | HorizontalResizeImage |
| Source kinds | xml_frames |
| Example locations | PartyCast: LOL_BAR |
| XML usage count | 1 |
| XML attribute usage count | 1 |
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

- name
- texture
- textureScale

## Common Inherits

- none

## Common Structural Child Elements

- [Left](element_Left.md)
- [Middle](element_Middle.md)
- [Right](element_Right.md)

## Attribute Reference

| Attribute | Required | Usage % | Sample Values |
| --- | --- | --- | --- |
| `texture` | **required** | 100% | EA_Training_Specialization |
| `textureScale` | **required** | 100% | 0.89 |
## Seen In

- PartyCast

## Examples

- PartyCast: LOL_BAR -> HorizontalResizeImage LOL_BAR

## Related APIs

- none

## Used With

- none

## Triggered By

- none

## Affects

- none
