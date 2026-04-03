# HelpTips.SetFocusOnWindow

- Category: Global Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 4 addons

## Confidence Assessment

- Level: HIGH

- Score: 100/100

- Rationale: Promoted as HIGH confidence because seen in 4 or more addons, matches a known engine namespace, called globally with no local definition.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +10 Argument pattern is consistent: Observed argument positions remain stable.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | Calling, DuffTimer, EA_UiDebugTools, Pocket Palette |
| Files seen in | CallingNotification.lua, DuffTimer.lua, PocketPalette.lua, Source/Debug.lua |
| Namespaces detected | HelpTips |
| Source kinds | lua_calls |
| Example locations | Calling: HighlightWindow, DuffTimer: HighlightWindow, EA_UiDebugTools: hw, Pocket Palette: HighlightWindow |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 4 |
| Global usage count | 4 |
| Local definition count | 0 |
| Documentation references | 0 |
| Initialization flow references | 0 |
| Known engine namespace | yes |
| Default UI presence | no |
| Event binding presence | no |
| Observed in XML and Lua | no |
| Consistent role | yes |
| Consistent arguments | yes |
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

## Signature (inferred)

```lua
HelpTips.SetFocusOnWindow(arg1)
```

## Description

Observed as a global function across 4 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a runtime window or control identifier. | Observed values: (text), window, window_name |

## Returns

- Not confidently inferable from contract artifacts alone.

## Side Effects

- Mutates engine-visible UI state or presentation.

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

- Canonical entry built from observed call sites, not from engine source or decompiled definitions.
