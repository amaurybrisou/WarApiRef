# Game event fan-in

- Category: events
- Confidence: HIGH

## Description

Observed multiple addons converging on a small set of runtime events such as LOADING_END, RELOAD_INTERFACE, and player update events.

## Involved APIs

- [SystemData.Events.ENTER_WORLD](../systemdata/fields/systemdata_SystemData.Events.ENTER_WORLD.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.ENTER_WORLD](../events/game_events/game_event_SystemData.Events.ENTER_WORLD.md) (HIGH 100/100) - Game Event
- [SystemData.Events.EXIT_GAME](../events/game_events/game_event_SystemData.Events.EXIT_GAME.md) (HIGH 100/100) - Game Event
- [SystemData.Events.EXIT_GAME](../systemdata/fields/systemdata_SystemData.Events.EXIT_GAME.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.INTERACT_COMPLETE_QUEST](../events/game_events/game_event_SystemData.Events.INTERACT_COMPLETE_QUEST.md) (HIGH 100/100) - Game Event
- [SystemData.Events.INTERACT_COMPLETE_QUEST](../systemdata/fields/systemdata_SystemData.Events.INTERACT_COMPLETE_QUEST.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.INTERACT_DEFAULT](../events/game_events/game_event_SystemData.Events.INTERACT_DEFAULT.md) (HIGH 100/100) - Game Event
- [SystemData.Events.INTERACT_DEFAULT](../systemdata/fields/systemdata_SystemData.Events.INTERACT_DEFAULT.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.INTERACT_DONE](../systemdata/fields/systemdata_SystemData.Events.INTERACT_DONE.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.INTERACT_DONE](../events/game_events/game_event_SystemData.Events.INTERACT_DONE.md) (HIGH 100/100) - Game Event
- [SystemData.Events.INTERACT_LOOT_ROLL_FIRST_ITEM](../events/game_events/game_event_SystemData.Events.INTERACT_LOOT_ROLL_FIRST_ITEM.md) (HIGH 100/100) - Game Event
- [SystemData.Events.INTERACT_LOOT_ROLL_FIRST_ITEM](../systemdata/fields/systemdata_SystemData.Events.INTERACT_LOOT_ROLL_FIRST_ITEM.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.INTERACT_SHOW_INFLUENCE_REWARDS](../events/game_events/game_event_SystemData.Events.INTERACT_SHOW_INFLUENCE_REWARDS.md) (HIGH 100/100) - Game Event
- [SystemData.Events.INTERACT_SHOW_INFLUENCE_REWARDS](../systemdata/fields/systemdata_SystemData.Events.INTERACT_SHOW_INFLUENCE_REWARDS.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.INTERACT_SHOW_LOOT](../events/game_events/game_event_SystemData.Events.INTERACT_SHOW_LOOT.md) (HIGH 100/100) - Game Event
- [SystemData.Events.INTERACT_SHOW_LOOT](../systemdata/fields/systemdata_SystemData.Events.INTERACT_SHOW_LOOT.md) (HIGH 100/100) - SystemData Field

## Flow Diagram

```text
SystemData.Events.ENTER_WORLD <-> SystemData.Events.INTERACT_DONE
```

## Example Code

```lua
SystemData.Events.ENTER_WORLD (HIGH)
```

## Evidence

- SystemData.Events.ENTER_WORLD (HIGH)
- SystemData.Events.EXIT_GAME (HIGH)
- SystemData.Events.INTERACT_COMPLETE_QUEST (HIGH)
- SystemData.Events.INTERACT_DEFAULT (HIGH)
- SystemData.Events.INTERACT_DONE (HIGH)
- SystemData.Events.INTERACT_LOOT_ROLL_FIRST_ITEM (HIGH)
- SystemData.Events.INTERACT_SHOW_INFLUENCE_REWARDS (HIGH)
- SystemData.Events.INTERACT_SHOW_LOOT (HIGH)
