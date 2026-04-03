# Window creation

- Category: window
- Confidence: HIGH

## What is this pattern

The window creation pattern describes how addons instantiate UI windows from XML definitions. [CreateWindow](../globals/functions/global_CreateWindow.md) loads the window declaration from XML and makes it visible (or creates it hidden for later display). Windows are the fundamental UI container in the engine.

## Why it exists

Windows are the foundational UI abstraction:
- **Containment**: Windows are parent containers for all child UI elements
- **Lifecycle**: Windows have show/hide, input, and event handling
- **Independent state**: Each window maintains separate render state, focus, anchor configuration
- **Persistence**: Windows persist across gameplay (unless explicitly destroyed)

## When it appears

- **Addon initialization**: Main windows created early, usually before event registration
- **Dynamic UI**: Additional windows created in response to user actions
- **On-demand display**: Windows created early but hidden, then shown on user action

## Minimal example

```lua
-- Create window from XML and make immediately visible
CreateWindow("MyAddonMainWindow", true)

-- Or create hidden and show later
CreateWindow("MyAddonSettingsWindow", false)
WindowSetShowing("MyAddonSettingsWindow", true)  -- Show when needed
```

## Annotated real example

From AdvancedPetAssist (observed):

```lua
-- Create main options window visible
CreateWindow("APAOptions", true)

-- Create HUD windows (head-up displays) for gameplay
CreateWindow("APAFollowTargetHUD", true)
CreateWindow("APAInstantOnlyHUD", true)

-- Position them immediately
WindowAddAnchor("APAFollowTargetHUD", "topleft", "UIParent", "topleft", 100, 100)
WindowAddAnchor("APAInstantOnlyHUD", "topleft", "UIParent", "topleft", 200, 100)
```

**Key observations**:
- Window names match XML definitions
- Visibility passed in `CreateWindow` call (true = show immediately)
- Positioning happens immediately after creation via anchors

## Detection signals / evidence

**Observe**:
- `CreateWindow(windowName, visible)` calls matching XML window definitions
- Windows always followed by positional anchoring
- Typically appear in addon initialization code
- Window names must exactly match `<Window name=` in XML

**Confidence indicators**:
- Window name matches XML definition
- Window successfully renders (visible or can be shown)
- Window accepts input and events as expected

## Related patterns

- [UI creation pattern](./conventions_UI_creation_pattern.md) — broader pattern encompassing window creation
- [Runtime anchoring](./window_Runtime_anchoring.md) — positioning windows after creation
- [Template instantiation](./window_Template_instantiation.md) — creating windows from templates
- [Initialization pattern](./conventions_Initialization_pattern.md) — when windows are created

## Common pitfalls

1. **Window name mismatch**: Using different name in CreateWindow than XML declaration.
   ```lua
   -- ❌ Wrong: name doesn't match XML
   CreateWindow("MyWindow", true)  -- but XML has <Window name="MyAddonWindow">
   
   -- ✓ Correct: exact match
   CreateWindow("MyAddonWindow", true)
   ```

2. **Missing anchor after creation**: Window created but never positioned (appears at origin 0,0).
   ```lua
   -- ❌ Wrong: created but not positioned
   CreateWindow("MyWindow", true)
   
   -- ✓ Correct: position immediately
   CreateWindow("MyWindow", true)
   WindowAddAnchor("MyWindow", "topleft", "UIParent", "topleft", 100, 100)
   ```

3. **Creating before XML loads**: Attempting window creation outside initialization when XML not yet parsed.

## Involved APIs

- [Window](../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type
- [CreateWindow](../globals/functions/global_CreateWindow.md) (HIGH 75/100) - Global Function
