# ButtonGetText

- Category: Window Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 3 addons

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
| Addons seen in | AdvancedRenownTrainer, EA_UiDebugTools, TastyButtons |
| Files seen in | AdvancedRenownTraining.lua, Source/devpad/DebugWindowCodePad.lua, TastyButtonsOptions.lua |
| Namespaces detected | ButtonGetText |
| Source kinds | lua_calls |
| Example locations | AdvancedRenownTrainer: Respecialize, EA_UiDebugTools: OnFileButtonUp, EA_UiDebugTools: TestPrint, TastyButtons: ButtonDelProfile, TastyButtons: ButtonLoadProfile, TastyButtons: OnComboStateChanged |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 13 |
| Global usage count | 13 |
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

Observed as a window function across 3 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a function or method reference. | Observed values: (comboBoxButtonName["ComboProfile"]..ComboBoxGetSelectedMenuItem(comboBoxButtonName["ComboProfileBox"])), WindowName.."RespecializeButton", comboBoxButtonName["ComboShowing"]..ComboBoxGetSelectedMenuItem(comboBoxButtonName["ComboShowingBox"]) |

## Returns

- Not confidently inferable from contract artifacts alone.

## Side Effects

- No side effect is confidently inferable from contract artifacts alone.

## Seen In

- AdvancedRenownTrainer
- EA_UiDebugTools
- TastyButtons

## Examples

- AdvancedRenownTrainer: Respecialize -> ButtonGetText(WindowName.."RespecializeButton")
- EA_UiDebugTools: OnFileButtonUp -> ButtonGetText(windowName)
- EA_UiDebugTools: TestPrint -> ButtonGetText(windowName)
- TastyButtons: ButtonDelProfile -> ButtonGetText((comboBoxButtonName["ComboProfile"]..ComboBoxGetSelectedMenuItem(comboBoxButtonName["ComboProfileBox"])))
- TastyButtons: ButtonLoadProfile -> ButtonGetText((comboBoxButtonName["ComboProfile"]..ComboBoxGetSelectedMenuItem(comboBoxButtonName["ComboProfileBox"])))
- TastyButtons: OnComboStateChanged -> ButtonGetText(comboBoxButtonName["ComboStateAdd"]..ComboBoxGetSelectedMenuItem(windowNameTable["StateView"].."ComboStateAdd"))

## Used With

- [DialogManager.MakeOneButtonDialog](../../globals/functions/global_DialogManager.MakeOneButtonDialog.md) (HIGH 100/100) - Global Function
- [DialogManager.MakeTwoButtonDialog](../../globals/functions/global_DialogManager.MakeTwoButtonDialog.md) (HIGH 100/100) - Global Function
- [EA_Window_InteractionRenownTraining.GetPointsAlreadyPurchased](../../globals/functions/global_EA_Window_InteractionRenownTraining.GetPointsAlreadyPurchased.md) (HIGH 98/100) - Global Function
- [GameData.Player.GetRenownRefundCost](../../globals/functions/global_GameData.Player.GetRenownRefundCost.md) (HIGH 80/100) - Global Function

## Affects

- [GameData.Player.GetRenownRefundCost](../../globals/functions/global_GameData.Player.GetRenownRefundCost.md) (HIGH 80/100) - Global Function

## Notes

- none
