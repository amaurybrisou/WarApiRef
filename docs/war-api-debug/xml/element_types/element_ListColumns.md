# ListColumns

- Category: XML Element Type
- Confidence level: LOW
- Confidence score: 15/100

## Confidence Assessment

- Level: LOW

- Score: 15/100

- Rationale: unknown

## Evidence Signals

- none

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | none |
| Source kinds | none |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 0 |
| Global usage count | 0 |
| Local definition count | 0 |
| Documentation references | 0 |
| Initialization flow references | 0 |
| Known engine namespace | no |
| Default UI presence | no |
| Event binding presence | no |
| Observed in XML and Lua | no |
| Consistent role | no |
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

ListColumns is a structural XML container that declares list column mappings. It commonly appears under ListBox and contains ListColumn entries.

## Common Attributes

- none

## Common Inherits

- none

## Common Parent Elements

- [ListData](element_ListData.md) — 2× (HIGH)

## Common Structural Child Elements

- [ListColumn](element_ListColumn.md) — 4× (HIGH)

## Structural Sub-Elements

### [ListColumn](element_ListColumn.md)

Observed 4 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `format` | **required** | wstring |
| `variable` | **required** | name, tactics, zone, set |
| `windowname` | **required** | Name, Tactics, Zone, Set |
## Recursive Hierarchy

- Root: [ListColumns](element_ListColumns.md)
- [ListColumn](element_ListColumn.md) (structural, 4×, HIGH)

## Seen In

- CM_ClosetGoblin

## Examples

- none

## Related APIs

- [ListData](element_ListData.md) (HIGH 90/100) - XML Element Type
- [ListColumn](element_ListColumn.md) (LOW 15/100) - XML Element Type
