# RegisterEventHandler

- Category: events
- Confidence: HIGH

## Description

Observed registering global runtime handlers against shared event identifiers.

## Involved APIs

- [SystemData.Events](../systemdata/fields/systemdata_SystemData.Events.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.ENTER_WORLD](../systemdata/fields/systemdata_SystemData.Events.ENTER_WORLD.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.LOADING_END](../systemdata/fields/systemdata_SystemData.Events.LOADING_END.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.LOADING_END](../events/game_events/game_event_SystemData.Events.LOADING_END.md) (HIGH 100/100) - Game Event
- [SystemData.Events.PLAYER_COMBAT_FLAG_UPDATED](../systemdata/fields/systemdata_SystemData.Events.PLAYER_COMBAT_FLAG_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.RELOAD_INTERFACE](../systemdata/fields/systemdata_SystemData.Events.RELOAD_INTERFACE.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.RELOAD_INTERFACE](../events/game_events/game_event_SystemData.Events.RELOAD_INTERFACE.md) (HIGH 100/100) - Game Event
- [SystemData.Events.WORLD_OBJ_COMBAT_EVENT](../events/game_events/game_event_SystemData.Events.WORLD_OBJ_COMBAT_EVENT.md) (HIGH 100/100) - Game Event
- [SystemData.Events.WORLD_OBJ_COMBAT_EVENT](../systemdata/fields/systemdata_SystemData.Events.WORLD_OBJ_COMBAT_EVENT.md) (HIGH 100/100) - SystemData Field
- [Window](../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type
- [RegisterEventHandler](../globals/functions/global_RegisterEventHandler.md) (MEDIUM 68/100) - Global Function
- [EventHandler](../xml/element_types/element_EventHandler.md) (MEDIUM 45/100) - XML Element Type

## Flow Diagram

```text
SystemData.Events.LOADING_END <-> SystemData.Events.RELOAD_INTERFACE
```

## Example Code

```lua
AbilityAlert: RegisterEventHandler(SystemData.Events.WORLD_OBJ_COMBAT_EVENT, "AbilityAlert.CombatEvent")
```

## Evidence

- AbilityAlert: RegisterEventHandler(SystemData.Events.WORLD_OBJ_COMBAT_EVENT, "AbilityAlert.CombatEvent")
- ActionBarHide: RegisterEventHandler(SystemData.Events.LOADING_END, "ActionBarHide.OnLoad")
- ActionBarHide: RegisterEventHandler(SystemData.Events.RELOAD_INTERFACE, "ActionBarHide.OnLoad")
- ActionBarHide: RegisterEventHandler(SystemData.Events.PLAYER_COMBAT_FLAG_UPDATED, "ActionBarHide.Combat")
- ActionFraction: RegisterEventHandler(SystemData.Events.LOADING_END, "ActionFractionWindow.UpdateVisibility")
- ActionFraction: RegisterEventHandler(SystemData.Events.ENTER_WORLD, "ActionFraction.Initialize")
