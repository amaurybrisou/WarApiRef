# wstring.len

- Category: Global Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 14 addons

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
| Addons seen in | CleanUnitFrames, CombatTextNames, Effigy, Enemy, HealGrid, I HATE YOU THIS MUCH, Kwestor, LibJson |
| Files seen in | CleanTargetUnitFrame.lua, Code/Core/Utils.lua, Elements/EffigyLabel.lua, Gui/HealGridGuiTabSpellTrack.lua, HealGrid.lua, HealGridTags.lua, IHYTM.lua, Json.lua |
| Namespaces detected | wstring |
| Source kinds | lua_calls |
| Example locations | CleanUnitFrames: UpdateUnit, CombatTextNames: SetupText, Effigy: wstringStartsWith, Enemy: FormatNumber, HealGrid: ExtractArg, HealGrid: FindLabel |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 25 |
| Global usage count | 25 |
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
wstring.len(arg1)
```

## Description

Observed as a global function across 14 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a runtime window or control identifier. | Observed values: GameData.Player.name, Start, TargetInfo:UnitName("selffriendlytarget") |

## Returns

- Not confidently inferable from contract artifacts alone.

## Side Effects

- No side effect is confidently inferable from contract artifacts alone.

## Seen In

- CleanUnitFrames
- CombatTextNames
- Effigy
- Enemy
- HealGrid
- I HATE YOU THIS MUCH
- Kwestor
- LibJson
- MyReasons
- NaturalLog
- Queue Queuer
- Shinies
- WSCT
- WarBoard_WarWhisperer

## Examples

- CleanUnitFrames: UpdateUnit -> wstring.len(nameText)
- CombatTextNames: SetupText -> wstring.len(hitData.name)
- Effigy: wstringStartsWith -> wstring.len(Start)
- Enemy: FormatNumber -> wstring.len(number_string)
- HealGrid: ExtractArg -> wstring.len(s)
- HealGrid: FindLabel -> wstring.len(s)

## Related APIs

- [OnUpdate](../../xml/handlers/handler_OnUpdate.md) (HIGH 88/100) - XML Event

## Used With

- [wstring.sub](global_wstring.sub.md) (HIGH 75/100) - Global Function

## Notes

- Canonical entry built from observed call sites, not from engine source or decompiled definitions.
