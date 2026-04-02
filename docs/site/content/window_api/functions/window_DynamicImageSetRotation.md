# DynamicImageSetRotation

- Category: Window Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 7 addons

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
| Addons seen in | Ace, Aura, Effigy, GCDsaver, LibWBToggler, Shinies, WoH-Reticle |
| Files seen in | `/workspace_addons/Ace/LibGUI.lua:1294`, `/workspace_addons/Aura/Source/AuraHelpers.lua:33`, `/workspace_addons/Effigy/LibGUI.lua:1291`, `/workspace_addons/GCDsaver/libs/LibGUI.lua:1291`, `/workspace_addons/LibWarBoardToggler/libs/LibGUI.lua:1291`, `/workspace_addons/Shinies/Libraries/LibGUI.lua:1291`, `/workspace_addons/WoH-Reticle/libs/LibGUI.lua:1291` |
| Namespaces detected | DynamicImageSetRotation |
| Source kinds | lua_calls |
| Example locations | Ace: LIBGUI_Image:Rotate, Aura: AuraHelpers.SetDynamicImageTexture, Effigy: LIBGUI_Image:Rotate, GCDsaver: LIBGUI_Image:Rotate, LibWBToggler: LIBGUI_Image:Rotate, Shinies: LIBGUI_Image:Rotate |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 7 |
| Global usage count | 7 |
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
DynamicImageSetRotation(arg1, arg2)
```

## Description

Observed mutating runtime image resources on existing controls.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a function or method reference. | Observed values: self.name, window |
| arg2 | Observed as a runtime window or control identifier. | Observed values: angle, rotation |

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
- Shinies
- WoH-Reticle

## Examples

- Ace: LIBGUI_Image:Rotate -> DynamicImageSetRotation(self.name, angle)
- Aura: AuraHelpers.SetDynamicImageTexture -> DynamicImageSetRotation(window, rotation)
- Effigy: LIBGUI_Image:Rotate -> DynamicImageSetRotation(self.name, angle)
- GCDsaver: LIBGUI_Image:Rotate -> DynamicImageSetRotation(self.name, angle)
- LibWBToggler: LIBGUI_Image:Rotate -> DynamicImageSetRotation(self.name, angle)
- Shinies: LIBGUI_Image:Rotate -> DynamicImageSetRotation(self.name, angle)

## Related APIs

- none

## Used With

- [DynamicImageSetTexture](window_DynamicImageSetTexture.md) (HIGH 100/100) - Window Function
- [DynamicImageSetTextureOrientation](window_DynamicImageSetTextureOrientation.md) (HIGH 100/100) - Window Function
- [DynamicImageSetTextureSlice](window_DynamicImageSetTextureSlice.md) (HIGH 100/100) - Window Function
- [WindowSetAlpha](window_WindowSetAlpha.md) (HIGH 100/100) - Window Function
- [DoesWindowExist](../../globals/functions/global_DoesWindowExist.md) (HIGH 83/100) - Global Function

## Triggered By

- none

## Affects

- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Notes

- none
