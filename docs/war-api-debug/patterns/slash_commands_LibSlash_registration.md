# LibSlash registration

- Category: slash_commands
- Confidence: MEDIUM

## Description

Observed slash commands being registered through the shared LibSlash table.

## Involved APIs

- [LibSlash](../globals/tables/table_LibSlash.md) (MEDIUM 55/100) - Global Table
- [LibSlash.RegisterSlashCmd](../globals/functions/global_LibSlash.RegisterSlashCmd.md) (MEDIUM 55/100) - Global Function

## Flow Diagram

```text
LibSlash.RegisterSlashCmd <-> WindowSetParent
```

## Example Code

```lua
CM_ClosetGoblin: LibSlash.RegisterSlashCmd("cg", ClosetGoblin.OnSlashCommand)
```

## Evidence

- CM_ClosetGoblin: LibSlash.RegisterSlashCmd("cg", ClosetGoblin.OnSlashCommand)
