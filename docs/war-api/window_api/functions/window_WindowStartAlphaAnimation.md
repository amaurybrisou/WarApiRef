# WindowStartAlphaAnimation

- Category: Window Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 35 addons

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
| Addons seen in | AbilityAlert, AbilityNotifier, ActionBarHide, ActionFraction, AdjustTheTip, AggroMeter, Amethyst, Aura |
| Files seen in | AbilityAlert.lua, AbilityNotifier.lua, ActionBarHide.lua, AdjustTheTip.lua, AggroMeter.lua, Amethyst.lua, Busted.lua, CCTV.lua |
| Namespaces detected | WindowStartAlphaAnimation |
| Source kinds | lua_calls |
| Example locations | AbilityAlert: Display, AbilityNotifier: Notify, ActionBarHide: ChangeAlpha, ActionFraction: OnUpdate, ActionFraction: UpdateActionFractionVisibility, AdjustTheTip: UpdateUnit |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 127 |
| Global usage count | 127 |
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
WindowStartAlphaAnimation(arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8)
```

## Description

Observed as a window function across 35 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a text or wstring payload. | Observed values: "AAWindow", "AbHelpWindow", "AggroMeterWindow"..MobID |
| arg2 | Observed as a function or method reference. | Observed values: 1, 2, SINGLE_NO_RESET |
| arg3 | Observed as a function or method reference. | Observed values: 0, 0.0, 0.1 |
| arg4 | Observed as a function or method reference. | Observed values: 0, 0.0, 0.2 |
| arg5 | Observed as a function or method reference. | Observed values: (CCTV.SETTINGS.FADEOUT)+0.1, .02, .2 |
| arg6 | Observed as a boolean toggle. | Observed values: anim.setStartBeforeDelay, false, hidden~=nil |
| arg7 | Observed as a numeric value. | Observed values: 0, 0.0, 0.5 |
| arg8 | Observed as a numeric value. | Observed values: 0, 1, anim.loopCount |

## Returns

- Not confidently inferable from contract artifacts alone.

## Side Effects

- No side effect is confidently inferable from contract artifacts alone.

## Seen In

- AbilityAlert
- AbilityNotifier
- ActionBarHide
- ActionFraction
- AdjustTheTip
- AggroMeter
- Amethyst
- Aura
- Busted
- CCTV
- CDown
- ChattyCathy
- Countdown
- DAoCBuff
- EA_LoadingScreen
- EA_OpenPartyWindow
- EA_UiDebugTools
- Emojii
- Enemy
- EveryBodyGuard
- FlagCap
- GCDTracker
- KeyBar
- LibAddonButton
- MapPin
- Minmap
- Pure
- RVMOD_PlayerStatus
- RVMOD_SquaredDistances
- RealmStatus
- RoR_SoR
- SNT_CASTBAR
- TurretRange
- Vectors
- Wikki's Cooldown Pulse

## Examples

- AbilityAlert: Display -> WindowStartAlphaAnimation("AAWindow", Window.AnimationType.EASE_OUT, 1.0, 0.0, 2, false, 0, 0)
- AbilityNotifier: Notify -> WindowStartAlphaAnimation("AbHelpWindow", Window.AnimationType.EASE_OUT, 1, 0, 2, false, 0, 1)
- ActionBarHide: ChangeAlpha -> WindowStartAlphaAnimation("EA_ActionBar"..k, Window.AnimationType.SINGLE_NO_RESET, alphastart, targetalpha, ActionBarHide.Settings.FadeInTime, false, ActionBarHide.Settings.FadeInDelay, 0)
- ActionBarHide: ChangeAlpha -> WindowStartAlphaAnimation("EA_ActionBar"..k, Window.AnimationType.SINGLE_NO_RESET, alphastart, targetalpha, ActionBarHide.Settings.FadeOutTime, false, ActionBarHide.Settings.FadeOutDelay, 0)
- ActionFraction: OnUpdate -> WindowStartAlphaAnimation(windowName.."LabelCurrentAP", Window.AnimationType.SINGLE_NO_RESET_HIDE, 1.0, 0.0, 2.0, false, 0, 0)
- ActionFraction: OnUpdate -> WindowStartAlphaAnimation(windowName.."LabelMaximumAP", Window.AnimationType.SINGLE_NO_RESET_HIDE, 1.0, 0.0, 2.0, false, 0, 0)

## Related APIs

- [OnUpdate](../../xml/handlers/handler_OnUpdate.md) (HIGH 88/100) - XML Event

## Used With

- [WindowGetAlpha](window_WindowGetAlpha.md) (HIGH 100/100) - Window Function
- [WindowSetShowing](window_WindowSetShowing.md) (HIGH 100/100) - Window Function
- [WindowStopAlphaAnimation](window_WindowStopAlphaAnimation.md) (HIGH 100/100) - Window Function
- [AlertTextWindow.AddLine](../../globals/functions/global_AlertTextWindow.AddLine.md) (HIGH 75/100) - Global Function

## Notes

- none
