# UI creation pattern

- Category: conventions
- Confidence: HIGH

## Description

UI is commonly created from XML, then positioned in Lua through CreateWindow or CreateWindowFromTemplate followed by anchor calls.

## Involved APIs

- [Button](../xml/element_types/element_Button.md) (HIGH 100/100) - XML Element Type
- [EA_Button_Default](../globals/constants/constant_EA_Button_Default.md) (HIGH 100/100) - Constant
- [EA_Button_DefaultWindowClose](../globals/constants/constant_EA_Button_DefaultWindowClose.md) (HIGH 100/100) - Constant
- [Window](../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type
- [CreateWindow](../globals/functions/global_CreateWindow.md) (HIGH 75/100) - Global Function
- [CreateWindowFromTemplate](../globals/functions/global_CreateWindowFromTemplate.md) (HIGH 75/100) - Global Function
- [windowName](../events/game_events/game_event_windowName.md) (MEDIUM 43/100) - Game Event

## Flow Diagram

```text
CreateWindow -> CreateWindowFromTemplate
```

## Example Code

```lua
Window creation: AbilityAlert: CreateWindow("AAWindow", true)
```

## Evidence

- Window creation: AbilityAlert: CreateWindow("AAWindow", true)
- Window creation: ActionFraction: CreateWindow(windowName, true)
- Window creation: AdvancedPetAssist: CreateWindow("APAOptions", true)
- Window creation: AdvancedRenownTrainer: CreateWindow("AdvancedRenownTrainingPresetsWindow", false)
- Window creation: AdvancedRenownTrainer: CreateWindow(ImportWindowName, false)
- Window creation: AdvancedRenownTrainer: CreateWindow(ImportNameInputWindowName, false)
- Template instantiation: Ace: CreateWindowFromTemplate(w.name, base, w.parent)
- Template instantiation: Ace: CreateWindowFromTemplate(w.name, "EA_Button_DefaultWindowClose", w.parent)
