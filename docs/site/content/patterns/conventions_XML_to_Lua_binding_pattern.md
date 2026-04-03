# XML to Lua binding pattern

- Category: conventions
- Confidence: HIGH

## What is this pattern

The XML to Lua binding pattern establishes the direct connection between XML event handler attributes and Lua functions. An XML event attribute (e.g., `OnSelChanged="MyAddon.OnSelectionChanged"`) routes engine UI events directly into Lua function calls, bridging the XML and Lua layers.

## Why it exists

XML declarations of handlers provide:
- **Declarative binding**: Handler mapping visible in XML, not hidden in Lua registration code
- **Type safety**: Handler names checked at XML parse time
- **Coupling**: Clear relationship between UI element and logic

## When it appears

- **UI element events**: Buttons, comboboxes, textboxes declare `On*` handlers in XML
- **Window lifecycle**: Windows declare `OnInitialize`, `OnHidden` handlers
- **Event routing**: Engine events (click, selection change) route to handlers

## Minimal example

```xml
<ComboBox name="MyCombo" OnSelChanged="MyAddon.OnComboChanged" />
```

```lua
function MyAddon.OnComboChanged()
  local selectedIndex = ComboBoxGetSelectedIndex("MyCombo")
  print("Selected index: " .. selectedIndex)
end
```

## Annotated real example

From AdvancedPetAssist (observed):

```xml
<!-- XML declares multiple handler bindings -->
<ComboBox name="APAComboAttackBind"
          OnSelChanged="APAGui.OnComboChanged"
          default_index="0" />

<ComboBox name="APAComboCastDelay"
          OnSelChanged="APAGui.OnComboChanged" />
```

```lua
-- Single handler processes events from multiple UI elements
function APAGui.OnComboChanged()
  -- Get the window that triggered the event
  local sourceWindow = SystemData.UI.ChatWindow  -- or use event context
  
  -- Update addon state based on the selection
  local selectedIndex = ComboBoxGetSelectedIndex(sourceWindow)
  APA.currentSettings[sourceWindow] = selectedIndex
end
```

**Key observations**:
- XML declares handler name as a qualified string (e.g., `"APAGui.OnComboChanged"`)
- Same handler can be reused across multiple UI elements
- Handler function must be defined at addon load time
- Handler is invoked with the window as implicit first argument in some cases

## Detection signals / evidence

**Observe**:
- XML `<Element>` tags include `On<EventName>=` attributes
- Attribute values are quoted strings matching Lua function paths
- Lua functions with matching names exist in the addon namespace
- Event handlers invoked whenever UI element event fires (button click, selection change)

**Confidence indicators**:
- Handler name matches between XML attribute and Lua function definition
- Handler function is callable at runtime (properly qualified)
- Handler is invoked consistently with expected event semantics

## Related patterns

- [Window event hooks](../events/events_Window_event_hooks.md) — which events are available
- [Event registration pattern](./conventions_Event_registration_pattern.md) — global event registration (different from XML handlers)

## Common pitfalls

1. **Unqualified function names**: Handler attribute uses local function name instead of module-qualified path.
   ```xml
   <!-- ❌ Wrong: local function won't resolve -->
   <Button OnLButtonUp="OnClick" />
   
   <!-- ✓ Correct: use qualified path -->
   <Button OnLButtonUp="MyAddon.OnClick" />
   ```

2. **Missing function definition**: Handler specified in XML but function not defined at runtime.
   ```lua
   -- ❌ Wrong: function never defined
   function OnClick() end  -- Wrong: not qualified
   
   -- ✓ Correct: define in addon namespace
   function MyAddon.OnClick() end
   ```

3. **Confusing XML handlers with global events**: XML handlers (routed through UI elements) and global event handlers (via `RegisterEventHandler`) are distinct registration mechanisms.

## Involved APIs

- [OnSelChanged](../xml/handlers/handler_OnSelChanged.md) (HIGH 100/100) - XML Event
- [OnSelChanged](../events/window_events/window_event_OnSelChanged.md) (HIGH 100/100) - Window Event
