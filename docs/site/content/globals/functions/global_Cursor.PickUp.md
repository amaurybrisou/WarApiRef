# Cursor.PickUp

- Category: Global Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 9 addons

## Confidence Assessment

- Level: HIGH

- Score: 100/100

- Rationale: Promoted as HIGH confidence because seen in 4 or more addons, matches a known engine namespace, called globally with no local definition.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +10 Argument pattern is consistent: Observed argument positions remain stable.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | AnywhereTrainerAdditions, CastSequence, Enemy, FozAuction, HotbarMorale, NerfedButtons, Sequencer, bigger_MacroWindow |
| Files seen in | AnywhereTrainerAdditions.lua, Code/Assist/Assist.lua, Code/Core/ConfigurationWindow.lua, Code/ScenarioInfo/ScenarioInfo.lua, HotbarMorale.lua, Sequencer.lua, Setup/Button.lua, Source/auctionwindowsellcontrols.lua |
| Namespaces detected | Cursor |
| Source kinds | lua_calls |
| Example locations | AnywhereTrainerAdditions: EquipmentLButtonDown, CastSequence: BeginDrag, Enemy: ConfigurationWindow_OnMacroMouseDrag, Enemy: OnMacroMouseDrag, FozAuction: PickupItemIfPossible, HotbarMorale: CursorSwap |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 18 |
| Global usage count | 18 |
| Local definition count | 0 |
| Documentation references | 0 |
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

Observed as a global function across 9 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a function or method reference. | Observed values: Cursor.SOURCE_ACTION_LIST, Cursor.SOURCE_BANK, Cursor.SOURCE_INVENTORY |
| arg2 | Observed as a runtime window or control identifier. | Observed values: 0, AuctionWindowSellControls.itemInventorySlot.slot, Cursor.SOURCE_ACTION_LIST |
| arg3 | Observed as a runtime window or control identifier. | Observed values: abData.id, abilityData.id, abilityId |
| arg4 | Observed as a function or method reference. | Observed values: abData.iconNum, abilityData.iconNum, buttonData.Icon |
| arg5 | Observed as a boolean toggle. | Observed values: Cursor.AUTO_PICKUP_ON_LBUTTON_UP, false, true |

## Returns

- Not confidently inferable from contract artifacts alone.

## Side Effects

- No side effect is confidently inferable from contract artifacts alone.

## Seen In

- AnywhereTrainerAdditions
- CastSequence
- Enemy
- FozAuction
- HotbarMorale
- NerfedButtons
- Sequencer
- bigger_MacroWindow
- zMailMod

## Examples

- AnywhereTrainerAdditions: EquipmentLButtonDown -> Cursor.PickUp(Cursor.SOURCE_BANK, slot, itemData.uniqueID, itemData.iconNum, true)
- CastSequence: BeginDrag -> Cursor.PickUp(action, 0, actionId, buttonData.Icon, Cursor.AUTO_PICKUP_ON_LBUTTON_UP)
- Enemy: ConfigurationWindow_OnMacroMouseDrag -> Cursor.PickUp(Cursor.SOURCE_MACRO, p.macroId, p.macroId, p.macroIconNum, false)
- Enemy: OnMacroMouseDrag -> Cursor.PickUp(Cursor.SOURCE_MACRO, macro_id, macro_id, macros[macro_id].iconNum, false)
- FozAuction: PickupItemIfPossible -> Cursor.PickUp(cursor, AuctionWindowSellControls.itemInventorySlot.slot, itemData.uniqueID, itemData.iconNum, true)
- HotbarMorale: CursorSwap -> Cursor.PickUp(slotType, 0, slotId, slotIcon, Cursor.AUTO_PICKUP_ON_LBUTTON_UP)

## Related APIs

- [IconLButtonUp](../../events/game_events/game_event_IconLButtonUp.md) (MEDIUM 43/100) - Game Event

## Used With

- [Cursor.Clear](global_Cursor.Clear.md) (HIGH 100/100) - Global Function
- [Cursor.IconOnCursor](global_Cursor.IconOnCursor.md) (HIGH 100/100) - Global Function
- [EA_BackpackUtilsMediator.GetCurrentBackpackType](global_EA_BackpackUtilsMediator.GetCurrentBackpackType.md) (HIGH 100/100) - Global Function
- [EA_BackpackUtilsMediator.GetCursorForBackpack](global_EA_BackpackUtilsMediator.GetCursorForBackpack.md) (HIGH 100/100) - Global Function
- [EA_BackpackUtilsMediator.GetItemsFromBackpack](global_EA_BackpackUtilsMediator.GetItemsFromBackpack.md) (HIGH 100/100) - Global Function
- [EA_ChatWindow.InsertItemLink](global_EA_ChatWindow.InsertItemLink.md) (HIGH 100/100) - Global Function
- [EA_Window_Macro.UpdateDetails](global_EA_Window_Macro.UpdateDetails.md) (HIGH 100/100) - Global Function
- [EA_BackpackUtilsMediator.ReleaseLockForSlot](global_EA_BackpackUtilsMediator.ReleaseLockForSlot.md) (HIGH 90/100) - Global Function
- [BankWindow.GetItem](global_BankWindow.GetItem.md) (HIGH 71/100) - Global Function
- [BankWindow.GetSlotNumberForButtonIndex](global_BankWindow.GetSlotNumberForButtonIndex.md) (HIGH 71/100) - Global Function

## Affects

- [GameData.PlayerActions.DO_ABILITY](../../gamedata/fields/gamedata_GameData.PlayerActions.DO_ABILITY.md) (HIGH 100/100) - GameData Field

## Notes

- Canonical entry built from observed call sites, not from engine source or decompiled definitions.
