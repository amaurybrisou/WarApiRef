# CraftingSystem.GetCraftingData

- Category: Global Function
- Confidence level: HIGH
- Confidence score: 83/100
- Seen in: 4 addons

## Confidence Assessment

- Level: HIGH

- Score: 83/100

- Rationale: Promoted as HIGH confidence because seen in 4 or more addons, called globally with no local definition, role is consistent across addons.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +10 Argument pattern is consistent: Observed argument positions remain stable.
- +8 Return usage is consistent: Observed as a stable query-style API.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | EZCraftX, JunkDump, Motion, nLootLink |
| Files seen in | JunkDump.lua, Source/EZCraftX.lua, Source/Motion.lua, source/nLootLinkData.lua |
| Namespaces detected | CraftingSystem |
| Source kinds | lua_calls |
| Example locations | EZCraftX: item_allowed_in_slot, JunkDump: professionChecks, JunkDump: professionChecks2, Motion: determineCraftSubKey, nLootLink: addListData |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 6 |
| Global usage count | 6 |
| Local definition count | 0 |
| Documentation references | 0 |
| Initialization flow references | 0 |
| Known engine namespace | no |
| Default UI presence | no |
| Event binding presence | no |
| Observed in XML and Lua | no |
| Consistent role | yes |
| Consistent arguments | yes |
| Consistent returns | yes |
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
CraftingSystem.GetCraftingData(arg1)
```

## Description

Observed as a global function across 4 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a runtime window or control identifier. | Observed values: itemData |

## Returns

- Observed as a query-style API. The concrete return shape is not inferable from contract artifacts alone.

## Side Effects

- No side effect is confidently inferable from contract artifacts alone.

## Seen In

- EZCraftX
- JunkDump
- Motion
- nLootLink

## Examples

- EZCraftX: item_allowed_in_slot -> CraftingSystem.GetCraftingData(itemData)
- JunkDump: professionChecks -> CraftingSystem.GetCraftingData(itemData)
- JunkDump: professionChecks2 -> CraftingSystem.GetCraftingData(itemData)
- Motion: determineCraftSubKey -> CraftingSystem.GetCraftingData(itemData)
- nLootLink: addListData -> CraftingSystem.GetCraftingData(itemData)

## Notes

- Canonical entry built from observed call sites, not from engine source or decompiled definitions.
