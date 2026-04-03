# ListData

- Category: XML Element Type
- Confidence level: HIGH
- Confidence score: 90/100

## Confidence Assessment

- Level: HIGH

- Score: 90/100

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, used directly in xml handler attributes, matches a known engine namespace.

## Evidence Signals

- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +30 Used directly in XML handler attributes: XML exposure suggests an engine-level contract.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | CM_ClosetGoblin |
| Namespaces detected | ListData |
| Source kinds | xml_frames |
| Example locations | CM_ClosetGoblin: ClosetGoblinCharacterWindowContentsSetList, CM_ClosetGoblin: ClosetGoblinZoneWindowContentsSetList |
| XML usage count | 2 |
| XML attribute usage count | 2 |
| Lua usage count | 0 |
| Global usage count | 0 |
| Local definition count | 0 |
| Documentation references | 0 |
| Initialization flow references | 0 |
| Known engine namespace | yes |
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

ListData is a structural XML sub-element that binds list controls to Lua-backed row data. It commonly appears under ListBox as list data wiring.

## Common Attributes

- populationfunction
- table

## Common Inherits

- none

## Common Parent Elements

- [ListBox](element_ListBox.md) — 2× (HIGH)

## Common Structural Child Elements

- [ListColumns](element_ListColumns.md) — 2× (HIGH)

## Attribute Reference

| Attribute | Required | Usage % | Sample Values |
| --- | --- | --- | --- |
| `populationfunction` | **required** | 100% | ClosetGoblinCharacterWindow.UpdateSetRow, ClosetGoblinZoneWindow.UpdateSetRow |
| `table` | **required** | 100% | ClosetGoblin.currentProfile.sets, ClosetGoblinZoneWindow.associationTable |
## Structural Sub-Elements

### [ListColumns](element_ListColumns.md)

Observed 2 times as an unnamed child.

## Recursive Hierarchy

- Root: [ListData](element_ListData.md)
- [ListColumns](element_ListColumns.md) (structural, 2×, HIGH)
  - [ListColumn](element_ListColumn.md) (structural, 4×, HIGH)

## Seen In

- CM_ClosetGoblin

## Examples

- CM_ClosetGoblin: ClosetGoblinCharacterWindowContentsSetList -> ListData in ListBox ClosetGoblinCharacterWindowContentsSetList
- CM_ClosetGoblin: ClosetGoblinZoneWindowContentsSetList -> ListData in ListBox ClosetGoblinZoneWindowContentsSetList

## Related APIs

- [ListBox](element_ListBox.md) (HIGH 90/100) - XML Element Type
- [ListColumns](element_ListColumns.md) (LOW 15/100) - XML Element Type
