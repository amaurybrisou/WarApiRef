# ScrollWindowUpdateScrollRect

- Category: Window Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 7 addons

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
| Addons seen in | DAoCBuff, Enemy, Killer, Miracle Grow Remix, RVMOD_Manager, Shinies, wbLeadHelper |
| Files seen in | `/workspace/DAoCBuff/Source/DAoCBuff.lua:752`, `/workspace/Enemy/Code/CombatLog/CombatLog.lua:1422`, `/workspace/Enemy/Code/GroupIcons/GroupIcons.lua:522`, `/workspace/Enemy/Code/Guard/Guard.lua:562`, `/workspace/Enemy/Code/TalismanAlerter/TalismanAlerter.lua:182`, `/workspace/Enemy/Code/Timer/Timer.lua:275`, `/workspace/Enemy/Code/UnitFrames/ClickCasting.lua:202`, `/workspace/Enemy/Code/UnitFrames/EffectsIndicator.lua:613` |
| Namespaces detected | ScrollWindowUpdateScrollRect |
| Source kinds | lua_calls |
| Example locations | DAoCBuff: DAoCBuff.ShowMessageWindow, Enemy: Enemy.CombatLogUI_ConfigDialog_OnInitialize, Enemy: Enemy.GroupIconsUI_ConfigDialog_OnInitialize, Enemy: Enemy.GuardUI_ConfigDialog_OnInitialize, Enemy: Enemy.TalismanAlerterUI_ConfigDialog_OnInitialize, Enemy: Enemy.TimerUI_ConfigDialog_OnInitialize |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 21 |
| Global usage count | 21 |
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
ScrollWindowUpdateScrollRect(arg1)
```

## Description

Observed as a window function across 7 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a function or method reference. | Observed values: "DAoCBuffMessageWindowScrollWindow", "EnemyClickCastingDialogContent", "EnemyEffectsIndicatorDialogContent" |

## Returns

- Not confidently inferable from API_Ref alone.

## Side Effects

- No side effect is confidently inferable from API_Ref alone.

## Seen In

- DAoCBuff
- Enemy
- Killer
- Miracle Grow Remix
- RVMOD_Manager
- Shinies
- wbLeadHelper

## Examples

- DAoCBuff: DAoCBuff.ShowMessageWindow -> ScrollWindowUpdateScrollRect("DAoCBuffMessageWindowScrollWindow")
- Enemy: Enemy.CombatLogUI_ConfigDialog_OnInitialize -> ScrollWindowUpdateScrollRect(config_dlg.section.windowName.."Content")
- Enemy: Enemy.GroupIconsUI_ConfigDialog_OnInitialize -> ScrollWindowUpdateScrollRect(dlg.section.windowName.."Content")
- Enemy: Enemy.GuardUI_ConfigDialog_OnInitialize -> ScrollWindowUpdateScrollRect(config_dlg.section.windowName.."Content")
- Enemy: Enemy.TalismanAlerterUI_ConfigDialog_OnInitialize -> ScrollWindowUpdateScrollRect(config_dlg.section.windowName.."Content")
- Enemy: Enemy.TimerUI_ConfigDialog_OnInitialize -> ScrollWindowUpdateScrollRect(config_dlg.section.windowName.."Content")

## Related APIs

- none

## Used With

- [ButtonGetPressedFlag](window_ButtonGetPressedFlag.md) (HIGH 100/100) - Window Function
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
- [ScrollWindowSetOffset](window_ScrollWindowSetOffset.md) (HIGH 100/100) - Window Function
- [TextEditBoxGetText](window_TextEditBoxGetText.md) (HIGH 100/100) - Window Function
- [TextEditBoxSetText](window_TextEditBoxSetText.md) (HIGH 100/100) - Window Function
- [TextEditBoxSetTextColor](window_TextEditBoxSetTextColor.md) (HIGH 100/100) - Window Function
- [WindowSetHandleInput](window_WindowSetHandleInput.md) (HIGH 100/100) - Window Function
- [WindowSetShowing](window_WindowSetShowing.md) (HIGH 100/100) - Window Function
- [CreateWindow](../../globals/functions/global_CreateWindow.md) (HIGH 75/100) - Global Function
- [towstring](../../globals/functions/global_towstring.md) (HIGH 75/100) - Global Function

## Triggered By

- none

## Affects

- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Notes

- none
