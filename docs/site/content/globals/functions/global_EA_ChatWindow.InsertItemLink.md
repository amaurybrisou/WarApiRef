# EA_ChatWindow.InsertItemLink

- Category: Global Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 2 addons

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 123

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, matches a known engine namespace, called globally with no local definition.

## Evidence Signals

- +18 Seen in 2 to 3 addons: Cross-addon spread is present but limited.
- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +10 Argument pattern is consistent: Observed argument positions remain stable.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | AnywhereTrainerAdditions, nLootLink |
| Files seen in | AnywhereTrainerAdditions.lua, source/nLootLinkGUI.lua |
| Namespaces detected | EA_ChatWindow |
| Source kinds | lua_calls |
| Example locations | AnywhereTrainerAdditions: EquipmentLButtonDown, nLootLink: itemMouseDownL |
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
EA_ChatWindow.InsertItemLink(arg1)
```

## Description

Observed as a global function across 2 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a runtime window or control identifier. | Observed values: itemData |

## Returns

- Not confidently inferable from contract artifacts alone.

## Side Effects

- No side effect is confidently inferable from contract artifacts alone.

## Seen In

- AnywhereTrainerAdditions
- nLootLink

## Examples

- AnywhereTrainerAdditions: EquipmentLButtonDown -> EA_ChatWindow.InsertItemLink(itemData)
- nLootLink: itemMouseDownL -> EA_ChatWindow.InsertItemLink(itemData)

## Used With

- [Cursor.Clear](global_Cursor.Clear.md) (HIGH 100/100) - Global Function
- [Cursor.IconOnCursor](global_Cursor.IconOnCursor.md) (HIGH 100/100) - Global Function
- [Cursor.PickUp](global_Cursor.PickUp.md) (HIGH 100/100) - Global Function
- [BankWindow.GetItem](global_BankWindow.GetItem.md) (HIGH 71/100) - Global Function
- [BankWindow.GetSlotNumberForButtonIndex](global_BankWindow.GetSlotNumberForButtonIndex.md) (HIGH 71/100) - Global Function

## Notes

- Canonical entry built from observed call sites, not from engine source or decompiled definitions.
