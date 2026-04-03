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
| Addons seen in | Aura, AutoMark, Brizio's Crappy Computer Medic, BuffHead, CCTV, CDown, Calling, CleansingBuddy |
| Files seen in | Code/Assist/Assist.xml, Code/CombatLog/CombatLogIDS.xml, Code/Core/Common.xml, Code/Core/Icon.xml, Code/GroupIcons/GroupIcons.xml, Code/Guard/GuardDistanceIndicator.xml, Code/Marks/MarkTemplate.xml, Code/Marks/Marks.xml |
| Namespaces detected | TexDims |
| Source kinds | xml_frames |
| Example locations | Aura: AuraFrameImageCircle, Aura: AuraFrameImageSquare, Aura: AuraScreenFlashFrameImage, AutoMark: T_AutoMark_Marker_Border, AutoMark: T_AutoMark_Marker_Icon, Brizio's Crappy Computer Medic: CCMGUIB1 |
| XML usage count | 189 |
| XML attribute usage count | 189 |
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

TexDims is a structural XML sub-element. It commonly appears under AnimatedImage and Button.

## Common Attributes

- x
- y

## Common Inherits

- none

## Common Parent Elements

- [DynamicImage](element_DynamicImage.md) — 156× (HIGH)
- [Button](element_Button.md) — 18× (HIGH)
- [CircleImage](element_CircleImage.md) — 11× (HIGH)
- [FullResizeImage](element_FullResizeImage.md) — 3× (MEDIUM)
- [AnimatedImage](element_AnimatedImage.md) — 1× (LOW)

## Attribute Reference

| Attribute | Required | Usage % | Sample Values |
| --- | --- | --- | --- |
| `x` | **required** | 100% | 64, 256, 32, 48, ... |
| `y` | **required** | 100% | 64, 256, 32, 48, ... |
## Seen In

- Aura
- AutoMark
- Brizio's Crappy Computer Medic
- BuffHead
- CCTV
- CDown
- Calling
- CleansingBuddy
- Compass3D
- CoolDownLine
- Crusher
- DAoCBuff
- DetauntHelper
- DuffTimer
- EA_LoadingScreen
- EZCraftX
- Enemy
- EveryBodyGuard
- Group Icons SG
- GroupRange
- GuardBot
- HealGrid
- KillStreak
- Minmap
- Motion
- NaturalLog
- NerfedButtons
- Obsidian
- Paint the leader
- PartyCast
- PotionBar
- Pure
- Pure Careerbar
- Queue Queuer
- RVMOD_Targets
- Shinies
- Squared
- Swift Assist
- TacticSetNames
- TastyButtons
- TexturedButtons
- TidyRoll
- Vectors
- WaaaghBar
- Warbuilder
- Wikki's Cooldown Bar
- Wikki's Cooldown Pulse
- compass
- yAssistHelper

## Examples

- Aura: AuraFrameImageCircle -> TexDims in CircleImage AuraFrameImageCircle
- Aura: AuraFrameImageSquare -> TexDims in DynamicImage AuraFrameImageSquare
- Aura: AuraScreenFlashFrameImage -> TexDims in DynamicImage AuraScreenFlashFrameImage
- AutoMark: T_AutoMark_Marker_Border -> TexDims in DynamicImage T_AutoMark_Marker_Border
- AutoMark: T_AutoMark_Marker_Icon -> TexDims in DynamicImage T_AutoMark_Marker_Icon
- Brizio's Crappy Computer Medic: CCMGUIB1 -> TexDims in DynamicImage CCMGUIB1

## Related APIs

- [AnimatedImage](element_AnimatedImage.md) (HIGH 100/100) - XML Element Type
- [Button](element_Button.md) (HIGH 100/100) - XML Element Type
- [CircleImage](element_CircleImage.md) (HIGH 100/100) - XML Element Type
- [DynamicImage](element_DynamicImage.md) (HIGH 100/100) - XML Element Type
- [DynamicImageSetTextureDimensions](../../window_api/functions/window_DynamicImageSetTextureDimensions.md) (HIGH 100/100) - Window Function
- [FullResizeImage](element_FullResizeImage.md) (HIGH 100/100) - XML Element Type

## Used With

- [DynamicImageSetTextureDimensions](../../window_api/functions/window_DynamicImageSetTextureDimensions.md) (HIGH 100/100) - Window Function
