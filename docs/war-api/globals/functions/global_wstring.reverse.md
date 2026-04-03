# wstring.reverse

- Category: Global Function
- Confidence level: HIGH
- Confidence score: 88/100
- Seen in: 3 addons

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
| Addons seen in | I HATE YOU THIS MUCH, Queue Queuer, WarBoard_WarWhisperer |
| Files seen in | IHYTM.lua, QueueQueuer.lua, warwhisperer.lua |
| Namespaces detected | wstring |
| Source kinds | lua_calls |
| Example locations | I HATE YOU THIS MUCH: FixName, I HATE YOU THIS MUCH: SendSyncMessage, Queue Queuer: FixName, WarBoard_WarWhisperer: FixName |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 4 |
| Global usage count | 4 |
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
wstring.reverse(arg1)
```

## Description

Observed as a global function across 3 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a runtime window or control identifier. | Observed values: msg, name |

## Returns

- Not confidently inferable from contract artifacts alone.

## Side Effects

- No side effect is confidently inferable from contract artifacts alone.

## Seen In

- I HATE YOU THIS MUCH
- Queue Queuer
- WarBoard_WarWhisperer

## Examples

- I HATE YOU THIS MUCH: FixName -> wstring.reverse(name)
- I HATE YOU THIS MUCH: SendSyncMessage -> wstring.reverse(msg)
- Queue Queuer: FixName -> wstring.reverse(name)
- WarBoard_WarWhisperer: FixName -> wstring.reverse(name)

## Used With

- [wstring.find](global_wstring.find.md) (HIGH 75/100) - Global Function
- [wstring.sub](global_wstring.sub.md) (HIGH 75/100) - Global Function

## Notes

- Canonical entry built from observed call sites, not from engine source or decompiled definitions.
