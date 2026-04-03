# Sizes

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
| Addons seen in | CM_ClosetGoblin |
| Namespaces detected | Sizes |
| Source kinds | xml_frames |
| Example locations | CM_ClosetGoblin: ClosetGoblinDefaultBG |
| XML usage count | 1 |
| XML attribute usage count | 1 |
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

Sizes is a structural XML sub-element. It commonly appears under FullResizeImage. It is typically used to organize structural children such as BottomRight, Middle, TopLeft.

## Common Attributes

- none

## Common Inherits

- none

## Common Parent Elements

- [FullResizeImage](element_FullResizeImage.md) — 1× (HIGH)

## Common Structural Child Elements

- [BottomRight](element_BottomRight.md) — 1× (HIGH)
- [Middle](element_Middle.md) — 1× (HIGH)
- [TopLeft](element_TopLeft.md) — 1× (HIGH)

## Structural Sub-Elements

### [BottomRight](element_BottomRight.md)

Observed 1 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `x` | **required** | 16, 350 |
| `y` | **required** | 15, 197 |
### [Middle](element_Middle.md)

Observed 1 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `x` | **required** | 78 |
| `y` | **required** | 25 |
### [TopLeft](element_TopLeft.md)

Observed 1 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `x` | **required** | 16, 256 |
| `y` | **required** | 15, 157 |
## Recursive Hierarchy

- Root: [Sizes](element_Sizes.md)
- [BottomRight](element_BottomRight.md) (structural, 1×, HIGH)
- [Middle](element_Middle.md) (structural, 1×, HIGH)
- [TopLeft](element_TopLeft.md) (structural, 1×, HIGH)

## Seen In

- CM_ClosetGoblin

## Examples

- CM_ClosetGoblin: ClosetGoblinDefaultBG -> Sizes in FullResizeImage ClosetGoblinDefaultBG

## Related APIs

- [FullResizeImage](element_FullResizeImage.md) (HIGH 90/100) - XML Element Type
- [BottomRight](element_BottomRight.md) (LOW 5/100) - XML Element Type
- [Middle](element_Middle.md) (LOW 5/100) - XML Element Type
- [TopLeft](element_TopLeft.md) (LOW 5/100) - XML Element Type
