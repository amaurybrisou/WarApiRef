# Game event fan-in

- Category: events
- Confidence: HIGH

## What is this pattern

The Game event fan-in pattern describes the observed convergence of many addons on a small set of frequently-monitored game events. Events like `LOADING_END`, `RELOAD_INTERFACE`, `UPDATE_PROCESSED`, and player state changes (`PLAYER_COMBAT_FLAG_UPDATED`, `PLAYER_CAREER_CATEGORY_UPDATED`) are registered by many addons for core initialization and reactive updates.

## Why it exists

This pattern emerges naturally because:
- **Initialization**: All addons need `LOADING_END` to initialize after game loads
- **Reload safety**: `RELOAD_INTERFACE` allows clean teardown before UI reload
- **Player state**: Core events track combat, career changes, inventory updates
- **Efficiency**: Centralizing on few events reduces engine event dispatch overhead

## When it appears

- **Addon startup**: All addons register for `LOADING_END`, `UPDATE_PROCESSED`
- **Player tracking**: Multiple addons listen to player state changes
- **Interface reload**: Both addons and engine infrastructure use `RELOAD_INTERFACE`
- **Custom events**: Addons broadcast custom events that fans out to many listeners

## Minimal example

```lua
-- Addon A registers for core event
RegisterEventHandler(SystemData.Events.LOADING_END, "AddonA.OnLoad")

-- Addon B registers for same event
RegisterEventHandler(SystemData.Events.LOADING_END, "AddonB.OnLoad")

-- Both handlers fire when engine broadcasts LOADING_END
```

## Annotated real example

Observed across multiple addons:

```lua
-- AdvancedRenownTrainer
RegisterEventHandler(SystemData.Events.LOADING_END, "AdvancedRenownTraining.OnReload")
RegisterEventHandler(SystemData.Events.PLAYER_CAREER_CATEGORY_UPDATED, "AdvancedRenownTraining.CreateDataTable")
RegisterEventHandler(SystemData.Events.RELOAD_INTERFACE, "AdvancedRenownTraining.OnReload")

-- AdvancedPetAssist
RegisterEventHandler(SystemData.Events.LOADING_END, LOADING_END_HANDLER)

-- Multiple addons converge on same events
-- LOADING_END fires -> calls AdvancedRenownTraining.OnReload, LOADING_END_HANDLER, ... (many others)
```

**Key observations**:
- Same event registered by many independent addons
- Each addon's handler executes synchronously when event fires
- High-traffic events like `UPDATE_PROCESSED` may have dozens of listeners
- Event fan-out is natural consequence of centralized event dispatch

## Detection signals / evidence

**Observe**:
- Event ID appears in multiple addons' `RegisterEventHandler` calls
- Same event type used across different addon categories
- `LOADING_END`, `RELOAD_INTERFACE`, `UPDATE_PROCESSED` are extremely common
- Player state events (`PLAYER_COMBAT_FLAG_UPDATED`, `PLAYER_CAREER_CATEGORY_UPDATED`) appear in many addons

**Confidence indicators**:
- Cross-addon handler registration for same event
- High frequency of event in addon ecosystem
- Event semantic meaning (initialization, player state) explains registrations

## Related patterns

- [RegisterEventHandler](./events_RegisterEventHandler.md) — individual handler registration
- [Event registration pattern](./conventions_Event_registration_pattern.md) — the pattern these events follow
- [Initialization pattern](./conventions_Initialization_pattern.md) — where `LOADING_END` is used

## Common pitfalls

1. **Handler blocking**: One slow handler delays all subsequent handlers in fan-out chain.
   ```lua
   -- ❌ Wrong: expensive operation blocks other handlers
   function SlowAddon.OnUpdate()
     for i = 1, 1000000 do
       expensive_computation()
     end
   end
   
   -- ✓ Correct: defer to background job
   function FastAddon.OnUpdate()
     schedule_async_work()
   end
   ```

2. **Handler order assumptions**: Assuming specific handler execution order in fan-out (order is undefined).

3. **Performance impact**: Registering many handlers for high-frequency events (`UPDATE_PROCESSED`) without care for performance.

## Involved APIs

- [RegisterEventHandler](../globals/functions/global_RegisterEventHandler.md) (HIGH 93/100) - Global Function
- [SystemData.Events.LOADING_END](../events/game_events/game_event_SystemData.Events.LOADING_END.md) (HIGH 100/100) - Game Event
- [SystemData.Events.RELOAD_INTERFACE](../events/game_events/game_event_SystemData.Events.RELOAD_INTERFACE.md) (HIGH 100/100) - Game Event
- [SystemData.Events.UPDATE_PROCESSED](../events/game_events/game_event_SystemData.Events.UPDATE_PROCESSED.md) (HIGH 100/100) - Game Event
- [SystemData.Events.PLAYER_COMBAT_FLAG_UPDATED](../events/game_events/game_event_SystemData.Events.PLAYER_COMBAT_FLAG_UPDATED.md) (HIGH 100/100) - Game Event
- [SystemData.Events.PLAYER_CAREER_CATEGORY_UPDATED](../events/game_events/game_event_SystemData.Events.PLAYER_CAREER_CATEGORY_UPDATED.md) (HIGH 100/100) - Game Event
