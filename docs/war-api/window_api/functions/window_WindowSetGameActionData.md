# WindowSetGameActionData

- Category: Window Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 13 addons

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
| Addons seen in | Ace, AggroMeter, Effigy, Enemy, GCDsaver, LibWBToggler, Miracle Grow Remix, MiracleGrow |
| Files seen in | `/workspace_addons/Ace/LibGUI.lua:574`, `/workspace_addons/AggroMeter/AggroMeter.lua:242`, `/workspace_addons/Effigy/Effigy.lua:337`, `/workspace_addons/Effigy/Effigy.lua:365`, `/workspace_addons/Effigy/Elements/EffigyIndicator.lua:132`, `/workspace_addons/Effigy/LibGUI.lua:573`, `/workspace_addons/Effigy/States/EffigyStateTargets.lua:325`, `/workspace_addons/Enemy/Code/GroupIcons/GroupIcon.lua:115` |
| Namespaces detected | WindowSetGameActionData |
| Source kinds | lua_calls |
| Example locations | Ace: LIBGUI_Button:Action, AggroMeter: AggroMeter.SelectChar, Effigy: Effigy.LButtonDown, Effigy: Effigy.LButtonUp, Effigy: Effigy.UpdateTarget, Effigy: EffigyIndicator:addInteractive |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 34 |
| Global usage count | 34 |
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
| windowName | Observed as a target window name. | Observed values: "EffigyTLMOAction", "EnemyGuardDistanceIndicator", WINDOW_NAME |
| arg2 | Observed as a function or method reference. | Observed values: 0, GameData.PlayerActions.ASSIST_PLAYER, GameData.PlayerActions.DO_ABILITY |
| arg3 | Observed as a numeric value. | Observed values: 0, GameData.TradeSkills.CULTIVATION, numArg |
| arg4 | Observed as a text or wstring payload. | Observed values: EffigyState:GetState(bar.state):GetTitle(), Enemy.isNil(g.guardPlayerName,L ""), GameData.Player.name |

## Returns

- Not confidently inferable from addon-api docs alone.

## Side Effects

- Mutates engine-visible UI state or presentation.

## Seen In

- Ace
- AggroMeter
- Effigy
- Enemy
- GCDsaver
- LibWBToggler
- Miracle Grow Remix
- MiracleGrow
- MiracleGrowLight
- Shinies
- Targets
- WoH-Reticle
- followTheLeader

## Examples

- Ace: LIBGUI_Button:Action -> WindowSetGameActionData(self.name, actionType, numArg, wstringArg)
- AggroMeter: AggroMeter.SelectChar -> WindowSetGameActionData(tostring(WindowName), GameData.PlayerActions.SET_TARGET, 0, towstring(AggroMeter.AggroHolder[tostring(MobNumber)][tonumber(LabelNumber)].name))
- Effigy: Effigy.LButtonDown -> WindowSetGameActionData(barname, GameData.PlayerActions.DO_MACRO, unitframeMacroId, L "")
- Effigy: Effigy.LButtonUp -> WindowSetGameActionData(barname, GameData.PlayerActions.SET_TARGET, 0, EffigyState:GetState(bar.state):GetTitle())
- Effigy: Effigy.UpdateTarget -> WindowSetGameActionData("EffigyTLMOAction", GameData.PlayerActions.SET_TARGET, 0, unitName)
- Effigy: EffigyIndicator:addInteractive -> WindowSetGameActionData(self.name, GameData.PlayerActions.SET_TARGET, 0, GameData.Player.name)

## Related APIs

- none

## Used With

- [GameData.PlayerActions.SET_TARGET](../../gamedata/fields/gamedata_GameData.PlayerActions.SET_TARGET.md) (HIGH 100/100) - GameData Field

## Triggered By

- [OnLButtonDown](../../xml/handlers/handler_OnLButtonDown.md) (HIGH 100/100) - XML Handler
- [OnLButtonDown](../../events/window_events/window_event_OnLButtonDown.md) (HIGH 100/100) - Window Event
- [OnLButtonUp](../../xml/handlers/handler_OnLButtonUp.md) (HIGH 100/100) - XML Handler
- [OnLButtonUp](../../events/window_events/window_event_OnLButtonUp.md) (HIGH 100/100) - Window Event
- [OnMouseOver](../../xml/handlers/handler_OnMouseOver.md) (HIGH 100/100) - XML Handler
- [OnMouseOver](../../events/window_events/window_event_OnMouseOver.md) (HIGH 100/100) - Window Event

## Affects

- [GameData.Player.level](../../gamedata/fields/gamedata_GameData.Player.level.md) (HIGH 100/100) - GameData Field
- [GameData.Player.name](../../gamedata/fields/gamedata_GameData.Player.name.md) (HIGH 100/100) - GameData Field
- [GameData.PlayerActions.DO_MACRO](../../gamedata/fields/gamedata_GameData.PlayerActions.DO_MACRO.md) (HIGH 100/100) - GameData Field
- [GameData.PlayerActions.SET_TARGET](../../gamedata/fields/gamedata_GameData.PlayerActions.SET_TARGET.md) (HIGH 100/100) - GameData Field
- [Label](../../xml/element_types/element_Label.md) (HIGH 100/100) - XML Element Type
- [SystemData.ActiveWindow.name](../../systemdata/fields/systemdata_SystemData.ActiveWindow.name.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.TARGET_PET](../../systemdata/fields/systemdata_SystemData.Events.TARGET_PET.md) (HIGH 100/100) - SystemData Field
- [SystemData.MouseOverWindow.name](../../systemdata/fields/systemdata_SystemData.MouseOverWindow.name.md) (HIGH 100/100) - SystemData Field
- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Notes

- none
