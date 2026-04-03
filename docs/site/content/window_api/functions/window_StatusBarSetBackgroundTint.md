# StatusBarSetBackgroundTint

- Category: Window Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 31 addons

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
| Addons seen in | Ace, ActionBarHide, AdjustTheTip, Amethyst, Crusher, EZCraftX, EZGuard, Effigy |
| Files seen in | AdjustTheTip.lua, LibGUI.lua, LibGui.lua, Libraries/LibGUI.lua, Libs/LibGUI.lua, PartyCast.lua, TidyRollFrame.lua, libs/LibGUI.lua |
| Namespaces detected | StatusBarSetBackgroundTint |
| Source kinds | lua_calls |
| Example locations | Ace: BackColor, ActionBarHide: BackColor, AdjustTheTip: AddTargetHealthToMouseOver, Amethyst: BackColor, Crusher: BackColor, EZCraftX: BackColor |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 36 |
| Global usage count | 36 |
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
StatusBarSetBackgroundTint(arg1, arg2, arg3, arg4)
```

## Description

Observed as a window function across 31 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a function or method reference. | Observed values: "PartyCastWindow"..PlayerNumber.."TimerBar", c_HEALTH_BAR_CONTAINER.."HealthPercentBarBar", info.."StatusBar" |
| arg2 | Observed as a runtime window or control identifier. | Observed values: 0, 128, DefaultColor.BLACK.r |
| arg3 | Observed as a runtime window or control identifier. | Observed values: 0, 128, DefaultColor.BLACK.g |
| arg4 | Observed as a runtime window or control identifier. | Observed values: 0, 128, DefaultColor.BLACK.b |

## Returns

- Not confidently inferable from contract artifacts alone.

## Side Effects

- Mutates engine-visible UI state or presentation.

## Seen In

- Ace
- ActionBarHide
- AdjustTheTip
- Amethyst
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
- Shinies
- TargetInfoRing
- TargetRing
- TidyRoll
- Tokens
- WarTriage
- Wikki's Cooldown Bar
- Wikki's Cooldown Pulse
- WoH-Reticle
- scenarioInfo
- xHUD
- xPanels

## Examples

- Ace: BackColor -> StatusBarSetBackgroundTint(self.name, red, green, blue)
- ActionBarHide: BackColor -> StatusBarSetBackgroundTint(self.name, red, green, blue)
- AdjustTheTip: AddTargetHealthToMouseOver -> StatusBarSetBackgroundTint(c_HEALTH_BAR_CONTAINER.."HealthPercentBarBar", 128, 128, 128)
- Amethyst: BackColor -> StatusBarSetBackgroundTint(self.name, red, green, blue)
- Crusher: BackColor -> StatusBarSetBackgroundTint(self.name, red, green, blue)
- EZCraftX: BackColor -> StatusBarSetBackgroundTint(self.name, red, green, blue)

## Used With

- [StatusBarSetForegroundTint](window_StatusBarSetForegroundTint.md) (HIGH 100/100) - Window Function
- [StatusBarSetMaximumValue](window_StatusBarSetMaximumValue.md) (HIGH 100/100) - Window Function
- [WindowAddAnchor](window_WindowAddAnchor.md) (HIGH 100/100) - Window Function
- [WindowGetDimensions](window_WindowGetDimensions.md) (HIGH 100/100) - Window Function

## Notes

- none
