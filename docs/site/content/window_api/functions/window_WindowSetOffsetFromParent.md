# WindowSetOffsetFromParent

- Category: Window Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 3 addons

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
| Addons seen in | GetStats, MapPin, RoR_SoR |
| Files seen in | `/workspace/GetStats/GetStats.lua:119`, `/workspace/MapPin/source/MapPin.lua:1122`, `/workspace/MapPin/source/MapPin.lua:144`, `/workspace/RoR_SoR/RoR_SoR.lua:2881` |
| Namespaces detected | WindowSetOffsetFromParent |
| Source kinds | lua_calls |
| Example locations | GetStats: GetStats.OnChatLogUpdated, MapPin: MapPin.local.OnHyperLinkLButtonUp2, MapPin: MapPin.test, MapPin: OnHyperLinkLButtonUp2, RoR_SoR: RoR_SoR.DefaultSettings |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 8 |
| Global usage count | 8 |
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
WindowSetOffsetFromParent(windowName, arg2, arg3)
```

## Description

Observed mutating runtime window state or presentation.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| windowName | Observed as a target window name. | Observed values: "GetStatsCompareLine"..i, "GetStatsLine"..GetStats.LineNumber, "MapPinCommonPin" |
| arg2 | Observed as a runtime window or control identifier. | Observed values: 350, 5, tonumber(Splitter[4]) |
| arg3 | Observed as a runtime window or control identifier. | Observed values: 100, 6+(tonumber(GetStats.LineNumber)*GetStats.RowHeight), 6+(tonumber(i)*GetStats.RowHeight) |

## Returns

- Not confidently inferable from API_Ref alone.

## Side Effects

- Mutates engine-visible UI state or presentation.

## Seen In

- GetStats
- MapPin
- RoR_SoR

## Examples

- GetStats: GetStats.OnChatLogUpdated -> WindowSetOffsetFromParent("GetStatsCompareLine"..i, 5, 6+(tonumber(i)*GetStats.RowHeight))
- GetStats: GetStats.OnChatLogUpdated -> WindowSetOffsetFromParent("GetStatsLine"..GetStats.LineNumber, 5, 6+(tonumber(GetStats.LineNumber)*GetStats.RowHeight))
- MapPin: MapPin.local.OnHyperLinkLButtonUp2 -> WindowSetOffsetFromParent("MapPinWBPin"..WBid, tonumber(Splitter[4]), tonumber(Splitter[5]))
- MapPin: MapPin.local.OnHyperLinkLButtonUp2 -> WindowSetOffsetFromParent("MapPinCommonPin", tonumber(Splitter[4]), tonumber(Splitter[5]))
- MapPin: MapPin.test -> WindowSetOffsetFromParent("MapPinWBPin"..WBid, tonumber(Splitter[4]), tonumber(Splitter[5]))
- MapPin: OnHyperLinkLButtonUp2 -> WindowSetOffsetFromParent("MapPinWBPin"..WBid, tonumber(Splitter[4]), tonumber(Splitter[5]))

## Related APIs

- none

## Used With

- [AlertText](../../globals/tables/table_AlertText.md) (HIGH 100/100) - Global Table
- [DialogManager.MakeTextEntryDialog](../../globals/functions/global_DialogManager.MakeTextEntryDialog.md) (HIGH 100/100) - Global Function
- [EA_ChatTabManager.GetTabName](../../globals/functions/global_EA_ChatTabManager.GetTabName.md) (HIGH 100/100) - Global Function
- [EA_ChatWindowGroups](../../globals/tables/table_EA_ChatWindowGroups.md) (HIGH 100/100) - Global Table
- [GameData.ChatData.name](../../gamedata/fields/gamedata_GameData.ChatData.name.md) (HIGH 100/100) - GameData Field
- [GameData.ChatData.text](../../gamedata/fields/gamedata_GameData.ChatData.text.md) (HIGH 100/100) - GameData Field
- [GameData.Sound.QUEST_ABANDONED](../../gamedata/fields/gamedata_GameData.Sound.QUEST_ABANDONED.md) (HIGH 100/100) - GameData Field
- [LabelSetText](window_LabelSetText.md) (HIGH 100/100) - Window Function
- [PartyUtils.GetWarbandLeader](../../globals/functions/global_PartyUtils.GetWarbandLeader.md) (HIGH 100/100) - Global Function
- [PartyUtils.IsPartyActive](../../globals/functions/global_PartyUtils.IsPartyActive.md) (HIGH 100/100) - Global Function
- [SystemData.ActiveWindow.name](../../systemdata/fields/systemdata_SystemData.ActiveWindow.name.md) (HIGH 100/100) - SystemData Field
- [WindowGetId](window_WindowGetId.md) (HIGH 100/100) - Window Function
- [WindowSetParent](window_WindowSetParent.md) (HIGH 100/100) - Window Function
- [WindowSetShowing](window_WindowSetShowing.md) (HIGH 100/100) - Window Function
- [wstring.match](../../globals/functions/global_wstring.match.md) (HIGH 100/100) - Global Function
- [AlertTextWindow.AddAlert](../../globals/functions/global_AlertTextWindow.AddAlert.md) (HIGH 88/100) - Global Function
- [DoesWindowExist](../../globals/functions/global_DoesWindowExist.md) (HIGH 83/100) - Global Function
- [CreateWindowFromTemplate](../../globals/functions/global_CreateWindowFromTemplate.md) (HIGH 75/100) - Global Function
- [DestroyWindow](../../globals/functions/global_DestroyWindow.md) (HIGH 75/100) - Global Function
- [towstring](../../globals/functions/global_towstring.md) (HIGH 75/100) - Global Function
- [wstring.sub](../../globals/functions/global_wstring.sub.md) (HIGH 75/100) - Global Function

## Triggered By

- none

## Affects

- [AlertText](../../globals/tables/table_AlertText.md) (HIGH 100/100) - Global Table
- [EA_ChatTabManager.GetTabName](../../globals/functions/global_EA_ChatTabManager.GetTabName.md) (HIGH 100/100) - Global Function
- [EA_ChatWindowGroups](../../globals/tables/table_EA_ChatWindowGroups.md) (HIGH 100/100) - Global Table
- [GameData.ChatData.name](../../gamedata/fields/gamedata_GameData.ChatData.name.md) (HIGH 100/100) - GameData Field
- [GameData.ChatData.text](../../gamedata/fields/gamedata_GameData.ChatData.text.md) (HIGH 100/100) - GameData Field
- [GameData.Sound.QUEST_ABANDONED](../../gamedata/fields/gamedata_GameData.Sound.QUEST_ABANDONED.md) (HIGH 100/100) - GameData Field
- [SystemData.ActiveWindow.name](../../systemdata/fields/systemdata_SystemData.ActiveWindow.name.md) (HIGH 100/100) - SystemData Field
- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Notes

- none
