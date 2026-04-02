# CooldownDisplay

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
| Addons seen in | MoraleCircle |
| Files seen in | `/workspace_addons/MoraleCircle/MoraleCircle.xml:116` |
| Namespaces detected | CooldownDisplay |
| Source kinds | xml_frames |
| Example locations | MoraleCircle: MoraleTemplateAbilityCooldown |
| XML usage count | 1 |
| XML attribute usage count | 1 |
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

- cooldownshape
- handleinput
- layer
- movable
- name
- popable
- segments
- sticky

## Common Handlers

- none

## Common Inherits

- none

## Seen In

- MoraleCircle

## Examples

- MoraleCircle: MoraleTemplateAbilityCooldown -> CooldownDisplay MoraleTemplateAbilityCooldown

## Related APIs

- none

## Used With

- none

## Triggered By

- none

## Affects

- none
