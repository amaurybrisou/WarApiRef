# TexDims

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
| Addons seen in | Aura, AutoMark, BuffHead, DAoCBuff, Enemy, PartyCast, PotionBar, Shinies |
| Files seen in | `/workspace/data/raw/Aura/Source/Templates.xml:0`, `/workspace/data/raw/AutoMark/Source/AutoMark.xml:0`, `/workspace/data/raw/BuffHead/Display.xml:0`, `/workspace/data/raw/BuffHead/Setup/General.xml:0`, `/workspace/data/raw/DAoCBuff/Source/DAoCBuff.xml:0`, `/workspace/data/raw/DAoCBuff/Source/DAoCBuffMsgWindow.xml:0`, `/workspace/data/raw/DAoCBuff/Source/DAoCBuffSettingsTabs.xml:1079`, `/workspace/data/raw/DAoCBuff/Source/DAoCBuffSettingsTabs.xml:1119` |
| Namespaces detected | TexDims |
| Source kinds | xml_frames |
| Example locations | Aura: AuraFrameImageCircle, Aura: AuraFrameImageSquare, Aura: AuraScreenFlashFrameImage, AutoMark: T_AutoMark_Marker_Border, AutoMark: T_AutoMark_Marker_Icon, BuffHead: BuffHeadBuffTemplateIconBorder |
| XML usage count | 74 |
| XML attribute usage count | 74 |
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

Observed XML element type instantiated by 11 addons.

## Common Attributes

- x
- y

## Common Inherits

- none

## Common Parent Elements

- [DynamicImage](element_DynamicImage.md) — 65× (HIGH)
- [Button](element_Button.md) — 4× (MEDIUM)
- [FullResizeImage](element_FullResizeImage.md) — 3× (MEDIUM)
- [CircleImage](element_CircleImage.md) — 2× (LOW)

## Attribute Reference

| Attribute | Required | Usage % | Sample Values |
| --- | --- | --- | --- |
| `x` | **required** | 100% | 64, 256, 32, 430, ... |
| `y` | **required** | 100% | 64, 256, 32, 430, ... |
## Seen In

- Aura
- AutoMark
- BuffHead
- DAoCBuff
- Enemy
- PartyCast
- PotionBar
- Shinies
- Swift Assist
- TexturedButtons
- TidyRoll

## Examples

- Aura: AuraFrameImageCircle -> TexDims in CircleImage AuraFrameImageCircle
- Aura: AuraFrameImageSquare -> TexDims in DynamicImage AuraFrameImageSquare
- Aura: AuraScreenFlashFrameImage -> TexDims in DynamicImage AuraScreenFlashFrameImage
- AutoMark: T_AutoMark_Marker_Border -> TexDims in DynamicImage T_AutoMark_Marker_Border
- AutoMark: T_AutoMark_Marker_Icon -> TexDims in DynamicImage T_AutoMark_Marker_Icon
- BuffHead: BuffHeadBuffTemplateIconBorder -> TexDims in DynamicImage BuffHeadBuffTemplateIconBorder

## Related APIs

- [Button](element_Button.md) (HIGH 100/100) - XML Element Type
- [CircleImage](element_CircleImage.md) (HIGH 100/100) - XML Element Type
- [DynamicImage](element_DynamicImage.md) (HIGH 100/100) - XML Element Type
- [FullResizeImage](element_FullResizeImage.md) (HIGH 100/100) - XML Element Type

## Used With

- none

## Triggered By

- none

## Affects

- none
