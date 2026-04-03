# Rotation

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
| Addons seen in | RVMOD_3DPortrait |
| Namespaces detected | Rotation |
| Source kinds | xml_frames |
| Example locations | RVMOD_3DPortrait: 3DPortraitSceneTemplate, RVMOD_3DPortrait: 3DPortrait_ARCHMAGE_FEMALE, RVMOD_3DPortrait: 3DPortrait_ZEALOT_MALE, RVMOD_3DPortrait: CharacterWindowScene |
| XML usage count | 4 |
| XML attribute usage count | 4 |
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

Rotation is a XML UI element. It commonly appears under NifDisplay.

## Common Attributes

- x
- y
- z

## Common Inherits

- none

## Common Parent Elements

- [NifDisplay](element_NifDisplay.md) — 4× (HIGH)

## Attribute Reference

| Attribute | Required | Usage % | Sample Values |
| --- | --- | --- | --- |
| `x` | **required** | 100% | 0.0, 20.0 |
| `y` | **required** | 100% | 180.0 |
| `z` | **required** | 100% | 180.0, 155.0, 170.0 |
## Seen In

- RVMOD_3DPortrait

## Examples

- RVMOD_3DPortrait: 3DPortraitSceneTemplate -> Rotation in NifDisplay 3DPortraitSceneTemplate
- RVMOD_3DPortrait: 3DPortrait_ARCHMAGE_FEMALE -> Rotation in NifDisplay 3DPortrait_ARCHMAGE_FEMALE
- RVMOD_3DPortrait: 3DPortrait_ZEALOT_MALE -> Rotation in NifDisplay 3DPortrait_ZEALOT_MALE
- RVMOD_3DPortrait: CharacterWindowScene -> Rotation in NifDisplay CharacterWindowScene

## Related APIs

- [NifDisplay](element_NifDisplay.md) (MEDIUM 45/100) - XML Element Type
