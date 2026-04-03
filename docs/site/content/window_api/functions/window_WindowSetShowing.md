# WindowSetShowing

- Category: Window Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 8 addons

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
| Addons seen in | InfoScroller, Moth, PartyCast, Soloq, TidyChat, TidyRoll, ZCurse_Profiler, minesweep |
| Files seen in | `/workspace/data/raw/CurseProfiler/CurseProfilerCompiled.lua:1065`, `/workspace/data/raw/InfoScroller/InfoScroller.lua:95`, `/workspace/data/raw/InfoScroller/libs/LibGUI.lua:74`, `/workspace/data/raw/InfoScroller/libs/LibGUI.lua:79`, `/workspace/data/raw/Moth/Moth.lua:167`, `/workspace/data/raw/Moth/Moth.lua:267`, `/workspace/data/raw/Moth/Moth.lua:418`, `/workspace/data/raw/Moth/Moth.lua:460` |
| Namespaces detected | WindowSetShowing |
| Source kinds | lua_calls |
| Example locations | InfoScroller: InfoScroller.CreateCard, InfoScroller: LIBGUI_ELEMENT:Hide, InfoScroller: LIBGUI_ELEMENT:Show, Moth: Moth.Clear, Moth: Moth.HealthBar, Moth: Moth.HideBorders |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 67 |
| Global usage count | 67 |
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
WindowSetShowing(windowName, arg2)
```

## Description

Observed mutating runtime window state or presentation.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| windowName | Observed as a target window name. | Observed values: "DefaultTooltip", "EA_Window_LootRoll", "EA_Window_WorldMap" |
| arg2 | Observed as a boolean toggle. | Observed values: (NameOfTargetBG~=L "")and PartyCast.Settings.ShowTarget, InfoScroller.Settings.Alignment=="left" or InfoScroller.Settings.Alignment=="center", InfoScroller.Settings.Alignment=="right" or InfoScroller.Settings.Alignment=="center" |

## Returns

- Not confidently inferable from addon-api docs alone.

## Side Effects

- Mutates engine-visible UI state or presentation.

## Seen In

- InfoScroller
- Moth
- PartyCast
- Soloq
- TidyChat
- TidyRoll
- ZCurse_Profiler
- minesweep

## Examples

- InfoScroller: InfoScroller.CreateCard -> WindowSetShowing(WindowName.."BackGroundStart", InfoScroller.Settings.Alignment=="right" or InfoScroller.Settings.Alignment=="center")
- InfoScroller: InfoScroller.CreateCard -> WindowSetShowing(WindowName.."BackGroundEnd", InfoScroller.Settings.Alignment=="left" or InfoScroller.Settings.Alignment=="center")
- InfoScroller: InfoScroller.CreateCard -> WindowSetShowing(WindowName.."BackGround", InfoScroller.Settings.ShowBG)
- InfoScroller: LIBGUI_ELEMENT:Hide -> WindowSetShowing(self.name, false)
- InfoScroller: LIBGUI_ELEMENT:Show -> WindowSetShowing(self.name, true)
- Moth: Moth.Clear -> WindowSetShowing("Moth", false)

## Related APIs

- none

## Used With

- [DynamicImageSetTextureSlice](window_DynamicImageSetTextureSlice.md) (HIGH 100/100) - Window Function
- [LibSlash.RegisterWSlashCmd](../../globals/functions/global_LibSlash.RegisterWSlashCmd.md) (HIGH 100/100) - Global Function
- [SystemData.Events.LOADING_END](../../events/game_events/game_event_SystemData.Events.LOADING_END.md) (HIGH 100/100) - Game Event
- [SystemData.Events.RELOAD_INTERFACE](../../events/game_events/game_event_SystemData.Events.RELOAD_INTERFACE.md) (HIGH 100/100) - Game Event
- [WindowAddAnchor](window_WindowAddAnchor.md) (HIGH 100/100) - Window Function
- [WindowClearAnchors](window_WindowClearAnchors.md) (HIGH 100/100) - Window Function
- [WindowGetDimensions](window_WindowGetDimensions.md) (HIGH 100/100) - Window Function
- [WindowSetDimensions](window_WindowSetDimensions.md) (HIGH 100/100) - Window Function
- [WindowSetTintColor](window_WindowSetTintColor.md) (HIGH 100/100) - Window Function
- [WindowUnregisterCoreEventHandler](window_WindowUnregisterCoreEventHandler.md) (HIGH 100/100) - Window Function

## Triggered By

- none

## Affects

- [InterfaceCore.GetScale](../../globals/functions/global_InterfaceCore.GetScale.md) (HIGH 100/100) - Global Function
- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Notes

- none
