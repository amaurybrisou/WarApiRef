# Runtime anchoring

- Category: window
- Confidence: MEDIUM

## Description

Observed layouts being finalized in Lua by clearing and re-adding anchors after window creation.

## Involved APIs

- [Window](../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type
- [WindowAddAnchor](../window_api/functions/window_WindowAddAnchor.md) (HIGH 100/100) - Window Function

## Flow Diagram

```text
WindowAddAnchor <-> WindowClearAnchors
```

## Example Code

```lua
Moth: WindowAddAnchor("Moth", "bottomleft", "CursorWindow", "topleft", 0, 0)
```

## Evidence

- Moth: WindowAddAnchor("Moth", "bottomleft", "CursorWindow", "topleft", 0, 0)
- Moth: WindowAddAnchor("MothHealthBar", "bottomleft", "MothBackground", "topleft", 0, yOffset)
- TidyChat: WindowAddAnchor(c_TEXT_ENTRY_WINDOW.."EntryBox", "topleft", c_TEXT_ENTRY_WINDOW, "topleft", EntryBoxAnchorXOffset, 0)
- TidyChat: WindowAddAnchor(c_TEXT_ENTRY_WINDOW.."EntryBox", "bottomright", c_TEXT_ENTRY_WINDOW, "bottomright", 0, 0)
- TidyChat: WindowAddAnchor(c_TEXT_ENTRY_WINDOW, TEAnchorPoint.."left", TEAnchorRelativeWindow, TEAnchorRelativePoint.."left", 0, 0)
- TidyChat: WindowAddAnchor(c_TEXT_ENTRY_WINDOW, TEAnchorPoint.."right", TEAnchorRelativeWindow, TEAnchorRelativePoint.."right", 0, 0)
