# LibSlash registration

- Category: slash_commands
- Confidence: HIGH

## Description

Observed slash commands being registered through the shared LibSlash table.

## Involved APIs

- [LibSlash](../globals/tables/table_LibSlash.md) (HIGH 100/100) - Global Table
- [Window](../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type
- [LibSlash.RegisterSlashCmd](../globals/functions/global_LibSlash.RegisterSlashCmd.md) (HIGH 85/100) - Global Function

## Flow Diagram

```text
OnInitialize -> LibSlash.RegisterSlashCmd
```

## Example Code

```lua
ActionBarCD: LibSlash.RegisterSlashCmd("abcd", function(args)ActionBarCD.SlashHandler(args)end)
```

## Evidence

- ActionBarCD: LibSlash.RegisterSlashCmd("abcd", function(args)ActionBarCD.SlashHandler(args)end)
- ActionBarHide: LibSlash.RegisterSlashCmd("abh", function()ActionBarHide.OptionsWindow()end)
- AdvancedPetAssist: LibSlash.RegisterSlashCmd("apa", function(input)APA.SlashHandler(input)end)
- Amethyst: LibSlash.RegisterSlashCmd("amt", function(msg)Amethyst.Slash(msg)end)
- Amethyst: LibSlash.RegisterSlashCmd("amethyst", function(msg)Amethyst.Slash(msg)end)
- ArmorGraphicToggle: LibSlash.RegisterSlashCmd("agt", function(msg)ArmorGraphicToggle.Slash(msg)end)
