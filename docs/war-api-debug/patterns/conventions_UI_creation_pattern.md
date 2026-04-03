# UI creation pattern

- Category: conventions
- Confidence: HIGH

## Description

UI is commonly created from XML, then positioned in Lua through CreateWindow or CreateWindowFromTemplate followed by anchor calls.

## Involved APIs

- [Window](../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type
- [Button](../xml/element_types/element_Button.md) (HIGH 90/100) - XML Element Type

## Flow Diagram

```text
Button <-> Window
```

## Example Code

```lua
Window creation: CM_ClosetGoblin: CreateWindow("yClosetGoblinButton", true)
```

## Evidence

- Window creation: CM_ClosetGoblin: CreateWindow("yClosetGoblinButton", true)
- Window creation: Clock: CreateWindow("ClockWindow", true)
- XML to Lua binding: CM_ClosetGoblin: .OnLButtonUp -> ClosetGoblinZoneWindow.OnClickZoneRow
- XML to Lua binding: CM_ClosetGoblin: .OnRButtonUp -> ClosetGoblinZoneWindow.OnSetZoneRowContextMenu
- XML to Lua binding: CM_ClosetGoblin: .OnLButtonUp -> ClosetGoblinZoneWindow.OnClickZoneRow
- XML to Lua binding: CM_ClosetGoblin: .OnRButtonUp -> ClosetGoblinZoneWindow.OnSetZoneRowContextMenu
- XML to Lua binding: CM_ClosetGoblin: .OnLButtonUp -> ClosetGoblinCharacterWindow.HotbarPageUpProxy
- XML to Lua binding: CM_ClosetGoblin: .OnLButtonUp -> ClosetGoblinCharacterWindow.HotbarPageDownProxy
