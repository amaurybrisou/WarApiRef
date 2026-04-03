# ScrollWindowUpdateScrollRect

- Category: Window Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 4 addons

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
| Addons seen in | DAoCBuff, Enemy, Killer, Shinies |
| Files seen in | `/workspace/data/raw/DAoCBuff/Source/DAoCBuff.lua:752`, `/workspace/data/raw/Enemy/Code/CombatLog/CombatLog.lua:1422`, `/workspace/data/raw/Enemy/Code/GroupIcons/GroupIcons.lua:522`, `/workspace/data/raw/Enemy/Code/Guard/Guard.lua:562`, `/workspace/data/raw/Enemy/Code/TalismanAlerter/TalismanAlerter.lua:182`, `/workspace/data/raw/Enemy/Code/Timer/Timer.lua:275`, `/workspace/data/raw/Enemy/Code/UnitFrames/ClickCasting.lua:202`, `/workspace/data/raw/Enemy/Code/UnitFrames/EffectsIndicator.lua:613` |
| Namespaces detected | ScrollWindowUpdateScrollRect |
| Source kinds | lua_calls |
| Example locations | DAoCBuff: DAoCBuff.ShowMessageWindow, Enemy: Enemy.CombatLogUI_ConfigDialog_OnInitialize, Enemy: Enemy.GroupIconsUI_ConfigDialog_OnInitialize, Enemy: Enemy.GuardUI_ConfigDialog_OnInitialize, Enemy: Enemy.TalismanAlerterUI_ConfigDialog_OnInitialize, Enemy: Enemy.TimerUI_ConfigDialog_OnInitialize |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 12 |
| Global usage count | 12 |
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

Observed as a window function across 4 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a function or method reference. | Observed values: "DAoCBuffMessageWindowScrollWindow", "EnemyClickCastingDialogContent", "EnemyEffectsIndicatorDialogContent" |

## Returns

- Not confidently inferable from addon-api docs alone.

## Side Effects

- No side effect is confidently inferable from addon-api docs alone.

## Seen In

- DAoCBuff
- Enemy
- Killer
- Shinies

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

- none

## Triggered By

- none

## Affects

- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Notes

- none
