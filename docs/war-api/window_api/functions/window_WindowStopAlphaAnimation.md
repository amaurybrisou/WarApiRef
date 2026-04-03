# WindowStopAlphaAnimation

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
| Addons seen in | Aura, DAoCBuff, Enemy, Pocket Palette, RoR_SoR, TidyChat, TurretRange |
| Files seen in | `/workspace/data/raw/Aura/Source/Aura.lua:282`, `/workspace/data/raw/DAoCBuff/Source/DAoCBuffFrames.lua:141`, `/workspace/data/raw/Enemy/Code/Guard/Guard.lua:323`, `/workspace/data/raw/Enemy/Code/TalismanAlerter/TalismanAlerter.lua:77`, `/workspace/data/raw/PocketPalette/PocketPalette.lua:91`, `/workspace/data/raw/RoR_SoR/RoR_SoR.lua:1055`, `/workspace/data/raw/TidyChat/TidyChat.lua:344`, `/workspace/data/raw/TurrentRange/Display.lua:338` |
| Namespaces detected | WindowStopAlphaAnimation |
| Source kinds | lua_calls |
| Example locations | Aura: Aura:Deactivate, DAoCBuff: DAoCBuffFrame:StopFading, Enemy: Enemy.Guard_GuardIndicator_Update, Enemy: Enemy.TalismanAlerter_Update, Pocket Palette: HighlightWindow, RoR_SoR: RoR_SoR.SET_KEEP |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 22 |
| Global usage count | 22 |
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
WindowStopAlphaAnimation(arg1)
```

## Description

Observed as a window function across 7 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a text or wstring payload. | Observed values: "EnemyGuardDistanceIndicator", "EnemyTalismanAlerterIndicator", "SoR_"..Window_Name.."KEEP1KEEPDOOR1" |

## Returns

- Not confidently inferable from addon-api docs alone.

## Side Effects

- No side effect is confidently inferable from addon-api docs alone.

## Seen In

- Aura
- DAoCBuff
- Enemy
- Pocket Palette
- RoR_SoR
- TidyChat
- TurretRange

## Examples

- Aura: Aura:Deactivate -> WindowStopAlphaAnimation(self:Get("internal-runtimewindowid"))
- Aura: Aura:Deactivate -> WindowStopAlphaAnimation(self:Get("internal-runtimetimerwindowid"))
- Aura: Aura:Deactivate -> WindowStopAlphaAnimation(self:Get("internal-runtimeflashwindowid"))
- DAoCBuff: DAoCBuffFrame:StopFading -> WindowStopAlphaAnimation(self.icon)
- DAoCBuff: DAoCBuffFrame:StopFading -> WindowStopAlphaAnimation(self.frame)
- Enemy: Enemy.Guard_GuardIndicator_Update -> WindowStopAlphaAnimation("EnemyGuardDistanceIndicator")

## Related APIs

- none

## Used With

- [WindowSetShowing](window_WindowSetShowing.md) (HIGH 100/100) - Window Function
- [WindowStartAlphaAnimation](window_WindowStartAlphaAnimation.md) (HIGH 100/100) - Window Function

## Triggered By

- none

## Affects

- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Notes

- none
