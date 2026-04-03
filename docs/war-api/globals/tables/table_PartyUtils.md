# PartyUtils

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
| Addons seen in | Arsenal Rank, AutoBand, CMap, DeepSleep, EA_OpenPartyWindow, EZGuard, Enemy, EveryBodyGuard |
| Files seen in | Code/Core/Groups/EnemyPlayer.lua, Code/Core/Groups/Groups.lua, Source/Hopper.lua, Source/LibGuard.lua, Source/PureGroup.lua, Source/PureGroupPet.lua, source/MapPin.lua, source/openpartywindow.lua |
| Namespaces detected | PartyUtils |
| Source kinds | lua_calls |
| Example locations | Arsenal Rank: isInGroup, AutoBand: cmd_dump_party_data, CMap: CmapLButton, CMap: TidyQueueLButton, DeepSleep: AddGroupMenuItems, EA_OpenPartyWindow: CountWarband |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 97 |
| Global usage count | 15 |
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

Shared function table with 15 member functions; the primary API surface for 33 addons.

## Functions

- PartyUtils.GetBolsterMenuText
- PartyUtils.GetPartyData
- PartyUtils.GetPartyMember
- PartyUtils.GetWarbandData
- PartyUtils.GetWarbandLeader
- PartyUtils.GetWarbandMainAssist
- PartyUtils.GetWarbandMember
- PartyUtils.GetWarbandParty
- PartyUtils.IsPartyActive
- PartyUtils.IsPartyMemberValid
- PartyUtils.IsPlayerInWarband
- PartyUtils.IsWarbandFull
- PartyUtils.MoveWarbandMember
- PartyUtils.SetMasterLooter
- PartyUtils.SwapWarbandMembers

## Observed Members

- none

## Seen In

- Arsenal Rank
- AutoBand
- CMap
- DeepSleep
- EA_OpenPartyWindow
- EZGuard
- Enemy
- EveryBodyGuard
- Group Icons SG
- GuardLine
- Hopper
- Info_DeathBlow
- LibGroup
- LibGuard
- MapPin
- MegaphonePlusPlus
- NoUselessMods-Assist
- PartyCast
- Pure
- Queue Queuer
- QuickWarReport
- Refer
- ResHelp
- RoR_SoR
- SocialWindow 2.0
- Squared
- Swift Assist
- TokenMachine
- WarTriage
- WhoHealedMe
- followTheLeader
- rorAutoInviter
- wbLeadHelper

## Examples

- Arsenal Rank: isInGroup -> PartyUtils.GetPartyData()
- AutoBand: cmd_dump_party_data -> PartyUtils.GetPartyData()
- CMap: CmapLButton -> PartyUtils.GetWarbandLeader()
- CMap: CmapLButton -> PartyUtils.IsPartyActive()
- CMap: TidyQueueLButton -> PartyUtils.GetWarbandLeader()
- CMap: TidyQueueLButton -> PartyUtils.IsPartyActive()

## Notes

- none
