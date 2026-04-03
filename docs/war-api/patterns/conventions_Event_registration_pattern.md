# Event registration pattern

- Category: conventions
- Confidence: HIGH

## What is this pattern

The event registration pattern establishes how addons subscribe to and respond to runtime game events via the [RegisterEventHandler](../globals/functions/global_RegisterEventHandler.md) API. It pairs an event identifier (typically from [SystemData.Events](../systemdata/fields/systemdata_SystemData.Events.md)) with a handler function, establishing a reactive contract between the engine and addon code.

## Why it exists

Event-driven architecture allows addons to respond to game state changes without polling. By registering handlers at predictable lifecycle moments, addons can initialize UI, refresh data, or perform cleanup reactively—ensuring responsiveness to player actions, combat state, interface reloads, and other system events.

## When it appears

This pattern appears during addon initialization and teardown:
- **Initialization**: After the addon's core state is loaded, before UI is visible
- **Reactive updates**: Throughout runtime whenever engine events fire (e.g., `LOADING_END`, `PLAYER_CAREER_CATEGORY_UPDATED`)
- **Cleanup**: During addon shutdown or interface reload to prevent stale handlers

## Minimal example

```lua
function MyAddon:OnInitialize()
  RegisterEventHandler(SystemData.Events.LOADING_END, "MyAddon.OnGameLoaded")
end

function MyAddon:OnGameLoaded()
  print("Game loaded, refreshing UI...")
end
```

## Annotated real example

From AdvancedRenownTrainer (observed):

```lua
-- Register handler for career changes (updates UI when player switches careers)
RegisterEventHandler(
  SystemData.Events.PLAYER_CAREER_CATEGORY_UPDATED,
  "AdvancedRenownTraining.CreateDataTable"
)

-- Register handler for interface reload (reset tables when /reloadui is invoked)
RegisterEventHandler(
  SystemData.Events.RELOAD_INTERFACE,
  "AdvancedRenownTraining.OnReload"
)

-- Register handler for initial load (populate initial data at login)
RegisterEventHandler(
  SystemData.Events.LOADING_END,
  "AdvancedRenownTraining.OnReload"
)
```

**Key observations**:
- Handler names are module-qualified (e.g., `"AdvancedRenownTraining.OnReload"`) for proper function lookup
- Multiple handlers can be registered for different event types
- The same handler can be reused for logically similar events

## Detection signals / evidence

**Observe**:
- `RegisterEventHandler(eventId, handlerName)` calls during addon initialization
- Event identifiers always from `SystemData.Events.*` namespace
- Handler functions referenced by qualified name (string or direct reference)
- Matching `UnregisterEventHandler` calls during cleanup paths

**Confidence indicators**:
- Event IDs verified in [SystemData.Events](../systemdata/fields/systemdata_SystemData.Events.md)
- Handler names resolve to top-level addon functions
- Registration and unregistration are paired (no leaks)

## Related patterns

- [RegisterEventHandler](./events_RegisterEventHandler.md) — specific API pattern
- [UnregisterEventHandler](./events_UnregisterEventHandler.md) — paired cleanup  
- [Initialization pattern](./conventions_Initialization_pattern.md) — where registration occurs
- [Game event fan-in](./events_Game_event_fan-in.md) — which events are heavily registered

## Common pitfalls

1. **Unqualified handler names**: Using local functions instead of module-qualified names—handlers will fail to resolve.
   ```lua
   -- ❌ Wrong: function name won't resolve at runtime
   RegisterEventHandler(SystemData.Events.LOADING_END, OnLoad)
   
   -- ✓ Correct: use qualified module path
   RegisterEventHandler(SystemData.Events.LOADING_END, "MyModule.OnLoad")
   ```

2. **Orphaned handlers**: Registering handlers without corresponding `UnregisterEventHandler` calls during teardown—causes memory leaks and stale callbacks.

3. **Handler assumptions**: Assuming handlers fire in isolation—handlers for related events may fire in sequence with shared state changes.

## Involved APIs

- [OnUpdate](../events/window_events/window_event_OnUpdate.md) (HIGH 100/100) - Window Event
- [OnUpdate](../xml/handlers/handler_OnUpdate.md) (HIGH 100/100) - XML Event
- [SystemData.Events.LOADING_END](../events/game_events/game_event_SystemData.Events.LOADING_END.md) (HIGH 100/100) - Game Event
- [SystemData.Events.LOADING_END](../systemdata/fields/systemdata_SystemData.Events.LOADING_END.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_CAREER_CATEGORY_UPDATED](../systemdata/fields/systemdata_SystemData.Events.PLAYER_CAREER_CATEGORY_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_CAREER_CATEGORY_UPDATED](../events/game_events/game_event_SystemData.Events.PLAYER_CAREER_CATEGORY_UPDATED.md) (HIGH 100/100) - Game Event
- [SystemData.Events.RELOAD_INTERFACE](../events/game_events/game_event_SystemData.Events.RELOAD_INTERFACE.md) (HIGH 100/100) - Game Event
- [SystemData.Events.RELOAD_INTERFACE](../systemdata/fields/systemdata_SystemData.Events.RELOAD_INTERFACE.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.UPDATE_PROCESSED](../systemdata/fields/systemdata_SystemData.Events.UPDATE_PROCESSED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.UPDATE_PROCESSED](../events/game_events/game_event_SystemData.Events.UPDATE_PROCESSED.md) (HIGH 100/100) - Game Event
- [RegisterEventHandler](../globals/functions/global_RegisterEventHandler.md) (HIGH 93/100) - Global Function
- [UnregisterEventHandler](../globals/functions/global_UnregisterEventHandler.md) (HIGH 93/100) - Global Function
- [EventHandler](../xml/element_types/element_EventHandler.md) (MEDIUM 45/100) - XML Element Type
