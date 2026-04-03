# Window creation

- Category: window
- Confidence: HIGH

## Description

Observed top-level UI windows being created from XML definitions at initialization time.

## Involved APIs

- [Window](../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type
- [CreateWindow](../globals/functions/global_CreateWindow.md) (HIGH 75/100) - Global Function
- [windowName](../events/game_events/game_event_windowName.md) (MEDIUM 43/100) - Game Event

## Flow Diagram

```text
windowName
  -> handlers: windowName
```

## Example Code

```lua
AbilityAlert: CreateWindow("AAWindow", true)
```

## Evidence

- AbilityAlert: CreateWindow("AAWindow", true)
- ActionFraction: CreateWindow(windowName, true)
- AdvancedPetAssist: CreateWindow("APAOptions", true)
- AdvancedRenownTrainer: CreateWindow("AdvancedRenownTrainingPresetsWindow", false)
- AdvancedRenownTrainer: CreateWindow(ImportWindowName, false)
- AdvancedRenownTrainer: CreateWindow(ImportNameInputWindowName, false)
