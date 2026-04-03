# Icons.GetCareerIconIDFromCareerLine

- Category: Global Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 3 addons

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 121

- Rationale: Promoted as HIGH confidence because matches a known engine namespace, referenced by generated docs or reference files, called globally with no local definition.

## Evidence Signals

- +18 Seen in 2 to 3 addons: Cross-addon spread is present but limited.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +10 Argument pattern is consistent: Observed argument positions remain stable.
- +8 Return usage is consistent: Observed as a stable query-style API.
- +25 Referenced by generated docs or reference files: The symbol is reinforced outside a single call page.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | Enemy, PartyCast, Swift Assist |
| Files seen in | `/workspace/data/raw/Enemy/Code/Assist/Assist.lua:195`, `/workspace/data/raw/Enemy/Code/CombatLog/CombatLogIDS.lua:211`, `/workspace/data/raw/Enemy/Code/GroupIcons/GroupIcon.lua:115`, `/workspace/data/raw/Enemy/Code/Marks/MarkTemplate.lua:85`, `/workspace/data/raw/Enemy/Code/UnitFrames/Parts/CareerIcon.lua:13`, `/workspace/data/raw/PartyCast/PartyCast.lua:284`, `/workspace/data/raw/swift-assist/SwiftAssist.lua:232`, `/workspace/data/raw/swift-assist/SwiftAssist.lua:99` |
| Namespaces detected | Icons |
| Source kinds | globals, lua_calls |
| Example locations | Enemy: Enemy.AssistUI_Target_Show, Enemy: Enemy.CombatLogUI_IDS_Update, Enemy: Enemy.UnitFramesParts_CareerIconInitialize, Enemy: EnemyGroupIcon:Attach, Enemy: EnemyMarkTemplate:Apply, PartyCast: PartyCast.Target |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 13 |
| Global usage count | 13 |
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

Observed as a global function across 3 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a runtime window or control identifier. | Observed values: Addon.career, PlayerCareer, cache.career |

## Returns

- Observed as a query-style API. The concrete return shape is not inferable from addon-api docs alone.

## Side Effects

- No side effect is confidently inferable from addon-api docs alone.

## Seen In

- Enemy
- PartyCast
- Swift Assist

## Examples

- Enemy: Enemy.AssistUI_Target_Show -> Icons.GetCareerIconIDFromCareerLine(careerId)
- Enemy: Enemy.CombatLogUI_IDS_Update -> Icons.GetCareerIconIDFromCareerLine(career)
- Enemy: Enemy.UnitFramesParts_CareerIconInitialize -> Icons.GetCareerIconIDFromCareerLine(cache.career)
- Enemy: EnemyGroupIcon:Attach -> Icons.GetCareerIconIDFromCareerLine(self.playerCareer)
- Enemy: EnemyMarkTemplate:Apply -> Icons.GetCareerIconIDFromCareerLine(careerId)
- PartyCast: PartyCast.Target -> Icons.GetCareerIconIDFromCareerLine(PlayerCareer)

## Related APIs

- [DynamicImageSetTexture](../../window_api/functions/window_DynamicImageSetTexture.md) (HIGH 100/100) - Window Function
- [GetIconData](global_GetIconData.md) (HIGH 83/100) - Global Function

## Used With

- [DynamicImageSetTexture](../../window_api/functions/window_DynamicImageSetTexture.md) (HIGH 100/100) - Window Function
- [WindowSetShowing](../../window_api/functions/window_WindowSetShowing.md) (HIGH 100/100) - Window Function
- [GetIconData](global_GetIconData.md) (HIGH 83/100) - Global Function
- [CreateWindowFromTemplate](global_CreateWindowFromTemplate.md) (HIGH 75/100) - Global Function

## Triggered By

- none

## Affects

- [GameData.PlayerActions.SET_TARGET](../../gamedata/fields/gamedata_GameData.PlayerActions.SET_TARGET.md) (HIGH 100/100) - GameData Field

## Notes

- Canonical entry built from observed call sites, not from engine source or decompiled definitions.
