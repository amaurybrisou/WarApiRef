# DynamicImageSetTextureDimensions

- Category: Window Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 2 addons

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 123

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, matches a known engine namespace, called globally with no local definition.

## Evidence Signals

- +18 Seen in 2 to 3 addons: Cross-addon spread is present but limited.
- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +10 Argument pattern is consistent: Observed argument positions remain stable.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | InfoScroller, PartyCast |
| Files seen in | `/workspace/data/raw/InfoScroller/libs/LibGUI.lua:1301`, `/workspace/data/raw/PartyCast/PartyCast.lua:399`, `/workspace/data/raw/PartyCast/libs/LibGUI.lua:1301` |
| Namespaces detected | DynamicImageSetTextureDimensions |
| Source kinds | lua_calls |
| Example locations | InfoScroller: LIBGUI_Image:TexDims, PartyCast: LIBGUI_Image:TexDims, PartyCast: PartyCast.FetchedText |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 4 |
| Global usage count | 4 |
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
| arg1 | Observed as a function or method reference. | Observed values: "PartyCastWindow"..PlayerNumber.."TargetWindowIcon", self.name |
| arg2 | Observed as a numeric value. | Observed values: 0, 64, width |
| arg3 | Observed as a runtime window or control identifier. | Observed values: 0, 64, height |

## Returns

- Not confidently inferable from addon-api docs alone.

## Side Effects

- Mutates engine-visible UI state or presentation.

## Seen In

- InfoScroller
- PartyCast

## Examples

- InfoScroller: LIBGUI_Image:TexDims -> DynamicImageSetTextureDimensions(self.name, width, height)
- PartyCast: LIBGUI_Image:TexDims -> DynamicImageSetTextureDimensions(self.name, width, height)
- PartyCast: PartyCast.FetchedText -> DynamicImageSetTextureDimensions("PartyCastWindow"..PlayerNumber.."TargetWindowIcon", 64, 64)
- PartyCast: PartyCast.FetchedText -> DynamicImageSetTextureDimensions("PartyCastWindow"..PlayerNumber.."TargetWindowIcon", 0, 0)

## Related APIs

- none

## Used With

- [LabelSetText](window_LabelSetText.md) (HIGH 100/100) - Window Function
- [LabelSetTextColor](window_LabelSetTextColor.md) (HIGH 100/100) - Window Function
- [DynamicImageSetTexture](window_DynamicImageSetTexture.md) (HIGH 98/100) - Window Function
- [towstring](../../globals/functions/global_towstring.md) (HIGH 75/100) - Global Function
- [GetIconData](../../globals/functions/global_GetIconData.md) (HIGH 71/100) - Global Function

## Triggered By

- none

## Affects

- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Notes

- none
