# WindowSetDimensions

- Category: Window Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 6 addons

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 110

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, seen in 4 or more addons, matches a known engine namespace.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- -15 Conflicting signatures across usages: Observed arity or argument shape conflicts across usages.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | InfoScroller, Moth, PartyCast, TidyChat, TidyRoll, minesweep |
| Files seen in | `/workspace/data/raw/InfoScroller/InfoScroller.lua:95`, `/workspace/data/raw/InfoScroller/libs/LibGUI.lua:1171`, `/workspace/data/raw/InfoScroller/libs/LibGUI.lua:1258`, `/workspace/data/raw/InfoScroller/libs/LibGUI.lua:367`, `/workspace/data/raw/InfoScroller/libs/LibGUI.lua:419`, `/workspace/data/raw/InfoScroller/libs/LibGUI.lua:527`, `/workspace/data/raw/InfoScroller/libs/LibGUI.lua:649`, `/workspace/data/raw/InfoScroller/libs/LibGUI.lua:703` |
| Namespaces detected | WindowSetDimensions |
| Source kinds | lua_calls |
| Example locations | InfoScroller: InfoScroller.CreateCard, InfoScroller: LIBGUI_Button:Resize, InfoScroller: LIBGUI_Image:Resize, InfoScroller: LIBGUI_Label:Resize, InfoScroller: LIBGUI_MultiTextbox:Resize, InfoScroller: LIBGUI_Scrollbar:Resize |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 32 |
| Global usage count | 32 |
| Local definition count | 0 |
| Documentation references | 0 |
| Initialization flow references | 0 |
| Known engine namespace | yes |
| Default UI presence | yes |
| Event binding presence | no |
| Observed in XML and Lua | no |
| Consistent role | yes |
| Consistent arguments | no |
| Consistent returns | no |
| Slash command presence | no |
| Weak usage only | no |
| Project-specific name | no |
| Placeholder or computed name | no |
| Conflicting signatures | yes |
| Conflicting roles | no |
| Wrapper likely | no |
| Never outside local graph | no |
| Local helper only | no |

## Signature (inferred)

```lua
WindowSetDimensions(windowName, arg2, arg3)
```

## Description

Observed mutating runtime window state or presentation.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| windowName | Observed as a target window name. | Observed values: "MineSweepWindow", "MothHealthBar", WindowName.."Image" |
| arg2 | Observed as a runtime window or control identifier. | Observed values: 10+(fieldX*28), 140, 32 |
| arg3 | Observed as a function or method reference. | Observed values: 20, 30, 32 |

## Returns

- Not confidently inferable from addon-api docs alone.

## Side Effects

- Mutates engine-visible UI state or presentation.

## Seen In

- InfoScroller
- Moth
- PartyCast
- TidyChat
- TidyRoll
- minesweep

## Examples

- InfoScroller: InfoScroller.CreateCard -> WindowSetDimensions(WindowName.."Image", width*scale, height*scale)
- InfoScroller: LIBGUI_Button:Resize -> WindowSetDimensions(self.name, width, self.height)
- InfoScroller: LIBGUI_Image:Resize -> WindowSetDimensions(self.name, width, height)
- InfoScroller: LIBGUI_Label:Resize -> WindowSetDimensions(self.name, self.width, self.height)
- InfoScroller: LIBGUI_MultiTextbox:Resize -> WindowSetDimensions(self.name, self.width, self.height)
- InfoScroller: LIBGUI_Scrollbar:Resize -> WindowSetDimensions(self.name, self.width, self.height)

## Related APIs

- none

## Used With

- [LabelSetFont](window_LabelSetFont.md) (HIGH 100/100) - Window Function
- [LabelSetText](window_LabelSetText.md) (HIGH 100/100) - Window Function
- [LabelSetTextColor](window_LabelSetTextColor.md) (HIGH 100/100) - Window Function
- [WindowGetDimensions](window_WindowGetDimensions.md) (HIGH 100/100) - Window Function
- [WindowSetShowing](window_WindowSetShowing.md) (HIGH 100/100) - Window Function
- [WindowSetTintColor](window_WindowSetTintColor.md) (HIGH 100/100) - Window Function
- [LabelGetTextDimensions](window_LabelGetTextDimensions.md) (HIGH 80/100) - Window Function

## Triggered By

- none

## Affects

- [InterfaceCore.GetScale](../../globals/functions/global_InterfaceCore.GetScale.md) (HIGH 100/100) - Global Function
- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Notes

- none
