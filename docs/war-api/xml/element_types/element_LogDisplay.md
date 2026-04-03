# LogDisplay

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
| Addons seen in | Shinies |
| Files seen in | `/workspace/data/raw/Shinies/Modules/UI/Shinies-UI-Auto.xml:0` |
| Namespaces detected | LogDisplay |
| Source kinds | xml_frames |
| Example locations | Shinies: ShiniesAutoUI_AutoSummary |
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

- background
- font
- maxchars
- maxentries
- name
- scrollbar
- scrollbarPosition

## Common Inherits

- none

## Common Parent Elements

- [Windows](element_Windows.md) — 1× (HIGH)

## Common Structural Child Elements

- [Windows](element_Windows.md) — 1× (HIGH)

## Typical XML Structure

```xml
<LogDisplay font="font_chat_text" maxchars="4096" maxentries="256" name="..." scrollbar="$parentScrollbar" scrollbarPosition="right">
  <Windows/>
</LogDisplay>
```

## Attribute Reference

| Attribute | Required | Usage % | Sample Values |
| --- | --- | --- | --- |
| `background` | **required** | 100% |  |
| `font` | **required** | 100% | font_chat_text |
| `maxchars` | **required** | 100% | 4096 |
| `maxentries` | **required** | 100% | 256 |
| `scrollbar` | **required** | 100% | $parentScrollbar |
| `scrollbarPosition` | **required** | 100% | right |
## Structural Sub-Elements

### [Windows](element_Windows.md)

Observed 1 times as an unnamed child.

## Seen In

- Shinies

## Examples

- Shinies: ShiniesAutoUI_AutoSummary -> LogDisplay ShiniesAutoUI_AutoSummary

## Related APIs

- [Windows](element_Windows.md) (HIGH 100/100) - XML Element Type

## Used With

- none

## Triggered By

- none

## Affects

- none
