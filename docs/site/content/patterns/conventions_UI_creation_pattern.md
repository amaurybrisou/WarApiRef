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
OnLButtonUp
  -> ui: Button, ColorPicker, DynamicImage, Label, ListBox, MapDisplay, Window
```

## Example Code

```lua
Window creation: AdvancedPetAssist: CreateWindow("APAOptions", true)
```

## Evidence

- Window creation: AdvancedPetAssist: CreateWindow("APAOptions", true)
- Window creation: AdvancedRenownTrainer: CreateWindow("AdvancedRenownTrainingPresetsWindow", false)
- Window creation: AdvancedRenownTrainer: CreateWindow(ImportWindowName, false)
- Window creation: AdvancedRenownTrainer: CreateWindow(ImportNameInputWindowName, false)
- Window creation: AdvancedRenownTrainer: CreateWindow(ExportWindowName, false)
- Window creation: AdvancedRenownTrainer: CreateWindow(LinkWindowName, false)
- Template instantiation: Ace: CreateWindowFromTemplate(w.name, base, w.parent)
- Template instantiation: Ace: CreateWindowFromTemplate(w.name, base, w.parent)
