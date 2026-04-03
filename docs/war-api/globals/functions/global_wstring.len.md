# wstring.len

- Category: Global Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 4 addons

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
| Addons seen in | CombatTextNames, Enemy, Shinies, WSCT |
| Files seen in | `/workspace/data/raw/Enemy/Code/Core/Utils.lua:684`, `/workspace/data/raw/Shinies/Modules/UI/Shinies-UI-Browse.lua:656`, `/workspace/data/raw/combattextnames/combattextnames.lua:150`, `/workspace/data/raw/wsct/wsct.lua:596` |
| Namespaces detected | wstring |
| Source kinds | lua_calls |
| Example locations | CombatTextNames: CombatTextNames.SetupText, Enemy: Enemy.FormatNumber, Shinies: Criteria_BuildQuery, WSCT: WSCT.ShortenString |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 5 |
| Global usage count | 5 |
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

Observed as a global function across 4 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a runtime window or control identifier. | Observed values: hitData.name, maxItemLevel, minItemLevel |

## Returns

- Not confidently inferable from addon-api docs alone.

## Side Effects

- No side effect is confidently inferable from addon-api docs alone.

## Seen In

- CombatTextNames
- Enemy
- Shinies
- WSCT

## Examples

- CombatTextNames: CombatTextNames.SetupText -> wstring.len(hitData.name)
- Enemy: Enemy.FormatNumber -> wstring.len(number_string)
- Shinies: Criteria_BuildQuery -> wstring.len(minItemLevel)
- Shinies: Criteria_BuildQuery -> wstring.len(maxItemLevel)
- WSCT: WSCT.ShortenString -> wstring.len(strString)

## Related APIs

- none

## Used With

- [wstring.sub](global_wstring.sub.md) (HIGH 100/100) - Global Function

## Triggered By

- none

## Affects

- none

## Notes

- Canonical entry built from observed call sites, not from engine source or decompiled definitions.
