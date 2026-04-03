# EA_Window_ContextMenu.AddCascadingMenuItem

- Category: Global Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 12 addons

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 110

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, seen in 4 or more addons, matches a known engine namespace.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- -15 Conflicting signatures across usages: Observed arity or argument shape conflicts across usages.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | ActionFraction, AdjustTheTip, AutoBand, FozAuction, LibAddonButton, MapPin, MarkBuff, MoraleCircle |
| Files seen in | AdjustTheTip.lua, AutoBand.lua, Gui.lua, Manager/Manager.lua, MoraleCircle.lua, QuickNameActionsRessurected.lua, RoR_SoR.lua, Source/auctionwindowsearchcontrols.lua |
| Namespaces detected | EA_Window_ContextMenu |
| Source kinds | lua_calls |
| Example locations | ActionFraction: RightClick, AdjustTheTip: CreateContextMenu, AutoBand: AutoBand_L_AddContextMenuCascade, FozAuction: OnLButtonUpAdditonalFilters, LibAddonButton: CreateRemoveContextMenu, LibAddonButton: ShowOptions |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 37 |
| Global usage count | 37 |
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
EA_Window_ContextMenu.AddCascadingMenuItem(arg1, arg2, arg3, arg4)
```

## Description

Observed as a global function across 12 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a function or method reference. | Observed values: GetStringFormatFromTable("AuctionHouseStrings",StringTables.AuctionHouse.CONTEXT_MENU_ADDITIONAL_FILTERS_CAREER_X,{careerSelectionName}), GetStringFormatFromTable("AuctionHouseStrings",StringTables.AuctionHouse.CONTEXT_MENU_ADDITIONAL_FILTERS_RESTRICTION_X,{restrictionSelectionName}), GetStringFormatFromTable("AuctionHouseStrings",StringTables.AuctionHouse.CONTEXT_MENU_ADDITIONAL_FILTERS_STATISTIC_X,{statisticSelectionName}) |
| arg2 | Observed as a runtime window or control identifier. | Observed values: ActionFractionWindow.SetFontSelectionMenu, ActionFractionWindow.SetPresetLocation, AdjustTheTip.SpawnAnchorContextMenu |
| arg3 | Observed as a boolean toggle. | Observed values: false |
| arg4 | Observed as a numeric value. | Observed values: 1, 2, EA_Window_ContextMenu.CONTEXT_MENU_1 |

## Returns

- Not confidently inferable from contract artifacts alone.

## Side Effects

- No side effect is confidently inferable from contract artifacts alone.

## Seen In

- ActionFraction
- AdjustTheTip
- AutoBand
- FozAuction
- LibAddonButton
- MapPin
- MarkBuff
- MoraleCircle
- QuickNameActions+
- RoR_SoR
- ShowMeTheBubbles
- XpStatus+G

## Examples

- ActionFraction: RightClick -> EA_Window_ContextMenu.AddCascadingMenuItem(LOC_TEXT.PRESET_LOCATIONS, ActionFractionWindow.SetPresetLocation, false, EA_Window_ContextMenu.CONTEXT_MENU_1)
- ActionFraction: RightClick -> EA_Window_ContextMenu.AddCascadingMenuItem(LOC_TEXT.SET_FONT_SIZE, ActionFractionWindow.SetFontSelectionMenu, false, EA_Window_ContextMenu.CONTEXT_MENU_1)
- AdjustTheTip: CreateContextMenu -> EA_Window_ContextMenu.AddCascadingMenuItem(L "Delayed Fading", AdjustTheTip.SpawnDelayedFadingContextMenu, false)
- AdjustTheTip: CreateContextMenu -> EA_Window_ContextMenu.AddCascadingMenuItem(L "Ability Tooltip Anchor", AdjustTheTip.SpawnAnchorContextMenu, false)
- AutoBand: AutoBand_L_AddContextMenuCascade -> EA_Window_ContextMenu.AddCascadingMenuItem(label_w, callback, false, menuId)
- FozAuction: OnLButtonUpAdditonalFilters -> EA_Window_ContextMenu.AddCascadingMenuItem(GetStringFormatFromTable("AuctionHouseStrings",StringTables.AuctionHouse.CONTEXT_MENU_ADDITIONAL_FILTERS_STATISTIC_X,{statisticSelectionName}), SpawnStatisticMenu)

## Used With

- [EA_Window_ContextMenu.AddMenuDivider](global_EA_Window_ContextMenu.AddMenuDivider.md) (HIGH 100/100) - Global Function
- [EA_Window_ContextMenu.AddMenuItem](global_EA_Window_ContextMenu.AddMenuItem.md) (HIGH 100/100) - Global Function
- [EA_Window_ContextMenu.AddUserDefinedMenuItem](global_EA_Window_ContextMenu.AddUserDefinedMenuItem.md) (HIGH 100/100) - Global Function
- [EA_Window_ContextMenu.CreateContextMenu](global_EA_Window_ContextMenu.CreateContextMenu.md) (HIGH 100/100) - Global Function
- [EA_Window_ContextMenu.Finalize](global_EA_Window_ContextMenu.Finalize.md) (HIGH 100/100) - Global Function

## Affects

- [GameData.Auction.RESTRICTION_NONE](../../gamedata/fields/gamedata_GameData.Auction.RESTRICTION_NONE.md) (HIGH 100/100) - GameData Field
- [GameData.BonusTypes.EBONUS_NONE](../../gamedata/fields/gamedata_GameData.BonusTypes.EBONUS_NONE.md) (HIGH 100/100) - GameData Field

## Notes

- Canonical entry built from observed call sites, not from engine source or decompiled definitions.
- Advanced return analysis: No strong return evidence observed
