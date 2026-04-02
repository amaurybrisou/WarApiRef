# EA_Window_WorldMap

- Category: Global Table
- Confidence level: HIGH
- Confidence score: 100/100

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 118

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, matches a known engine namespace, referenced by generated docs or reference files.

## Evidence Signals

- +18 Seen in 2 to 3 addons: Cross-addon spread is present but limited.
- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +25 Referenced by generated docs or reference files: The symbol is reinforced outside a single call page.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | MapMonster, MapPin |
| Files seen in | `/workspace_addons/MapMonster/Source/MapMonster_EditorWindow.lua:822`, `/workspace_addons/MapMonster/Source/MapMonster_Pins.lua:2492`, `/workspace_addons/MapPin/source/MapPin.lua:2280`, `/workspace_addons/MapPin/source/MapPin.lua:4227` |
| Namespaces detected | EA_Window_WorldMap |
| Source kinds | globals, lua_calls |
| Example locations | MapMonster: MapMonster.Editor.OnZoneNameChange, MapMonster: MapMonsterAPI.HighlightMapPin, MapPin: MapPin.LButtonUp, MapPin: MapPin.QuickWP |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 4 |
| Global usage count | 2 |
| Local definition count | 2 |
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

Observed shared global table or namespace surfaced in 2 addons.

## Functions

- EA_Window_WorldMap.SetMap
- EA_Window_WorldMap.ShowZone

## Observed Members

- none

## Seen In

- MapMonster
- MapPin

## Examples

- MapMonster: MapMonster.Editor.OnZoneNameChange -> EA_Window_WorldMap.SetMap(GameDefs.MapLevel.ZONE_MAP, zoneComboOrder[selectedZone])
- MapMonster: MapMonsterAPI.HighlightMapPin -> EA_Window_WorldMap.SetMap(GameDefs.MapLevel.ZONE_MAP, pinData.zoneId)
- MapPin: MapPin.LButtonUp -> EA_Window_WorldMap.ShowZone(MapPin.Pins.WayPoint[ButtonNumber-9].Zone)
- MapPin: MapPin.QuickWP -> EA_Window_WorldMap.ShowZone(MapPin.Pins.WayPoint[ClickedButton].Zone)

## Related APIs

- [EA_ChatWindow.InsertText](../functions/global_EA_ChatWindow.InsertText.md) (HIGH 100/100) - Global Function
- [EA_Window_WorldMap.ShowZone](../functions/global_EA_Window_WorldMap.ShowZone.md) (HIGH 100/100) - Global Function
- [PartyUtils.GetWarbandLeader](../functions/global_PartyUtils.GetWarbandLeader.md) (HIGH 100/100) - Global Function
- [PartyUtils.IsPartyActive](../functions/global_PartyUtils.IsPartyActive.md) (HIGH 100/100) - Global Function

## Used With

- [DialogManager.MakeTextEntryDialog](../functions/global_DialogManager.MakeTextEntryDialog.md) (HIGH 100/100) - Global Function
- [EA_ChatWindow.InsertText](../functions/global_EA_ChatWindow.InsertText.md) (HIGH 100/100) - Global Function
- [EA_Window_WorldMap.ShowZone](../functions/global_EA_Window_WorldMap.ShowZone.md) (HIGH 100/100) - Global Function
- [OnLButtonUp](../../events/window_events/window_event_OnLButtonUp.md) (HIGH 100/100) - Window Event

## Triggered By

- none

## Affects

- none

## Notes

- none
