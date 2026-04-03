# WindowSetAlpha

- Category: Window Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 3 addons

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 123

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, matches a known engine namespace, called globally with no local definition.

## Evidence Signals

- +18 Seen in 2 to 3 addons: Cross-addon spread is present but limited.
- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +10 Argument pattern is consistent: Observed argument positions remain stable.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | InfoScroller, PartyCast, TidyChat |
| Files seen in | `/workspace/data/raw/InfoScroller/InfoScroller.lua:95`, `/workspace/data/raw/InfoScroller/libs/LibGUI.lua:200`, `/workspace/data/raw/PartyCast/PartyCast.lua:399`, `/workspace/data/raw/PartyCast/PartyCast.lua:655`, `/workspace/data/raw/PartyCast/libs/LibGUI.lua:200`, `/workspace/data/raw/TidyChat/TidyChat.lua:329`, `/workspace/data/raw/TidyChat/TidyChat.lua:344`, `/workspace/data/raw/TidyChat/TidyChat.lua:404` |
| Namespaces detected | WindowSetAlpha |
| Source kinds | lua_calls |
| Example locations | InfoScroller: InfoScroller.CreateCard, InfoScroller: LIBGUI_ELEMENT:Alpha, PartyCast: LIBGUI_ELEMENT:Alpha, PartyCast: PartyCast.FetchedText, PartyCast: PartyCast.Update, TidyChat: TidyChatCore.SetTextEntry |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 24 |
| Global usage count | 24 |
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
WindowSetAlpha(windowName, arg2)
```

## Description

Observed mutating runtime window state or presentation.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| windowName | Observed as a target window name. | Observed values: "PartyCastWindow"..PlayerNumber, "PartyCastWindow"..PlayerNumber.."Timer", "PartyCastWindow"..i |
| arg2 | Observed as a runtime window or control identifier. | Observed values: (chatwindow_tabs_handle_input~=false and 1)or 0, (scrollbar_position==c_SCROLLBAR_POSITION_HIDDEN and 0)or scrollbar_alpha, 0 |

## Returns

- Not confidently inferable from addon-api docs alone.

## Side Effects

- Mutates engine-visible UI state or presentation.

## Seen In

- InfoScroller
- PartyCast
- TidyChat

## Examples

- InfoScroller: InfoScroller.CreateCard -> WindowSetAlpha(WindowName.."BackGroundBG", InfoScroller.Settings.BGAlpha)
- InfoScroller: InfoScroller.CreateCard -> WindowSetAlpha(WindowName.."BackGroundStart", InfoScroller.Settings.BGAlpha)
- InfoScroller: InfoScroller.CreateCard -> WindowSetAlpha(WindowName.."BackGroundEnd", InfoScroller.Settings.BGAlpha)
- InfoScroller: InfoScroller.CreateCard -> WindowSetAlpha(WindowName.."Image", alpha)
- InfoScroller: LIBGUI_ELEMENT:Alpha -> WindowSetAlpha(self.name, alpha)
- PartyCast: LIBGUI_ELEMENT:Alpha -> WindowSetAlpha(self.name, alpha)

## Related APIs

- none

## Used With

- [WindowGetAlpha](window_WindowGetAlpha.md) (HIGH 100/100) - Window Function

## Triggered By

- none

## Affects

- [InterfaceCore.GetScale](../../globals/functions/global_InterfaceCore.GetScale.md) (HIGH 100/100) - Global Function
- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Notes

- none
