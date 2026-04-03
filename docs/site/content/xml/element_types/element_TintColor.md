# TintColor

- Category: XML Element Type
- Confidence level: HIGH
- Confidence score: 100/100

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 110

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, seen in 4 or more addons, used directly in xml handler attributes.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +30 Used directly in XML handler attributes: XML exposure suggests an engine-level contract.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | InfoScroller, Moth, PartyCast, Soloq, minesweep |
| Files seen in | `/workspace/data/raw/InfoScroller/InfoScroller.xml:236`, `/workspace/data/raw/InfoScroller/InfoScroller.xml:42`, `/workspace/data/raw/Moth/Moth.xml:128`, `/workspace/data/raw/Moth/Moth.xml:157`, `/workspace/data/raw/Moth/Moth.xml:18`, `/workspace/data/raw/Moth/Moth.xml:30`, `/workspace/data/raw/PartyCast/PartyCast.xml:305`, `/workspace/data/raw/PartyCast/PartyCast.xml:348` |
| Namespaces detected | TintColor |
| Source kinds | xml_frames |
| Example locations | InfoScroller: InfoScrollerTemplateBackGroundBG, InfoScroller: InfoScrollerTemplateSeperatorBG, Moth: MothBackground, Moth: MothBordertronned, Moth: MothCellTemplateBackground, Moth: MothRowTemplateBackground |
| XML usage count | 12 |
| XML attribute usage count | 12 |
| Lua usage count | 0 |
| Global usage count | 0 |
| Local definition count | 0 |
| Documentation references | 0 |
| Initialization flow references | 0 |
| Known engine namespace | no |
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

Observed XML element type instantiated by 5 addons.

## Common Attributes

- b
- g
- r
- a

## Common Inherits

- none

## Common Parent Elements

- [FullResizeImage](element_FullResizeImage.md) — 10× (HIGH)
- [DynamicImage](element_DynamicImage.md) — 2× (MEDIUM)

## Attribute Reference

| Attribute | Required | Usage % | Sample Values |
| --- | --- | --- | --- |
| `b` | **required** | 100% | 0, 36, 166, 200, ... |
| `g` | **required** | 100% | 0, 28, 84, 200, ... |
| `r` | **required** | 100% | 0, 237, 200, 55 |
| `a` | optional | 50% | 1, 0.4, 0.5, 0.8 |
## Seen In

- InfoScroller
- Moth
- PartyCast
- Soloq
- minesweep

## Examples

- InfoScroller: InfoScrollerTemplateBackGroundBG -> TintColor in FullResizeImage InfoScrollerTemplateBackGroundBG
- InfoScroller: InfoScrollerTemplateSeperatorBG -> TintColor in DynamicImage InfoScrollerTemplateSeperatorBG
- Moth: MothBackground -> TintColor in FullResizeImage MothBackground
- Moth: MothBordertronned -> TintColor in FullResizeImage MothBordertronned
- Moth: MothCellTemplateBackground -> TintColor in FullResizeImage MothCellTemplateBackground
- Moth: MothRowTemplateBackground -> TintColor in FullResizeImage MothRowTemplateBackground

## Related APIs

- [DynamicImage](element_DynamicImage.md) (HIGH 100/100) - XML Element Type
- [FullResizeImage](element_FullResizeImage.md) (HIGH 100/100) - XML Element Type

## Used With

- none

## Triggered By

- none

## Affects

- none
