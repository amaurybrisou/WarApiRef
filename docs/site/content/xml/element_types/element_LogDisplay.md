# LogDisplay

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
| Addons seen in | EA_UiDebugTools, Shinies |
| Files seen in | `/workspace_addons/Shinies/Modules/UI/Shinies-UI-Auto.xml:182`, `/workspace_addons/ea_uidebugtools/Source/DebugWindow.xml:102` |
| Namespaces detected | LogDisplay |
| Source kinds | xml_frames |
| Example locations | EA_UiDebugTools: DebugWindowText, Shinies: ShiniesAutoUI_AutoSummary |
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

Observed XML element type instantiated by 2 addons.

## Common Attributes

- background
- font
- maxchars
- maxentries
- name
- scrollbar
- scrollbarPosition
- autoHideScrollBar
- handleinput
- textFadeTime

## Common Inherits

- none

## Common Parent Elements

- [Window](element_Window.md)

## Common Named Child Elements

- [VerticalScrollbar](element_VerticalScrollbar.md)
- [Button](element_Button.md)


## Attribute Reference

| Attribute | Role | Observed Values | Count |
|-----------|------|-----------------|-------|
| `background` | string | - | 2 |
| `font` | string | `font_clear_small_bold`, `font_chat_text` | 2 |
| `maxchars` | number | `4096` | 2 |
| `maxentries` | number | `-1`, `256` | 2 |
| `name` | string | `DebugWindowText`, `ShiniesAutoUI_AutoSummary` | 2 |
| `scrollbar` | frame-ref | `$parentScrollbar` | 2 |
| `scrollbarPosition` | string | `right` | 2 |
| `autoHideScrollBar` | boolean | `true` | 1 |
| `handleinput` | boolean | `true` | 1 |
| `textFadeTime` | number | `30` | 1 |

## Seen In

- EA_UiDebugTools
- Shinies

## Examples

- EA_UiDebugTools: DebugWindowText -> LogDisplay DebugWindowText
- Shinies: ShiniesAutoUI_AutoSummary -> LogDisplay ShiniesAutoUI_AutoSummary

## Related APIs

- none

## Used With

- none

## Triggered By

- none

## Affects

- none
