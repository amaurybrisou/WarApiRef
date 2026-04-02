# ButtonGetText

- Category: Window Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 2 addons

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 123

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, matches a known engine namespace, called globally with no local definition.

## Evidence Signals

- +18 Seen in 2 to 3 addons: Cross-addon spread is present but limited.
- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +10 Argument pattern is consistent: Observed argument positions remain stable.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | AdvancedRenownTrainer, EA_UiDebugTools |
| Files seen in | `/workspace_addons/advancedrenowntrainer/AdvancedRenownTraining.lua:763`, `/workspace_addons/ea_uidebugtools/Source/devpad/DebugWindowCodePad.lua:320`, `/workspace_addons/ea_uidebugtools/Source/devpad/DebugWindowCodePad.lua:785` |
| Namespaces detected | ButtonGetText |
| Source kinds | lua_calls |
| Example locations | AdvancedRenownTrainer: AdvancedRenownTraining.Respecialize, EA_UiDebugTools: DevPad.TestPrint, EA_UiDebugTools: DevPadWindow.OnFileButtonUp |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 3 |
| Global usage count | 3 |
| Local definition count | 0 |
| Documentation references | 0 |
| Initialization flow references | 0 |
| Known engine namespace | yes |
| Default UI presence | yes |
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
ButtonGetText(arg1)
```

## Description

Observed as a window function across 2 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a runtime window or control identifier. | Observed values: WindowName.."RespecializeButton", windowName |

## Returns

- Not confidently inferable from addon-api docs alone.

## Side Effects

- No side effect is confidently inferable from addon-api docs alone.

## Seen In

- AdvancedRenownTrainer
- EA_UiDebugTools

## Examples

- AdvancedRenownTrainer: AdvancedRenownTraining.Respecialize -> ButtonGetText(WindowName.."RespecializeButton")
- EA_UiDebugTools: DevPad.TestPrint -> ButtonGetText(windowName)
- EA_UiDebugTools: DevPadWindow.OnFileButtonUp -> ButtonGetText(windowName)

## Related APIs

- none

## Used With

- [Button](../../xml/element_types/element_Button.md) (HIGH 100/100) - XML Element Type
- [DialogManager.MakeOneButtonDialog](../../globals/functions/global_DialogManager.MakeOneButtonDialog.md) (HIGH 100/100) - Global Function
- [DialogManager.MakeTwoButtonDialog](../../globals/functions/global_DialogManager.MakeTwoButtonDialog.md) (HIGH 100/100) - Global Function
- [EA_Window_InteractionRenownTraining](../../globals/tables/table_EA_Window_InteractionRenownTraining.md) (HIGH 100/100) - Global Table
- [GameData.Player.GetRenownRefundCost](../../globals/functions/global_GameData.Player.GetRenownRefundCost.md) (HIGH 100/100) - Global Function
- [LabelGetText](window_LabelGetText.md) (HIGH 100/100) - Window Function
- [OnLButtonUp](../../events/window_events/window_event_OnLButtonUp.md) (HIGH 100/100) - Window Event
- [OnLButtonUp](../../xml/handlers/handler_OnLButtonUp.md) (HIGH 100/100) - XML Handler
- [SystemData.ActiveWindow.name](../../systemdata/fields/systemdata_SystemData.ActiveWindow.name.md) (HIGH 100/100) - SystemData Field
- [wstring.gsub](../../globals/functions/global_wstring.gsub.md) (HIGH 100/100) - Global Function
- [wstring.match](../../globals/functions/global_wstring.match.md) (HIGH 100/100) - Global Function
- [towstring](../../globals/functions/global_towstring.md) (HIGH 75/100) - Global Function

## Triggered By

- [OnLButtonUp](../../xml/handlers/handler_OnLButtonUp.md) (HIGH 100/100) - XML Handler
- [OnLButtonUp](../../events/window_events/window_event_OnLButtonUp.md) (HIGH 100/100) - Window Event

## Affects

- [Button](../../xml/element_types/element_Button.md) (HIGH 100/100) - XML Element Type
- [EA_Window_InteractionRenownTraining](../../globals/tables/table_EA_Window_InteractionRenownTraining.md) (HIGH 100/100) - Global Table
- [GameData.Player.GetRenownRefundCost](../../globals/functions/global_GameData.Player.GetRenownRefundCost.md) (HIGH 100/100) - Global Function
- [SystemData.ActiveWindow.name](../../systemdata/fields/systemdata_SystemData.ActiveWindow.name.md) (HIGH 100/100) - SystemData Field
- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Notes

- none
