# ListColumns

- Category: XML Element Type
- Confidence level: MEDIUM
- Confidence score: 55/100

## Confidence Assessment

- Level: MEDIUM

- Score: 55/100

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

ListColumns is a structural XML container used by list controls to define how backing-table fields map to row windows.

## Common Attributes

- none

## Common Inherits

- none

## Common Parent Elements

- [ListData](element_ListData.md) — 25× (HIGH)

## Common Structural Child Elements

- [ListColumn](element_ListColumn.md) — 42× (HIGH)

## Structural Sub-Elements

### [ListColumn](element_ListColumn.md)

Observed 42 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `format` | **required** | wstring, number, icon |
| `variable` | **required** | Name, RankN, Character, Type |
| `windowname` | **required** | Name, Rank, Character, Type |
## Recursive Hierarchy

- Root: [ListColumns](element_ListColumns.md)
- [ListColumn](element_ListColumn.md) (structural, 42×, HIGH)

## Seen In

- AggroMeter
- Aura
- BuffHead
- CM_ClosetGoblin
- DAoCBuff
- Enemy
- Pocket Palette
- Shinies
- TexturedButtons
- TidyChat
- TidyRoll
- TurretRange

## Examples

- none

## Related APIs

- [ListData](element_ListData.md) (HIGH 100/100) - XML Element Type
- [ListColumn](element_ListColumn.md) (MEDIUM 45/100) - XML Element Type

## Used With

- none

## Triggered By

- none

## Affects

- none
