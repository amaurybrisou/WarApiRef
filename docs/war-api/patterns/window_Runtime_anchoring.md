# Runtime anchoring

- Category: window
- Confidence: HIGH

## What is this pattern

The runtime anchoring pattern describes dynamically positioning windows in Lua after creation. [WindowAddAnchor](../window_api/functions/window_WindowAddAnchor.md) establishes or modifies window position relative to a parent window, allowing programmatic layout changes after XML static positioning.

## Why it exists

Runtime anchoring enables:
- **Dynamic layouts**: Position windows based on runtime data
- **Responsive design**: Adjust positions based on content size or screen resolution
- **Layout chains**: Multiple windows anchored to each other in dependency chains
- **Repositioning**: Move windows without re-creating them

## When it appears

- **Post-creation**: Immediately after window creation (`CreateWindow` or `CreateWindowFromTemplate`)
- **Dynamic content**: After calculating content dimensions or layout dimensions
- **Visual adjustments**: When repositioning windows based on user interaction or content changes
- **Template instance positioning**: Each template instance repositioned uniquely

## Minimal example

```lua
-- Create window
CreateWindow("MyWindow", true)

-- Position it relative to another window
WindowAddAnchor(
  "MyWindow",          -- Window to position
  "topleft",           -- Point on our window
  "UIParent",          -- Reference window
  "topleft",           -- Point on reference
  100, 100             -- X, Y offset
)
```

## Annotated real example

From AdvancedPetAssist (observed):

```lua
-- Create main window
CreateWindow("APAOptions", true)

-- Position main window at specific screen location
WindowAddAnchor(
  "APAOptions",
  "topleft",
  "UIParent",
  "topleft",
  50,
  50
)

-- Create and position child elements sequentially
for i, element in ipairs(elements) do
  local elementName = "APAOption_" .. i
  
  CreateWindowFromTemplate(elementName, "APAOptionTemplate", "APAOptionsContent")
  
  -- Position each element relative to container
  WindowAddAnchor(
    elementName,
    "topleft",
    "APAOptionsContent",
    "topleft",
    0,
    (i - 1) * 30  -- Stack vertically
  )
end
```

Another pattern from AdvancedRenownTrainer (observed):

```lua
-- Clear and re-add anchors (repositioning)
WindowClearAnchors("PresetWindow")
WindowAddAnchor(
  "PresetWindow",
  "topleft",
  presetContainer,
  "topleft",
  newX,
  newY
)
```

**Key observations**:
- Anchor specifies both source and target "points" (e.g., "topleft", "center")
- X/Y offsets allow fine pixel positioning
- Multiple anchors can position a window (for stretching/scaling)
- Anchors are re-added to reposition dynamically

## Detection signals / evidence

**Observe**:
- `WindowAddAnchor(windowName, sourcePoint, targetWindow, targetPoint, offsetX, offsetY)` calls
- Always appear after window creation
- Source windows just created or in layout pass
- Target windows are parents or reference elements

**Confidence indicators**:
- Window successfully renders at specified position
- Target window exists and is visible
- Multiple anchors chain windows together

## Related patterns

- [Window creation](./window_Window_creation.md) — windows are created then anchored
- [Template instantiation](./window_Template_instantiation.md) — each instance is anchored separately
- [UI creation pattern](./conventions_UI_creation_pattern.md) — broader pattern
- [XML runtime caveats](./conventions_XML_runtime_caveats.md) — edge cases with anchoring

## Common pitfalls

1. **Target window doesn't exist**: Anchoring to parent that hasn't been created yet.
   ```lua
   -- ❌ Wrong: target doesn't exist
   WindowAddAnchor("child", "topleft", "nonExistentParent", "topleft", 0, 0)
   
   -- ✓ Correct: parent created first
   CreateWindow("parent", true)
   CreateWindow("child", true)
   WindowAddAnchor("child", "topleft", "parent", "topleft", 0, 0)
   ```

2. **Wrong anchor point names**: Using invalid point constants (must be cardinal/ordinal: "topleft", "center", etc.).
   ```lua
   -- ❌ Wrong: invalid point name
   WindowAddAnchor("myWin", "top-left", ...)  -- Hyphen invalid
   
   -- ✓ Correct: use valid point
   WindowAddAnchor("myWin", "topleft", ...)
   ```

3. **Not clearing before re-adding**: Calling `WindowAddAnchor` without first calling `WindowClearAnchors` when repositioning.
   ```lua
   -- ❌ Wrong: previous anchor still active
   WindowAddAnchor("win", "topleft", "UIParent", "topleft", 0, 0)
   WindowAddAnchor("win", "topleft", "UIParent", "topleft", 100, 0)  -- May have both anchors
   
   -- ✓ Correct: clear before re-anchoring
   WindowAddAnchor("win", "topleft", "UIParent", "topleft", 0, 0)
   WindowClearAnchors("win")
   WindowAddAnchor("win", "topleft", "UIParent", "topleft", 100, 0)
   ```

## Involved APIs

- [Window](../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type
- [WindowAddAnchor](../window_api/functions/window_WindowAddAnchor.md) (HIGH 100/100) - Window Function
- [Anchor](../xml/element_types/element_Anchor.md) (MEDIUM 55/100) - XML Element Type
