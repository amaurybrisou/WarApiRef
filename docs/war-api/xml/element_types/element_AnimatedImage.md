# AnimatedImage

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
| Addons seen in | TidyRoll |
| Files seen in | `/workspace/data/raw/TidyRoll/TidyRoll.xml:137`, `/workspace/data/raw/TidyRoll/TidyRoll.xml:156` |
| Namespaces detected | AnimatedImage |
| Source kinds | xml_frames |
| Example locations | TidyRoll: TRollOverlayFlash, TidyRoll: TRollOverlayGlow |
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

Observed XML element type instantiated by 1 addons.

## Common Attributes

- fps
- handleinput
- layer
- name
- sticky
- texture
- textureScale
- texturescale

## Common Inherits

- none

## Common Parent Elements

- [Window](element_Window.md)

## Common Structural Child Elements

- [AnimFrames](element_AnimFrames.md) — 2× (HIGH)

## Attribute Reference

| Attribute | Required | Usage % | Sample Values |
| --- | --- | --- | --- |
| `fps` | **required** | 100% | 10 |
| `handleinput` | **required** | 100% | false |
| `layer` | **required** | 100% | overlay |
| `sticky` | **required** | 100% | false |
| `texture` | **required** | 100% | EA_ActionBarAnim_Casting, recharge_flash_anim |
| `textureScale` | optional | 50% | 0.9444 |
| `texturescale` | optional | 50% | 1 |
## Structural Sub-Elements

### [AnimFrames](element_AnimFrames.md)

Observed 2 times as an unnamed child.

## Seen In

- TidyRoll

## Examples

- TidyRoll: TRollOverlayFlash -> AnimatedImage TRollOverlayFlash
- TidyRoll: TRollOverlayGlow -> AnimatedImage TRollOverlayGlow

## Related APIs

- none

## Used With

- none

## Triggered By

- none

## Affects

- none
