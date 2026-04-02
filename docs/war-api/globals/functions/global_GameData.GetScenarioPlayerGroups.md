# GameData.GetScenarioPlayerGroups

- Category: Global Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 4 addons

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 168

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, seen in 4 or more addons, matches a known engine namespace.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +10 Argument pattern is consistent: Observed argument positions remain stable.
- +8 Return usage is consistent: Observed as a stable query-style API.
- +25 Referenced by generated docs or reference files: The symbol is reinforced outside a single call page.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | Enemy, LibGroup, Targets, followTheLeader |
| Files seen in | `/workspace_addons/Enemy/Code/Core/Groups/Groups.lua:281`, `/workspace_addons/LibGroup/LibGroup.lua:553`, `/workspace_addons/followTheLeader/followTheLeader.lua:237`, `/workspace_addons/targets/Targets.lua:100` |
| Namespaces detected | GameData |
| Source kinds | globals, lua_calls |
| Example locations | Enemy: Enemy._GroupsUpdate, LibGroup: GetScenarioData, LibGroup: LibGroup.local.GetScenarioData, Targets: Targets.filter_scenario, followTheLeader: followTheLeader.GetScenarioLeader |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 5 |
| Global usage count | 5 |
| Local definition count | 0 |
| Documentation references | 1 |
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

Observed as a global function across 4 addons.

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
- Targets
- followTheLeader

## Examples

- Enemy: Enemy._GroupsUpdate -> GameData.GetScenarioPlayerGroups()
- LibGroup: GetScenarioData -> GameData.GetScenarioPlayerGroups()
- LibGroup: LibGroup.local.GetScenarioData -> GameData.GetScenarioPlayerGroups()
- Targets: Targets.filter_scenario -> GameData.GetScenarioPlayerGroups()
- followTheLeader: followTheLeader.GetScenarioLeader -> GameData.GetScenarioPlayerGroups()

## Related APIs

- [PartyUtils.GetPartyData](global_PartyUtils.GetPartyData.md) (HIGH 100/100) - Global Function
- [PartyUtils.GetWarbandData](global_PartyUtils.GetWarbandData.md) (HIGH 100/100) - Global Function

## Used With

- none

## Triggered By

- none

## Affects

- [PartyUtils.GetPartyData](global_PartyUtils.GetPartyData.md) (HIGH 100/100) - Global Function

## Notes

- Canonical entry built from observed call sites, not from engine source or decompiled definitions.
