# wstring.upper

- Category: Global Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 7 addons

## Confidence Assessment

- Level: HIGH

- Score: 100/100

- Rationale: Promoted as HIGH confidence because seen in 4 or more addons, matches a known engine namespace, called globally with no local definition.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +10 Argument pattern is consistent: Observed argument positions remain stable.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | EA_LoadingScreen, Effigy, HealGrid, Info_Loot, Moth, Queue Queuer, Refer |
| Files seen in | Elements/EffigyLabel.lua, HealGridTags.lua, Info_Loot.lua, MothHelpers.lua, QueueQueuer.lua, Refer.lua, source/nodataloadingscreen.lua, source/scenarioenterloadingscreen.lua |
| Namespaces detected | wstring |
| Source kinds | lua_calls |
| Example locations | EA_LoadingScreen: BeginNoDataLoadScreen, EA_LoadingScreen: BeginScenarioEnterLoadScreen, EA_LoadingScreen: BeginStandardLoadScreen, Effigy: Render, HealGrid: ProcessTagUpper, Info_Loot: TextArrived |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 10 |
| Global usage count | 10 |
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

Observed as a global function across 7 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a runtime window or control identifier. | Observed values: LabelGetText("EA_Window_ScenarioJoinPromptBoxName"), lootbag, scenarioName |

## Returns

- Not confidently inferable from contract artifacts alone.

## Side Effects

- No side effect is confidently inferable from contract artifacts alone.

## Seen In

- EA_LoadingScreen
- Effigy
- HealGrid
- Info_Loot
- Moth
- Queue Queuer
- Refer

## Examples

- EA_LoadingScreen: BeginNoDataLoadScreen -> wstring.upper(text)
- EA_LoadingScreen: BeginScenarioEnterLoadScreen -> wstring.upper(scenarioName)
- EA_LoadingScreen: BeginStandardLoadScreen -> wstring.upper(zoneName)
- Effigy: Render -> wstring.upper(text)
- HealGrid: ProcessTagUpper -> wstring.upper(str)
- Info_Loot: TextArrived -> wstring.upper(lootbag)

## Related APIs

- [OnUpdate](../../xml/handlers/handler_OnUpdate.md) (HIGH 88/100) - XML Event

## Used With

- [InterfaceCore.GetResolutionScale](global_InterfaceCore.GetResolutionScale.md) (HIGH 100/100) - Global Function
- [LabelSetText](../../window_api/functions/window_LabelSetText.md) (HIGH 100/100) - Window Function
- [WindowGetDimensions](../../window_api/functions/window_WindowGetDimensions.md) (HIGH 100/100) - Window Function

## Affects

- [GameData.ChatData.text](../../gamedata/fields/gamedata_GameData.ChatData.text.md) (HIGH 100/100) - GameData Field

## Notes

- Canonical entry built from observed call sites, not from engine source or decompiled definitions.
