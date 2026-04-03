# WindowSetTintColor

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
| Files seen in | `/workspace/data/raw/InfoScroller/libs/LibConfig.lua:583`, `/workspace/data/raw/InfoScroller/libs/LibGUI.lua:249`, `/workspace/data/raw/Moth/Moth.lua:267`, `/workspace/data/raw/Moth/Moth.lua:588`, `/workspace/data/raw/PartyCast/PartyCast.lua:399`, `/workspace/data/raw/PartyCast/PartyCast.lua:655`, `/workspace/data/raw/PartyCast/libs/LibConfig.lua:583`, `/workspace/data/raw/PartyCast/libs/LibGUI.lua:249` |
| Namespaces detected | WindowSetTintColor |
| Source kinds | lua_calls |
| Example locations | InfoScroller: LIBGUI_ELEMENT:Tint, InfoScroller: LibConfig.OnUpdate, Moth: Moth.UpdateHealthBar, Moth: Moth.UpdateLevel, PartyCast: LIBGUI_ELEMENT:Tint, PartyCast: LibConfig.OnUpdate |
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
WindowSetTintColor(windowName, arg2, arg3, arg4)
```

## Description

Observed mutating runtime window state or presentation.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| windowName | Observed as a target window name. | Observed values: "MothHealthBar", "PartyCastWindow"..PlayerNumber.."Button", "PartyCastWindow"..PlayerNumber.."ButtonIcon" |
| arg2 | Observed as a function or method reference. | Observed values: 0, 175, 200 |
| arg3 | Observed as a function or method reference. | Observed values: 0, 175, 200 |
| arg4 | Observed as a function or method reference. | Observed values: 0, 200, 255 |

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

- InfoScroller: LIBGUI_ELEMENT:Tint -> WindowSetTintColor(self.name, r, g, b)
- InfoScroller: LibConfig.OnUpdate -> WindowSetTintColor(LibConfig.colorizer.object.name, LibConfig.colorizer.r:GetValue(), LibConfig.colorizer.g:GetValue(), LibConfig.colorizer.b:GetValue())
- Moth: Moth.UpdateHealthBar -> WindowSetTintColor("MothHealthBar", Moth.Helpers.DefaultColorConvert(DefaultColor.GREEN))
- Moth: Moth.UpdateHealthBar -> WindowSetTintColor("MothHealthBar", Moth.Helpers.DefaultColorConvert(DefaultColor.YELLOW))
- Moth: Moth.UpdateHealthBar -> WindowSetTintColor("MothHealthBar", Moth.Helpers.DefaultColorConvert(DefaultColor.ORANGE))
- Moth: Moth.UpdateHealthBar -> WindowSetTintColor("MothHealthBar", Moth.Helpers.DefaultColorConvert(DefaultColor.RED))

## Related APIs

- none

## Used With

- [WindowGetDimensions](window_WindowGetDimensions.md) (HIGH 100/100) - Window Function
- [WindowSetDimensions](window_WindowSetDimensions.md) (HIGH 100/100) - Window Function
- [WindowSetShowing](window_WindowSetShowing.md) (HIGH 100/100) - Window Function

## Triggered By

- [SystemData.Events.UPDATE_PROCESSED](../../events/game_events/game_event_SystemData.Events.UPDATE_PROCESSED.md) (HIGH 100/100) - Game Event

## Affects

- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Notes

- none
