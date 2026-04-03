# Eventhandlers

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
| Addons seen in | MapMonster |
| Files seen in | Source/MapMonster_Templates.xml |
| Namespaces detected | Eventhandlers |
| Source kinds | xml_frames |
| Example locations | MapMonster: MapFilterContextSubMenuChoice |
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

Eventhandlers is a structural XML sub-element. It commonly appears under Button. It is typically used to organize structural children such as EventHandler.

## Common Attributes

- none

## Common Inherits

- none

## Common Parent Elements

- [Button](element_Button.md) — 1× (HIGH)

## Common Structural Child Elements

- [EventHandler](element_EventHandler.md) — 2× (HIGH)

## Structural Sub-Elements

### [EventHandler](element_EventHandler.md)

Observed 2 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `event` | **required** | OnShown, OnHidden, OnLButtonUp, OnSelChanged |
| `function` | **required** | APAGui.OnShown, APAGui.OnHidden, APAGui.Hide, APAGui.OnTabButtonUp |
## Recursive Hierarchy

- Root: [Eventhandlers](element_Eventhandlers.md)
- [EventHandler](element_EventHandler.md) (structural, 2×, HIGH)

## Seen In

- MapMonster

## Examples

- MapMonster: MapFilterContextSubMenuChoice -> Eventhandlers in Button MapFilterContextSubMenuChoice

## Related APIs

- [Button](element_Button.md) (HIGH 100/100) - XML Element Type
- [EventHandler](element_EventHandler.md) (MEDIUM 45/100) - XML Element Type
