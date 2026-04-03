# Cursor

- Category: Global Table
- Confidence level: HIGH
- Confidence score: 100/100

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 125

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, seen in 4 or more addons, matches a known engine namespace.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | AnywhereTrainerAdditions, AuctionStats, BankWindowFix, CM_ClosetGoblin, CastSequence, Crusher, DetauntHelper, Enemy |
| Files seen in | Code/Assist/Assist.lua, Code/Core/ConfigurationWindow.lua, Code/UnitFrames/ClickCasting.lua, Configuration/WCDBConfig_Filters.lua, Configuration/WCDPConfig_Filters.lua, Modules/UI/Shinies-UI-Auto.lua, Modules/UI/Shinies-UI-Browse.lua, Modules/UI/Shinies-UI-Post.lua |
| Namespaces detected | Cursor |
| Source kinds | lua_calls |
| Example locations | AnywhereTrainerAdditions: EquipmentLButtonDown, AuctionStats: AuctionHouseWindowCreateSearchItemName_LButtonDown, AuctionStats: EA_Window_Backpack_EquipmentLButtonDown_Override, AuctionStats: OnLButtonUpItem, BankWindowFix: BankEquipmentRButtonDown, CM_ClosetGoblin: HandleDrag |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 89 |
| Global usage count | 3 |
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
| Conflicting signatures | no |
| Conflicting roles | no |
| Wrapper likely | no |
| Never outside local graph | no |
| Local helper only | no |

## Description

Shared function table with 3 member functions; the primary API surface for 24 addons.

## Functions

- Cursor.Clear
- Cursor.IconOnCursor
- Cursor.PickUp

## Observed Members

- none

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
- AnywhereTrainerAdditions: EquipmentLButtonDown -> Cursor.Clear()
- AnywhereTrainerAdditions: EquipmentLButtonDown -> Cursor.PickUp(Cursor.SOURCE_BANK, slot, itemData.uniqueID, itemData.iconNum, true)
- AuctionStats: AuctionHouseWindowCreateSearchItemName_LButtonDown -> Cursor.IconOnCursor()
- AuctionStats: AuctionHouseWindowCreateSearchItemName_LButtonDown -> Cursor.Clear()
- AuctionStats: EA_Window_Backpack_EquipmentLButtonDown_Override -> Cursor.IconOnCursor()

## Notes

- none
