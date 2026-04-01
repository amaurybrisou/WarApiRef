# WindowGetAnchor

- Category: Window Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 10 addons

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
| Addons seen in | Ace, BuffHead, Effigy, GCDsaver, LibWBToggler, Moth, Pocket Palette, Shinies |
| Files seen in | `/workspace/Ace/LibGUI.lua:172`, `/workspace/BuffHead/Setup/LayoutFrame.lua:105`, `/workspace/Effigy/LibGUI.lua:172`, `/workspace/GCDsaver/libs/LibGUI.lua:172`, `/workspace/LibWarBoardToggler/libs/LibGUI.lua:172`, `/workspace/Moth/Moth.lua:578`, `/workspace/PocketPalette/PocketPalette.lua:152`, `/workspace/Shinies/Libraries/LibGUI.lua:172` |
| Namespaces detected | WindowGetAnchor |
| Source kinds | lua_calls |
| Example locations | Ace: LIBGUI_ELEMENT:GetPosition, BuffHead: BuffHead.Setup.LayoutFrame:UpdatePosition, Effigy: LIBGUI_ELEMENT:GetPosition, GCDsaver: LIBGUI_ELEMENT:GetPosition, LibWBToggler: LIBGUI_ELEMENT:GetPosition, Moth: Moth.HealthBar |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 11 |
| Global usage count | 11 |
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
WindowGetAnchor(windowName, arg2)
```

## Description

Observed querying runtime window state or metadata.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| windowName | Observed as a target window name. | Observed values: "MothHealthBar", PP.settings.windows["main"].name, self.name |
| arg2 | Observed as a numeric value. | Observed values: 1 |

## Returns

- Not confidently inferable from API_Ref alone.

## Side Effects

- No side effect is confidently inferable from API_Ref alone.

## Seen In

- Ace
- BuffHead
- Effigy
- GCDsaver
- LibWBToggler
- Moth
- Pocket Palette
- Shinies
- WarTriage
- WoH-Reticle

## Examples

- Ace: LIBGUI_ELEMENT:GetPosition -> WindowGetAnchor(self.name, 1)
- BuffHead: BuffHead.Setup.LayoutFrame:UpdatePosition -> WindowGetAnchor(self:GetName(), 1)
- Effigy: LIBGUI_ELEMENT:GetPosition -> WindowGetAnchor(self.name, 1)
- GCDsaver: LIBGUI_ELEMENT:GetPosition -> WindowGetAnchor(self.name, 1)
- LibWBToggler: LIBGUI_ELEMENT:GetPosition -> WindowGetAnchor(self.name, 1)
- Moth: Moth.HealthBar -> WindowGetAnchor("MothHealthBar", 1)

## Related APIs

- none

## Used With

- [InterfaceCore.GetScale](../../globals/functions/global_InterfaceCore.GetScale.md) (HIGH 100/100) - Global Function
- [WindowAddAnchor](window_WindowAddAnchor.md) (HIGH 100/100) - Window Function
- [WindowClearAnchors](window_WindowClearAnchors.md) (HIGH 100/100) - Window Function
- [WindowSetShowing](window_WindowSetShowing.md) (HIGH 100/100) - Window Function

## Triggered By

- none

## Affects

- [InterfaceCore.GetScale](../../globals/functions/global_InterfaceCore.GetScale.md) (HIGH 100/100) - Global Function
- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Notes

- none
