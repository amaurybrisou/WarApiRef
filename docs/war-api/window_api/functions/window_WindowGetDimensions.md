# WindowGetDimensions

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
| Files seen in | `/workspace/data/raw/InfoScroller/libs/LibGUI.lua:1043`, `/workspace/data/raw/InfoScroller/libs/LibGUI.lua:1146`, `/workspace/data/raw/InfoScroller/libs/LibGUI.lua:1234`, `/workspace/data/raw/InfoScroller/libs/LibGUI.lua:506`, `/workspace/data/raw/InfoScroller/libs/LibGUI.lua:619`, `/workspace/data/raw/InfoScroller/libs/LibGUI.lua:679`, `/workspace/data/raw/InfoScroller/libs/LibGUI.lua:734`, `/workspace/data/raw/InfoScroller/libs/LibGUI.lua:803` |
| Namespaces detected | WindowGetDimensions |
| Source kinds | lua_calls |
| Example locations | InfoScroller: LIBGUI_Button:New, InfoScroller: LIBGUI_Checkbox:New, InfoScroller: LIBGUI_Combobox:New, InfoScroller: LIBGUI_Image:New, InfoScroller: LIBGUI_MultiTextbox:New, InfoScroller: LIBGUI_Optionbutton:New |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 23 |
| Global usage count | 23 |
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
WindowGetDimensions(windowName)
```

## Description

Observed querying runtime window state or metadata.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| windowName | Observed as a target window name. | Observed values: "MothBackground", c_CHANNEL_MENU, cellName.."NPCIcon" |

## Returns

- Not confidently inferable from addon-api docs alone.

## Side Effects

- No side effect is confidently inferable from addon-api docs alone.

## Seen In

- InfoScroller
- Moth
- PartyCast
- TidyChat

## Examples

- InfoScroller: LIBGUI_Button:New -> WindowGetDimensions(w.name)
- InfoScroller: LIBGUI_Checkbox:New -> WindowGetDimensions(w.name)
- InfoScroller: LIBGUI_Combobox:New -> WindowGetDimensions(w.name)
- InfoScroller: LIBGUI_Image:New -> WindowGetDimensions(w.name)
- InfoScroller: LIBGUI_MultiTextbox:New -> WindowGetDimensions(w.name)
- InfoScroller: LIBGUI_Optionbutton:New -> WindowGetDimensions(w.name)

## Related APIs

- none

## Used With

- [ButtonSetCheckButtonFlag](window_ButtonSetCheckButtonFlag.md) (HIGH 100/100) - Window Function
- [WindowSetDimensions](window_WindowSetDimensions.md) (HIGH 100/100) - Window Function
- [WindowSetShowing](window_WindowSetShowing.md) (HIGH 100/100) - Window Function
- [WindowSetTintColor](window_WindowSetTintColor.md) (HIGH 100/100) - Window Function
- [CreateWindowFromTemplate](../../globals/functions/global_CreateWindowFromTemplate.md) (HIGH 75/100) - Global Function

## Triggered By

- none

## Affects

- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Notes

- none
