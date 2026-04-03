# PartyUtils.GetWarbandData

- Category: Global Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 9 addons

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
| Addons seen in | EA_OpenPartyWindow, Enemy, Hopper, LibGroup, LibGuard, MegaphonePlusPlus, Squared, WarTriage |
| Files seen in | Code/Core/Groups/Groups.lua, LibGroup.lua, MegaphonePlusPlus.lua, Source/Hopper.lua, Source/LibGuard.lua, SquaredWarband.lua, WHMCore.lua, WarTriage.lua |
| Namespaces detected | PartyUtils |
| Source kinds | lua_calls |
| Example locations | EA_OpenPartyWindow: CountWarband, EA_OpenPartyWindow: UpdateWarband, Enemy: _GroupsUpdate, Hopper: UpdateWarbandOpenSlots, LibGroup: GetWarbandData, LibGuard: GetIdFromName |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 11 |
| Global usage count | 11 |
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
PartyUtils.GetWarbandData()
```

## Description

Observed as a global function across 9 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |

## Returns

- Observed as a query-style API. The concrete return shape is not inferable from contract artifacts alone.

## Side Effects

- No side effect is confidently inferable from contract artifacts alone.

## Seen In

- EA_OpenPartyWindow
- Enemy
- Hopper
- LibGroup
- LibGuard
- MegaphonePlusPlus
- Squared
- WarTriage
- WhoHealedMe

## Examples

- EA_OpenPartyWindow: CountWarband -> PartyUtils.GetWarbandData()
- EA_OpenPartyWindow: UpdateWarband -> PartyUtils.GetWarbandData()
- Enemy: _GroupsUpdate -> PartyUtils.GetWarbandData()
- Hopper: UpdateWarbandOpenSlots -> PartyUtils.GetWarbandData()
- LibGroup: GetWarbandData -> PartyUtils.GetWarbandData()
- LibGuard: GetIdFromName -> PartyUtils.GetWarbandData()

## Affects

- [GameData.GetScenarioPlayerGroups](global_GameData.GetScenarioPlayerGroups.md) (HIGH 100/100) - Global Function

## Notes

- Canonical entry built from observed call sites, not from engine source or decompiled definitions.
- Advanced return analysis: No strong return evidence observed
