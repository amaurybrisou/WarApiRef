# GameData.GetScenarioPlayers

- Category: Global Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 4 addons

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
| Addons seen in | Arsenal Rank, DetauntHelper, Enemy, Trakario |
| Files seen in | ArsenalRank.lua, Code/ScenarioInfo/ScenarioInfo.lua, Source/TargetInfo.lua, trakario.lua |
| Namespaces detected | GameData |
| Source kinds | lua_calls |
| Example locations | Arsenal Rank: update, DetauntHelper: _UpdateScenarioInfo, Enemy: ScenarioInfoUpdateData, Trakario: GetCareerIDByName, Trakario: GetDeathCountByName |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 5 |
| Global usage count | 5 |
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
GameData.GetScenarioPlayers()
```

## Description

Observed as a global function across 4 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |

## Returns

- Observed as a query-style API. The concrete return shape is not inferable from contract artifacts alone.

## Side Effects

- No side effect is confidently inferable from contract artifacts alone.

## Seen In

- Arsenal Rank
- DetauntHelper
- Enemy
- Trakario

## Examples

- Arsenal Rank: update -> GameData.GetScenarioPlayers()
- DetauntHelper: _UpdateScenarioInfo -> GameData.GetScenarioPlayers()
- Enemy: ScenarioInfoUpdateData -> GameData.GetScenarioPlayers()
- Trakario: GetCareerIDByName -> GameData.GetScenarioPlayers()
- Trakario: GetDeathCountByName -> GameData.GetScenarioPlayers()

## Affects

- [GameData.Account.ServerName](../../gamedata/fields/gamedata_GameData.Account.ServerName.md) (HIGH 100/100) - GameData Field
- [GameData.Player.name](../../gamedata/fields/gamedata_GameData.Player.name.md) (HIGH 100/100) - GameData Field
- [GameData.ScenarioData](../../gamedata/fields/gamedata_GameData.ScenarioData.md) (HIGH 100/100) - GameData Field
- [GameData.ScenarioData.id](../../gamedata/fields/gamedata_GameData.ScenarioData.id.md) (HIGH 100/100) - GameData Field
- [SystemData.Events.SCENARIO_START_UPDATING_PLAYERS_STATS](../../systemdata/fields/systemdata_SystemData.Events.SCENARIO_START_UPDATING_PLAYERS_STATS.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.SCENARIO_STOP_UPDATING_PLAYERS_STATS](../../systemdata/fields/systemdata_SystemData.Events.SCENARIO_STOP_UPDATING_PLAYERS_STATS.md) (HIGH 100/100) - SystemData Field

## Notes

- Canonical entry built from observed call sites, not from engine source or decompiled definitions.
- Advanced return analysis: No strong return evidence observed
