# WindowClearAnchors

- Category: Window Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 4 addons

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
| Addons seen in | InfoScroller, Moth, PartyCast, TidyChat |
| Files seen in | `/workspace/data/raw/InfoScroller/InfoScroller.lua:95`, `/workspace/data/raw/InfoScroller/libs/LibGUI.lua:152`, `/workspace/data/raw/Moth/Moth.lua:205`, `/workspace/data/raw/Moth/Moth.lua:575`, `/workspace/data/raw/PartyCast/PartyCast.lua:655`, `/workspace/data/raw/PartyCast/libs/LibGUI.lua:152`, `/workspace/data/raw/TidyChat/TidyChat.lua:239`, `/workspace/data/raw/TidyChat/TidyChat.lua:344` |
| Namespaces detected | WindowClearAnchors |
| Source kinds | lua_calls |
| Example locations | InfoScroller: InfoScroller.CreateCard, InfoScroller: LIBGUI_ELEMENT:ClearAnchors, Moth: Moth.Anchor, Moth: Moth.HealthBar, PartyCast: LIBGUI_ELEMENT:ClearAnchors, PartyCast: PartyCast.Update |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 21 |
| Global usage count | 21 |
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
WindowClearAnchors(windowName)
```

## Description

Observed resetting a window layout before applying new runtime anchors.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| windowName | Observed as a target window name. | Observed values: "Moth", "MothHealthBar", "PartyCastWindow"..i |

## Returns

- Not confidently inferable from addon-api docs alone.

## Side Effects

- Mutates runtime window layout state.

## Seen In

- InfoScroller
- Moth
- PartyCast
- TidyChat

## Examples

- InfoScroller: InfoScroller.CreateCard -> WindowClearAnchors(WindowName.."Image")
- InfoScroller: LIBGUI_ELEMENT:ClearAnchors -> WindowClearAnchors(self.name)
- Moth: Moth.Anchor -> WindowClearAnchors("Moth")
- Moth: Moth.HealthBar -> WindowClearAnchors("MothHealthBar")
- PartyCast: LIBGUI_ELEMENT:ClearAnchors -> WindowClearAnchors(self.name)
- PartyCast: PartyCast.Update -> WindowClearAnchors("PartyCastWindow"..i)

## Related APIs

- none

## Used With

- [LayoutEditor.RegisterWindow](window_LayoutEditor.RegisterWindow.md) (HIGH 100/100) - Window Function
- [StatusBarGetCurrentValue](window_StatusBarGetCurrentValue.md) (HIGH 100/100) - Window Function
- [WindowAddAnchor](window_WindowAddAnchor.md) (HIGH 100/100) - Window Function
- [WindowGetAlpha](window_WindowGetAlpha.md) (HIGH 100/100) - Window Function
- [WindowGetShowing](window_WindowGetShowing.md) (HIGH 100/100) - Window Function
- [WindowSetHandleInput](window_WindowSetHandleInput.md) (HIGH 100/100) - Window Function
- [WindowSetShowing](window_WindowSetShowing.md) (HIGH 100/100) - Window Function
- [DoesWindowExist](../../globals/functions/global_DoesWindowExist.md) (HIGH 83/100) - Global Function
- [StatusBarGetMaximumValue](window_StatusBarGetMaximumValue.md) (HIGH 80/100) - Window Function
- [CreateWindow](../../globals/functions/global_CreateWindow.md) (HIGH 75/100) - Global Function

## Triggered By

- none

## Affects

- [InterfaceCore.GetScale](../../globals/functions/global_InterfaceCore.GetScale.md) (HIGH 100/100) - Global Function
- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Notes

- none
