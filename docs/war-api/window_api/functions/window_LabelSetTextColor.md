# LabelSetTextColor

- Category: Window Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 4 addons

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
| Addons seen in | InfoScroller, Moth, PartyCast, minesweep |
| Files seen in | `/workspace/data/raw/InfoScroller/libs/LibGUI.lua:463`, `/workspace/data/raw/Moth/Moth.lua:215`, `/workspace/data/raw/Moth/Moth.lua:227`, `/workspace/data/raw/PartyCast/PartyCast.lua:399`, `/workspace/data/raw/PartyCast/libs/LibGUI.lua:463`, `/workspace/data/raw/minesweep/minesweep.lua:109` |
| Namespaces detected | LabelSetTextColor |
| Source kinds | lua_calls |
| Example locations | InfoScroller: LIBGUI_Label:Color, Moth: Moth.SetCellText, Moth: Moth.SetCellTextIcon, PartyCast: LIBGUI_Label:Color, PartyCast: PartyCast.FetchedText, minesweep: minesweep.ShowCell |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 10 |
| Global usage count | 10 |
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
LabelSetTextColor(arg1, arg2, arg3, arg4)
```

## Description

Observed updating label text or label styling on existing controls.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a text or wstring payload. | Observed values: "PartyCastWindow"..PlayerNumber.."Name", WindowName.."Text", cellText |
| arg2 | Observed as a function or method reference. | Observed values: DefaultColor.GOLD.r, PartyCast.Settings.Colors.Ability[1], PartyCast.Settings.Colors.Failed[1] |
| arg3 | Observed as a function or method reference. | Observed values: DefaultColor.GOLD.g, PartyCast.Settings.Colors.Ability[2], PartyCast.Settings.Colors.Failed[2] |
| arg4 | Observed as a function or method reference. | Observed values: DefaultColor.GOLD.b, PartyCast.Settings.Colors.Ability[3], PartyCast.Settings.Colors.Failed[3] |

## Returns

- Not confidently inferable from addon-api docs alone.

## Side Effects

- Mutates engine-visible UI state or presentation.

## Seen In

- InfoScroller
- Moth
- PartyCast
- minesweep

## Examples

- InfoScroller: LIBGUI_Label:Color -> LabelSetTextColor(self.name, red, green, blue)
- Moth: Moth.SetCellText -> LabelSetTextColor(cellText, color.r, color.g, color.b)
- Moth: Moth.SetCellTextIcon -> LabelSetTextColor(cellText, color.r, color.g, color.b)
- PartyCast: LIBGUI_Label:Color -> LabelSetTextColor(self.name, red, green, blue)
- PartyCast: PartyCast.FetchedText -> LabelSetTextColor("PartyCastWindow"..PlayerNumber.."Name", DefaultColor.GOLD.r, DefaultColor.GOLD.g, DefaultColor.GOLD.b)
- PartyCast: PartyCast.FetchedText -> LabelSetTextColor("PartyCastWindow"..PlayerNumber.."Name", PartyCast.Settings.Colors.Ability[1], PartyCast.Settings.Colors.Ability[2], PartyCast.Settings.Colors.Ability[3])

## Related APIs

- none

## Used With

- [DynamicImageSetTextureDimensions](window_DynamicImageSetTextureDimensions.md) (HIGH 100/100) - Window Function
- [LabelGetTextColor](window_LabelGetTextColor.md) (HIGH 100/100) - Window Function
- [LabelSetFont](window_LabelSetFont.md) (HIGH 100/100) - Window Function
- [LabelSetText](window_LabelSetText.md) (HIGH 100/100) - Window Function
- [WindowSetDimensions](window_WindowSetDimensions.md) (HIGH 100/100) - Window Function
- [DynamicImageSetTexture](window_DynamicImageSetTexture.md) (HIGH 98/100) - Window Function
- [Icons.GetCareerIconIDFromCareerLine](../../globals/functions/global_Icons.GetCareerIconIDFromCareerLine.md) (HIGH 88/100) - Global Function
- [LabelGetTextDimensions](window_LabelGetTextDimensions.md) (HIGH 80/100) - Window Function
- [towstring](../../globals/functions/global_towstring.md) (HIGH 75/100) - Global Function
- [GetIconData](../../globals/functions/global_GetIconData.md) (HIGH 71/100) - Global Function

## Triggered By

- none

## Affects

- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Notes

- none
