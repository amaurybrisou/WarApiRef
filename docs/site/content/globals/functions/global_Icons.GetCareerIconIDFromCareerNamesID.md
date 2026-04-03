# Icons.GetCareerIconIDFromCareerNamesID

- Category: Global Function
- Confidence level: HIGH
- Confidence score: 96/100
- Seen in: 3 addons

## Confidence Assessment

- Level: HIGH

- Score: 96/100

- Rationale: Promoted as HIGH confidence because matches a known engine namespace, called globally with no local definition, seen in 2 to 3 addons.

## Evidence Signals

- +18 Seen in 2 to 3 addons: Cross-addon spread is present but limited.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +10 Argument pattern is consistent: Observed argument positions remain stable.
- +8 Return usage is consistent: Observed as a stable query-style API.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | DetauntHelper, EA_ScenarioGroupWindow, Trakario |
| Files seen in | Source/Config/BarSettings.lua, Source/ScenarioGroupWindow.lua, Source/TargetInfo.lua, trakario.lua |
| Namespaces detected | Icons |
| Source kinds | lua_calls |
| Example locations | DetauntHelper: GetCareerIconDataFromCareerID, DetauntHelper: InitUI, DetauntHelper: _GetCareerIcon, EA_ScenarioGroupWindow: UpdateSingleMemberWindow, Trakario: SetCareerIcon |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 6 |
| Global usage count | 6 |
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
Icons.GetCareerIconIDFromCareerNamesID(arg1)
```

## Description

Observed as a global function across 3 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a function or method reference. | Observed values: GameData.Player.career.id, careerId, cid |

## Returns

- Observed as a query-style API. The concrete return shape is not inferable from contract artifacts alone.

## Side Effects

- No side effect is confidently inferable from contract artifacts alone.

## Seen In

- DetauntHelper
- EA_ScenarioGroupWindow
- Trakario

## Examples

- DetauntHelper: GetCareerIconDataFromCareerID -> Icons.GetCareerIconIDFromCareerNamesID(cid)
- DetauntHelper: InitUI -> Icons.GetCareerIconIDFromCareerNamesID(GameData.Player.career.id)
- DetauntHelper: _GetCareerIcon -> Icons.GetCareerIconIDFromCareerNamesID(self.pinfo.careerId)
- EA_ScenarioGroupWindow: UpdateSingleMemberWindow -> Icons.GetCareerIconIDFromCareerNamesID(player.careerId)
- Trakario: SetCareerIcon -> Icons.GetCareerIconIDFromCareerNamesID(careerId)

## Affects

- [GameData.PlayerActions.SET_TARGET](../../gamedata/fields/gamedata_GameData.PlayerActions.SET_TARGET.md) (HIGH 100/100) - GameData Field

## Notes

- Canonical entry built from observed call sites, not from engine source or decompiled definitions.
