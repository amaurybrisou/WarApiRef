# GameData.GetScenarioPlayerGroups

- Category: Global Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 11 addons

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 143

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, seen in 4 or more addons, matches a known engine namespace.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +10 Argument pattern is consistent: Observed argument positions remain stable.
- +8 Return usage is consistent: Observed as a stable query-style API.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | AutoBand, Calling, EA_ScenarioGroupWindow, Enemy, HealGrid, Hopper, LibGroup, Squared |
| Files seen in | AB_util.lua, AB_wb.lua, Calling.lua, Code/Core/Groups/Groups.lua, GridFactories/HealGridGridFactory.lua, HealGridUtility.lua, LibGroup.lua, Source/Hopper.lua |
| Namespaces detected | GameData |
| Source kinds | lua_calls |
| Example locations | AutoBand: get_hot_scdata, AutoBand: load_from_scdata, Calling: IsPlayerInGroup, EA_ScenarioGroupWindow: ScenarioGroupHasOpenSlot, EA_ScenarioGroupWindow: UpdateMainAssist, EA_ScenarioGroupWindow: UpdatePlayerData |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 18 |
| Global usage count | 18 |
| Local definition count | 0 |
| Documentation references | 0 |
| Initialization flow references | 0 |
| Known engine namespace | yes |
| Default UI presence | yes |
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
GameData.GetScenarioPlayerGroups()
```

## Description

Observed as a global function across 11 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |

## Returns

- Observed as a query-style API. The concrete return shape is not inferable from contract artifacts alone.

## Side Effects

- No side effect is confidently inferable from contract artifacts alone.

## Seen In

- AutoBand
- Calling
- EA_ScenarioGroupWindow
- Enemy
- HealGrid
- Hopper
- LibGroup
- Squared
- Targets
- WarTriage
- followTheLeader

## Examples

- AutoBand: get_hot_scdata -> GameData.GetScenarioPlayerGroups()
- AutoBand: load_from_scdata -> GameData.GetScenarioPlayerGroups()
- Calling: IsPlayerInGroup -> GameData.GetScenarioPlayerGroups()
- EA_ScenarioGroupWindow: ScenarioGroupHasOpenSlot -> GameData.GetScenarioPlayerGroups()
- EA_ScenarioGroupWindow: UpdateMainAssist -> GameData.GetScenarioPlayerGroups()
- EA_ScenarioGroupWindow: UpdatePlayerData -> GameData.GetScenarioPlayerGroups()

## Related APIs

- [PartyUtils.GetPartyData](global_PartyUtils.GetPartyData.md) (HIGH 100/100) - Global Function
- [PartyUtils.GetWarbandData](global_PartyUtils.GetWarbandData.md) (HIGH 100/100) - Global Function

## Notes

- Canonical entry built from observed call sites, not from engine source or decompiled definitions.
