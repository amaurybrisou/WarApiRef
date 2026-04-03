# ComboBoxSetSelectedMenuItem

- Category: Window Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 18 addons

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
| Addons seen in | Ace, AdvancedPetAssist, BankArkel, BuffHead, DAoCBuff, Enemy, Killer, LibWBToggler |
| Files seen in | `/workspace/data/raw/Ace/LibGUI.lua:1116`, `/workspace/data/raw/AdvancedPetAssist/APAGui.lua:735`, `/workspace/data/raw/BankArkel/BankArkel.lua:364`, `/workspace/data/raw/BankArkel/BankArkel.lua:513`, `/workspace/data/raw/BuffHead/Setup/SetupAdvancedContainersItemProperties.lua:129`, `/workspace/data/raw/BuffHead/Setup/SetupAdvancedContainersItemProperties.lua:139`, `/workspace/data/raw/BuffHead/Setup/SetupAdvancedContainersItemProperties.lua:195`, `/workspace/data/raw/BuffHead/Setup/SetupAdvancedContainersItemProperties.lua:205` |
| Namespaces detected | ComboBoxSetSelectedMenuItem |
| Source kinds | lua_calls |
| Example locations | Ace: LIBGUI_Combobox:SelectIndex, AdvancedPetAssist: AdvancedPetAssist.local.FillCombo, AdvancedPetAssist: FillCombo, BankArkel: BankArkel.PackClose, BankArkel: BankArkel.SetupCombos, BuffHead: BuffHead.Setup.AdvancedContainersItem.Properties.ShowProperties |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 205 |
| Global usage count | 205 |
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
ComboBoxSetSelectedMenuItem(arg1, arg2)
```

## Description

Observed as a window function across 18 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a function or method reference. | Observed values: "BankArkelBackpackCombo", "EnemyChooseChannelDialogChannelList", "EnemyClickCastingDialogContentScrollChildAction" |
| arg2 | Observed as a function or method reference. | Observed values: #ACTIVE_CONDITION, #CONDITIONS, #CONDITIONS-1 |

## Returns

- Not confidently inferable from addon-api docs alone.

## Side Effects

- Mutates engine-visible UI state or presentation.

## Seen In

- Ace
- AdvancedPetAssist
- BankArkel
- BuffHead
- DAoCBuff
- Enemy
- Killer
- LibWBToggler
- PartyCast
- Pocket Palette
- PotionBar
- Shinies
- TexturedButtons
- TidyChat
- TidyRoll
- TurretRange
- WhoHealedMe
- WoH-Reticle

## Examples

- Ace: LIBGUI_Combobox:SelectIndex -> ComboBoxSetSelectedMenuItem(self.name, index)
- AdvancedPetAssist: AdvancedPetAssist.local.FillCombo -> ComboBoxSetSelectedMenuItem(comboName, selIdx)
- AdvancedPetAssist: FillCombo -> ComboBoxSetSelectedMenuItem(comboName, selIdx)
- BankArkel: BankArkel.PackClose -> ComboBoxSetSelectedMenuItem("BankArkelBackpackCombo", 1)
- BankArkel: BankArkel.SetupCombos -> ComboBoxSetSelectedMenuItem("BankArkelBackpackCombo", 1)
- BuffHead: BuffHead.Setup.AdvancedContainersItem.Properties.ShowProperties -> ComboBoxSetSelectedMenuItem(windowName.."PropertiesComboBox", propertyToIndex[propertyId]or 1)

## Related APIs

- none

## Used With

- [ButtonSetPressedFlag](window_ButtonSetPressedFlag.md) (HIGH 100/100) - Window Function
- [ButtonSetText](window_ButtonSetText.md) (HIGH 100/100) - Window Function
- [ComboBoxAddMenuItem](window_ComboBoxAddMenuItem.md) (HIGH 100/100) - Window Function
- [ComboBoxClearMenuItems](window_ComboBoxClearMenuItems.md) (HIGH 100/100) - Window Function
- [LabelSetText](window_LabelSetText.md) (HIGH 100/100) - Window Function
- [CreateWindow](../../globals/functions/global_CreateWindow.md) (HIGH 75/100) - Global Function

## Triggered By

- [OnLButtonUp](../../xml/handlers/handler_OnLButtonUp.md) (HIGH 100/100) - XML Event
- [OnLButtonUp](../../events/window_events/window_event_OnLButtonUp.md) (HIGH 100/100) - Window Event

## Affects

- [Button](../../xml/element_types/element_Button.md) (HIGH 100/100) - XML Element Type
- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Notes

- none
