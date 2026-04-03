# WindowSetHandleInput

- Category: Window Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 47 addons

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
| Addons seen in | Ace, ActionBarHide, ActionFraction, AdjustTheTip, Amethyst, BuffHead, CCTV, CleanUnitFrames |
| Files seen in | AdjustTheTip.lua, CCTV.lua, CleanGroupMemberUnitFrame.lua, Code/Assist/Assist.lua, Code/GroupIcons/GroupIcon.lua, Code/Guard/Guard.lua, Code/Marks/MarkTemplate.lua, Code/UnitFrames/UnitFramePart.lua |
| Namespaces detected | WindowSetHandleInput |
| Source kinds | lua_calls |
| Example locations | Ace: CaptureInput, Ace: IgnoreInput, Ace: MakeMovable, ActionBarHide: CaptureInput, ActionBarHide: IgnoreInput, ActionBarHide: MakeMovable |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 178 |
| Global usage count | 178 |
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
WindowSetHandleInput(windowName, arg2)
```

## Description

Observed mutating runtime window state or presentation.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| windowName | Observed as a target window name. | Observed values: "CCTVRootWindow", "CCTVSnareWindow", "CCTVStaggerWindow" |
| arg2 | Observed as a boolean toggle. | Observed values: (handleInput==true), Pure.Get("group-frame-handleinput"), Pure.Get("grouppet-frame-handleinput") |

## Returns

- Not confidently inferable from contract artifacts alone.

## Side Effects

- Mutates engine-visible UI state or presentation.

## Seen In

- Ace
- ActionBarHide
- ActionFraction
- AdjustTheTip
- Amethyst
- BuffHead
- CCTV
- CleanUnitFrames
- Crusher
- DAoCBuff
- DuffTimer
- EA_ScenarioGroupWindow
- EZCraftX
- EZGuard
- Effigy
- Enemy
- EveryBodyGuard
- FastFriends
- GCDsaver
- GroupSpotter
- Hopper
- InfoScroller
- KeyBar
- LibWBToggler
- Map
- Miracle Grow Remix
- Motion
- NaturalLog
- PartyCast
- Pure
- Pure Careerbar
- RealmStatus
- Shinies
- Statdoll Remix
- TargetRing
- TidyChat
- TidyRoll
- Tokens
- Vectors
- WarTriage
- Wikki's Cooldown Bar
- Wikki's Cooldown Pulse
- WoH-Reticle
- nRarity
- scenarioInfo
- xHUD
- xPanels

## Examples

- Ace: CaptureInput -> WindowSetHandleInput(self.name, true)
- Ace: IgnoreInput -> WindowSetHandleInput(self.name, false)
- Ace: MakeMovable -> WindowSetHandleInput(self.name, true)
- ActionBarHide: CaptureInput -> WindowSetHandleInput(self.name, true)
- ActionBarHide: IgnoreInput -> WindowSetHandleInput(self.name, false)
- ActionBarHide: MakeMovable -> WindowSetHandleInput(self.name, true)

## Used With

- [WindowSetShowing](window_WindowSetShowing.md) (HIGH 100/100) - Window Function

## Notes

- none
