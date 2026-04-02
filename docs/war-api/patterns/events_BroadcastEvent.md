# BroadcastEvent

- Category: events
- Confidence: HIGH

## Description

Observed triggering a runtime event so existing handlers are notified.

## Involved APIs

- [SystemData.Events.GROUP_LEAVE](../systemdata/fields/systemdata_SystemData.Events.GROUP_LEAVE.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.GROUP_LEAVE](../events/game_events/game_event_SystemData.Events.GROUP_LEAVE.md) (HIGH 100/100) - Game Event
- [SystemData.Events.SCENARIO_START_UPDATING_PLAYERS_STATS](../systemdata/fields/systemdata_SystemData.Events.SCENARIO_START_UPDATING_PLAYERS_STATS.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.SCENARIO_STOP_UPDATING_PLAYERS_STATS](../events/game_events/game_event_SystemData.Events.SCENARIO_STOP_UPDATING_PLAYERS_STATS.md) (HIGH 100/100) - Game Event
- [SystemData.Events.SCENARIO_STOP_UPDATING_PLAYERS_STATS](../systemdata/fields/systemdata_SystemData.Events.SCENARIO_STOP_UPDATING_PLAYERS_STATS.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.TARGET_PET](../systemdata/fields/systemdata_SystemData.Events.TARGET_PET.md) (HIGH 100/100) - SystemData Field
- [BroadcastEvent](../globals/functions/global_BroadcastEvent.md) (HIGH 93/100) - Global Function

## Flow Diagram

```text
OnLButtonUp <-> OnLButtonUp
```

## Example Code

```lua
Effigy: BroadcastEvent(SystemData.Events.TARGET_PET)
```

## Evidence

- Effigy: BroadcastEvent(SystemData.Events.TARGET_PET)
- Effigy: BroadcastEvent(SystemData.Events.GROUP_LEAVE)
- Effigy: BroadcastEvent(SystemData.Events.GROUP_LEAVE)
- Enemy: BroadcastEvent(custom_target_event)
- Enemy: BroadcastEvent(SystemData.Events.SCENARIO_START_UPDATING_PLAYERS_STATS)
- Enemy: BroadcastEvent(SystemData.Events.SCENARIO_STOP_UPDATING_PLAYERS_STATS)
