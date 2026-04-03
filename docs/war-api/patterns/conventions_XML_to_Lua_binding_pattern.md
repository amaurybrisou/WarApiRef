# XML to Lua binding pattern

- Category: conventions
- Confidence: HIGH

## Description

XML handler names map directly to Lua functions and can be cross-checked through the bindings page.

## Involved APIs

- [Button](../xml/element_types/element_Button.md) (HIGH 100/100) - XML Element Type
- [OnHidden](../events/window_events/window_event_OnHidden.md) (HIGH 100/100) - Window Event
- [OnHidden](../xml/handlers/handler_OnHidden.md) (HIGH 100/100) - XML Event
- [OnLButtonUp](../events/window_events/window_event_OnLButtonUp.md) (HIGH 100/100) - Window Event
- [OnLButtonUp](../xml/handlers/handler_OnLButtonUp.md) (HIGH 100/100) - XML Event
- [OnMouseWheel](../events/window_events/window_event_OnMouseWheel.md) (HIGH 100/100) - Window Event
- [OnMouseWheel](../xml/handlers/handler_OnMouseWheel.md) (HIGH 100/100) - XML Event
- [OnSelChanged](../events/window_events/window_event_OnSelChanged.md) (HIGH 100/100) - Window Event
- [OnSelChanged](../xml/handlers/handler_OnSelChanged.md) (HIGH 100/100) - XML Event
- [OnShown](../events/window_events/window_event_OnShown.md) (HIGH 100/100) - Window Event
- [OnShown](../xml/handlers/handler_OnShown.md) (HIGH 100/100) - XML Event
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
- TidyChat: TidyChatCopyClose.OnLButtonUp -> TidyChat.Copy.OnClose
- TidyChat: TidyChatCopyNext.OnLButtonUp -> TidyChat.Copy.CopyNext
