# PartyUtils.SwapWarbandMembers

- Category: Global Function
- Confidence level: MEDIUM
- Confidence score: 45/100
- Seen in: 1 addons

## Confidence Assessment

- Level: MEDIUM

- Score: 45/100

- Rationale: Promoted as MEDIUM confidence because matches a known engine namespace, called globally with no local definition.

## Evidence Signals

- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | EA_OpenPartyWindow |
| Files seen in | source/openpartywindowtabmanage.lua |
| Namespaces detected | PartyUtils |
| Source kinds | lua_calls |
| Example locations | EA_OpenPartyWindow: OnLButtonUpProcessed |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 1 |
| Global usage count | 1 |
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
| Conflicting signatures | no |
| Conflicting roles | no |
| Wrapper likely | no |
| Never outside local graph | no |
| Local helper only | no |

## Signature (inferred)

```lua
PartyUtils.SwapWarbandMembers(arg1, arg2)
```

## Description

Observed as a global function across 1 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a function or method reference. | Observed values: member.name |
| arg2 | Observed as a function or method reference. | Observed values: memberToSwap.name |

## Returns

- Not confidently inferable from contract artifacts alone.

## Side Effects

- No side effect is confidently inferable from contract artifacts alone.

## Seen In

- EA_OpenPartyWindow

## Examples

- EA_OpenPartyWindow: OnLButtonUpProcessed -> PartyUtils.SwapWarbandMembers(member.name, memberToSwap.name)

## Used With

- [PartyUtils.GetWarbandMember](global_PartyUtils.GetWarbandMember.md) (HIGH 100/100) - Global Function
- [WindowClearAnchors](../../window_api/functions/window_WindowClearAnchors.md) (HIGH 100/100) - Window Function
- [WindowGetId](../../window_api/functions/window_WindowGetId.md) (HIGH 100/100) - Window Function
- [PartyUtils.MoveWarbandMember](global_PartyUtils.MoveWarbandMember.md) (HIGH 88/100) - Global Function

## Affects

- [SystemData.MouseOverWindow.name](../../systemdata/fields/systemdata_SystemData.MouseOverWindow.name.md) (HIGH 100/100) - SystemData Field

## Notes

- Only one addon surfaced this symbol in the current corpus.
- Canonical entry built from observed call sites, not from engine source or decompiled definitions.
