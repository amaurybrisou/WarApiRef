# GetIconData

- Category: Global Function
- Confidence level: HIGH
- Confidence score: 71/100
- Seen in: 3 addons

## Confidence Assessment

- Level: HIGH

- Score: 71/100

- Rationale: Promoted as HIGH confidence because called globally with no local definition, seen in 2 to 3 addons, role is consistent across addons.

## Evidence Signals

- +18 Seen in 2 to 3 addons: Cross-addon spread is present but limited.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +10 Argument pattern is consistent: Observed argument positions remain stable.
- +8 Return usage is consistent: Observed as a stable query-style API.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | Moth, PartyCast, TidyRoll |
| Files seen in | `/workspace/data/raw/Moth/Moth.lua:227`, `/workspace/data/raw/PartyCast/PartyCast.lua:399`, `/workspace/data/raw/TidyRoll/TidyRollFrame.lua:134`, `/workspace/data/raw/TidyRoll/TidyRollFrame.lua:180` |
| Namespaces detected | GetIconData |
| Source kinds | lua_calls |
| Example locations | Moth: Moth.SetCellTextIcon, PartyCast: PartyCast.FetchedText, TidyRoll: TidyRollFrame:SetCareerIcon, TidyRoll: TidyRollFrame:SetIcon |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 5 |
| Global usage count | 5 |
| Local definition count | 0 |
| Documentation references | 0 |
| Initialization flow references | 0 |
| Known engine namespace | no |
| Default UI presence | no |
| Event binding presence | no |
| Observed in XML and Lua | no |
| Consistent role | yes |
| Consistent arguments | yes |
| Consistent returns | yes |
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
GetIconData(arg1)
```

## Description

Observed as a shared query API across 3 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a runtime window or control identifier. | Observed values: Icons.GetCareerIconIDFromCareerLine(text), careerIconNum, iconNum |

## Returns

- Observed as a query-style API. The concrete return shape is not inferable from addon-api docs alone.

## Side Effects

- No side effect is confidently inferable from addon-api docs alone.

## Seen In

- Moth
- PartyCast
- TidyRoll

## Examples

- Moth: Moth.SetCellTextIcon -> GetIconData(Icons.GetCareerIconIDFromCareerLine(text))
- PartyCast: PartyCast.FetchedText -> GetIconData(tonumber(SPLIT_TEXT[7]))
- PartyCast: PartyCast.FetchedText -> GetIconData(tonumber(Target_Icon))
- TidyRoll: TidyRollFrame:SetCareerIcon -> GetIconData(careerIconNum)
- TidyRoll: TidyRollFrame:SetIcon -> GetIconData(iconNum)

## Related APIs

- none

## Used With

- [DynamicImageSetTextureDimensions](../../window_api/functions/window_DynamicImageSetTextureDimensions.md) (HIGH 100/100) - Window Function
- [LabelSetFont](../../window_api/functions/window_LabelSetFont.md) (HIGH 100/100) - Window Function
- [LabelSetText](../../window_api/functions/window_LabelSetText.md) (HIGH 100/100) - Window Function
- [LabelSetTextColor](../../window_api/functions/window_LabelSetTextColor.md) (HIGH 100/100) - Window Function
- [DynamicImageSetTexture](../../window_api/functions/window_DynamicImageSetTexture.md) (HIGH 98/100) - Window Function
- [Icons.GetCareerIconIDFromCareerLine](global_Icons.GetCareerIconIDFromCareerLine.md) (HIGH 88/100) - Global Function
- [towstring](global_towstring.md) (HIGH 75/100) - Global Function

## Triggered By

- none

## Affects

- [DefaultColor](../tables/table_DefaultColor.md) (HIGH 100/100) - Global Table

## Notes

- Canonical entry built from observed call sites, not from engine source or decompiled definitions.
