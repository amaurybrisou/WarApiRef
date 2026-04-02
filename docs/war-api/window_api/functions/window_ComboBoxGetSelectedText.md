# ComboBoxGetSelectedText

- Category: Window Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 6 addons

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
| Addons seen in | AdvancedRenownTrainer, BankArkel, Busted, EA_UiDebugTools, PotionBar, TidyRoll |
| Files seen in | `/workspace_addons/BankArkel/BankArkel.lua:482`, `/workspace_addons/Busted/Busted.lua:248`, `/workspace_addons/PotionBar/settings/Settings.lua:151`, `/workspace_addons/TidyRoll/TidyRollOptions.lua:782`, `/workspace_addons/advancedrenowntrainer/AdvancedRenownTraining.lua:920`, `/workspace_addons/advancedrenowntrainer/AdvancedRenownTraining.lua:964`, `/workspace_addons/advancedrenowntrainer/AdvancedRenownTraining.lua:978`, `/workspace_addons/advancedrenowntrainer/AdvancedRenownTrainingImportExport.lua:192` |
| Namespaces detected | ComboBoxGetSelectedText |
| Source kinds | lua_calls |
| Example locations | AdvancedRenownTrainer: AdvancedRenownTraining.DeletePreset, AdvancedRenownTrainer: AdvancedRenownTraining.ExportToLink, AdvancedRenownTrainer: AdvancedRenownTraining.OnExportButtonPressed, AdvancedRenownTrainer: AdvancedRenownTraining.OnLoadPressed, AdvancedRenownTrainer: AdvancedRenownTraining.SelectedItemChanged, BankArkel: BankArkel.PackCombo |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 14 |
| Global usage count | 14 |
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
ComboBoxGetSelectedText(arg1)
```

## Description

Observed as a window function across 6 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a function or method reference. | Observed values: "BankArkelBackpackCombo", "BustedGUIAddonSelect", PresetWindowName.."LoadComboBox" |

## Returns

- Not confidently inferable from addon-api docs alone.

## Side Effects

- No side effect is confidently inferable from addon-api docs alone.

## Seen In

- AdvancedRenownTrainer
- BankArkel
- Busted
- EA_UiDebugTools
- PotionBar
- TidyRoll

## Examples

- AdvancedRenownTrainer: AdvancedRenownTraining.DeletePreset -> ComboBoxGetSelectedText(PresetWindowName.."LoadComboBox")
- AdvancedRenownTrainer: AdvancedRenownTraining.ExportToLink -> ComboBoxGetSelectedText(PresetWindowName.."LoadComboBox")
- AdvancedRenownTrainer: AdvancedRenownTraining.OnExportButtonPressed -> ComboBoxGetSelectedText(PresetWindowName.."LoadComboBox")
- AdvancedRenownTrainer: AdvancedRenownTraining.OnLoadPressed -> ComboBoxGetSelectedText(PresetWindowName.."LoadComboBox")
- AdvancedRenownTrainer: AdvancedRenownTraining.SelectedItemChanged -> ComboBoxGetSelectedText(PresetWindowName.."LoadComboBox")
- BankArkel: BankArkel.PackCombo -> ComboBoxGetSelectedText("BankArkelBackpackCombo")

## Related APIs

- none

## Used With

- [Button](../../xml/element_types/element_Button.md) (HIGH 100/100) - XML Element Type
- [ComboBox](../../xml/element_types/element_ComboBox.md) (HIGH 100/100) - XML Element Type
- [ComboBoxClearMenuItems](window_ComboBoxClearMenuItems.md) (HIGH 100/100) - Window Function
- [ComboBoxGetSelectedMenuItem](window_ComboBoxGetSelectedMenuItem.md) (HIGH 100/100) - Window Function
- [EA_ChatWindow.InsertText](../../globals/functions/global_EA_ChatWindow.InsertText.md) (HIGH 100/100) - Global Function
- [OnLButtonUp](../../events/window_events/window_event_OnLButtonUp.md) (HIGH 100/100) - Window Event
- [OnLButtonUp](../../xml/handlers/handler_OnLButtonUp.md) (HIGH 100/100) - XML Handler
- [OnSelChanged](../../events/window_events/window_event_OnSelChanged.md) (HIGH 100/100) - Window Event
- [OnSelChanged](../../xml/handlers/handler_OnSelChanged.md) (HIGH 100/100) - XML Handler
- [WindowSetShowing](window_WindowSetShowing.md) (HIGH 100/100) - Window Function

## Triggered By

- [OnLButtonUp](../../xml/handlers/handler_OnLButtonUp.md) (HIGH 100/100) - XML Handler
- [OnLButtonUp](../../events/window_events/window_event_OnLButtonUp.md) (HIGH 100/100) - Window Event
- [OnSelChanged](../../xml/handlers/handler_OnSelChanged.md) (HIGH 100/100) - XML Handler
- [OnSelChanged](../../events/window_events/window_event_OnSelChanged.md) (HIGH 100/100) - Window Event

## Affects

- [Button](../../xml/element_types/element_Button.md) (HIGH 100/100) - XML Element Type
- [ComboBox](../../xml/element_types/element_ComboBox.md) (HIGH 100/100) - XML Element Type
- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Notes

- none
