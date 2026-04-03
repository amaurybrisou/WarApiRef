# RegisterEventHandler

- Category: events
- Confidence: MEDIUM

## Description

Observed registering global runtime handlers against shared event identifiers.

## Involved APIs

- [SystemData.Events.LOADING_END](../systemdata/fields/systemdata_SystemData.Events.LOADING_END.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.LOADING_END](../events/game_events/game_event_SystemData.Events.LOADING_END.md) (HIGH 100/100) - Game Event
- [SystemData.Events.L_BUTTON_UP_PROCESSED](../events/game_events/game_event_SystemData.Events.L_BUTTON_UP_PROCESSED.md) (HIGH 100/100) - Game Event
- [SystemData.Events.L_BUTTON_UP_PROCESSED](../systemdata/fields/systemdata_SystemData.Events.L_BUTTON_UP_PROCESSED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.RELOAD_INTERFACE](../systemdata/fields/systemdata_SystemData.Events.RELOAD_INTERFACE.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.RELOAD_INTERFACE](../events/game_events/game_event_SystemData.Events.RELOAD_INTERFACE.md) (HIGH 100/100) - Game Event
- [RegisterEventHandler](../globals/functions/global_RegisterEventHandler.md) (HIGH 81/100) - Global Function

## Flow Diagram

```text
Button <-> OnLButtonUp
```

## Example Code

```lua
TidyChat: RegisterEventHandler(SystemData.Events.LOADING_END, "TidyChat.OnLoad")
```

## Evidence

- TidyChat: RegisterEventHandler(SystemData.Events.LOADING_END, "TidyChat.OnLoad")
- TidyChat: RegisterEventHandler(SystemData.Events.RELOAD_INTERFACE, "TidyChat.OnLoad")
- TidyChat: RegisterEventHandler(SystemData.Events.L_BUTTON_UP_PROCESSED, "TidyChat.OnLBU")
- TidyChat: RegisterEventHandler(chatLogEventId, "TidyChat.OnChatLogUpdated")
- TidyChat: RegisterEventHandler(combatLogEventId, "TidyChat.OnCombatLogUpdated")
- TidyChat: RegisterEventHandler(systemLogEventId, "TidyChat.OnSystemLogUpdated")
