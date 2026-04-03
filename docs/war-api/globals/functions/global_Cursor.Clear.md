# Cursor.Clear

- Category: Global Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 22 addons

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
| Addons seen in | AnywhereTrainerAdditions, AuctionStats, CM_ClosetGoblin, CastSequence, Crusher, DetauntHelper, Enemy, FixGit |
| Files seen in | AnywhereTrainerAdditions.lua, AuctionAssist.lua, ClosetGoblinCharacterWindow.lua, Code/UnitFrames/ClickCasting.lua, Configuration/WCDBConfig_Filters.lua, Configuration/WCDPConfig_Filters.lua, CustomAutoRoll.lua, Fixes.lua |
| Namespaces detected | Cursor |
| Source kinds | lua_calls |
| Example locations | AnywhereTrainerAdditions: EquipmentLButtonDown, AuctionStats: AuctionHouseWindowCreateSearchItemName_LButtonDown, AuctionStats: OnLButtonUpItem, CM_ClosetGoblin: HandleDrag, CastSequence: OnLButtonUp, Crusher: OnLButtonUp |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 47 |
| Global usage count | 47 |
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
Cursor.Clear()
```

## Description

Observed as a global function across 22 addons.

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
- zMailMod

## Examples

- AnywhereTrainerAdditions: EquipmentLButtonDown -> Cursor.Clear()
- AuctionStats: AuctionHouseWindowCreateSearchItemName_LButtonDown -> Cursor.Clear()
- AuctionStats: OnLButtonUpItem -> Cursor.Clear()
- CM_ClosetGoblin: HandleDrag -> Cursor.Clear()
- CastSequence: OnLButtonUp -> Cursor.Clear()
- Crusher: OnLButtonUp -> Cursor.Clear()

## Related APIs

- [OnLButtonUp](../../xml/handlers/handler_OnLButtonUp.md) (HIGH 100/100) - XML Event
- [OnLButtonDown](../../xml/handlers/handler_OnLButtonDown.md) (HIGH 88/100) - XML Event

## Used With

- [Cursor.IconOnCursor](global_Cursor.IconOnCursor.md) (HIGH 100/100) - Global Function
- [Cursor.PickUp](global_Cursor.PickUp.md) (HIGH 100/100) - Global Function
- [EA_BackpackUtilsMediator.GetCurrentBackpackType](global_EA_BackpackUtilsMediator.GetCurrentBackpackType.md) (HIGH 100/100) - Global Function
- [EA_BackpackUtilsMediator.GetCursorForBackpack](global_EA_BackpackUtilsMediator.GetCursorForBackpack.md) (HIGH 100/100) - Global Function
- [EA_BackpackUtilsMediator.GetItemsFromBackpack](global_EA_BackpackUtilsMediator.GetItemsFromBackpack.md) (HIGH 100/100) - Global Function
- [EA_ChatWindow.InsertItemLink](global_EA_ChatWindow.InsertItemLink.md) (HIGH 100/100) - Global Function
- [SystemData.TrialAlert.ALERT_AUCTION](../../systemdata/fields/systemdata_SystemData.TrialAlert.ALERT_AUCTION.md) (HIGH 100/100) - SystemData Field
- [WindowSetShowing](../../window_api/functions/window_WindowSetShowing.md) (HIGH 100/100) - Window Function
- [DoesWindowExist](global_DoesWindowExist.md) (HIGH 83/100) - Global Function
- [BankWindow.GetItem](global_BankWindow.GetItem.md) (HIGH 71/100) - Global Function
- [BankWindow.GetSlotNumberForButtonIndex](global_BankWindow.GetSlotNumberForButtonIndex.md) (HIGH 71/100) - Global Function

## Affects

- [SystemData.ActiveWindow.name](../../systemdata/fields/systemdata_SystemData.ActiveWindow.name.md) (HIGH 100/100) - SystemData Field
- [SystemData.MouseOverWindow.name](../../systemdata/fields/systemdata_SystemData.MouseOverWindow.name.md) (HIGH 100/100) - SystemData Field

## Notes

- Canonical entry built from observed call sites, not from engine source or decompiled definitions.
- Advanced return analysis: No strong return evidence observed
