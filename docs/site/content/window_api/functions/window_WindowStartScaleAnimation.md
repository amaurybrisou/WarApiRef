# WindowStartScaleAnimation

- Category: Window Function
- Confidence level: HIGH
- Confidence score: 90/100
- Seen in: 1 addons

## Confidence Assessment

- Level: HIGH

- Score: 90/100

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, matches a known engine namespace, called globally with no local definition.

## Evidence Signals

- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +10 Argument pattern is consistent: Observed argument positions remain stable.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | MapPin |
| Files seen in | `/workspace/MapPin/source/MapPin.lua:1359`, `/workspace/MapPin/source/MapPin.lua:2861`, `/workspace/MapPin/source/MapPin.lua:2887` |
| Namespaces detected | WindowStartScaleAnimation |
| Source kinds | lua_calls |
| Example locations | MapPin: MapPin.CreateMainContextMenu, MapPin: MapPin.Map.UpdatePinCoordinates, MapPin: MapPin.local.SendCommand, MapPin: SendCommand |
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
| Consistent role | no |
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
WindowStartScaleAnimation(arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8)
```

## Description

Observed as a window function across 1 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a text or wstring payload. | Observed values: "MapPinNotePin"..WindowsNumber, "MapPinRequestPin"..RequestWindowsNumber, "MapPinWayPointPin"..MapPin.WayPoints |
| arg2 | Observed as a numeric value. | Observed values: 2 |
| arg3 | Observed as a function or method reference. | Observed values: 2.5*Scale |
| arg4 | Observed as a runtime window or control identifier. | Observed values: 1*Scale, MapPin.Pins.Notes[EA_Window_WorldMap.currentMap][WindowsNumber].PinScale*Scale, MapPin.Pins.Request[RequestWindowsNumber].PinScale*Scale |
| arg5 | Observed as a function or method reference. | Observed values: 0.5 |
| arg6 | Observed as a boolean toggle. | Observed values: true |
| arg7 | Observed as a function or method reference. | Observed values: 0.0 |
| arg8 | Observed as a numeric value. | Observed values: 0 |

## Returns

- Not confidently inferable from API_Ref alone.

## Side Effects

- No side effect is confidently inferable from API_Ref alone.

## Seen In

- MapPin

## Examples

- MapPin: MapPin.CreateMainContextMenu -> WindowStartScaleAnimation("MapPinWayPointPin"..MapPin.WayPoints, 2, 2.5*Scale, 1*Scale, 0.5, true, 0.0, 0)
- MapPin: MapPin.Map.UpdatePinCoordinates -> WindowStartScaleAnimation("MapPinNotePin"..WindowsNumber, 2, 2.5*Scale, MapPin.Pins.Notes[EA_Window_WorldMap.currentMap][WindowsNumber].PinScale*Scale, 0.5, true, 0.0, 0)
- MapPin: MapPin.Map.UpdatePinCoordinates -> WindowStartScaleAnimation("MapPinRequestPin"..RequestWindowsNumber, 2, 2.5*Scale, MapPin.Pins.Request[RequestWindowsNumber].PinScale*Scale, 0.5, true, 0.0, 0)
- MapPin: MapPin.local.SendCommand -> WindowStartScaleAnimation("MapPinWayPointPin"..MapPin.WayPoints, 2, 2.5*Scale, 1*Scale, 0.5, true, 0.0, 0)
- MapPin: SendCommand -> WindowStartScaleAnimation("MapPinWayPointPin"..MapPin.WayPoints, 2, 2.5*Scale, 1*Scale, 0.5, true, 0.0, 0)

## Related APIs

- none

## Used With

- [EA_Window_ContextMenu.AddCascadingMenuItem](../../globals/functions/global_EA_Window_ContextMenu.AddCascadingMenuItem.md) (HIGH 100/100) - Global Function
- [EA_Window_ContextMenu.AddMenuDivider](../../globals/functions/global_EA_Window_ContextMenu.AddMenuDivider.md) (HIGH 100/100) - Global Function
- [EA_Window_ContextMenu.AddMenuItem](../../globals/functions/global_EA_Window_ContextMenu.AddMenuItem.md) (HIGH 100/100) - Global Function
- [EA_Window_ContextMenu.CreateContextMenu](../../globals/functions/global_EA_Window_ContextMenu.CreateContextMenu.md) (HIGH 100/100) - Global Function
- [EA_Window_ContextMenu.Finalize](../../globals/functions/global_EA_Window_ContextMenu.Finalize.md) (HIGH 100/100) - Global Function
- [PartyUtils.IsPartyActive](../../globals/functions/global_PartyUtils.IsPartyActive.md) (HIGH 100/100) - Global Function
- [WindowGetDimensions](window_WindowGetDimensions.md) (HIGH 100/100) - Window Function
- [WindowGetScale](window_WindowGetScale.md) (HIGH 100/100) - Window Function
- [WindowGetShowing](window_WindowGetShowing.md) (HIGH 100/100) - Window Function
- [WindowSetScale](window_WindowSetScale.md) (HIGH 100/100) - Window Function
- [WindowSetShowing](window_WindowSetShowing.md) (HIGH 100/100) - Window Function
- [WindowSetTintColor](window_WindowSetTintColor.md) (HIGH 100/100) - Window Function
- [WindowStartAlphaAnimation](window_WindowStartAlphaAnimation.md) (HIGH 100/100) - Window Function
- [DoesWindowExist](../../globals/functions/global_DoesWindowExist.md) (HIGH 83/100) - Global Function
- [GetIconData](../../globals/functions/global_GetIconData.md) (HIGH 83/100) - Global Function
- [towstring](../../globals/functions/global_towstring.md) (HIGH 75/100) - Global Function

## Triggered By

- none

## Affects

- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Notes

- Only one addon surfaced this symbol in the current corpus.
