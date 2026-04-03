# Window event hooks

- Category: events
- Confidence: HIGH

## What is this pattern

The Window event hooks pattern describes how XML-declared `On*` handler attributes route engine UI events into Lua functions. Handlers like `OnLButtonUp`, `OnSelChanged`, `OnInitialize` bridge the UI engine and addon code, executing in response to user interactions and window lifecycle events.

## Why it exists

UI event hooks allow addons to:
- Respond to button clicks, selection changes, text input
- Execute initialization code when windows appear
- Handle focus, hover, and input events
- Map UI interactions to addon logic

## When it appears

- **Initialization**: `OnInitialize` when window is created
- **User interaction**: `OnLButtonUp`, `OnRButtonUp` for clicks; `OnSelChanged` for selection
- **Lifecycle**: `OnHidden` when window disappears; `OnShown` when visible
- **Input**: `OnKeyEnter`, `OnKeyEscape` for keyboard

## Minimal example

```xml
<Button name="MyButton" OnLButtonUp="MyAddon.OnButtonClick" />
```

```lua
function MyAddon.OnButtonClick()
  print("Button clicked!")
end
```

## Annotated real example

From various addons (observed):

```xml
<!-- Window with multiple event hooks -->
<Window name="MyWindow" OnInitialize="MyAddon.OnWindowInit" OnHidden="MyAddon.OnWindowHidden">
  <Button name="CloseButton" OnLButtonUp="MyAddon.OnCloseClicked" />
  <ComboBox name="StrategyCombo" OnSelChanged="MyAddon.OnStrategySelected" />
</Window>
```

```lua
function MyAddon.OnWindowInit()
  -- Execute when window is created
  PopulateCombo("StrategyCombo")
end

function MyAddon.OnWindowHidden()
  -- Execute when window is hidden
  SaveWindowState()
end

function MyAddon.OnCloseClicked()
  -- Execute when button is clicked
  WindowSetShowing("MyWindow", false)
end

function MyAddon.OnStrategySelected()
  -- Execute when combo selection changes
  ApplySelectedStrategy()
end
```

**Key observations**:
- Handler names declared in XML as `On<EventName>=` attributes
- Functions must be module-qualified and defined in Lua
- Different event types (lifecycle, input, state change) share same mechanism
- Handlers execute synchronously when event occurs

## Detection signals / evidence

**Observe**:
- XML elements declare `On<EventName>=` attributes
- Attribute values are quoted Lua function paths (e.g., `"Module.Handler"`)
- Corresponding Lua functions exist and are callable
- Handler functions invoked when events occur (button click, window shown, etc.)

**Confidence indicators**:
- Handler name matches XML attribute and Lua function
- Handler executes at appropriate event time
- Multiple different `On*` handlers on same element

## Related patterns

- [XML to Lua binding pattern](./conventions_XML_to_Lua_binding_pattern.md) — general handler binding
- [Event registration pattern](./conventions_Event_registration_pattern.md) — global event registration (different mechanism)

## Common pitfalls

1. **Unqualified handler names**: Using local function name instead of module path.
   ```xml
   <!-- ❌ Wrong: won't resolve -->
   <Button OnLButtonUp="OnClick" />
   
   <!-- ✓ Correct: use module path -->
   <Button OnLButtonUp="MyAddon.OnClick" />
   ```

2. **Missing function definition**: Handler declared in XML but function not defined in Lua.

3. **Confusing with global events**: XML handlers and `RegisterEventHandler` are different registration mechanisms.

4. **Event handler signature mismatch**: Handler expects different number or type of arguments than engine provides.

## Involved APIs

- [OnLButtonUp](../events/window_events/window_event_OnLButtonUp.md) (HIGH 100/100) - Window Event
- [OnInitialize](../events/window_events/window_event_OnInitialize.md) (HIGH 100/100) - Window Event
- [OnHidden](../events/window_events/window_event_OnHidden.md) (HIGH 100/100) - Window Event
- [OnSelChanged](../events/window_events/window_event_OnSelChanged.md) (HIGH 100/100) - Window Event
