# wstring.sub

- Category: Global Function
- Confidence level: MEDIUM
- Confidence score: 40/100
- Seen in: 1 addons

## Confidence Assessment

- Level: MEDIUM

- Final score: 40/100

- Raw weighted score: 30

- Rationale: Promoted as MEDIUM confidence because matches a known engine namespace, called globally with no local definition, known namespace override.

## Evidence Signals

- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- -15 Conflicting signatures across usages: Observed arity or argument shape conflicts across usages.
- +10 Known namespace override: Known engine namespaces are promoted to at least MEDIUM when strong contradictory evidence is absent.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | Moth |
| Files seen in | `/workspace/data/raw/Moth/Moth.lua:270`, `/workspace/data/raw/Moth/MothHelpers.lua:39` |
| Namespaces detected | wstring |
| Source kinds | lua_calls |
| Example locations | Moth: Moth.UpdateLevel, Moth: MothHelpers.CapitalizeWString |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 3 |
| Global usage count | 3 |
| Local definition count | 0 |
| Documentation references | 0 |
| Initialization flow references | 0 |
| Known engine namespace | yes |
| Default UI presence | no |
| Event binding presence | no |
| Observed in XML and Lua | no |
| Consistent role | no |
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
wstring.sub(arg1, arg2, arg3)
```

## Description

Observed as a global function across 1 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a runtime window or control identifier. | Observed values: unitTierDesc, wstr |
| arg2 | Observed as a numeric value. | Observed values: 1, 2 |
| arg3 | Observed as a numeric value. | Observed values: 1 |

## Returns

- Not confidently inferable from addon-api docs alone.

## Side Effects

- No side effect is confidently inferable from addon-api docs alone.

## Seen In

- Moth

## Examples

- Moth: Moth.UpdateLevel -> wstring.sub(unitTierDesc, 1, 1)
- Moth: MothHelpers.CapitalizeWString -> wstring.sub(wstr, 2)
- Moth: MothHelpers.CapitalizeWString -> wstring.sub(wstr, 1, 1)

## Related APIs

- none

## Used With

- [wstring.lower](global_wstring.lower.md) (MEDIUM 45/100) - Global Function
- [wstring.upper](global_wstring.upper.md) (MEDIUM 45/100) - Global Function

## Triggered By

- none

## Affects

- none

## Notes

- Only one addon surfaced this symbol in the current corpus.
- Canonical entry built from observed call sites, not from engine source or decompiled definitions.
