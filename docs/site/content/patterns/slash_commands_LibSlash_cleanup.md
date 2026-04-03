# LibSlash cleanup

- Category: slash_commands
- Confidence: HIGH

## What is this pattern

The LibSlash cleanup pattern describes unregistering previously registered slash commands via [LibSlash.UnregisterSlashCmd](../globals/functions/global_LibSlash.UnregisterSlashCmd.md). This is the cleanup counterpart to [LibSlash.RegisterSlashCmd](../globals/functions/global_LibSlash.RegisterSlashCmd.md), ensuring slash commands don't persist after addon shutdown.

## Why it exists

Unregistering slash commands:
- Prevents stale command handlers after addon unload
- Allows clean addon reload without duplicate commands
- Implements proper resource lifecycle management
- Avoids runtime errors from missing handlers

## When it appears

- **Addon shutdown**: When addon is disabled or unloaded
- **Interface reload**: When player invokes `/reloadui`
- **Cleanup phase**: During telemetry or cleanup batch operations

## Minimal example

```lua
-- Register during initialization
function MyAddon:Initialize()
  LibSlash.RegisterSlashCmd("mycommand", "MyAddon.OnSlash")
end

-- Unregister during cleanup
function MyAddon:Cleanup()
  LibSlash.UnregisterSlashCmd("mycommand")
end
```

## Annotated real example

From DAoCBuff (observed):

```lua
-- Register multiple slash commands
LibSlash.RegisterSlashCmd("daocbuff", DAoCBuff.OnSlash)
LibSlash.RegisterSlashCmd("resetdaocbuff", DAoCBuff.OnReset)

-- ... later, during teardown ...

-- Unregister each command individually
LibSlash.UnregisterSlashCmd("daocbuff")
LibSlash.UnregisterSlashCmd("resetdaocbuff")
```

Another example from WhoHealedMe (observed):

```lua
-- Register
LibSlash.RegisterSlashCmd("whm", WhoHealedMe.OnSlash)

-- Unregister on shutdown
LibSlash.UnregisterSlashCmd("whm")
```

**Key observations**:
- Command name used in unregister must exactly match registered command
- Each registered command is unregistered individually
- Unregistration typically happens during addon teardown or reload

## Detection signals / evidence

**Observe**:
- `LibSlash.UnregisterSlashCmd(commandName)` calls
- Command names match prior `LibSlash.RegisterSlashCmd` calls
- Unregistration occurs during cleanup/shutdown paths
- No subsequent command invocations after unregistration

**Confidence indicators**:
- Command name matches registered command
- Unregistration is paired with registration (no orphans)
- Cleanup happens at appropriate lifecycle points

## Related patterns

- [LibSlash registration](./slash_commands_LibSlash_registration.md) — paired registration
- [Event registration pattern](./conventions_Event_registration_pattern.md) — similar unregister pattern

## Common pitfalls

1. **Mismatched command name**: Unregistering with different name than registered.
   ```lua
   -- ❌ Wrong: command name doesn't match
   LibSlash.RegisterSlashCmd("cmd", MyAddon.Handler)
   LibSlash.UnregisterSlashCmd("command")  -- Typo!
   
   -- ✓ Correct: exact match
   LibSlash.RegisterSlashCmd("cmd", MyAddon.Handler)
   LibSlash.UnregisterSlashCmd("cmd")
   ```

2. **Skipped unregistration**: Removing registration call but forgetting cleanup.
   ```lua
   -- ❌ Wrong: never unregisters
   LibSlash.RegisterSlashCmd("cmd", MyAddon.Handler)
   -- Missing: LibSlash.UnregisterSlashCmd("cmd")
   
   -- ✓ Correct: pair register and unregister
   LibSlash.RegisterSlashCmd("cmd", MyAddon.Handler)
   -- ... later ...
   LibSlash.UnregisterSlashCmd("cmd")
   ```

3. **Double unregistration**: Unregistering same command twice causes errors.

## Involved APIs

- [LibSlash](../globals/tables/table_LibSlash.md) (HIGH 100/100) - Global Table
- [LibSlash.UnregisterSlashCmd](../globals/functions/global_LibSlash.UnregisterSlashCmd.md) (HIGH 100/100) - Global Function
