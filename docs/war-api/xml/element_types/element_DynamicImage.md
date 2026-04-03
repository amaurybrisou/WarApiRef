# DynamicImage

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
| Addons seen in | InfoScroller, Moth, PartyCast, Soloq, TidyRoll, minesweep |
| Files seen in | `/workspace/data/raw/InfoScroller/InfoScroller.xml:111`, `/workspace/data/raw/InfoScroller/InfoScroller.xml:122`, `/workspace/data/raw/InfoScroller/InfoScroller.xml:145`, `/workspace/data/raw/InfoScroller/InfoScroller.xml:156`, `/workspace/data/raw/InfoScroller/InfoScroller.xml:179`, `/workspace/data/raw/InfoScroller/InfoScroller.xml:190`, `/workspace/data/raw/InfoScroller/InfoScroller.xml:214`, `/workspace/data/raw/InfoScroller/InfoScroller.xml:225` |
| Namespaces detected | DynamicImage |
| Source kinds | xml_frames |
| Example locations | InfoScroller: InfoScrollerTemplateBackGroundEnd, InfoScroller: InfoScrollerTemplateBackGroundStart, InfoScroller: InfoScrollerTemplateIcon1LeftIcon, InfoScroller: InfoScrollerTemplateIcon1RightIcon, InfoScroller: InfoScrollerTemplateIcon2LeftIcon, InfoScroller: InfoScrollerTemplateIcon2RightIcon |
| XML usage count | 39 |
| XML attribute usage count | 39 |
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

Observed XML element type instantiated by 6 addons.

## Common Attributes

- name
- handleinput
- texture
- textureScale
- layer
- alpha
- slice
- texturescale
- inherits
- popable
- mirrorTexCoords

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

## Common Structural Child Elements

- [TintColor](element_TintColor.md) — 2× (MEDIUM)

## Common Template Bases

- EA_Image_DefaultIcon
- EA_Image_DefaultIconFrame
- TargetLevelBackgroundTemplate


> **Note**: This element type commonly acts as a template base.

## Attribute Reference

| Attribute | Required | Usage % | Sample Values |
| --- | --- | --- | --- |
| `handleinput` | **required** | 94% | false, true |
| `texture` | optional | 61% | SEPERATOR, Fade, EA_HUD_01, map_markers01, ... |
| `textureScale` | optional | 58% | 1, 0.75, 1.0, 0.45, ... |
| `layer` | optional | 46% | default, background, overlay |
| `alpha` | optional | 43% | 1, 0.4 |
| `slice` | optional | 12% | RvR-Flag, LordHeroSpecial-Skull, tactic-white, Neutral-Frame |
| `texturescale` | optional | 10% | 1, 0.8, 0.425 |
| `inherits` | optional | 7% | TargetLevelBackgroundTemplate, EA_Image_DefaultIcon, EA_Image_DefaultIconFrame |
| `popable` | optional | 5% | false |
| `mirrorTexCoords` | optional | 2% | true |
## Structural Sub-Elements

### [TintColor](element_TintColor.md)

Observed 2 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `b` | **required** | 0, 36, 166, 200 |
| `g` | **required** | 0, 28, 84, 200 |
| `r` | **required** | 0, 237, 200, 55 |
| `a` | optional | 1, 0.4, 0.5, 0.8 |
## Lua Functions Manipulating This Type

- Moth.UpdateTarget

## Seen In

- InfoScroller
- Moth
- PartyCast
- Soloq
- TidyRoll
- minesweep

## Examples

- InfoScroller: InfoScrollerTemplateBackGroundEnd -> DynamicImage InfoScrollerTemplateBackGroundEnd
- InfoScroller: InfoScrollerTemplateBackGroundStart -> DynamicImage InfoScrollerTemplateBackGroundStart
- InfoScroller: InfoScrollerTemplateIcon1LeftIcon -> DynamicImage InfoScrollerTemplateIcon1LeftIcon
- InfoScroller: InfoScrollerTemplateIcon1RightIcon -> DynamicImage InfoScrollerTemplateIcon1RightIcon
- InfoScroller: InfoScrollerTemplateIcon2LeftIcon -> DynamicImage InfoScrollerTemplateIcon2LeftIcon
- InfoScroller: InfoScrollerTemplateIcon2RightIcon -> DynamicImage InfoScrollerTemplateIcon2RightIcon

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
