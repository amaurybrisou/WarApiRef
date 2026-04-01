# GameData.PlayerActions.SET_TARGET

- Category: GameData Field
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 7 addons

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 150

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, seen in 4 or more addons, matches a known engine namespace.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +25 Referenced by generated docs or reference files: The symbol is reinforced outside a single call page.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | AggroMeter, Effigy, Enemy, RetAlert, Targets, WarTriage, followTheLeader |
| Files seen in | `/workspace/AggroMeter/AggroMeter.lua:242`, `/workspace/Effigy/Effigy.lua:365`, `/workspace/Effigy/Elements/EffigyIndicator.lua:132`, `/workspace/Effigy/States/EffigyStateTargets.lua:325`, `/workspace/Enemy/Code/GroupIcons/GroupIcon.lua:115`, `/workspace/Enemy/Code/Guard/Guard.lua:104`, `/workspace/Enemy/Code/Marks/MarkTemplate.lua:85`, `/workspace/Enemy/Code/UnitFrames/UnitFrame.lua:344` |
| Namespaces detected | GameData |
| Source kinds | lua_call |
| Example locations | AggroMeter.SelectChar, Effigy.LButtonUp, Effigy.UpdateTarget, EffigyIndicator:addInteractive, Enemy.Guard_SetGuardPlayerName, Enemy.UnitFramesUI_UnitFrame_OnLButtonDown |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 13 |
| Global usage count | 13 |
| Local definition count | 0 |
| Documentation references | 1 |
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
| Conflicting signatures | no |
| Conflicting roles | no |
| Wrapper likely | no |
| Never outside local graph | no |
| Local helper only | no |

## Description

Observed GameData field used by 7 addons through generated function calls, event pages, or lifecycle evidence.

## Seen In

- AggroMeter
- Effigy
- Enemy
- RetAlert
- Targets
- WarTriage
- followTheLeader

## Related APIs

- [LabelGetTextDimensions](../../window_api/functions/window_LabelGetTextDimensions.md) (HIGH 100/100) - Window Function
- [WindowGetParent](../../window_api/functions/window_WindowGetParent.md) (HIGH 100/100) - Window Function
- [WindowSetGameActionData](../../window_api/functions/window_WindowSetGameActionData.md) (HIGH 100/100) - Window Function
- [WindowSetGameActionTrigger](../../window_api/functions/window_WindowSetGameActionTrigger.md) (HIGH 100/100) - Window Function

## Used With

- [GameData.Player.level](gamedata_GameData.Player.level.md) (HIGH 100/100) - GameData Field
- [Label](../../xml/element_types/element_Label.md) (HIGH 100/100) - XML Element Type
- [OnLButtonUp](../../events/window_events/window_event_OnLButtonUp.md) (HIGH 100/100) - Window Event
- [OnLButtonUp](../../xml/handlers/handler_OnLButtonUp.md) (HIGH 100/100) - XML Handler
- [OnMouseOver](../../events/window_events/window_event_OnMouseOver.md) (HIGH 100/100) - Window Event
- [OnMouseOver](../../xml/handlers/handler_OnMouseOver.md) (HIGH 100/100) - XML Handler
- [SystemData.MouseOverWindow.name](../../systemdata/fields/systemdata_SystemData.MouseOverWindow.name.md) (HIGH 100/100) - SystemData Field
- [WindowGetParent](../../window_api/functions/window_WindowGetParent.md) (HIGH 100/100) - Window Function
- [WindowSetGameActionData](../../window_api/functions/window_WindowSetGameActionData.md) (HIGH 100/100) - Window Function
- [WindowSetGameActionTrigger](../../window_api/functions/window_WindowSetGameActionTrigger.md) (HIGH 100/100) - Window Function
- [WindowSetHandleInput](../../window_api/functions/window_WindowSetHandleInput.md) (HIGH 100/100) - Window Function
- [towstring](../../globals/functions/global_towstring.md) (HIGH 75/100) - Global Function

## Triggered By

- none

## Affects

- none

## Notes

- Observed in contexts: AggroMeter.SelectChar, Effigy.LButtonUp, Effigy.UpdateTarget, EffigyIndicator:addInteractive, Enemy.Guard_SetGuardPlayerName, Enemy.UnitFramesUI_UnitFrame_OnLButtonDown
