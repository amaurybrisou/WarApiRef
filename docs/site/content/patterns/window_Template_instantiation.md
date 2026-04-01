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
- [Window](../globals/tables/table_Window.md) (HIGH 100/100) - Global Table
- [CreateWindow](../globals/functions/global_CreateWindow.md) (HIGH 75/100) - Global Function
- [CreateWindowFromTemplate](../globals/functions/global_CreateWindowFromTemplate.md) (HIGH 75/100) - Global Function

## Flow Diagram

```text
OnLButtonUp
  -> handlers: DebugWindow.ClearTextLog, InterfaceCore.ReloadUI
  -> ui: Button, ColorPicker, ComboBox, DynamicImage, FullResizeImage, Label, ListBox, MapDisplay, Window
```

## Example Code

```lua
Ace: CreateWindowFromTemplate(w.name, base, w.parent)
```

## Evidence

- Ace: CreateWindowFromTemplate(w.name, base, w.parent)
- Ace: CreateWindowFromTemplate(w.name, base, w.parent)
- Ace: CreateWindowFromTemplate(w.name, "EA_Button_DefaultWindowClose", w.parent)
- Ace: CreateWindowFromTemplate(w.name, base, w.parent)
- Ace: CreateWindowFromTemplate(w.name, base, w.parent)
- Ace: CreateWindowFromTemplate(w.name, base, w.parent)
