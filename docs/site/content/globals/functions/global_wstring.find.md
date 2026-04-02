# wstring.find

- Category: Global Function
- Confidence level: MEDIUM
- Confidence score: 63/100
- Seen in: 2 addons

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
| Addons seen in | Effigy, TidyChat |
| Files seen in | `/workspace_addons/Effigy/Elements/EffigyLabel.lua:16`, `/workspace_addons/Effigy/States/EffigyStatePlayer.lua:294`, `/workspace_addons/TidyChat/TidyChat.lua:1469` |
| Namespaces detected | wstring |
| Source kinds | lua_calls |
| Example locations | Effigy: Effigy.UpdateTitle, Effigy: Effigy.local.WStringSplitInclDelimiter, Effigy: WStringSplitInclDelimiter, TidyChat: TidyChatLogs.ProcessLootRollEntry |
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
wstring.find(arg1, arg2)
```

## Description

Observed as a global function across 2 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a runtime window or control identifier. | Observed values: inString, text, titleData.name |
| arg2 | Observed as a runtime window or control identifier. | Observed values: L ",", delimiter, find_TEXT_NEEDGREED_ROLL_HEADER |

## Returns

- Not confidently inferable from addon-api docs alone.

## Side Effects

- No side effect is confidently inferable from addon-api docs alone.

## Seen In

- Effigy
- TidyChat

## Examples

- Effigy: Effigy.UpdateTitle -> wstring.find(titleData.name, L ",")
- Effigy: Effigy.local.WStringSplitInclDelimiter -> wstring.find(inString, delimiter, pos, true)
- Effigy: WStringSplitInclDelimiter -> wstring.find(inString, delimiter, pos, true)
- TidyChat: TidyChatLogs.ProcessLootRollEntry -> wstring.find(text, find_TEXT_SELECT)
- TidyChat: TidyChatLogs.ProcessLootRollEntry -> wstring.find(text, find_TEXT_WINNER_HEADER)
- TidyChat: TidyChatLogs.ProcessLootRollEntry -> wstring.find(text, find_TEXT_NEEDGREED_ROLL_HEADER)

## Related APIs

- none

## Used With

- [GameData.Player.activeTitle](../../gamedata/fields/gamedata_GameData.Player.activeTitle.md) (HIGH 100/100) - GameData Field
- [GameData.Player.worldObjNum](../../gamedata/fields/gamedata_GameData.Player.worldObjNum.md) (HIGH 100/100) - GameData Field
- [wstring.match](global_wstring.match.md) (HIGH 100/100) - Global Function
- [wstring.sub](global_wstring.sub.md) (HIGH 75/100) - Global Function

## Triggered By

- none

## Affects

- [GameData.Player.activeTitle](../../gamedata/fields/gamedata_GameData.Player.activeTitle.md) (HIGH 100/100) - GameData Field
- [GameData.Player.worldObjNum](../../gamedata/fields/gamedata_GameData.Player.worldObjNum.md) (HIGH 100/100) - GameData Field

## Notes

- Canonical entry built from observed call sites, not from engine source or decompiled definitions.
