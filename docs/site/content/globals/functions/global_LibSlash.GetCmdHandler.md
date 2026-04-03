# LibSlash.GetCmdHandler

- Category: Global Function
- Confidence level: HIGH
- Confidence score: 73/100
- Seen in: 1 addons

## Confidence Assessment

- Level: HIGH

- Score: 73/100

- Rationale: Promoted as HIGH confidence because matches a known engine namespace, called globally with no local definition, argument pattern is consistent.

## Evidence Signals

- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +10 Argument pattern is consistent: Observed argument positions remain stable.
- +8 Return usage is consistent: Observed as a stable query-style API.
- +10 Appears in slash command registration patterns: Observed in shared command registration flows.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | FastFriends |
| Files seen in | FastFriends.lua |
| Namespaces detected | LibSlash |
| Source kinds | lua_calls |
| Example locations | FastFriends: ff_isSlashRegistered |
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
| Consistent returns | yes |
| Slash command presence | yes |
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
LibSlash.GetCmdHandler(arg1)
```

## Description

Observed as a global function across 1 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a runtime window or control identifier. | Observed values: alias |

## Returns

- Observed as a query-style API. The concrete return shape is not inferable from contract artifacts alone.

## Side Effects

- No side effect is confidently inferable from contract artifacts alone.

## Seen In

- FastFriends

## Examples

- FastFriends: ff_isSlashRegistered -> LibSlash.GetCmdHandler(alias)

## Notes

- Only one addon surfaced this symbol in the current corpus.
- Canonical entry built from observed call sites, not from engine source or decompiled definitions.
