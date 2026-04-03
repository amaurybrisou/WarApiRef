# ButtonGetText

- Category: Window Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 5 addons

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 135

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, seen in 4 or more addons, matches a known engine namespace.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +10 Argument pattern is consistent: Observed argument positions remain stable.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | AdvancedRenownTrainer, EA_UiDebugTools, MoraleSet, TastyButtons, WarBoard_WarWhisperer |
| Files seen in | AdvancedRenownTraining.lua, Source/MoraleSet.lua, Source/devpad/DebugWindowCodePad.lua, TastyButtonsOptions.lua, warwhisperer.lua |
| Namespaces detected | ButtonGetText |
| Source kinds | lua_calls |
| Example locations | AdvancedRenownTrainer: Respecialize, EA_UiDebugTools: OnFileButtonUp, EA_UiDebugTools: TestPrint, MoraleSet: OnRButtonUpSetMenu, TastyButtons: ButtonDelProfile, TastyButtons: ButtonLoadProfile |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 18 |
| Global usage count | 18 |
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

Observed as a window function across 5 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a function or method reference. | Observed values: "EA_TacticsEditorSetMenuSelectedButton", (comboBoxButtonName["ComboProfile"]..ComboBoxGetSelectedMenuItem(comboBoxButtonName["ComboProfileBox"])), SystemData.ActiveWindow.name |

## Returns

- Not confidently inferable from contract artifacts alone.

## Side Effects

- No side effect is confidently inferable from contract artifacts alone.

## Seen In

- AdvancedRenownTrainer
- EA_UiDebugTools
- MoraleSet
- TastyButtons
- WarBoard_WarWhisperer

## Examples

- AdvancedRenownTrainer: Respecialize -> ButtonGetText(WindowName.."RespecializeButton")
- EA_UiDebugTools: OnFileButtonUp -> ButtonGetText(windowName)
- EA_UiDebugTools: TestPrint -> ButtonGetText(windowName)
- MoraleSet: OnRButtonUpSetMenu -> ButtonGetText("EA_TacticsEditorSetMenuSelectedButton")
- TastyButtons: ButtonDelProfile -> ButtonGetText((comboBoxButtonName["ComboProfile"]..ComboBoxGetSelectedMenuItem(comboBoxButtonName["ComboProfileBox"])))
- TastyButtons: ButtonLoadProfile -> ButtonGetText((comboBoxButtonName["ComboProfile"]..ComboBoxGetSelectedMenuItem(comboBoxButtonName["ComboProfileBox"])))

## Used With

- [DialogManager.MakeOneButtonDialog](../../globals/functions/global_DialogManager.MakeOneButtonDialog.md) (HIGH 100/100) - Global Function
- [DialogManager.MakeTwoButtonDialog](../../globals/functions/global_DialogManager.MakeTwoButtonDialog.md) (HIGH 100/100) - Global Function
- [EA_Window_InteractionRenownTraining.GetPointsAlreadyPurchased](../../globals/functions/global_EA_Window_InteractionRenownTraining.GetPointsAlreadyPurchased.md) (HIGH 98/100) - Global Function
- [GameData.Player.GetRenownRefundCost](../../globals/functions/global_GameData.Player.GetRenownRefundCost.md) (HIGH 80/100) - Global Function

## Affects

- [GameData.Player.GetRenownRefundCost](../../globals/functions/global_GameData.Player.GetRenownRefundCost.md) (HIGH 80/100) - Global Function

## Notes

- none
