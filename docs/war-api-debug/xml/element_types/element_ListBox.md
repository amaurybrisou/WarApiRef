# ListBox

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
| Namespaces detected | ListBox |
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

ListBox is a container-style XML element for row-based UI lists. It commonly works with structural children such as ListData and ListColumns and binds events to Lua callbacks.

## Common Attributes

- layer
- name
- rowcount
- rowdef
- rowspacing
- scrollbar
- visiblerows

## Common Inherits

- none

## Common Parent Elements

- [Windows](element_Windows.md) — 2× (HIGH)

## Common Structural Child Elements

- [Anchors](element_Anchors.md) — 2× (HIGH)
- [ListData](element_ListData.md) — 2× (HIGH)
- [Size](element_Size.md) — 1× (HIGH)

## Attribute Reference

| Attribute | Required | Usage % | Sample Values |
| --- | --- | --- | --- |
| `layer` | **required** | 100% | secondary |
| `rowcount` | **required** | 100% | 100 |
| `rowdef` | **required** | 100% | ClosetGoblinSetRow, ClosetGoblinZoneRow |
| `rowspacing` | **required** | 100% | 1 |
| `scrollbar` | **required** | 100% | EA_ScrollBar_DefaultVerticalChain |
| `visiblerows` | **required** | 100% | 16, 20 |
## Structural Sub-Elements

### [Anchors](element_Anchors.md)

Observed 2 times as an unnamed child.

### [ListData](element_ListData.md)

Observed 2 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `populationfunction` | **required** | ClosetGoblinCharacterWindow.UpdateSetRow, ClosetGoblinZoneWindow.UpdateSetRow |
| `table` | **required** | ClosetGoblin.currentProfile.sets, ClosetGoblinZoneWindow.associationTable |
### [Size](element_Size.md)

Observed 1 times as an unnamed child.

## Recursive Hierarchy

- Root: [ListBox](element_ListBox.md)
- [Anchors](element_Anchors.md) (structural, 2×, HIGH)
  - [Anchor](element_Anchor.md) (structural, 167×, HIGH)
    - [AbsPoint](element_AbsPoint.md) (structural, 166×, HIGH)
- [ListData](element_ListData.md) (structural, 2×, HIGH)
  - [ListColumns](element_ListColumns.md) (structural, 2×, HIGH)
    - [ListColumn](element_ListColumn.md) (structural, 4×, HIGH)
- [Size](element_Size.md) (structural, 1×, HIGH)
  - [AbsPoint](element_AbsPoint.md) (structural, 39×, HIGH)

## Lua Functions Manipulating This Type

- ClosetGoblinCharacterWindow.UpdateHighlightOnRow
- ClosetGoblinZoneWindow.UpdateHighlightOnRow

## Seen In

- CM_ClosetGoblin

## Examples

- CM_ClosetGoblin: ClosetGoblinCharacterWindowContentsSetList -> ListBox ClosetGoblinCharacterWindowContentsSetList
- CM_ClosetGoblin: ClosetGoblinZoneWindowContentsSetList -> ListBox ClosetGoblinZoneWindowContentsSetList

## Related APIs

- [Windows](element_Windows.md) (HIGH 100/100) - XML Element Type
- [Anchors](element_Anchors.md) (HIGH 98/100) - XML Element Type
- [Size](element_Size.md) (HIGH 98/100) - XML Element Type
- [ListData](element_ListData.md) (HIGH 90/100) - XML Element Type
