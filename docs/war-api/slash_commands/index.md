# Slash Commands

## LibSlash registration

- Confidence: HIGH

- Description: Observed slash commands being registered through the shared LibSlash table.

- Evidence:

- InfoScroller: LibSlash.RegisterSlashCmd("infoscroller", function(input)InfoScroller_config.Slash(input)end)
  - InfoScroller: LibSlash.RegisterSlashCmd("info", function(input)InfoScroller_config.Slash(input)end)
  - NPC Item Sale Price: LibSlash.RegisterSlashCmd("nisp", function(args)Nisp.SlashHandler(args)end)
  - PartyCast: LibSlash.RegisterSlashCmd("pc", function(input)PartyCast.Command(input)end)
  - PartyCast: LibSlash.RegisterSlashCmd("partycast", function(input)PartyCast.Command(input)end)
  - Soloq: LibSlash.RegisterSlashCmd("soloq", function(args)Soloq.SlashCmd(args)end)
