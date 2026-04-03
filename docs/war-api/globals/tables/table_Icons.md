# Icons

- Category: Global Table
- Confidence level: HIGH
- Confidence score: 100/100

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 140

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, seen in 4 or more addons, matches a known engine namespace.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +10 Referenced from initialization flow: Lifecycle reconstruction references this symbol.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +25 Referenced by generated docs or reference files: The symbol is reinforced outside a single call page.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | AutoMark, Enemy, PartyCast, Swift Assist |
| Files seen in | `/workspace/data/raw/Enemy/Code/Assist/Assist.lua:195`, `/workspace/data/raw/Enemy/Code/CombatLog/CombatLogIDS.lua:211`, `/workspace/data/raw/Enemy/Code/GroupIcons/GroupIcon.lua:115`, `/workspace/data/raw/Enemy/Code/Marks/MarkTemplate.lua:85`, `/workspace/data/raw/Enemy/Code/UnitFrames/Parts/CareerIcon.lua:13`, `/workspace/data/raw/PartyCast/PartyCast.lua:284`, `/workspace/data/raw/swift-assist/SwiftAssist.lua:232`, `/workspace/data/raw/swift-assist/SwiftAssist.lua:99` |
| Namespaces detected | Icons |
| Source kinds | globals, lua_calls |
| Example locations | Enemy: Enemy.AssistUI_Target_Show, Enemy: Enemy.CombatLogUI_IDS_Update, Enemy: Enemy.UnitFramesParts_CareerIconInitialize, Enemy: EnemyGroupIcon:Attach, Enemy: EnemyMarkTemplate:Apply, PartyCast: PartyCast.Target |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 10 |
| Global usage count | 1 |
| Local definition count | 1 |
| Documentation references | 1 |
| Initialization flow references | 1 |
| Known engine namespace | yes |
| Default UI presence | yes |
| Event binding presence | no |
| Observed in XML and Lua | no |
| Consistent role | yes |
| Consistent arguments | no |
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

## Description

Observed shared global table or namespace surfaced in 4 addons.

## Functions

- Icons.GetCareerIconIDFromCareerLine

## Observed Members

- none

## Seen In

- AutoMark
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

- none

## Used With

- none

## Triggered By

- none

## Affects

- none

## Notes

- none
