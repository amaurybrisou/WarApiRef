# DynamicImageSetTextureScale

- Category: Window Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 10 addons

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
| Addons seen in | Ace, Effigy, GCDsaver, LibWBToggler, MapMonster, Miracle Grow Remix, Moth, Shinies |
| Files seen in | `/workspace_addons/Ace/LibGUI.lua:1326`, `/workspace_addons/Effigy/Elements/EffigyBar.lua:135`, `/workspace_addons/Effigy/LibGUI.lua:1323`, `/workspace_addons/Effigy/Textures/TextureManager.lua:125`, `/workspace_addons/GCDsaver/libs/LibGUI.lua:1323`, `/workspace_addons/LibWarBoardToggler/LibWBToggler.lua:16`, `/workspace_addons/LibWarBoardToggler/libs/LibGUI.lua:1323`, `/workspace_addons/MGRemix/MGRemix.lua:797` |
| Namespaces detected | DynamicImageSetTextureScale |
| Source kinds | lua_calls |
| Example locations | Ace: LIBGUI_Image:TexScale, Effigy: Effigy.TextureManager.SetTexture, Effigy: EffigyBar:setup, Effigy: LIBGUI_Image:TexScale, GCDsaver: LIBGUI_Image:TexScale, LibWBToggler: LIBGUI_Image:TexScale |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 23 |
| Global usage count | 23 |
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
| arg1 | Observed as a function or method reference. | Observed values: editorWindow.."MapPinIcon", entryData.windowName.."MapIcon", iconWindow |
| arg2 | Observed as a function or method reference. | Observed values: .125, 1, Addon.texture_list[t].scale |

## Returns

- Not confidently inferable from addon-api docs alone.

## Side Effects

- Mutates engine-visible UI state or presentation.

## Seen In

- Ace
- Effigy
- GCDsaver
- LibWBToggler
- MapMonster
- Miracle Grow Remix
- Moth
- Shinies
- TexturedButtons
- WoH-Reticle

## Examples

- Ace: LIBGUI_Image:TexScale -> DynamicImageSetTextureScale(self.name, scale)
- Effigy: Effigy.TextureManager.SetTexture -> DynamicImageSetTextureScale(win_name, Addon.texture_list[t].scale)
- Effigy: EffigyBar:setup -> DynamicImageSetTextureScale(runtimeCache.Names.fg.."Fill", self.fg.texture.scale)
- Effigy: LIBGUI_Image:TexScale -> DynamicImageSetTextureScale(self.name, scale)
- GCDsaver: LIBGUI_Image:TexScale -> DynamicImageSetTextureScale(self.name, scale)
- LibWBToggler: LIBGUI_Image:TexScale -> DynamicImageSetTextureScale(self.name, scale)

## Related APIs

- none

## Used With

- [DynamicImageSetTexture](window_DynamicImageSetTexture.md) (HIGH 100/100) - Window Function
- [DynamicImageSetTextureDimensions](window_DynamicImageSetTextureDimensions.md) (HIGH 100/100) - Window Function
- [DynamicImageSetTextureSlice](window_DynamicImageSetTextureSlice.md) (HIGH 100/100) - Window Function
- [LayoutEditor.UnregisterWindow](window_LayoutEditor.UnregisterWindow.md) (HIGH 100/100) - Window Function
- [WindowAddAnchor](window_WindowAddAnchor.md) (HIGH 100/100) - Window Function
- [WindowClearAnchors](window_WindowClearAnchors.md) (HIGH 100/100) - Window Function
- [WindowSetAlpha](window_WindowSetAlpha.md) (HIGH 100/100) - Window Function
- [WindowSetDimensions](window_WindowSetDimensions.md) (HIGH 100/100) - Window Function
- [WindowSetHandleInput](window_WindowSetHandleInput.md) (HIGH 100/100) - Window Function
- [WindowSetLayer](window_WindowSetLayer.md) (HIGH 100/100) - Window Function
- [WindowSetParent](window_WindowSetParent.md) (HIGH 100/100) - Window Function
- [WindowSetScale](window_WindowSetScale.md) (HIGH 100/100) - Window Function
- [WindowSetShowing](window_WindowSetShowing.md) (HIGH 100/100) - Window Function
- [WindowSetTintColor](window_WindowSetTintColor.md) (HIGH 100/100) - Window Function
- [DoesWindowExist](../../globals/functions/global_DoesWindowExist.md) (HIGH 83/100) - Global Function
- [CreateWindowFromTemplate](../../globals/functions/global_CreateWindowFromTemplate.md) (HIGH 75/100) - Global Function
- [towstring](../../globals/functions/global_towstring.md) (HIGH 75/100) - Global Function

## Triggered By

- none

## Affects

- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Notes

- none
