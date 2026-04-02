# DynamicImageSetTextureDimensions

- Category: Window Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 12 addons

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
| Addons seen in | Ace, BuffHead, Effigy, Enemy, GCDsaver, LibWBToggler, Miracle Grow Remix, PotionBar |
| Files seen in | `/workspace_addons/Ace/LibGUI.lua:1304`, `/workspace_addons/BuffHead/Setup/SetupLayoutProperties.lua:117`, `/workspace_addons/Effigy/LibGUI.lua:1301`, `/workspace_addons/Effigy/Textures/TextureManager.lua:125`, `/workspace_addons/Effigy/Textures/TextureManager.lua:64`, `/workspace_addons/Enemy/Code/UnitFrames/Parts/CareerIcon.lua:13`, `/workspace_addons/Enemy/Code/UnitFrames/UnitFramePart.lua:393`, `/workspace_addons/Enemy/Code/UnitFrames/UnitFramePart.lua:407` |
| Namespaces detected | DynamicImageSetTextureDimensions |
| Source kinds | lua_calls |
| Example locations | Ace: LIBGUI_Image:TexDims, BuffHead: BuffHead.local.SetTextureButton, BuffHead: SetTextureButton, Effigy: Effigy.TextureManager.SetTexture, Effigy: Effigy.TextureManager.SetTextureFill, Effigy: LIBGUI_Image:TexDims |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 30 |
| Global usage count | 30 |
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
DynamicImageSetTextureDimensions(arg1, arg2, arg3)
```

## Description

Observed mutating runtime image resources on existing controls.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a function or method reference. | Observed values: adat.iconname, iconWindowName, mg.sWindowName.."LayoutBarBack" |
| arg2 | Observed as a function or method reference. | Observed values: 256, 32, 64 |
| arg3 | Observed as a function or method reference. | Observed values: 256, 32, 64 |

## Returns

- Not confidently inferable from addon-api docs alone.

## Side Effects

- Mutates engine-visible UI state or presentation.

## Seen In

- Ace
- BuffHead
- Effigy
- Enemy
- GCDsaver
- LibWBToggler
- Miracle Grow Remix
- PotionBar
- Shinies
- TexturedButtons
- WSCT
- WoH-Reticle

## Examples

- Ace: LIBGUI_Image:TexDims -> DynamicImageSetTextureDimensions(self.name, width, height)
- BuffHead: BuffHead.local.SetTextureButton -> DynamicImageSetTextureDimensions(textureButton.."Texture", dimensions.Width, dimensions.Height)
- BuffHead: SetTextureButton -> DynamicImageSetTextureDimensions(textureButton.."Texture", dimensions.Width, dimensions.Height)
- Effigy: Effigy.TextureManager.SetTexture -> DynamicImageSetTextureDimensions(win_name, Addon.texture_list[t].width, Addon.texture_list[t].height)
- Effigy: Effigy.TextureManager.SetTextureFill -> DynamicImageSetTextureDimensions(win_name, perc*Addon.texture_list[t].width, Addon.texture_list[t].height)
- Effigy: Effigy.TextureManager.SetTextureFill -> DynamicImageSetTextureDimensions(win_name, Addon.texture_list[t].width, Addon.texture_list[t].height*perc)

## Related APIs

- none

## Used With

- [DynamicImageSetTexture](window_DynamicImageSetTexture.md) (HIGH 100/100) - Window Function

## Triggered By

- none

## Affects

- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Notes

- none
