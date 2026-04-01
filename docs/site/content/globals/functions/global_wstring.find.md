# wstring.find

- Category: Global Function
- Confidence level: MEDIUM
- Confidence score: 63/100
- Seen in: 3 addons

## Confidence Assessment

- Level: MEDIUM

- Score: 63/100

- Rationale: Promoted as MEDIUM confidence because matches a known engine namespace, called globally with no local definition, seen in 2 to 3 addons.

## Evidence Signals

- +18 Seen in 2 to 3 addons: Cross-addon spread is present but limited.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- -15 Conflicting signatures across usages: Observed arity or argument shape conflicts across usages.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | Effigy, Queue Queuer, TidyChat |
| Files seen in | `/workspace/Effigy/Elements/EffigyLabel.lua:16`, `/workspace/Effigy/States/EffigyStatePlayer.lua:294`, `/workspace/QueueQueuer/QueueQueuer.lua:338`, `/workspace/QueueQueuer/QueueQueuer.lua:377`, `/workspace/TidyChat/TidyChat.lua:1469` |
| Namespaces detected | wstring |
| Source kinds | lua_calls |
| Example locations | Effigy: Effigy.UpdateTitle, Effigy: Effigy.local.WStringSplitInclDelimiter, Effigy: WStringSplitInclDelimiter, Queue Queuer: QueueQueuer.CompareWStrings, Queue Queuer: QueueQueuer.FixName, TidyChat: TidyChatLogs.ProcessLootRollEntry |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 13 |
| Global usage count | 13 |
| Local definition count | 0 |
| Documentation references | 0 |
| Initialization flow references | 0 |
| Known engine namespace | yes |
| Default UI presence | no |
| Event binding presence | no |
| Observed in XML and Lua | no |
| Consistent role | yes |
| Consistent arguments | no |
| Consistent returns | no |
| Slash command presence | no |
| Weak usage only | no |
| Project-specific name | no |
| Placeholder or computed name | no |
| Conflicting signatures | yes |
| Conflicting roles | no |
| Wrapper likely | no |
| Never outside local graph | no |
| Local helper only | no |

## Signature (inferred)

```lua
wstring.find(arg1, arg2, arg3, arg4)
```

## Description

Observed as a global function across 3 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a runtime window or control identifier. | Observed values: first, inString, name |
| arg2 | Observed as a text or wstring payload. | Observed values: L "	", L " ", L "," |
| arg3 | Observed as a numeric value. | Observed values: 1, pos |
| arg4 | Observed as a boolean toggle. | Observed values: true |

## Returns

- Not confidently inferable from API_Ref alone.

## Side Effects

- No side effect is confidently inferable from API_Ref alone.

## Seen In

- Effigy
- Queue Queuer
- TidyChat

## Examples

- Effigy: Effigy.UpdateTitle -> wstring.find(titleData.name, L ",")
- Effigy: Effigy.local.WStringSplitInclDelimiter -> wstring.find(inString, delimiter, pos, true)
- Effigy: WStringSplitInclDelimiter -> wstring.find(inString, delimiter, pos, true)
- Queue Queuer: QueueQueuer.CompareWStrings -> wstring.find(first, second, 1, true)
- Queue Queuer: QueueQueuer.FixName -> wstring.find(name, L "^", 1, true)
- Queue Queuer: QueueQueuer.FixName -> wstring.find(name, L " ", 1, true)

## Related APIs

- none

## Used With

- [GameData.Player.activeTitle](../../gamedata/fields/gamedata_GameData.Player.activeTitle.md) (HIGH 100/100) - GameData Field
- [GameData.Player.worldObjNum](../../gamedata/fields/gamedata_GameData.Player.worldObjNum.md) (HIGH 100/100) - GameData Field
- [wstring.len](global_wstring.len.md) (HIGH 100/100) - Global Function
- [towstring](global_towstring.md) (HIGH 75/100) - Global Function
- [wstring.sub](global_wstring.sub.md) (HIGH 75/100) - Global Function
- [wstring.reverse](global_wstring.reverse.md) (MEDIUM 45/100) - Global Function

## Triggered By

- none

## Affects

- [GameData.Player.activeTitle](../../gamedata/fields/gamedata_GameData.Player.activeTitle.md) (HIGH 100/100) - GameData Field
- [GameData.Player.worldObjNum](../../gamedata/fields/gamedata_GameData.Player.worldObjNum.md) (HIGH 100/100) - GameData Field

## Notes

- Canonical entry built from observed call sites, not from engine source or decompiled definitions.
