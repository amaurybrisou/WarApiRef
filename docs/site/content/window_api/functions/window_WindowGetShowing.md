# WindowGetShowing

- Category: Window Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 6 addons

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
| Addons seen in | InfoScroller, PartyCast, Soloq, TidyChat, TidyRoll, ZCurse_Profiler |
| Files seen in | `/workspace/data/raw/CurseProfiler/CurseProfilerCompiled.lua:1065`, `/workspace/data/raw/InfoScroller/libs/LibGUI.lua:84`, `/workspace/data/raw/PartyCast/PartyCast.lua:399`, `/workspace/data/raw/PartyCast/PartyCast.lua:655`, `/workspace/data/raw/PartyCast/libs/LibGUI.lua:84`, `/workspace/data/raw/Soloq/Soloq.lua:166`, `/workspace/data/raw/Soloq/ui/Overview.lua:38`, `/workspace/data/raw/TidyChat/TidyChat.lua:1089` |
| Namespaces detected | WindowGetShowing |
| Source kinds | lua_calls |
| Example locations | InfoScroller: LIBGUI_ELEMENT:Showing, PartyCast: LIBGUI_ELEMENT:Showing, PartyCast: PartyCast.FetchedText, PartyCast: PartyCast.Update, Soloq: Soloq.onEnterRankedScenario, Soloq: Soloq.toggleOverviewWindow |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 18 |
| Global usage count | 18 |
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
WindowGetShowing(windowName)
```

## Description

Observed querying runtime window state or metadata.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| windowName | Observed as a target window name. | Observed values: "EA_Window_WorldMap", "PartyCastWindow"..PlayerNumber, "PartyCastWindow_Dynamic"..i |

## Returns

- Not confidently inferable from addon-api docs alone.

## Side Effects

- No side effect is confidently inferable from addon-api docs alone.

## Seen In

- InfoScroller
- PartyCast
- Soloq
- TidyChat
- TidyRoll
- ZCurse_Profiler

## Examples

- InfoScroller: LIBGUI_ELEMENT:Showing -> WindowGetShowing(self.name)
- PartyCast: LIBGUI_ELEMENT:Showing -> WindowGetShowing(self.name)
- PartyCast: PartyCast.FetchedText -> WindowGetShowing("PartyCastWindow"..PlayerNumber)
- PartyCast: PartyCast.Update -> WindowGetShowing("PartyCastWindow_Dynamic"..i)
- Soloq: Soloq.onEnterRankedScenario -> WindowGetShowing("SoloqOverviewWindow")
- Soloq: Soloq.toggleOverviewWindow -> WindowGetShowing(overviewWindowName)

## Related APIs

- none

## Used With

- [LayoutEditor.RegisterWindow](window_LayoutEditor.RegisterWindow.md) (HIGH 100/100) - Window Function
- [StatusBarGetCurrentValue](window_StatusBarGetCurrentValue.md) (HIGH 100/100) - Window Function
- [WindowAddAnchor](window_WindowAddAnchor.md) (HIGH 100/100) - Window Function
- [WindowClearAnchors](window_WindowClearAnchors.md) (HIGH 100/100) - Window Function
- [WindowGetAlpha](window_WindowGetAlpha.md) (HIGH 100/100) - Window Function
- [DoesWindowExist](../../globals/functions/global_DoesWindowExist.md) (HIGH 83/100) - Global Function
- [StatusBarGetMaximumValue](window_StatusBarGetMaximumValue.md) (HIGH 80/100) - Window Function

## Triggered By

- none

## Affects

- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Notes

- none
