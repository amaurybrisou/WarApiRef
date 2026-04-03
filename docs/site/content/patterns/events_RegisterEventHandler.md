# RegisterEventHandler

- Category: events
- Confidence: HIGH

## What is this pattern

The RegisterEventHandler pattern describes registering addon-level global event handlers via [RegisterEventHandler](../globals/functions/global_RegisterEventHandler.md). This is complementary to XML-scoped event handlers and allows addons to listen for game and custom events at the global level.

## Why it exists

Global event handlers allow addons to:
- Respond to game-wide events without UI context
- React to player combat state, interface reload, data updates
- Coordinate behavior across multiple windows
- Implement background monitoring and periodic updates

## When it appears

- **Player events**: LOADING_END, PLAYER_COMBAT_FLAG_UPDATED, PLAYER_CAREER_CATEGORY_UPDATED
- **Interface events**: RELOAD_INTERFACE, UPDATE_PROCESSED
- **Custom events**: Addon-defined events broadcast by other addons

## Minimal example

```lua
function MyAddon:Initialize()
  RegisterEventHandler(
    SystemData.Events.LOADING_END,
    "MyAddon.OnGameLoaded"
  )
end

function MyAddon.OnGameLoaded()
  -- Refresh addon data after game loads
end
```

## Annotated real example

From AdvancedRenownTrainer (observed):

```lua
-- Multiple handlers for different event types
RegisterEventHandler(
  SystemData.Events.LOADING_END,
  "AdvancedRenownTraining.OnReload"
)

RegisterEventHandler(
  SystemData.Events.PLAYER_CAREER_CATEGORY_UPDATED,
  "AdvancedRenownTraining.CreateDataTable"
)

RegisterEventHandler(
  SystemData.Events.RELOAD_INTERFACE,
  "AdvancedRenownTraining.OnReload"
)

-- Handlers execute when events fire
function AdvancedRenownTraining.OnReload()
  -- Rebuild data tables
end

function AdvancedRenownTraining.CreateDataTable()
  -- Update on career change
end
```

**Key observations**:
- Multiple handlers can listen to the same event
- Same handler can be registered for different events
- Handlers must be qualified function names

## Detection signals / evidence

**Observe**:
- `RegisterEventHandler(eventId, "Module.Handler")` calls
- Event IDs from `SystemData.Events.*`
- Handler function names as qualified strings
- Matching `UnregisterEventHandler` calls for cleanup

**Confidence indicators**:
- Event ID exists in SystemData.Events
- Handler function is defined and accessible
- Registration/unregistration are paired

## Related patterns

- [Event registration pattern](./conventions_Event_registration_pattern.md) — the overall pattern
- [UnregisterEventHandler](./events_UnregisterEventHandler.md) — paired cleanup
- [BroadcastEvent](./events_BroadcastEvent.md) — triggering events

## Common pitfalls

1. **Unqualified names**: Using local function name instead of module-qualified path.

2. **Orphaned handlers**: Registering without corresponding unregistration leads to memory leaks.

3. **Handler reuse pitfall**: Reusing one handler for multiple event types without considering payload differences.

## Involved APIs

- [RegisterEventHandler](../globals/functions/global_RegisterEventHandler.md) (HIGH 93/100) - Global Function
- [SystemData.Events.LOADING_END](../events/game_events/game_event_SystemData.Events.LOADING_END.md) (HIGH 100/100) - Game Event
- [SystemData.Events.PLAYER_CAREER_CATEGORY_UPDATED](../events/game_events/game_event_SystemData.Events.PLAYER_CAREER_CATEGORY_UPDATED.md) (HIGH 100/100) - Game Event
- [SystemData.Events.RELOAD_INTERFACE](../events/game_events/game_event_SystemData.Events.RELOAD_INTERFACE.md) (HIGH 100/100) - Game Event
