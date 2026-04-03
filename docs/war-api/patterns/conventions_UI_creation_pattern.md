# UI creation pattern

- Category: conventions
- Confidence: HIGH

## Description

UI is commonly created from XML, then positioned in Lua through CreateWindow or CreateWindowFromTemplate followed by anchor calls.

## Involved APIs

- [Window](../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type
- [CreateWindow](../globals/functions/global_CreateWindow.md) (HIGH 75/100) - Global Function
- [CreateWindowFromTemplate](../globals/functions/global_CreateWindowFromTemplate.md) (HIGH 75/100) - Global Function

## Flow Diagram

```text
CreateWindow <-> CreateWindowFromTemplate
```

## Example Code

```lua
Window creation: InfoScroller: CreateWindow("InfoScrollerMainWindow", true)
```

## Evidence

- Window creation: InfoScroller: CreateWindow("InfoScrollerMainWindow", true)
- Window creation: Moth: CreateWindow("Moth", true)
- Window creation: Soloq: CreateWindow(overviewWindowName, false)
- Window creation: TidyChat: CreateWindow(c_TEXT_ENTRY_ANCHOR, false)
- Window creation: TidyRoll: CreateWindow(c_TROLL_AUTO_ROLL_WINDOW, false)
- Window creation: TidyRoll: CreateWindow(c_TIDY_ROLL_ANCHOR, false)
- Template instantiation: InfoScroller: CreateWindowFromTemplate(WindowName, "InfoScrollerTemplate", "InfoScrollerMainWindow")
- Template instantiation: InfoScroller: CreateWindowFromTemplate(w.name, base, w.parent)
