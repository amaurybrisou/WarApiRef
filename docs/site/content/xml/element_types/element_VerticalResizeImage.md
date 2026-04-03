# VerticalResizeImage

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
| Addons seen in | BuffHead |
| Files seen in | `/workspace/data/raw/BuffHead/Setup/SetupLayout.xml:0` |
| Namespaces detected | VerticalResizeImage |
| Source kinds | xml_frames |
| Example locations | BuffHead: BuffHeadLayoutVerticalButtonHighlight, BuffHead: BuffHeadLayoutVerticalButtonNormal, BuffHead: BuffHeadLayoutVerticalButtonPressed, BuffHead: BuffHeadLayoutVerticalResizeImage |
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
- texture

## Common Inherits

- BuffHeadLayoutVerticalResizeImage

## Common Parent Elements

- [Windows](element_Windows.md) — 4× (HIGH)

## Common Structural Child Elements

- [TintColor](element_TintColor.md) — 3× (HIGH)
- [Sizes](element_Sizes.md) — 1× (LOW)
- [TexCoords](element_TexCoords.md) — 1× (LOW)

## Common Template Bases

- BuffHeadLayoutVerticalResizeImage

## Typical XML Structure

```xml
<VerticalResizeImage name="..." texture="EA_TintableImage">
  <Sizes bottom="0" middle="30" top="0"/>
  <TexCoords>
    <Middle x="424" y="762"/>
  </TexCoords>
</VerticalResizeImage>
```

## Attribute Reference

| Attribute | Required | Usage % | Sample Values |
| --- | --- | --- | --- |
| `inherits` | optional | 75% | BuffHeadLayoutVerticalResizeImage |
| `texture` | optional | 25% | EA_TintableImage |
## Structural Sub-Elements

### [TintColor](element_TintColor.md)

Observed 3 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `b` | **required** | 0, 110, 130, 50 |
| `g` | **required** | 0, 200, 185, 130 |
| `r` | **required** | 255, 0, 180, 200 |
| `a` | optional | 30, 255, 0.5, 200 |
### [Sizes](element_Sizes.md)

Observed 1 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `middle` | optional | 30, 200, 61, 50 |
| `left` | optional | 0, 7 |
| `right` | optional | 0 |
| `bottom` | optional | 0 |
| `top` | optional | 0 |
### [TexCoords](element_TexCoords.md)

Observed 1 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `x` | optional | 0, 133, 88, 32 |
| `y` | optional | 0, 163, 51, 32 |
## Recursive Hierarchy

- Root: [VerticalResizeImage](element_VerticalResizeImage.md)
- [Sizes](element_Sizes.md) (structural, 1×, LOW)
  - [BottomRight](element_BottomRight.md) (structural, 14×, HIGH)
  - [Middle](element_Middle.md) (structural, 14×, HIGH)
  - [TopLeft](element_TopLeft.md) (structural, 14×, HIGH)
- [TexCoords](element_TexCoords.md) (structural, 1×, LOW)
  - [BottomCenter](element_BottomCenter.md) (structural, 9×, MEDIUM)
  - [BottomLeft](element_BottomLeft.md) (structural, 9×, MEDIUM)
  - [BottomRight](element_BottomRight.md) (structural, 9×, MEDIUM)
  - [Disabled](element_Disabled.md) (structural, 7×, MEDIUM)
  - [Left](element_Left.md) (structural, 7×, MEDIUM)
  - [Middle](element_Middle.md) (structural, 9×, MEDIUM)
  - [MiddleCenter](element_MiddleCenter.md) (structural, 9×, MEDIUM)
  - [MiddleLeft](element_MiddleLeft.md) (structural, 9×, MEDIUM)
  - [MiddleRight](element_MiddleRight.md) (structural, 9×, MEDIUM)
  - [Normal](element_Normal.md) (structural, 7×, MEDIUM)
  - [NormalHighlit](element_NormalHighlit.md) (structural, 7×, MEDIUM)
  - [Pressed](element_Pressed.md) (structural, 7×, MEDIUM)
  - [PressedHighlit](element_PressedHighlit.md) (structural, 1×, LOW)
  - [Right](element_Right.md) (structural, 7×, MEDIUM)
  - [TopCenter](element_TopCenter.md) (structural, 9×, MEDIUM)
  - [TopLeft](element_TopLeft.md) (structural, 9×, MEDIUM)
  - [TopRight](element_TopRight.md) (structural, 9×, MEDIUM)
- [TintColor](element_TintColor.md) (structural, 3×, HIGH)

## Seen In

- BuffHead

## Examples

- BuffHead: BuffHeadLayoutVerticalButtonHighlight -> VerticalResizeImage BuffHeadLayoutVerticalButtonHighlight
- BuffHead: BuffHeadLayoutVerticalButtonNormal -> VerticalResizeImage BuffHeadLayoutVerticalButtonNormal
- BuffHead: BuffHeadLayoutVerticalButtonPressed -> VerticalResizeImage BuffHeadLayoutVerticalButtonPressed
- BuffHead: BuffHeadLayoutVerticalResizeImage -> VerticalResizeImage BuffHeadLayoutVerticalResizeImage

## Related APIs

- [Sizes](element_Sizes.md) (HIGH 100/100) - XML Element Type
- [TexCoords](element_TexCoords.md) (HIGH 100/100) - XML Element Type
- [TintColor](element_TintColor.md) (HIGH 100/100) - XML Element Type
- [Windows](element_Windows.md) (HIGH 100/100) - XML Element Type

## Used With

- none

## Triggered By

- none

## Affects

- none
