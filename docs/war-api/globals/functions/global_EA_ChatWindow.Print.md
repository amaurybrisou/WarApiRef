# EA_ChatWindow.Print

- Category: Global Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 1 addons

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 115

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, matches a known engine namespace, referenced by generated docs or reference files.

## Evidence Signals

- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +10 Argument pattern is consistent: Observed argument positions remain stable.
- +25 Referenced by generated docs or reference files: The symbol is reinforced outside a single call page.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | NPC Item Sale Price |
| Files seen in | `/workspace/data/raw/nisp/Source/Nisp.lua:120`, `/workspace/data/raw/nisp/Source/Nisp.lua:184`, `/workspace/data/raw/nisp/Source/Nisp.lua:213`, `/workspace/data/raw/nisp/Source/Nisp.lua:26`, `/workspace/data/raw/nisp/Source/Nisp.lua:51` |
| Namespaces detected | EA_ChatWindow |
| Source kinds | globals, lua_calls |
| Example locations | NPC Item Sale Price: Nisp.DumpClear, NPC Item Sale Price: Nisp.DumpItem, NPC Item Sale Price: Nisp.Init, NPC Item Sale Price: Nisp.SetItemTooltipData, NPC Item Sale Price: Nisp.SlashHandler |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 19 |
| Global usage count | 19 |
| Local definition count | 0 |
| Documentation references | 1 |
| Initialization flow references | 0 |
| Known engine namespace | yes |
| Default UI presence | yes |
| Event binding presence | no |
| Observed in XML and Lua | no |
| Consistent role | no |
| Consistent arguments | yes |
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

## Signature (inferred)

```lua
EA_ChatWindow.Print(arg1)
```

## Description

Observed as a global function across 1 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a text or wstring payload. | Observed values: L "/nisp debug dumpclear - to clear dumped items", L "/nisp debug off - to turn debugging off", L "/nisp debug on - to turn debugging on" |

## Returns

- Not confidently inferable from addon-api docs alone.

## Side Effects

- Writes output to the chat or UI surface.

## Seen In

- NPC Item Sale Price

## Examples

- NPC Item Sale Price: Nisp.DumpClear -> EA_ChatWindow.Print(L "Dump Table Cleared")
- NPC Item Sale Price: Nisp.DumpItem -> EA_ChatWindow.Print(L "Dumped item: "..itemData.name)
- NPC Item Sale Price: Nisp.DumpItem -> EA_ChatWindow.Print(L "Item Dump contains item: "..itemData.name)
- NPC Item Sale Price: Nisp.Init -> EA_ChatWindow.Print(L "Nisp Installed and Enabled")
- NPC Item Sale Price: Nisp.Init -> EA_ChatWindow.Print(L "Nisp Initialized and Enabled (/nisp for commands)")
- NPC Item Sale Price: Nisp.Init -> EA_ChatWindow.Print(L "Nisp Initialized, but disabled (/nisp for commands)")

## Related APIs

- none

## Used With

- [LibSlash.RegisterSlashCmd](global_LibSlash.RegisterSlashCmd.md) (HIGH 80/100) - Global Function

## Triggered By

- none

## Affects

- none

## Notes

- Only one addon surfaced this symbol in the current corpus.
- Canonical entry built from observed call sites, not from engine source or decompiled definitions.
- Advanced return analysis: No strong return evidence observed
