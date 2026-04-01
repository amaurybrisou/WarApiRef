# EA_Window_InteractionRenownTraining

- Category: Global Table
- Confidence level: HIGH
- Confidence score: 100/100

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 118

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, matches a known engine namespace, referenced by generated docs or reference files.

## Evidence Signals

- +18 Seen in 2 to 3 addons: Cross-addon spread is present but limited.
- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +25 Referenced by generated docs or reference files: The symbol is reinforced outside a single call page.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | AdvancedRenownTrainer, AnywhereTrainer |
| Files seen in | `/workspace/AnywhereTrainer/source/AnywhereTrainer.lua:243` |
| Namespaces detected | EA_Window_InteractionRenownTraining |
| Source kinds | globals, lua_calls |
| Example locations | AnywhereTrainer: AnywhereTrainer.OnLeftClickRenown |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 2 |
| Global usage count | 2 |
| Local definition count | 1 |
| Documentation references | 1 |
| Initialization flow references | 0 |
| Known engine namespace | yes |
| Default UI presence | yes |
| Event binding presence | no |
| Observed in XML and Lua | no |
| Consistent role | yes |
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

## Description

Observed shared global table or namespace surfaced in 2 addons.

## Functions

- EA_Window_InteractionRenownTraining.Hide
- EA_Window_InteractionRenownTraining.Show

## Observed Members

- none

## Seen In

- AdvancedRenownTrainer
- AnywhereTrainer

## Examples

- AnywhereTrainer: AnywhereTrainer.OnLeftClickRenown -> EA_Window_InteractionRenownTraining.Show()
- AnywhereTrainer: AnywhereTrainer.OnLeftClickRenown -> EA_Window_InteractionRenownTraining.Hide()

## Related APIs

- [ButtonGetText](../../window_api/functions/window_ButtonGetText.md) (HIGH 100/100) - Window Function
- [DialogManager.MakeOneButtonDialog](../functions/global_DialogManager.MakeOneButtonDialog.md) (HIGH 100/100) - Global Function
- [DialogManager.MakeTwoButtonDialog](../functions/global_DialogManager.MakeTwoButtonDialog.md) (HIGH 100/100) - Global Function
- [GameData.Player.GetRenownRefundCost](../functions/global_GameData.Player.GetRenownRefundCost.md) (HIGH 100/100) - Global Function

## Used With

- [Button](../../xml/element_types/element_Button.md) (HIGH 100/100) - XML Element Type
- [ButtonGetText](../../window_api/functions/window_ButtonGetText.md) (HIGH 100/100) - Window Function
- [DialogManager.MakeOneButtonDialog](../functions/global_DialogManager.MakeOneButtonDialog.md) (HIGH 100/100) - Global Function
- [DialogManager.MakeTwoButtonDialog](../functions/global_DialogManager.MakeTwoButtonDialog.md) (HIGH 100/100) - Global Function
- [GameData.Player.GetRenownRefundCost](../functions/global_GameData.Player.GetRenownRefundCost.md) (HIGH 100/100) - Global Function
- [OnLButtonUp](../../events/window_events/window_event_OnLButtonUp.md) (HIGH 100/100) - Window Event
- [OnLButtonUp](../../xml/handlers/handler_OnLButtonUp.md) (HIGH 100/100) - XML Handler

## Triggered By

- none

## Affects

- none

## Notes

- none
