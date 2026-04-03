# DestroyWindow

- Category: Global Function
- Confidence level: HIGH
- Confidence score: 75/100
- Seen in: 4 addons

## Confidence Assessment

- Level: HIGH

- Score: 75/100

- Rationale: Promoted as HIGH confidence because seen in 4 or more addons, called globally with no local definition, role is consistent across addons.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +10 Argument pattern is consistent: Observed argument positions remain stable.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | InfoScroller, PartyCast, TidyRoll, minesweep |
| Files seen in | `/workspace/data/raw/InfoScroller/libs/LibGUI.lua:120`, `/workspace/data/raw/PartyCast/libs/LibGUI.lua:120`, `/workspace/data/raw/TidyRoll/TidyRollFrame.lua:99`, `/workspace/data/raw/minesweep/minesweep.lua:11`, `/workspace/data/raw/minesweep/minesweep.lua:156` |
| Namespaces detected | DestroyWindow |
| Source kinds | lua_calls |
| Example locations | InfoScroller: LIBGUI_ELEMENT:Destroy, PartyCast: LIBGUI_ELEMENT:Destroy, TidyRoll: TidyRollFrame:Destroy, minesweep: minesweep.Close, minesweep: minesweep.MakeField |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 5 |
| Global usage count | 5 |
| Local definition count | 0 |
| Documentation references | 0 |
| Initialization flow references | 0 |
| Known engine namespace | no |
| Default UI presence | no |
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
DestroyWindow(windowName)
```

## Description

Observed tearing down runtime-created windows.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| windowName | Observed as a target window name. | Observed values: "MineSweepWindow", self.name, self:GetName() |

## Returns

- Not confidently inferable from addon-api docs alone.

## Side Effects

- Destroys a runtime-created window instance.

## Seen In

- InfoScroller
- PartyCast
- TidyRoll
- minesweep

## Examples

- InfoScroller: LIBGUI_ELEMENT:Destroy -> DestroyWindow(self.name)
- PartyCast: LIBGUI_ELEMENT:Destroy -> DestroyWindow(self.name)
- TidyRoll: TidyRollFrame:Destroy -> DestroyWindow(self:GetName())
- minesweep: minesweep.Close -> DestroyWindow("MineSweepWindow")
- minesweep: minesweep.MakeField -> DestroyWindow("MineSweepWindow")

## Related APIs

- none

## Used With

- none

## Triggered By

- [OnLButtonUp](../../xml/handlers/handler_OnLButtonUp.md) (HIGH 100/100) - XML Event
- [OnLButtonUp](../../events/window_events/window_event_OnLButtonUp.md) (HIGH 100/100) - Window Event

## Affects

- [Button](../../xml/element_types/element_Button.md) (HIGH 100/100) - XML Element Type
- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Notes

- Canonical entry built from observed call sites, not from engine source or decompiled definitions.
