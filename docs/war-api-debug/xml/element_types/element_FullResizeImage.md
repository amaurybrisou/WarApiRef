# FullResizeImage

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
| Addons seen in | CM_ClosetGoblin |
| Namespaces detected | FullResizeImage |
| Source kinds | xml_frames |
| Example locations | CM_ClosetGoblin: ClosetGoblinCharacterWindowContentsBG, CM_ClosetGoblin: ClosetGoblinDefaultBG, CM_ClosetGoblin: ClosetGoblinSetRowBackgroundName, CM_ClosetGoblin: ClosetGoblinSetRowSelectionBorder, CM_ClosetGoblin: ClosetGoblinZoneRowBackgroundZone, CM_ClosetGoblin: ClosetGoblinZoneRowSelectionBorder |
| XML usage count | 7 |
| XML attribute usage count | 7 |
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

FullResizeImage is a container-style XML element. It commonly appears under Window. It is typically used to organize structural children such as Anchors, Size, Sizes.

## Common Attributes

- handleinput
- inherits
- layer
- name
- texture

## Common Inherits

- ClosetGoblinDefaultBG
- DefaultWindowBackground
- EA_FullResizeImage_TintableSolidBackground

## Common Parent Elements

- [Windows](element_Windows.md) — 7× (HIGH)

## Common Structural Child Elements

- [Anchors](element_Anchors.md) — 6× (HIGH)
- [Size](element_Size.md) — 2× (MEDIUM)
- [Sizes](element_Sizes.md) — 1× (LOW)
- [TexCoords](element_TexCoords.md) — 1× (LOW)

## Common Template Bases

- ClosetGoblinDefaultBG
- DefaultWindowBackground
- EA_FullResizeImage_TintableSolidBackground

## Attribute Reference

| Attribute | Required | Usage % | Sample Values |
| --- | --- | --- | --- |
| `inherits` | **required** | 85% | EA_FullResizeImage_TintableSolidBackground, DefaultWindowBackground, ClosetGoblinDefaultBG |
| `handleinput` | optional | 42% | false |
| `layer` | optional | 42% | background |
| `texture` | optional | 14% | shared_01 |
## Structural Sub-Elements

### [Anchors](element_Anchors.md)

Observed 6 times as an unnamed child.

### [Size](element_Size.md)

Observed 2 times as an unnamed child.

### [Sizes](element_Sizes.md)

Observed 1 times as an unnamed child.

### [TexCoords](element_TexCoords.md)

Observed 1 times as an unnamed child.

## Recursive Hierarchy

- Root: [FullResizeImage](element_FullResizeImage.md)
- [Anchors](element_Anchors.md) (structural, 6×, HIGH)
  - [Anchor](element_Anchor.md) (structural, 167×, HIGH)
    - [AbsPoint](element_AbsPoint.md) (structural, 166×, HIGH)
- [Size](element_Size.md) (structural, 2×, MEDIUM)
  - [AbsPoint](element_AbsPoint.md) (structural, 39×, HIGH)
- [Sizes](element_Sizes.md) (structural, 1×, LOW)
  - [BottomRight](element_BottomRight.md) (structural, 1×, HIGH)
  - [Middle](element_Middle.md) (structural, 1×, HIGH)
  - [TopLeft](element_TopLeft.md) (structural, 1×, HIGH)
- [TexCoords](element_TexCoords.md) (structural, 1×, LOW)
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

- CM_ClosetGoblin: ClosetGoblinCharacterWindowContentsBG -> FullResizeImage ClosetGoblinCharacterWindowContentsBG
- CM_ClosetGoblin: ClosetGoblinDefaultBG -> FullResizeImage ClosetGoblinDefaultBG
- CM_ClosetGoblin: ClosetGoblinSetRowBackgroundName -> FullResizeImage ClosetGoblinSetRowBackgroundName
- CM_ClosetGoblin: ClosetGoblinSetRowSelectionBorder -> FullResizeImage ClosetGoblinSetRowSelectionBorder
- CM_ClosetGoblin: ClosetGoblinZoneRowBackgroundZone -> FullResizeImage ClosetGoblinZoneRowBackgroundZone
- CM_ClosetGoblin: ClosetGoblinZoneRowSelectionBorder -> FullResizeImage ClosetGoblinZoneRowSelectionBorder

## Related APIs

- [Windows](element_Windows.md) (HIGH 100/100) - XML Element Type
- [Anchors](element_Anchors.md) (HIGH 98/100) - XML Element Type
- [Size](element_Size.md) (HIGH 98/100) - XML Element Type
- [EA_FullResizeImage_TintableSolidBackground](../../globals/constants/constant_EA_FullResizeImage_TintableSolidBackground.md) (HIGH 90/100) - Constant
- [Sizes](element_Sizes.md) (MEDIUM 45/100) - XML Element Type
- [TexCoords](element_TexCoords.md) (MEDIUM 45/100) - XML Element Type

## Used With

- [Button](element_Button.md) (HIGH 90/100) - XML Element Type
