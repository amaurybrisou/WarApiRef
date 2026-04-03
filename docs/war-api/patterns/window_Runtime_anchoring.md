# Runtime anchoring

- Category: window
- Confidence: HIGH

## Description

Observed layouts being finalized in Lua by clearing and re-adding anchors after window creation.

## Involved APIs

- [Window](../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type
- [WindowAddAnchor](../window_api/functions/window_WindowAddAnchor.md) (HIGH 100/100) - Window Function
- [Anchor](../xml/element_types/element_Anchor.md) (MEDIUM 55/100) - XML Element Type

## Flow Diagram

```text
OnLButtonUp
  -> ui: Button, ColorPicker, DynamicImage, Label, ListBox, MapDisplay, Window
```

## Example Code

```lua
Ace: WindowAddAnchor(self.name, pointOnAnchor, anchorWindow, pointOnSelf, xOffset, yOffset)
```

## Evidence

- Ace: WindowAddAnchor(self.name, pointOnAnchor, anchorWindow, pointOnSelf, xOffset, yOffset)
- AdvancedPetAssist: WindowAddAnchor(name, "topleft", "APAOptionsContent", "topleft", x, y)
- AdvancedPetAssist: WindowAddAnchor(name, "topleft", "APAOptionsContent", "topleft", x, y)
- AdvancedRenownTrainer: WindowAddAnchor(t.windowName, t.relativePoint, t.relativeTo, t.point, t.x, t.y)
- AdvancedRenownTrainer: WindowAddAnchor(t.windowName, t.relativePoint, t.relativeTo, t.point, t.x, t.y)
- AnywhereTrainerAdditions: WindowAddAnchor(tab.Name, tab.Anchor.Point, tab.Anchor.RelativeTo, tab.Anchor.RelativePoint, tab.Anchor.X, tab.Anchor.Y)
