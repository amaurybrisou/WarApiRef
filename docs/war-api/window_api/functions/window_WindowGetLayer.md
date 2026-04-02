# WindowGetLayer

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
| Addons seen in | Ace, BuffHead, Effigy, GCDsaver, LibWBToggler, Shinies, TexturedButtons, WoH-Reticle |
| Files seen in | `/workspace_addons/Ace/LibGUI.lua:178`, `/workspace_addons/BuffHead/Setup/SelectColor.lua:27`, `/workspace_addons/BuffHead/Setup/SelectTexture.lua:59`, `/workspace_addons/Effigy/LibGUI.lua:178`, `/workspace_addons/GCDsaver/libs/LibGUI.lua:178`, `/workspace_addons/LibWarBoardToggler/libs/LibGUI.lua:178`, `/workspace_addons/Shinies/Libraries/LibGUI.lua:178`, `/workspace_addons/TexturedButtons/Setup/SelectColor.lua:27` |
| Namespaces detected | WindowGetLayer |
| Source kinds | lua_calls |
| Example locations | Ace: LIBGUI_ELEMENT:Layer, BuffHead: BuffHead.Setup.SelectColor.Show, BuffHead: BuffHead.Setup.SelectTexture.Show, Effigy: LIBGUI_ELEMENT:Layer, GCDsaver: LIBGUI_ELEMENT:Layer, LibWBToggler: LIBGUI_ELEMENT:Layer |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 10 |
| Global usage count | 10 |
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
WindowGetLayer(windowName)
```

## Description

Observed querying runtime window state or metadata.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| windowName | Observed as a target window name. | Observed values: anchorName, self.name, window |

## Returns

- Not confidently inferable from addon-api docs alone.

## Side Effects

- No side effect is confidently inferable from addon-api docs alone.

## Seen In

- Ace
- BuffHead
- Effigy
- GCDsaver
- LibWBToggler
- Shinies
- TexturedButtons
- WoH-Reticle
- nRarity

## Examples

- Ace: LIBGUI_ELEMENT:Layer -> WindowGetLayer(self.name)
- BuffHead: BuffHead.Setup.SelectColor.Show -> WindowGetLayer(window)
- BuffHead: BuffHead.Setup.SelectTexture.Show -> WindowGetLayer(window)
- Effigy: LIBGUI_ELEMENT:Layer -> WindowGetLayer(self.name)
- GCDsaver: LIBGUI_ELEMENT:Layer -> WindowGetLayer(self.name)
- LibWBToggler: LIBGUI_ELEMENT:Layer -> WindowGetLayer(self.name)

## Related APIs

- none

## Used With

- [WindowClearAnchors](window_WindowClearAnchors.md) (HIGH 100/100) - Window Function
- [WindowGetShowing](window_WindowGetShowing.md) (HIGH 100/100) - Window Function
- [WindowSetLayer](window_WindowSetLayer.md) (HIGH 100/100) - Window Function

## Triggered By

- none

## Affects

- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Notes

- none
