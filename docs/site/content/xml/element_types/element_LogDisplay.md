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
| Files seen in | `/workspace/Shinies/Modules/UI/Shinies-UI-Auto.xml:182`, `/workspace/ea_uidebugtools/Source/DebugWindow.xml:102` |
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

## Common Handlers

- none

## Common Inherits

- none

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
