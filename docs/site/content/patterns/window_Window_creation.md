# Window creation

- Category: window
- Confidence: HIGH

## Description

Observed top-level UI windows being created from XML definitions at initialization time.

## Involved APIs

- [Window](../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type
- [Window](../globals/tables/table_Window.md) (HIGH 100/100) - Global Table
- [CreateWindow](../globals/functions/global_CreateWindow.md) (HIGH 75/100) - Global Function

## Flow Diagram

```text
OnLButtonUp
  -> handlers: DebugWindow.ClearTextLog, InterfaceCore.ReloadUI
  -> ui: Button, ColorPicker, DynamicImage, FullResizeImage, Label, ListBox, MapDisplay, Window
```

## Example Code

```lua
AdvancedPetAssist: CreateWindow("APAOptions", true)
```

## Evidence

- AdvancedPetAssist: CreateWindow("APAOptions", true)
- AdvancedRenownTrainer: CreateWindow("AdvancedRenownTrainingPresetsWindow", false)
- AdvancedRenownTrainer: CreateWindow(ImportWindowName, false)
- AdvancedRenownTrainer: CreateWindow(ImportNameInputWindowName, false)
- AdvancedRenownTrainer: CreateWindow(ExportWindowName, false)
- AdvancedRenownTrainer: CreateWindow(LinkWindowName, false)
