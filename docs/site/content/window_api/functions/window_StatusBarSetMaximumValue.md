# StatusBarSetMaximumValue

- Category: Window Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 37 addons

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
| Addons seen in | Ace, ActionBarHide, AdjustTheTip, Amethyst, BlackBox, CleanUnitFrames, Crusher, EZCraftX |
| Files seen in | AdjustTheTip.lua, BlackBox.lua, CleanGroupMemberUnitFrame.lua, CleanTargetUnitFrame.lua, FlagCap.lua, LibGUI.lua, LibGui.lua, Libraries/LibGUI.lua |
| Namespaces detected | StatusBarSetMaximumValue |
| Source kinds | lua_calls |
| Example locations | Ace: New, ActionBarHide: New, AdjustTheTip: AddTargetHealthToMouseOver, Amethyst: New, BlackBox: OnApplicationTwoButtonDialogHook, CleanUnitFrames: Create |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 46 |
| Global usage count | 46 |
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
StatusBarSetMaximumValue(arg1, arg2)
```

## Description

Observed as a window function across 37 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a function or method reference. | Observed values: "PartyCastWindow"..PlayerNumber.."TimerBar", "RespawnTimerWindowBar", "SoR_"..Window_Name.."VPDESTROBAR" |
| arg2 | Observed as a numeric value. | Observed values: -1, 1, 100 |

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

- Ace: New -> StatusBarSetMaximumValue(w.name, 1)
- ActionBarHide: New -> StatusBarSetMaximumValue(w.name, 1)
- AdjustTheTip: AddTargetHealthToMouseOver -> StatusBarSetMaximumValue(c_HEALTH_BAR_CONTAINER.."HealthPercentBarBar", 100)
- Amethyst: New -> StatusBarSetMaximumValue(w.name, 1)
- BlackBox: OnApplicationTwoButtonDialogHook -> StatusBarSetMaximumValue("RespawnTimerWindowBar", respawnTimeLeft)
- CleanUnitFrames: Create -> StatusBarSetMaximumValue(newUnitFrame:GetName().."StatusContainerHealthPercentBar", 100)

## Related APIs

- [OnInitialize](../../xml/handlers/handler_OnInitialize.md) (HIGH 88/100) - XML Event

## Used With

- [StatusBarSetBackgroundTint](window_StatusBarSetBackgroundTint.md) (HIGH 100/100) - Window Function
- [StatusBarSetForegroundTint](window_StatusBarSetForegroundTint.md) (HIGH 100/100) - Window Function
- [WindowAddAnchor](window_WindowAddAnchor.md) (HIGH 100/100) - Window Function
- [WindowGetDimensions](window_WindowGetDimensions.md) (HIGH 100/100) - Window Function
- [CreateWindowFromTemplate](../../globals/functions/global_CreateWindowFromTemplate.md) (HIGH 75/100) - Global Function

## Affects

- [GameData.BuffTargetType.GROUP_MEMBER_START](../../gamedata/fields/gamedata_GameData.BuffTargetType.GROUP_MEMBER_START.md) (HIGH 100/100) - GameData Field

## Notes

- none
