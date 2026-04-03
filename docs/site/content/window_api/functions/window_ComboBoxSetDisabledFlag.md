# ComboBoxSetDisabledFlag

- Category: Window Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 38 addons

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
| Addons seen in | Ace, ActionBarHide, AdvancedRenownTrainer, Amethyst, Crusher, DAoCBuff, DammazKron, EZCraftX |
| Files seen in | AdvancedRenownTrainingImportExport.lua, Conf/DK_Config.lua, Gui/HealGridGuiUtility.lua, LibGUI.lua, LibGui.lua, Libraries/LibGUI.lua, Libs/LibGUI.lua, Settings.lua |
| Namespaces detected | ComboBoxSetDisabledFlag |
| Source kinds | lua_calls |
| Example locations | Ace: SetEnabled, ActionBarHide: SetEnabled, AdvancedRenownTrainer: OnExportHidden, AdvancedRenownTrainer: OnExportShown, Amethyst: SetEnabled, Crusher: SetEnabled |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 63 |
| Global usage count | 63 |
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

Observed as a window function across 38 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a function or method reference. | Observed values: "AdvancedRenownTrainingPresetsWindowLoadComboBox", "ZonePOPWndComboBoxZoneNames", FilterWindow.."ClassTableComboBox" |
| arg2 | Observed as a runtime window or control identifier. | Observed values: active, disable, false |

## Returns

- Not confidently inferable from contract artifacts alone.

## Side Effects

- Mutates engine-visible UI state or presentation.

## Seen In

- Ace
- ActionBarHide
- AdvancedRenownTrainer
- Amethyst
- Crusher
- DAoCBuff
- DammazKron
- EZCraftX
- EZGuard
- Effigy
- GCDsaver
- HealGrid
- Hopper
- InfoScroller
- LibWBToggler
- Map
- Motion
- NaturalLog
- PartyCast
- PotionBar
- Pure
- Pure Careerbar
- RealmStatus
- Shinies
- TargetRing
- TastyButtons
- TidyChat
- Tokens
- Vectors
- WarTriage
- Wikki's Cooldown Bar
- Wikki's Cooldown Pulse
- WoH-Reticle
- ZonePOP
- nLootLink
- scenarioInfo
- xHUD
- xPanels

## Examples

- Ace: SetEnabled -> ComboBoxSetDisabledFlag(self.name, not flag)
- ActionBarHide: SetEnabled -> ComboBoxSetDisabledFlag(self.name, not flag)
- AdvancedRenownTrainer: OnExportHidden -> ComboBoxSetDisabledFlag("AdvancedRenownTrainingPresetsWindowLoadComboBox", false)
- AdvancedRenownTrainer: OnExportShown -> ComboBoxSetDisabledFlag("AdvancedRenownTrainingPresetsWindowLoadComboBox", true)
- Amethyst: SetEnabled -> ComboBoxSetDisabledFlag(self.name, not flag)
- Crusher: SetEnabled -> ComboBoxSetDisabledFlag(self.name, not flag)

## Used With

- [ButtonSetDisabledFlag](window_ButtonSetDisabledFlag.md) (HIGH 100/100) - Window Function

## Notes

- none
