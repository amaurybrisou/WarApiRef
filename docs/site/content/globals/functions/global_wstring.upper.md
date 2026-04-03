# wstring.upper

- Category: Global Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 6 addons

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
| Addons seen in | Effigy, HealGrid, Info_Loot, Moth, Queue Queuer, Refer |
| Files seen in | Elements/EffigyLabel.lua, HealGridTags.lua, Info_Loot.lua, MothHelpers.lua, QueueQueuer.lua, Refer.lua |
| Namespaces detected | wstring |
| Source kinds | lua_calls |
| Example locations | Effigy: Render, HealGrid: ProcessTagUpper, Info_Loot: TextArrived, Moth: CapitalizeWString, Queue Queuer: OnUpdate, Refer: FormatPlayerName |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 7 |
| Global usage count | 7 |
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

Observed as a global function across 6 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a runtime window or control identifier. | Observed values: LabelGetText("EA_Window_ScenarioJoinPromptBoxName"), lootbag, str |

## Returns

- Not confidently inferable from contract artifacts alone.

## Side Effects

- No side effect is confidently inferable from contract artifacts alone.

## Seen In

- Effigy
- HealGrid
- Info_Loot
- Moth
- Queue Queuer
- Refer

## Examples

- Effigy: Render -> wstring.upper(text)
- HealGrid: ProcessTagUpper -> wstring.upper(str)
- Info_Loot: TextArrived -> wstring.upper(lootbag)
- Moth: CapitalizeWString -> wstring.upper(wstring.sub(wstr,1,1))
- Queue Queuer: OnUpdate -> wstring.upper(LabelGetText("EA_Window_ScenarioJoinPromptBoxName"))
- Refer: FormatPlayerName -> wstring.upper(wstring.sub(nameLower,1,1))

## Related APIs

- [OnUpdate](../../xml/handlers/handler_OnUpdate.md) (HIGH 88/100) - XML Event

## Used With

- [wstring.lower](global_wstring.lower.md) (HIGH 100/100) - Global Function
- [wstring.sub](global_wstring.sub.md) (HIGH 75/100) - Global Function

## Affects

- [GameData.ChatData.text](../../gamedata/fields/gamedata_GameData.ChatData.text.md) (HIGH 100/100) - GameData Field
- [GameData.Player.zone](../../gamedata/fields/gamedata_GameData.Player.zone.md) (HIGH 100/100) - GameData Field
- [SystemData.Events.INTERACT_SELECT_SCENARIO_QUEUE_LIST](../../systemdata/fields/systemdata_SystemData.Events.INTERACT_SELECT_SCENARIO_QUEUE_LIST.md) (HIGH 100/100) - SystemData Field

## Notes

- Canonical entry built from observed call sites, not from engine source or decompiled definitions.
