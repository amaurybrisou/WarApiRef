# XML to Lua binding

- Category: window
- Confidence: HIGH

## Description

Observed XML handler events routing directly into addon Lua functions through shared engine bindings.

## Involved APIs

- [OnSelChanged](../xml/handlers/handler_OnSelChanged.md) (HIGH 100/100) - XML Handler
- [OnSelChanged](../events/window_events/window_event_OnSelChanged.md) (HIGH 100/100) - Window Event

## Flow Diagram

```text
OnSelChanged <-> OnSelChanged
```

## Example Code

```lua
AdvancedPetAssist: APAComboAttackBind.OnSelChanged -> APAGui.OnComboChanged
```

## Evidence

- AdvancedPetAssist: APAComboAttackBind.OnSelChanged -> APAGui.OnComboChanged
- AdvancedPetAssist: APAComboAutoReattack.OnSelChanged -> APAGui.OnComboChanged
- AdvancedPetAssist: APAComboAutoReattackDelay.OnSelChanged -> APAGui.OnComboChanged
- AdvancedPetAssist: APAComboCastDelay.OnSelChanged -> APAGui.OnComboChanged
- AdvancedPetAssist: APAComboCastOnAcquire.OnSelChanged -> APAGui.OnComboChanged
- AdvancedPetAssist: APAComboCombatExitDelay.OnSelChanged -> APAGui.OnComboChanged
