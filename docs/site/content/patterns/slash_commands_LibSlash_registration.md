# LibSlash registration

- Category: slash_commands
- Confidence: HIGH

## What is this pattern

The LibSlash registration pattern describes how addons register slash commands through the shared [LibSlash](../globals/tables/table_LibSlash.md) library. Addons call [LibSlash.RegisterSlashCmd](../globals/functions/global_LibSlash.RegisterSlashCmd.md) to register handlers for player-typed commands like `/apa`, `/aura`, `/bagomatic`, etc.

## Why it exists

Centralized slash command handling through LibSlash provides:
- **Unified interface**: Single library manages all addon slash commands
- **Easy parsing**: Library handles command parsing and argument splitting
- **No conflicts**: Central registry prevents command name collisions
- **Convention**: All addons use same registration pattern

## When it appears

- **Addon initialization**: Slash commands registered during addon load
- **Configuration access**: Players invoke commands to access addon UI or settings
- **Debugging**: Commands can dump state, run diagnostics, reload specific systems

## Minimal example

```lua
function MyAddon.OnSlashCommand(input)
  print("Slash command received: " .. input)
end

LibSlash.RegisterSlashCmd(
  "myaddon",
  MyAddon.OnSlashCommand
)

-- Now player can type: /myaddon <arguments>
```

## Annotated real example

From AdvancedPetAssist (observed):

```lua
-- Register slash command with function reference
LibSlash.RegisterSlashCmd(
  "apa",
  function(input)
    APA.SlashHandler(input)
  end
)

-- Handler function parses the command
function APA.SlashHandler(input)
  if input == "" then
    -- No arguments: open options window
    CreateWindow("APAOptions", true)
  elseif input == "reset" then
    -- Reset settings
    APA:ResetSettings()
  else
    -- Show usage
    print("/apa - Open options")
    print("/apa reset - Reset settings")
  end
end
```

Another example from Aura (observed):

```lua
-- Multiple slash command registration for same handler
LibSlash.RegisterSlashCmd("aura", AuraAddon.Slash)
LibSlash.RegisterSlashCmd("auraconfig", AuraAddon.Slash)
LibSlash.RegisterSlashCmd("showaura", AuraAddon.Slash)

-- Single handler handles all variations
function AuraAddon.Slash(input)
  -- Handler logic
end
```

**Key observations**:
- Command names are short and lowercase (e.g., `"apa"`, `"aura"`, `"bagomatic"`)
- Handlers can be named functions or anonymous lambdas
- Multiple command names can route to same handler
- Handlers receive unparsed input string

## Detection signals / evidence

**Observe**:
- `LibSlash.RegisterSlashCmd(commandName, handler)` calls
- Command names are player-friendly (lowercase, short)
- Handler is function reference or lambda
- Handlers parse input string and dispatch subcommands

**Confidence indicators**:
- Command name matches what player would type
- Handler function exists and is callable
- Handler parses common subcommands (help, config, reset)

## Related patterns

- [LibSlash cleanup](./slash_commands_LibSlash_cleanup.md) — unregistration/cleanup
- [Event registration pattern](./conventions_Event_registration_pattern.md) — similar registration pattern

## Common pitfalls

1. **Duplicate command names**: Registering same command name twice (second overwrites first).
   ```lua
   -- ❌ Wrong: second registration overwrites first
   LibSlash.RegisterSlashCmd("cmd", AddonA.OnSlash)
   LibSlash.RegisterSlashCmd("cmd", AddonB.OnSlash)  -- AddonA's handler lost
   
   -- ✓ Correct: use unique command names
   LibSlash.RegisterSlashCmd("addonA", AddonA.OnSlash)
   LibSlash.RegisterSlashCmd("addonB", AddonB.OnSlash)
   ```

2. **Handler not defined**: Registering with function that doesn't exist.
   ```lua
   -- ❌ Wrong: unknown function
   LibSlash.RegisterSlashCmd("cmd", MyAddon.UnknownHandler)
   
   -- ✓ Correct: define handler first
   function MyAddon.SlashHandler(input) end
   LibSlash.RegisterSlashCmd("cmd", MyAddon.SlashHandler)
   ```

3. **Not unregistering on shutdown**: Leaving slash commands registered after addon unload causes stale references.
   ```lua
   -- Initialize
   LibSlash.RegisterSlashCmd("cmd", MyAddon.Handler)
   
   -- Cleanup (missing!)
   -- Should call:
   -- LibSlash.UnregisterSlashCmd("cmd")
   ```

## Involved APIs

- [LibSlash](../globals/tables/table_LibSlash.md) (HIGH 100/100) - Global Table
- [LibSlash.RegisterSlashCmd](../globals/functions/global_LibSlash.RegisterSlashCmd.md) (HIGH 100/100) - Global Function
