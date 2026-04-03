# WindowSetGameActionData

- Category: Window Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 53 addons

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
| Addons seen in | Ace, AggroMeter, Amethyst, Assist, BuddyBind, CMap, CastSequence, CleanUnitFrames |
| Files seen in | AggroMeter.lua, BuddyBind.lua, CleanGroupMemberUnitFrame.lua, Code/GroupIcons/GroupIcon.lua, Code/Guard/Guard.lua, Code/Marks/MarkTemplate.lua, Code/UnitFrames/ClickCasting.lua, Code/UnitFrames/UnitFrame.lua |
| Namespaces detected | WindowSetGameActionData |
| Source kinds | lua_calls |
| Example locations | Ace: Action, AggroMeter: SelectChar, Amethyst: Action, Assist: updTargets, BuddyBind: GrabName, BuddyBind: init |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 109 |
| Global usage count | 109 |
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
WindowSetGameActionData(windowName, arg2, arg3, arg4)
```

## Description

Observed mutating runtime window state or presentation.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| windowName | Observed as a target window name. | Observed values: "AssistWindow1BTN", "AssistWindow2BTN", "AssistWindow3BTN" |
| arg2 | Observed as a function or method reference. | Observed values: 0, 4, GameData.PlayerActions.ASSIST_PLAYER |
| arg3 | Observed as a runtime window or control identifier. | Observed values: 0, 46, 47 |
| arg4 | Observed as a text or wstring payload. | Observed values: EffigyState:GetState(bar.state):GetTitle(), Enemy.isNil(g.guardPlayerName,L ""), GameData.Player.name |

## Returns

- Not confidently inferable from contract artifacts alone.

## Side Effects

- Mutates engine-visible UI state or presentation.

## Seen In

- Ace
- AggroMeter
- Amethyst
- Assist
- BuddyBind
- CMap
- CastSequence
- CleanUnitFrames
- Crusher
- DetauntHelper
- EA_ScenarioGroupWindow
- EZCraftX
- EZGuard
- Effigy
- Enemy
- EveryBodyGuard
- GCDsaver
- GroupSpotter
- HealGrid
- HealHoverAssist
- Hopper
- InfoScroller
- KeyBar
- LibWBToggler
- Map
- MarkBuff
- Miracle Grow Remix
- MiracleGrow
- MiracleGrowLight
- Motion
- NaturalLog
- PartyCast
- Pure
- Pure Careerbar
- RealmStatus
- Refer
- ResHelp
- Shinies
- Squared
- SquaredClick
- TargetRing
- Targets
- TastyButtons
- Tokens
- Twister
- WarBoard_WarWhisperer
- WarTriage
- Wikki's Cooldown Bar
- Wikki's Cooldown Pulse
- WoH-Reticle
- followTheLeader
- scenarioInfo
- xHUD

## Examples

- Ace: Action -> WindowSetGameActionData(self.name, actionType, numArg, wstringArg)
- AggroMeter: SelectChar -> WindowSetGameActionData(tostring(WindowName), GameData.PlayerActions.SET_TARGET, 0, towstring(AggroMeter.AggroHolder[tostring(MobNumber)][tonumber(LabelNumber)].name))
- Amethyst: Action -> WindowSetGameActionData(self.name, actionType, numArg, wstringArg)
- Assist: updTargets -> WindowSetGameActionData("AssistWindow1BTN", 4, 46, L "")
- Assist: updTargets -> WindowSetGameActionData("AssistWindow2BTN", 4, 47, L "")
- Assist: updTargets -> WindowSetGameActionData("AssistWindow3BTN", 4, 48, L "")

## Related APIs

- [OnLButtonDown](../../xml/handlers/handler_OnLButtonDown.md) (HIGH 88/100) - XML Event
- [OnUpdate](../../xml/handlers/handler_OnUpdate.md) (HIGH 88/100) - XML Event

## Used With

- [GameData.PlayerActions.SET_TARGET](../../gamedata/fields/gamedata_GameData.PlayerActions.SET_TARGET.md) (HIGH 100/100) - GameData Field

## Affects

- [GameData.PlayerActions.SET_TARGET](../../gamedata/fields/gamedata_GameData.PlayerActions.SET_TARGET.md) (HIGH 100/100) - GameData Field
- [SystemData.MouseOverWindow.name](../../systemdata/fields/systemdata_SystemData.MouseOverWindow.name.md) (HIGH 100/100) - SystemData Field

## Notes

- none
