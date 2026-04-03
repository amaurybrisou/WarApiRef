# Cursor.PickUp

- Category: Global Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 3 addons

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 113

- Rationale: Promoted as HIGH confidence because matches a known engine namespace, referenced by generated docs or reference files, called globally with no local definition.

## Evidence Signals

- +18 Seen in 2 to 3 addons: Cross-addon spread is present but limited.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +10 Argument pattern is consistent: Observed argument positions remain stable.
- +25 Referenced by generated docs or reference files: The symbol is reinforced outside a single call page.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | AnywhereTrainerAdditions, Enemy, HotbarMorale |
| Files seen in | `/workspace/data/raw/AnywhereTrainerAdditions/AnywhereTrainerAdditions.lua:218`, `/workspace/data/raw/Enemy/Code/Assist/Assist.lua:363`, `/workspace/data/raw/Enemy/Code/Core/ConfigurationWindow.lua:65`, `/workspace/data/raw/Enemy/Code/ScenarioInfo/ScenarioInfo.lua:1447`, `/workspace/data/raw/HotbarMorale/HotbarMorale.lua:10`, `/workspace/data/raw/HotbarMorale/HotbarMorale.lua:5` |
| Namespaces detected | Cursor |
| Source kinds | globals, lua_calls |
| Example locations | AnywhereTrainerAdditions: AnywhereTrainerAdditions.EquipmentLButtonDown, Enemy: Enemy.ConfigurationWindow_OnMacroMouseDrag, Enemy: Enemy.local.OnMacroMouseDrag, Enemy: OnMacroMouseDrag, HotbarMorale: ActionButton.CursorSwap, HotbarMorale: CursorSwap |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 7 |
| Global usage count | 7 |
| Local definition count | 0 |
| Documentation references | 1 |
| Initialization flow references | 0 |
| Known engine namespace | yes |
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
Cursor.PickUp(arg1, arg2, arg3, arg4, arg5)
```

## Description

Observed as a global function across 3 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a function or method reference. | Observed values: Cursor.SOURCE_BANK, Cursor.SOURCE_MACRO, slotType |
| arg2 | Observed as a runtime window or control identifier. | Observed values: 0, macro_id, p.macroId |
| arg3 | Observed as a runtime window or control identifier. | Observed values: itemData.uniqueID, macro_id, p.macroId |
| arg4 | Observed as a function or method reference. | Observed values: itemData.iconNum, macros[macro_id].iconNum, p.macroIconNum |
| arg5 | Observed as a boolean toggle. | Observed values: Cursor.AUTO_PICKUP_ON_LBUTTON_UP, false, true |

## Returns

- Not confidently inferable from addon-api docs alone.

## Side Effects

- No side effect is confidently inferable from addon-api docs alone.

## Seen In

- AnywhereTrainerAdditions
- Enemy
- HotbarMorale

## Examples

- AnywhereTrainerAdditions: AnywhereTrainerAdditions.EquipmentLButtonDown -> Cursor.PickUp(Cursor.SOURCE_BANK, slot, itemData.uniqueID, itemData.iconNum, true)
- Enemy: Enemy.ConfigurationWindow_OnMacroMouseDrag -> Cursor.PickUp(Cursor.SOURCE_MACRO, p.macroId, p.macroId, p.macroIconNum, false)
- Enemy: Enemy.local.OnMacroMouseDrag -> Cursor.PickUp(Cursor.SOURCE_MACRO, macro_id, macro_id, macros[macro_id].iconNum, false)
- Enemy: OnMacroMouseDrag -> Cursor.PickUp(Cursor.SOURCE_MACRO, macro_id, macro_id, macros[macro_id].iconNum, false)
- HotbarMorale: ActionButton.CursorSwap -> Cursor.PickUp(slotType, 0, slotId, slotIcon, Cursor.AUTO_PICKUP_ON_LBUTTON_UP)
- HotbarMorale: CursorSwap -> Cursor.PickUp(slotType, 0, slotId, slotIcon, Cursor.AUTO_PICKUP_ON_LBUTTON_UP)

## Related APIs

- none

## Used With

- [BankWindow](../tables/table_BankWindow.md) (HIGH 100/100) - Global Table
- [Cursor.Clear](global_Cursor.Clear.md) (HIGH 100/100) - Global Function
- [Cursor.IconOnCursor](global_Cursor.IconOnCursor.md) (HIGH 100/100) - Global Function
- [EA_ChatWindow.InsertItemLink](global_EA_ChatWindow.InsertItemLink.md) (HIGH 100/100) - Global Function
- [GameData.PlayerActions.DO_ABILITY](../../gamedata/fields/gamedata_GameData.PlayerActions.DO_ABILITY.md) (HIGH 100/100) - GameData Field
- [OnMouseOver](../../xml/handlers/handler_OnMouseOver.md) (HIGH 100/100) - XML Event

## Triggered By

- [OnMouseDrag](../../xml/handlers/handler_OnMouseDrag.md) (HIGH 100/100) - XML Event
- [OnMouseDrag](../../events/window_events/window_event_OnMouseDrag.md) (HIGH 100/100) - Window Event

## Affects

- [BankWindow](../tables/table_BankWindow.md) (HIGH 100/100) - Global Table
- [Button](../../xml/element_types/element_Button.md) (HIGH 100/100) - XML Element Type
- [GameData.PlayerActions.DO_ABILITY](../../gamedata/fields/gamedata_GameData.PlayerActions.DO_ABILITY.md) (HIGH 100/100) - GameData Field

## Notes

- Canonical entry built from observed call sites, not from engine source or decompiled definitions.
