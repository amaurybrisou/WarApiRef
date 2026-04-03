# DynamicImageSetTextureSlice

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
| Addons seen in | InfoScroller, Moth, PartyCast, minesweep |
| Files seen in | `/workspace/data/raw/InfoScroller/InfoScroller.lua:95`, `/workspace/data/raw/InfoScroller/libs/LibGUI.lua:1281`, `/workspace/data/raw/Moth/Moth.lua:418`, `/workspace/data/raw/PartyCast/libs/LibGUI.lua:1281`, `/workspace/data/raw/minesweep/minesweep.lua:128`, `/workspace/data/raw/minesweep/minesweep.lua:82` |
| Namespaces detected | DynamicImageSetTextureSlice |
| Source kinds | lua_calls |
| Example locations | InfoScroller: InfoScroller.CreateCard, InfoScroller: LIBGUI_Image:TexSlice, Moth: Moth.UpdateNPCIcon, PartyCast: LIBGUI_Image:TexSlice, minesweep: minesweep.LButtonUp, minesweep: minesweep.RButtonUp |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 6 |
| Global usage count | 6 |
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
DynamicImageSetTextureSlice(arg1, arg2)
```

## Description

Observed mutating runtime image resources on existing controls.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a function or method reference. | Observed values: WindowName.."Icon", WindowName.."Image", iconWindow |
| arg2 | Observed as a text or wstring payload. | Observed values: "BombDestruction", "GuildStandard", TicketTable.Image.slice |

## Returns

- Not confidently inferable from addon-api docs alone.

## Side Effects

- Mutates engine-visible UI state or presentation.

## Seen In

- InfoScroller
- Moth
- PartyCast
- minesweep

## Examples

- InfoScroller: InfoScroller.CreateCard -> DynamicImageSetTextureSlice(WindowName.."Image", TicketTable.Image.slice)
- InfoScroller: LIBGUI_Image:TexSlice -> DynamicImageSetTextureSlice(self.name, slice)
- Moth: Moth.UpdateNPCIcon -> DynamicImageSetTextureSlice(iconWindow, filterData.slice)
- PartyCast: LIBGUI_Image:TexSlice -> DynamicImageSetTextureSlice(self.name, slice)
- minesweep: minesweep.LButtonUp -> DynamicImageSetTextureSlice(WindowName.."Icon", "BombDestruction")
- minesweep: minesweep.RButtonUp -> DynamicImageSetTextureSlice(WindowName.."Icon", "GuildStandard")

## Related APIs

- none

## Used With

- [WindowSetShowing](window_WindowSetShowing.md) (HIGH 100/100) - Window Function

## Triggered By

- [OnLButtonUp](../../xml/handlers/handler_OnLButtonUp.md) (HIGH 100/100) - XML Event
- [OnLButtonUp](../../events/window_events/window_event_OnLButtonUp.md) (HIGH 100/100) - Window Event
- [OnRButtonUp](../../xml/handlers/handler_OnRButtonUp.md) (HIGH 100/100) - XML Event
- [OnRButtonUp](../../events/window_events/window_event_OnRButtonUp.md) (HIGH 100/100) - Window Event

## Affects

- [Button](../../xml/element_types/element_Button.md) (HIGH 100/100) - XML Element Type
- [InterfaceCore.GetScale](../../globals/functions/global_InterfaceCore.GetScale.md) (HIGH 100/100) - Global Function
- [SystemData.ActiveWindow.name](../../systemdata/fields/systemdata_SystemData.ActiveWindow.name.md) (HIGH 100/100) - SystemData Field
- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Notes

- none
