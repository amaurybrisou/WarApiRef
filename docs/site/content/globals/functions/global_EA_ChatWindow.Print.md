# EA_ChatWindow.Print

- Category: Global Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 5 addons

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 135

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, seen in 4 or more addons, matches a known engine namespace.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +25 Referenced by generated docs or reference files: The symbol is reinforced outside a single call page.
- -15 Conflicting signatures across usages: Observed arity or argument shape conflicts across usages.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | Lib RuString, NPC Item Sale Price, Soloq, TimeToDie, ZCurse_Profiler |
| Files seen in | `/workspace/data/raw/CurseProfiler/CurseProfilerCompiled.lua:1952`, `/workspace/data/raw/CurseProfiler/CurseProfilerCompiled.lua:3343`, `/workspace/data/raw/CurseProfiler/CurseProfilerCompiled.lua:62`, `/workspace/data/raw/CurseProfiler/CurseProfilerCompiled.lua:68`, `/workspace/data/raw/RuStringLib/RuStringLib.lua:233`, `/workspace/data/raw/Soloq/Utils.lua:7`, `/workspace/data/raw/TimeToDie/TimeToDie.lua:3`, `/workspace/data/raw/nisp/Source/Nisp.lua:120` |
| Namespaces detected | EA_ChatWindow |
| Source kinds | globals, lua_calls |
| Example locations | Lib RuString: LibRuString.ToggleHook, NPC Item Sale Price: Nisp.DumpClear, NPC Item Sale Price: Nisp.DumpItem, NPC Item Sale Price: Nisp.Init, NPC Item Sale Price: Nisp.SetItemTooltipData, NPC Item Sale Price: Nisp.SlashHandler |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 29 |
| Global usage count | 29 |
| Local definition count | 0 |
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
| Conflicting signatures | yes |
| Conflicting roles | no |
| Wrapper likely | no |
| Never outside local graph | no |
| Local helper only | no |

## Signature (inferred)

```lua
EA_ChatWindow.Print(arg1)
```

## Description

Observed as a global function across 5 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a text or wstring payload. | Observed values: CurseProfiler.trueor(text), L "/nisp debug dumpclear - to clear dumped items", L "/nisp debug off - to turn debugging off" |

## Returns

- Not confidently inferable from addon-api docs alone.

## Side Effects

- Writes output to the chat or UI surface.

## Seen In

- Lib RuString
- NPC Item Sale Price
- Soloq
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

- [SystemData.Events.RELOAD_INTERFACE](../../systemdata/fields/systemdata_SystemData.Events.RELOAD_INTERFACE.md) (HIGH 100/100) - SystemData Field

## Notes

- Canonical entry built from observed call sites, not from engine source or decompiled definitions.
- Advanced return analysis: No strong return evidence observed
