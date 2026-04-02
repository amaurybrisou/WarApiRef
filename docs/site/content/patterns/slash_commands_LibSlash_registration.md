# LibSlash registration

- Category: slash_commands
- Confidence: MEDIUM

## Description

Observed slash commands being registered through the shared LibSlash table.

## Involved APIs

- [LibSlash](../globals/tables/table_LibSlash.md) (HIGH 100/100) - Global Table
- [LibSlash.RegisterSlashCmd](../globals/functions/global_LibSlash.RegisterSlashCmd.md) (HIGH 80/100) - Global Function

## Flow Diagram

```text
OnLButtonUp <-> OnLButtonUp
```

## Example Code

```lua
NPC Item Sale Price: LibSlash.RegisterSlashCmd("nisp", function(args)Nisp.SlashHandler(args)end)
```

## Evidence

- NPC Item Sale Price: LibSlash.RegisterSlashCmd("nisp", function(args)Nisp.SlashHandler(args)end)
