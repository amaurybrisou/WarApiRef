# wstring.upper

- Category: Global Function
- Confidence level: HIGH
- Confidence score: 88/100
- Seen in: 2 addons

## Confidence Assessment

- Level: HIGH

- Score: 88/100

- Rationale: Promoted as HIGH confidence because matches a known engine namespace, called globally with no local definition, seen in 2 to 3 addons.

## Evidence Signals

- +18 Seen in 2 to 3 addons: Cross-addon spread is present but limited.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +10 Argument pattern is consistent: Observed argument positions remain stable.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | Moth, Queue Queuer |
| Files seen in | `/workspace/Moth/MothHelpers.lua:39`, `/workspace/QueueQueuer/QueueQueuer.lua:773` |
| Namespaces detected | wstring |
| Source kinds | lua_calls |
| Example locations | Moth: MothHelpers.CapitalizeWString, Queue Queuer: QueueQueuer.OnUpdate |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 2 |
| Global usage count | 2 |
| Local definition count | 0 |
| Documentation references | 0 |
| Initialization flow references | 0 |
| Known engine namespace | yes |
| Default UI presence | no |
| Event binding presence | no |
| Observed in XML and Lua | no |
| Consistent role | yes |
| Consistent arguments | yes |
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

## Signature (inferred)

```lua
wstring.upper(arg1)
```

## Description

Observed as a global function across 2 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a function or method reference. | Observed values: LabelGetText("EA_Window_ScenarioJoinPromptBoxName"), wstring.sub(wstr,1,1) |

## Returns

- Not confidently inferable from addon-api docs alone.

## Side Effects

- No side effect is confidently inferable from addon-api docs alone.

## Seen In

- Moth
- Queue Queuer

## Examples

- Moth: MothHelpers.CapitalizeWString -> wstring.upper(wstring.sub(wstr,1,1))
- Queue Queuer: QueueQueuer.OnUpdate -> wstring.upper(LabelGetText("EA_Window_ScenarioJoinPromptBoxName"))

## Related APIs

- none

## Used With

- [wstring.sub](global_wstring.sub.md) (HIGH 75/100) - Global Function
- [wstring.lower](global_wstring.lower.md) (MEDIUM 45/100) - Global Function

## Triggered By

- none

## Affects

- [GameData.Player.zone](../../gamedata/fields/gamedata_GameData.Player.zone.md) (HIGH 100/100) - GameData Field
- [SystemData.Events.INTERACT_SELECT_SCENARIO_QUEUE_LIST](../../systemdata/fields/systemdata_SystemData.Events.INTERACT_SELECT_SCENARIO_QUEUE_LIST.md) (HIGH 100/100) - SystemData Field

## Notes

- Canonical entry built from observed call sites, not from engine source or decompiled definitions.
