# DynamicImageSetTextureDimensions

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
| Addons seen in | Ace, BuffHead, Enemy, LibWBToggler, PartyCast, PotionBar, Shinies, TexturedButtons |
| Files seen in | `/workspace/data/raw/Ace/LibGUI.lua:1304`, `/workspace/data/raw/BuffHead/Setup/SetupLayoutProperties.lua:117`, `/workspace/data/raw/Enemy/Code/UnitFrames/Parts/CareerIcon.lua:13`, `/workspace/data/raw/Enemy/Code/UnitFrames/UnitFramePart.lua:393`, `/workspace/data/raw/Enemy/Code/UnitFrames/UnitFramePart.lua:407`, `/workspace/data/raw/LibWarBoardToggler/libs/LibGUI.lua:1301`, `/workspace/data/raw/PartyCast/PartyCast.lua:399`, `/workspace/data/raw/PartyCast/libs/LibGUI.lua:1301` |
| Namespaces detected | DynamicImageSetTextureDimensions |
| Source kinds | lua_calls |
| Example locations | Ace: LIBGUI_Image:TexDims, BuffHead: BuffHead.local.SetTextureButton, BuffHead: SetTextureButton, Enemy: Enemy.UnitFramePart_ApplyTexture, Enemy: Enemy.UnitFramePart_ApplyTexturePercentResize, Enemy: Enemy.UnitFramesParts_CareerIconInitialize |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 18 |
| Global usage count | 18 |
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
| arg1 | Observed as a function or method reference. | Observed values: "PartyCastWindow"..PlayerNumber.."TargetWindowIcon", adat.iconname, iconWindowName |
| arg2 | Observed as a runtime window or control identifier. | Observed values: 0, 256, 32 |
| arg3 | Observed as a runtime window or control identifier. | Observed values: 0, 256, 32 |

## Returns

- Not confidently inferable from addon-api docs alone.

## Side Effects

- Mutates engine-visible UI state or presentation.

## Seen In

- Ace
- BuffHead
- Enemy
- LibWBToggler
- PartyCast
- PotionBar
- Shinies
- TexturedButtons
- WSCT
- WoH-Reticle

## Examples

- Ace: LIBGUI_Image:TexDims -> DynamicImageSetTextureDimensions(self.name, width, height)
- BuffHead: BuffHead.local.SetTextureButton -> DynamicImageSetTextureDimensions(textureButton.."Texture", dimensions.Width, dimensions.Height)
- BuffHead: SetTextureButton -> DynamicImageSetTextureDimensions(textureButton.."Texture", dimensions.Width, dimensions.Height)
- Enemy: Enemy.UnitFramePart_ApplyTexture -> DynamicImageSetTextureDimensions(windowName, texture.dx, texture.dy)
- Enemy: Enemy.UnitFramePart_ApplyTexturePercentResize -> DynamicImageSetTextureDimensions(t.windowName, value*texture.dx, texture.dy)
- Enemy: Enemy.UnitFramePart_ApplyTexturePercentResize -> DynamicImageSetTextureDimensions(t.windowName, texture.dx, value*texture.dy)

## Related APIs

- none

## Used With

- [DynamicImageSetTexture](window_DynamicImageSetTexture.md) (HIGH 100/100) - Window Function
- [LabelSetText](window_LabelSetText.md) (HIGH 100/100) - Window Function
- [PartyUtils.GetPartyData](../../globals/functions/global_PartyUtils.GetPartyData.md) (HIGH 100/100) - Global Function
- [GetIconData](../../globals/functions/global_GetIconData.md) (HIGH 83/100) - Global Function
- [towstring](../../globals/functions/global_towstring.md) (HIGH 75/100) - Global Function

## Triggered By

- none

## Affects

- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Notes

- none
