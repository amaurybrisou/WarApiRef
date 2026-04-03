# ComboBoxGetSelectedText

- Category: Window Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 9 addons

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
| Addons seen in | AdvancedRenownTrainer, BankArkel, Busted, ChattyCathy, DPSMeter, EA_OpenPartyWindow, EA_UiDebugTools, PotionBar |
| Files seen in | AdvancedRenownTraining.lua, AdvancedRenownTrainingImportExport.lua, BankArkel.lua, Busted.lua, ChattyCathy.lua, DPSMeterWindow.lua, Source/bust/Busted.lua, Source/devpad/DebugWindowCodePad.lua |
| Namespaces detected | ComboBoxGetSelectedText |
| Source kinds | lua_calls |
| Example locations | AdvancedRenownTrainer: DeletePreset, AdvancedRenownTrainer: ExportToLink, AdvancedRenownTrainer: OnExportButtonPressed, AdvancedRenownTrainer: OnLoadPressed, AdvancedRenownTrainer: SelectedItemChanged, BankArkel: PackCombo |
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
ComboBoxGetSelectedText(arg1)
```

## Description

Observed as a window function across 9 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a function or method reference. | Observed values: "BankArkelBackpackCombo", "BustedGUIAddonSelect", "ChattyCathyOptToCombo" |

## Returns

- Not confidently inferable from contract artifacts alone.

## Side Effects

- No side effect is confidently inferable from contract artifacts alone.

## Seen In

- AdvancedRenownTrainer
- BankArkel
- Busted
- ChattyCathy
- DPSMeter
- EA_OpenPartyWindow
- EA_UiDebugTools
- PotionBar
- TidyRoll

## Examples

- AdvancedRenownTrainer: DeletePreset -> ComboBoxGetSelectedText(PresetWindowName.."LoadComboBox")
- AdvancedRenownTrainer: ExportToLink -> ComboBoxGetSelectedText(PresetWindowName.."LoadComboBox")
- AdvancedRenownTrainer: OnExportButtonPressed -> ComboBoxGetSelectedText(PresetWindowName.."LoadComboBox")
- AdvancedRenownTrainer: OnLoadPressed -> ComboBoxGetSelectedText(PresetWindowName.."LoadComboBox")
- AdvancedRenownTrainer: SelectedItemChanged -> ComboBoxGetSelectedText(PresetWindowName.."LoadComboBox")
- BankArkel: PackCombo -> ComboBoxGetSelectedText("BankArkelBackpackCombo")

## Used With

- [WindowSetShowing](window_WindowSetShowing.md) (HIGH 100/100) - Window Function

## Notes

- none
