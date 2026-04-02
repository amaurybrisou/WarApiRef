# Window creation

- Category: window
- Confidence: MEDIUM

## Description

Observed top-level UI windows being created from XML definitions at initialization time.

## Involved APIs

- [Window](../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Flow Diagram

```text
OnLButtonUp
  -> ui: Button, ListBox, Window
```

## Example Code

```lua
Moth: CreateWindow("Moth", true)
```

## Evidence

- Moth: CreateWindow("Moth", true)
- TidyChat: CreateWindow(c_TEXT_ENTRY_ANCHOR, false)
- TidyRoll: CreateWindow(c_TROLL_AUTO_ROLL_WINDOW, false)
- TidyRoll: CreateWindow(c_TIDY_ROLL_ANCHOR, false)
- TidyRoll: CreateWindow(c_TIDY_ROLL_TIMER, false)
- TidyRoll: CreateWindow(c_TIDY_ROLL_ESC, false)
