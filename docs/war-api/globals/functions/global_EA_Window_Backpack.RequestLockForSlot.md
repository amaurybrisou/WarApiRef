# EA_Window_Backpack.RequestLockForSlot

- Category: Global Function
- Confidence level: HIGH
- Confidence score: 98/100
- Seen in: 2 addons

## Confidence Assessment

- Level: HIGH

- Score: 98/100

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, matches a known engine namespace, called globally with no local definition.

## Evidence Signals

- +18 Seen in 2 to 3 addons: Cross-addon spread is present but limited.
- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- -15 Conflicting signatures across usages: Observed arity or argument shape conflicts across usages.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | AuctionStats, zMailMod |
| Files seen in | AuctionAssist.lua, zMailModMassMail.lua |
| Namespaces detected | EA_Window_Backpack |
| Source kinds | lua_calls |
| Example locations | AuctionStats: PutUpForAuction, zMailMod: LockBackpackSlot |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 2 |
| Global usage count | 2 |
| Local definition count | 0 |
| Documentation references | 0 |
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
EA_Window_Backpack.RequestLockForSlot(arg1, arg2, arg3)
```

## Description

Observed as a global function across 2 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a function or method reference. | Observed values: inventorySlot.slot, slotB |
| arg2 | Observed as a runtime window or control identifier. | Observed values: backpack, inventorySlot.backpack |
| arg3 | Observed as a text or wstring payload. | Observed values: "AuctionHouseWindowCreateAuction", "EA_Window_BackpackIconViewSection"..bag.."Buttons" |

## Returns

- Not confidently inferable from contract artifacts alone.

## Side Effects

- No side effect is confidently inferable from contract artifacts alone.

## Seen In

- AuctionStats
- zMailMod

## Examples

- AuctionStats: PutUpForAuction -> EA_Window_Backpack.RequestLockForSlot(inventorySlot.slot, inventorySlot.backpack, "AuctionHouseWindowCreateAuction")
- zMailMod: LockBackpackSlot -> EA_Window_Backpack.RequestLockForSlot(slotB, backpack, "EA_Window_BackpackIconViewSection"..bag.."Buttons", {r=0,g=255,b=0})

## Notes

- Canonical entry built from observed call sites, not from engine source or decompiled definitions.
