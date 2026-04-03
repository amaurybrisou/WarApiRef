# ActiveZoneOffset

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
| Addons seen in | EA_UiDebugTools, SpamBayes |
| Files seen in | Source/DebugWindowVerticalScroll.xml |
| Namespaces detected | ActiveZoneOffset |
| Source kinds | xml_frames |
| Example locations | EA_UiDebugTools: EA_ScrollBar_ChatVertical, SpamBayes: SpamBayesWindowLogTextScrollbar |
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

ActiveZoneOffset is a XML UI element. It commonly appears under VerticalScrollbar.

## Common Attributes

- x
- y

## Common Inherits

- none

## Common Parent Elements

- [VerticalScrollbar](element_VerticalScrollbar.md) — 2× (HIGH)

## Attribute Reference

| Attribute | Required | Usage % | Sample Values |
| --- | --- | --- | --- |
| `x` | **required** | 100% | 100 |
| `y` | **required** | 100% | 0 |
## Seen In

- EA_UiDebugTools
- SpamBayes

## Examples

- EA_UiDebugTools: EA_ScrollBar_ChatVertical -> ActiveZoneOffset in VerticalScrollbar EA_ScrollBar_ChatVertical
- SpamBayes: SpamBayesWindowLogTextScrollbar -> ActiveZoneOffset in VerticalScrollbar SpamBayesWindowLogTextScrollbar

## Related APIs

- [VerticalScrollbar](element_VerticalScrollbar.md) (HIGH 100/100) - XML Element Type
