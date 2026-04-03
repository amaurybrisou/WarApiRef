# Initialization pattern

- Category: conventions
- Confidence: HIGH

## What is this pattern

The initialization pattern describes the predictable sequence through which addons transition from manifest loading into functional state: manifest parsing → initialization hooks → XML UI creation → event registration → runtime readiness.

## Why it exists

Addons must initialize in a specific order to satisfy dependencies: plugins load first, initialization runs before UI is visible, and event handlers are registered only after core state exists. This sequencing prevents race conditions and uninitialized state errors.

## When it appears

During addon load and startup:
- **Manifest load**: Engine loads `.xml` and parses addon metadata
- **Initialize hook**: Addon code runs, typically establishing state and creating UI
- **XML instantiation**: UI windows are created from XML definitions
- **Event registration**: Handlers are registered for runtime state changes
- **Ready**: Addon is fully functional with all UI visible and event loops active

## Minimal example

```lua
MyAddon = {}

function MyAddon:Initialize()
  self.state = {}
  self.window = CreateWindow("MyAddonWindow", true)
end

MyAddon:Initialize()
```

## Annotated real example

From AdvancedPetAssist (observed):

```lua
-- Initialize addon state and windows during load
function APAFollowTargetHUD()
  -- Create state tables (state management pattern)
  APA_Settings = APA_Settings or {}
  
  -- Create UI from XML templates (window creation pattern)
  CreateWindow("APAFollowTargetHUDWindow", true)
  CreateWindow("APAInstantOnlyHUDWindow", true)
  
  -- Register event handlers (event registration pattern)
  RegisterEventHandler(SystemData.Events.LOADING_END, "APA.OnReload")
end
```

**Key observations**:
- State is initialized before UI creation
- Windows are created early in the sequence
- Event handlers registered near the end, after state and UI exist

## Detection signals / evidence

**Observe**:
- Addon initialization happens in a single entry function (often unqualified or named `Initialize`)
- State tables created early in the function
- UI window creation via `CreateWindow` or `CreateWindowFromTemplate`
- Event registration via `RegisterEventHandler` near the end
- Sequence is: state → UI → events

**Confidence indicators**:
- Initialization order is consistent across addons
- No events registered before state initialization
- All UI created before runtime hooks are attached

## Related patterns

- [Event registration pattern](./conventions_Event_registration_pattern.md) — where handlers are registered
- [State management pattern](./conventions_State_management_pattern.md) — where state is initialized
- [UI creation pattern](./conventions_UI_creation_pattern.md) — where windows are created
- [Window creation](./window_Window_creation.md) — specific window instantiation

## Common pitfalls

1. **Registering handlers before state**: Event handlers firing before addon state is ready causes nil reference errors.
   ```lua
   -- ❌ Wrong: handler registered before state
   RegisterEventHandler(SystemData.Events.LOADING_END, "MyAddon.OnLoad")
   MyAddon.state = {}
   
   -- ✓ Correct: state first, then handlers
   MyAddon.state = {}
   RegisterEventHandler(SystemData.Events.LOADING_END, "MyAddon.OnLoad")
   ```

2. **Creating UI in wrong context**: Attempting window creation outside initialization hooks may fail if XML is not yet loaded.

3. **Missing error handling**: No try/catch around initialization allows single initialization failure to cascade into runtime errors.

## Involved APIs

- [SystemData.Events.PLAYER_COMBAT_FLAG_UPDATED](../events/game_events/game_event_SystemData.Events.PLAYER_COMBAT_FLAG_UPDATED.md) (HIGH 100/100) - Game Event
- [SystemData.Events.UPDATE_PROCESSED](../events/game_events/game_event_SystemData.Events.UPDATE_PROCESSED.md) (HIGH 100/100) - Game Event
