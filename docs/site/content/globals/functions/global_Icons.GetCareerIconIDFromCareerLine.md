# Icons.GetCareerIconIDFromCareerLine

- Category: Global Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 7 addons

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 133

- Rationale: Promoted as HIGH confidence because seen in 4 or more addons, matches a known engine namespace, referenced by generated docs or reference files.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +10 Argument pattern is consistent: Observed argument positions remain stable.
- +8 Return usage is consistent: Observed as a stable query-style API.
- +25 Referenced by generated docs or reference files: The symbol is reinforced outside a single call page.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | AutoMark, Effigy, Enemy, MapPin, Moth, Swift Assist, Targets |
| Files seen in | `/workspace_addons/AutoMark/Source/AutoMark.lua:7`, `/workspace_addons/Effigy/Elements/EffigyIcons.lua:138`, `/workspace_addons/Effigy/Elements/EffigyIcons.lua:172`, `/workspace_addons/Enemy/Code/Assist/Assist.lua:195`, `/workspace_addons/Enemy/Code/CombatLog/CombatLogIDS.lua:211`, `/workspace_addons/Enemy/Code/GroupIcons/GroupIcon.lua:115`, `/workspace_addons/Enemy/Code/Marks/MarkTemplate.lua:85`, `/workspace_addons/Enemy/Code/UnitFrames/Parts/CareerIcon.lua:13` |
| Namespaces detected | Icons |
| Source kinds | globals, lua_calls |
| Example locations | AutoMark: AutoMark.local.CreateMarker, AutoMark: CreateMarker, Effigy: Effigy.local.RenderCareerIcon, Effigy: Effigy.local.SetupCareerIcon, Effigy: RenderCareerIcon, Effigy: SetupCareerIcon |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 22 |
| Global usage count | 22 |
| Local definition count | 0 |
| Documentation references | 1 |
| Initialization flow references | 0 |
| Known engine namespace | yes |
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
Icons.GetCareerIconIDFromCareerLine(arg1)
```

## Description

Observed as a global function across 7 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a runtime window or control identifier. | Observed values: Addon.career, GameData.Player.career.line, cache.career |

## Returns

- Observed as a query-style API. The concrete return shape is not inferable from addon-api docs alone.

## Side Effects

- No side effect is confidently inferable from addon-api docs alone.

## Seen In

- AutoMark
- Effigy
- Enemy
- MapPin
- Moth
- Swift Assist
- Targets

## Examples

- AutoMark: AutoMark.local.CreateMarker -> Icons.GetCareerIconIDFromCareerLine(career_id)
- AutoMark: CreateMarker -> Icons.GetCareerIconIDFromCareerLine(career_id)
- Effigy: Effigy.local.RenderCareerIcon -> Icons.GetCareerIconIDFromCareerLine(value)
- Effigy: Effigy.local.SetupCareerIcon -> Icons.GetCareerIconIDFromCareerLine(GameData.Player.career.line)
- Effigy: RenderCareerIcon -> Icons.GetCareerIconIDFromCareerLine(value)
- Effigy: SetupCareerIcon -> Icons.GetCareerIconIDFromCareerLine(GameData.Player.career.line)

## Related APIs

- [DynamicImageSetTexture](../../window_api/functions/window_DynamicImageSetTexture.md) (HIGH 100/100) - Window Function
- [GetIconData](global_GetIconData.md) (HIGH 83/100) - Global Function

## Used With

- [DynamicImageSetTexture](../../window_api/functions/window_DynamicImageSetTexture.md) (HIGH 100/100) - Window Function
- [WindowSetShowing](../../window_api/functions/window_WindowSetShowing.md) (HIGH 100/100) - Window Function
- [GetIconData](global_GetIconData.md) (HIGH 83/100) - Global Function
- [towstring](global_towstring.md) (HIGH 75/100) - Global Function

## Triggered By

- none

## Affects

- [GameData.Player.career.line](../../gamedata/fields/gamedata_GameData.Player.career.line.md) (HIGH 100/100) - GameData Field

## Notes

- Canonical entry built from observed call sites, not from engine source or decompiled definitions.
