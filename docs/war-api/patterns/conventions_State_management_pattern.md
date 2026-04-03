# State management pattern

- Category: conventions
- Confidence: MEDIUM

## What is this pattern

The state management pattern establishes how addons store persistent and runtime configuration state. State is rooted in addon-owned global tables (often registered with saved variables systems), initialized during addon load, and then accessed throughout runtime.

## Why it exists

Addons need to maintain configuration across sessions (saved variables) and runtime state (in-memory globals). By centralizing state in named globals, addons can:
- Recover settings after reload or restart
- Share state between UI elements and event handlers
- Expose configuration to debugging and the console

## When it appears

- **Initialization**: State tables created during addon load
- **Configuration**: Settings loaded from saved variables or defaults
- **Runtime**: State read and updated by UI handlers and event callbacks
- **Persistence**: State written back to saved variables on change or shutdown

## Minimal example

```lua
MyAddon = {}
MyAddon.state = {}
MyAddon.settings = {
  enabled = true,
  windowPos = { x = 100, y = 100 }
}
```

## Annotated real example

From AdvancedPetAssist (observed):

```lua
-- Root addon namespace contains mutable settings
APA_Settings = APA_Settings or {
  targetingMode = "auto",
  autoReattack = true,
  delayMs = 500
}

-- Separate table for runtime state (not saved)
APA.currentPet = nil
APA.lastCastTime = 0
APA.isEnabled = true
```

**Key observations**:
- Settings table (`APA_Settings`) uses global naming convention for saved variables
- Runtime state uses local namespaces (`APA.*`)
- Configuration is initialized `or {}`  to support upgrade paths

## Detection signals / evidence

**Observe**:
- Top-level global tables named after addon (e.g., `AddonName_Settings`, `AddonName.state`)
- Tables initialized with `or {}` syntax for safe defaults
- Settings accessed throughout UI handlers and event callbacks
- Runtime state stored in subordinate tables

**Confidence indicators**:
- Global names follow addon naming conventions
- State is initialized before event handlers fire
- Same state variables referenced in multiple handler functions

## Related patterns

- [Initialization pattern](./conventions_Initialization_pattern.md) — where state is created
- [UI creation pattern](./conventions_UI_creation_pattern.md) — UI reads from state
- [Event registration pattern](./conventions_Event_registration_pattern.md) — handlers access state

## Common pitfalls

1. **Missing safe initialization**: State accessed before `or {}` initialization leads to nil errors.
   ```lua
   -- ❌ Wrong: assumes state already exists
   MyAddon.settings.enabled = true
   
   -- ✓ Correct: safe initialization
   MyAddon.settings = MyAddon.settings or {}
   MyAddon.settings.enabled = true
   ```

2. **Mixing saved and runtime state**: Confusing which data persists across reloads.
   ```lua
   -- ❌ Wrong: runtime cache saved to disk
   SavedAddon.cache = compute_expensive_data()
   
   -- ✓ Correct: cache separate from saved state
   Addon.cache = compute_expensive_data()
   ```

3. **Global namespace pollution**: Using unqualified globals instead of addon-owned tables.

## Involved APIs

- (Global tables with no specific API dependency)
