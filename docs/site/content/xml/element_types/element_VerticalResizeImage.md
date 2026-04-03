# VerticalResizeImage

- Category: XML Element Type
- Confidence level: HIGH
- Confidence score: 100/100

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 135

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, seen in 4 or more addons, used directly in xml handler attributes.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +30 Used directly in XML handler attributes: XML exposure suggests an engine-level contract.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | BuffHead, CMap, DammazKron, EA_UiDebugTools, Minmap, NerfedButtons, RVMOD_PlayerStatus, SpamBayes |
| Files seen in | Core/Tome/DK_TOP.xml, Setup/SetupLayout.xml, Source/DebugWindowVerticalScroll.xml, UI/NBSBCore.xml |
| Namespaces detected | VerticalResizeImage |
| Source kinds | xml_frames |
| Example locations | BuffHead: BuffHeadLayoutVerticalButtonHighlight, BuffHead: BuffHeadLayoutVerticalButtonNormal, BuffHead: BuffHeadLayoutVerticalButtonPressed, BuffHead: BuffHeadLayoutVerticalResizeImage, CMap: HUDInfluenceBarBackground, CMap: HUDInfluenceBarFill |
| XML usage count | 27 |
| XML attribute usage count | 27 |
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

VerticalResizeImage is a container-style XML element. It commonly appears under Window. It is typically used to organize structural children such as Anchors, Size, Sizes.

## Common Attributes

- handleinput
- inherits
- layer
- name
- reverseFill
- reversefill
- savesettings
- sticky
- texture
- textureScale
- virtual

## Common Inherits

- BuffHeadLayoutVerticalResizeImage
- DK_VertLine
- EA_VerticalResizeImage_DefaultVerticalSeparatorMiddle
- RVMOD_PlayerStatusVerticalBarTintableTemplate
- TomeVertResizeLine
- VerticalMoraleBarSlice

## Common Parent Elements

- [Windows](element_Windows.md) — 27× (HIGH)

## Common Structural Child Elements

- [Anchors](element_Anchors.md) — 14× (HIGH)
- [Size](element_Size.md) — 10× (HIGH)
- [Sizes](element_Sizes.md) — 8× (HIGH)
- [TexCoords](element_TexCoords.md) — 5× (MEDIUM)
- [TintColor](element_TintColor.md) — 4× (MEDIUM)
- [TexSlices](element_TexSlices.md) — 3× (MEDIUM)

## Common Template Bases

- BuffHeadLayoutVerticalResizeImage
- DK_VertLine
- EA_VerticalResizeImage_DefaultVerticalSeparatorMiddle
- RVMOD_PlayerStatusVerticalBarTintableTemplate
- TomeVertResizeLine
- VerticalMoraleBarSlice


> **Note**: This element type commonly acts as a template base.

## Attribute Reference

| Attribute | Required | Usage % | Sample Values |
| --- | --- | --- | --- |
| `inherits` | optional | 70% | BuffHeadLayoutVerticalResizeImage, TomeVertResizeLine, DK_VertLine, EA_VerticalResizeImage_DefaultVerticalSeparatorMiddle, ... |
| `texture` | optional | 29% | EA_TintableImage, ui_main, EA_HUD_01, spambayes_tint_scrollbar, ... |
| `handleinput` | optional | 22% | false |
| `layer` | optional | 7% | secondary, background |
| `savesettings` | optional | 7% | false |
| `textureScale` | optional | 7% | 0.3333, 0.6666 |
| `virtual` | optional | 7% | true |
| `reverseFill` | optional | 3% | true |
| `reversefill` | optional | 3% | true |
| `sticky` | optional | 3% | false |
## Structural Sub-Elements

### [Anchors](element_Anchors.md)

Observed 14 times as an unnamed child.

### [Size](element_Size.md)

Observed 10 times as an unnamed child.

### [Sizes](element_Sizes.md)

Observed 8 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `middle` | optional | 30, 200, 90, 50 |
| `left` | optional | 0, 7, 13, 15 |
| `right` | optional | 0, 7, 13, 15 |
| `bottom` | optional | 0, 13, 3 |
| `top` | optional | 0, 13, 3 |
### [TexCoords](element_TexCoords.md)

Observed 5 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `x` | optional | 0, 133, 88, 113 |
| `y` | optional | 0, 163, 51, 90 |
### [TintColor](element_TintColor.md)

Observed 4 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `b` | **required** | 0, 110, 130, 50 |
| `g` | **required** | 0, 200, 185, 130 |
| `r` | **required** | 255, 0, 180, 200 |
| `a` | optional | 30, 255, 0.8, 175 |
### [TexSlices](element_TexSlices.md)

Observed 3 times as an unnamed child.

## Recursive Hierarchy

- Root: [VerticalResizeImage](element_VerticalResizeImage.md)
- [Anchors](element_Anchors.md) (structural, 14×, HIGH)
  - [AbsPoint](element_AbsPoint.md) (structural, 6×, MEDIUM)
  - [Anchor](element_Anchor.md) (structural, 14161×, HIGH)
    - [AbsPoint](element_AbsPoint.md) (structural, 12549×, HIGH)
    - [Anchor](element_Anchor.md) (structural, 29×, HIGH)
      - (cycle)
- [Size](element_Size.md) (structural, 10×, HIGH)
  - [AbsPoint](element_AbsPoint.md) (structural, 9073×, HIGH)
