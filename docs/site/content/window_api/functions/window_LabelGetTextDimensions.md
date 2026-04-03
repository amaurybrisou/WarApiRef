# LabelGetTextDimensions

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
| Addons seen in | Enemy, Killer, RoR_SoR, TurretRange, WSCT, WhoHealedMe |
| Files seen in | `/workspace/data/raw/Enemy/Code/Marks/MarkTemplate.lua:85`, `/workspace/data/raw/Killer/KillerConfigWindow.lua:125`, `/workspace/data/raw/Killer/KillerEvents.lua:133`, `/workspace/data/raw/Killer/KillerFeed.lua:380`, `/workspace/data/raw/RoR_SoR/RoR_SoR.lua:1055`, `/workspace/data/raw/RoR_SoR/RoR_SoR.lua:364`, `/workspace/data/raw/RoR_SoR/RoR_SoR.lua:461`, `/workspace/data/raw/RoR_SoR/RoR_SoR.lua:536` |
| Namespaces detected | LabelGetTextDimensions |
| Source kinds | lua_calls |
| Example locations | Enemy: EnemyMarkTemplate:Apply, Killer: Killer.RefreshPersonalCounter, Killer: Killer.ShowRowTooltip, Killer: Killer.local.MeasureLabelWidth, Killer: MeasureLabelWidth, RoR_SoR: RoR_SoR.SET_CAMPAIGN |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 20 |
| Global usage count | 20 |
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
LabelGetTextDimensions(arg1)
```

## Description

Observed as a window function across 6 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a text or wstring payload. | Observed values: "KillerPersonalCounterCount", "KillerTooltipAbility", "SoR_"..Window_Name.."CLAIM_WINDOW1TEXTPROXY" |

## Returns

- Not confidently inferable from addon-api docs alone.

## Side Effects

- No side effect is confidently inferable from addon-api docs alone.

## Seen In

- Enemy
- Killer
- RoR_SoR
- TurretRange
- WSCT
- WhoHealedMe

## Examples

- Enemy: EnemyMarkTemplate:Apply -> LabelGetTextDimensions(windowName.."ContentText")
- Killer: Killer.RefreshPersonalCounter -> LabelGetTextDimensions("KillerPersonalCounterCount")
- Killer: Killer.ShowRowTooltip -> LabelGetTextDimensions("KillerTooltipAbility")
- Killer: Killer.local.MeasureLabelWidth -> LabelGetTextDimensions(labelName)
- Killer: MeasureLabelWidth -> LabelGetTextDimensions(labelName)
- RoR_SoR: RoR_SoR.SET_CAMPAIGN -> LabelGetTextDimensions("SoR_"..Window_Name.."_BannerLabel")

## Related APIs

- none

## Used With

- none

## Triggered By

- none

## Affects

- [GameData.PlayerActions.SET_TARGET](../../gamedata/fields/gamedata_GameData.PlayerActions.SET_TARGET.md) (HIGH 100/100) - GameData Field
- [SystemData.MouseOverWindow.name](../../systemdata/fields/systemdata_SystemData.MouseOverWindow.name.md) (HIGH 100/100) - SystemData Field
- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Notes

- none
