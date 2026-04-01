# ComboBoxSetDisabledFlag

- Category: Window Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 11 addons

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
| Addons seen in | Ace, AdvancedRenownTrainer, DAoCBuff, Effigy, GCDsaver, LibWBToggler, PotionBar, Shinies |
| Files seen in | `/workspace/Ace/LibGUI.lua:1084`, `/workspace/DAoCBuff/Source/DAoCBuffSettings.lua:491`, `/workspace/DAoCBuff/Source/DAoCBuffSettings.lua:804`, `/workspace/DAoCBuff/Source/DAoCBuffSettings.lua:834`, `/workspace/DAoCBuff/Source/DAoCBuffSettings.lua:898`, `/workspace/DAoCBuff/Source/DAoCBuffSettings2ndTier.lua:1056`, `/workspace/DAoCBuff/Source/DAoCBuffSettings2ndTier.lua:1224`, `/workspace/DAoCBuff/Source/DAoCBuffSettings2ndTier.lua:401` |
| Namespaces detected | ComboBoxSetDisabledFlag |
| Source kinds | lua_calls |
| Example locations | Ace: LIBGUI_Combobox:SetEnabled, AdvancedRenownTrainer: AdvancedRenownTraining.OnExportHidden, AdvancedRenownTrainer: AdvancedRenownTraining.OnExportShown, DAoCBuff: DAoCBuffSettings.ActivateStickCombos, DAoCBuff: DAoCBuffSettings.ActivateType, DAoCBuff: DAoCBuffSettings.PopulateTarget |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 30 |
| Global usage count | 30 |
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
ComboBoxSetDisabledFlag(arg1, arg2)
```

## Description

Observed as a window function across 11 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a function or method reference. | Observed values: "AdvancedRenownTrainingPresetsWindowLoadComboBox", FilterWindow.."ClassTableComboBox", FilterWindow.."FilterPropertyComboBox" |
| arg2 | Observed as a runtime window or control identifier. | Observed values: active, disable, false |

## Returns

- Not confidently inferable from addon-api docs alone.

## Side Effects

- Mutates engine-visible UI state or presentation.

## Seen In

- Ace
- AdvancedRenownTrainer
- DAoCBuff
- Effigy
- GCDsaver
- LibWBToggler
- PotionBar
- Shinies
- TidyChat
- WarTriage
- WoH-Reticle

## Examples

- Ace: LIBGUI_Combobox:SetEnabled -> ComboBoxSetDisabledFlag(self.name, not flag)
- AdvancedRenownTrainer: AdvancedRenownTraining.OnExportHidden -> ComboBoxSetDisabledFlag("AdvancedRenownTrainingPresetsWindowLoadComboBox", false)
- AdvancedRenownTrainer: AdvancedRenownTraining.OnExportShown -> ComboBoxSetDisabledFlag("AdvancedRenownTrainingPresetsWindowLoadComboBox", true)
- DAoCBuff: DAoCBuffSettings.ActivateStickCombos -> ComboBoxSetDisabledFlag(FrameTab.."GrowLeftComboBox", active)
- DAoCBuff: DAoCBuffSettings.ActivateStickCombos -> ComboBoxSetDisabledFlag(FrameTab.."GrowUpComboBox", active)
- DAoCBuff: DAoCBuffSettings.ActivateStickCombos -> ComboBoxSetDisabledFlag(FrameTab.."GrowHorizontalComboBox", active)

## Related APIs

- none

## Used With

- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Triggered By

- [OnHidden](../../xml/handlers/handler_OnHidden.md) (HIGH 100/100) - XML Handler
- [OnHidden](../../events/window_events/window_event_OnHidden.md) (HIGH 100/100) - Window Event
- [OnShown](../../xml/handlers/handler_OnShown.md) (HIGH 100/100) - XML Handler
- [OnShown](../../events/window_events/window_event_OnShown.md) (HIGH 100/100) - Window Event

## Affects

- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Notes

- none
