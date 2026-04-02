# BroadcastEvent

- Category: Global Function
- Confidence level: HIGH
- Confidence score: 93/100
- Seen in: 8 addons

## Confidence Assessment

- Level: HIGH

- Score: 93/100

- Rationale: Promoted as HIGH confidence because seen in 4 or more addons, called globally with no local definition, used in event registration or dispatch.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +18 Used in event registration or dispatch: Observed in event-driven engine hooks.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +10 Argument pattern is consistent: Observed argument positions remain stable.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | Effigy, Enemy, FastInteract, GCDsaver, JunkDump, LoyalPet, RVMOD_Manager, Targets |
| Files seen in | `/workspace_addons/Effigy/Effigy.lua:337`, `/workspace_addons/Effigy/Effigy.lua:376`, `/workspace_addons/Enemy/Code/GroupIcons/GroupIcons.lua:200`, `/workspace_addons/Enemy/Code/ScenarioInfo/ScenarioInfo.lua:371`, `/workspace_addons/Enemy/Code/ScenarioInfo/ScenarioInfo.lua:507`, `/workspace_addons/Enemy/Code/ScenarioInfo/ScenarioInfo.lua:519`, `/workspace_addons/FastInteract/FastInteract.lua:20`, `/workspace_addons/GCDsaver/GCDsaver.lua:76` |
| Namespaces detected | BroadcastEvent |
| Source kinds | lua_calls |
| Example locations | Effigy: Effigy.LButtonDown, Effigy: Effigy.local.OnMenuClickLeaveGroup, Effigy: OnMenuClickLeaveGroup, Enemy: Enemy.GroupIcons_GroupIcon_OnLButtonDown, Enemy: Enemy.ScenarioInfoCheckBroadcast, Enemy: Enemy.ScenarioInfoUI_ScenarioInfoDialog_OnHidden |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 16 |
| Global usage count | 16 |
| Local definition count | 0 |
| Documentation references | 0 |
| Initialization flow references | 0 |
| Known engine namespace | no |
| Default UI presence | no |
| Event binding presence | yes |
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
BroadcastEvent(eventId)
```

## Description

Observed triggering a runtime event so existing handlers are notified.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| eventId | Observed as a runtime event identifier dispatched to listeners. | Observed values: SystemData.Events.GROUP_LEAVE, SystemData.Events.INTERACT_ACCEPT_QUEST, SystemData.Events.INTERACT_COMPLETE_QUEST |

## Returns

- Not confidently inferable from addon-api docs alone.

## Side Effects

- Dispatches an event into the runtime event system.

## Seen In

- Effigy
- Enemy
- FastInteract
- GCDsaver
- JunkDump
- LoyalPet
- RVMOD_Manager
- Targets

## Examples

- Effigy: Effigy.LButtonDown -> BroadcastEvent(SystemData.Events.TARGET_PET)
- Effigy: Effigy.local.OnMenuClickLeaveGroup -> BroadcastEvent(SystemData.Events.GROUP_LEAVE)
- Effigy: OnMenuClickLeaveGroup -> BroadcastEvent(SystemData.Events.GROUP_LEAVE)
- Enemy: Enemy.GroupIcons_GroupIcon_OnLButtonDown -> BroadcastEvent(custom_target_event)
- Enemy: Enemy.ScenarioInfoCheckBroadcast -> BroadcastEvent(SystemData.Events.SCENARIO_START_UPDATING_PLAYERS_STATS)
- Enemy: Enemy.ScenarioInfoUI_ScenarioInfoDialog_OnHidden -> BroadcastEvent(SystemData.Events.SCENARIO_STOP_UPDATING_PLAYERS_STATS)

## Related APIs

- [AlertText](../tables/table_AlertText.md) (HIGH 100/100) - Global Table

## Used With

- [GameData.PlayerActions.DO_MACRO](../../gamedata/fields/gamedata_GameData.PlayerActions.DO_MACRO.md) (HIGH 100/100) - GameData Field
- [OnLButtonDown](../../events/window_events/window_event_OnLButtonDown.md) (HIGH 100/100) - Window Event
- [OnLButtonDown](../../xml/handlers/handler_OnLButtonDown.md) (HIGH 100/100) - XML Handler
- [SystemData.ActiveWindow.name](../../systemdata/fields/systemdata_SystemData.ActiveWindow.name.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.GROUP_LEAVE](../../systemdata/fields/systemdata_SystemData.Events.GROUP_LEAVE.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.TARGET_PET](../../systemdata/fields/systemdata_SystemData.Events.TARGET_PET.md) (HIGH 100/100) - SystemData Field
- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type
- [WindowSetGameActionData](../../window_api/functions/window_WindowSetGameActionData.md) (HIGH 100/100) - Window Function

## Triggered By

- [OnHidden](../../xml/handlers/handler_OnHidden.md) (HIGH 100/100) - XML Handler
- [OnHidden](../../events/window_events/window_event_OnHidden.md) (HIGH 100/100) - Window Event
- [OnLButtonDown](../../xml/handlers/handler_OnLButtonDown.md) (HIGH 100/100) - XML Handler
- [OnLButtonDown](../../events/window_events/window_event_OnLButtonDown.md) (HIGH 100/100) - Window Event
- [SystemData.Events.SCENARIO_STOP_UPDATING_PLAYERS_STATS](../../events/game_events/game_event_SystemData.Events.SCENARIO_STOP_UPDATING_PLAYERS_STATS.md) (HIGH 100/100) - Game Event

## Affects

- [GameData.PlayerActions.DO_MACRO](../../gamedata/fields/gamedata_GameData.PlayerActions.DO_MACRO.md) (HIGH 100/100) - GameData Field
- [SystemData.ActiveWindow.name](../../systemdata/fields/systemdata_SystemData.ActiveWindow.name.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.GROUP_LEAVE](../../systemdata/fields/systemdata_SystemData.Events.GROUP_LEAVE.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.SCENARIO_START_UPDATING_PLAYERS_STATS](../../systemdata/fields/systemdata_SystemData.Events.SCENARIO_START_UPDATING_PLAYERS_STATS.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.SCENARIO_STOP_UPDATING_PLAYERS_STATS](../../systemdata/fields/systemdata_SystemData.Events.SCENARIO_STOP_UPDATING_PLAYERS_STATS.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.TARGET_PET](../../systemdata/fields/systemdata_SystemData.Events.TARGET_PET.md) (HIGH 100/100) - SystemData Field
- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Notes

- Canonical entry built from observed call sites, not from engine source or decompiled definitions.
