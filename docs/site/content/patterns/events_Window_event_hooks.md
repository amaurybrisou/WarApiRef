# Window event hooks

- Category: events
- Confidence: HIGH

## Description

Observed XML and window-scoped handlers using On* hooks to bridge engine UI events into Lua.

## Involved APIs

- [Button](../xml/element_types/element_Button.md) (HIGH 100/100) - XML Element Type
- [OnHidden](../xml/handlers/handler_OnHidden.md) (HIGH 100/100) - XML Event
- [OnHidden](../events/window_events/window_event_OnHidden.md) (HIGH 100/100) - Window Event
- [OnLButtonUp](../xml/handlers/handler_OnLButtonUp.md) (HIGH 100/100) - XML Event
- [OnLButtonUp](../events/window_events/window_event_OnLButtonUp.md) (HIGH 100/100) - Window Event
- [OnMButtonUp](../events/window_events/window_event_OnMButtonUp.md) (HIGH 100/100) - Window Event
- [OnMouseWheel](../events/window_events/window_event_OnMouseWheel.md) (HIGH 100/100) - Window Event
- [OnMouseWheel](../xml/handlers/handler_OnMouseWheel.md) (HIGH 100/100) - XML Event
- [OnRButtonUp](../events/window_events/window_event_OnRButtonUp.md) (HIGH 100/100) - Window Event
- [Window](../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type
- [OnInitialize](../xml/handlers/handler_OnInitialize.md) (HIGH 73/100) - XML Event
- [OnLButtonDown](../events/window_events/window_event_OnLButtonDown.md) (HIGH 73/100) - Window Event
- [OnMButtonUp](../xml/handlers/handler_OnMButtonUp.md) (HIGH 73/100) - XML Event
- [OnMouseOver](../xml/handlers/handler_OnMouseOver.md) (HIGH 73/100) - XML Event
- [OnRButtonUp](../xml/handlers/handler_OnRButtonUp.md) (HIGH 73/100) - XML Event
- [OnInitialize](../events/window_events/window_event_OnInitialize.md) (MEDIUM 53/100) - Window Event
- [OnMouseOver](../events/window_events/window_event_OnMouseOver.md) (MEDIUM 53/100) - Window Event

## Flow Diagram

```text
OnLButtonUp
  -> ui: Button, ListBox, Window
```

## Example Code

```lua
OnHidden (HIGH)
```

## Evidence

- OnHidden (HIGH)
- OnInitialize (MEDIUM)
- OnLButtonDown (HIGH)
- OnLButtonUp (HIGH)
- OnMButtonUp (HIGH)
- OnMouseOver (MEDIUM)
- OnMouseWheel (HIGH)
- OnRButtonUp (HIGH)
