# ScrollWindow

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
| Addons seen in | PartyCast |
| Files seen in | `/workspace/data/raw/PartyCast/PartyCast.xml:550` |
| Namespaces detected | ScrollWindow |
| Source kinds | xml_frames |
| Example locations | PartyCast: PartyCastWindow_Template_SpecialNewCastScroll |
| XML usage count | 1 |
| XML attribute usage count | 1 |
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

Observed XML element type instantiated by 1 addons.

## Common Attributes

- autoHideScrollBar
- childscrollwindow
- lineheight
- name

## Common Inherits

- none

## Common Parent Elements

- [Window](element_Window.md)

## Common Named Child Elements

- [Window](element_Window.md)


> **Note**: This element type commonly acts as a template base.

## Attribute Reference

| Attribute | Required | Usage % | Sample Values |
| --- | --- | --- | --- |
| `autoHideScrollBar` | **required** | 100% | true |
| `childscrollwindow` | **required** | 100% | $parent_scroll |
| `lineheight` | **required** | 100% | 38 |
## Seen In

- PartyCast

## Examples

- PartyCast: PartyCastWindow_Template_SpecialNewCastScroll -> ScrollWindow PartyCastWindow_Template_SpecialNewCastScroll

## Related APIs

- [Window](element_Window.md) (HIGH 100/100) - XML Element Type

## Used With

- none

## Triggered By

- none

## Affects

- none
