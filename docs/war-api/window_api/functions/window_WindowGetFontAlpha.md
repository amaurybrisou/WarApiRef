# WindowGetFontAlpha

- Category: Window Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 5 addons

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 135

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, seen in 4 or more addons, matches a known engine namespace.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +10 Argument pattern is consistent: Observed argument positions remain stable.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | Ace, LibWBToggler, PartyCast, Shinies, WoH-Reticle |
| Files seen in | `/workspace/data/raw/Ace/LibGUI.lua:213`, `/workspace/data/raw/LibWarBoardToggler/libs/LibGUI.lua:213`, `/workspace/data/raw/PartyCast/libs/LibGUI.lua:213`, `/workspace/data/raw/Shinies/Libraries/LibGUI.lua:213`, `/workspace/data/raw/WoH-Reticle/libs/LibGUI.lua:213` |
| Namespaces detected | WindowGetFontAlpha |
| Source kinds | lua_calls |
| Example locations | Ace: LIBGUI_ELEMENT:FontAlpha, LibWBToggler: LIBGUI_ELEMENT:FontAlpha, PartyCast: LIBGUI_ELEMENT:FontAlpha, Shinies: LIBGUI_ELEMENT:FontAlpha, WoH-Reticle: LIBGUI_ELEMENT:FontAlpha |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 5 |
| Global usage count | 5 |
| Local definition count | 0 |
| Documentation references | 0 |
| Initialization flow references | 0 |
| Known engine namespace | yes |
| Default UI presence | yes |
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
WindowGetFontAlpha(windowName)
```

## Description

Observed querying runtime window state or metadata.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| windowName | Observed as a target window name. | Observed values: self.name |

## Returns

- Not confidently inferable from addon-api docs alone.

## Side Effects

- No side effect is confidently inferable from addon-api docs alone.

## Seen In

- Ace
- LibWBToggler
- PartyCast
- Shinies
- WoH-Reticle

## Examples

- Ace: LIBGUI_ELEMENT:FontAlpha -> WindowGetFontAlpha(self.name)
- LibWBToggler: LIBGUI_ELEMENT:FontAlpha -> WindowGetFontAlpha(self.name)
- PartyCast: LIBGUI_ELEMENT:FontAlpha -> WindowGetFontAlpha(self.name)
- Shinies: LIBGUI_ELEMENT:FontAlpha -> WindowGetFontAlpha(self.name)
- WoH-Reticle: LIBGUI_ELEMENT:FontAlpha -> WindowGetFontAlpha(self.name)

## Related APIs

- none

## Used With

- [WindowSetFontAlpha](window_WindowSetFontAlpha.md) (HIGH 100/100) - Window Function

## Triggered By

- none

## Affects

- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Notes

- none
