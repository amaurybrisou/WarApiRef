# Template instantiation

- Category: window
- Confidence: HIGH

## Description

Observed repeated UI elements being instantiated from XML templates at runtime.

## Involved APIs

- [Button](../xml/element_types/element_Button.md) (HIGH 100/100) - XML Element Type
- [EA_Button_Default](../globals/constants/constant_EA_Button_Default.md) (HIGH 100/100) - Constant
- [EA_Button_DefaultWindowClose](../globals/constants/constant_EA_Button_DefaultWindowClose.md) (HIGH 100/100) - Constant
- [Window](../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type
- [CreateWindow](../globals/functions/global_CreateWindow.md) (HIGH 75/100) - Global Function
- [CreateWindowFromTemplate](../globals/functions/global_CreateWindowFromTemplate.md) (HIGH 75/100) - Global Function

## Flow Diagram

```text
OnLButtonUp
  -> ui: Button, ListBox, Window
```

## Example Code

```lua
InfoScroller: CreateWindowFromTemplate(WindowName, "InfoScrollerTemplate", "InfoScrollerMainWindow")
```

## Evidence

- InfoScroller: CreateWindowFromTemplate(WindowName, "InfoScrollerTemplate", "InfoScrollerMainWindow")
- InfoScroller: CreateWindowFromTemplate(w.name, base, w.parent)
- InfoScroller: CreateWindowFromTemplate(w.name, base, w.parent)
- InfoScroller: CreateWindowFromTemplate(w.name, "EA_Button_DefaultWindowClose", w.parent)
- InfoScroller: CreateWindowFromTemplate(w.name, base, w.parent)
- InfoScroller: CreateWindowFromTemplate(w.name, base, w.parent)
