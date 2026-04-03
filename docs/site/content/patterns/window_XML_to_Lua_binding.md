# XML to Lua binding

- Category: window
- Confidence: HIGH

## Description

Observed XML handler events routing directly into addon Lua functions through shared engine bindings.

## Involved APIs

- [Button](../xml/element_types/element_Button.md) (HIGH 100/100) - XML Element Type
- [OnLButtonUp](../xml/handlers/handler_OnLButtonUp.md) (HIGH 100/100) - XML Event
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
