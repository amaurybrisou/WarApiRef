# Visual

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
| Addons seen in | BankArkel |
| Files seen in | `/workspace/data/raw/BankArkel/BankArkel.xml:0` |
| Namespaces detected | Visual |
| Source kinds | xml_frames |
| Example locations | BankArkel: BankArkelBackpack |
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

- none

## Common Inherits

- none

## Common Parent Elements

- [Window](element_Window.md) — 1× (HIGH)

## Common Structural Child Elements

- [Color](element_Color.md) — 1× (HIGH)

## Structural Sub-Elements

### [Color](element_Color.md)

Observed 1 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `b` | **required** | 73, 225, 255, 55 |
| `g` | **required** | 218, 255, 155, 55 |
| `r` | **required** | 255, 155, 245, 175 |
| `a` | optional | 255, 0, 0.5, 0.8 |
## Recursive Hierarchy

- Root: [Visual](element_Visual.md)
- [Color](element_Color.md) (structural, 1×, HIGH)

## Seen In

- BankArkel

## Examples

- BankArkel: BankArkelBackpack -> Visual in Window BankArkelBackpack

## Related APIs

- [Window](element_Window.md) (HIGH 100/100) - XML Element Type
- [Color](element_Color.md) (MEDIUM 45/100) - XML Element Type

## Used With

- none

## Triggered By

- none

## Affects

- none
