# UI creation pattern

- Category: conventions
- Confidence: HIGH

## What is this pattern

The UI creation pattern describes how addons instantiate user interface windows from XML definitions. Windows are created at initialization time using [CreateWindow](../globals/functions/global_CreateWindow.md) or [CreateWindowFromTemplate](../globals/functions/global_CreateWindowFromTemplate.md), then positioned and configured in Lua.

## Why it exists

Separating UI layout (XML) from logic (Lua) provides:
- **Declarative UI**: XML declares structure without imperative Lua code
- **Reusability**: Window templates can be instantiated multiple times
- **Dynamic positioning**: Lua anchors allow runtime layout adjustment
- **Separation of concerns**: Designers modify XML; programmers modify behavior

## When it appears

- **Addon initialization**: Windows created early, before event registration
- **Dynamic UI**: Additional windows created at runtime in response to user actions
- **Template reuse**: Multiple windows instantiated from the same template

## Minimal example

```lua
-- Create a window from XML definition
MyAddon.mainWindow = CreateWindow("MyAddonMainWindow", true)

-- Position it dynamically
WindowAddAnchor("MyAddonMainWindow", "topleft", "UIParent", "topleft", 100, 100)
```

## Annotated real example

From AdvancedPetAssist (observed):

```lua
-- Create main options window from XML
CreateWindow("APAOptions", true)

-- Create additional windows from templates
for i, preset in ipairs(presets) do
  local windowName = "APAPresetWindow" .. i
  CreateWindowFromTemplate(
    windowName,
    "APAPresetTemplate",
    "APAOptionsContent"
  )
  
  -- Position each window dynamically
  WindowAddAnchor(
    windowName,
    "topleft",
    "APAOptionsContent",
    "topleft",
    0,
    (i - 1) * 20
  )
end
```

**Key observations**:
- `CreateWindow` loads from XML and makes the window visible
- `CreateWindowFromTemplate` instantiates a template multiple times
- `WindowAddAnchor` applies dynamic positioning ratios

## Detection signals / evidence

**Observe**:
- `CreateWindow(windowName, visible)` or `CreateWindowFromTemplate(name, template, parent)` calls
- Calls always match window names declared in XML
- Followed by `WindowAddAnchor` calls for positioning
- Typically appear in initialization functions, before event registration

**Confidence indicators**:
- Window names match XML definitions
- All created windows are anchored
- Windows created before event handlers register

## Related patterns

- [Runtime anchoring](./window_Runtime_anchoring.md) — dynamic positioning after creation
- [Template instantiation](./window_Template_instantiation.md) — reusing window templates
- [Initialization pattern](./conventions_Initialization_pattern.md) — where creation occurs
- [Window creation](./window_Window_creation.md) — detailed window API

## Common pitfalls

1. **Window not visible**: Creating window with `visible=false` but never calling show.
   ```lua
   -- ❌ Wrong: window created but never shown
   CreateWindow("MyWindow", false)
   
   -- ✓ Correct: make visible or show later
   CreateWindow("MyWindow", true)
   -- or
   CreateWindow("MyWindow", false)
   WindowSetShowing("MyWindow", true)  -- shown later
   ```

2. **Missing anchor after creation**: Window positioned at origin instead of desired location.

3. **Template not found**: Using template name that doesn't match XML definition.

## Involved APIs

- [Window](../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type
- [CreateWindow](../globals/functions/global_CreateWindow.md) (HIGH 75/100) - Global Function
- [CreateWindowFromTemplate](../globals/functions/global_CreateWindowFromTemplate.md) (HIGH 75/100) - Global Function
