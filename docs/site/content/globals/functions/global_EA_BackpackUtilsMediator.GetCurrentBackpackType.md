# EA_BackpackUtilsMediator.GetCurrentBackpackType

- Category: Global Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 3 addons

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
| Addons seen in | BagOMatic, TidyRoll, zMailMod |
| Files seen in | BagOMatic.lua, CustomAutoRoll.lua, zMailMod.lua, zMailModMassMail.lua |
| Namespaces detected | EA_BackpackUtilsMediator |
| Source kinds | lua_calls |
| Example locations | BagOMatic: findItemInBagPack, TidyRoll: OnListLbuttonUp, zMailMod: InventoryLButtonDown, zMailMod: InventoryLButtonUp, zMailMod: InventoryRButtonUp, zMailMod: OnSlotLButtonUp |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 6 |
| Global usage count | 6 |
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
EA_BackpackUtilsMediator.GetCurrentBackpackType()
```

## Description

Observed as a global function across 3 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |

## Returns

- Observed as a query-style API. The concrete return shape is not inferable from contract artifacts alone.

## Side Effects

- No side effect is confidently inferable from contract artifacts alone.

## Seen In

- BagOMatic
- TidyRoll
- zMailMod

## Examples

- BagOMatic: findItemInBagPack -> EA_BackpackUtilsMediator.GetCurrentBackpackType()
- TidyRoll: OnListLbuttonUp -> EA_BackpackUtilsMediator.GetCurrentBackpackType()
- zMailMod: InventoryLButtonDown -> EA_BackpackUtilsMediator.GetCurrentBackpackType()
- zMailMod: InventoryLButtonUp -> EA_BackpackUtilsMediator.GetCurrentBackpackType()
- zMailMod: InventoryRButtonUp -> EA_BackpackUtilsMediator.GetCurrentBackpackType()
- zMailMod: OnSlotLButtonUp -> EA_BackpackUtilsMediator.GetCurrentBackpackType()

## Used With

- [Cursor.Clear](global_Cursor.Clear.md) (HIGH 100/100) - Global Function
- [Cursor.IconOnCursor](global_Cursor.IconOnCursor.md) (HIGH 100/100) - Global Function
- [EA_BackpackUtilsMediator.GetItemsFromBackpack](global_EA_BackpackUtilsMediator.GetItemsFromBackpack.md) (HIGH 100/100) - Global Function
- [EA_Window_Backpack.GetSlotFromActionButtonGroup](global_EA_Window_Backpack.GetSlotFromActionButtonGroup.md) (HIGH 100/100) - Global Function
- [SystemData.ActiveWindow.name](../../systemdata/fields/systemdata_SystemData.ActiveWindow.name.md) (HIGH 100/100) - SystemData Field

## Affects

- [SystemData.ActiveWindow.name](../../systemdata/fields/systemdata_SystemData.ActiveWindow.name.md) (HIGH 100/100) - SystemData Field

## Notes

- Canonical entry built from observed call sites, not from engine source or decompiled definitions.
