# ComboBoxClearMenuItems

- Category: Window Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 17 addons

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
| Addons seen in | Ace, AdvancedPetAssist, AdvancedRenownTrainer, BankArkel, BuffHead, DAoCBuff, Enemy, Killer |
| Files seen in | `/workspace/data/raw/Ace/LibGUI.lua:1078`, `/workspace/data/raw/AdvancedPetAssist/APAGui.lua:735`, `/workspace/data/raw/BankArkel/BankArkel.lua:513`, `/workspace/data/raw/BuffHead/Setup/SetupAdvancedCompressionItemEffect.lua:70`, `/workspace/data/raw/BuffHead/Setup/SetupAdvancedContainersItem.lua:342`, `/workspace/data/raw/BuffHead/Setup/SetupAdvancedContainersItemProperties.lua:205`, `/workspace/data/raw/BuffHead/Setup/SetupAdvancedContainersItemProperties.lua:228`, `/workspace/data/raw/BuffHead/Setup/SetupAdvancedContainersItemProperties.lua:283` |
| Namespaces detected | ComboBoxClearMenuItems |
| Source kinds | lua_calls |
| Example locations | Ace: LIBGUI_Combobox:Clear, AdvancedPetAssist: AdvancedPetAssist.local.FillCombo, AdvancedPetAssist: FillCombo, AdvancedRenownTrainer: AdvancedRenownTrainer.local.SaveCurrentSpecAsPreset, AdvancedRenownTrainer: AdvancedRenownTraining.DeletePreset, AdvancedRenownTrainer: AdvancedRenownTraining.SavePreset |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 121 |
| Global usage count | 121 |
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
ComboBoxClearMenuItems(arg1)
```

## Description

Observed as a window function across 17 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a function or method reference. | Observed values: "AdvancedRenownTrainingPresetsWindowLoadComboBox", "BankArkelBackpackCombo", "EnemyChooseChannelDialogChannelList" |

## Returns

- Not confidently inferable from addon-api docs alone.

## Side Effects

- No side effect is confidently inferable from addon-api docs alone.

## Seen In

- Ace
- AdvancedPetAssist
- AdvancedRenownTrainer
- BankArkel
- BuffHead
- DAoCBuff
- Enemy
- Killer
- LibWBToggler
- PartyCast
- PotionBar
- Shinies
- TexturedButtons
- TidyRoll
- TurretRange
- WhoHealedMe
- WoH-Reticle

## Examples

- Ace: LIBGUI_Combobox:Clear -> ComboBoxClearMenuItems(self.name)
- AdvancedPetAssist: AdvancedPetAssist.local.FillCombo -> ComboBoxClearMenuItems(comboName)
- AdvancedPetAssist: FillCombo -> ComboBoxClearMenuItems(comboName)
- AdvancedRenownTrainer: AdvancedRenownTrainer.local.SaveCurrentSpecAsPreset -> ComboBoxClearMenuItems(PresetWindowName.."LoadComboBox")
- AdvancedRenownTrainer: AdvancedRenownTraining.DeletePreset -> ComboBoxClearMenuItems(PresetWindowName.."LoadComboBox")
- AdvancedRenownTrainer: AdvancedRenownTraining.SavePreset -> ComboBoxClearMenuItems(PresetWindowName.."LoadComboBox")

## Related APIs

- none

## Used With

- [ButtonSetText](window_ButtonSetText.md) (HIGH 100/100) - Window Function
- [ComboBoxAddMenuItem](window_ComboBoxAddMenuItem.md) (HIGH 100/100) - Window Function
- [ComboBoxSetSelectedMenuItem](window_ComboBoxSetSelectedMenuItem.md) (HIGH 100/100) - Window Function

## Triggered By

- [OnLButtonUp](../../xml/handlers/handler_OnLButtonUp.md) (HIGH 100/100) - XML Event
- [OnLButtonUp](../../events/window_events/window_event_OnLButtonUp.md) (HIGH 100/100) - Window Event

## Affects

- [Button](../../xml/element_types/element_Button.md) (HIGH 100/100) - XML Element Type
- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Notes

- none
