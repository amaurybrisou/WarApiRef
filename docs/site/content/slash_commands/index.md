# Slash Commands

## LibSlash registration

- Confidence: HIGH

- Description: Observed slash commands being registered through the shared LibSlash table.

- Evidence:

- ActionBarCD: LibSlash.RegisterSlashCmd("abcd", function(args)ActionBarCD.SlashHandler(args)end)
  - ActionBarHide: LibSlash.RegisterSlashCmd("abh", function()ActionBarHide.OptionsWindow()end)
  - AdvancedPetAssist: LibSlash.RegisterSlashCmd("apa", function(input)APA.SlashHandler(input)end)
  - Amethyst: LibSlash.RegisterSlashCmd("amt", function(msg)Amethyst.Slash(msg)end)
  - Amethyst: LibSlash.RegisterSlashCmd("amethyst", function(msg)Amethyst.Slash(msg)end)
  - ArmorGraphicToggle: LibSlash.RegisterSlashCmd("agt", function(msg)ArmorGraphicToggle.Slash(msg)end)

## LibSlash cleanup

- Confidence: HIGH

- Description: Observed slash commands being explicitly unregistered during addon shutdown paths.

- Evidence:

- ActionFraction: LibSlash.UnregisterSlashCmd("af")
  - ActionFraction: LibSlash.UnregisterSlashCmd("ActionFraction")
  - AuctionStats: LibSlash.UnregisterSlashCmd("au")
  - AuctionStats: LibSlash.UnregisterSlashCmd("uc")
  - AuctionStats: LibSlash.UnregisterSlashCmd("undercut")
  - CDown: LibSlash.UnregisterSlashCmd("CDown")
