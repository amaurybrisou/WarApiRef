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
OnLButtonUp
  -> ui: Button, ColorPicker, DynamicImage, Label, ListBox, MapDisplay, Window
```

## Example Code

```lua
DAoCBuff: LibSlash.UnregisterSlashCmd("daocbuff")
```

## Evidence

- DAoCBuff: LibSlash.UnregisterSlashCmd("daocbuff")
- DAoCBuff: LibSlash.UnregisterSlashCmd("resetdaocbuff")
- Killer: LibSlash.UnregisterSlashCmd("killer")
- PotionBar: LibSlash.UnregisterSlashCmd(PotionBar.LibSlashCommand)
- Shinies: LibSlash.UnregisterSlashCmd("shinies")
- WhoHealedMe: LibSlash.UnregisterSlashCmd("whm")
