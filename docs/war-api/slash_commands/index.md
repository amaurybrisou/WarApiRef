# Slash Commands

## LibSlash registration

- Confidence: HIGH

- Description: Observed slash commands being registered through the shared LibSlash table.

- Evidence:

- AdvancedPetAssist: LibSlash.RegisterSlashCmd("apa", function(input)APA.SlashHandler(input)end)
  - Aura: LibSlash.RegisterSlashCmd("aura", AuraAddon.Slash)
  - Aura: LibSlash.RegisterSlashCmd("auraconfig", AuraAddon.Slash)
  - Aura: LibSlash.RegisterSlashCmd("showaura", AuraAddon.Slash)
  - AutoBand: LibSlash.RegisterSlashCmd("autoband", function(msg)AutoBand.parse_cmd(msg)end)
  - AutoBand: LibSlash.RegisterSlashCmd("ab", function(msg)AutoBand.parse_cmd(msg)end)

## LibSlash cleanup

- Confidence: HIGH

- Description: Observed slash commands being explicitly unregistered during addon shutdown paths.

- Evidence:

- DAoCBuff: LibSlash.UnregisterSlashCmd("daocbuff")
  - DAoCBuff: LibSlash.UnregisterSlashCmd("resetdaocbuff")
  - DaemonAssist: LibSlash.UnregisterSlashCmd("da")
  - DeepSleep: LibSlash.UnregisterSlashCmd("ds")
  - DeepSleep: LibSlash.UnregisterSlashCmd("DeepSleep")
  - Killer: LibSlash.UnregisterSlashCmd("killer")
