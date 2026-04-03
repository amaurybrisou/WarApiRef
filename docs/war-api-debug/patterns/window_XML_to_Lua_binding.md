# XML to Lua binding

- Category: window
- Confidence: HIGH

## Description

Observed XML handler events routing directly into addon Lua functions through shared engine bindings.

## Involved APIs

- [Window](../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type
- [Button](../xml/element_types/element_Button.md) (HIGH 90/100) - XML Element Type

## Flow Diagram

```text
Button <-> Window
```

## Example Code

```lua
CM_ClosetGoblin: .OnLButtonUp -> ClosetGoblinZoneWindow.OnClickZoneRow
```

## Evidence

- CM_ClosetGoblin: .OnLButtonUp -> ClosetGoblinZoneWindow.OnClickZoneRow
- CM_ClosetGoblin: .OnRButtonUp -> ClosetGoblinZoneWindow.OnSetZoneRowContextMenu
- CM_ClosetGoblin: .OnLButtonUp -> ClosetGoblinZoneWindow.OnClickZoneRow
- CM_ClosetGoblin: .OnRButtonUp -> ClosetGoblinZoneWindow.OnSetZoneRowContextMenu
- CM_ClosetGoblin: .OnLButtonUp -> ClosetGoblinCharacterWindow.HotbarPageUpProxy
- CM_ClosetGoblin: .OnLButtonUp -> ClosetGoblinCharacterWindow.HotbarPageDownProxy
