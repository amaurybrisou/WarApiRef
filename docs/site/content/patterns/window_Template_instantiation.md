# Template instantiation

- Category: window
- Confidence: HIGH

## Description

Observed repeated UI elements being instantiated from XML templates at runtime.

## Involved APIs

- [Button](../xml/element_types/element_Button.md) (HIGH 100/100) - XML Element Type
- [ComboBox](../xml/element_types/element_ComboBox.md) (HIGH 100/100) - XML Element Type
- [EA_Button_Default](../globals/constants/constant_EA_Button_Default.md) (HIGH 100/100) - Constant
- [EA_Button_DefaultCheckBox](../globals/constants/constant_EA_Button_DefaultCheckBox.md) (HIGH 100/100) - Constant
- [EA_Button_DefaultWindowClose](../globals/constants/constant_EA_Button_DefaultWindowClose.md) (HIGH 100/100) - Constant
- [EA_ComboBox_DefaultResizable](../globals/constants/constant_EA_ComboBox_DefaultResizable.md) (HIGH 100/100) - Constant
- [EA_EditBox_DefaultFrame](../globals/constants/constant_EA_EditBox_DefaultFrame.md) (HIGH 100/100) - Constant
- [EditBox](../xml/element_types/element_EditBox.md) (HIGH 100/100) - XML Element Type
- [Window](../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type
- [CreateWindow](../globals/functions/global_CreateWindow.md) (HIGH 75/100) - Global Function
- [CreateWindowFromTemplate](../globals/functions/global_CreateWindowFromTemplate.md) (HIGH 75/100) - Global Function

## Flow Diagram

```text
CreateWindow -> CreateWindowFromTemplate
```

## Example Code

```lua
Ace: CreateWindowFromTemplate(w.name, base, w.parent)
```

## Evidence

- Ace: CreateWindowFromTemplate(w.name, base, w.parent)
- Ace: CreateWindowFromTemplate(w.name, "EA_Button_DefaultWindowClose", w.parent)
- ActionBarHide: CreateWindowFromTemplate(w.name, "EA_ComboBox_DefaultResizable", w.parent)
- ActionBarHide: CreateWindowFromTemplate(w.name, "EA_Button_DefaultCheckBox", w.parent)
- ActionBarHide: CreateWindowFromTemplate(w.name, "EA_Button_DefaultWindowClose", w.parent)
- ActionBarHide: CreateWindowFromTemplate(w.name, "EA_EditBox_DefaultFrame", w.parent)
