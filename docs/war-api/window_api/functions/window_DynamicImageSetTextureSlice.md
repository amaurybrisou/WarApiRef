# DynamicImageSetTextureSlice

- Category: Window Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 16 addons

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
| Addons seen in | Ace, Aura, Effigy, GCDsaver, LibWBToggler, MapMonster, MapPin, Miracle Grow Remix |
| Files seen in | `/workspace/Ace/LibGUI.lua:1284`, `/workspace/Aura/Source/AuraHelpers.lua:33`, `/workspace/Effigy/LibGUI.lua:1281`, `/workspace/Effigy/Textures/TextureManager.lua:125`, `/workspace/GCDsaver/libs/LibGUI.lua:1281`, `/workspace/LibWarBoardToggler/LibWBToggler.lua:16`, `/workspace/LibWarBoardToggler/libs/LibGUI.lua:1281`, `/workspace/MGRemix/MGRemix.lua:797` |
| Namespaces detected | DynamicImageSetTextureSlice |
| Source kinds | lua_calls |
| Example locations | Ace: LIBGUI_Image:TexSlice, Aura: AuraHelpers.SetDynamicImageTexture, Effigy: Effigy.TextureManager.SetTexture, Effigy: LIBGUI_Image:TexSlice, GCDsaver: LIBGUI_Image:TexSlice, LibWBToggler: LIBGUI_Image:TexSlice |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 58 |
| Global usage count | 58 |
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
| arg1 | Observed as a function or method reference. | Observed values: "MapPinTooltipsIcon", "SoR_"..Window_Name.."BG", "SoR_"..Window_Name.."CITY_RANK" |
| arg2 | Observed as a text or wstring payload. | Observed values: "Black-Slot", "Dirt", "Dirt-Slot" |

## Returns

- Not confidently inferable from addon-api docs alone.

## Side Effects

- Mutates engine-visible UI state or presentation.

## Seen In

- Ace
- Aura
- Effigy
- GCDsaver
- LibWBToggler
- MapMonster
- MapPin
- Miracle Grow Remix
- MiracleGrow
- MiracleGrowLight
- Moth
- RoR_SoR
- Shinies
- TexturedButtons
- WarTriage
- WoH-Reticle

## Examples

- Ace: LIBGUI_Image:TexSlice -> DynamicImageSetTextureSlice(self.name, slice)
- Aura: AuraHelpers.SetDynamicImageTexture -> DynamicImageSetTextureSlice(window, slice)
- Effigy: Effigy.TextureManager.SetTexture -> DynamicImageSetTextureSlice(win_name, Addon.texture_list[t].slice)
- Effigy: LIBGUI_Image:TexSlice -> DynamicImageSetTextureSlice(self.name, slice)
- GCDsaver: LIBGUI_Image:TexSlice -> DynamicImageSetTextureSlice(self.name, slice)
- LibWBToggler: LIBGUI_Image:TexSlice -> DynamicImageSetTextureSlice(self.name, slice)

## Related APIs

- none

## Used With

- [DynamicImageSetRotation](window_DynamicImageSetRotation.md) (HIGH 100/100) - Window Function
- [DynamicImageSetTexture](window_DynamicImageSetTexture.md) (HIGH 100/100) - Window Function
- [DynamicImageSetTextureDimensions](window_DynamicImageSetTextureDimensions.md) (HIGH 100/100) - Window Function
- [DynamicImageSetTextureOrientation](window_DynamicImageSetTextureOrientation.md) (HIGH 100/100) - Window Function
- [DynamicImageSetTextureScale](window_DynamicImageSetTextureScale.md) (HIGH 100/100) - Window Function
- [WindowSetAlpha](window_WindowSetAlpha.md) (HIGH 100/100) - Window Function
- [WindowSetTintColor](window_WindowSetTintColor.md) (HIGH 100/100) - Window Function
- [DoesWindowExist](../../globals/functions/global_DoesWindowExist.md) (HIGH 83/100) - Global Function

## Triggered By

- none

## Affects

- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Notes

- none
