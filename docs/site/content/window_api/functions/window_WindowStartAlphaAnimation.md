# WindowStartAlphaAnimation

- Category: Window Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 10 addons

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
| Addons seen in | AggroMeter, Aura, Busted, DAoCBuff, EA_UiDebugTools, Enemy, MapPin, PeaceOut |
| Files seen in | `/workspace/AggroMeter/AggroMeter.lua:68`, `/workspace/Aura/Source/Aura.lua:186`, `/workspace/Aura/Source/Aura.lua:282`, `/workspace/Busted/Busted.lua:162`, `/workspace/DAoCBuff/Source/DAoCBuffFrames.lua:150`, `/workspace/DAoCBuff/Source/DAoCBuffSettings.lua:112`, `/workspace/DAoCBuff/Source/DAoCBuffSettings.lua:161`, `/workspace/DAoCBuff/Source/DAoCBuffSettings.lua:169` |
| Namespaces detected | WindowStartAlphaAnimation |
| Source kinds | lua_calls |
| Example locations | AggroMeter: AggroMeter.SplitText, Aura: Aura:Activate, Aura: Aura:Deactivate, Busted: BustedGUI.NewErrorRecorded, DAoCBuff: DAoCBuffFrame:StartFading, DAoCBuff: DAoCBuffSettings.Disable |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 63 |
| Global usage count | 63 |
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

Observed as a window function across 10 addons.

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

- Not confidently inferable from API_Ref alone.

## Side Effects

- No side effect is confidently inferable from API_Ref alone.

## Seen In

- AggroMeter
- Aura
- Busted
- DAoCBuff
- EA_UiDebugTools
- Enemy
- MapPin
- PeaceOut
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
- [EA_Window_ContextMenu.AddCascadingMenuItem](../../globals/functions/global_EA_Window_ContextMenu.AddCascadingMenuItem.md) (HIGH 100/100) - Global Function
- [EA_Window_ContextMenu.AddMenuDivider](../../globals/functions/global_EA_Window_ContextMenu.AddMenuDivider.md) (HIGH 100/100) - Global Function
- [EA_Window_ContextMenu.AddMenuItem](../../globals/functions/global_EA_Window_ContextMenu.AddMenuItem.md) (HIGH 100/100) - Global Function
- [EA_Window_ContextMenu.CreateContextMenu](../../globals/functions/global_EA_Window_ContextMenu.CreateContextMenu.md) (HIGH 100/100) - Global Function
- [EA_Window_ContextMenu.Finalize](../../globals/functions/global_EA_Window_ContextMenu.Finalize.md) (HIGH 100/100) - Global Function
- [LabelSetText](window_LabelSetText.md) (HIGH 100/100) - Window Function
- [PartyUtils.IsPartyActive](../../globals/functions/global_PartyUtils.IsPartyActive.md) (HIGH 100/100) - Global Function
- [SystemData.Events.UPDATE_PROCESSED](../../systemdata/fields/systemdata_SystemData.Events.UPDATE_PROCESSED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.UPDATE_PROCESSED](../../events/game_events/game_event_SystemData.Events.UPDATE_PROCESSED.md) (HIGH 100/100) - Game Event
- [WindowGetAlpha](window_WindowGetAlpha.md) (HIGH 100/100) - Window Function
- [WindowGetDimensions](window_WindowGetDimensions.md) (HIGH 100/100) - Window Function
- [WindowGetHandleInput](window_WindowGetHandleInput.md) (HIGH 100/100) - Window Function
- [WindowGetScale](window_WindowGetScale.md) (HIGH 100/100) - Window Function
- [WindowGetShowing](window_WindowGetShowing.md) (HIGH 100/100) - Window Function
- [WindowSetHandleInput](window_WindowSetHandleInput.md) (HIGH 100/100) - Window Function
- [WindowSetScale](window_WindowSetScale.md) (HIGH 100/100) - Window Function
- [WindowSetShowing](window_WindowSetShowing.md) (HIGH 100/100) - Window Function
- [WindowSetTintColor](window_WindowSetTintColor.md) (HIGH 100/100) - Window Function
- [WindowStopAlphaAnimation](window_WindowStopAlphaAnimation.md) (HIGH 100/100) - Window Function
- [WindowStartScaleAnimation](window_WindowStartScaleAnimation.md) (HIGH 90/100) - Window Function
- [DoesWindowExist](../../globals/functions/global_DoesWindowExist.md) (HIGH 83/100) - Global Function
- [GetIconData](../../globals/functions/global_GetIconData.md) (HIGH 83/100) - Global Function
- [CreateWindowFromTemplate](../../globals/functions/global_CreateWindowFromTemplate.md) (HIGH 75/100) - Global Function
- [towstring](../../globals/functions/global_towstring.md) (HIGH 75/100) - Global Function
- [UnregisterEventHandler](../../globals/functions/global_UnregisterEventHandler.md) (MEDIUM 68/100) - Global Function

## Triggered By

- none

## Affects

- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Notes

- none
