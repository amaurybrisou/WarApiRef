# ButtonSetPressedFlag

- Category: Window Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 21 addons

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
| Addons seen in | Ace, AdvancedRenownTrainer, AggroMeter, Aura, BuffHead, CM_ClosetGoblin, DAoCBuff, Enemy |
| Files seen in | `/workspace/data/raw/Ace/LibGUI.lua:761`, `/workspace/data/raw/Ace/LibGUI.lua:831`, `/workspace/data/raw/AggroMeter/AggroMeter.lua:426`, `/workspace/data/raw/AggroMeter/AggroMeter.lua:5`, `/workspace/data/raw/Aura/Source/AuraShares.lua:365`, `/workspace/data/raw/Aura/Source/AuraShares.lua:372`, `/workspace/data/raw/Aura/Source/AuraShares.lua:68`, `/workspace/data/raw/BuffHead/Setup/SetupAdvancedContainersItemProperties.lua:283` |
| Namespaces detected | ButtonSetPressedFlag |
| Source kinds | lua_calls |
| Example locations | Ace: LIBGUI_Checkbox:SetValue, Ace: LIBGUI_Optionbutton:SetValue, AdvancedRenownTrainer: AdvancedRenownTrainer.local.SetLabels, AdvancedRenownTrainer: AdvancedRenownTraining.ChangeTab, AdvancedRenownTrainer: AdvancedRenownTraining.OnButtonPressedActiveTab, AdvancedRenownTrainer: AdvancedRenownTraining.OnButtonPressedAdvancedTab |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 302 |
| Global usage count | 302 |
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
ButtonSetPressedFlag(arg1, arg2)
```

## Description

Observed mutating button text or pressed state on existing controls.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a function or method reference. | Observed values: "AggroMeterGrayWindowBlackTab", "AggroMeterGrayWindowListTab", "AggroMeterGrayWindowWhiteTab" |
| arg2 | Observed as a function or method reference. | Observed values: (LibGroup.Settings.Distance.Enabled==true), (activeTracker.Enabled==true), (activeTracker.OnTargetChange.ClearAlwaysShow==true) |

## Returns

- Not confidently inferable from addon-api docs alone.

## Side Effects

- Mutates engine-visible UI state or presentation.

## Seen In

- Ace
- AdvancedRenownTrainer
- AggroMeter
- Aura
- BuffHead
- CM_ClosetGoblin
- DAoCBuff
- Enemy
- Killer
- LibGroup
- LibWBToggler
- PartyCast
- Pocket Palette
- PotionBar
- Shinies
- TexturedButtons
- TidyChat
- TidyRoll
- TurretRange
- WSCT
- WoH-Reticle

## Examples

- Ace: LIBGUI_Checkbox:SetValue -> ButtonSetPressedFlag(self.name, checkState)
- Ace: LIBGUI_Optionbutton:SetValue -> ButtonSetPressedFlag(self.name, checkState)
- AdvancedRenownTrainer: AdvancedRenownTrainer.local.SetLabels -> ButtonSetPressedFlag(PresetWindowName.."SaveSelectedCheckBox", CheckButton)
- AdvancedRenownTrainer: AdvancedRenownTraining.ChangeTab -> ButtonSetPressedFlag(WindowName..AdvancedRenownTraining.ADVANCED, false)
- AdvancedRenownTrainer: AdvancedRenownTraining.ChangeTab -> ButtonSetPressedFlag(WindowName..AdvancedRenownTraining.ACTIVE, false)
- AdvancedRenownTrainer: AdvancedRenownTraining.ChangeTab -> ButtonSetPressedFlag(WindowName..AdvancedRenownTraining.BASIC, false)

## Related APIs

- none

## Used With

- [ButtonGetPressedFlag](window_ButtonGetPressedFlag.md) (HIGH 100/100) - Window Function
- [ButtonSetText](window_ButtonSetText.md) (HIGH 100/100) - Window Function
- [ComboBoxAddMenuItem](window_ComboBoxAddMenuItem.md) (HIGH 100/100) - Window Function
- [ComboBoxSetSelectedMenuItem](window_ComboBoxSetSelectedMenuItem.md) (HIGH 100/100) - Window Function
- [LabelSetText](window_LabelSetText.md) (HIGH 100/100) - Window Function
- [OnLButtonUp](../../xml/handlers/handler_OnLButtonUp.md) (HIGH 100/100) - XML Event
- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type
- [WindowSetShowing](window_WindowSetShowing.md) (HIGH 100/100) - Window Function

## Triggered By

- none

## Affects

- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Notes

- none
