# Icon

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
| Addons seen in | NerfedButtons |
| Namespaces detected | Icon |
| Source kinds | xml_frames |
| Example locations | NerfedButtons: Cycle through friendlies in party or warband, NerfedButtons: Cycle through hurt friendlies in party or warband, NerfedButtons: Set the Bunny Hop Target, NerfedButtons: Target most hurt friendly in your party or warband, NerfedButtons: Target the highest priority Detaunt Helper target |
| XML usage count | 5 |
| XML attribute usage count | 5 |
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

Icon is a XML UI element. It commonly appears under Assets.

## Common Attributes

- id
- name
- texture

## Common Inherits

- none

## Common Parent Elements

- [Assets](element_Assets.md) — 114× (HIGH)

## Attribute Reference

| Attribute | Required | Usage % | Sample Values |
| --- | --- | --- | --- |
| `id` | **required** | 100% | 31001, 31004, 31006, 31007, ... |
| `texture` | **required** | 100% | emojis/001.tga, emojis/004.tga, emojis/006.tga, emojis/007.tga, ... |
## Seen In

- NerfedButtons

## Examples

- NerfedButtons: Cycle through friendlies in party or warband -> Icon Cycle through friendlies in party or warband
- NerfedButtons: Cycle through hurt friendlies in party or warband -> Icon Cycle through hurt friendlies in party or warband
- NerfedButtons: Set the Bunny Hop Target -> Icon Set the Bunny Hop Target
- NerfedButtons: Target most hurt friendly in your party or warband -> Icon Target most hurt friendly in your party or warband
- NerfedButtons: Target the highest priority Detaunt Helper target -> Icon Target the highest priority Detaunt Helper target

## Related APIs

- [Assets](element_Assets.md) (MEDIUM 45/100) - XML Element Type
