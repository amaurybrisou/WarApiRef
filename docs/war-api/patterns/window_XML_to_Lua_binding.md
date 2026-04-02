# XML to Lua binding

- Category: window
- Confidence: HIGH

## Description

Observed XML handler events routing directly into addon Lua functions through shared engine bindings.

## Involved APIs

- [Button](../xml/element_types/element_Button.md) (HIGH 100/100) - XML Element Type
- [OnHidden](../events/window_events/window_event_OnHidden.md) (HIGH 100/100) - Window Event
- [OnHidden](../xml/handlers/handler_OnHidden.md) (HIGH 100/100) - XML Handler
- [OnLButtonUp](../xml/handlers/handler_OnLButtonUp.md) (HIGH 100/100) - XML Handler
- [OnLButtonUp](../events/window_events/window_event_OnLButtonUp.md) (HIGH 100/100) - Window Event
- [OnMouseWheel](../xml/handlers/handler_OnMouseWheel.md) (HIGH 100/100) - XML Handler
- [OnMouseWheel](../events/window_events/window_event_OnMouseWheel.md) (HIGH 100/100) - Window Event
- [OnSelChanged](../xml/handlers/handler_OnSelChanged.md) (HIGH 100/100) - XML Handler
- [OnSelChanged](../events/window_events/window_event_OnSelChanged.md) (HIGH 100/100) - Window Event
- [OnShown](../events/window_events/window_event_OnShown.md) (HIGH 100/100) - Window Event
- [OnShown](../xml/handlers/handler_OnShown.md) (HIGH 100/100) - XML Handler
- [Window](../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Flow Diagram

```text
OnLButtonUp
  -> ui: Button, ListBox, Window
```

## Example Code

```lua
TidyChat: TChatCheckboxTemplate.OnLButtonUp -> TidyChat.Options.OnCheckboxLBU
```

## Evidence

- TidyChat: TChatCheckboxTemplate.OnLButtonUp -> TidyChat.Options.OnCheckboxLBU
- TidyChat: TChatTabButton.OnLButtonUp -> TidyChat.Options.OnTabLBU
- TidyChat: TChatTabWindowsTemplateSelectWindowCombo.OnSelChanged -> TidyChat.Options.UpdateGroupTabs
- TidyChat: TidyChatCopy.OnHidden -> TidyChat.Copy.OnHidden
- TidyChat: TidyChatCopy.OnMouseWheel -> TidyChat.Copy.OnMouseWheel
- TidyChat: TidyChatCopy.OnShown -> TidyChat.Copy.OnShown
