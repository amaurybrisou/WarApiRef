# EA_BackpackUtilsMediator.GetItemsFromBackpack

- Category: Global Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 2 addons

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 156

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, matches a known engine namespace, referenced by generated docs or reference files.

## Evidence Signals

- +18 Seen in 2 to 3 addons: Cross-addon spread is present but limited.
- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +10 Argument pattern is consistent: Observed argument positions remain stable.
- +8 Return usage is consistent: Observed as a stable query-style API.
- +25 Referenced by generated docs or reference files: The symbol is reinforced outside a single call page.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | BagOMatic, TidyRoll |
| Files seen in | `/workspace/data/raw/TidyRoll/CustomAutoRoll.lua:400`, `/workspace/data/raw/bagomatic/BagOMatic.lua:818`, `/workspace/data/raw/bagomatic/BagOMatic.lua:888` |
| Namespaces detected | EA_BackpackUtilsMediator |
| Source kinds | globals, lua_calls |
| Example locations | BagOMatic: BagOMatic.SalvageHook, BagOMatic: BagOMatic.findItemInBagPack, TidyRoll: TidyRoll.CustomAutoRoll.OnListLbuttonUp |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 3 |
| Global usage count | 3 |
| Local definition count | 0 |
| Documentation references | 1 |
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
EA_BackpackUtilsMediator.GetItemsFromBackpack(arg1)
```

## Description

Observed as a global function across 2 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a runtime window or control identifier. | Observed values: backpackType |

## Returns

- Observed as a query-style API. The concrete return shape is not inferable from addon-api docs alone.

## Side Effects

- No side effect is confidently inferable from addon-api docs alone.

## Seen In

- BagOMatic
- TidyRoll

## Examples

- BagOMatic: BagOMatic.SalvageHook -> EA_BackpackUtilsMediator.GetItemsFromBackpack(backpackType)
- BagOMatic: BagOMatic.findItemInBagPack -> EA_BackpackUtilsMediator.GetItemsFromBackpack(backpackType)
- TidyRoll: TidyRoll.CustomAutoRoll.OnListLbuttonUp -> EA_BackpackUtilsMediator.GetItemsFromBackpack(backpackType)

## Related APIs

- [EA_BackpackUtilsMediator.GetBackpack](global_EA_BackpackUtilsMediator.GetBackpack.md) (HIGH 100/100) - Global Function
- [EA_BackpackUtilsMediator.GetCurrentBackpackType](global_EA_BackpackUtilsMediator.GetCurrentBackpackType.md) (HIGH 100/100) - Global Function
- [EA_BackpackUtilsMediator.GetCursorForBackpack](global_EA_BackpackUtilsMediator.GetCursorForBackpack.md) (HIGH 100/100) - Global Function

## Used With

- none

## Triggered By

- [OnLButtonUp](../../xml/handlers/handler_OnLButtonUp.md) (HIGH 100/100) - XML Event
- [OnLButtonUp](../../events/window_events/window_event_OnLButtonUp.md) (HIGH 100/100) - Window Event

## Affects

- [Cursor](../tables/table_Cursor.md) (HIGH 100/100) - Global Table
- [EA_BackpackUtilsMediator.GetBackpack](global_EA_BackpackUtilsMediator.GetBackpack.md) (HIGH 100/100) - Global Function
- [EA_BackpackUtilsMediator.GetCurrentBackpackType](global_EA_BackpackUtilsMediator.GetCurrentBackpackType.md) (HIGH 100/100) - Global Function
- [EA_BackpackUtilsMediator.GetCursorForBackpack](global_EA_BackpackUtilsMediator.GetCursorForBackpack.md) (HIGH 100/100) - Global Function
- [ListBox](../../xml/element_types/element_ListBox.md) (HIGH 100/100) - XML Element Type

## Notes

- Canonical entry built from observed call sites, not from engine source or decompiled definitions.
