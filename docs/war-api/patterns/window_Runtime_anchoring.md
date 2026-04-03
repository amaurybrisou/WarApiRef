# Runtime anchoring

- Category: window
- Confidence: HIGH

## Description

Observed layouts being finalized in Lua by clearing and re-adding anchors after window creation.

## Involved APIs

- [Window](../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type
- [WindowAddAnchor](../window_api/functions/window_WindowAddAnchor.md) (HIGH 100/100) - Window Function

## Flow Diagram

```text
OnLButtonUp
  -> ui: Button, ListBox, Window
```

## Example Code

```lua
InfoScroller: WindowAddAnchor(WindowName.."Image", anchor, WindowName, anchor, T_X, T_Y)
```

## Evidence

- InfoScroller: WindowAddAnchor(WindowName.."Image", anchor, WindowName, anchor, T_X, T_Y)
- InfoScroller: WindowAddAnchor(self.name, pointOnAnchor, anchorWindow, pointOnSelf, xOffset, yOffset)
- Moth: WindowAddAnchor("Moth", "bottomleft", "CursorWindow", "topleft", 0, 0)
- Moth: WindowAddAnchor("MothHealthBar", "bottomleft", "MothBackground", "topleft", 0, yOffset)
- PartyCast: WindowAddAnchor(self.name, pointOnAnchor, anchorWindow, pointOnSelf, xOffset, yOffset)
- PartyCast: WindowAddAnchor("PartyCastWindow"..i, Frame_Anchor, "PartyCastWindow_Dynamic"..i, Frame_Anchor, 0, PartyCast.Settings.Offset)
