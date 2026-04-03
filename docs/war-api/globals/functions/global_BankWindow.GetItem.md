# BankWindow.GetItem

- Category: Global Function
- Confidence level: HIGH
- Confidence score: 71/100
- Seen in: 2 addons

## Confidence Assessment

- Level: HIGH

- Score: 71/100

- Rationale: Promoted as HIGH confidence because called globally with no local definition, seen in 2 to 3 addons, role is consistent across addons.

## Evidence Signals

- +18 Seen in 2 to 3 addons: Cross-addon spread is present but limited.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +10 Argument pattern is consistent: Observed argument positions remain stable.
- +8 Return usage is consistent: Observed as a stable query-style API.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | AnywhereTrainerAdditions, BankWindowFix |
| Files seen in | AnywhereTrainerAdditions.lua, Source/BankWindowFix.lua |
| Namespaces detected | BankWindow |
| Source kinds | lua_calls |
| Example locations | AnywhereTrainerAdditions: EquipmentLButtonDown, BankWindowFix: BankEquipmentRButtonDown |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 2 |
| Global usage count | 2 |
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
BankWindow.GetItem(arg1)
```

## Description

Observed as a global function across 2 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a runtime window or control identifier. | Observed values: slot |

## Returns

- Observed as a query-style API. The concrete return shape is not inferable from contract artifacts alone.

## Side Effects

- No side effect is confidently inferable from contract artifacts alone.

## Seen In

- AnywhereTrainerAdditions
- BankWindowFix

## Examples

- AnywhereTrainerAdditions: EquipmentLButtonDown -> BankWindow.GetItem(slot)
- BankWindowFix: BankEquipmentRButtonDown -> BankWindow.GetItem(slot)

## Used With

- [Cursor.Clear](global_Cursor.Clear.md) (HIGH 100/100) - Global Function
- [Cursor.IconOnCursor](global_Cursor.IconOnCursor.md) (HIGH 100/100) - Global Function
- [Cursor.PickUp](global_Cursor.PickUp.md) (HIGH 100/100) - Global Function
- [EA_ChatWindow.InsertItemLink](global_EA_ChatWindow.InsertItemLink.md) (HIGH 100/100) - Global Function
- [BankWindow.GetSlotNumberForButtonIndex](global_BankWindow.GetSlotNumberForButtonIndex.md) (HIGH 71/100) - Global Function

## Notes

- Canonical entry built from observed call sites, not from engine source or decompiled definitions.
