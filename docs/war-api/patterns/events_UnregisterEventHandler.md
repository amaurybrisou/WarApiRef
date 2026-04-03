# UnregisterEventHandler

- Category: events
- Confidence: MEDIUM

## Description

Observed removing previously registered global runtime handlers.

## Involved APIs

- [SystemData.Events.INTERACT_DONE](../events/game_events/game_event_SystemData.Events.INTERACT_DONE.md) (HIGH 100/100) - Game Event
- [SystemData.Events.INTERACT_DONE](../systemdata/fields/systemdata_SystemData.Events.INTERACT_DONE.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_BEGIN_CAST](../events/game_events/game_event_SystemData.Events.PLAYER_BEGIN_CAST.md) (HIGH 100/100) - Game Event
- [SystemData.Events.PLAYER_BEGIN_CAST](../systemdata/fields/systemdata_SystemData.Events.PLAYER_BEGIN_CAST.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_CAST_TIMER_SETBACK](../systemdata/fields/systemdata_SystemData.Events.PLAYER_CAST_TIMER_SETBACK.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_CAST_TIMER_SETBACK](../events/game_events/game_event_SystemData.Events.PLAYER_CAST_TIMER_SETBACK.md) (HIGH 100/100) - Game Event
- [SystemData.Events.PLAYER_DEATH](../events/game_events/game_event_SystemData.Events.PLAYER_DEATH.md) (HIGH 100/100) - Game Event
- [SystemData.Events.PLAYER_DEATH](../systemdata/fields/systemdata_SystemData.Events.PLAYER_DEATH.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_END_CAST](../systemdata/fields/systemdata_SystemData.Events.PLAYER_END_CAST.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_END_CAST](../events/game_events/game_event_SystemData.Events.PLAYER_END_CAST.md) (HIGH 100/100) - Game Event
- [SystemData.Events.PLAYER_START_INTERACT_TIMER](../systemdata/fields/systemdata_SystemData.Events.PLAYER_START_INTERACT_TIMER.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_START_INTERACT_TIMER](../events/game_events/game_event_SystemData.Events.PLAYER_START_INTERACT_TIMER.md) (HIGH 100/100) - Game Event
- [UnregisterEventHandler](../globals/functions/global_UnregisterEventHandler.md) (HIGH 81/100) - Global Function

## Flow Diagram

```text
SystemData.Events.ENTER_WORLD <-> SystemData.Events.INTERACT_DONE
```

## Example Code

```lua
PartyCast: UnregisterEventHandler(SystemData.Events.PLAYER_START_INTERACT_TIMER, "PartyCast.StartInteract")
```

## Evidence

- PartyCast: UnregisterEventHandler(SystemData.Events.PLAYER_START_INTERACT_TIMER, "PartyCast.StartInteract")
- PartyCast: UnregisterEventHandler(SystemData.Events.INTERACT_DONE, "PartyCast.EndCast")
- PartyCast: UnregisterEventHandler(SystemData.Events.PLAYER_BEGIN_CAST, "PartyCast.StartCast")
- PartyCast: UnregisterEventHandler(SystemData.Events.PLAYER_END_CAST, "PartyCast.EndCast")
- PartyCast: UnregisterEventHandler(SystemData.Events.PLAYER_CAST_TIMER_SETBACK, "PartyCast.SetbackCast")
- PartyCast: UnregisterEventHandler(SystemData.Events.PLAYER_DEATH, "PartyCast.ON_DEATH")
