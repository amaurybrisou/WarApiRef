# Cursor.Clear

- Category: Global Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 6 addons

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 125

- Rationale: Promoted as HIGH confidence because seen in 4 or more addons, matches a known engine namespace, referenced by generated docs or reference files.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +10 Argument pattern is consistent: Observed argument positions remain stable.
- +25 Referenced by generated docs or reference files: The symbol is reinforced outside a single call page.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | AnywhereTrainerAdditions, CM_ClosetGoblin, Enemy, HotbarMorale, RandomMount, Shinies |
| Files seen in | `/workspace_addons/AnywhereTrainerAdditions/AnywhereTrainerAdditions.lua:218`, `/workspace_addons/ClosetGoblin/ClosetGoblinCharacterWindow.lua:512`, `/workspace_addons/Enemy/Code/UnitFrames/ClickCasting.lua:379`, `/workspace_addons/HotbarMorale/HotbarMorale.lua:10`, `/workspace_addons/HotbarMorale/HotbarMorale.lua:5`, `/workspace_addons/RandomMount/RandomMountUI.lua:415`, `/workspace_addons/Shinies/Modules/UI/Shinies-UI-Browse.lua:1105`, `/workspace_addons/Shinies/Modules/UI/Shinies-UI-Browse.lua:1141` |
| Namespaces detected | Cursor |
| Source kinds | globals, lua_calls |
| Example locations | AnywhereTrainerAdditions: AnywhereTrainerAdditions.EquipmentLButtonDown, CM_ClosetGoblin: ClosetGoblinCharacterWindow.HandleDrag, Enemy: Enemy.UnitFramesUI_UnitFrameClickCastingDialog_OnActionConfig1ActionLButtonUp, HotbarMorale: ActionButton.CursorSwap, HotbarMorale: CursorSwap, HotbarMorale: HotbarMorale.Initialize |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 12 |
| Global usage count | 12 |
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
Cursor.Clear()
```

## Description

Observed as a global function across 6 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |

## Returns

- Not confidently inferable from addon-api docs alone.

## Side Effects

- No side effect is confidently inferable from addon-api docs alone.

## Seen In

- AnywhereTrainerAdditions
- CM_ClosetGoblin
- Enemy
- HotbarMorale
- RandomMount
- Shinies

## Examples

- AnywhereTrainerAdditions: AnywhereTrainerAdditions.EquipmentLButtonDown -> Cursor.Clear()
- CM_ClosetGoblin: ClosetGoblinCharacterWindow.HandleDrag -> Cursor.Clear()
- Enemy: Enemy.UnitFramesUI_UnitFrameClickCastingDialog_OnActionConfig1ActionLButtonUp -> Cursor.Clear()
- HotbarMorale: ActionButton.CursorSwap -> Cursor.Clear()
- HotbarMorale: CursorSwap -> Cursor.Clear()
- HotbarMorale: HotbarMorale.Initialize -> Cursor.Clear()

## Related APIs

- [OnLButtonDown](../../xml/handlers/handler_OnLButtonDown.md) (HIGH 100/100) - XML Handler

## Used With

- [BankWindow](../tables/table_BankWindow.md) (HIGH 100/100) - Global Table
- [Cursor](../tables/table_Cursor.md) (HIGH 100/100) - Global Table
- [Cursor.IconOnCursor](global_Cursor.IconOnCursor.md) (HIGH 100/100) - Global Function
- [Cursor.PickUp](global_Cursor.PickUp.md) (HIGH 100/100) - Global Function
- [DynamicImage](../../xml/element_types/element_DynamicImage.md) (HIGH 100/100) - XML Element Type
- [EA_BackpackUtilsMediator.GetCurrentBackpackType](global_EA_BackpackUtilsMediator.GetCurrentBackpackType.md) (HIGH 100/100) - Global Function
- [EA_BackpackUtilsMediator.GetCursorForBackpack](global_EA_BackpackUtilsMediator.GetCursorForBackpack.md) (HIGH 100/100) - Global Function
- [EA_BackpackUtilsMediator.GetItemsFromBackpack](global_EA_BackpackUtilsMediator.GetItemsFromBackpack.md) (HIGH 100/100) - Global Function
- [EA_ChatWindow.InsertItemLink](global_EA_ChatWindow.InsertItemLink.md) (HIGH 100/100) - Global Function
- [GameData.PlayerActions.DO_ABILITY](../../gamedata/fields/gamedata_GameData.PlayerActions.DO_ABILITY.md) (HIGH 100/100) - GameData Field
- [ListBox](../../xml/element_types/element_ListBox.md) (HIGH 100/100) - XML Element Type
- [OnLButtonUp](../../events/window_events/window_event_OnLButtonUp.md) (HIGH 100/100) - Window Event
- [OnLButtonUp](../../xml/handlers/handler_OnLButtonUp.md) (HIGH 100/100) - XML Handler
- [OnMouseOver](../../xml/handlers/handler_OnMouseOver.md) (HIGH 100/100) - XML Handler
- [SystemData.ActiveWindow.name](../../systemdata/fields/systemdata_SystemData.ActiveWindow.name.md) (HIGH 100/100) - SystemData Field
- [WindowGetId](../../window_api/functions/window_WindowGetId.md) (HIGH 100/100) - Window Function
- [WindowSetShowing](../../window_api/functions/window_WindowSetShowing.md) (HIGH 100/100) - Window Function
- [BankWindow.GetItem](global_BankWindow.GetItem.md) (HIGH 96/100) - Global Function
- [BankWindow.GetSlotNumberForButtonIndex](global_BankWindow.GetSlotNumberForButtonIndex.md) (HIGH 96/100) - Global Function

## Triggered By

- [OnLButtonUp](../../xml/handlers/handler_OnLButtonUp.md) (HIGH 100/100) - XML Handler
- [OnLButtonUp](../../events/window_events/window_event_OnLButtonUp.md) (HIGH 100/100) - Window Event

## Affects

- [BankWindow](../tables/table_BankWindow.md) (HIGH 100/100) - Global Table
- [DynamicImage](../../xml/element_types/element_DynamicImage.md) (HIGH 100/100) - XML Element Type
- [GameData.PlayerActions.DO_ABILITY](../../gamedata/fields/gamedata_GameData.PlayerActions.DO_ABILITY.md) (HIGH 100/100) - GameData Field
- [SystemData.ActiveWindow.name](../../systemdata/fields/systemdata_SystemData.ActiveWindow.name.md) (HIGH 100/100) - SystemData Field
- [BankWindow.GetItem](global_BankWindow.GetItem.md) (HIGH 96/100) - Global Function
- [BankWindow.GetSlotNumberForButtonIndex](global_BankWindow.GetSlotNumberForButtonIndex.md) (HIGH 96/100) - Global Function

## Notes

- Canonical entry built from observed call sites, not from engine source or decompiled definitions.
- Advanced return analysis: No strong return evidence observed
