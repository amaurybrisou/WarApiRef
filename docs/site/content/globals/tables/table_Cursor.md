# Cursor

- Category: Global Table
- Confidence level: HIGH
- Confidence score: 100/100

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 130

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, seen in 4 or more addons, matches a known engine namespace.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +25 Referenced by generated docs or reference files: The symbol is reinforced outside a single call page.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | AnywhereTrainerAdditions, BankWindowFix, CM_ClosetGoblin, Enemy, HotbarMorale, RandomMount, Shinies, TidyRoll |
| Files seen in | `/workspace/AnywhereTrainerAdditions/AnywhereTrainerAdditions.lua:218`, `/workspace/BankWindowFix/Source/BankWindowFix.lua:30`, `/workspace/ClosetGoblin/ClosetGoblinCharacterWindow.lua:512`, `/workspace/Enemy/Code/Assist/Assist.lua:363`, `/workspace/Enemy/Code/Core/ConfigurationWindow.lua:65`, `/workspace/Enemy/Code/ScenarioInfo/ScenarioInfo.lua:1447`, `/workspace/Enemy/Code/UnitFrames/ClickCasting.lua:379`, `/workspace/HotbarMorale/HotbarMorale.lua:10` |
| Namespaces detected | Cursor |
| Source kinds | globals, lua_calls |
| Example locations | AnywhereTrainerAdditions: AnywhereTrainerAdditions.EquipmentLButtonDown, BankWindowFix: BankWindowFix.BankEquipmentRButtonDown, CM_ClosetGoblin: ClosetGoblinCharacterWindow.HandleDrag, Enemy: Enemy.ConfigurationWindow_OnMacroMouseDrag, Enemy: Enemy.UnitFramesUI_UnitFrameClickCastingDialog_OnActionConfig1ActionLButtonUp, Enemy: Enemy.local.OnMacroMouseDrag |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 33 |
| Global usage count | 3 |
| Local definition count | 1 |
| Documentation references | 1 |
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
| Conflicting signatures | no |
| Conflicting roles | no |
| Wrapper likely | no |
| Never outside local graph | no |
| Local helper only | no |

## Description

Observed shared global table or namespace surfaced in 8 addons.

## Functions

- Cursor.Clear
- Cursor.IconOnCursor
- Cursor.PickUp

## Observed Members

- none

## Seen In

- AnywhereTrainerAdditions
- BankWindowFix
- CM_ClosetGoblin
- Enemy
- HotbarMorale
- RandomMount
- Shinies
- TidyRoll

## Examples

- AnywhereTrainerAdditions: AnywhereTrainerAdditions.EquipmentLButtonDown -> Cursor.IconOnCursor()
- AnywhereTrainerAdditions: AnywhereTrainerAdditions.EquipmentLButtonDown -> Cursor.Clear()
- AnywhereTrainerAdditions: AnywhereTrainerAdditions.EquipmentLButtonDown -> Cursor.PickUp(Cursor.SOURCE_BANK, slot, itemData.uniqueID, itemData.iconNum, true)
- BankWindowFix: BankWindowFix.BankEquipmentRButtonDown -> Cursor.IconOnCursor()
- CM_ClosetGoblin: ClosetGoblinCharacterWindow.HandleDrag -> Cursor.IconOnCursor()
- CM_ClosetGoblin: ClosetGoblinCharacterWindow.HandleDrag -> Cursor.Clear()

## Related APIs

- [EA_BackpackUtilsMediator.GetCurrentBackpackType](../functions/global_EA_BackpackUtilsMediator.GetCurrentBackpackType.md) (HIGH 100/100) - Global Function
- [EA_BackpackUtilsMediator.GetCursorForBackpack](../functions/global_EA_BackpackUtilsMediator.GetCursorForBackpack.md) (HIGH 100/100) - Global Function
- [EA_BackpackUtilsMediator.GetItemsFromBackpack](../functions/global_EA_BackpackUtilsMediator.GetItemsFromBackpack.md) (HIGH 100/100) - Global Function

## Used With

- [Cursor.Clear](../functions/global_Cursor.Clear.md) (HIGH 100/100) - Global Function
- [Cursor.IconOnCursor](../functions/global_Cursor.IconOnCursor.md) (HIGH 100/100) - Global Function
- [EA_BackpackUtilsMediator.GetCurrentBackpackType](../functions/global_EA_BackpackUtilsMediator.GetCurrentBackpackType.md) (HIGH 100/100) - Global Function
- [EA_BackpackUtilsMediator.GetCursorForBackpack](../functions/global_EA_BackpackUtilsMediator.GetCursorForBackpack.md) (HIGH 100/100) - Global Function
- [EA_BackpackUtilsMediator.GetItemsFromBackpack](../functions/global_EA_BackpackUtilsMediator.GetItemsFromBackpack.md) (HIGH 100/100) - Global Function
- [ListBox](../../xml/element_types/element_ListBox.md) (HIGH 100/100) - XML Element Type
- [OnLButtonUp](../../events/window_events/window_event_OnLButtonUp.md) (HIGH 100/100) - Window Event
- [OnLButtonUp](../../xml/handlers/handler_OnLButtonUp.md) (HIGH 100/100) - XML Handler

## Triggered By

- none

## Affects

- none

## Notes

- none
