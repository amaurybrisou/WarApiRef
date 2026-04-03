# BroadcastEvent

- Category: events
- Confidence: MEDIUM

## Description

Observed triggering a runtime event so existing handlers are notified.

## Involved APIs

- [SystemData.Events.SCENARIO_START_UPDATING_PLAYERS_STATS](../systemdata/fields/systemdata_SystemData.Events.SCENARIO_START_UPDATING_PLAYERS_STATS.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.SCENARIO_STOP_UPDATING_PLAYERS_STATS](../events/game_events/game_event_SystemData.Events.SCENARIO_STOP_UPDATING_PLAYERS_STATS.md) (HIGH 100/100) - Game Event
- [SystemData.Events.SCENARIO_STOP_UPDATING_PLAYERS_STATS](../systemdata/fields/systemdata_SystemData.Events.SCENARIO_STOP_UPDATING_PLAYERS_STATS.md) (HIGH 100/100) - SystemData Field

## Flow Diagram

```text
OnLButtonUp
  -> ui: Button, ColorPicker, DynamicImage, Label, ListBox, MapDisplay, Window
```

## Example Code

```lua
Enemy: BroadcastEvent(custom_target_event)
```

## Evidence

- Enemy: BroadcastEvent(custom_target_event)
- Enemy: BroadcastEvent(SystemData.Events.SCENARIO_START_UPDATING_PLAYERS_STATS)
- Enemy: BroadcastEvent(SystemData.Events.SCENARIO_STOP_UPDATING_PLAYERS_STATS)
- Enemy: BroadcastEvent(SystemData.Events.SCENARIO_START_UPDATING_PLAYERS_STATS)
