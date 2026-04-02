# UI creation pattern

- Category: conventions
- Confidence: HIGH

## Description

UI is commonly created from XML, then positioned in Lua through CreateWindow or CreateWindowFromTemplate followed by anchor calls.

## Involved APIs

- [Button](../xml/element_types/element_Button.md) (HIGH 100/100) - XML Element Type
- [Window](../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Flow Diagram

```text
OnLButtonUp
  -> ui: Button, ListBox, Window
```

## Example Code

```lua
Window creation: Moth: CreateWindow("Moth", true)
```

## Evidence

- Window creation: Moth: CreateWindow("Moth", true)
- Window creation: TidyChat: CreateWindow(c_TEXT_ENTRY_ANCHOR, false)
- Window creation: TidyRoll: CreateWindow(c_TROLL_AUTO_ROLL_WINDOW, false)
- Window creation: TidyRoll: CreateWindow(c_TIDY_ROLL_ANCHOR, false)
- Window creation: TidyRoll: CreateWindow(c_TIDY_ROLL_TIMER, false)
- Window creation: TidyRoll: CreateWindow(c_TIDY_ROLL_ESC, false)
- Template instantiation: TidyChat: CreateWindowFromTemplate(c_CHANNEL_MENU.."AllianceButton", "ChannelMenuButton", "ChatChannelSelectionWindow")
- Template instantiation: TidyChat: CreateWindowFromTemplate(c_CHANNEL_MENU.."AdviceButton", "ChannelMenuButton", "ChatChannelSelectionWindow")
