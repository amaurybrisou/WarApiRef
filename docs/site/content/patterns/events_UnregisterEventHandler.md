# UnregisterEventHandler

- Category: events
- Confidence: HIGH

## What is this pattern

The UnregisterEventHandler pattern describes removing previously registered event handlers via [UnregisterEventHandler](../globals/functions/global_UnregisterEventHandler.md). This is the cleanup counterpart to [RegisterEventHandler](../globals/functions/global_RegisterEventHandler.md), ensuring handlers don't persist after addon shutdown.

## Why it exists

Unregistering handlers:
- Prevents stale callbacks after addon unload
- Avoids memory leaks (orphaned GC roots)
- Allows clean addon reload without duplicate handlers
- Implements proper resource lifecycle management

## When it appears

- **Addon shutdown**: When addon is disabled or unloaded
- **Interface reload**: When player invokes `/reloadui`
- **Cleanup**: During teardown procedures

## Minimal example

```lua
-- Register during initialization
function MyAddon:Initialize()
  RegisterEventHandler(SystemData.Events.UPDATE_PROCESSED, "MyAddon.OnUpdate")
end

-- Unregister during cleanup
function MyAddon:Cleanup()
  UnregisterEventHandler(SystemData.Events.UPDATE_PROCESSED, "MyAddon.OnUpdate")
end
```

## Annotated real example

From Ace (observed):

```lua
-- Register core update handler
RegisterEventHandler(
  SystemData.Events.UPDATE_PROCESSED,
  "AceAddon_OnUpdate_DONOTTOUCH"
)

-- ... later, during teardown ...

-- Unregister with matching event and handler name
UnregisterEventHandler(
  SystemData.Events.UPDATE_PROCESSED,
  "AceAddon_OnUpdate_DONOTTOUCH"
)
```

Another example from AdvancedRenownTrainer:

```lua
-- Cleanup multiple handlers
UnregisterEventHandler(
  SystemData.Events.PLAYER_CAREER_CATEGORY_UPDATED,
  "AdvancedRenownTraining.CreateDataTable"
)
```

**Key observations**:
- Event ID and handler name must exactly match the registration
- All registered handlers should have corresponding unregistration calls
- Cleanup typically happens at addon shutdown or interface reload

## Detection signals / evidence

**Observe**:
- `UnregisterEventHandler(eventId, "Module.Handler")` calls
- Event ID and handler name match a prior `RegisterEventHandler` call
- Unregistration occurs during teardown/cleanup paths
- No subsequent handler invocations after unregistration

**Confidence indicators**:
- Event/handler pair matches registered pairs
- Unregistration is paired with registration (no orphans)
- Cleanup happens at appropriate lifecycle points

## Related patterns

- [Event registration pattern](./conventions_Event_registration_pattern.md) — the overall pattern
- [RegisterEventHandler](./events_RegisterEventHandler.md) — paired registration

## Common pitfalls

1. **Mismatched event/handler**: Event ID or handler name doesn't exactly match registration call.
   ```lua
   -- ❌ Wrong: handler name doesn't match
   RegisterEventHandler(SystemData.Events.UPDATE_PROCESSED, "MyAddon.OnUpdate")
   UnregisterEventHandler(SystemData.Events.UPDATE_PROCESSED, "MyAddon.OnFrameUpdate")
   
   -- ✓ Correct: exact match
   RegisterEventHandler(SystemData.Events.UPDATE_PROCESSED, "MyAddon.OnUpdate")
   UnregisterEventHandler(SystemData.Events.UPDATE_PROCESSED, "MyAddon.OnUpdate")
   ```

2. **Skipped unregistration**: Removing handler registration call but forgetting unregistration.

3. **Double unregistration**: Unregistering same handler twice causes errors.

## Involved APIs

- [UnregisterEventHandler](../globals/functions/global_UnregisterEventHandler.md) (HIGH 93/100) - Global Function
- [SystemData.Events.UPDATE_PROCESSED](../events/game_events/game_event_SystemData.Events.UPDATE_PROCESSED.md) (HIGH 100/100) - Game Event
