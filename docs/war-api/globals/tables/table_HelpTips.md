# HelpTips

- Category: Global Table
- Confidence level: HIGH
- Confidence score: 100/100

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 125

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, seen in 4 or more addons, matches a known engine namespace.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | Calling, DuffTimer, EA_UiDebugTools, Pocket Palette |
| Files seen in | Source/Debug.lua |
| Namespaces detected | HelpTips |
| Source kinds | lua_calls |
| Example locations | Calling: HighlightWindow, DuffTimer: HighlightWindow, EA_UiDebugTools: hw, Pocket Palette: HighlightWindow |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 4 |
| Global usage count | 1 |
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

Shared function table with 1 member functions; the primary API surface for 4 addons.

## Functions

- HelpTips.SetFocusOnWindow

## Observed Members

- none

## Seen In

- Calling
- DuffTimer
- EA_UiDebugTools
- Pocket Palette

## Examples

- Calling: HighlightWindow -> HelpTips.SetFocusOnWindow(window)
- DuffTimer: HighlightWindow -> HelpTips.SetFocusOnWindow(window)
- EA_UiDebugTools: hw -> HelpTips.SetFocusOnWindow((text))
- Pocket Palette: HighlightWindow -> HelpTips.SetFocusOnWindow(window_name)

## Notes

- none
