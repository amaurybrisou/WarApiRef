# DynamicImageSetTextureScale

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
| Addons seen in | InfoScroller, Moth, PartyCast |
| Files seen in | `/workspace/data/raw/InfoScroller/InfoScroller.lua:95`, `/workspace/data/raw/InfoScroller/libs/LibGUI.lua:1323`, `/workspace/data/raw/Moth/Moth.lua:418`, `/workspace/data/raw/PartyCast/libs/LibGUI.lua:1323` |
| Namespaces detected | DynamicImageSetTextureScale |
| Source kinds | lua_calls |
| Example locations | InfoScroller: InfoScroller.CreateCard, InfoScroller: LIBGUI_Image:TexScale, Moth: Moth.UpdateNPCIcon, PartyCast: LIBGUI_Image:TexScale |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 4 |
| Global usage count | 4 |
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
DynamicImageSetTextureScale(arg1, arg2)
```

## Description

Observed mutating runtime image resources on existing controls.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a function or method reference. | Observed values: WindowName.."Image", iconWindow, self.name |
| arg2 | Observed as a function or method reference. | Observed values: TicketTable.Image.scale, filterData.scale, scale |

## Returns

- Not confidently inferable from addon-api docs alone.

## Side Effects

- Mutates engine-visible UI state or presentation.

## Seen In

- InfoScroller
- Moth
- PartyCast

## Examples

- InfoScroller: InfoScroller.CreateCard -> DynamicImageSetTextureScale(WindowName.."Image", TicketTable.Image.scale)
- InfoScroller: LIBGUI_Image:TexScale -> DynamicImageSetTextureScale(self.name, scale)
- Moth: Moth.UpdateNPCIcon -> DynamicImageSetTextureScale(iconWindow, filterData.scale)
- PartyCast: LIBGUI_Image:TexScale -> DynamicImageSetTextureScale(self.name, scale)

## Related APIs

- none

## Used With

- [DynamicImageSetTextureOrientation](window_DynamicImageSetTextureOrientation.md) (HIGH 100/100) - Window Function
- [InterfaceCore.GetScale](../../globals/functions/global_InterfaceCore.GetScale.md) (HIGH 100/100) - Global Function
- [DynamicImageSetTexture](window_DynamicImageSetTexture.md) (HIGH 98/100) - Window Function
- [DoesWindowExist](../../globals/functions/global_DoesWindowExist.md) (HIGH 83/100) - Global Function
- [CreateWindowFromTemplate](../../globals/functions/global_CreateWindowFromTemplate.md) (HIGH 75/100) - Global Function

## Triggered By

- none

## Affects

- [InterfaceCore.GetScale](../../globals/functions/global_InterfaceCore.GetScale.md) (HIGH 100/100) - Global Function
- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Notes

- none
