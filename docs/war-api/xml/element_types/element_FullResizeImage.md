# FullResizeImage

- Category: XML Element Type
- Confidence level: HIGH
- Confidence score: 100/100

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 123

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, used directly in xml handler attributes, matches a known engine namespace.

## Evidence Signals

- +18 Seen in 2 to 3 addons: Cross-addon spread is present but limited.
- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +30 Used directly in XML handler attributes: XML exposure suggests an engine-level contract.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | Moth, TidyChat, TidyRoll |
| Files seen in | `/workspace/data/raw/Moth/Moth.xml:109`, `/workspace/data/raw/Moth/Moth.xml:128`, `/workspace/data/raw/Moth/Moth.xml:157`, `/workspace/data/raw/Moth/Moth.xml:18`, `/workspace/data/raw/Moth/Moth.xml:30`, `/workspace/data/raw/Moth/Moth.xml:42`, `/workspace/data/raw/Moth/Moth.xml:53`, `/workspace/data/raw/Moth/Moth.xml:64` |
| Namespaces detected | FullResizeImage |
| Source kinds | xml_frames |
| Example locations | Moth: MothBackground, Moth: MothBorderDark, Moth: MothBorderDecorative, Moth: MothBorderGold, Moth: MothBorderSilver, Moth: MothBorderTan |
| XML usage count | 17 |
| XML attribute usage count | 17 |
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

Observed XML element type instantiated by 3 addons.

## Common Attributes

- inherits
- name
- handleinput
- alpha
- layer

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

## Common Structural Child Elements

- [TintColor](element_TintColor.md) — 4× (MEDIUM)

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
| `inherits` | **required** | 100% | EA_FullResizeImage_TintableSolidBackground, EA_Button_ResizeIconFrameNormal, EA_Button_ResizeIconFramePressed, EA_Button_ResizeIconFrameHighlight, ... |
| `handleinput` | optional | 23% | false |
| `alpha` | optional | 17% | 0.5, 0 |
| `layer` | optional | 5% | overlay |
## Structural Sub-Elements

### [TintColor](element_TintColor.md)

Observed 4 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `b` | **required** | 0, 36, 166 |
| `g` | **required** | 0, 28, 84 |
| `r` | **required** | 0, 237 |
## Lua Functions Manipulating This Type

- Moth.HideBorders
- Moth.Clear
- Moth.HealthBar
- Moth.UpdateHealthBar

## Seen In

- Moth
- TidyChat
- TidyRoll

## Examples

- Moth: MothBackground -> FullResizeImage MothBackground
- Moth: MothBorderDark -> FullResizeImage MothBorderDark
- Moth: MothBorderDecorative -> FullResizeImage MothBorderDecorative
- Moth: MothBorderGold -> FullResizeImage MothBorderGold
- Moth: MothBorderSilver -> FullResizeImage MothBorderSilver
- Moth: MothBorderTan -> FullResizeImage MothBorderTan

## Related APIs

- [Window](element_Window.md) (HIGH 100/100) - XML Element Type
- [TintColor](element_TintColor.md) (MEDIUM 45/100) - XML Element Type

## Used With

- none

## Triggered By

- none

## Affects

- none
