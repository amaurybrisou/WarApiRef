# WindowAddAnchor

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
| Addons seen in | InfoScroller, Moth, PartyCast, TidyChat, TidyRoll |
| Files seen in | `/workspace/data/raw/InfoScroller/InfoScroller.lua:95`, `/workspace/data/raw/InfoScroller/libs/LibGUI.lua:140`, `/workspace/data/raw/Moth/Moth.lua:205`, `/workspace/data/raw/Moth/Moth.lua:575`, `/workspace/data/raw/PartyCast/PartyCast.lua:655`, `/workspace/data/raw/PartyCast/libs/LibGUI.lua:140`, `/workspace/data/raw/TidyChat/TidyChat.lua:239`, `/workspace/data/raw/TidyChat/TidyChat.lua:344` |
| Namespaces detected | WindowAddAnchor |
| Source kinds | lua_calls |
| Example locations | InfoScroller: InfoScroller.CreateCard, InfoScroller: LIBGUI_ELEMENT:AddAnchor, Moth: Moth.Anchor, Moth: Moth.HealthBar, PartyCast: LIBGUI_ELEMENT:AddAnchor, PartyCast: PartyCast.Update |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 96 |
| Global usage count | 96 |
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
WindowAddAnchor(windowName, point, relativeTo, relativePoint, offsetX, offsetY)
```

## Description

Observed positioning windows relative to other runtime UI elements.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| windowName | Observed as the window being positioned. | Observed values: "Moth", "MothHealthBar", "PartyCastWindow"..i |
| point | Observed as an anchor point string. | Observed values: "bottom", "bottomleft", "bottomright" |
| relativeTo | Observed as a reference window name. | Observed values: "CursorWindow", "MothBackground", "PartyCastWindow"..i |
| relativePoint | Observed as a reference anchor point string. | Observed values: "bottom", "bottomleft", "bottomright" |
| offsetX | Observed as a numeric horizontal offset. | Observed values: -10, -2, -24 |
| offsetY | Observed as a numeric vertical offset. | Observed values: -(55-(50*perc_comp)), -10, -2 |

## Returns

- Not confidently inferable from addon-api docs alone.

## Side Effects

- Mutates runtime window layout state.

## Seen In

- InfoScroller
- Moth
- PartyCast
- TidyChat
- TidyRoll

## Examples

- InfoScroller: InfoScroller.CreateCard -> WindowAddAnchor(WindowName.."Image", anchor, WindowName, anchor, T_X, T_Y)
- InfoScroller: LIBGUI_ELEMENT:AddAnchor -> WindowAddAnchor(self.name, pointOnAnchor, anchorWindow, pointOnSelf, xOffset, yOffset)
- Moth: Moth.Anchor -> WindowAddAnchor("Moth", "bottomleft", "CursorWindow", "topleft", 0, 0)
- Moth: Moth.HealthBar -> WindowAddAnchor("MothHealthBar", "bottomleft", "MothBackground", "topleft", 0, yOffset)
- PartyCast: LIBGUI_ELEMENT:AddAnchor -> WindowAddAnchor(self.name, pointOnAnchor, anchorWindow, pointOnSelf, xOffset, yOffset)
- PartyCast: PartyCast.Update -> WindowAddAnchor("PartyCastWindow"..i, Frame_Anchor, "PartyCastWindow_Dynamic"..i, Frame_Anchor, 0, PartyCast.Settings.Offset)

## Related APIs

- none

## Used With

- [LayoutEditor.RegisterWindow](window_LayoutEditor.RegisterWindow.md) (HIGH 100/100) - Window Function
- [StatusBarGetCurrentValue](window_StatusBarGetCurrentValue.md) (HIGH 100/100) - Window Function
- [WindowClearAnchors](window_WindowClearAnchors.md) (HIGH 100/100) - Window Function
- [WindowGetAlpha](window_WindowGetAlpha.md) (HIGH 100/100) - Window Function
- [WindowGetShowing](window_WindowGetShowing.md) (HIGH 100/100) - Window Function
- [WindowSetHandleInput](window_WindowSetHandleInput.md) (HIGH 100/100) - Window Function
- [WindowSetShowing](window_WindowSetShowing.md) (HIGH 100/100) - Window Function
- [DoesWindowExist](../../globals/functions/global_DoesWindowExist.md) (HIGH 83/100) - Global Function
- [StatusBarGetMaximumValue](window_StatusBarGetMaximumValue.md) (HIGH 80/100) - Window Function

## Triggered By

- none

## Affects

- [InterfaceCore.GetScale](../../globals/functions/global_InterfaceCore.GetScale.md) (HIGH 100/100) - Global Function
- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Notes

- none
