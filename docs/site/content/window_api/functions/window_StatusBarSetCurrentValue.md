# StatusBarSetCurrentValue

- Category: Window Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 38 addons

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
| Addons seen in | Ace, ActionBarHide, AdjustTheTip, Amethyst, BlackBox, CleanUnitFrames, Crusher, EA_ScenarioGroupWindow |
| Files seen in | AdjustTheTip.lua, BlackBox.lua, CleanGroupMemberUnitFrame.lua, CleanTargetUnitFrame.lua, FlagCap.lua, LibGUI.lua, LibGui.lua, Libraries/LibGUI.lua |
| Namespaces detected | StatusBarSetCurrentValue |
| Source kinds | lua_calls |
| Example locations | Ace: SetValue, ActionBarHide: SetValue, AdjustTheTip: SetMouseOverTargetHealth, Amethyst: SetValue, BlackBox: UpdateTimer, CleanUnitFrames: UpdateActionPoints |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 51 |
| Global usage count | 51 |
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
StatusBarSetCurrentValue(arg1, arg2)
```

## Description

Observed as a window function across 38 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a function or method reference. | Observed values: "PartyCastWindow"..PlayerNumber.."TimerBar", "RespawnTimerWindowBar", "SoR_"..Window_Name.."VPDESTROBAR" |
| arg2 | Observed as a runtime window or control identifier. | Observed values: 0, 1, controlFill |

## Returns

- Not confidently inferable from contract artifacts alone.

## Side Effects

- Mutates engine-visible UI state or presentation.

## Seen In

- Ace
- ActionBarHide
- AdjustTheTip
- Amethyst
- BlackBox
- CleanUnitFrames
- Crusher
- EA_ScenarioGroupWindow
- EZCraftX
- EZGuard
- Effigy
- FlagCap
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
- Targets
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

- Ace: SetValue -> StatusBarSetCurrentValue(self.name, value)
- ActionBarHide: SetValue -> StatusBarSetCurrentValue(self.name, value)
- AdjustTheTip: SetMouseOverTargetHealth -> StatusBarSetCurrentValue(c_HEALTH_BAR_CONTAINER.."HealthPercentBarBar", val)
- Amethyst: SetValue -> StatusBarSetCurrentValue(self.name, value)
- BlackBox: UpdateTimer -> StatusBarSetCurrentValue("RespawnTimerWindowBar", respawnTimeLeft)
- CleanUnitFrames: UpdateActionPoints -> StatusBarSetCurrentValue(self:GetName().."StatusContainerAPPercentBar", newActionPoints)

## Notes

- none
