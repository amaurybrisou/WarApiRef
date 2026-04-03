# DynamicImageSetTexture

- Category: Window Function
- Confidence level: HIGH
- Confidence score: 98/100
- Seen in: 3 addons

## Confidence Assessment

- Level: HIGH

- Score: 98/100

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, matches a known engine namespace, called globally with no local definition.

## Evidence Signals

- +18 Seen in 2 to 3 addons: Cross-addon spread is present but limited.
- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- -15 Conflicting signatures across usages: Observed arity or argument shape conflicts across usages.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | InfoScroller, Moth, PartyCast |
| Files seen in | `/workspace/data/raw/InfoScroller/InfoScroller.lua:95`, `/workspace/data/raw/InfoScroller/libs/LibGUI.lua:1267`, `/workspace/data/raw/Moth/Moth.lua:227`, `/workspace/data/raw/Moth/Moth.lua:418`, `/workspace/data/raw/PartyCast/PartyCast.lua:399`, `/workspace/data/raw/PartyCast/libs/LibGUI.lua:1267` |
| Namespaces detected | DynamicImageSetTexture |
| Source kinds | lua_calls |
| Example locations | InfoScroller: InfoScroller.CreateCard, InfoScroller: LIBGUI_Image:Texture, Moth: Moth.SetCellTextIcon, Moth: Moth.UpdateNPCIcon, PartyCast: LIBGUI_Image:Texture, PartyCast: PartyCast.FetchedText |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 8 |
| Global usage count | 8 |
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
DynamicImageSetTexture(windowName, texture, textureX, textureY)
```

## Description

Observed mutating runtime image resources on existing controls.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| windowName | Observed as a target image control name. | Observed values: "PartyCastWindow"..PlayerNumber.."ButtonIcon", "PartyCastWindow"..PlayerNumber.."TargetWindowIcon", WindowName.."Image" |
| texture | Observed as a texture resource name. | Observed values: "", "map_markers01", Icon_texture |
| textureX | Observed as a numeric texture coordinate or atlas offset. | Observed values: 0, 32, 64 |
| textureY | Observed as a numeric texture coordinate or atlas offset. | Observed values: 0, 32, 64 |

## Returns

- Not confidently inferable from addon-api docs alone.

## Side Effects

- Mutates engine-visible UI state or presentation.

## Seen In

- InfoScroller
- Moth
- PartyCast

## Examples

- InfoScroller: InfoScroller.CreateCard -> DynamicImageSetTexture(WindowName.."Image", T_Name, width*scale, height*scale)
- InfoScroller: LIBGUI_Image:Texture -> DynamicImageSetTexture(self.name, texture, x, y)
- Moth: Moth.SetCellTextIcon -> DynamicImageSetTexture(cellName.."CareerIcon", careerIconId, 0, 0)
- Moth: Moth.UpdateNPCIcon -> DynamicImageSetTexture(iconWindow, "map_markers01", 32, 32)
- PartyCast: LIBGUI_Image:Texture -> DynamicImageSetTexture(self.name, texture, x, y)
- PartyCast: PartyCast.FetchedText -> DynamicImageSetTexture("PartyCastWindow"..PlayerNumber.."ButtonIcon", texture, 64, 64)

## Related APIs

- none

## Used With

- [DynamicImageSetTextureDimensions](window_DynamicImageSetTextureDimensions.md) (HIGH 100/100) - Window Function
- [DynamicImageSetTextureOrientation](window_DynamicImageSetTextureOrientation.md) (HIGH 100/100) - Window Function
- [DynamicImageSetTextureScale](window_DynamicImageSetTextureScale.md) (HIGH 100/100) - Window Function
- [InterfaceCore.GetScale](../../globals/functions/global_InterfaceCore.GetScale.md) (HIGH 100/100) - Global Function
- [LabelSetFont](window_LabelSetFont.md) (HIGH 100/100) - Window Function
- [LabelSetText](window_LabelSetText.md) (HIGH 100/100) - Window Function
- [LabelSetTextColor](window_LabelSetTextColor.md) (HIGH 100/100) - Window Function
- [Icons.GetCareerIconIDFromCareerLine](../../globals/functions/global_Icons.GetCareerIconIDFromCareerLine.md) (HIGH 88/100) - Global Function
- [DoesWindowExist](../../globals/functions/global_DoesWindowExist.md) (HIGH 83/100) - Global Function
- [CreateWindowFromTemplate](../../globals/functions/global_CreateWindowFromTemplate.md) (HIGH 75/100) - Global Function
- [towstring](../../globals/functions/global_towstring.md) (HIGH 75/100) - Global Function
- [GetIconData](../../globals/functions/global_GetIconData.md) (HIGH 71/100) - Global Function

## Triggered By

- none

## Affects

- [InterfaceCore.GetScale](../../globals/functions/global_InterfaceCore.GetScale.md) (HIGH 100/100) - Global Function
- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Notes

- none
