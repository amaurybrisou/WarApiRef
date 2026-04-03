# Color

- Category: XML Element Type
- Confidence level: HIGH
- Confidence score: 98/100

## Confidence Assessment

- Level: HIGH

- Score: 98/100

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, used directly in xml handler attributes, seen in 2 to 3 addons.

## Evidence Signals

- +18 Seen in 2 to 3 addons: Cross-addon spread is present but limited.
- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +30 Used directly in XML handler attributes: XML exposure suggests an engine-level contract.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | CM_ClosetGoblin, Clock |
| Namespaces detected | Color |
| Source kinds | xml_frames |
| Example locations | CM_ClosetGoblin: ClosetGoblinActionBarPageSelectorCurrentPageText, CM_ClosetGoblin: ClosetGoblinTalismanLabel, CM_ClosetGoblin: ClosetGoblinZoneWindowContentsZoneOption, Clock: ClockWindowText |
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

Color is a XML UI element. It commonly appears under Label.

## Common Attributes

- a
- b
- g
- r

## Common Inherits

- none

## Common Parent Elements

- [Label](element_Label.md) — 4× (HIGH)

## Attribute Reference

| Attribute | Required | Usage % | Sample Values |
| --- | --- | --- | --- |
| `a` | **required** | 100% | 255 |
| `b` | **required** | 100% | 255 |
| `g` | **required** | 100% | 255 |
| `r` | **required** | 100% | 255 |
## Seen In

- CM_ClosetGoblin
- Clock

## Examples

- CM_ClosetGoblin: ClosetGoblinActionBarPageSelectorCurrentPageText -> Color in Label ClosetGoblinActionBarPageSelectorCurrentPageText
- CM_ClosetGoblin: ClosetGoblinTalismanLabel -> Color in Label ClosetGoblinTalismanLabel
- CM_ClosetGoblin: ClosetGoblinZoneWindowContentsZoneOption -> Color in Label ClosetGoblinZoneWindowContentsZoneOption
- Clock: ClockWindowText -> Color in Label ClockWindowText

## Related APIs

- [Label](element_Label.md) (HIGH 100/100) - XML Element Type
