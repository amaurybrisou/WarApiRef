# Game event fan-in

- Category: events
- Confidence: HIGH

## Description

Observed multiple addons converging on a small set of runtime events such as LOADING_END, RELOAD_INTERFACE, and player update events.

## Involved APIs

- [SystemData.Events.INTERACT_LOOT_ROLL_FIRST_ITEM](../systemdata/fields/systemdata_SystemData.Events.INTERACT_LOOT_ROLL_FIRST_ITEM.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.INTERACT_LOOT_ROLL_FIRST_ITEM](../events/game_events/game_event_SystemData.Events.INTERACT_LOOT_ROLL_FIRST_ITEM.md) (HIGH 100/100) - Game Event
- [SystemData.Events.INTERACT_SHOW_LOOT_ROLL_DATA](../events/game_events/game_event_SystemData.Events.INTERACT_SHOW_LOOT_ROLL_DATA.md) (HIGH 100/100) - Game Event
- [SystemData.Events.INTERACT_SHOW_LOOT_ROLL_DATA](../systemdata/fields/systemdata_SystemData.Events.INTERACT_SHOW_LOOT_ROLL_DATA.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.LOADING_END](../events/game_events/game_event_SystemData.Events.LOADING_END.md) (HIGH 100/100) - Game Event
- [SystemData.Events.LOADING_END](../systemdata/fields/systemdata_SystemData.Events.LOADING_END.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.L_BUTTON_UP_PROCESSED](../systemdata/fields/systemdata_SystemData.Events.L_BUTTON_UP_PROCESSED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.L_BUTTON_UP_PROCESSED](../events/game_events/game_event_SystemData.Events.L_BUTTON_UP_PROCESSED.md) (HIGH 100/100) - Game Event
- [SystemData.Events.PLAYER_TARGET_UPDATED](../events/game_events/game_event_SystemData.Events.PLAYER_TARGET_UPDATED.md) (HIGH 100/100) - Game Event
- [SystemData.Events.PLAYER_TARGET_UPDATED](../systemdata/fields/systemdata_SystemData.Events.PLAYER_TARGET_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.RELOAD_INTERFACE](../systemdata/fields/systemdata_SystemData.Events.RELOAD_INTERFACE.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.RELOAD_INTERFACE](../events/game_events/game_event_SystemData.Events.RELOAD_INTERFACE.md) (HIGH 100/100) - Game Event
- [TextLogGetUpdateEventId](../events/game_events/game_event_TextLogGetUpdateEventId.md) (MEDIUM 63/100) - Game Event

## Flow Diagram

```text
Button <-> OnLButtonUp
```

## Example Code

```lua
SystemData.Events.INTERACT_LOOT_ROLL_FIRST_ITEM (HIGH)
```

## Evidence

- SystemData.Events.INTERACT_LOOT_ROLL_FIRST_ITEM (HIGH)
- SystemData.Events.INTERACT_SHOW_LOOT_ROLL_DATA (HIGH)
- SystemData.Events.LOADING_END (HIGH)
- SystemData.Events.L_BUTTON_UP_PROCESSED (HIGH)
- SystemData.Events.PLAYER_TARGET_UPDATED (HIGH)
- SystemData.Events.RELOAD_INTERFACE (HIGH)
- TextLogGetUpdateEventId (MEDIUM)
