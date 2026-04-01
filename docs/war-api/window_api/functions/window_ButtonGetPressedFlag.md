# ButtonGetPressedFlag

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
| Addons seen in | Ace, BuffHead, Crafting Info Tooltip, DAoCBuff, Effigy, Enemy, GCDsaver, Killer |
| Files seen in | `/workspace/Ace/LibGUI.lua:756`, `/workspace/Ace/LibGUI.lua:826`, `/workspace/BuffHead/Setup/SetupAdvancedContainersItemProperties.lua:701`, `/workspace/BuffHead/Setup/SetupAdvancedContainersItemProperties.lua:711`, `/workspace/BuffHead/Setup/SetupAdvancedContainersItemProperties.lua:835`, `/workspace/BuffHead/Setup/SetupAdvancedContainersItemProperties.lua:845`, `/workspace/BuffHead/Setup/SetupDisplay.lua:102`, `/workspace/BuffHead/Setup/SetupEffectCache.lua:419` |
| Namespaces detected | ButtonGetPressedFlag |
| Source kinds | lua_calls |
| Example locations | Ace: LIBGUI_Checkbox:GetValue, Ace: LIBGUI_Optionbutton:GetValue, BuffHead: BuffHead.Setup.AdvancedContainersItem.Properties.OnEffectsAlwaysShowEnableLUp, BuffHead: BuffHead.Setup.AdvancedContainersItem.Properties.OnEffectsPermanentLUp, BuffHead: BuffHead.Setup.AdvancedContainersItem.Properties.OnHandleInputEnableRemovableLUp, BuffHead: BuffHead.Setup.AdvancedContainersItem.Properties.OnHandleInputShowTooltipsLUp |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 138 |
| Global usage count | 138 |
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
ButtonGetPressedFlag(arg1)
```

## Description

Observed as a window function across 25 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a function or method reference. | Observed values: "EnemyClickCastingDialogContentScrollChildArchetype1", "EnemyClickCastingDialogContentScrollChildArchetype2", "EnemyClickCastingDialogContentScrollChildArchetype3" |

## Returns

- Not confidently inferable from addon-api docs alone.

## Side Effects

- No side effect is confidently inferable from addon-api docs alone.

## Seen In

- Ace
- BuffHead
- Crafting Info Tooltip
- DAoCBuff
- Effigy
- Enemy
- GCDsaver
- Killer
- LibGroup
- LibWBToggler
- MapMonster
- MegaphonePlusPlus
- Miracle Grow Remix
- PotionBar
- Shinies
- TexturedButtons
- TidyChat
- TidyQueue
- TidyRoll
- TurretRange
- WSCT
- WarTriage
- WoH-Reticle
- nRarity
- wbLeadHelper

## Examples

- Ace: LIBGUI_Checkbox:GetValue -> ButtonGetPressedFlag(self.name)
- Ace: LIBGUI_Optionbutton:GetValue -> ButtonGetPressedFlag(self.name)
- BuffHead: BuffHead.Setup.AdvancedContainersItem.Properties.OnEffectsAlwaysShowEnableLUp -> ButtonGetPressedFlag(windowName.."ElementEffectsAlwaysShowEnable".."Button")
- BuffHead: BuffHead.Setup.AdvancedContainersItem.Properties.OnEffectsPermanentLUp -> ButtonGetPressedFlag(windowName.."ElementEffectsPermanent".."Button")
- BuffHead: BuffHead.Setup.AdvancedContainersItem.Properties.OnHandleInputEnableRemovableLUp -> ButtonGetPressedFlag(windowName.."ElementHandleInputEnableRemovable".."Button")
- BuffHead: BuffHead.Setup.AdvancedContainersItem.Properties.OnHandleInputShowTooltipsLUp -> ButtonGetPressedFlag(windowName.."ElementHandleInputShowTooltips".."Button")

## Related APIs

- none

## Used With

- [ButtonSetDisabledFlag](window_ButtonSetDisabledFlag.md) (HIGH 100/100) - Window Function
- [ButtonSetPressedFlag](window_ButtonSetPressedFlag.md) (HIGH 100/100) - Window Function
- [ButtonSetText](window_ButtonSetText.md) (HIGH 100/100) - Window Function
- [ComboBoxAddMenuItem](window_ComboBoxAddMenuItem.md) (HIGH 100/100) - Window Function
- [ComboBoxClearMenuItems](window_ComboBoxClearMenuItems.md) (HIGH 100/100) - Window Function
- [ComboBoxGetSelectedMenuItem](window_ComboBoxGetSelectedMenuItem.md) (HIGH 100/100) - Window Function
- [ComboBoxSetSelectedMenuItem](window_ComboBoxSetSelectedMenuItem.md) (HIGH 100/100) - Window Function
- [LabelSetText](window_LabelSetText.md) (HIGH 100/100) - Window Function
- [LabelSetTextColor](window_LabelSetTextColor.md) (HIGH 100/100) - Window Function
- [LayoutEditor.RegisterWindow](window_LayoutEditor.RegisterWindow.md) (HIGH 100/100) - Window Function
- [LayoutEditor.UnregisterWindow](window_LayoutEditor.UnregisterWindow.md) (HIGH 100/100) - Window Function
- [OnLButtonUp](../../events/window_events/window_event_OnLButtonUp.md) (HIGH 100/100) - Window Event
- [OnLButtonUp](../../xml/handlers/handler_OnLButtonUp.md) (HIGH 100/100) - XML Handler
- [ScrollWindowUpdateScrollRect](window_ScrollWindowUpdateScrollRect.md) (HIGH 100/100) - Window Function
- [TextEditBoxGetText](window_TextEditBoxGetText.md) (HIGH 100/100) - Window Function
- [TextEditBoxSetText](window_TextEditBoxSetText.md) (HIGH 100/100) - Window Function
- [TextEditBoxSetTextColor](window_TextEditBoxSetTextColor.md) (HIGH 100/100) - Window Function
- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type
- [WindowSetHandleInput](window_WindowSetHandleInput.md) (HIGH 100/100) - Window Function
- [WindowSetShowing](window_WindowSetShowing.md) (HIGH 100/100) - Window Function
- [towstring](../../globals/functions/global_towstring.md) (HIGH 75/100) - Global Function

## Triggered By

- [OnLButtonUp](../../xml/handlers/handler_OnLButtonUp.md) (HIGH 100/100) - XML Handler
- [OnLButtonUp](../../events/window_events/window_event_OnLButtonUp.md) (HIGH 100/100) - Window Event

## Affects

- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Notes

- none
