# AnimFrame

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
| Namespaces detected | AnimFrame |
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

- id
- x
- y

## Common Inherits

- none

## Common Parent Elements

- [AnimFrames](element_AnimFrames.md) — 14× (HIGH)

## Attribute Reference

| Attribute | Required | Usage % | Sample Values |
| --- | --- | --- | --- |
| `id` | **required** | 100% | 1, 2, 3, 4, ... |
| `x` | **required** | 100% | 0, 72, 144, 216, ... |
| `y` | **required** | 100% | 0, 85, 170 |
## Seen In

- TidyRoll

## Examples

- TidyRoll: TRollOverlayFlash -> AnimFrame in AnimatedImage TRollOverlayFlash
- TidyRoll: TRollOverlayGlow -> AnimFrame in AnimatedImage TRollOverlayGlow

## Related APIs

- [AnimFrames](element_AnimFrames.md) (MEDIUM 45/100) - XML Element Type

## Used With

- none

## Triggered By

- none

## Affects

- none
