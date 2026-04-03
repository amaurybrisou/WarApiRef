# WindowGetAlpha

- Category: Window Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 45 addons

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
| Addons seen in | Ace, ActionBarHide, ActionFraction, AdjustTheTip, Amethyst, CCTV, ChattyCathy, Crusher |
| Files seen in | ActionBarHide.lua, AdjustTheTip.lua, CCTV.lua, ChattyCathy.lua, Core.lua, Fader.lua, GroupIcons.lua, GuardLine.lua |
| Namespaces detected | WindowGetAlpha |
| Source kinds | lua_calls |
| Example locations | Ace: Alpha, ActionBarHide: Alpha, ActionBarHide: ChangeAlpha, ActionFraction: OnUpdate, ActionFraction: UpdateActionFractionVisibility, AdjustTheTip: OnMouseOverTargetWindowClick |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 84 |
| Global usage count | 84 |
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
WindowGetAlpha(windowName)
```

## Description

Observed querying runtime window state or metadata.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| windowName | Observed as a target window name. | Observed values: "CCTVRootWindow", "CCTVSnareWindow", "CCTVStaggerWindow" |

## Returns

- Not confidently inferable from contract artifacts alone.

## Side Effects

- No side effect is confidently inferable from contract artifacts alone.

## Seen In

- Ace
- ActionBarHide
- ActionFraction
- AdjustTheTip
- Amethyst
- CCTV
- ChattyCathy
- Crusher
- DAoCBuff
- EA_ScenarioGroupWindow
- EZCraftX
- EZGuard
- Effigy
- EveryBodyGuard
- GCDsaver
- Group Icons
- GuardLine
- Hopper
- InfoScroller
- KeyBar
- LibAddonButton
- LibWBToggler
- Map
- Motion
- NaturalLog
- PartyCast
- Pure
- Pure Careerbar
- RVMOD_SquaredDistances
- RealmStatus
- RetAlert
- RoR_SoR
- SNT_CASTBAR
- Shinies
- Statdoll
- TargetRing
- Tokens
- Vectors
- WarTriage
- Wikki's Cooldown Bar
- Wikki's Cooldown Pulse
- WoH-Reticle
- scenarioInfo
- xHUD
- xPanels

## Examples

- Ace: Alpha -> WindowGetAlpha(self.name)
- ActionBarHide: Alpha -> WindowGetAlpha(self.name)
- ActionBarHide: ChangeAlpha -> WindowGetAlpha("EA_ActionBar"..k)
- ActionFraction: OnUpdate -> WindowGetAlpha(windowName.."LabelMaximumAP")
- ActionFraction: OnUpdate -> WindowGetAlpha(windowName.."LabelCurrentAP")
- ActionFraction: UpdateActionFractionVisibility -> WindowGetAlpha(windowName.."LabelCurrentAP")

## Related APIs

- [OnInitialize](../../xml/handlers/handler_OnInitialize.md) (HIGH 88/100) - XML Event
- [OnUpdate](../../xml/handlers/handler_OnUpdate.md) (HIGH 88/100) - XML Event
- [DoesWindowExist](../../globals/functions/global_DoesWindowExist.md) (HIGH 83/100) - Global Function

## Used With

- [WindowSetAlpha](window_WindowSetAlpha.md) (HIGH 100/100) - Window Function
- [WindowSetShowing](window_WindowSetShowing.md) (HIGH 100/100) - Window Function
- [WindowStartAlphaAnimation](window_WindowStartAlphaAnimation.md) (HIGH 100/100) - Window Function

## Notes

- Advanced return analysis: No strong return evidence observed
