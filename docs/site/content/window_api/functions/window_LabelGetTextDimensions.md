# LabelGetTextDimensions

- Category: Window Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 21 addons

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
| Addons seen in | Calling, DammazKron, DetauntHelper, Enemy, GetStats, HealGrid, Killer, LibAddonButton |
| Files seen in | CallingNotification.lua, Code/Marks/MarkTemplate.lua, Core/ToolTip/DK_Tooltip.lua, Display.lua, GetStats.lua, KillerConfigWindow.lua, KillerEvents.lua, KillerFeed.lua |
| Namespaces detected | LabelGetTextDimensions |
| Source kinds | lua_calls |
| Example locations | Calling: ManageNotification, DammazKron: SetData, DetauntHelper: OnUpdateFontCombo, Enemy: Apply, GetStats: GetSafeTextWidth, HealGrid: Display |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 59 |
| Global usage count | 59 |
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

Observed as a window function across 21 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a function or method reference. | Observed values: "CallWindowText", "KillerPersonalCounterCount", "KillerTooltipAbility" |

## Returns

- Not confidently inferable from contract artifacts alone.

## Side Effects

- No side effect is confidently inferable from contract artifacts alone.

## Seen In

- Calling
- DammazKron
- DetauntHelper
- Enemy
- GetStats
- HealGrid
- Killer
- LibAddonButton
- MapMonster
- MapPin
- Moth
- NaturalLog
- RoR_SoR
- TacticSetNames
- TurretRange
- WSCT
- WarBoard_Loc
- WarBoard_TogglerRankedLeaderboard
- WarBoard_TogglerWARCommander
- WhoHealedMe
- warboard_TogglerWarBuilder

## Examples

- Calling: ManageNotification -> LabelGetTextDimensions("CallWindowText")
- DammazKron: SetData -> LabelGetTextDimensions(windowName.."Name")
- DammazKron: SetData -> LabelGetTextDimensions(windowName.."Career")
- DammazKron: SetData -> LabelGetTextDimensions(windowName.."Level")
- DammazKron: SetData -> LabelGetTextDimensions(windowName.."DeathblowValue")
- DammazKron: SetData -> LabelGetTextDimensions(windowName.."DeathblowTitle")

## Used With

- [DynamicImageSetTextureScale](window_DynamicImageSetTextureScale.md) (HIGH 100/100) - Window Function
- [Icons.GetCareerIconIDFromCareerLine](../../globals/functions/global_Icons.GetCareerIconIDFromCareerLine.md) (HIGH 100/100) - Global Function
- [LabelSetText](window_LabelSetText.md) (HIGH 100/100) - Window Function
- [GetIconData](../../globals/functions/global_GetIconData.md) (HIGH 83/100) - Global Function

## Notes

- none
