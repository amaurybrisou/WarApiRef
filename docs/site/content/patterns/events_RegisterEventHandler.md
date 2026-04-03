# RegisterEventHandler

- Category: events
- Confidence: HIGH

## Description

Observed registering global runtime handlers against shared event identifiers.

## Involved APIs

- [SystemData.Events.INTERACT_DONE](../systemdata/fields/systemdata_SystemData.Events.INTERACT_DONE.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.INTERACT_DONE](../events/game_events/game_event_SystemData.Events.INTERACT_DONE.md) (HIGH 100/100) - Game Event
- [SystemData.Events.LOADING_END](../systemdata/fields/systemdata_SystemData.Events.LOADING_END.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.LOADING_END](../events/game_events/game_event_SystemData.Events.LOADING_END.md) (HIGH 100/100) - Game Event
- [SystemData.Events.PLAYER_BEGIN_CAST](../systemdata/fields/systemdata_SystemData.Events.PLAYER_BEGIN_CAST.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_BEGIN_CAST](../events/game_events/game_event_SystemData.Events.PLAYER_BEGIN_CAST.md) (HIGH 100/100) - Game Event
- [SystemData.Events.PLAYER_END_CAST](../events/game_events/game_event_SystemData.Events.PLAYER_END_CAST.md) (HIGH 100/100) - Game Event
- [SystemData.Events.PLAYER_END_CAST](../systemdata/fields/systemdata_SystemData.Events.PLAYER_END_CAST.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_START_INTERACT_TIMER](../events/game_events/game_event_SystemData.Events.PLAYER_START_INTERACT_TIMER.md) (HIGH 100/100) - Game Event
- [SystemData.Events.PLAYER_START_INTERACT_TIMER](../systemdata/fields/systemdata_SystemData.Events.PLAYER_START_INTERACT_TIMER.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.RELOAD_INTERFACE](../events/game_events/game_event_SystemData.Events.RELOAD_INTERFACE.md) (HIGH 100/100) - Game Event
- [SystemData.Events.RELOAD_INTERFACE](../systemdata/fields/systemdata_SystemData.Events.RELOAD_INTERFACE.md) (HIGH 100/100) - SystemData Field
- [RegisterEventHandler](../globals/functions/global_RegisterEventHandler.md) (HIGH 93/100) - Global Function

## Flow Diagram

```text
SystemData.Events.ENTER_WORLD <-> SystemData.Events.INTERACT_DONE
```

## Example Code

```lua
Lib RuString: RegisterEventHandler(SystemData.Events.LOADING_END, "LibRuString.OnLoad")
```

## Evidence

- Lib RuString: RegisterEventHandler(SystemData.Events.LOADING_END, "LibRuString.OnLoad")
- Lib RuString: RegisterEventHandler(SystemData.Events.RELOAD_INTERFACE, "LibRuString.OnLoad")
- PartyCast: RegisterEventHandler(SystemData.Events.PLAYER_START_INTERACT_TIMER, "PartyCast.StartInteract")
- PartyCast: RegisterEventHandler(SystemData.Events.INTERACT_DONE, "PartyCast.EndCast")
- PartyCast: RegisterEventHandler(SystemData.Events.PLAYER_BEGIN_CAST, "PartyCast.StartCast")
- PartyCast: RegisterEventHandler(SystemData.Events.PLAYER_END_CAST, "PartyCast.EndCast")
