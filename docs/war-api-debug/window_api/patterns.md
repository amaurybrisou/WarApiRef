# Window Patterns

## Window creation

- Confidence: MEDIUM

- Description: Observed top-level UI windows being created from XML definitions at initialization time.

- Evidence:

- CM_ClosetGoblin: CreateWindow("yClosetGoblinButton", true)
  - Clock: CreateWindow("ClockWindow", true)

## XML to Lua binding

- Confidence: HIGH

- Description: Observed XML handler events routing directly into addon Lua functions through shared engine bindings.

- Evidence:

- CM_ClosetGoblin: .OnLButtonUp -> ClosetGoblinZoneWindow.OnClickZoneRow
  - CM_ClosetGoblin: .OnRButtonUp -> ClosetGoblinZoneWindow.OnSetZoneRowContextMenu
  - CM_ClosetGoblin: .OnLButtonUp -> ClosetGoblinZoneWindow.OnClickZoneRow
  - CM_ClosetGoblin: .OnRButtonUp -> ClosetGoblinZoneWindow.OnSetZoneRowContextMenu
  - CM_ClosetGoblin: .OnLButtonUp -> ClosetGoblinCharacterWindow.HotbarPageUpProxy
  - CM_ClosetGoblin: .OnLButtonUp -> ClosetGoblinCharacterWindow.HotbarPageDownProxy
