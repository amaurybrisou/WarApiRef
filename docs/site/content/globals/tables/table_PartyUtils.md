# PartyUtils

- Category: Global Table
- Confidence level: HIGH
- Confidence score: 100/100

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 130

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, seen in 4 or more addons, matches a known engine namespace.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +25 Referenced by generated docs or reference files: The symbol is reinforced outside a single call page.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | Enemy, GuardLine, LibGroup, LibGuard, PartyCast, RoR_SoR, Swift Assist, WhoHealedMe |
| Files seen in | `/workspace/data/raw/Enemy/Code/Core/Groups/EnemyPlayer.lua:57`, `/workspace/data/raw/Enemy/Code/Core/Groups/Groups.lua:281`, `/workspace/data/raw/Enemy/Code/Core/Groups/Groups.lua:523`, `/workspace/data/raw/Enemy/Code/Core/Groups/Groups.lua:558`, `/workspace/data/raw/GuardLine/GuardLine.lua:197`, `/workspace/data/raw/LibGroup/LibGroup.lua:538`, `/workspace/data/raw/LibGroup/LibGroup.lua:660`, `/workspace/data/raw/LibGroup/LibGroup.lua:788` |
| Namespaces detected | PartyUtils |
| Source kinds | globals, lua_calls |
| Example locations | Enemy: Enemy.Groups_OnBattlegroupMemberUpdated, Enemy: Enemy.Groups_OnGroupStatusUpdated, Enemy: Enemy._GroupsUpdate, Enemy: EnemyPlayer:IsMainAssist, GuardLine: GuardLine.update, LibGroup: GetPartyData |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 27 |
| Global usage count | 7 |
| Local definition count | 6 |
| Documentation references | 1 |
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

Observed shared global table or namespace surfaced in 9 addons.

## Functions

- PartyUtils.GetPartyData
- PartyUtils.GetPartyMember
- PartyUtils.GetWarbandData
- PartyUtils.GetWarbandLeader
- PartyUtils.GetWarbandMainAssist
- PartyUtils.GetWarbandMember
- PartyUtils.IsPartyActive

## Observed Members

- none

## Seen In

- Enemy
- GuardLine
- LibGroup
- LibGuard
- PartyCast
- RoR_SoR
- Swift Assist
- WhoHealedMe
- followTheLeader

## Examples

- Enemy: Enemy.Groups_OnBattlegroupMemberUpdated -> PartyUtils.GetWarbandMember(groupIndex, memberIndex)
- Enemy: Enemy.Groups_OnGroupStatusUpdated -> PartyUtils.GetPartyMember(memberIndex)
- Enemy: Enemy._GroupsUpdate -> PartyUtils.GetWarbandData()
- Enemy: Enemy._GroupsUpdate -> PartyUtils.GetPartyData()
- Enemy: EnemyPlayer:IsMainAssist -> PartyUtils.GetWarbandMainAssist()
- GuardLine: GuardLine.update -> PartyUtils.IsPartyActive()

## Related APIs

- [PartyUtils.GetPartyMember](../functions/global_PartyUtils.GetPartyMember.md) (HIGH 100/100) - Global Function
- [PartyUtils.GetWarbandMember](../functions/global_PartyUtils.GetWarbandMember.md) (HIGH 100/100) - Global Function

## Used With

- none

## Triggered By

- none

## Affects

- none

## Notes

- none
