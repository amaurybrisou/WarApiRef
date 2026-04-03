# LibSlash cleanup

- Category: slash_commands
- Confidence: HIGH

## Description

Observed slash commands being explicitly unregistered during addon shutdown paths.

## Involved APIs

- [LibSlash](../globals/tables/table_LibSlash.md) (HIGH 100/100) - Global Table
- [LibSlash.UnregisterSlashCmd](../globals/functions/global_LibSlash.UnregisterSlashCmd.md) (HIGH 100/100) - Global Function

## Flow Diagram

```text
OnInitialize -> RegisterEventHandler
```

## Example Code

```lua
ActionFraction: LibSlash.UnregisterSlashCmd("af")
```

## Evidence

- ActionFraction: LibSlash.UnregisterSlashCmd("af")
- ActionFraction: LibSlash.UnregisterSlashCmd("ActionFraction")
- AuctionStats: LibSlash.UnregisterSlashCmd("au")
- AuctionStats: LibSlash.UnregisterSlashCmd("uc")
- AuctionStats: LibSlash.UnregisterSlashCmd("undercut")
- CDown: LibSlash.UnregisterSlashCmd("CDown")
