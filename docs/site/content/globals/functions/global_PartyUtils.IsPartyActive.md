# PartyUtils.IsPartyActive

- Category: Global Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 6 addons

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 133

- Rationale: Promoted as HIGH confidence because seen in 4 or more addons, matches a known engine namespace, referenced by generated docs or reference files.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +10 Argument pattern is consistent: Observed argument positions remain stable.
- +8 Return usage is consistent: Observed as a stable query-style API.
- +25 Referenced by generated docs or reference files: The symbol is reinforced outside a single call page.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | GuardLine, LibGuard, MapPin, QuickWarReport, RoR_SoR, Swift Assist |
| Files seen in | `/workspace_addons/GuardLine/GuardLine.lua:197`, `/workspace_addons/LibGuard/Source/LibGuard.lua:186`, `/workspace_addons/LibGuard/Source/LibGuard.lua:362`, `/workspace_addons/MapPin/source/MapPin.lua:1122`, `/workspace_addons/MapPin/source/MapPin.lua:2280`, `/workspace_addons/MapPin/source/MapPin.lua:2861`, `/workspace_addons/MapPin/source/MapPin.lua:3008`, `/workspace_addons/MapPin/source/MapPin.lua:3011` |
| Namespaces detected | PartyUtils |
| Source kinds | globals, lua_calls |
| Example locations | GuardLine: GuardLine.update, LibGuard: LibGuard.GetIdFromName, LibGuard: LibGuard.UpdateStateMachine, MapPin: MapPin.CreateMainContextMenu, MapPin: MapPin.CreateRequestContextMenu, MapPin: MapPin.LButtonUp |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 15 |
| Global usage count | 15 |
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
PartyUtils.IsPartyActive()
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

- GuardLine
- LibGuard
- MapPin
- QuickWarReport
- RoR_SoR
- Swift Assist

## Examples

- GuardLine: GuardLine.update -> PartyUtils.IsPartyActive()
- LibGuard: LibGuard.GetIdFromName -> PartyUtils.IsPartyActive()
- LibGuard: LibGuard.UpdateStateMachine -> PartyUtils.IsPartyActive()
- MapPin: MapPin.CreateMainContextMenu -> PartyUtils.IsPartyActive()
- MapPin: MapPin.CreateRequestContextMenu -> PartyUtils.IsPartyActive()
- MapPin: MapPin.LButtonUp -> PartyUtils.IsPartyActive()

## Related APIs

- none

## Used With

- [AlertText](../tables/table_AlertText.md) (HIGH 100/100) - Global Table
- [DialogManager.MakeTextEntryDialog](global_DialogManager.MakeTextEntryDialog.md) (HIGH 100/100) - Global Function
- [EA_ChatWindow.InsertText](global_EA_ChatWindow.InsertText.md) (HIGH 100/100) - Global Function
- [EA_Window_ContextMenu.AddCascadingMenuItem](global_EA_Window_ContextMenu.AddCascadingMenuItem.md) (HIGH 100/100) - Global Function
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
- [PartyUtils.GetPartyData](global_PartyUtils.GetPartyData.md) (HIGH 100/100) - Global Function
- [PartyUtils.GetWarbandData](global_PartyUtils.GetWarbandData.md) (HIGH 100/100) - Global Function
- [PartyUtils.GetWarbandLeader](global_PartyUtils.GetWarbandLeader.md) (HIGH 100/100) - Global Function
- [SystemData.MouseOverWindow.name](../../systemdata/fields/systemdata_SystemData.MouseOverWindow.name.md) (HIGH 100/100) - SystemData Field
- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type
- [WindowGetId](../../window_api/functions/window_WindowGetId.md) (HIGH 100/100) - Window Function
- [WindowGetScale](../../window_api/functions/window_WindowGetScale.md) (HIGH 100/100) - Window Function
- [WindowGetShowing](../../window_api/functions/window_WindowGetShowing.md) (HIGH 100/100) - Window Function
- [WindowSetOffsetFromParent](../../window_api/functions/window_WindowSetOffsetFromParent.md) (HIGH 100/100) - Window Function
- [WindowSetShowing](../../window_api/functions/window_WindowSetShowing.md) (HIGH 100/100) - Window Function
- [WindowStartAlphaAnimation](../../window_api/functions/window_WindowStartAlphaAnimation.md) (HIGH 100/100) - Window Function
- [wstring.match](global_wstring.match.md) (HIGH 100/100) - Global Function
- [WindowStartScaleAnimation](../../window_api/functions/window_WindowStartScaleAnimation.md) (HIGH 90/100) - Window Function
- [AlertTextWindow.AddAlert](global_AlertTextWindow.AddAlert.md) (HIGH 88/100) - Global Function
- [GetIconData](global_GetIconData.md) (HIGH 83/100) - Global Function
- [towstring](global_towstring.md) (HIGH 75/100) - Global Function

## Triggered By

- [OnLButtonUp](../../xml/handlers/handler_OnLButtonUp.md) (HIGH 100/100) - XML Handler
- [OnLButtonUp](../../events/window_events/window_event_OnLButtonUp.md) (HIGH 100/100) - Window Event

## Affects

- [EA_Window_WorldMap](../tables/table_EA_Window_WorldMap.md) (HIGH 100/100) - Global Table
- [GameData.Player.Pet.name](../../gamedata/fields/gamedata_GameData.Player.Pet.name.md) (HIGH 100/100) - GameData Field
- [GameData.Player.name](../../gamedata/fields/gamedata_GameData.Player.name.md) (HIGH 100/100) - GameData Field
- [PartyUtils.GetPartyData](global_PartyUtils.GetPartyData.md) (HIGH 100/100) - Global Function
- [PartyUtils.GetWarbandData](global_PartyUtils.GetWarbandData.md) (HIGH 100/100) - Global Function
- [PartyUtils.GetWarbandLeader](global_PartyUtils.GetWarbandLeader.md) (HIGH 100/100) - Global Function
- [SystemData.MouseOverWindow.name](../../systemdata/fields/systemdata_SystemData.MouseOverWindow.name.md) (HIGH 100/100) - SystemData Field
- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Notes

- Canonical entry built from observed call sites, not from engine source or decompiled definitions.
