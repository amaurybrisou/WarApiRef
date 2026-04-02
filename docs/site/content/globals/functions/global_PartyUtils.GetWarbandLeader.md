# PartyUtils.GetWarbandLeader

- Category: Global Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 3 addons

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 121

- Rationale: Promoted as HIGH confidence because matches a known engine namespace, referenced by generated docs or reference files, called globally with no local definition.

## Evidence Signals

- +18 Seen in 2 to 3 addons: Cross-addon spread is present but limited.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +10 Argument pattern is consistent: Observed argument positions remain stable.
- +8 Return usage is consistent: Observed as a stable query-style API.
- +25 Referenced by generated docs or reference files: The symbol is reinforced outside a single call page.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | MapPin, QuickWarReport, followTheLeader |
| Files seen in | `/workspace_addons/MapPin/source/MapPin.lua:1122`, `/workspace_addons/MapPin/source/MapPin.lua:2280`, `/workspace_addons/MapPin/source/MapPin.lua:3008`, `/workspace_addons/MapPin/source/MapPin.lua:3011`, `/workspace_addons/MapPin/source/MapPin.lua:4452`, `/workspace_addons/QuickWarReport/QWRComms.lua:28`, `/workspace_addons/followTheLeader/followTheLeader.lua:209` |
| Namespaces detected | PartyUtils |
| Source kinds | globals, lua_calls |
| Example locations | MapPin: MapPin.CreateRequestContextMenu, MapPin: MapPin.LButtonUp, MapPin: MapPin.QuickRequest, MapPin: MapPin.local.RequestCommand, MapPin: MapPin.test, MapPin: RequestCommand |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 9 |
| Global usage count | 9 |
| Local definition count | 0 |
| Documentation references | 1 |
| Initialization flow references | 0 |
| Known engine namespace | yes |
| Default UI presence | no |
| Event binding presence | no |
| Observed in XML and Lua | no |
| Consistent role | yes |
| Consistent arguments | yes |
| Consistent returns | yes |
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
PartyUtils.GetWarbandLeader()
```

## Description

Observed as a global function across 3 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |

## Returns

- Observed as a query-style API. The concrete return shape is not inferable from addon-api docs alone.

## Side Effects

- No side effect is confidently inferable from addon-api docs alone.

## Seen In

- MapPin
- QuickWarReport
- followTheLeader

## Examples

- MapPin: MapPin.CreateRequestContextMenu -> PartyUtils.GetWarbandLeader()
- MapPin: MapPin.LButtonUp -> PartyUtils.GetWarbandLeader()
- MapPin: MapPin.QuickRequest -> PartyUtils.GetWarbandLeader()
- MapPin: MapPin.local.RequestCommand -> PartyUtils.GetWarbandLeader()
- MapPin: MapPin.test -> PartyUtils.GetWarbandLeader()
- MapPin: RequestCommand -> PartyUtils.GetWarbandLeader()

## Related APIs

- [EA_ChatWindow.InsertText](global_EA_ChatWindow.InsertText.md) (HIGH 100/100) - Global Function
- [EA_Window_WorldMap.ShowZone](global_EA_Window_WorldMap.ShowZone.md) (HIGH 100/100) - Global Function
- [PartyUtils.GetPartyData](global_PartyUtils.GetPartyData.md) (HIGH 100/100) - Global Function
- [PartyUtils.IsPartyActive](global_PartyUtils.IsPartyActive.md) (HIGH 100/100) - Global Function

## Used With

- [AlertText](../tables/table_AlertText.md) (HIGH 100/100) - Global Table
- [DialogManager.MakeTextEntryDialog](global_DialogManager.MakeTextEntryDialog.md) (HIGH 100/100) - Global Function
- [EA_ChatWindow.InsertText](global_EA_ChatWindow.InsertText.md) (HIGH 100/100) - Global Function
- [EA_Window_ContextMenu.AddMenuDivider](global_EA_Window_ContextMenu.AddMenuDivider.md) (HIGH 100/100) - Global Function
- [EA_Window_ContextMenu.AddMenuItem](global_EA_Window_ContextMenu.AddMenuItem.md) (HIGH 100/100) - Global Function
- [EA_Window_ContextMenu.CreateContextMenu](global_EA_Window_ContextMenu.CreateContextMenu.md) (HIGH 100/100) - Global Function
- [EA_Window_ContextMenu.Finalize](global_EA_Window_ContextMenu.Finalize.md) (HIGH 100/100) - Global Function
- [EA_Window_WorldMap](../tables/table_EA_Window_WorldMap.md) (HIGH 100/100) - Global Table
- [EA_Window_WorldMap.ShowZone](global_EA_Window_WorldMap.ShowZone.md) (HIGH 100/100) - Global Function
- [GameData.ChatData.name](../../gamedata/fields/gamedata_GameData.ChatData.name.md) (HIGH 100/100) - GameData Field
- [GameData.ChatData.text](../../gamedata/fields/gamedata_GameData.ChatData.text.md) (HIGH 100/100) - GameData Field
- [GameData.Sound.QUEST_ABANDONED](../../gamedata/fields/gamedata_GameData.Sound.QUEST_ABANDONED.md) (HIGH 100/100) - GameData Field
- [Icons.GetCareerIconIDFromCareerLine](global_Icons.GetCareerIconIDFromCareerLine.md) (HIGH 100/100) - Global Function
- [OnLButtonUp](../../events/window_events/window_event_OnLButtonUp.md) (HIGH 100/100) - Window Event
- [OnLButtonUp](../../xml/handlers/handler_OnLButtonUp.md) (HIGH 100/100) - XML Handler
- [PartyUtils.IsPartyActive](global_PartyUtils.IsPartyActive.md) (HIGH 100/100) - Global Function
- [SystemData.MouseOverWindow.name](../../systemdata/fields/systemdata_SystemData.MouseOverWindow.name.md) (HIGH 100/100) - SystemData Field
- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type
- [WindowGetId](../../window_api/functions/window_WindowGetId.md) (HIGH 100/100) - Window Function
- [WindowGetScale](../../window_api/functions/window_WindowGetScale.md) (HIGH 100/100) - Window Function
- [WindowGetShowing](../../window_api/functions/window_WindowGetShowing.md) (HIGH 100/100) - Window Function
- [WindowSetOffsetFromParent](../../window_api/functions/window_WindowSetOffsetFromParent.md) (HIGH 100/100) - Window Function
- [WindowSetShowing](../../window_api/functions/window_WindowSetShowing.md) (HIGH 100/100) - Window Function
- [wstring.match](global_wstring.match.md) (HIGH 100/100) - Global Function
- [AlertTextWindow.AddAlert](global_AlertTextWindow.AddAlert.md) (HIGH 88/100) - Global Function
- [GetIconData](global_GetIconData.md) (HIGH 83/100) - Global Function
- [towstring](global_towstring.md) (HIGH 75/100) - Global Function

## Triggered By

- [OnLButtonUp](../../xml/handlers/handler_OnLButtonUp.md) (HIGH 100/100) - XML Handler
- [OnLButtonUp](../../events/window_events/window_event_OnLButtonUp.md) (HIGH 100/100) - Window Event

## Affects

- [AlertText](../tables/table_AlertText.md) (HIGH 100/100) - Global Table
- [EA_Window_WorldMap](../tables/table_EA_Window_WorldMap.md) (HIGH 100/100) - Global Table
- [GameData.ChatData.name](../../gamedata/fields/gamedata_GameData.ChatData.name.md) (HIGH 100/100) - GameData Field
- [GameData.ChatData.text](../../gamedata/fields/gamedata_GameData.ChatData.text.md) (HIGH 100/100) - GameData Field
- [GameData.Player.zone](../../gamedata/fields/gamedata_GameData.Player.zone.md) (HIGH 100/100) - GameData Field
- [GameData.Sound.QUEST_ABANDONED](../../gamedata/fields/gamedata_GameData.Sound.QUEST_ABANDONED.md) (HIGH 100/100) - GameData Field
- [SystemData.MouseOverWindow.name](../../systemdata/fields/systemdata_SystemData.MouseOverWindow.name.md) (HIGH 100/100) - SystemData Field
- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Notes

- Canonical entry built from observed call sites, not from engine source or decompiled definitions.
