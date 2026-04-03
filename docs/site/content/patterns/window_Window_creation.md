# Window creation

- Category: window
- Confidence: HIGH

## Description

Observed top-level UI windows being created from XML definitions at initialization time.

## Involved APIs

- [Window](../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type
- [CreateWindow](../globals/functions/global_CreateWindow.md) (HIGH 75/100) - Global Function

## Flow Diagram

```text
OnLButtonUp
  -> ui: Button, ListBox, Window
```

## Example Code

```lua
InfoScroller: CreateWindow("InfoScrollerMainWindow", true)
```

## Evidence

- InfoScroller: CreateWindow("InfoScrollerMainWindow", true)
- Moth: CreateWindow("Moth", true)
- Soloq: CreateWindow(overviewWindowName, false)
- TidyChat: CreateWindow(c_TEXT_ENTRY_ANCHOR, false)
- TidyRoll: CreateWindow(c_TROLL_AUTO_ROLL_WINDOW, false)
- TidyRoll: CreateWindow(c_TIDY_ROLL_ANCHOR, false)
