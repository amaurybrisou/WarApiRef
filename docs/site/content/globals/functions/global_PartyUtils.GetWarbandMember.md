# PartyUtils.GetWarbandMember

- Category: Global Function
- Confidence level: HIGH
- Confidence score: 96/100
- Seen in: 3 addons

## Confidence Assessment

- Level: HIGH

- Score: 96/100

- Rationale: Promoted as HIGH confidence because matches a known engine namespace, called globally with no local definition, seen in 2 to 3 addons.

## Evidence Signals

- +18 Seen in 2 to 3 addons: Cross-addon spread is present but limited.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +10 Argument pattern is consistent: Observed argument positions remain stable.
- +8 Return usage is consistent: Observed as a stable query-style API.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | Enemy, LibGroup, Squared |
| Files seen in | Code/Core/Groups/Groups.lua, LibGroup.lua, SquaredWarband.lua |
| Namespaces detected | PartyUtils |
| Source kinds | lua_calls |
| Example locations | Enemy: Groups_OnBattlegroupMemberUpdated, LibGroup: OnWarbandMemberUpdated, Squared: UpdateMember |
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
| Consistent role | yes |
| Consistent arguments | yes |
| Consistent returns | yes |
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
PartyUtils.GetWarbandMember(arg1, arg2)
```

## Description

Observed as a global function across 3 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a runtime window or control identifier. | Observed values: groupInWarband, groupIndex, partyIndex |
| arg2 | Observed as a runtime window or control identifier. | Observed values: memberIndex, slotInGroup |

## Returns

- Observed as a query-style API. The concrete return shape is not inferable from contract artifacts alone.

## Side Effects

- No side effect is confidently inferable from contract artifacts alone.

## Seen In

- Enemy
- LibGroup
- Squared

## Examples

- Enemy: Groups_OnBattlegroupMemberUpdated -> PartyUtils.GetWarbandMember(groupIndex, memberIndex)
- LibGroup: OnWarbandMemberUpdated -> PartyUtils.GetWarbandMember(groupInWarband, slotInGroup)
- Squared: UpdateMember -> PartyUtils.GetWarbandMember(partyIndex, memberIndex)

## Notes

- Canonical entry built from observed call sites, not from engine source or decompiled definitions.
