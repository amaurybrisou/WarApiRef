# ScenarioGroupWindow.LeaveGroup

- Category: Global Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 1 addons

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 105

- Rationale: Promoted as HIGH confidence because used directly in xml handler attributes, referenced by generated docs or reference files, called globally with no local definition.

## Evidence Signals

- +30 Used directly in XML handler attributes: XML exposure suggests an engine-level contract.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +20 Observed in both XML and Lua paths: Cross-source linkage reinforces platform-level usage.
- +10 Argument pattern is consistent: Observed argument positions remain stable.
- +25 Referenced by generated docs or reference files: The symbol is reinforced outside a single call page.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | Effigy |
| Files seen in | Effigy.lua |
| Namespaces detected | ScenarioGroupWindow |
| Source kinds | bindings, lua_calls |
| Example locations | EA_ScenarioGroupWindow: .OnLButtonUp, Effigy: OnMenuClickLeaveScenarioGroup |
| XML usage count | 1 |
| XML attribute usage count | 1 |
| Lua usage count | 2 |
| Global usage count | 2 |
| Local definition count | 0 |
| Documentation references | 1 |
| Initialization flow references | 0 |
| Known engine namespace | no |
| Default UI presence | no |
| Event binding presence | no |
| Observed in XML and Lua | yes |
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
ScenarioGroupWindow.LeaveGroup()
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

- Effigy

## Examples

- Effigy: OnMenuClickLeaveScenarioGroup -> ScenarioGroupWindow.LeaveGroup()

## Related APIs

- [OnLButtonUp](../../xml/handlers/handler_OnLButtonUp.md) (HIGH 100/100) - XML Event

## Triggered By

- [OnLButtonUp](../../xml/handlers/handler_OnLButtonUp.md) (HIGH 100/100) - XML Event

## Notes

- Only one addon surfaced this symbol in the current corpus.
- Canonical entry built from observed call sites, not from engine source or decompiled definitions.
- Advanced return analysis: No strong return evidence observed
