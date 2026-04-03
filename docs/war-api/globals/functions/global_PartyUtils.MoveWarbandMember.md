# PartyUtils.MoveWarbandMember

- Category: Global Function
- Confidence level: HIGH
- Confidence score: 88/100
- Seen in: 2 addons

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
| Addons seen in | EA_OpenPartyWindow, Hopper |
| Files seen in | Source/Hopper.lua, source/openpartywindowtabmanage.lua |
| Namespaces detected | PartyUtils |
| Source kinds | lua_calls |
| Example locations | EA_OpenPartyWindow: OnLButtonUpProcessed, Hopper: SetGroup |
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
PartyUtils.MoveWarbandMember(arg1, arg2)
```

## Description

Observed as a global function across 2 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a runtime window or control identifier. | Observed values: member.name, playerName |
| arg2 | Observed as a runtime window or control identifier. | Observed values: groupIndex, newPartyIndex |

## Returns

- Not confidently inferable from contract artifacts alone.

## Side Effects

- No side effect is confidently inferable from contract artifacts alone.

## Seen In

- EA_OpenPartyWindow
- Hopper

## Examples

- EA_OpenPartyWindow: OnLButtonUpProcessed -> PartyUtils.MoveWarbandMember(member.name, newPartyIndex)
- Hopper: SetGroup -> PartyUtils.MoveWarbandMember(playerName, groupIndex)

## Used With

- [PartyUtils.GetWarbandMember](global_PartyUtils.GetWarbandMember.md) (HIGH 100/100) - Global Function
- [WindowClearAnchors](../../window_api/functions/window_WindowClearAnchors.md) (HIGH 100/100) - Window Function
- [WindowGetId](../../window_api/functions/window_WindowGetId.md) (HIGH 100/100) - Window Function
- [PartyUtils.SwapWarbandMembers](global_PartyUtils.SwapWarbandMembers.md) (MEDIUM 45/100) - Global Function

## Affects

- [SystemData.MouseOverWindow.name](../../systemdata/fields/systemdata_SystemData.MouseOverWindow.name.md) (HIGH 100/100) - SystemData Field

## Notes

- Canonical entry built from observed call sites, not from engine source or decompiled definitions.
