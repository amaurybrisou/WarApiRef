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
| Addons seen in | AutoMark, Effigy, Enemy, MapPin, Moth, Swift Assist, Targets |
| Files seen in | `/workspace_addons/AutoMark/Source/AutoMark.lua:7`, `/workspace_addons/Effigy/Elements/EffigyIcons.lua:138`, `/workspace_addons/Effigy/Elements/EffigyIcons.lua:172`, `/workspace_addons/Enemy/Code/Assist/Assist.lua:195`, `/workspace_addons/Enemy/Code/CombatLog/CombatLogIDS.lua:211`, `/workspace_addons/Enemy/Code/GroupIcons/GroupIcon.lua:115`, `/workspace_addons/Enemy/Code/Marks/MarkTemplate.lua:85`, `/workspace_addons/Enemy/Code/UnitFrames/Parts/CareerIcon.lua:13` |
| Namespaces detected | Icons |
| Source kinds | globals, lua_calls |
| Example locations | AutoMark: AutoMark.local.CreateMarker, AutoMark: CreateMarker, Effigy: Effigy.local.RenderCareerIcon, Effigy: Effigy.local.SetupCareerIcon, Effigy: RenderCareerIcon, Effigy: SetupCareerIcon |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 22 |
| Global usage count | 1 |
| Local definition count | 4 |
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

Observed shared global table or namespace surfaced in 7 addons.

## Functions

- Icons.GetCareerIconIDFromCareerLine

## Observed Members

- none

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

- none

## Used With

- none

## Triggered By

- none

## Affects

- none

## Notes

- none
