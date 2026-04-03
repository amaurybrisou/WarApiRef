# EA_Window_Backpack.GetCursorForBackpack

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
| Addons seen in | BankWindowFix, Mass Refine, Miracle Grow Remix |
| Files seen in | MGRemix.lua, MassRefine.lua, Source/BankWindowFix.lua |
| Namespaces detected | EA_Window_Backpack |
| Source kinds | lua_calls |
| Example locations | BankWindowFix: BagEquipmentRButtonUp, Mass Refine: RefineItem, Miracle Grow Remix: GetCursorForBackpack |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 3 |
| Global usage count | 3 |
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
EA_Window_Backpack.GetCursorForBackpack(arg1)
```

## Description

Observed as a global function across 3 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a runtime window or control identifier. | Observed values: EA_Window_Backpack.currentMode, nBackpackType, nLoc |

## Returns

- Observed as a query-style API. The concrete return shape is not inferable from contract artifacts alone.

## Side Effects

- No side effect is confidently inferable from contract artifacts alone.

## Seen In

- BankWindowFix
- Mass Refine
- Miracle Grow Remix

## Examples

- BankWindowFix: BagEquipmentRButtonUp -> EA_Window_Backpack.GetCursorForBackpack(EA_Window_Backpack.currentMode)
- Mass Refine: RefineItem -> EA_Window_Backpack.GetCursorForBackpack(nLoc)
- Miracle Grow Remix: GetCursorForBackpack -> EA_Window_Backpack.GetCursorForBackpack(nBackpackType)

## Used With

- [DialogManager.MakeTwoButtonDialog](global_DialogManager.MakeTwoButtonDialog.md) (HIGH 100/100) - Global Function
- [GameData.ItemLocs.INVENTORY](../../gamedata/fields/gamedata_GameData.ItemLocs.INVENTORY.md) (HIGH 100/100) - GameData Field
- [EA_Window_Backpack.AutoAddCraftingItemIfPossible](global_EA_Window_Backpack.AutoAddCraftingItemIfPossible.md) (HIGH 90/100) - Global Function
- [EA_Window_Backpack.ConfirmThenRefine](global_EA_Window_Backpack.ConfirmThenRefine.md) (HIGH 80/100) - Global Function
- [BankWindow.IsShowing](global_BankWindow.IsShowing.md) (HIGH 71/100) - Global Function

## Affects

- [GameData.ItemLocs.INVENTORY](../../gamedata/fields/gamedata_GameData.ItemLocs.INVENTORY.md) (HIGH 100/100) - GameData Field

## Notes

- Canonical entry built from observed call sites, not from engine source or decompiled definitions.
- Advanced return analysis: No strong return evidence observed
