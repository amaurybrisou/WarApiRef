# BroadcastEvent

- Category: events
- Confidence: MEDIUM

## What is this pattern

The BroadcastEvent pattern describes how addons send custom or game events to notify other event handlers. [BroadcastEvent](../globals/functions/global_BroadcastEvent.md) invokes all handlers registered for a given event, triggering coordinated responses across the addon ecosystem.

## Why it exists

Broadcasting allows addons to:
- Signal custom state changes to other handlers
- Trigger coordinated behavior across multiple subsystems
- Implement loose coupling (broadcast listeners don't know about each other)
- Follow the same event-driven pattern as engine-fired events

## When it appears

- **Custom events**: Addon-specific state changes (e.g., "buff applied", "strategy changed")
- **Engine event wrapping**: Forwarding game events with extra context
- **Coordination**: Notifying multiple handlers of atomic state changes

## Minimal example

```lua
-- Define a custom event
local MY_CUSTOM_EVENT = "MyAddon_StrategyChanged"

-- Broadcast it
function MyAddon:ChangeStrategy(strategyName)
  self.currentStrategy = strategyName
  BroadcastEvent(MY_CUSTOM_EVENT, strategyName)
end

-- Handlers register to listen
RegisterEventHandler(MY_CUSTOM_EVENT, "OtherAddon.OnStrategyChange")
```

## Annotated real example

From Enemy (observed):

```lua
-- Broadcast custom target event to notify other add-ons
function Enemy:NotifyTargetChange()
  BroadcastEvent(self.TARGET_CHANGED_EVENT)
end

-- Also broadcast engine events with addon context
function Enemy:OnGameEvent(eventId, ...)
  if eventId == SystemData.Events.SCENARIO_START_UPDATING_PLAYERS_STATS then
    -- Forward to listeners with extra context
    BroadcastEvent("Enemy_ScenarioStarting", ...)
  end
end
```

**Key observations**:
- Custom events can be arbitrary strings (convention is `AddonName_EventName`)
- Both custom and engine events can be broadcasted
- Broadcast immediately notifies all registered handlers

## Detection signals / evidence

**Observe**:
- `BroadcastEvent(eventId, ...)` calls with event name and optional payload
- Event can be game event (e.g., `SystemData.Events.*`) or custom string
- Corresponding `RegisterEventHandler` calls with same event ID
- Handlers execute synchronously when event is broadcast

**Confidence indicators**:
- Event is registered and handled elsewhere in codebase
- Broadcast is called during meaningful state changes
- Handlers access the broadcast payload correctly

## Related patterns

- [Event registration pattern](./conventions_Event_registration_pattern.md) — how handlers listen
- [RegisterEventHandler](./events_RegisterEventHandler.md) — related API pattern
- [Game event fan-in](./events_Game_event_fan-in.md) — which events are broadcast

## Common pitfalls

1. **Broadcasting before handlers register**: Broadcasting an event before any handlers are registered means no listeners.
   ```lua
   -- ❌ Wrong: broadcast before handlers listen
   BroadcastEvent("MyEvent", data)
   RegisterEventHandler("MyEvent", "Listener.OnEvent")
   
   -- ✓ Correct: register listeners first
   RegisterEventHandler("MyEvent", "Listener.OnEvent")
   BroadcastEvent("MyEvent", data)
   ```

2. **Synchronous execution surprises**: All handlers execute immediately; handlers block each other.

3. **Event name typos**: Custom event ID in `BroadcastEvent` doesn't match registered handler events.

## Involved APIs

- [BroadcastEvent](../globals/functions/global_BroadcastEvent.md) (HIGH 100/100) - Global Function
- [SystemData.Events.SCENARIO_START_UPDATING_PLAYERS_STATS](../systemdata/fields/systemdata_SystemData.Events.SCENARIO_START_UPDATING_PLAYERS_STATS.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.SCENARIO_STOP_UPDATING_PLAYERS_STATS](../events/game_events/game_event_SystemData.Events.SCENARIO_STOP_UPDATING_PLAYERS_STATS.md) (HIGH 100/100) - Game Event
