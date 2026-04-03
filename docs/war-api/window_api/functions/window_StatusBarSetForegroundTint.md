# StatusBarSetForegroundTint

- Category: Window Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 35 addons

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 110

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, seen in 4 or more addons, matches a known engine namespace.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- -15 Conflicting signatures across usages: Observed arity or argument shape conflicts across usages.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | Ace, ActionBarHide, AdjustTheTip, Amethyst, CMap, Crusher, EZCraftX, EZGuard |
| Files seen in | AdjustTheTip.lua, LibGUI.lua, LibGui.lua, Libraries/LibGUI.lua, Libs/LibGUI.lua, PartyCast.lua, RoR_SoR.lua, TidyRollFrame.lua |
| Namespaces detected | StatusBarSetForegroundTint |
| Source kinds | lua_calls |
| Example locations | Ace: ForeColor, ActionBarHide: ForeColor, AdjustTheTip: AddTargetHealthToMouseOver, Amethyst: ForeColor, CMap: ForeColor, Crusher: ForeColor |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 43 |
| Global usage count | 43 |
| Local definition count | 0 |
| Documentation references | 0 |
| Initialization flow references | 0 |
| Known engine namespace | yes |
| Default UI presence | yes |
| Event binding presence | no |
| Observed in XML and Lua | no |
| Consistent role | yes |
| Consistent arguments | no |
| Consistent returns | no |
| Slash command presence | no |
| Weak usage only | no |
| Project-specific name | no |
| Placeholder or computed name | no |
| Conflicting signatures | yes |
| Conflicting roles | no |
| Wrapper likely | no |
| Never outside local graph | no |
| Local helper only | no |

## Signature (inferred)

```lua
StatusBarSetForegroundTint(arg1, arg2, arg3, arg4)
```

## Description

Observed as a window function across 35 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a function or method reference. | Observed values: "PartyCastWindow"..PlayerNumber.."TimerBar", "SoR_"..Window_Name.."VPDESTROBAR", "SoR_"..Window_Name.."VPORDERBAR" |
| arg2 | Observed as a runtime window or control identifier. | Observed values: 0, DefaultColor.GREEN.r, PartyCast.Settings.Colors.Casting[1] |
| arg3 | Observed as a runtime window or control identifier. | Observed values: 255, DefaultColor.GREEN.g, PartyCast.Settings.Colors.Casting[2] |
| arg4 | Observed as a runtime window or control identifier. | Observed values: 0, DefaultColor.GREEN.b, PartyCast.Settings.Colors.Casting[3] |

## Returns

- Not confidently inferable from contract artifacts alone.

## Side Effects

- Mutates engine-visible UI state or presentation.

## Seen In

- Ace
- ActionBarHide
- AdjustTheTip
- Amethyst
- CMap
- Crusher
- EZCraftX
- EZGuard
- Effigy
- GCDsaver
- Hopper
- InfoScroller
- LibWBToggler
- Map
- Motion
- NaturalLog
- PartyCast
- Pure
- Pure Careerbar
- RealmStatus
- RoR_SoR
- SNT_CASTBAR
- Shinies
- TargetInfoRing
- TargetRing
- TidyRoll
- Tokens
- WarBoard_WarWhisperer
- WarTriage
- Wikki's Cooldown Bar
- Wikki's Cooldown Pulse
- WoH-Reticle
- scenarioInfo
- xHUD
- xPanels

## Examples

- Ace: ForeColor -> StatusBarSetForegroundTint(self.name, red, green, blue)
- ActionBarHide: ForeColor -> StatusBarSetForegroundTint(self.name, red, green, blue)
- AdjustTheTip: AddTargetHealthToMouseOver -> StatusBarSetForegroundTint(c_HEALTH_BAR_CONTAINER.."HealthPercentBarBar", 0, 255, 0)
- Amethyst: ForeColor -> StatusBarSetForegroundTint(self.name, red, green, blue)
- CMap: ForeColor -> StatusBarSetForegroundTint(self.name, red, green, blue)
- Crusher: ForeColor -> StatusBarSetForegroundTint(self.name, red, green, blue)

## Used With

- [StatusBarSetBackgroundTint](window_StatusBarSetBackgroundTint.md) (HIGH 100/100) - Window Function
- [StatusBarSetMaximumValue](window_StatusBarSetMaximumValue.md) (HIGH 100/100) - Window Function
- [WindowAddAnchor](window_WindowAddAnchor.md) (HIGH 100/100) - Window Function
- [WindowGetDimensions](window_WindowGetDimensions.md) (HIGH 100/100) - Window Function

## Notes

- none
