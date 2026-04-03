# WindowStopAlphaAnimation

- Category: Window Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 22 addons

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
| Addons seen in | AdjustTheTip, Amethyst, Aura, Busted, CCTV, CDown, Calling, DAoCBuff |
| Files seen in | AdjustTheTip.lua, Amethyst.lua, Busted.lua, CCTV.lua, CDownFrames.lua, CallingNotification.lua, Code/Guard/Guard.lua, Code/TalismanAlerter/TalismanAlerter.lua |
| Namespaces detected | WindowStopAlphaAnimation |
| Source kinds | lua_calls |
| Example locations | AdjustTheTip: UpdateUnit, Amethyst: ShowCastBar, Amethyst: ShowDummyCastBar, Amethyst: StartInteractTimer, Aura: Deactivate, Busted: ClearAlertFlash |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 53 |
| Global usage count | 53 |
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

Observed as a window function across 22 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a text or wstring payload. | Observed values: "BustedMiniGUIText", "CCTVRootWindow", "CCTVRootWindowBackground2" |

## Returns

- Not confidently inferable from contract artifacts alone.

## Side Effects

- No side effect is confidently inferable from contract artifacts alone.

## Seen In

- AdjustTheTip
- Amethyst
- Aura
- Busted
- CCTV
- CDown
- Calling
- DAoCBuff
- DuffTimer
- EA_UiDebugTools
- Enemy
- LibAddonButton
- Minmap
- Pocket Palette
- Pure
- RVMOD_PlayerStatus
- RVMOD_SquaredDistances
- RoR_SoR
- SNT_CASTBAR
- TidyChat
- TurretRange
- Vectors

## Examples

- AdjustTheTip: UpdateUnit -> WindowStopAlphaAnimation(self:GetName())
- Amethyst: ShowCastBar -> WindowStopAlphaAnimation(C.name)
- Amethyst: ShowDummyCastBar -> WindowStopAlphaAnimation(C.name)
- Amethyst: StartInteractTimer -> WindowStopAlphaAnimation(C.name)
- Aura: Deactivate -> WindowStopAlphaAnimation(self:Get("internal-runtimewindowid"))
- Aura: Deactivate -> WindowStopAlphaAnimation(self:Get("internal-runtimetimerwindowid"))

## Used With

- [WindowSetShowing](window_WindowSetShowing.md) (HIGH 100/100) - Window Function
- [WindowStartAlphaAnimation](window_WindowStartAlphaAnimation.md) (HIGH 100/100) - Window Function

## Notes

- none
