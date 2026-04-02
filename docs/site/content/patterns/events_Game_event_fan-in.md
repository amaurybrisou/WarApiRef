# Game event fan-in

- Category: events
- Confidence: HIGH

## Description

Observed multiple addons converging on a small set of runtime events such as LOADING_END, RELOAD_INTERFACE, and player update events.

## Involved APIs

- [EA_CareerResourceWindow](../events/game_events/game_event_EA_CareerResourceWindow.md) (HIGH 100/100) - Game Event
- [Text](../xml/element_types/element_Text.md) (HIGH 100/100) - XML Element Type
- [Window](../globals/tables/table_Window.md) (HIGH 100/100) - Global Table
- [Window](../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type
- [CombatLogNewCombatEvent](../events/game_events/game_event_CombatLogNewCombatEvent.md) (HIGH 93/100) - Game Event
- [CombatLogSessionsUpdated](../events/game_events/game_event_CombatLogSessionsUpdated.md) (HIGH 93/100) - Game Event
- [CombatLogSettingsChanged](../events/game_events/game_event_CombatLogSettingsChanged.md) (HIGH 93/100) - Game Event
- [ConfigDialogInitializeSections](../events/game_events/game_event_ConfigDialogInitializeSections.md) (HIGH 93/100) - Game Event
- [SettingsChanged](../events/game_events/game_event_SettingsChanged.md) (HIGH 93/100) - Game Event
- [BroadcastMessageInvite](../events/game_events/game_event_BroadcastMessageInvite.md) (HIGH 73/100) - Game Event
- [ChatTextArrived](../events/game_events/game_event_ChatTextArrived.md) (HIGH 73/100) - Game Event
- [Global](../events/game_events/game_event_Global.md) (HIGH 73/100) - Game Event

## Flow Diagram

```text
Global
  -> handlers: Global
```

## Example Code

```lua
BroadcastMessageInvite (HIGH)
```

## Evidence

- BroadcastMessageInvite (HIGH)
- ChatTextArrived (HIGH)
- CombatLogNewCombatEvent (HIGH)
- CombatLogSessionsUpdated (HIGH)
- CombatLogSettingsChanged (HIGH)
- ConfigDialogInitializeSections (HIGH)
- EA_CareerResourceWindow (HIGH)
- Global (HIGH)
