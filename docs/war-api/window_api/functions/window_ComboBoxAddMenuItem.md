# ComboBoxAddMenuItem

- Category: Window Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 14 addons

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
| Addons seen in | Ace, BankArkel, BuffHead, DAoCBuff, Enemy, Killer, LibWBToggler, PartyCast |
| Files seen in | `/workspace/data/raw/Ace/LibGUI.lua:1068`, `/workspace/data/raw/BankArkel/BankArkel.lua:513`, `/workspace/data/raw/BuffHead/Setup/SetupAdvancedCompressionItemEffect.lua:70`, `/workspace/data/raw/BuffHead/Setup/SetupAdvancedContainersItem.lua:342`, `/workspace/data/raw/BuffHead/Setup/SetupAdvancedContainersItemProperties.lua:205`, `/workspace/data/raw/BuffHead/Setup/SetupAdvancedContainersItemProperties.lua:283`, `/workspace/data/raw/BuffHead/Setup/SetupAdvancedContainersItemProperties.lua:431`, `/workspace/data/raw/BuffHead/Setup/SetupContainer.lua:187` |
| Namespaces detected | ComboBoxAddMenuItem |
| Source kinds | lua_calls |
| Example locations | Ace: LIBGUI_Combobox:Add, BankArkel: BankArkel.SetupCombos, BuffHead: BuffHead.Setup.AdvancedCompressionItemEffect.Initialize, BuffHead: BuffHead.Setup.AdvancedContainersItem.Initialize, BuffHead: BuffHead.Setup.AdvancedContainersItem.Properties.Initialize, BuffHead: BuffHead.Setup.Container.Initialize |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 295 |
| Global usage count | 295 |
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
ComboBoxAddMenuItem(arg1, arg2)
```

## Description

Observed as a window function across 14 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a function or method reference. | Observed values: "BankArkelBackpackCombo", "EnemyChooseChannelDialogChannelList", "EnemyClickCastingDialogContentScrollChildAction" |
| arg2 | Observed as a function or method reference. | Observed values: BankArkel.db.Entry[i].Name, L "---", L "0.#" |

## Returns

- Not confidently inferable from addon-api docs alone.

## Side Effects

- No side effect is confidently inferable from addon-api docs alone.

## Seen In

- Ace
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
- WoH-Reticle

## Examples

- Ace: LIBGUI_Combobox:Add -> ComboBoxAddMenuItem(self.name, itemText)
- BankArkel: BankArkel.SetupCombos -> ComboBoxAddMenuItem("BankArkelBackpackCombo", StringToWString(tbNone))
- BankArkel: BankArkel.SetupCombos -> ComboBoxAddMenuItem("BankArkelBackpackCombo", StringToWString(tbSelf))
- BankArkel: BankArkel.SetupCombos -> ComboBoxAddMenuItem("BankArkelBackpackCombo", BankArkel.db.Entry[i].Name)
- BuffHead: BuffHead.Setup.AdvancedCompressionItemEffect.Initialize -> ComboBoxAddMenuItem(windowName.."CastByComboBox", localization["Setup.AdvancedCompressionItemEffect.CastBy.Self"])
- BuffHead: BuffHead.Setup.AdvancedCompressionItemEffect.Initialize -> ComboBoxAddMenuItem(windowName.."CastByComboBox", localization["Setup.AdvancedCompressionItemEffect.CastBy.Others"])

## Related APIs

- none

## Used With

- [ButtonSetPressedFlag](window_ButtonSetPressedFlag.md) (HIGH 100/100) - Window Function
- [ButtonSetText](window_ButtonSetText.md) (HIGH 100/100) - Window Function
- [ComboBoxClearMenuItems](window_ComboBoxClearMenuItems.md) (HIGH 100/100) - Window Function
- [ComboBoxSetSelectedMenuItem](window_ComboBoxSetSelectedMenuItem.md) (HIGH 100/100) - Window Function
- [LabelSetText](window_LabelSetText.md) (HIGH 100/100) - Window Function

## Triggered By

- none

## Affects

- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Notes

- none
