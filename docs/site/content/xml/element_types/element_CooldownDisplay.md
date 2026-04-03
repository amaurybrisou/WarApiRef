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
| Files seen in | `/workspace/data/raw/MoraleCircle/MoraleCircle.xml:0` |
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

## Common Inherits

- none

## Common Parent Elements

- [Windows](element_Windows.md) — 1× (HIGH)

## Common Structural Child Elements

- [Size](element_Size.md) — 1× (HIGH)


> **Note**: This element type commonly acts as a template base.

## Typical XML Structure

```xml
<CooldownDisplay cooldownshape="circle" handleinput="false" layer="overlay" movable="true" name="..." popable="false" segments="32" sticky="true">
  <Size/>
</CooldownDisplay>
```

## Attribute Reference

| Attribute | Required | Usage % | Sample Values |
| --- | --- | --- | --- |
| `cooldownshape` | **required** | 100% | circle |
| `handleinput` | **required** | 100% | false |
| `layer` | **required** | 100% | overlay |
| `movable` | **required** | 100% | true |
| `popable` | **required** | 100% | false |
| `segments` | **required** | 100% | 32 |
| `sticky` | **required** | 100% | true |
## Structural Sub-Elements

### [Size](element_Size.md)

Observed 1 times as an unnamed child.

## Seen In

- MoraleCircle

## Examples

- MoraleCircle: MoraleTemplateAbilityCooldown -> CooldownDisplay MoraleTemplateAbilityCooldown

## Related APIs

- [Size](element_Size.md) (HIGH 100/100) - XML Element Type
- [Windows](element_Windows.md) (HIGH 100/100) - XML Element Type

## Used With

- none

## Triggered By

- none

## Affects

- none
