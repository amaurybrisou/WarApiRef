# PartyUtils.GetPartyData

- Category: Global Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 6 addons

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
| Addons seen in | Enemy, LibGroup, LibGuard, WhoHealedMe, followTheLeader, wbLeadHelper |
| Files seen in | `/workspace_addons/Enemy/Code/Core/Groups/Groups.lua:281`, `/workspace_addons/LibGroup/LibGroup.lua:538`, `/workspace_addons/LibGuard/Source/LibGuard.lua:186`, `/workspace_addons/WhoHealedMe/WHMCore.lua:201`, `/workspace_addons/followTheLeader/followTheLeader.lua:209`, `/workspace_addons/followTheLeader/followTheLeader.lua:262`, `/workspace_addons/wbLeadHelper/wbLeadHelper.lua:996` |
| Namespaces detected | PartyUtils |
| Source kinds | globals, lua_calls |
| Example locations | Enemy: Enemy._GroupsUpdate, LibGroup: GetPartyData, LibGroup: LibGroup.local.GetPartyData, LibGuard: LibGuard.GetIdFromName, WhoHealedMe: WHMCore.IsHealerInPartyOrWarband, followTheLeader: followTheLeader.GetDynamicLeaderName |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 8 |
| Global usage count | 8 |
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
PartyUtils.GetPartyData()
```

## Description

Observed as a global function across 6 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |

## Returns

- Observed as a query-style API. The concrete return shape is not inferable from addon-api docs alone.

## Side Effects

- No side effect is confidently inferable from addon-api docs alone.

## Seen In

- Enemy
- LibGroup
- LibGuard
- WhoHealedMe
- followTheLeader
- wbLeadHelper

## Examples

- Enemy: Enemy._GroupsUpdate -> PartyUtils.GetPartyData()
- LibGroup: GetPartyData -> PartyUtils.GetPartyData()
- LibGroup: LibGroup.local.GetPartyData -> PartyUtils.GetPartyData()
- LibGuard: LibGuard.GetIdFromName -> PartyUtils.GetPartyData()
- WhoHealedMe: WHMCore.IsHealerInPartyOrWarband -> PartyUtils.GetPartyData()
- followTheLeader: followTheLeader.GetDynamicLeaderName -> PartyUtils.GetPartyData()

## Related APIs

- [GameData.GetScenarioPlayerGroups](global_GameData.GetScenarioPlayerGroups.md) (HIGH 100/100) - Global Function
- [PartyUtils.GetWarbandData](global_PartyUtils.GetWarbandData.md) (HIGH 100/100) - Global Function
- [PartyUtils.IsPartyActive](global_PartyUtils.IsPartyActive.md) (HIGH 100/100) - Global Function

## Used With

- [PartyUtils.GetWarbandData](global_PartyUtils.GetWarbandData.md) (HIGH 100/100) - Global Function

## Triggered By

- none

## Affects

- [GameData.GetScenarioPlayerGroups](global_GameData.GetScenarioPlayerGroups.md) (HIGH 100/100) - Global Function
- [GameData.Player.name](../../gamedata/fields/gamedata_GameData.Player.name.md) (HIGH 100/100) - GameData Field
- [PartyUtils.GetWarbandData](global_PartyUtils.GetWarbandData.md) (HIGH 100/100) - Global Function
- [PartyUtils.GetWarbandLeader](global_PartyUtils.GetWarbandLeader.md) (HIGH 100/100) - Global Function

## Notes

- Canonical entry built from observed call sites, not from engine source or decompiled definitions.
- Advanced return analysis: No strong return evidence observed
