# EA_Window_Backpack.GetSlotFromActionButtonGroup

- Category: Global Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 2 addons

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 131

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, matches a known engine namespace, called globally with no local definition.

## Evidence Signals

- +18 Seen in 2 to 3 addons: Cross-addon spread is present but limited.
- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +10 Argument pattern is consistent: Observed argument positions remain stable.
- +8 Return usage is consistent: Observed as a stable query-style API.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | NPC Item Sale Price, zMailMod |
| Files seen in | Source/Nisp.lua, zMailMod.lua |
| Namespaces detected | EA_Window_Backpack |
| Source kinds | lua_calls |
| Example locations | NPC Item Sale Price: RButtonUp, zMailMod: InventoryLButtonDown, zMailMod: InventoryLButtonUp, zMailMod: InventoryRButtonUp |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 4 |
| Global usage count | 4 |
| Local definition count | 0 |
| Documentation references | 0 |
| Initialization flow references | 0 |
| Known engine namespace | yes |
| Default UI presence | yes |
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
EA_Window_Backpack.GetSlotFromActionButtonGroup(arg1, arg2)
```

## Description

Observed as a global function across 2 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a function or method reference. | Observed values: SystemData.ActiveWindow.name |
| arg2 | Observed as a runtime window or control identifier. | Observed values: buttonId |

## Returns

- Observed as a query-style API. The concrete return shape is not inferable from contract artifacts alone.

## Side Effects

- No side effect is confidently inferable from contract artifacts alone.

## Seen In

- NPC Item Sale Price
- zMailMod

## Examples

- NPC Item Sale Price: RButtonUp -> EA_Window_Backpack.GetSlotFromActionButtonGroup(SystemData.ActiveWindow.name, buttonId)
- zMailMod: InventoryLButtonDown -> EA_Window_Backpack.GetSlotFromActionButtonGroup(SystemData.ActiveWindow.name, buttonId)
- zMailMod: InventoryLButtonUp -> EA_Window_Backpack.GetSlotFromActionButtonGroup(SystemData.ActiveWindow.name, buttonId)
- zMailMod: InventoryRButtonUp -> EA_Window_Backpack.GetSlotFromActionButtonGroup(SystemData.ActiveWindow.name, buttonId)

## Used With

- [EA_BackpackUtilsMediator.GetCurrentBackpackType](global_EA_BackpackUtilsMediator.GetCurrentBackpackType.md) (HIGH 100/100) - Global Function
- [SystemData.ActiveWindow.name](../../systemdata/fields/systemdata_SystemData.ActiveWindow.name.md) (HIGH 100/100) - SystemData Field

## Affects

- [SystemData.ActiveWindow.name](../../systemdata/fields/systemdata_SystemData.ActiveWindow.name.md) (HIGH 100/100) - SystemData Field

## Notes

- Canonical entry built from observed call sites, not from engine source or decompiled definitions.
