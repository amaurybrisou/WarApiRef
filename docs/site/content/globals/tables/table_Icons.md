# Icons

- Category: Global Table
- Confidence level: HIGH
- Confidence score: 100/100

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 125

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, seen in 4 or more addons, matches a known engine namespace.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | AutoMark, BuddyBind, Calling, CleanUnitFrames, DammazKron, DetauntHelper, Ding, EA_ScenarioGroupWindow |
| Files seen in | Code/Assist/Assist.lua, Code/CombatLog/CombatLogIDS.lua, Code/GroupIcons/GroupIcon.lua, Code/Marks/MarkTemplate.lua, Code/UnitFrames/Parts/CareerIcon.lua, Core/Tome/DK_Tome.lua, Core/ToolTip/DK_Tooltip.lua, Elements/EffigyIcons.lua |
| Namespaces detected | Icons |
| Source kinds | lua_calls |
| Example locations | AutoMark: CreateMarker, BuddyBind: GrabName, Calling: SetListLine, Calling: ShowCallerIcon, Calling: ShowTargetIcon, CleanUnitFrames: SetCareerIcon |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 82 |
| Global usage count | 2 |
| Local definition count | 0 |
| Documentation references | 0 |
| Initialization flow references | 0 |
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

Shared function table with 2 member functions; the primary API surface for 31 addons.

## Functions

- Icons.GetCareerIconIDFromCareerLine
- Icons.GetCareerIconIDFromCareerNamesID

## Observed Members

- none

## Seen In

- AutoMark
- BuddyBind
- Calling
- CleanUnitFrames
- DammazKron
- DetauntHelper
- Ding
- EA_ScenarioGroupWindow
- Effigy
- Enemy
- EveryBodyGuard
- Group Icons
- Group Icons SG
- GuildWarden
- HealGrid
- MapPin
- MarkBuff
- Moth
- PartyCast
- Pure
- QuickNameActions+
- ResHelp
- SessionRPs
- Squared
- Swift Assist
- TargetInfoRing
- Targets
- ThinkOutLoud
- Trakario
- xHUD
- yAssistHelper

## Examples

- AutoMark: CreateMarker -> Icons.GetCareerIconIDFromCareerLine(career_id)
- BuddyBind: GrabName -> Icons.GetCareerIconIDFromCareerLine(TargetInfo:UnitCareer("selffriendlytarget"))
- Calling: SetListLine -> Icons.GetCareerIconIDFromCareerLine(call.CallerCareerID)
- Calling: SetListLine -> Icons.GetCareerIconIDFromCareerLine(call.TargetCareerID)
- Calling: ShowCallerIcon -> Icons.GetCareerIconIDFromCareerLine(careerId)
- Calling: ShowTargetIcon -> Icons.GetCareerIconIDFromCareerLine(careerId)

## Notes

- none
