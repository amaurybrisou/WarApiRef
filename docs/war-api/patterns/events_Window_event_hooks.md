# Window event hooks

- Category: events
- Confidence: HIGH

## Description

Observed XML and window-scoped handlers using On* hooks to bridge engine UI events into Lua.

## Involved APIs

- [Button](../xml/element_types/element_Button.md) (HIGH 100/100) - XML Element Type
- [OnHidden](../events/window_events/window_event_OnHidden.md) (HIGH 100/100) - Window Event
- [OnHidden](../xml/handlers/handler_OnHidden.md) (HIGH 100/100) - XML Handler
- [OnHyperLinkLButtonUp](../xml/handlers/handler_OnHyperLinkLButtonUp.md) (HIGH 100/100) - XML Handler
- [OnHyperLinkLButtonUp](../events/window_events/window_event_OnHyperLinkLButtonUp.md) (HIGH 100/100) - Window Event
- [OnHyperLinkRButtonUp](../events/window_events/window_event_OnHyperLinkRButtonUp.md) (HIGH 100/100) - Window Event
- [OnHyperLinkRButtonUp](../xml/handlers/handler_OnHyperLinkRButtonUp.md) (HIGH 100/100) - XML Handler
- [OnInitialize](../events/window_events/window_event_OnInitialize.md) (HIGH 100/100) - Window Event
- [OnInitialize](../xml/handlers/handler_OnInitialize.md) (HIGH 100/100) - XML Handler
- [Window](../globals/tables/table_Window.md) (HIGH 100/100) - Global Table
- [Window](../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type
- [OnActionButtonLButtonDown](../xml/handlers/handler_OnActionButtonLButtonDown.md) (HIGH 73/100) - XML Handler
- [OnActionButtonLButtonUp](../xml/handlers/handler_OnActionButtonLButtonUp.md) (HIGH 73/100) - XML Handler
- [OnActionButtonMouseOver](../xml/handlers/handler_OnActionButtonMouseOver.md) (HIGH 73/100) - XML Handler
- [OnActionButtonRButtonDown](../xml/handlers/handler_OnActionButtonRButtonDown.md) (HIGH 73/100) - XML Handler
- [OnActionButtonLButtonDown](../events/window_events/window_event_OnActionButtonLButtonDown.md) (MEDIUM 53/100) - Window Event
- [OnActionButtonLButtonUp](../events/window_events/window_event_OnActionButtonLButtonUp.md) (MEDIUM 53/100) - Window Event
- [OnActionButtonMouseOver](../events/window_events/window_event_OnActionButtonMouseOver.md) (MEDIUM 53/100) - Window Event
- [OnActionButtonRButtonDown](../events/window_events/window_event_OnActionButtonRButtonDown.md) (MEDIUM 53/100) - Window Event

## Flow Diagram

```text
OnInitialize
  -> ui: Button, Window
```

## Example Code

```lua
OnActionButtonLButtonDown (MEDIUM)
```

## Evidence

- OnActionButtonLButtonDown (MEDIUM)
- OnActionButtonLButtonUp (MEDIUM)
- OnActionButtonMouseOver (MEDIUM)
- OnActionButtonRButtonDown (MEDIUM)
- OnHidden (HIGH)
- OnHyperLinkLButtonUp (HIGH)
- OnHyperLinkRButtonUp (HIGH)
- OnInitialize (HIGH)
