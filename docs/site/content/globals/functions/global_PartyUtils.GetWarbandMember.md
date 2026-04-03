# PartyUtils.GetWarbandMember

- Category: Global Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 5 addons

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 108

- Rationale: Promoted as HIGH confidence because seen in 4 or more addons, matches a known engine namespace, called globally with no local definition.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +10 Argument pattern is consistent: Observed argument positions remain stable.
- +8 Return usage is consistent: Observed as a stable query-style API.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | EA_OpenPartyWindow, Enemy, LibGroup, NoUselessMods-Assist, Squared |
| Files seen in | Code/Core/Groups/Groups.lua, LibGroup.lua, SquaredWarband.lua, no-useless-mods-warband.lua, source/openpartywindowtabmanage.lua |
| Namespaces detected | PartyUtils |
| Source kinds | lua_calls |
| Example locations | EA_OpenPartyWindow: OnLButtonUpProcessed, EA_OpenPartyWindow: OnRButtonUpPlayerRow, EA_OpenPartyWindow: SingleMemberUpdate, Enemy: Groups_OnBattlegroupMemberUpdated, LibGroup: OnWarbandMemberUpdated, NoUselessMods-Assist: GetWarbandMember |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 9 |
| Global usage count | 9 |
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

Observed as a global function across 5 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a runtime window or control identifier. | Observed values: groupInWarband, groupIndex, newPartyIndex |
| arg2 | Observed as a runtime window or control identifier. | Observed values: memberIndex, newMemberIndex, slotInGroup |

## Returns

- Observed as a query-style API. The concrete return shape is not inferable from contract artifacts alone.

## Side Effects

- No side effect is confidently inferable from contract artifacts alone.

## Seen In

- EA_OpenPartyWindow
- Enemy
- LibGroup
- NoUselessMods-Assist
- Squared

## Examples

- EA_OpenPartyWindow: OnLButtonUpProcessed -> PartyUtils.GetWarbandMember(partyIndex, memberIndex)
- EA_OpenPartyWindow: OnLButtonUpProcessed -> PartyUtils.GetWarbandMember(newPartyIndex, newMemberIndex)
- EA_OpenPartyWindow: OnRButtonUpPlayerRow -> PartyUtils.GetWarbandMember(partyIndex, memberIndex)
- EA_OpenPartyWindow: SingleMemberUpdate -> PartyUtils.GetWarbandMember(partyIndex, memberIndex)
- Enemy: Groups_OnBattlegroupMemberUpdated -> PartyUtils.GetWarbandMember(groupIndex, memberIndex)
- LibGroup: OnWarbandMemberUpdated -> PartyUtils.GetWarbandMember(groupInWarband, slotInGroup)

## Used With

- [WindowClearAnchors](../../window_api/functions/window_WindowClearAnchors.md) (HIGH 100/100) - Window Function
- [WindowGetId](../../window_api/functions/window_WindowGetId.md) (HIGH 100/100) - Window Function
- [PartyUtils.MoveWarbandMember](global_PartyUtils.MoveWarbandMember.md) (HIGH 88/100) - Global Function
- [PartyUtils.SwapWarbandMembers](global_PartyUtils.SwapWarbandMembers.md) (MEDIUM 45/100) - Global Function

## Affects

- [SystemData.MouseOverWindow.name](../../systemdata/fields/systemdata_SystemData.MouseOverWindow.name.md) (HIGH 100/100) - SystemData Field

## Notes

- Canonical entry built from observed call sites, not from engine source or decompiled definitions.
- Advanced return analysis: No strong return evidence observed
