# EA_ChatWindow

- Category: Global Table
- Confidence level: HIGH
- Confidence score: 100/100

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 130

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, seen in 4 or more addons, matches a known engine namespace.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +25 Referenced by generated docs or reference files: The symbol is reinforced outside a single call page.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | Lib RuString, LibSkillicon, NPC Item Sale Price, Soloq, TidyChat, TimeToDie, ZCurse_Profiler |
| Files seen in | `/workspace/data/raw/CurseProfiler/CurseProfilerCompiled.lua:1952`, `/workspace/data/raw/CurseProfiler/CurseProfilerCompiled.lua:3343`, `/workspace/data/raw/CurseProfiler/CurseProfilerCompiled.lua:62`, `/workspace/data/raw/CurseProfiler/CurseProfilerCompiled.lua:68`, `/workspace/data/raw/RuStringLib/RuStringLib.lua:233`, `/workspace/data/raw/Soloq/Utils.lua:7`, `/workspace/data/raw/TidyChat/TidyChat.lua:305`, `/workspace/data/raw/TidyChat/TidyChat.lua:489` |
| Namespaces detected | EA_ChatWindow |
| Source kinds | globals, lua_calls |
| Example locations | Lib RuString: LibRuString.ToggleHook, NPC Item Sale Price: Nisp.DumpClear, NPC Item Sale Price: Nisp.DumpItem, NPC Item Sale Price: Nisp.Init, NPC Item Sale Price: Nisp.SetItemTooltipData, NPC Item Sale Price: Nisp.SlashHandler |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 32 |
| Global usage count | 4 |
| Local definition count | 3 |
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

Observed shared global table or namespace surfaced in 7 addons.

## Functions

- EA_ChatWindow.OnSettingsChanged
- EA_ChatWindow.Print
- EA_ChatWindow.ShowTabCycleButtons
- EA_ChatWindow.UpdateTabScrollWindow

## Observed Members

- none

## Seen In

- Lib RuString
- LibSkillicon
- NPC Item Sale Price
- Soloq
- TidyChat
- TimeToDie
- ZCurse_Profiler

## Examples

- Lib RuString: LibRuString.ToggleHook -> EA_ChatWindow.Print(WRu("RuStringsLib: ýýýýý ýýýýýýýýý ýýýýýýýý ý ýýýý, ýýýýýýýýýýýýý ýýýýýýýýý. ýýý ýýýýý ýýýýýýý ýýýýýýýý '/reloadui'."))
- NPC Item Sale Price: Nisp.DumpClear -> EA_ChatWindow.Print(L "Dump Table Cleared")
- NPC Item Sale Price: Nisp.DumpItem -> EA_ChatWindow.Print(L "Dumped item: "..itemData.name)
- NPC Item Sale Price: Nisp.DumpItem -> EA_ChatWindow.Print(L "Item Dump contains item: "..itemData.name)
- NPC Item Sale Price: Nisp.Init -> EA_ChatWindow.Print(L "Nisp Installed and Enabled")
- NPC Item Sale Price: Nisp.Init -> EA_ChatWindow.Print(L "Nisp Initialized and Enabled (/nisp for commands)")

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
