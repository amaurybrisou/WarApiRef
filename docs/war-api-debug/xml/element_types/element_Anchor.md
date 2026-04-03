# Anchor

- Category: XML Element Type
- Confidence level: MEDIUM
- Confidence score: 45/100

## Confidence Assessment

- Level: MEDIUM

- Score: 45/100

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

Anchor is a structural XML sub-element. It commonly appears under Anchors. It is typically used to organize structural children such as AbsPoint.

## Common Attributes

- point
- relativePoint
- relativeTo

## Common Inherits

- none

## Common Parent Elements

- [Anchors](element_Anchors.md) — 167× (HIGH)

## Common Structural Child Elements

- [AbsPoint](element_AbsPoint.md) — 166× (HIGH)

## Attribute Reference

| Attribute | Required | Usage % | Sample Values |
| --- | --- | --- | --- |
| `point` | **required** | 100% | left, topleft, bottomright, topright, ... |
| `relativePoint` | **required** | 100% | left, topleft, bottomright, top, ... |
| `relativeTo` | **required** | 98% | $parentName, $parentTactics, $parentZone, $parentSet, ... |
## Structural Sub-Elements

### [AbsPoint](element_AbsPoint.md)

Observed 166 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `x` | **required** | 75, 65, 355, 340 |
| `y` | **required** | 75, 20, 26, 0 |
## Recursive Hierarchy

- Root: [Anchor](element_Anchor.md)
- [AbsPoint](element_AbsPoint.md) (structural, 166×, HIGH)

## Seen In

- CM_ClosetGoblin
- Clock

## Examples

- none

## Related APIs

- [Anchors](element_Anchors.md) (HIGH 98/100) - XML Element Type
- [AbsPoint](element_AbsPoint.md) (MEDIUM 35/100) - XML Element Type
