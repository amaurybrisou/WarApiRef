# UnregisterEventHandler

- Category: events
- Confidence: MEDIUM

## Description

Observed removing previously registered global runtime handlers.

## Involved APIs

- [SystemData.Events.INTERACT_SHOW_LOOT_ROLL_DATA](../events/game_events/game_event_SystemData.Events.INTERACT_SHOW_LOOT_ROLL_DATA.md) (HIGH 100/100) - Game Event
- [SystemData.Events.INTERACT_SHOW_LOOT_ROLL_DATA](../systemdata/fields/systemdata_SystemData.Events.INTERACT_SHOW_LOOT_ROLL_DATA.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.LOADING_END](../events/game_events/game_event_SystemData.Events.LOADING_END.md) (HIGH 100/100) - Game Event
- [SystemData.Events.LOADING_END](../systemdata/fields/systemdata_SystemData.Events.LOADING_END.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.RELOAD_INTERFACE](../systemdata/fields/systemdata_SystemData.Events.RELOAD_INTERFACE.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.RELOAD_INTERFACE](../events/game_events/game_event_SystemData.Events.RELOAD_INTERFACE.md) (HIGH 100/100) - Game Event

## Flow Diagram

```text
OnLButtonUp <-> OnLButtonUp
```

## Example Code

```lua
TidyRoll: UnregisterEventHandler(SystemData.Events.LOADING_END, "TidyRoll.OnLoad")
```

## Evidence

- TidyRoll: UnregisterEventHandler(SystemData.Events.LOADING_END, "TidyRoll.OnLoad")
- TidyRoll: UnregisterEventHandler(SystemData.Events.RELOAD_INTERFACE, "TidyRoll.OnLoad")
- TidyRoll: UnregisterEventHandler(SystemData.Events.INTERACT_SHOW_LOOT_ROLL_DATA, "TidyRoll.UpdateLootRollData")
