# DynamicImageSetTextureOrientation

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
| Addons seen in | Ace, Aura, Effigy, GCDsaver, LibWBToggler, RoR_SoR, Shinies, WoH-Reticle |
| Files seen in | `/workspace_addons/Ace/LibGUI.lua:1316`, `/workspace_addons/Aura/Source/AuraHelpers.lua:33`, `/workspace_addons/Effigy/LibGUI.lua:1313`, `/workspace_addons/GCDsaver/libs/LibGUI.lua:1313`, `/workspace_addons/LibWarBoardToggler/libs/LibGUI.lua:1313`, `/workspace_addons/RoR_SoR/RoR_SoR.lua:592`, `/workspace_addons/Shinies/Libraries/LibGUI.lua:1313`, `/workspace_addons/WoH-Reticle/libs/LibGUI.lua:1313` |
| Namespaces detected | DynamicImageSetTextureOrientation |
| Source kinds | lua_calls |
| Example locations | Ace: LIBGUI_Image:TexFlip, Aura: AuraHelpers.SetDynamicImageTexture, Effigy: LIBGUI_Image:TexFlip, GCDsaver: LIBGUI_Image:TexFlip, LibWBToggler: LIBGUI_Image:TexFlip, RoR_SoR: RoR_SoR.SET_CITY |
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
DynamicImageSetTextureOrientation(arg1, arg2)
```

## Description

Observed mutating runtime image resources on existing controls.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a function or method reference. | Observed values: "SoR_"..Window_Name.."Flames", self.name, window |
| arg2 | Observed as a runtime window or control identifier. | Observed values: Window_Name=="162", flipped, mirror |

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
- RoR_SoR
- Shinies
- WoH-Reticle

## Examples

- Ace: LIBGUI_Image:TexFlip -> DynamicImageSetTextureOrientation(self.name, flipped)
- Aura: AuraHelpers.SetDynamicImageTexture -> DynamicImageSetTextureOrientation(window, mirror)
- Effigy: LIBGUI_Image:TexFlip -> DynamicImageSetTextureOrientation(self.name, flipped)
- GCDsaver: LIBGUI_Image:TexFlip -> DynamicImageSetTextureOrientation(self.name, flipped)
- LibWBToggler: LIBGUI_Image:TexFlip -> DynamicImageSetTextureOrientation(self.name, flipped)
- RoR_SoR: RoR_SoR.SET_CITY -> DynamicImageSetTextureOrientation("SoR_"..Window_Name.."Flames", Window_Name=="162")

## Related APIs

- none

## Used With

- [DynamicImageSetRotation](window_DynamicImageSetRotation.md) (HIGH 100/100) - Window Function
- [DynamicImageSetTextureSlice](window_DynamicImageSetTextureSlice.md) (HIGH 100/100) - Window Function
- [WindowSetAlpha](window_WindowSetAlpha.md) (HIGH 100/100) - Window Function

## Triggered By

- none

## Affects

- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Notes

- none
