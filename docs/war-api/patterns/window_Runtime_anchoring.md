# Runtime anchoring

- Category: window
- Confidence: HIGH

## Description

Observed layouts being finalized in Lua by clearing and re-adding anchors after window creation.

## Involved APIs

- [Anchor](../xml/element_types/element_Anchor.md) (HIGH 100/100) - XML Element Type
- [Window](../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type
- [WindowAddAnchor](../window_api/functions/window_WindowAddAnchor.md) (HIGH 100/100) - Window Function
- [windowName](../events/game_events/game_event_windowName.md) (MEDIUM 43/100) - Game Event

## Flow Diagram

```text
windowName
  -> handlers: windowName
```

## Example Code

```lua
Ace: WindowAddAnchor(self.name, pointOnAnchor, anchorWindow, pointOnSelf, xOffset, yOffset)
```

## Evidence

- Ace: WindowAddAnchor(self.name, pointOnAnchor, anchorWindow, pointOnSelf, xOffset, yOffset)
- ActionBarHide: WindowAddAnchor(self.name, pointOnAnchor, anchorWindow, pointOnSelf, xOffset, yOffset)
- ActionFraction: WindowAddAnchor(windowName, "center", "Root", "center", 0, 0)
- ActionFraction: WindowAddAnchor(windowName, "center", "PlayerWindowStatusContainerAPPercentBar", "center", 2, 6)
- ActionFraction: WindowAddAnchor(windowName, "center", "Root", "center", 0, 0)
- AdjustTheTip: WindowAddAnchor(c_HEALTH_BAR_CONTAINER, "bottomright", "MouseOverTargetUnitWindow", "bottomright", -10, -10)
