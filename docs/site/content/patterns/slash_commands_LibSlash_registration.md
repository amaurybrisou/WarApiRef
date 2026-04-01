# LibSlash registration

- Category: slash_commands
- Confidence: HIGH

## Description

Observed slash commands being registered through the shared LibSlash table.

## Involved APIs

- [LibSlash](../globals/tables/table_LibSlash.md) (HIGH 100/100) - Global Table
- [LibSlash.RegisterSlashCmd](../globals/functions/global_LibSlash.RegisterSlashCmd.md) (HIGH 100/100) - Global Function

## Flow Diagram

```text
OnLButtonUp <-> OnLButtonUp
```

## Example Code

```lua
AdvancedPetAssist: LibSlash.RegisterSlashCmd("apa", function(input)APA.SlashHandler(input)end)
```

## Evidence

- AdvancedPetAssist: LibSlash.RegisterSlashCmd("apa", function(input)APA.SlashHandler(input)end)
- Aura: LibSlash.RegisterSlashCmd("aura", AuraAddon.Slash)
- Aura: LibSlash.RegisterSlashCmd("auraconfig", AuraAddon.Slash)
- Aura: LibSlash.RegisterSlashCmd("showaura", AuraAddon.Slash)
- AutoBand: LibSlash.RegisterSlashCmd("autoband", function(msg)AutoBand.parse_cmd(msg)end)
- AutoBand: LibSlash.RegisterSlashCmd("ab", function(msg)AutoBand.parse_cmd(msg)end)
