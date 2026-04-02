# EA_ChatWindow

- Category: Global Table
- Confidence level: HIGH
- Confidence score: 100/100

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 118

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, matches a known engine namespace, referenced by generated docs or reference files.

## Evidence Signals

- +18 Seen in 2 to 3 addons: Cross-addon spread is present but limited.
- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +25 Referenced by generated docs or reference files: The symbol is reinforced outside a single call page.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | NPC Item Sale Price, TidyChat |
| Files seen in | `/workspace/data/raw/TidyChat/TidyChat.lua:305`, `/workspace/data/raw/TidyChat/TidyChat.lua:489`, `/workspace/data/raw/nisp/Source/Nisp.lua:120`, `/workspace/data/raw/nisp/Source/Nisp.lua:184`, `/workspace/data/raw/nisp/Source/Nisp.lua:213`, `/workspace/data/raw/nisp/Source/Nisp.lua:26`, `/workspace/data/raw/nisp/Source/Nisp.lua:51` |
| Namespaces detected | EA_ChatWindow |
| Source kinds | globals, lua_calls |
| Example locations | NPC Item Sale Price: Nisp.DumpClear, NPC Item Sale Price: Nisp.DumpItem, NPC Item Sale Price: Nisp.Init, NPC Item Sale Price: Nisp.SetItemTooltipData, NPC Item Sale Price: Nisp.SlashHandler, TidyChat: TidyChatCore.AddNewChannels |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 22 |
| Global usage count | 4 |
| Local definition count | 1 |
| Documentation references | 1 |
| Initialization flow references | 0 |
| Known engine namespace | yes |
| Default UI presence | yes |
| Event binding presence | no |
| Observed in XML and Lua | no |
| Consistent role | yes |
| Consistent arguments | no |
| Consistent returns | no |
| Slash command presence | no |
| Weak usage only | no |
| Project-specific name | no |
| Placeholder or computed name | no |
| Conflicting signatures | no |
| Conflicting roles | no |
| Wrapper likely | no |
| Never outside local graph | no |
| Local helper only | no |

## Description

Observed shared global table or namespace surfaced in 2 addons.

## Functions

- EA_ChatWindow.OnSettingsChanged
- EA_ChatWindow.Print
- EA_ChatWindow.ShowTabCycleButtons
- EA_ChatWindow.UpdateTabScrollWindow

## Observed Members

- none

## Seen In

- NPC Item Sale Price
- TidyChat

## Examples

- NPC Item Sale Price: Nisp.DumpClear -> EA_ChatWindow.Print(L "Dump Table Cleared")
- NPC Item Sale Price: Nisp.DumpItem -> EA_ChatWindow.Print(L "Dumped item: "..itemData.name)
- NPC Item Sale Price: Nisp.DumpItem -> EA_ChatWindow.Print(L "Item Dump contains item: "..itemData.name)
- NPC Item Sale Price: Nisp.Init -> EA_ChatWindow.Print(L "Nisp Installed and Enabled")
- NPC Item Sale Price: Nisp.Init -> EA_ChatWindow.Print(L "Nisp Initialized and Enabled (/nisp for commands)")
- NPC Item Sale Price: Nisp.Init -> EA_ChatWindow.Print(L "Nisp Initialized, but disabled (/nisp for commands)")

## Related APIs

- [RegisterEventHandler](../functions/global_RegisterEventHandler.md) (HIGH 81/100) - Global Function

## Used With

- none

## Triggered By

- none

## Affects

- none

## Notes

- none
