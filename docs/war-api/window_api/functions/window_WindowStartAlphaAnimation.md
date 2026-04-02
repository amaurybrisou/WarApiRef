# WindowStartAlphaAnimation

- Category: Window Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 9 addons

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
| Addons seen in | AggroMeter, Aura, Busted, DAoCBuff, EA_UiDebugTools, Enemy, MapPin, RoR_SoR |
| Files seen in | `/workspace_addons/AggroMeter/AggroMeter.lua:68`, `/workspace_addons/Aura/Source/Aura.lua:186`, `/workspace_addons/Aura/Source/Aura.lua:282`, `/workspace_addons/Busted/Busted.lua:162`, `/workspace_addons/DAoCBuff/Source/DAoCBuffFrames.lua:150`, `/workspace_addons/DAoCBuff/Source/DAoCBuffSettings.lua:112`, `/workspace_addons/DAoCBuff/Source/DAoCBuffSettings.lua:161`, `/workspace_addons/DAoCBuff/Source/DAoCBuffSettings.lua:169` |
| Namespaces detected | WindowStartAlphaAnimation |
| Source kinds | lua_calls |
| Example locations | AggroMeter: AggroMeter.SplitText, Aura: Aura:Activate, Aura: Aura:Deactivate, Busted: BustedGUI.NewErrorRecorded, DAoCBuff: DAoCBuffFrame:StartFading, DAoCBuff: DAoCBuffSettings.Disable |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 61 |
| Global usage count | 61 |
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

Observed as a window function across 9 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a text or wstring payload. | Observed values: "AggroMeterWindow"..MobID, "BustedMiniGUIText", "DAoCBuff_Settings" |
| arg2 | Observed as a function or method reference. | Observed values: 1, 2, Window.AnimationType.EASE_OUT |
| arg3 | Observed as a function or method reference. | Observed values: 0, 0.5, 0.7 |
| arg4 | Observed as a function or method reference. | Observed values: 0, 0.0, 0.2 |
| arg5 | Observed as a function or method reference. | Observed values: 0.2, 0.3, 0.4 |
| arg6 | Observed as a boolean toggle. | Observed values: anim.setStartBeforeDelay, false, true |
| arg7 | Observed as a function or method reference. | Observed values: 0, 0.0, 1.5 |
| arg8 | Observed as a numeric value. | Observed values: 0, 1, anim.loopCount |

## Returns

- Not confidently inferable from addon-api docs alone.

## Side Effects

- No side effect is confidently inferable from addon-api docs alone.

## Seen In

- AggroMeter
- Aura
- Busted
- DAoCBuff
- EA_UiDebugTools
- Enemy
- MapPin
- RoR_SoR
- TurretRange

## Examples

- AggroMeter: AggroMeter.SplitText -> WindowStartAlphaAnimation("AggroMeterWindow"..MobID, Window.AnimationType.SINGLE_NO_RESET, 0, 1, 0.5, false, 0, 0)
- Aura: Aura:Activate -> WindowStartAlphaAnimation(self:Get("internal-runtimeflashwindowid"), anim.type, anim.startAlpha, self:Get("texture-alpha"), anim.duration, anim.setStartBeforeDelay, anim.delay, anim.loopCount)
- Aura: Aura:Activate -> WindowStartAlphaAnimation(self:Get("internal-runtimewindowid"), anim.type, anim.startAlpha, anim.endAlpha, anim.duration, anim.setStartBeforeDelay, anim.delay, anim.loopCount)
- Aura: Aura:Deactivate -> WindowStartAlphaAnimation(self:Get("internal-runtimewindowid"), anim.type, anim.startAlpha, anim.endAlpha, anim.duration, anim.setStartBeforeDelay, anim.delay, anim.loopCount)
- Aura: Aura:Deactivate -> WindowStartAlphaAnimation(self:Get("internal-runtimetimerwindowid"), anim.type, anim.startAlpha, anim.endAlpha, anim.duration, anim.setStartBeforeDelay, anim.delay, anim.loopCount)
- Aura: Aura:Deactivate -> WindowStartAlphaAnimation(self:Get("internal-runtimeflashwindowid"), anim.type, anim.startAlpha, anim.endAlpha, anim.duration, anim.setStartBeforeDelay, anim.delay, anim.loopCount)

## Related APIs

- none

## Used With

- [AlertTextWindow.AddLine](../../globals/functions/global_AlertTextWindow.AddLine.md) (HIGH 100/100) - Global Function
- [WindowSetShowing](window_WindowSetShowing.md) (HIGH 100/100) - Window Function
- [WindowStopAlphaAnimation](window_WindowStopAlphaAnimation.md) (HIGH 100/100) - Window Function

## Triggered By

- none

## Affects

- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Notes

- none
