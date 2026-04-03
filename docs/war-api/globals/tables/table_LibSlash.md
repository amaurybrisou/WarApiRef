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
| Addons seen in | InfoScroller, Lib RuString, LibSlash, NPC Item Sale Price, PartyCast, Soloq, TidyChat, TidyRoll |
| Files seen in | `/workspace/data/raw/InfoScroller/InfoScroller.lua:30`, `/workspace/data/raw/PartyCast/PartyCast.lua:51`, `/workspace/data/raw/RuStringLib/RuStringLib.lua:282`, `/workspace/data/raw/Soloq/Soloq.lua:22`, `/workspace/data/raw/TidyChat/TidyChat.lua:189`, `/workspace/data/raw/TidyRoll/TidyRoll.lua:265`, `/workspace/data/raw/TimeToDie/TimeToDie.lua:229`, `/workspace/data/raw/minesweep/minesweep.lua:7` |
| Namespaces detected | LibSlash |
| Source kinds | globals, lua_calls |
| Example locations | InfoScroller: InfoScroller.OnInitialize, Lib RuString: LibRuString.OnLoad, NPC Item Sale Price: Nisp.Init, PartyCast: PartyCast.Init, Soloq: Soloq.OnInitialize, TidyChat: TidyChat.OnLoad |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 14 |
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

Observed shared global table or namespace surfaced in 10 addons.

## Functions

- LibSlash.RegisterSlashCmd
- LibSlash.RegisterWSlashCmd

## Observed Members

- none

## Seen In

- InfoScroller
- Lib RuString
- LibSlash
- NPC Item Sale Price
- PartyCast
- Soloq
- TidyChat
- TidyRoll
- TimeToDie
- minesweep

## Examples

- InfoScroller: InfoScroller.OnInitialize -> LibSlash.RegisterSlashCmd("infoscroller", function(input)InfoScroller_config.Slash(input)end)
- InfoScroller: InfoScroller.OnInitialize -> LibSlash.RegisterSlashCmd("info", function(input)InfoScroller_config.Slash(input)end)
- Lib RuString: LibRuString.OnLoad -> LibSlash.RegisterWSlashCmd("forcedrustrings", function(input)if(input==L "true")then LibRuString.ToggleHook(true)elseif(input==L "false")then LibRuString.ToggleHook(false)end end)
- NPC Item Sale Price: Nisp.Init -> LibSlash.RegisterSlashCmd("nisp", function(args)Nisp.SlashHandler(args)end)
- PartyCast: PartyCast.Init -> LibSlash.RegisterSlashCmd("pc", function(input)PartyCast.Command(input)end)
- PartyCast: PartyCast.Init -> LibSlash.RegisterSlashCmd("partycast", function(input)PartyCast.Command(input)end)

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
