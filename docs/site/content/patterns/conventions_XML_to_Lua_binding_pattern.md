# XML to Lua binding pattern

- Category: conventions
- Confidence: HIGH

## Description

XML handler names map directly to Lua functions and can be cross-checked through the bindings page.

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
- AdvancedPetAssist: APAComboDebug.OnSelChanged -> APAGui.OnComboChanged
- AdvancedPetAssist: APAComboEnabled.OnSelChanged -> APAGui.OnComboChanged