- [Sizes](element_Sizes.md) (structural, 8×, HIGH)
  - [BottomRight](element_BottomRight.md) (structural, 30×, HIGH)
  - [Middle](element_Middle.md) (structural, 30×, HIGH)
  - [TopLeft](element_TopLeft.md) (structural, 30×, HIGH)
- [TexCoords](element_TexCoords.md) (structural, 5×, MEDIUM)
  - [Bottom](element_Bottom.md) (structural, 2×, LOW)
  - [BottomCenter](element_BottomCenter.md) (structural, 19×, HIGH)
  - [BottomLeft](element_BottomLeft.md) (structural, 19×, HIGH)
  - [BottomRight](element_BottomRight.md) (structural, 19×, HIGH)
  - [Disabled](element_Disabled.md) (structural, 43×, HIGH)
  - [Left](element_Left.md) (structural, 33×, HIGH)
  - [Middle](element_Middle.md) (structural, 41×, HIGH)
  - [MiddleCenter](element_MiddleCenter.md) (structural, 19×, HIGH)
  - [MiddleLeft](element_MiddleLeft.md) (structural, 19×, HIGH)
  - [MiddleRight](element_MiddleRight.md) (structural, 19×, HIGH)
  - [Normal](element_Normal.md) (structural, 47×, HIGH)
  - [NormalHighlit](element_NormalHighlit.md) (structural, 43×, HIGH)
  - [Pressed](element_Pressed.md) (structural, 46×, HIGH)
  - [PressedHighlit](element_PressedHighlit.md) (structural, 10×, HIGH)
  - [Right](element_Right.md) (structural, 33×, HIGH)
  - [Top](element_Top.md) (structural, 2×, LOW)
  - [TopCenter](element_TopCenter.md) (structural, 19×, HIGH)
  - [TopLeft](element_TopLeft.md) (structural, 19×, HIGH)
  - [TopRight](element_TopRight.md) (structural, 19×, HIGH)
- [TexSlices](element_TexSlices.md) (structural, 3×, MEDIUM)
  - [Bottom](element_Bottom.md) (structural, 2×, LOW)
  - [BottomCenter](element_BottomCenter.md) (structural, 11×, HIGH)
  - [BottomLeft](element_BottomLeft.md) (structural, 11×, HIGH)
  - [BottomRight](element_BottomRight.md) (structural, 11×, HIGH)
  - [Disabled](element_Disabled.md) (structural, 53×, HIGH)
  - [DisabledPressed](element_DisabledPressed.md) (structural, 36×, HIGH)
  - [Left](element_Left.md) (structural, 3×, MEDIUM)
  - [Middle](element_Middle.md) (structural, 10×, HIGH)
  - [MiddleCenter](element_MiddleCenter.md) (structural, 11×, HIGH)
  - [MiddleLeft](element_MiddleLeft.md) (structural, 11×, HIGH)
  - [MiddleRight](element_MiddleRight.md) (structural, 11×, HIGH)
  - [Normal](element_Normal.md) (structural, 120×, HIGH)
  - [NormalHighlit](element_NormalHighlit.md) (structural, 104×, HIGH)
  - [Pressed](element_Pressed.md) (structural, 117×, HIGH)
  - [PressedHighlit](element_PressedHighlit.md) (structural, 95×, HIGH)
  - [Right](element_Right.md) (structural, 3×, MEDIUM)
  - [Top](element_Top.md) (structural, 2×, LOW)
  - [TopCenter](element_TopCenter.md) (structural, 11×, HIGH)
  - [TopLeft](element_TopLeft.md) (structural, 11×, HIGH)
  - [TopRight](element_TopRight.md) (structural, 11×, HIGH)
- [TintColor](element_TintColor.md) (structural, 4×, MEDIUM)

## Seen In

- BuffHead
- CMap
- DammazKron
- EA_UiDebugTools
- Minmap
- NerfedButtons
- RVMOD_PlayerStatus
- SpamBayes
- TastyButtons
- VerticalMorale

## Examples

- BuffHead: BuffHeadLayoutVerticalButtonHighlight -> VerticalResizeImage BuffHeadLayoutVerticalButtonHighlight
- BuffHead: BuffHeadLayoutVerticalButtonNormal -> VerticalResizeImage BuffHeadLayoutVerticalButtonNormal
- BuffHead: BuffHeadLayoutVerticalButtonPressed -> VerticalResizeImage BuffHeadLayoutVerticalButtonPressed
- BuffHead: BuffHeadLayoutVerticalResizeImage -> VerticalResizeImage BuffHeadLayoutVerticalResizeImage
- CMap: HUDInfluenceBarBackground -> VerticalResizeImage HUDInfluenceBarBackground
- CMap: HUDInfluenceBarFill -> VerticalResizeImage HUDInfluenceBarFill

## Related APIs

- [Anchors](element_Anchors.md) (HIGH 100/100) - XML Element Type
- [EA_VerticalResizeImage_DefaultVerticalSeparatorMiddle](../../globals/constants/constant_EA_VerticalResizeImage_DefaultVerticalSeparatorMiddle.md) (HIGH 100/100) - Constant
- [Size](element_Size.md) (HIGH 100/100) - XML Element Type
- [Sizes](element_Sizes.md) (HIGH 100/100) - XML Element Type
- [TexCoords](element_TexCoords.md) (HIGH 100/100) - XML Element Type
- [TexSlices](element_TexSlices.md) (HIGH 100/100) - XML Element Type
- [TintColor](element_TintColor.md) (HIGH 100/100) - XML Element Type
- [Windows](element_Windows.md) (HIGH 100/100) - XML Element Type
