# TexCoords

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
| Namespaces detected | TexCoords |
| Source kinds | xml_frames |
| Example locations | CM_ClosetGoblin: ClosetGoblinDefaultBG, CM_ClosetGoblin: ClosetGoblinDefaultButton |
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

TexCoords is a structural XML sub-element. It commonly appears under Button and FullResizeImage. It is typically used to organize structural children such as BottomCenter, BottomLeft, BottomRight.

## Common Attributes

- none

## Common Inherits

- none

## Common Parent Elements

- [Button](element_Button.md) — 1× (HIGH)
- [FullResizeImage](element_FullResizeImage.md) — 1× (HIGH)

## Common Structural Child Elements

- [BottomCenter](element_BottomCenter.md) — 1× (HIGH)
- [BottomLeft](element_BottomLeft.md) — 1× (HIGH)
- [BottomRight](element_BottomRight.md) — 1× (HIGH)
- [Disabled](element_Disabled.md) — 1× (HIGH)
- [MiddleCenter](element_MiddleCenter.md) — 1× (HIGH)
- [MiddleLeft](element_MiddleLeft.md) — 1× (HIGH)
- [MiddleRight](element_MiddleRight.md) — 1× (HIGH)
- [Normal](element_Normal.md) — 1× (HIGH)
- [NormalHighlit](element_NormalHighlit.md) — 1× (HIGH)
- [Pressed](element_Pressed.md) — 1× (HIGH)
- [TopCenter](element_TopCenter.md) — 1× (HIGH)
- [TopLeft](element_TopLeft.md) — 1× (HIGH)
- [TopRight](element_TopRight.md) — 1× (HIGH)

## Structural Sub-Elements

### [BottomCenter](element_BottomCenter.md)

Observed 1 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `x` | **required** | 272 |
| `y` | **required** | 197 |
### [BottomLeft](element_BottomLeft.md)

Observed 1 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `x` | **required** | 256 |
| `y` | **required** | 197 |
### [BottomRight](element_BottomRight.md)

Observed 1 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `x` | **required** | 16, 350 |
| `y` | **required** | 15, 197 |
### [Disabled](element_Disabled.md)

Observed 1 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `texture` | **required** | EA_SquareFrame |
| `x` | **required** | 0 |
| `y` | **required** | 0 |
### [MiddleCenter](element_MiddleCenter.md)

Observed 1 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `x` | **required** | 272 |
| `y` | **required** | 172 |
### [MiddleLeft](element_MiddleLeft.md)

Observed 1 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `x` | **required** | 256 |
| `y` | **required** | 172 |
### [MiddleRight](element_MiddleRight.md)

Observed 1 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `x` | **required** | 350 |
| `y` | **required** | 172 |
### [Normal](element_Normal.md)

Observed 1 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `texture` | **required** | EA_SquareFrame |
| `x` | **required** | 0 |
| `y` | **required** | 0 |
### [NormalHighlit](element_NormalHighlit.md)

Observed 1 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `texture` | **required** | EA_SquareFrame_Highlight |
| `x` | **required** | 0 |
| `y` | **required** | 0 |
### [Pressed](element_Pressed.md)

Observed 1 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `texture` | **required** | EA_SquareFrame_Pressed |
| `x` | **required** | 0 |
| `y` | **required** | 0 |
### [TopCenter](element_TopCenter.md)

Observed 1 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `x` | **required** | 272 |
| `y` | **required** | 157 |
### [TopLeft](element_TopLeft.md)

Observed 1 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `x` | **required** | 16, 256 |
| `y` | **required** | 15, 157 |
### [TopRight](element_TopRight.md)

Observed 1 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `x` | **required** | 350 |
| `y` | **required** | 157 |
## Recursive Hierarchy

- Root: [TexCoords](element_TexCoords.md)
- [BottomCenter](element_BottomCenter.md) (structural, 1×, HIGH)
- [BottomLeft](element_BottomLeft.md) (structural, 1×, HIGH)
- [BottomRight](element_BottomRight.md) (structural, 1×, HIGH)
- [Disabled](element_Disabled.md) (structural, 1×, HIGH)
- [MiddleCenter](element_MiddleCenter.md) (structural, 1×, HIGH)
- [MiddleLeft](element_MiddleLeft.md) (structural, 1×, HIGH)
- [MiddleRight](element_MiddleRight.md) (structural, 1×, HIGH)
- [Normal](element_Normal.md) (structural, 1×, HIGH)
- [NormalHighlit](element_NormalHighlit.md) (structural, 1×, HIGH)
- [Pressed](element_Pressed.md) (structural, 1×, HIGH)
- [TopCenter](element_TopCenter.md) (structural, 1×, HIGH)
- [TopLeft](element_TopLeft.md) (structural, 1×, HIGH)
- [TopRight](element_TopRight.md) (structural, 1×, HIGH)

## Seen In

- CM_ClosetGoblin

## Examples

- CM_ClosetGoblin: ClosetGoblinDefaultBG -> TexCoords in FullResizeImage ClosetGoblinDefaultBG
- CM_ClosetGoblin: ClosetGoblinDefaultButton -> TexCoords in Button ClosetGoblinDefaultButton

## Related APIs

- [Button](element_Button.md) (HIGH 90/100) - XML Element Type
- [FullResizeImage](element_FullResizeImage.md) (HIGH 90/100) - XML Element Type
- [BottomCenter](element_BottomCenter.md) (LOW 5/100) - XML Element Type
- [BottomLeft](element_BottomLeft.md) (LOW 5/100) - XML Element Type
- [BottomRight](element_BottomRight.md) (LOW 5/100) - XML Element Type
- [Disabled](element_Disabled.md) (LOW 5/100) - XML Element Type
- [MiddleCenter](element_MiddleCenter.md) (LOW 5/100) - XML Element Type
- [MiddleLeft](element_MiddleLeft.md) (LOW 5/100) - XML Element Type
- [MiddleRight](element_MiddleRight.md) (LOW 5/100) - XML Element Type
- [Normal](element_Normal.md) (LOW 5/100) - XML Element Type
- [NormalHighlit](element_NormalHighlit.md) (LOW 5/100) - XML Element Type
- [Pressed](element_Pressed.md) (LOW 5/100) - XML Element Type
- [TopCenter](element_TopCenter.md) (LOW 5/100) - XML Element Type
- [TopLeft](element_TopLeft.md) (LOW 5/100) - XML Element Type
- [TopRight](element_TopRight.md) (LOW 5/100) - XML Element Type
