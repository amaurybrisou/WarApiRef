# XML to Lua binding pattern

- Category: conventions
- Confidence: HIGH

## Description

XML handler names map directly to Lua functions and can be cross-checked through the bindings page.

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
- CM_ClosetGoblin: .OnShown -> ClosetGoblinCharacterWindow.OnShow
- CM_ClosetGoblin: .OnHidden -> ClosetGoblinCharacterWindow.OnHide
