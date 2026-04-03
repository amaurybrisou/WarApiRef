# DynamicImageSetTextureScale

- Category: Window Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 6 addons

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
| Addons seen in | Ace, LibWBToggler, PartyCast, Shinies, TexturedButtons, WoH-Reticle |
| Files seen in | `/workspace/data/raw/Ace/LibGUI.lua:1326`, `/workspace/data/raw/LibWarBoardToggler/LibWBToggler.lua:16`, `/workspace/data/raw/LibWarBoardToggler/libs/LibGUI.lua:1323`, `/workspace/data/raw/PartyCast/libs/LibGUI.lua:1323`, `/workspace/data/raw/Shinies/Libraries/LibGUI.lua:1323`, `/workspace/data/raw/TexturedButtons/CooldownPulse.lua:51`, `/workspace/data/raw/WoH-Reticle/libs/LibGUI.lua:1323` |
| Namespaces detected | DynamicImageSetTextureScale |
| Source kinds | lua_calls |
| Example locations | Ace: LIBGUI_Image:TexScale, LibWBToggler: LIBGUI_Image:TexScale, LibWBToggler: LibWBToggler.CreateToggler, PartyCast: LIBGUI_Image:TexScale, Shinies: LIBGUI_Image:TexScale, TexturedButtons: CooldownPulse:SetIcon |
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
DynamicImageSetTextureScale(arg1, arg2)
```

## Description

Observed mutating runtime image resources on existing controls.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a function or method reference. | Observed values: modName.."Icon", self.name, self:GetName().."Icon" |
| arg2 | Observed as a runtime window or control identifier. | Observed values: 1, scale, texScale |

## Returns

- Not confidently inferable from addon-api docs alone.

## Side Effects

- Mutates engine-visible UI state or presentation.

## Seen In

- Ace
- LibWBToggler
- PartyCast
- Shinies
- TexturedButtons
- WoH-Reticle

## Examples

- Ace: LIBGUI_Image:TexScale -> DynamicImageSetTextureScale(self.name, scale)
- LibWBToggler: LIBGUI_Image:TexScale -> DynamicImageSetTextureScale(self.name, scale)
- LibWBToggler: LibWBToggler.CreateToggler -> DynamicImageSetTextureScale(modName.."Icon", texScale)
- PartyCast: LIBGUI_Image:TexScale -> DynamicImageSetTextureScale(self.name, scale)
- Shinies: LIBGUI_Image:TexScale -> DynamicImageSetTextureScale(self.name, scale)
- TexturedButtons: CooldownPulse:SetIcon -> DynamicImageSetTextureScale(self:GetName().."Icon", 1)

## Related APIs

- none

## Used With

- none

## Triggered By

- none

## Affects

- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Notes

- none
