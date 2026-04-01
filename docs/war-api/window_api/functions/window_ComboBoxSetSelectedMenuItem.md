# ComboBoxSetSelectedMenuItem

- Category: Window Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 32 addons

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
| Addons seen in | Ace, AdvancedPetAssist, BankArkel, BuffHead, Busted, Crafting Info Tooltip, DAoCBuff, DaemonAssist |
| Files seen in | `/workspace/Ace/LibGUI.lua:1116`, `/workspace/AdvancedPetAssist/APAGui.lua:735`, `/workspace/BankArkel/BankArkel.lua:364`, `/workspace/BankArkel/BankArkel.lua:513`, `/workspace/BuffHead/Setup/SetupAdvancedContainersItemProperties.lua:129`, `/workspace/BuffHead/Setup/SetupAdvancedContainersItemProperties.lua:139`, `/workspace/BuffHead/Setup/SetupAdvancedContainersItemProperties.lua:195`, `/workspace/BuffHead/Setup/SetupAdvancedContainersItemProperties.lua:205` |
| Namespaces detected | ComboBoxSetSelectedMenuItem |
| Source kinds | lua_calls |
| Example locations | Ace: LIBGUI_Combobox:SelectIndex, AdvancedPetAssist: AdvancedPetAssist.local.FillCombo, AdvancedPetAssist: FillCombo, BankArkel: BankArkel.PackClose, BankArkel: BankArkel.SetupCombos, BuffHead: BuffHead.Setup.AdvancedContainersItem.Properties.ShowProperties |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 263 |
| Global usage count | 263 |
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

Observed as a window function across 32 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a function or method reference. | Observed values: "BankArkelBackpackCombo", "BustedGUIAddonSelect", "EnemyChooseChannelDialogChannelList" |
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
- Busted
- Crafting Info Tooltip
- DAoCBuff
- DaemonAssist
- DeepSleep
- EA_UiDebugTools
- Effigy
- Enemy
- GCDsaver
- Killer
- LibWBToggler
- LoyalPet
- MapMonster
- MapPin
- MegaphonePlusPlus
- Miracle Grow Remix
- Pocket Palette
- PotionBar
- Shinies
- TexturedButtons
- ThinkOutLoud
- TidyChat
- TidyRoll
- TurretRange
- WarTriage
- WhoHealedMe
- WoH-Reticle
- wbLeadHelper

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

- [Button](../../xml/element_types/element_Button.md) (HIGH 100/100) - XML Element Type
- [ButtonGetPressedFlag](window_ButtonGetPressedFlag.md) (HIGH 100/100) - Window Function
- [ButtonSetDisabledFlag](window_ButtonSetDisabledFlag.md) (HIGH 100/100) - Window Function
- [ButtonSetPressedFlag](window_ButtonSetPressedFlag.md) (HIGH 100/100) - Window Function
- [ButtonSetText](window_ButtonSetText.md) (HIGH 100/100) - Window Function
- [ComboBoxAddMenuItem](window_ComboBoxAddMenuItem.md) (HIGH 100/100) - Window Function
- [ComboBoxClearMenuItems](window_ComboBoxClearMenuItems.md) (HIGH 100/100) - Window Function
- [ComboBoxGetSelectedMenuItem](window_ComboBoxGetSelectedMenuItem.md) (HIGH 100/100) - Window Function
- [DialogManager.MakeTextEntryDialog](../../globals/functions/global_DialogManager.MakeTextEntryDialog.md) (HIGH 100/100) - Global Function
- [LabelSetText](window_LabelSetText.md) (HIGH 100/100) - Window Function
- [LabelSetTextColor](window_LabelSetTextColor.md) (HIGH 100/100) - Window Function
- [LayoutEditor.RegisterWindow](window_LayoutEditor.RegisterWindow.md) (HIGH 100/100) - Window Function
- [LayoutEditor.UnregisterWindow](window_LayoutEditor.UnregisterWindow.md) (HIGH 100/100) - Window Function
- [OnLButtonUp](../../xml/handlers/handler_OnLButtonUp.md) (HIGH 100/100) - XML Handler
- [OnLButtonUp](../../events/window_events/window_event_OnLButtonUp.md) (HIGH 100/100) - Window Event
- [ScrollWindowSetOffset](window_ScrollWindowSetOffset.md) (HIGH 100/100) - Window Function
- [ScrollWindowUpdateScrollRect](window_ScrollWindowUpdateScrollRect.md) (HIGH 100/100) - Window Function
- [TextEditBoxGetText](window_TextEditBoxGetText.md) (HIGH 100/100) - Window Function
- [TextEditBoxSetText](window_TextEditBoxSetText.md) (HIGH 100/100) - Window Function
- [TextEditBoxSetTextColor](window_TextEditBoxSetTextColor.md) (HIGH 100/100) - Window Function
- [WindowAddAnchor](window_WindowAddAnchor.md) (HIGH 100/100) - Window Function
- [WindowClearAnchors](window_WindowClearAnchors.md) (HIGH 100/100) - Window Function
- [WindowSetHandleInput](window_WindowSetHandleInput.md) (HIGH 100/100) - Window Function
- [WindowSetLayer](window_WindowSetLayer.md) (HIGH 100/100) - Window Function
- [WindowSetShowing](window_WindowSetShowing.md) (HIGH 100/100) - Window Function
- [WindowSetTintColor](window_WindowSetTintColor.md) (HIGH 100/100) - Window Function
- [CreateWindow](../../globals/functions/global_CreateWindow.md) (HIGH 75/100) - Global Function
- [towstring](../../globals/functions/global_towstring.md) (HIGH 75/100) - Global Function

## Triggered By

- [OnLButtonUp](../../xml/handlers/handler_OnLButtonUp.md) (HIGH 100/100) - XML Handler
- [OnLButtonUp](../../events/window_events/window_event_OnLButtonUp.md) (HIGH 100/100) - Window Event

## Affects

- [Button](../../xml/element_types/element_Button.md) (HIGH 100/100) - XML Element Type
- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Notes

- none
