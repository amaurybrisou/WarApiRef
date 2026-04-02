# Slash Commands

## LibSlash registration

- Confidence: HIGH

- Description: Observed slash commands being registered through the shared LibSlash table.

- Evidence:

- AdvancedPetAssist: LibSlash.RegisterSlashCmd("apa", function(input)APA.SlashHandler(input)end)
  - Aura: LibSlash.RegisterSlashCmd("aura", AuraAddon.Slash)
  - Aura: LibSlash.RegisterSlashCmd("auraconfig", AuraAddon.Slash)
  - Aura: LibSlash.RegisterSlashCmd("showaura", AuraAddon.Slash)
  - AutoMark: LibSlash.RegisterSlashCmd("automark", AutoMark.OnSlashCommand)
  - BagOMatic: LibSlash.RegisterSlashCmd("bagomatic", function(msg)BagOMatic.parse_cmd(msg)end)

## LibSlash cleanup

- Confidence: HIGH

- Description: Observed slash commands being explicitly unregistered during addon shutdown paths.

- Evidence:

- DAoCBuff: LibSlash.UnregisterSlashCmd("daocbuff")
  - DAoCBuff: LibSlash.UnregisterSlashCmd("resetdaocbuff")
  - DaemonAssist: LibSlash.UnregisterSlashCmd("da")
  - Killer: LibSlash.UnregisterSlashCmd("killer")
  - PotionBar: LibSlash.UnregisterSlashCmd(PotionBar.LibSlashCommand)
  - QuickWarReport: LibSlash.UnregisterSlashCmd("qwr")
