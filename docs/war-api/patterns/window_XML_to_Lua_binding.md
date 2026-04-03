# XML to Lua binding

- Category: window
- Confidence: HIGH

## Description

Observed XML handler events routing directly into addon Lua functions through shared engine bindings.

## Involved APIs

- [Button](../xml/element_types/element_Button.md) (HIGH 100/100) - XML Element Type
- [EA_ChatWindow](../globals/tables/table_EA_ChatWindow.md) (HIGH 100/100) - Global Table
- [Label](../xml/element_types/element_Label.md) (HIGH 100/100) - XML Element Type
- [OnHyperLinkLButtonUp](../events/window_events/window_event_OnHyperLinkLButtonUp.md) (HIGH 100/100) - Window Event
- [OnHyperLinkLButtonUp](../xml/handlers/handler_OnHyperLinkLButtonUp.md) (HIGH 100/100) - XML Event
- [OnHyperLinkRButtonUp](../events/window_events/window_event_OnHyperLinkRButtonUp.md) (HIGH 100/100) - Window Event
- [OnHyperLinkRButtonUp](../xml/handlers/handler_OnHyperLinkRButtonUp.md) (HIGH 100/100) - XML Event
- [Window](../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Flow Diagram

```text
Label <-> OnHyperLinkRButtonUp
```

## Example Code

```lua
InfoScroller: InfoScrollerTemplateLabel1.OnHyperLinkLButtonUp -> EA_ChatWindow.OnHyperLinkLButtonUp
```

## Evidence

- InfoScroller: InfoScrollerTemplateLabel1.OnHyperLinkLButtonUp -> EA_ChatWindow.OnHyperLinkLButtonUp
- InfoScroller: InfoScrollerTemplateLabel1.OnHyperLinkRButtonUp -> EA_ChatWindow.OnHyperLinkRButtonUp
- InfoScroller: InfoScrollerTemplateLabel2.OnHyperLinkLButtonUp -> EA_ChatWindow.OnHyperLinkLButtonUp
- InfoScroller: InfoScrollerTemplateLabel2.OnHyperLinkRButtonUp -> EA_ChatWindow.OnHyperLinkRButtonUp
- InfoScroller: InfoScrollerTemplateLabel3.OnHyperLinkLButtonUp -> EA_ChatWindow.OnHyperLinkLButtonUp
- InfoScroller: InfoScrollerTemplateLabel3.OnHyperLinkRButtonUp -> EA_ChatWindow.OnHyperLinkRButtonUp
