# PartyUtils.GetPartyData

- Category: Global Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 17 addons

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 108

- Rationale: Promoted as HIGH confidence because seen in 4 or more addons, matches a known engine namespace, called globally with no local definition.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +10 Argument pattern is consistent: Observed argument positions remain stable.
- +8 Return usage is consistent: Observed as a stable query-style API.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | Arsenal Rank, AutoBand, DeepSleep, EZGuard, Enemy, EveryBodyGuard, Info_DeathBlow, LibGroup |
| Files seen in | ArsenalRank.lua, AutoBand.lua, Code/Core/Groups/Groups.lua, EZGuard.lua, EveryBodyGuard.lua, Info_DeathBlow.lua, LibGroup.lua, Modules.lua |
| Namespaces detected | PartyUtils |
| Source kinds | lua_calls |
| Example locations | Arsenal Rank: isInGroup, AutoBand: cmd_dump_party_data, DeepSleep: AddGroupMenuItems, EZGuard: UpdateGroup, Enemy: _GroupsUpdate, EveryBodyGuard: ActivateGuard_Hard |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 30 |
| Global usage count | 30 |
| Local definition count | 0 |
| Documentation references | 0 |
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
PartyUtils.GetPartyData()
```

## Description

Observed as a global function across 17 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |

## Returns

- Observed as a query-style API. The concrete return shape is not inferable from contract artifacts alone.

## Side Effects

- No side effect is confidently inferable from contract artifacts alone.

## Seen In

- Arsenal Rank
- AutoBand
- DeepSleep
- EZGuard
- Enemy
- EveryBodyGuard
- Info_DeathBlow
- LibGroup
- LibGuard
- NoUselessMods-Assist
- PartyCast
- Pure
- Squared
- WarTriage
- WhoHealedMe
- followTheLeader
- wbLeadHelper

## Examples

- Arsenal Rank: isInGroup -> PartyUtils.GetPartyData()
- AutoBand: cmd_dump_party_data -> PartyUtils.GetPartyData()
- DeepSleep: AddGroupMenuItems -> PartyUtils.GetPartyData()
- EZGuard: UpdateGroup -> PartyUtils.GetPartyData()
- Enemy: _GroupsUpdate -> PartyUtils.GetPartyData()
- EveryBodyGuard: ActivateGuard_Hard -> PartyUtils.GetPartyData()

## Affects

- [GameData.GetScenarioPlayerGroups](global_GameData.GetScenarioPlayerGroups.md) (HIGH 100/100) - Global Function

## Notes

- Canonical entry built from observed call sites, not from engine source or decompiled definitions.
- Advanced return analysis: No strong return evidence observed
