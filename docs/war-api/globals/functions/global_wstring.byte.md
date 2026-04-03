# wstring.byte

- Category: Global Function
- Confidence level: MEDIUM
- Confidence score: 55/100
- Seen in: 1 addons

## Confidence Assessment

- Level: MEDIUM

- Score: 55/100

- Rationale: Promoted as MEDIUM confidence because matches a known engine namespace, called globally with no local definition, argument pattern is consistent.

## Evidence Signals

- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +10 Argument pattern is consistent: Observed argument positions remain stable.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | LibJson |
| Files seen in | Json.lua |
| Namespaces detected | wstring |
| Source kinds | lua_calls |
| Example locations | LibJson: encode_wstring |
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
| Consistent role | no |
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
wstring.byte(arg1, arg2, arg3)
```

## Description

Observed as a global function across 1 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a runtime window or control identifier. | Observed values: str |
| arg2 | Observed as a numeric value. | Observed values: 1 |
| arg3 | Observed as a function or method reference. | Observed values: wstring.len(str) |

## Returns

- Not confidently inferable from contract artifacts alone.

## Side Effects

- No side effect is confidently inferable from contract artifacts alone.

## Seen In

- LibJson

## Examples

- LibJson: encode_wstring -> wstring.byte(str, 1, wstring.len(str))

## Notes

- Only one addon surfaced this symbol in the current corpus.
- Canonical entry built from observed call sites, not from engine source or decompiled definitions.
