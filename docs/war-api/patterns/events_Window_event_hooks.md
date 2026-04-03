# Window event hooks

- Category: events
- Confidence: HIGH

## Description

Observed XML and window-scoped handlers using On* hooks to bridge engine UI events into Lua.

## Involved APIs

- [Button](../xml/element_types/element_Button.md) (HIGH 100/100) - XML Element Type
- [OnLButtonUp](../xml/handlers/handler_OnLButtonUp.md) (HIGH 100/100) - XML Event
- [Window](../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type
- [OnLButtonUp](../events/window_events/window_event_OnLButtonUp.md) (MEDIUM 53/100) - Window Event

## Flow Diagram

```text
OnLButtonUp
  -> handlers: AuctionWindow.Hide, EA_Window_OpenParty.ToggleFullWindow, ScenarioGroupWindow.LeaveGroup
  -> ui: Button
```

## Example Code

```lua
OnLButtonUp (MEDIUM)
```

## Evidence

- OnLButtonUp (MEDIUM)
