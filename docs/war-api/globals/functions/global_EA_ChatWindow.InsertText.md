# EA_ChatWindow.InsertText

- Category: Global Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 2 addons

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 148

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, matches a known engine namespace, referenced by generated docs or reference files.

## Evidence Signals

- +18 Seen in 2 to 3 addons: Cross-addon spread is present but limited.
- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +10 Argument pattern is consistent: Observed argument positions remain stable.
- +25 Referenced by generated docs or reference files: The symbol is reinforced outside a single call page.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | AdvancedRenownTrainer, MapPin |
| Files seen in | `/workspace_addons/MapPin/source/MapPin.lua:2280`, `/workspace_addons/MapPin/source/MapPin.lua:2465`, `/workspace_addons/MapPin/source/MapPin.lua:2602`, `/workspace_addons/advancedrenowntrainer/AdvancedRenownTrainingImportExport.lua:225` |
| Namespaces detected | EA_ChatWindow |
| Source kinds | globals, lua_calls |
| Example locations | AdvancedRenownTrainer: AdvancedRenownTraining.ExportToLink, MapPin: MapPin.LButtonUp, MapPin: MapPin.RButtonUp, MapPin: MapPin.local.SendLink, MapPin: SendLink |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 7 |
| Global usage count | 7 |
| Local definition count | 0 |
| Documentation references | 1 |
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
EA_ChatWindow.InsertText(arg1)
```

## Description

Observed as a global function across 2 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a function or method reference. | Observed values: CreateHyperLink(AddLink,L "WayPoints: "..towstring(GetStringFromTable("ZoneNames",MapPin.Pins.WayPoint[1].Zone))..L " ["..towstring(MapCordsX)..L "K,"..towstring(MapCordsY)..L "K]",{30,255,100},{}), CreateHyperLink(DataText,FormatTitle(PinData.Title)..L ": "..towstring(GetStringFromTable("ZoneNames",PinData.Zone))..L " ["..towstring(MapCordsX)..L "K,"..towstring(MapCordsY)..L "K]",{150,200,255},{}), hl |

## Returns

- Not confidently inferable from addon-api docs alone.

## Side Effects

- No side effect is confidently inferable from addon-api docs alone.

## Seen In

- AdvancedRenownTrainer
- MapPin

## Examples

- AdvancedRenownTrainer: AdvancedRenownTraining.ExportToLink -> EA_ChatWindow.InsertText(hl)
- MapPin: MapPin.LButtonUp -> EA_ChatWindow.InsertText(CreateHyperLink(DataText,FormatTitle(PinData.Title)..L ": "..towstring(GetStringFromTable("ZoneNames",PinData.Zone))..L " ["..towstring(MapCordsX)..L "K,"..towstring(MapCordsY)..L "K]",{150,200,255},{}))
- MapPin: MapPin.LButtonUp -> EA_ChatWindow.InsertText(CreateHyperLink(AddLink,L "WayPoints: "..towstring(GetStringFromTable("ZoneNames",MapPin.Pins.WayPoint[1].Zone))..L " ["..towstring(MapCordsX)..L "K,"..towstring(MapCordsY)..L "K]",{30,255,100},{}))
- MapPin: MapPin.RButtonUp -> EA_ChatWindow.InsertText(CreateHyperLink(AddLink,L "WayPoints: "..towstring(GetStringFromTable("ZoneNames",MapPin.Pins.WayPoint[1].Zone))..L " ["..towstring(MapCordsX)..L "K,"..towstring(MapCordsY)..L "K]",{30,255,100},{}))
- MapPin: MapPin.local.SendLink -> EA_ChatWindow.InsertText(CreateHyperLink(AddLink,L "WayPoints: "..towstring(GetStringFromTable("ZoneNames",MapPin.Pins.WayPoint[1].Zone))..L " ["..towstring(MapCordsX)..L "K,"..towstring(MapCordsY)..L "K]",{30,255,100},{}))
- MapPin: SendLink -> EA_ChatWindow.InsertText(CreateHyperLink(AddLink,L "WayPoints: "..towstring(GetStringFromTable("ZoneNames",MapPin.Pins.WayPoint[1].Zone))..L " ["..towstring(MapCordsX)..L "K,"..towstring(MapCordsY)..L "K]",{30,255,100},{}))

## Related APIs

- none

## Used With

- [Button](../../xml/element_types/element_Button.md) (HIGH 100/100) - XML Element Type
- [ComboBoxGetSelectedText](../../window_api/functions/window_ComboBoxGetSelectedText.md) (HIGH 100/100) - Window Function
- [DialogManager.MakeTextEntryDialog](global_DialogManager.MakeTextEntryDialog.md) (HIGH 100/100) - Global Function
- [EA_Window_WorldMap](../tables/table_EA_Window_WorldMap.md) (HIGH 100/100) - Global Table
- [EA_Window_WorldMap.ShowZone](global_EA_Window_WorldMap.ShowZone.md) (HIGH 100/100) - Global Function
- [Icons.GetCareerIconIDFromCareerLine](global_Icons.GetCareerIconIDFromCareerLine.md) (HIGH 100/100) - Global Function
- [OnLButtonUp](../../events/window_events/window_event_OnLButtonUp.md) (HIGH 100/100) - Window Event
- [OnLButtonUp](../../xml/handlers/handler_OnLButtonUp.md) (HIGH 100/100) - XML Handler
- [PartyUtils.GetWarbandLeader](global_PartyUtils.GetWarbandLeader.md) (HIGH 100/100) - Global Function
- [PartyUtils.IsPartyActive](global_PartyUtils.IsPartyActive.md) (HIGH 100/100) - Global Function
- [SystemData.MouseOverWindow.name](../../systemdata/fields/systemdata_SystemData.MouseOverWindow.name.md) (HIGH 100/100) - SystemData Field
- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type
- [WindowGetId](../../window_api/functions/window_WindowGetId.md) (HIGH 100/100) - Window Function
- [WindowGetScale](../../window_api/functions/window_WindowGetScale.md) (HIGH 100/100) - Window Function
- [WindowGetShowing](../../window_api/functions/window_WindowGetShowing.md) (HIGH 100/100) - Window Function
- [WindowSetShowing](../../window_api/functions/window_WindowSetShowing.md) (HIGH 100/100) - Window Function
- [GetIconData](global_GetIconData.md) (HIGH 83/100) - Global Function
- [towstring](global_towstring.md) (HIGH 75/100) - Global Function

## Triggered By

- [OnHyperLinkRButtonUp](../../xml/handlers/handler_OnHyperLinkRButtonUp.md) (HIGH 100/100) - XML Handler
- [OnHyperLinkRButtonUp](../../events/window_events/window_event_OnHyperLinkRButtonUp.md) (HIGH 100/100) - Window Event
- [OnLButtonUp](../../xml/handlers/handler_OnLButtonUp.md) (HIGH 100/100) - XML Handler
- [OnLButtonUp](../../events/window_events/window_event_OnLButtonUp.md) (HIGH 100/100) - Window Event
- [OnRButtonUp](../../xml/handlers/handler_OnRButtonUp.md) (HIGH 100/100) - XML Handler
- [OnRButtonUp](../../events/window_events/window_event_OnRButtonUp.md) (HIGH 100/100) - Window Event

## Affects

- [Button](../../xml/element_types/element_Button.md) (HIGH 100/100) - XML Element Type
- [EA_Window_WorldMap](../tables/table_EA_Window_WorldMap.md) (HIGH 100/100) - Global Table
- [GameData.Player.isGroupLeader](../../gamedata/fields/gamedata_GameData.Player.isGroupLeader.md) (HIGH 100/100) - GameData Field
- [Label](../../xml/element_types/element_Label.md) (HIGH 100/100) - XML Element Type
- [PartyUtils.GetWarbandLeader](global_PartyUtils.GetWarbandLeader.md) (HIGH 100/100) - Global Function
- [SystemData.MouseOverWindow.name](../../systemdata/fields/systemdata_SystemData.MouseOverWindow.name.md) (HIGH 100/100) - SystemData Field
- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Notes

- Canonical entry built from observed call sites, not from engine source or decompiled definitions.
