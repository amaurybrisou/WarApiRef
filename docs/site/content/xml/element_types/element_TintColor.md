# TintColor

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
| Addons seen in | Moth |
| Files seen in | `/workspace/data/raw/Moth/Moth.xml:128`, `/workspace/data/raw/Moth/Moth.xml:157`, `/workspace/data/raw/Moth/Moth.xml:18`, `/workspace/data/raw/Moth/Moth.xml:30` |
| Namespaces detected | TintColor |
| Source kinds | xml_frames |
| Example locations | Moth: MothBackground, Moth: MothBordertronned, Moth: MothCellTemplateBackground, Moth: MothRowTemplateBackground |
| XML usage count | 4 |
| XML attribute usage count | 4 |
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

- b
- g
- r

## Common Inherits

- none

## Common Parent Elements

- [FullResizeImage](element_FullResizeImage.md) — 4× (HIGH)

## Attribute Reference

| Attribute | Required | Usage % | Sample Values |
| --- | --- | --- | --- |
| `b` | **required** | 100% | 0, 36, 166 |
| `g` | **required** | 100% | 0, 28, 84 |
| `r` | **required** | 100% | 0, 237 |
## Seen In

- Moth

## Examples

- Moth: MothBackground -> TintColor in FullResizeImage MothBackground
- Moth: MothBordertronned -> TintColor in FullResizeImage MothBordertronned
- Moth: MothCellTemplateBackground -> TintColor in FullResizeImage MothCellTemplateBackground
- Moth: MothRowTemplateBackground -> TintColor in FullResizeImage MothRowTemplateBackground

## Related APIs

- none

## Used With

- none

## Triggered By

- none

## Affects

- none
