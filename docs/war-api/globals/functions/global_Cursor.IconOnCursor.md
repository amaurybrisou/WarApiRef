# Cursor.IconOnCursor

- Category: Global Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 24 addons

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
| Addons seen in | AnywhereTrainerAdditions, AuctionStats, BankWindowFix, CM_ClosetGoblin, CastSequence, Crusher, DetauntHelper, Enemy |
| Files seen in | AnywhereTrainerAdditions.lua, AuctionAssist.lua, ClosetGoblinCharacterWindow.lua, Code/Assist/Assist.lua, Code/Core/ConfigurationWindow.lua, Code/ScenarioInfo/ScenarioInfo.lua, Code/UnitFrames/ClickCasting.lua, Configuration/WCDBConfig_Filters.lua |
| Namespaces detected | Cursor |
| Source kinds | lua_calls |
| Example locations | AnywhereTrainerAdditions: EquipmentLButtonDown, AuctionStats: AuctionHouseWindowCreateSearchItemName_LButtonDown, AuctionStats: EA_Window_Backpack_EquipmentLButtonDown_Override, AuctionStats: OnLButtonUpItem, BankWindowFix: BankEquipmentRButtonDown, CM_ClosetGoblin: HandleDrag |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 53 |
| Global usage count | 53 |
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
Cursor.IconOnCursor()
```

## Description

Observed as a global function across 24 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |

## Returns

- Not confidently inferable from contract artifacts alone.

## Side Effects

- No side effect is confidently inferable from contract artifacts alone.

## Seen In

- AnywhereTrainerAdditions
- AuctionStats
- BankWindowFix
- CM_ClosetGoblin
- CastSequence
- Crusher
- DetauntHelper
- Enemy
- FixGit
- FozAuction
- HotbarMorale
- KeyBar
- NerfedButtons
- RandomMount
- Sequencer
- Shinies
- SquaredClick
- TastyButtons
- TidyRoll
- Twister
- Wikki's Cooldown Bar
- Wikki's Cooldown Pulse
- bigger_MacroWindow
- zMailMod

## Examples

- AnywhereTrainerAdditions: EquipmentLButtonDown -> Cursor.IconOnCursor()
- AuctionStats: AuctionHouseWindowCreateSearchItemName_LButtonDown -> Cursor.IconOnCursor()
- AuctionStats: EA_Window_Backpack_EquipmentLButtonDown_Override -> Cursor.IconOnCursor()
- AuctionStats: OnLButtonUpItem -> Cursor.IconOnCursor()
- BankWindowFix: BankEquipmentRButtonDown -> Cursor.IconOnCursor()
- CM_ClosetGoblin: HandleDrag -> Cursor.IconOnCursor()

## Related APIs

- [OnLButtonUp](../../xml/handlers/handler_OnLButtonUp.md) (HIGH 100/100) - XML Event
- [OnLButtonDown](../../xml/handlers/handler_OnLButtonDown.md) (HIGH 88/100) - XML Event
- [OnMouseOver](../../xml/handlers/handler_OnMouseOver.md) (HIGH 88/100) - XML Event

## Used With

- [Cursor.Clear](global_Cursor.Clear.md) (HIGH 100/100) - Global Function
- [Cursor.PickUp](global_Cursor.PickUp.md) (HIGH 100/100) - Global Function
- [EA_BackpackUtilsMediator.GetCurrentBackpackType](global_EA_BackpackUtilsMediator.GetCurrentBackpackType.md) (HIGH 100/100) - Global Function
- [EA_BackpackUtilsMediator.GetCursorForBackpack](global_EA_BackpackUtilsMediator.GetCursorForBackpack.md) (HIGH 100/100) - Global Function
- [EA_BackpackUtilsMediator.GetItemsFromBackpack](global_EA_BackpackUtilsMediator.GetItemsFromBackpack.md) (HIGH 100/100) - Global Function
- [EA_ChatWindow.InsertItemLink](global_EA_ChatWindow.InsertItemLink.md) (HIGH 100/100) - Global Function
- [EA_Window_Macro.UpdateDetails](global_EA_Window_Macro.UpdateDetails.md) (HIGH 100/100) - Global Function
- [SystemData.TrialAlert.ALERT_AUCTION](../../systemdata/fields/systemdata_SystemData.TrialAlert.ALERT_AUCTION.md) (HIGH 100/100) - SystemData Field
- [EA_BackpackUtilsMediator.ReleaseLockForSlot](global_EA_BackpackUtilsMediator.ReleaseLockForSlot.md) (HIGH 90/100) - Global Function
- [DoesWindowExist](global_DoesWindowExist.md) (HIGH 83/100) - Global Function
- [BankWindow.GetItem](global_BankWindow.GetItem.md) (HIGH 71/100) - Global Function
- [BankWindow.GetSlotNumberForButtonIndex](global_BankWindow.GetSlotNumberForButtonIndex.md) (HIGH 71/100) - Global Function

## Affects

- [SystemData.ActiveWindow.name](../../systemdata/fields/systemdata_SystemData.ActiveWindow.name.md) (HIGH 100/100) - SystemData Field
- [SystemData.MouseOverWindow.name](../../systemdata/fields/systemdata_SystemData.MouseOverWindow.name.md) (HIGH 100/100) - SystemData Field

## Notes

- Canonical entry built from observed call sites, not from engine source or decompiled definitions.
