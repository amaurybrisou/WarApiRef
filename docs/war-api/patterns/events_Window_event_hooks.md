# Window event hooks

- Category: events
- Confidence: HIGH

## Description

Observed XML and window-scoped handlers using On* hooks to bridge engine UI events into Lua.

## Involved APIs

- [Button](../xml/element_types/element_Button.md) (HIGH 100/100) - XML Element Type
- [OnHidden](../xml/handlers/handler_OnHidden.md) (HIGH 100/100) - XML Event
- [OnHidden](../events/window_events/window_event_OnHidden.md) (HIGH 100/100) - Window Event
- [OnHyperLinkLButtonUp](../events/window_events/window_event_OnHyperLinkLButtonUp.md) (HIGH 100/100) - Window Event
- [OnHyperLinkLButtonUp](../xml/handlers/handler_OnHyperLinkLButtonUp.md) (HIGH 100/100) - XML Event
- [OnHyperLinkRButtonUp](../xml/handlers/handler_OnHyperLinkRButtonUp.md) (HIGH 100/100) - XML Event
- [OnHyperLinkRButtonUp](../events/window_events/window_event_OnHyperLinkRButtonUp.md) (HIGH 100/100) - Window Event
- [OnInitialize](../xml/handlers/handler_OnInitialize.md) (HIGH 100/100) - XML Event
- [OnInitialize](../events/window_events/window_event_OnInitialize.md) (HIGH 100/100) - Window Event
- [OnKeyEnter](../events/window_events/window_event_OnKeyEnter.md) (HIGH 100/100) - Window Event
- [OnKeyEscape](../xml/handlers/handler_OnKeyEscape.md) (HIGH 100/100) - XML Event
- [OnKeyEscape](../events/window_events/window_event_OnKeyEscape.md) (HIGH 100/100) - Window Event
- [OnLButtonDown](../events/window_events/window_event_OnLButtonDown.md) (HIGH 100/100) - Window Event
- [OnLButtonDown](../xml/handlers/handler_OnLButtonDown.md) (HIGH 100/100) - XML Event
- [OnLButtonUp](../events/window_events/window_event_OnLButtonUp.md) (HIGH 100/100) - Window Event
- [OnLButtonUp](../xml/handlers/handler_OnLButtonUp.md) (HIGH 100/100) - XML Event
- [Window](../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type
- [OnKeyEnter](../xml/handlers/handler_OnKeyEnter.md) (HIGH 93/100) - XML Event

## Flow Diagram

```text
OnLButtonUp
  -> ui: Button, ColorPicker, DynamicImage, Label, ListBox, MapDisplay, Window
```

## Example Code

```lua
OnHidden (HIGH)
```

## Evidence

- OnHidden (HIGH)
- OnHyperLinkLButtonUp (HIGH)
- OnHyperLinkRButtonUp (HIGH)
- OnInitialize (HIGH)
- OnKeyEnter (HIGH)
- OnKeyEscape (HIGH)
- OnLButtonDown (HIGH)
- OnLButtonUp (HIGH)
