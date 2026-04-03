# DynamicImageSetTextureSlice

- Category: Window Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 9 addons

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
| Addons seen in | Ace, Aura, LibWBToggler, MiracleGrowLight, PartyCast, RoR_SoR, Shinies, TexturedButtons |
| Files seen in | `/workspace/data/raw/Ace/LibGUI.lua:1284`, `/workspace/data/raw/Aura/Source/AuraHelpers.lua:33`, `/workspace/data/raw/LibWarBoardToggler/LibWBToggler.lua:16`, `/workspace/data/raw/LibWarBoardToggler/libs/LibGUI.lua:1281`, `/workspace/data/raw/MiracleGrowLight/MiracleGrowLight.lua:168`, `/workspace/data/raw/MiracleGrowLight/MiracleGrowLight.lua:301`, `/workspace/data/raw/PartyCast/libs/LibGUI.lua:1281`, `/workspace/data/raw/RoR_SoR/RoR_SoR.lua:1055` |
| Namespaces detected | DynamicImageSetTextureSlice |
| Source kinds | lua_calls |
| Example locations | Ace: LIBGUI_Image:TexSlice, Aura: AuraHelpers.SetDynamicImageTexture, LibWBToggler: LIBGUI_Image:TexSlice, LibWBToggler: LibWBToggler.CreateToggler, MiracleGrowLight: MiracleGrowLight.Initialize, MiracleGrowLight: MiracleGrowLight.OnUpdate |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 24 |
| Global usage count | 24 |
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
| arg1 | Observed as a function or method reference. | Observed values: "SoR_"..Window_Name.."BG", "SoR_"..Window_Name.."CITY_RANK", "SoR_"..Window_Name.."KEEP1KEEPRANK" |
| arg2 | Observed as a text or wstring payload. | Observed values: "Black-Slot", "Dirt", "GreenCross" |

## Returns

- Not confidently inferable from addon-api docs alone.

## Side Effects

- Mutates engine-visible UI state or presentation.

## Seen In

- Ace
- Aura
- LibWBToggler
- MiracleGrowLight
- PartyCast
- RoR_SoR
- Shinies
- TexturedButtons
- WoH-Reticle

## Examples

- Ace: LIBGUI_Image:TexSlice -> DynamicImageSetTextureSlice(self.name, slice)
- Aura: AuraHelpers.SetDynamicImageTexture -> DynamicImageSetTextureSlice(window, slice)
- LibWBToggler: LIBGUI_Image:TexSlice -> DynamicImageSetTextureSlice(self.name, slice)
- LibWBToggler: LibWBToggler.CreateToggler -> DynamicImageSetTextureSlice(modName.."Icon", texSlice)
- MiracleGrowLight: MiracleGrowLight.Initialize -> DynamicImageSetTextureSlice(windowName.."Plant"..i.."ButtonFrame", "IconFrame-1")
- MiracleGrowLight: MiracleGrowLight.Initialize -> DynamicImageSetTextureSlice(windowName.."Plant"..i.."HarvestFrame", "IconFrame-1")

## Related APIs

- none

## Used With

- [DynamicImageSetRotation](window_DynamicImageSetRotation.md) (HIGH 100/100) - Window Function
- [DynamicImageSetTextureOrientation](window_DynamicImageSetTextureOrientation.md) (HIGH 100/100) - Window Function
- [WindowSetAlpha](window_WindowSetAlpha.md) (HIGH 100/100) - Window Function

## Triggered By

- none

## Affects

- [GameData.PlayerActions.PERFORM_CRAFTING](../../gamedata/fields/gamedata_GameData.PlayerActions.PERFORM_CRAFTING.md) (HIGH 100/100) - GameData Field
- [GameData.TradeSkills.CULTIVATION](../../gamedata/fields/gamedata_GameData.TradeSkills.CULTIVATION.md) (HIGH 100/100) - GameData Field
- [SystemData.Events.LOADING_END](../../systemdata/fields/systemdata_SystemData.Events.LOADING_END.md) (HIGH 100/100) - SystemData Field
- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Notes

- none
