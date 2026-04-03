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
Label <-> OnHyperLinkRButtonUp
```

## Example Code

```lua
InfoScroller: LibSlash.RegisterSlashCmd("infoscroller", function(input)InfoScroller_config.Slash(input)end)
```

## Evidence

- InfoScroller: LibSlash.RegisterSlashCmd("infoscroller", function(input)InfoScroller_config.Slash(input)end)
- InfoScroller: LibSlash.RegisterSlashCmd("info", function(input)InfoScroller_config.Slash(input)end)
- NPC Item Sale Price: LibSlash.RegisterSlashCmd("nisp", function(args)Nisp.SlashHandler(args)end)
- PartyCast: LibSlash.RegisterSlashCmd("pc", function(input)PartyCast.Command(input)end)
- PartyCast: LibSlash.RegisterSlashCmd("partycast", function(input)PartyCast.Command(input)end)
- Soloq: LibSlash.RegisterSlashCmd("soloq", function(args)Soloq.SlashCmd(args)end)
