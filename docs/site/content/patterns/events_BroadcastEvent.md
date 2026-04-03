# BroadcastEvent

- Category: events
- Confidence: HIGH

## Description

Observed triggering a runtime event so existing handlers are notified.

## Involved APIs

- [SystemData.Events](../systemdata/fields/systemdata_SystemData.Events.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.EXIT_GAME](../systemdata/fields/systemdata_SystemData.Events.EXIT_GAME.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.SCENARIO_START_UPDATING_PLAYERS_STATS](../systemdata/fields/systemdata_SystemData.Events.SCENARIO_START_UPDATING_PLAYERS_STATS.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.SCENARIO_STOP_UPDATING_PLAYERS_STATS](../systemdata/fields/systemdata_SystemData.Events.SCENARIO_STOP_UPDATING_PLAYERS_STATS.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.SEND_CHAT_TEXT](../systemdata/fields/systemdata_SystemData.Events.SEND_CHAT_TEXT.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.USER_SETTINGS_CHANGED](../systemdata/fields/systemdata_SystemData.Events.USER_SETTINGS_CHANGED.md) (HIGH 100/100) - SystemData Field
- [BroadcastEvent](../globals/functions/global_BroadcastEvent.md) (HIGH 93/100) - Global Function

## Flow Diagram

```text
OnUpdate -> BroadcastEvent
```

## Example Code

```lua
Assist: BroadcastEvent(SystemData.Events.SCENARIO_START_UPDATING_PLAYERS_STATS)
```

## Evidence

- Assist: BroadcastEvent(SystemData.Events.SCENARIO_START_UPDATING_PLAYERS_STATS)
- Assist: BroadcastEvent(SystemData.Events.SCENARIO_STOP_UPDATING_PLAYERS_STATS)
- AuctionStats: BroadcastEvent(SystemData.Events.EXIT_GAME)
- AutoFocus: BroadcastEvent(SystemData.Events.SEND_CHAT_TEXT)
- Brizio's Crappy Computer Medic: BroadcastEvent(SystemData.Events.USER_SETTINGS_CHANGED)
- Brizio's Crappy Computer Medic: BroadcastEvent(SystemData.Events.USER_SETTINGS_CHANGED)
