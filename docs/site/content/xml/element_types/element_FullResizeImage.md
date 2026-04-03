# FullResizeImage

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
| Addons seen in | InfoScroller, Moth, PartyCast, Soloq, TidyChat, TidyRoll, minesweep |
| Files seen in | `/workspace/data/raw/InfoScroller/InfoScroller.xml:236`, `/workspace/data/raw/Moth/Moth.xml:109`, `/workspace/data/raw/Moth/Moth.xml:128`, `/workspace/data/raw/Moth/Moth.xml:157`, `/workspace/data/raw/Moth/Moth.xml:18`, `/workspace/data/raw/Moth/Moth.xml:30`, `/workspace/data/raw/Moth/Moth.xml:42`, `/workspace/data/raw/Moth/Moth.xml:53` |
| Namespaces detected | FullResizeImage |
| Source kinds | xml_frames |
| Example locations | InfoScroller: InfoScrollerTemplateBackGroundBG, Moth: MothBackground, Moth: MothBorderDark, Moth: MothBorderDecorative, Moth: MothBorderGold, Moth: MothBorderSilver |
| XML usage count | 29 |
| XML attribute usage count | 29 |
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

Observed XML element type instantiated by 7 addons.

## Common Attributes

- name
- inherits
- handleinput
- alpha
- layer
- texture
- movable

## Common Inherits

- EA_FullResizeImage_TintableSolidBackground
- EA_Button_ResizeIconFrameNormal
- EA_FullResizeImage_BlackTransparent
- EA_Button_ResizeIconFrameHighlight
- EA_Button_ResizeIconFramePressed
- EA_FullResizeImage_DefaultFrame
- EA_FullResizeImage_MetalFill
- EA_FullResizeImage_TanBorder

## Common Parent Elements

- [Window](element_Window.md)
- [Button](element_Button.md)

## Common Structural Child Elements

- [TintColor](element_TintColor.md) — 10× (HIGH)

## Common Template Bases

- EA_Button_ResizeIconFrameHighlight
- EA_Button_ResizeIconFrameNormal
- EA_Button_ResizeIconFramePressed
- EA_FullResizeImage_BlackTransparent
- EA_FullResizeImage_DefaultFrame
- EA_FullResizeImage_MetalFill
- EA_FullResizeImage_TanBorder
- EA_FullResizeImage_TintableSolidBackground


> **Note**: This element type commonly acts as a template base.

## Attribute Reference

| Attribute | Required | Usage % | Sample Values |
| --- | --- | --- | --- |
| `inherits` | **required** | 89% | EA_FullResizeImage_TintableSolidBackground, EA_Button_ResizeIconFrameNormal, EA_Button_ResizeIconFramePressed, EA_Button_ResizeIconFrameHighlight, ... |
| `handleinput` | optional | 44% | false, true |
| `alpha` | optional | 31% | 0.4, 0.5, 0, 0.8, ... |
| `layer` | optional | 24% | background, overlay, default |
| `texture` | optional | 10% | EA_HUD_01 |
| `movable` | optional | 3% | false |
## Structural Sub-Elements

### [TintColor](element_TintColor.md)

Observed 10 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `b` | **required** | 0, 36, 166, 200 |
| `g` | **required** | 0, 28, 84, 200 |
| `r` | **required** | 0, 237, 200, 55 |
| `a` | optional | 1, 0.4, 0.5, 0.8 |
## Lua Functions Manipulating This Type

- Moth.Clear
- Moth.UpdateHealthBar
- Moth.HealthBar
- Moth.HideBorders

## Seen In

- InfoScroller
- Moth
- PartyCast
- Soloq
- TidyChat
- TidyRoll
- minesweep

## Examples

- InfoScroller: InfoScrollerTemplateBackGroundBG -> FullResizeImage InfoScrollerTemplateBackGroundBG
- Moth: MothBackground -> FullResizeImage MothBackground
- Moth: MothBorderDark -> FullResizeImage MothBorderDark
- Moth: MothBorderDecorative -> FullResizeImage MothBorderDecorative
- Moth: MothBorderGold -> FullResizeImage MothBorderGold
- Moth: MothBorderSilver -> FullResizeImage MothBorderSilver

## Related APIs

- [Button](element_Button.md) (HIGH 100/100) - XML Element Type
- [TintColor](element_TintColor.md) (HIGH 100/100) - XML Element Type
- [Window](element_Window.md) (HIGH 100/100) - XML Element Type

## Used With

- none

## Triggered By

- none

## Affects

- none
