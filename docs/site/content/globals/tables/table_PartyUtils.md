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
| Addons seen in | DeepSleep, Enemy, GuardLine, LibGroup, LibGuard, MapPin, MegaphonePlusPlus, Queue Queuer |
| Files seen in | `/workspace/DeepSleep/Modules.lua:99`, `/workspace/Enemy/Code/Core/Groups/EnemyPlayer.lua:57`, `/workspace/Enemy/Code/Core/Groups/Groups.lua:281`, `/workspace/Enemy/Code/Core/Groups/Groups.lua:523`, `/workspace/Enemy/Code/Core/Groups/Groups.lua:558`, `/workspace/GuardLine/GuardLine.lua:197`, `/workspace/LibGroup/LibGroup.lua:538`, `/workspace/LibGroup/LibGroup.lua:660` |
| Namespaces detected | PartyUtils |
| Source kinds | globals, lua_calls |
| Example locations | DeepSleep: AddGroupMenuItems, DeepSleep: PlayerMenuWindow.AddGroupMenuItems, Enemy: Enemy.Groups_OnBattlegroupMemberUpdated, Enemy: Enemy.Groups_OnGroupStatusUpdated, Enemy: Enemy._GroupsUpdate, Enemy: EnemyPlayer:IsMainAssist |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 50 |
| Global usage count | 8 |
| Local definition count | 9 |
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

Observed shared global table or namespace surfaced in 15 addons.

## Functions

- PartyUtils.GetPartyData
- PartyUtils.GetPartyMember
- PartyUtils.GetWarbandData
- PartyUtils.GetWarbandLeader
- PartyUtils.GetWarbandMainAssist
- PartyUtils.GetWarbandMember
- PartyUtils.IsPartyActive
- PartyUtils.IsPlayerInWarband

## Observed Members

- none

## Seen In

- DeepSleep
- Enemy
- GuardLine
- LibGroup
- LibGuard
- MapPin
- MegaphonePlusPlus
- Queue Queuer
- QuickWarReport
- RoR_SoR
- Swift Assist
- WarTriage
- WhoHealedMe
- followTheLeader
- wbLeadHelper

## Examples

- DeepSleep: AddGroupMenuItems -> PartyUtils.GetPartyData()
- DeepSleep: PlayerMenuWindow.AddGroupMenuItems -> PartyUtils.GetPartyData()
- Enemy: Enemy.Groups_OnBattlegroupMemberUpdated -> PartyUtils.GetWarbandMember(groupIndex, memberIndex)
- Enemy: Enemy.Groups_OnGroupStatusUpdated -> PartyUtils.GetPartyMember(memberIndex)
- Enemy: Enemy._GroupsUpdate -> PartyUtils.GetWarbandData()
- Enemy: Enemy._GroupsUpdate -> PartyUtils.GetPartyData()

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
