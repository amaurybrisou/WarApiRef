# EA_Window_ScenarioLobby.OnLeaveActiveQueueFromLobby

- Category: Global Function
- Confidence level: HIGH
- Confidence score: 90/100
- Seen in: 1 addons

## Confidence Assessment

- Level: HIGH

- Score: 90/100

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, matches a known engine namespace, called globally with no local definition.

## Evidence Signals

- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +10 Argument pattern is consistent: Observed argument positions remain stable.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | Queue Queuer |
| Files seen in | QueueQueuer.lua |
| Namespaces detected | EA_Window_ScenarioLobby |
| Source kinds | lua_calls |
| Example locations | Queue Queuer: LeaveQueues, Queue Queuer: OnCancel |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 2 |
| Global usage count | 2 |
| Local definition count | 0 |
| Documentation references | 0 |
| Initialization flow references | 0 |
| Known engine namespace | yes |
| Default UI presence | yes |
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
EA_Window_ScenarioLobby.OnLeaveActiveQueueFromLobby()
```

## Description

Observed as a global function across 1 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |

## Returns

- Not confidently inferable from contract artifacts alone.

## Side Effects

- No side effect is confidently inferable from contract artifacts alone.

## Seen In

- Queue Queuer

## Examples

- Queue Queuer: LeaveQueues -> EA_Window_ScenarioLobby.OnLeaveActiveQueueFromLobby()
- Queue Queuer: OnCancel -> EA_Window_ScenarioLobby.OnLeaveActiveQueueFromLobby()

## Affects

- [SystemData.Events.INTERACT_LEAVE_SCENARIO_QUEUE](../../systemdata/fields/systemdata_SystemData.Events.INTERACT_LEAVE_SCENARIO_QUEUE.md) (HIGH 100/100) - SystemData Field

## Notes

- Only one addon surfaced this symbol in the current corpus.
- Canonical entry built from observed call sites, not from engine source or decompiled definitions.
- Advanced return analysis: No strong return evidence observed
