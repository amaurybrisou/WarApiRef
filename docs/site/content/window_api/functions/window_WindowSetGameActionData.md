# WindowSetGameActionData

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
| Addons seen in | Ace, AggroMeter, Enemy, LibWBToggler, MiracleGrowLight, PartyCast, Shinies, WoH-Reticle |
| Files seen in | `/workspace/data/raw/Ace/LibGUI.lua:574`, `/workspace/data/raw/AggroMeter/AggroMeter.lua:242`, `/workspace/data/raw/Enemy/Code/GroupIcons/GroupIcon.lua:115`, `/workspace/data/raw/Enemy/Code/Guard/Guard.lua:104`, `/workspace/data/raw/Enemy/Code/Marks/MarkTemplate.lua:85`, `/workspace/data/raw/Enemy/Code/UnitFrames/ClickCasting.lua:147`, `/workspace/data/raw/Enemy/Code/UnitFrames/UnitFrame.lua:314`, `/workspace/data/raw/Enemy/Code/UnitFrames/UnitFrame.lua:344` |
| Namespaces detected | WindowSetGameActionData |
| Source kinds | lua_calls |
| Example locations | Ace: LIBGUI_Button:Action, AggroMeter: AggroMeter.SelectChar, Enemy: Enemy.Guard_SetGuardPlayerName, Enemy: Enemy.UnitFramesUI_UnitFrame_OnLButtonDown, Enemy: EnemyClickCasting:Proceed, Enemy: EnemyGroupIcon:Attach |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 17 |
| Global usage count | 17 |
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
| windowName | Observed as a target window name. | Observed values: "EnemyGuardDistanceIndicator", WINDOW_NAME, frame.windowName |
| arg2 | Observed as a function or method reference. | Observed values: 0, GameData.PlayerActions.ASSIST_PLAYER, GameData.PlayerActions.DO_ABILITY |
| arg3 | Observed as a numeric value. | Observed values: 0, GameData.TradeSkills.CULTIVATION, numArg |
| arg4 | Observed as a runtime window or control identifier. | Observed values: Enemy.isNil(g.guardPlayerName,L ""), L "", frame.playerName |

## Returns

- Not confidently inferable from addon-api docs alone.

## Side Effects

- Mutates engine-visible UI state or presentation.

## Seen In

- Ace
- AggroMeter
- Enemy
- LibWBToggler
- MiracleGrowLight
- PartyCast
- Shinies
- WoH-Reticle
- followTheLeader

## Examples

- Ace: LIBGUI_Button:Action -> WindowSetGameActionData(self.name, actionType, numArg, wstringArg)
- AggroMeter: AggroMeter.SelectChar -> WindowSetGameActionData(tostring(WindowName), GameData.PlayerActions.SET_TARGET, 0, towstring(AggroMeter.AggroHolder[tostring(MobNumber)][tonumber(LabelNumber)].name))
- Enemy: Enemy.Guard_SetGuardPlayerName -> WindowSetGameActionData("EnemyGuardDistanceIndicator", GameData.PlayerActions.SET_TARGET, 0, Enemy.isNil(g.guardPlayerName,L ""))
- Enemy: Enemy.UnitFramesUI_UnitFrame_OnLButtonDown -> WindowSetGameActionData(frame.windowName, GameData.PlayerActions.SET_TARGET, 0, frame.playerName)
- Enemy: EnemyClickCasting:Proceed -> WindowSetGameActionData(self.frame.windowName, GameData.PlayerActions.DO_ABILITY, self.abilityId, self.frame.playerName)
- Enemy: EnemyClickCasting:Proceed -> WindowSetGameActionData(self.frame.windowName, 0, 0, L "")

## Related APIs

- none

## Used With

- none

## Triggered By

- [OnLButtonDown](../../xml/handlers/handler_OnLButtonDown.md) (HIGH 100/100) - XML Event
- [OnLButtonDown](../../events/window_events/window_event_OnLButtonDown.md) (HIGH 100/100) - Window Event
- [OnLButtonUp](../../xml/handlers/handler_OnLButtonUp.md) (HIGH 100/100) - XML Event
- [OnLButtonUp](../../events/window_events/window_event_OnLButtonUp.md) (HIGH 100/100) - Window Event
- [OnMouseOver](../../xml/handlers/handler_OnMouseOver.md) (HIGH 100/100) - XML Event
- [OnMouseOver](../../events/window_events/window_event_OnMouseOver.md) (HIGH 100/100) - Window Event

## Affects

- [GameData.PlayerActions.ASSIST_PLAYER](../../gamedata/fields/gamedata_GameData.PlayerActions.ASSIST_PLAYER.md) (HIGH 100/100) - GameData Field
- [GameData.PlayerActions.DO_ABILITY](../../gamedata/fields/gamedata_GameData.PlayerActions.DO_ABILITY.md) (HIGH 100/100) - GameData Field
- [GameData.PlayerActions.SET_TARGET](../../gamedata/fields/gamedata_GameData.PlayerActions.SET_TARGET.md) (HIGH 100/100) - GameData Field
- [Label](../../xml/element_types/element_Label.md) (HIGH 100/100) - XML Element Type
- [SystemData.MouseOverWindow.name](../../systemdata/fields/systemdata_SystemData.MouseOverWindow.name.md) (HIGH 100/100) - SystemData Field
- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Notes

- none
