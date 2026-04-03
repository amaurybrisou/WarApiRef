# WindowGetAnchor

- Category: Window Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 7 addons

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
| Addons seen in | Ace, BuffHead, LibWBToggler, PartyCast, Pocket Palette, Shinies, WoH-Reticle |
| Files seen in | `/workspace/data/raw/Ace/LibGUI.lua:172`, `/workspace/data/raw/BuffHead/Setup/LayoutFrame.lua:105`, `/workspace/data/raw/LibWarBoardToggler/libs/LibGUI.lua:172`, `/workspace/data/raw/PartyCast/libs/LibGUI.lua:172`, `/workspace/data/raw/PocketPalette/PocketPalette.lua:152`, `/workspace/data/raw/Shinies/Libraries/LibGUI.lua:172`, `/workspace/data/raw/Shinies/Modules/Aggregator/Shinies-Aggregator-Tooltip.lua:452`, `/workspace/data/raw/WoH-Reticle/libs/LibGUI.lua:172` |
| Namespaces detected | WindowGetAnchor |
| Source kinds | lua_calls |
| Example locations | Ace: LIBGUI_ELEMENT:GetPosition, BuffHead: BuffHead.Setup.LayoutFrame:UpdatePosition, LibWBToggler: LIBGUI_ELEMENT:GetPosition, PartyCast: LIBGUI_ELEMENT:GetPosition, Pocket Palette: PP.CreateWindow, Shinies: LIBGUI_ELEMENT:GetPosition |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 8 |
| Global usage count | 8 |
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
| windowName | Observed as a target window name. | Observed values: PP.settings.windows["main"].name, self.name, self:GetName() |
| arg2 | Observed as a numeric value. | Observed values: 1 |

## Returns

- Not confidently inferable from addon-api docs alone.

## Side Effects

- No side effect is confidently inferable from addon-api docs alone.

## Seen In

- Ace
- BuffHead
- LibWBToggler
- PartyCast
- Pocket Palette
- Shinies
- WoH-Reticle

## Examples

- Ace: LIBGUI_ELEMENT:GetPosition -> WindowGetAnchor(self.name, 1)
- BuffHead: BuffHead.Setup.LayoutFrame:UpdatePosition -> WindowGetAnchor(self:GetName(), 1)
- LibWBToggler: LIBGUI_ELEMENT:GetPosition -> WindowGetAnchor(self.name, 1)
- PartyCast: LIBGUI_ELEMENT:GetPosition -> WindowGetAnchor(self.name, 1)
- Pocket Palette: PP.CreateWindow -> WindowGetAnchor(PP.settings.windows["main"].name, 1)
- Shinies: LIBGUI_ELEMENT:GetPosition -> WindowGetAnchor(self.name, 1)

## Related APIs

- none

## Used With

- none

## Triggered By

- none

## Affects

- [InterfaceCore.GetScale](../../globals/functions/global_InterfaceCore.GetScale.md) (HIGH 100/100) - Global Function
- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Notes

- none
