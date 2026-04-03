# SendUseItem

- Category: Global Function
- Confidence level: HIGH
- Confidence score: 75/100
- Seen in: 4 addons

## Confidence Assessment

- Level: HIGH

- Score: 75/100

- Rationale: Promoted as HIGH confidence because seen in 4 or more addons, called globally with no local definition, role is consistent across addons.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +10 Argument pattern is consistent: Observed argument positions remain stable.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | BankWindowFix, KeyBar, PotionBar, RandomMount |
| Files seen in | KeyBar.lua, RandomMountCore.lua, Source/BankWindowFix.lua, source/Floating.lua |
| Namespaces detected | SendUseItem |
| Source kinds | lua_calls |
| Example locations | BankWindowFix: BagEquipmentRButtonUp, KeyBar: OnLButtonUp, PotionBar: LButtonUp, RandomMount: Mount |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 5 |
| Global usage count | 5 |
| Local definition count | 0 |
| Documentation references | 0 |
| Initialization flow references | 0 |
| Known engine namespace | no |
| Default UI presence | no |
| Event binding presence | no |
| Observed in XML and Lua | no |
| Consistent role | yes |
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
SendUseItem(arg1, arg2, arg3, arg4, arg5)
```

## Description

Observed as a shared action API across 4 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a function or method reference. | Observed values: GameData.ItemLocs.INVENTORY, ItemLoc |
| arg2 | Observed as a function or method reference. | Observed values: CurrentItem, PotionBarFloating.Buttons[PotionBar.getNumberFromWindowName()].slotNum, RandomMount.MountBagSlotNumbers[i] |
| arg3 | Observed as a numeric value. | Observed values: 0 |
| arg4 | Observed as a numeric value. | Observed values: 0 |
| arg5 | Observed as a numeric value. | Observed values: 0 |

## Returns

- Not confidently inferable from contract artifacts alone.

## Side Effects

- Sends a client action request back into the engine.

## Seen In

- BankWindowFix
- KeyBar
- PotionBar
- RandomMount

## Examples

- BankWindowFix: BagEquipmentRButtonUp -> SendUseItem(GameData.ItemLocs.INVENTORY, slot, 0, 0, 0)
- KeyBar: OnLButtonUp -> SendUseItem(ItemLoc, CurrentItem, 0, 0, 0)
- PotionBar: LButtonUp -> SendUseItem(GameData.ItemLocs.INVENTORY, PotionBarFloating.Buttons[PotionBar.getNumberFromWindowName()].slotNum, 0, 0, 0)
- RandomMount: Mount -> SendUseItem(GameData.ItemLocs.INVENTORY, RandomMount.SquigMountSlot, 0, 0, 0)
- RandomMount: Mount -> SendUseItem(GameData.ItemLocs.INVENTORY, RandomMount.MountBagSlotNumbers[i], 0, 0, 0)

## Related APIs

- [OnLButtonUp](../../xml/handlers/handler_OnLButtonUp.md) (HIGH 100/100) - XML Event

## Affects

- [GameData.ItemLocs.INVENTORY](../../gamedata/fields/gamedata_GameData.ItemLocs.INVENTORY.md) (HIGH 100/100) - GameData Field

## Notes

- Canonical entry built from observed call sites, not from engine source or decompiled definitions.
