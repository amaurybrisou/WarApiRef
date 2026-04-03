# wstring.find

- Category: Global Function
- Confidence level: HIGH
- Confidence score: 75/100
- Seen in: 10 addons

## Confidence Assessment

- Level: HIGH

- Score: 75/100

- Rationale: Promoted as HIGH confidence because seen in 4 or more addons, matches a known engine namespace, called globally with no local definition.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- -15 Conflicting signatures across usages: Observed arity or argument shape conflicts across usages.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | Cram The Spam, DammazKron, Effigy, I HATE YOU THIS MUCH, Pure, Queue Queuer, ResHelp, TidyChat |
| Files seen in | Core/DK_Core.lua, CramTheSpam.lua, Elements/EffigyLabel.lua, IHYTM.lua, QueueQueuer.lua, ResHelp.lua, Source/PureTarget_UnitFrame.lua, Source/TomeTracker_Saga.lua |
| Namespaces detected | wstring |
| Source kinds | lua_calls |
| Example locations | Cram The Spam: FuzzyFilterTest, Cram The Spam: GenericSlashCommand, DammazKron: PrintRegister, Effigy: UpdateTitle, Effigy: WStringSplitInclDelimiter, I HATE YOU THIS MUCH: ChatHandler |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 68 |
| Global usage count | 68 |
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

Observed as a global function across 10 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a runtime window or control identifier. | Observed values: GameData.ChatData.text, LineString, aText |
| arg2 | Observed as a text or wstring payload. | Observed values: L "	", L " ", L " damage" |
| arg3 | Observed as a numeric value. | Observed values: 1, 2, pos |
| arg4 | Observed as a boolean toggle. | Observed values: true |

## Returns

- Not confidently inferable from contract artifacts alone.

## Side Effects

- No side effect is confidently inferable from contract artifacts alone.

## Seen In

- Cram The Spam
- DammazKron
- Effigy
- I HATE YOU THIS MUCH
- Pure
- Queue Queuer
- ResHelp
- TidyChat
- TomeTracker
- Trakario

## Examples

- Cram The Spam: FuzzyFilterTest -> wstring.find(wLowerString, L "%.([%w]*)%.com")
- Cram The Spam: GenericSlashCommand -> wstring.find(wStringParam, L "^([^ ]*) ?(.*)$")
- DammazKron: PrintRegister -> wstring.find(LineString, L "<br>")
- Effigy: UpdateTitle -> wstring.find(titleData.name, L ",")
- Effigy: WStringSplitInclDelimiter -> wstring.find(inString, delimiter, pos, true)
- I HATE YOU THIS MUCH: ChatHandler -> wstring.find(player, author, 1, true)

## Used With

- [wstring.lower](global_wstring.lower.md) (HIGH 100/100) - Global Function
- [wstring.sub](global_wstring.sub.md) (HIGH 75/100) - Global Function

## Affects

- [GameData.ChatData.name](../../gamedata/fields/gamedata_GameData.ChatData.name.md) (HIGH 100/100) - GameData Field
- [GameData.Player.activeTitle](../../gamedata/fields/gamedata_GameData.Player.activeTitle.md) (HIGH 100/100) - GameData Field
- [GameData.Player.worldObjNum](../../gamedata/fields/gamedata_GameData.Player.worldObjNum.md) (HIGH 100/100) - GameData Field

## Notes

- Canonical entry built from observed call sites, not from engine source or decompiled definitions.
