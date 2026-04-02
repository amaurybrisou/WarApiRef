# ComboBoxGetSelectedMenuItem

- Category: Window Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 25 addons

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
| Addons seen in | Ace, AdvancedRenownTrainer, BuffHead, Crafting Info Tooltip, DAoCBuff, EA_UiDebugTools, Effigy, Enemy |
| Files seen in | `/workspace_addons/Ace/LibGUI.lua:1110`, `/workspace_addons/BuffHead/Setup/SetupAdvancedContainersItem.lua:455`, `/workspace_addons/BuffHead/Setup/SetupAdvancedContainersItem.lua:470`, `/workspace_addons/BuffHead/Setup/SetupAdvancedContainersItemProperties.lua:581`, `/workspace_addons/BuffHead/Setup/SetupAdvancedContainersItemProperties.lua:594`, `/workspace_addons/BuffHead/Setup/SetupAdvancedContainersItemProperties.lua:619`, `/workspace_addons/BuffHead/Setup/SetupAdvancedContainersItemProperties.lua:631`, `/workspace_addons/BuffHead/Setup/SetupAdvancedContainersItemProperties.lua:693` |
| Namespaces detected | ComboBoxGetSelectedMenuItem |
| Source kinds | lua_calls |
| Example locations | Ace: LIBGUI_Combobox:SelectedIndex, AdvancedRenownTrainer: AdvancedRenownTraining.OnExportButtonPressed, BuffHead: BuffHead.Setup.AdvancedContainersItem.OnPositionChanged, BuffHead: BuffHead.Setup.AdvancedContainersItem.OnTargetTypeChanged, BuffHead: BuffHead.Setup.AdvancedContainersItem.Properties.OnEffectsBuffsChanged, BuffHead: BuffHead.Setup.AdvancedContainersItem.Properties.OnEffectsDebuffsChanged |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 245 |
| Global usage count | 245 |
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
ComboBoxGetSelectedMenuItem(arg1)
```

## Description

Observed as a window function across 25 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a function or method reference. | Observed values: "DyeWindowDyeOrderCombo", "DyeWindowDyeTypeCombo", "EnemyChooseChannelDialogChannelList" |

## Returns

- Not confidently inferable from addon-api docs alone.

## Side Effects

- No side effect is confidently inferable from addon-api docs alone.

## Seen In

- Ace
- AdvancedRenownTrainer
- BuffHead
- Crafting Info Tooltip
- DAoCBuff
- EA_UiDebugTools
- Effigy
- Enemy
- GCDsaver
- Killer
- LibWBToggler
- MapMonster
- MapPin
- Miracle Grow Remix
- Pocket Palette
- PotionBar
- Shinies
- TexturedButtons
- TidyChat
- TidyRoll
- TurretRange
- WSCT
- WarBoard
- WoH-Reticle
- wbLeadHelper

## Examples

- Ace: LIBGUI_Combobox:SelectedIndex -> ComboBoxGetSelectedMenuItem(self.name)
- AdvancedRenownTrainer: AdvancedRenownTraining.OnExportButtonPressed -> ComboBoxGetSelectedMenuItem(PresetWindowName.."LoadComboBox")
- BuffHead: BuffHead.Setup.AdvancedContainersItem.OnPositionChanged -> ComboBoxGetSelectedMenuItem(windowName.."PositionComboBox")
- BuffHead: BuffHead.Setup.AdvancedContainersItem.OnTargetTypeChanged -> ComboBoxGetSelectedMenuItem(windowName.."TargetComboBox")
- BuffHead: BuffHead.Setup.AdvancedContainersItem.Properties.OnEffectsBuffsChanged -> ComboBoxGetSelectedMenuItem(windowName.."ElementEffectsBuffsComboBox")
- BuffHead: BuffHead.Setup.AdvancedContainersItem.Properties.OnEffectsDebuffsChanged -> ComboBoxGetSelectedMenuItem(windowName.."ElementEffectsDebuffsComboBox")

## Related APIs

- none

## Used With

- [Button](../../xml/element_types/element_Button.md) (HIGH 100/100) - XML Element Type
- [ComboBox](../../xml/element_types/element_ComboBox.md) (HIGH 100/100) - XML Element Type
- [OnLButtonUp](../../xml/handlers/handler_OnLButtonUp.md) (HIGH 100/100) - XML Handler
- [OnSelChanged](../../events/window_events/window_event_OnSelChanged.md) (HIGH 100/100) - Window Event
- [OnSelChanged](../../xml/handlers/handler_OnSelChanged.md) (HIGH 100/100) - XML Handler

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
