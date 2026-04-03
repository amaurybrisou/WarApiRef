# Game event fan-in

- Category: events
- Confidence: HIGH

## Description

Observed multiple addons converging on a small set of runtime events such as LOADING_END, RELOAD_INTERFACE, and player update events.

## Involved APIs

- [Text](../xml/element_types/element_Text.md) (HIGH 98/100) - XML Element Type
- [CombatLogNewCombatEvent](../events/game_events/game_event_CombatLogNewCombatEvent.md) (HIGH 93/100) - Game Event
- [CombatLogSessionsUpdated](../events/game_events/game_event_CombatLogSessionsUpdated.md) (HIGH 93/100) - Game Event
- [CombatLogSettingsChanged](../events/game_events/game_event_CombatLogSettingsChanged.md) (HIGH 93/100) - Game Event
- [ConfigDialogInitializeSections](../events/game_events/game_event_ConfigDialogInitializeSections.md) (HIGH 93/100) - Game Event
- [GroupsPlayerDistanceUpdated](../events/game_events/game_event_GroupsPlayerDistanceUpdated.md) (HIGH 93/100) - Game Event
- [SettingsChanged](../events/game_events/game_event_SettingsChanged.md) (HIGH 93/100) - Game Event
- [BroadcastMessageInvite](../events/game_events/game_event_BroadcastMessageInvite.md) (HIGH 73/100) - Game Event
- [ChatTextArrived](../events/game_events/game_event_ChatTextArrived.md) (HIGH 73/100) - Game Event
- [GroupsPlayerCombatUpdated](../events/game_events/game_event_GroupsPlayerCombatUpdated.md) (HIGH 73/100) - Game Event

## Flow Diagram

```text
OnLButtonUp
  -> ui: Button, ColorPicker, DynamicImage, Label, ListBox, MapDisplay, Window
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
- GroupsPlayerCombatUpdated (HIGH)
- GroupsPlayerDistanceUpdated (HIGH)
