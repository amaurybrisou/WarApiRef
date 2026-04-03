# EA_Window_InteractionRenownTraining.GetPointsAvailable

- Category: Global Function
- Confidence level: HIGH
- Confidence score: 98/100
- Seen in: 1 addons

## Confidence Assessment

- Level: HIGH

- Score: 98/100

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, matches a known engine namespace, called globally with no local definition.

## Evidence Signals

- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +10 Argument pattern is consistent: Observed argument positions remain stable.
- +8 Return usage is consistent: Observed as a stable query-style API.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | AdvancedRenownTrainer |
| Files seen in | AdvancedRenownTraining.lua |
| Namespaces detected | EA_Window_InteractionRenownTraining |
| Source kinds | lua_calls |
| Example locations | AdvancedRenownTrainer: GetPointsAvailable, AdvancedRenownTrainer: Initialize, AdvancedRenownTrainer: OnShown, AdvancedRenownTrainer: UpdatePointsLabel |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 5 |
| Global usage count | 5 |
| Local definition count | 0 |
| Documentation references | 0 |
| Initialization flow references | 0 |
| Known engine namespace | yes |
| Default UI presence | yes |
| Event binding presence | no |
| Observed in XML and Lua | no |
| Consistent role | no |
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
EA_Window_InteractionRenownTraining.GetPointsAvailable()
```

## Description

Observed as a global function across 1 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |

## Returns

- Observed as a query-style API. The concrete return shape is not inferable from contract artifacts alone.

## Side Effects

- No side effect is confidently inferable from contract artifacts alone.

## Seen In

- AdvancedRenownTrainer

## Examples

- AdvancedRenownTrainer: GetPointsAvailable -> EA_Window_InteractionRenownTraining.GetPointsAvailable()
- AdvancedRenownTrainer: Initialize -> EA_Window_InteractionRenownTraining.GetPointsAvailable()
- AdvancedRenownTrainer: OnShown -> EA_Window_InteractionRenownTraining.GetPointsAvailable()
- AdvancedRenownTrainer: UpdatePointsLabel -> EA_Window_InteractionRenownTraining.GetPointsAvailable()

## Related APIs

- [OnShown](../../xml/handlers/handler_OnShown.md) (HIGH 88/100) - XML Event

## Used With

- [LabelSetText](../../window_api/functions/window_LabelSetText.md) (HIGH 100/100) - Window Function
- [EA_Window_InteractionRenownTraining.GetPointsSpent](global_EA_Window_InteractionRenownTraining.GetPointsSpent.md) (HIGH 98/100) - Global Function

## Affects

- [SystemData.Events.LOADING_END](../../systemdata/fields/systemdata_SystemData.Events.LOADING_END.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_CAREER_CATEGORY_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_CAREER_CATEGORY_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.RELOAD_INTERFACE](../../systemdata/fields/systemdata_SystemData.Events.RELOAD_INTERFACE.md) (HIGH 100/100) - SystemData Field

## Notes

- Only one addon surfaced this symbol in the current corpus.
- Canonical entry built from observed call sites, not from engine source or decompiled definitions.
- Advanced return analysis: No strong return evidence observed
