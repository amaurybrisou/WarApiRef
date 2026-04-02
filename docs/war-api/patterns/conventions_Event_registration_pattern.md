# Event registration pattern

- Category: conventions
- Confidence: HIGH

## Description

Runtime events are typically wired through RegisterEventHandler or window-scoped variants, with handler names staying module-qualified.

## Involved APIs

- [SystemData.Events.LOADING_END](../events/game_events/game_event_SystemData.Events.LOADING_END.md) (HIGH 100/100) - Game Event
- [SystemData.Events.LOADING_END](../systemdata/fields/systemdata_SystemData.Events.LOADING_END.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.L_BUTTON_UP_PROCESSED](../events/game_events/game_event_SystemData.Events.L_BUTTON_UP_PROCESSED.md) (HIGH 100/100) - Game Event
- [SystemData.Events.L_BUTTON_UP_PROCESSED](../systemdata/fields/systemdata_SystemData.Events.L_BUTTON_UP_PROCESSED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.RELOAD_INTERFACE](../events/game_events/game_event_SystemData.Events.RELOAD_INTERFACE.md) (HIGH 100/100) - Game Event
- [SystemData.Events.RELOAD_INTERFACE](../systemdata/fields/systemdata_SystemData.Events.RELOAD_INTERFACE.md) (HIGH 100/100) - SystemData Field
- [RegisterEventHandler](../globals/functions/global_RegisterEventHandler.md) (HIGH 81/100) - Global Function

## Flow Diagram

```text
OnLButtonUp <-> OnLButtonUp
```

## Example Code

```lua
RegisterEventHandler: TidyChat: RegisterEventHandler(SystemData.Events.LOADING_END, "TidyChat.OnLoad")
```

## Evidence

- RegisterEventHandler: TidyChat: RegisterEventHandler(SystemData.Events.LOADING_END, "TidyChat.OnLoad")
- RegisterEventHandler: TidyChat: RegisterEventHandler(SystemData.Events.RELOAD_INTERFACE, "TidyChat.OnLoad")
- RegisterEventHandler: TidyChat: RegisterEventHandler(SystemData.Events.L_BUTTON_UP_PROCESSED, "TidyChat.OnLBU")
- RegisterEventHandler: TidyChat: RegisterEventHandler(chatLogEventId, "TidyChat.OnChatLogUpdated")
- RegisterEventHandler: TidyChat: RegisterEventHandler(combatLogEventId, "TidyChat.OnCombatLogUpdated")
- RegisterEventHandler: TidyChat: RegisterEventHandler(systemLogEventId, "TidyChat.OnSystemLogUpdated")
- UnregisterEventHandler: TidyRoll: UnregisterEventHandler(SystemData.Events.LOADING_END, "TidyRoll.OnLoad")
- UnregisterEventHandler: TidyRoll: UnregisterEventHandler(SystemData.Events.RELOAD_INTERFACE, "TidyRoll.OnLoad")
