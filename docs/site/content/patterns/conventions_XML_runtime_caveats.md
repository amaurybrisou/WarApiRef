# XML runtime caveats

- Category: conventions
- Confidence: MEDIUM

## Description

Implementation-validated findings show that XML input and scroll layout behavior can depend on ancestor state and on outer-window sizing, even when child nodes appear correctly configured.

## Involved APIs

- [Button](../xml/element_types/element_Button.md) (HIGH 100/100) - XML Element Type
- [OnLButtonUp](../events/window_events/window_event_OnLButtonUp.md) (HIGH 100/100) - Window Event
- [OnLButtonUp](../xml/handlers/handler_OnLButtonUp.md) (HIGH 100/100) - XML Handler
- [ScrollWindow](../xml/element_types/element_ScrollWindow.md) (HIGH 100/100) - XML Element Type
- [ScrollWindowUpdateScrollRect](../window_api/functions/window_ScrollWindowUpdateScrollRect.md) (HIGH 100/100) - Window Function
- [Window](../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type
- [Window](../globals/tables/table_Window.md) (HIGH 100/100) - Global Table

## Flow Diagram

```text
OnLButtonUp
  -> handlers: DebugWindow.ClearTextLog, InterfaceCore.ReloadUI
  -> ui: Button, ColorPicker, DynamicImage, FullResizeImage, Label, ListBox, MapDisplay, Window
```

## Example Code

```lua
WhoHealedMe: a child `OnLButtonUp` target remained inert until the parent or template input chain was made input-enabled.
```

## Evidence

- WhoHealedMe: a child `OnLButtonUp` target remained inert until the parent or template input chain was made input-enabled.
- guidance: validate ancestor `handleinput` state across the clickable chain, not only on the child node.
- caveat: treat this as a reusable runtime warning, not a guaranteed engine contract.
- WhoHealedMe: nested scroll content dimensions initially under-reported usable space during early layout.
- guidance: compute size from the outer parent first, then resize child content and call `ScrollWindowUpdateScrollRect`.
