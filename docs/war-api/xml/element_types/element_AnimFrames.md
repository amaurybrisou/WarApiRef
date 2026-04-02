# AnimFrames

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
| Namespaces detected | AnimFrames |
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

- none

## Common Inherits

- none

## Common Parent Elements

- [AnimatedImage](element_AnimatedImage.md) — 2× (HIGH)

## Common Structural Child Elements

- [AnimFrame](element_AnimFrame.md) — 14× (HIGH)

## Structural Sub-Elements

### [AnimFrame](element_AnimFrame.md)

Observed 14 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `id` | **required** | 1, 2, 3, 4 |
| `x` | **required** | 0, 72, 144, 216 |
| `y` | **required** | 0, 85, 170 |
## Seen In

- TidyRoll

## Examples

- TidyRoll: TRollOverlayFlash -> AnimFrames in AnimatedImage TRollOverlayFlash
- TidyRoll: TRollOverlayGlow -> AnimFrames in AnimatedImage TRollOverlayGlow

## Related APIs

- none

## Used With

- none

## Triggered By

- none

## Affects

- none
