# XML to Lua binding pattern

- Category: conventions
- Confidence: HIGH

## Description

XML handler names map directly to Lua functions and can be cross-checked through the bindings page.

## Involved APIs

- [Button](../xml/element_types/element_Button.md) (HIGH 100/100) - XML Element Type
- [OnLButtonUp](../xml/handlers/handler_OnLButtonUp.md) (HIGH 100/100) - XML Event
- [OnSelChanged](../xml/handlers/handler_OnSelChanged.md) (HIGH 88/100) - XML Event
- [OnLButtonUp](../events/window_events/window_event_OnLButtonUp.md) (MEDIUM 53/100) - Window Event

## Flow Diagram

```text
OnLButtonUp
  -> handlers: ScenarioGroupWindow.LeaveGroup
  -> ui: Button
```

## Example Code

```lua
AdvancedPetAssist: .OnLButtonUp -> APAGui.OnTabButtonUp
```

## Evidence

- AdvancedPetAssist: .OnLButtonUp -> APAGui.OnTabButtonUp
- AdvancedPetAssist: .OnLButtonUp -> APAGui.OnTabButtonUp
- AdvancedPetAssist: .OnLButtonUp -> APAGui.OnTabButtonUp
- AdvancedPetAssist: .OnLButtonUp -> APAGui.OnTabButtonUp
- AdvancedPetAssist: .OnLButtonUp -> APAGui.OnTabButtonUp
- AdvancedPetAssist: .OnLButtonUp -> APAGui.OnTabButtonUp
- AdvancedPetAssist: .OnSelChanged -> APAGui.OnComboChanged
- AdvancedPetAssist: .OnSelChanged -> APAGui.OnComboChanged
