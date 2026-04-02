# WindowSetFontAlpha

- Category: Window Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 9 addons

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
| Addons seen in | Ace, DAoCBuff, Effigy, GCDsaver, LibWBToggler, Shinies, TurretRange, WSCT |
| Files seen in | `/workspace_addons/Ace/LibGUI.lua:213`, `/workspace_addons/DAoCBuff/Source/DAoCBuffFrames.lua:40`, `/workspace_addons/Effigy/LibGUI.lua:213`, `/workspace_addons/GCDsaver/libs/LibGUI.lua:213`, `/workspace_addons/LibWarBoardToggler/libs/LibGUI.lua:213`, `/workspace_addons/Shinies/Libraries/LibGUI.lua:213`, `/workspace_addons/TurrentRange/Display.lua:338`, `/workspace_addons/TurrentRange/Setup/SetupDisplay.lua:100` |
| Namespaces detected | WindowSetFontAlpha |
| Source kinds | lua_calls |
| Example locations | Ace: LIBGUI_ELEMENT:FontAlpha, DAoCBuff: DAoCBuffFrame:Create, Effigy: LIBGUI_ELEMENT:FontAlpha, GCDsaver: LIBGUI_ELEMENT:FontAlpha, LibWBToggler: LIBGUI_ELEMENT:FontAlpha, Shinies: LIBGUI_ELEMENT:FontAlpha |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 15 |
| Global usage count | 15 |
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
WindowSetFontAlpha(windowName, arg2)
```

## Description

Observed mutating runtime window state or presentation.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| windowName | Observed as a target window name. | Observed values: adat.textname, self.name, windowName |
| arg2 | Observed as a runtime window or control identifier. | Observed values: 1.0, adat.alpha, alpha |

## Returns

- Not confidently inferable from addon-api docs alone.

## Side Effects

- Mutates engine-visible UI state or presentation.

## Seen In

- Ace
- DAoCBuff
- Effigy
- GCDsaver
- LibWBToggler
- Shinies
- TurretRange
- WSCT
- WoH-Reticle

## Examples

- Ace: LIBGUI_ELEMENT:FontAlpha -> WindowSetFontAlpha(self.name, alpha)
- DAoCBuff: DAoCBuffFrame:Create -> WindowSetFontAlpha(windowName, 1.0)
- Effigy: LIBGUI_ELEMENT:FontAlpha -> WindowSetFontAlpha(self.name, alpha)
- GCDsaver: LIBGUI_ELEMENT:FontAlpha -> WindowSetFontAlpha(self.name, alpha)
- LibWBToggler: LIBGUI_ELEMENT:FontAlpha -> WindowSetFontAlpha(self.name, alpha)
- Shinies: LIBGUI_ELEMENT:FontAlpha -> WindowSetFontAlpha(self.name, alpha)

## Related APIs

- none

## Used With

- [WindowGetFontAlpha](window_WindowGetFontAlpha.md) (HIGH 100/100) - Window Function

## Triggered By

- none

## Affects

- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Notes

- none
