# LibSlash

- Category: Global Table
- Confidence level: HIGH
- Confidence score: 100/100

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 105

- Rationale: Promoted as HIGH confidence because seen in 4 or more addons, matches a known engine namespace, referenced by generated docs or reference files.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +10 Appears in slash command registration patterns: Observed in shared command registration flows.
- +25 Referenced by generated docs or reference files: The symbol is reinforced outside a single call page.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | LibSlash, NPC Item Sale Price, TidyChat, TidyRoll |
| Files seen in | `/workspace/data/raw/TidyChat/TidyChat.lua:189`, `/workspace/data/raw/TidyRoll/TidyRoll.lua:265`, `/workspace/data/raw/nisp/Source/Nisp.lua:26` |
| Namespaces detected | LibSlash |
| Source kinds | globals, lua_calls |
| Example locations | NPC Item Sale Price: Nisp.Init, TidyChat: TidyChat.OnLoad, TidyRoll: TidyRoll.OnLoad |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 3 |
| Global usage count | 2 |
| Local definition count | 1 |
| Documentation references | 1 |
| Initialization flow references | 0 |
| Known engine namespace | yes |
| Default UI presence | no |
| Event binding presence | no |
| Observed in XML and Lua | no |
| Consistent role | yes |
| Consistent arguments | no |
| Consistent returns | no |
| Slash command presence | yes |
| Weak usage only | no |
| Project-specific name | no |
| Placeholder or computed name | no |
| Conflicting signatures | no |
| Conflicting roles | no |
| Wrapper likely | no |
| Never outside local graph | no |
| Local helper only | no |

## Description

Observed shared global table or namespace surfaced in 4 addons.

## Functions

- LibSlash.RegisterSlashCmd
- LibSlash.RegisterWSlashCmd

## Observed Members

- none

## Seen In

- LibSlash
- NPC Item Sale Price
- TidyChat
- TidyRoll

## Examples

- NPC Item Sale Price: Nisp.Init -> LibSlash.RegisterSlashCmd("nisp", function(args)Nisp.SlashHandler(args)end)
- TidyChat: TidyChat.OnLoad -> LibSlash.RegisterWSlashCmd("tchat", TidyChat.ToggleOptions)
- TidyRoll: TidyRoll.OnLoad -> LibSlash.RegisterWSlashCmd("troll", TidyRoll.ToggleOptions)

## Related APIs

- none

## Used With

- none

## Triggered By

- none

## Affects

- none

## Notes

- none
