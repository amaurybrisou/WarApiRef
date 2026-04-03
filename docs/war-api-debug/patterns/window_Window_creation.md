# Window creation

- Category: window
- Confidence: MEDIUM

## Description

Observed top-level UI windows being created from XML definitions at initialization time.

## Involved APIs

- [Window](../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type
- [Button](../xml/element_types/element_Button.md) (HIGH 90/100) - XML Element Type

## Flow Diagram

```text
Button <-> Window
```

## Example Code

```lua
CM_ClosetGoblin: CreateWindow("yClosetGoblinButton", true)
```

## Evidence

- CM_ClosetGoblin: CreateWindow("yClosetGoblinButton", true)
- Clock: CreateWindow("ClockWindow", true)
