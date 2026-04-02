# PartyUtils.GetPartyMember

- Category: Global Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 2 addons

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 121

- Rationale: Promoted as HIGH confidence because matches a known engine namespace, referenced by generated docs or reference files, called globally with no local definition.

## Evidence Signals

- +18 Seen in 2 to 3 addons: Cross-addon spread is present but limited.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +10 Argument pattern is consistent: Observed argument positions remain stable.
- +8 Return usage is consistent: Observed as a stable query-style API.
- +25 Referenced by generated docs or reference files: The symbol is reinforced outside a single call page.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | Enemy, LibGroup |
| Files seen in | `/workspace_addons/Enemy/Code/Core/Groups/Groups.lua:558`, `/workspace_addons/LibGroup/LibGroup.lua:802` |
| Namespaces detected | PartyUtils |
| Source kinds | globals, lua_calls |
| Example locations | Enemy: Enemy.Groups_OnGroupStatusUpdated, LibGroup: LibGroup.OnGroupStatusUpdated |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 2 |
| Global usage count | 2 |
| Local definition count | 0 |
| Documentation references | 1 |
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
PartyUtils.GetPartyMember(arg1)
```

## Description

Observed as a global function across 2 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a runtime window or control identifier. | Observed values: indexInGroup, memberIndex |

## Returns

- Observed as a query-style API. The concrete return shape is not inferable from addon-api docs alone.

## Side Effects

- No side effect is confidently inferable from addon-api docs alone.

## Seen In

- Enemy
- LibGroup

## Examples

- Enemy: Enemy.Groups_OnGroupStatusUpdated -> PartyUtils.GetPartyMember(memberIndex)
- LibGroup: LibGroup.OnGroupStatusUpdated -> PartyUtils.GetPartyMember(indexInGroup)

## Related APIs

- none

## Used With

- none

## Triggered By

- [SystemData.Events.GROUP_STATUS_UPDATED](../../events/game_events/game_event_SystemData.Events.GROUP_STATUS_UPDATED.md) (HIGH 100/100) - Game Event

## Affects

- [PartyUtils](../tables/table_PartyUtils.md) (HIGH 100/100) - Global Table

## Notes

- Canonical entry built from observed call sites, not from engine source or decompiled definitions.
