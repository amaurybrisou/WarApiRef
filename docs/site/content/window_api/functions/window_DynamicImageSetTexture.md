# DynamicImageSetTexture

- Category: Window Function
- Confidence level: HIGH
- Confidence score: 90/100
- Seen in: 1 addons

## Confidence Assessment

- Level: HIGH

- Score: 90/100

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, matches a known engine namespace, called globally with no local definition.

## Evidence Signals

- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +10 Argument pattern is consistent: Observed argument positions remain stable.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | Moth |
| Files seen in | `/workspace/data/raw/Moth/Moth.lua:227`, `/workspace/data/raw/Moth/Moth.lua:421` |
| Namespaces detected | DynamicImageSetTexture |
| Source kinds | lua_calls |
| Example locations | Moth: Moth.SetCellTextIcon, Moth: Moth.UpdateNPCIcon |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 2 |
| Global usage count | 2 |
| Local definition count | 0 |
| Documentation references | 0 |
| Initialization flow references | 0 |
| Known engine namespace | yes |
| Default UI presence | yes |
| Event binding presence | no |
| Observed in XML and Lua | no |
| Consistent role | no |
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
DynamicImageSetTexture(windowName, texture, textureX, textureY)
```

## Description

Observed mutating runtime image resources on existing controls.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| windowName | Observed as a target image control name. | Observed values: cellName.."CareerIcon", iconWindow |
| texture | Observed as a texture resource name. | Observed values: "map_markers01", careerIconId |
| textureX | Observed as a numeric texture coordinate or atlas offset. | Observed values: 0, 32 |
| textureY | Observed as a numeric texture coordinate or atlas offset. | Observed values: 0, 32 |

## Returns

- Not confidently inferable from addon-api docs alone.

## Side Effects

- Mutates engine-visible UI state or presentation.

## Seen In

- Moth

## Examples

- Moth: Moth.SetCellTextIcon -> DynamicImageSetTexture(cellName.."CareerIcon", careerIconId, 0, 0)
- Moth: Moth.UpdateNPCIcon -> DynamicImageSetTexture(iconWindow, "map_markers01", 32, 32)

## Related APIs

- none

## Used With

- [LabelSetFont](window_LabelSetFont.md) (HIGH 100/100) - Window Function
- [LabelSetText](window_LabelSetText.md) (HIGH 100/100) - Window Function
- [WindowGetDimensions](window_WindowGetDimensions.md) (HIGH 100/100) - Window Function
- [WindowSetShowing](window_WindowSetShowing.md) (HIGH 100/100) - Window Function
- [WindowSetDimensions](window_WindowSetDimensions.md) (HIGH 98/100) - Window Function
- [LabelSetTextColor](window_LabelSetTextColor.md) (HIGH 90/100) - Window Function
- [DynamicImageSetTextureScale](window_DynamicImageSetTextureScale.md) (HIGH 80/100) - Window Function
- [DynamicImageSetTextureSlice](window_DynamicImageSetTextureSlice.md) (HIGH 80/100) - Window Function
- [GetIconData](../../globals/functions/global_GetIconData.md) (HIGH 71/100) - Global Function

## Triggered By

- none

## Affects

- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Notes

- Only one addon surfaced this symbol in the current corpus.
