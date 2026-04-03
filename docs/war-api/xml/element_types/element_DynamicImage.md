# DynamicImage

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
| Addons seen in | Moth, TidyRoll |
| Files seen in | `/workspace/data/raw/Moth/Moth.xml:165`, `/workspace/data/raw/Moth/Moth.xml:185`, `/workspace/data/raw/Moth/Moth.xml:196`, `/workspace/data/raw/Moth/Moth.xml:207`, `/workspace/data/raw/Moth/Moth.xml:219`, `/workspace/data/raw/Moth/Moth.xml:97`, `/workspace/data/raw/TidyRoll/CustomAutoRoll.xml:16`, `/workspace/data/raw/TidyRoll/CustomAutoRoll.xml:26` |
| Namespaces detected | DynamicImage |
| Source kinds | xml_frames |
| Example locations | Moth: MothCellTemplateCareerIcon, Moth: MothCellTemplateDiffMaskIcon, Moth: MothCellTemplateMapPinIconSlice, Moth: MothCellTemplateNPCIcon, Moth: MothCellTemplateRankBackground, Moth: MothRvRFlagIndicator |
| XML usage count | 10 |
| XML attribute usage count | 10 |
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

Observed XML element type instantiated by 2 addons.

## Common Attributes

- name
- handleinput
- layer
- inherits
- texture
- textureScale
- popable
- slice

## Common Inherits

- EA_Image_DefaultIcon
- EA_Image_DefaultIconFrame
- TargetLevelBackgroundTemplate

## Common Parent Elements

- [Window](element_Window.md)
- [Button](element_Button.md)
- [DynamicImage](element_DynamicImage.md)

## Common Named Child Elements

- [DynamicImage](element_DynamicImage.md)

## Common Template Bases

- EA_Image_DefaultIcon
- EA_Image_DefaultIconFrame
- TargetLevelBackgroundTemplate


> **Note**: This element type commonly acts as a template base.

## Attribute Reference

| Attribute | Required | Usage % | Sample Values |
| --- | --- | --- | --- |
| `handleinput` | optional | 80% | false |
| `layer` | optional | 80% | overlay, default, background |
| `inherits` | optional | 30% | TargetLevelBackgroundTemplate, EA_Image_DefaultIcon, EA_Image_DefaultIconFrame |
| `texture` | optional | 30% | EA_HUD_01, map_markers01 |
| `textureScale` | optional | 30% | 0.75, 1.0, 0.875 |
| `popable` | optional | 20% | false |
| `slice` | optional | 20% | RvR-Flag, LordHeroSpecial-Skull |
## Lua Functions Manipulating This Type

- Moth.UpdateTarget

## Seen In

- Moth
- TidyRoll

## Examples

- Moth: MothCellTemplateCareerIcon -> DynamicImage MothCellTemplateCareerIcon
- Moth: MothCellTemplateDiffMaskIcon -> DynamicImage MothCellTemplateDiffMaskIcon
- Moth: MothCellTemplateMapPinIconSlice -> DynamicImage MothCellTemplateMapPinIconSlice
- Moth: MothCellTemplateNPCIcon -> DynamicImage MothCellTemplateNPCIcon
- Moth: MothCellTemplateRankBackground -> DynamicImage MothCellTemplateRankBackground
- Moth: MothRvRFlagIndicator -> DynamicImage MothRvRFlagIndicator

## Related APIs

- [Button](element_Button.md) (HIGH 100/100) - XML Element Type
- [Window](element_Window.md) (HIGH 100/100) - XML Element Type

## Used With

- none

## Triggered By

- none

## Affects

- none
