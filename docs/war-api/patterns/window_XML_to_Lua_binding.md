# XML to Lua binding

- Category: window
- Confidence: HIGH

## What is this pattern

The XML to Lua binding pattern (window scope) describes how window-scoped XML event handler attributes connect to Lua functions. Unlike global [RegisterEventHandler](../globals/functions/global_RegisterEventHandler.md), XML handlers execute in window context and are automatically cleaned up when windows are destroyed.

## Why it exists

Window-scoped handlers provide:
- **Automatic lifecycle**: Handlers tied to window; destroyed with window
- **No memory leaks**: No manual unregistration needed
- **Local coupling**: Handler name visible in XML alongside window definition
- **Window context**: Handlers receive window as implicit context

## When it appears

- **Button clicks**: `OnLButtonUp`, `OnRButtonUp` handlers
- **Selection changes**: `OnSelChanged` for combobox, listbox
- **Window lifecycle**: `OnInitialize`, `OnHidden`, `OnShown`
- **Keyboard input**: `OnKeyEnter`, `OnKeyEscape`

## Minimal example

```xml
<Window name="MyWindow" OnInitialize="MyAddon.OnWindowInit">
  <Button name="OKButton" OnLButtonUp="MyAddon.OnOKClicked" />
</Window>
```

```lua
function MyAddon.OnWindowInit()
  print("Window initialized")
end

function MyAddon.OnOKClicked()
  print("OK button clicked")
end
```

## Annotated real example

From AdvancedPetAssist (observed):

```xml
<!-- Multiple handlers on window and child elements -->
<Window name="APAOptions" OnInitialize="APAGui.OnOptionsInit">
  <ComboBox name="APAComboAttackBind"
            OnSelChanged="APAGui.OnComboChanged" />
  <ComboBox name="APAComboCastDelay"
            OnSelChanged="APAGui.OnComboChanged" />
  <Button name="CloseButton"
          OnLButtonUp="APAGui.OnCloseClicked" />
</Window>
```

```lua
-- Single handler processes events from multiple UI elements
function APAGui.OnComboChanged()
  -- Handler can identify which combobox fired by checking context
  local selectedIndex = ComboBoxGetSelectedIndex(...)
  APA.ApplyComboSetting(selectedIndex)
end

function APAGui.OnCloseClicked()
  WindowSetShowing("APAOptions", false)
end

function APAGui.OnOptionsInit()
  -- Populate initial values
  PopulateComboBoxes()
end
```

**Key observations**:
- Multiple sibling handlers can be declared on same XML element
- Same handler function can be reused across different UI elements
- Handlers automatically cleaned up when window is destroyed
- No manual `UnregisterEventHandler` needed for window-scoped handlers

## Detection signals / evidence

**Observe**:
- XML `<*>` elements declare `On<EventName>=` attributes
- Attribute values are qualified Lua function names
- Corresponding Lua functions are defined in addon
- Handlers invoke when events fire (no explicit registration needed)
- Handlers automatically disappear when window destroyed

**Confidence indicators**:
- Handler name matches XML attribute and Lua function
- Handler executes when event occurs
- No orphaned handlers after window destruction

## Related patterns

- [XML to Lua binding pattern (conventions scope)](./conventions_XML_to_Lua_binding_pattern.md) — detailed pattern overview
- [Window event hooks](../events/events_Window_event_hooks.md) — which events are available
- [Window creation](./window_Window_creation.md) — windows are created with handlers active

## Common pitfalls

1. **Unqualified handler names**: Using local function name instead of module path.
   ```xml
   <!-- ❌ Wrong: won't resolve -->
   <Button OnLButtonUp="OnClick" />
   
   <!-- ✓ Correct: use module path -->
   <Button OnLButtonUp="MyAddon.OnClick" />
   ```

2. **Missing function definition**: Handler declared in XML but function not defined.
   ```lua
   -- ❌ Wrong: not in right namespace
   function OnClick() end
   
   -- ✓ Correct: define in addon namespace
   function MyAddon.OnClick() end
   ```

3. **Confusing with global event handlers**: XML handlers are window-scoped and auto-cleaned; global handlers require manual unregister.

## Involved APIs

- [OnSelChanged](../xml/handlers/handler_OnSelChanged.md) (HIGH 100/100) - XML Event
- [OnSelChanged](../events/window_events/window_event_OnSelChanged.md) (HIGH 100/100) - Window Event
- [Window](../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type
