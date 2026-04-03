# Game event fan-in

- Category: events
- Confidence: HIGH

## Description

Observed multiple addons converging on a small set of runtime events such as LOADING_END, RELOAD_INTERFACE, and player update events.

## Involved APIs

- [Text](../xml/element_types/element_Text.md) (HIGH 100/100) - XML Element Type
- [ConfigDialogInitializeSections](../events/game_events/game_event_ConfigDialogInitializeSections.md) (MEDIUM 63/100) - Game Event
- [GroupsPlayerDistanceUpdated](../events/game_events/game_event_GroupsPlayerDistanceUpdated.md) (MEDIUM 63/100) - Game Event
- [GroupsPlayerEffectsUpdated](../events/game_events/game_event_GroupsPlayerEffectsUpdated.md) (MEDIUM 63/100) - Game Event
- [BroadcastMessageInvite](../events/game_events/game_event_BroadcastMessageInvite.md) (MEDIUM 43/100) - Game Event
- [CREATED_PREFIX..index](../events/game_events/game_event_CREATED_PREFIX..index.md) (MEDIUM 43/100) - Game Event
- [ChatTextArrived](../events/game_events/game_event_ChatTextArrived.md) (MEDIUM 43/100) - Game Event
- [Global](../events/game_events/game_event_Global.md) (MEDIUM 43/100) - Game Event
- [GroupsPlayerCombatUpdated](../events/game_events/game_event_GroupsPlayerCombatUpdated.md) (MEDIUM 43/100) - Game Event

## Flow Diagram

```text
Global
  -> handlers: Global
```

## Example Code

```lua
BroadcastMessageInvite (MEDIUM)
```

## Evidence

- BroadcastMessageInvite (MEDIUM)
- CREATED_PREFIX..index (MEDIUM)
- ChatTextArrived (MEDIUM)
- ConfigDialogInitializeSections (MEDIUM)
- Global (MEDIUM)
- GroupsPlayerCombatUpdated (MEDIUM)
- GroupsPlayerDistanceUpdated (MEDIUM)
- GroupsPlayerEffectsUpdated (MEDIUM)
