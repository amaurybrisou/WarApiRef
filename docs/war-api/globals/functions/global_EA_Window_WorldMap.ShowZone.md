# EA_Window_WorldMap.ShowZone

- Category: Global Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 1 addons

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 115

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, matches a known engine namespace, referenced by generated docs or reference files.

## Evidence Signals

- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +10 Argument pattern is consistent: Observed argument positions remain stable.
- +25 Referenced by generated docs or reference files: The symbol is reinforced outside a single call page.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | MapPin |
| Files seen in | `/workspace_addons/MapPin/source/MapPin.lua:2280`, `/workspace_addons/MapPin/source/MapPin.lua:4227` |
| Namespaces detected | EA_Window_WorldMap |
| Source kinds | globals, lua_calls |
| Example locations | MapPin: MapPin.LButtonUp, MapPin: MapPin.QuickWP |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 2 |
| Global usage count | 2 |
| Local definition count | 0 |
| Documentation references | 1 |
| Initialization flow references | 0 |
| Known engine namespace | yes |
| Default UI presence | yes |
| Event binding presence | no |
| Observed in XML and Lua | no |
| Consistent role | no |
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
EA_Window_WorldMap.ShowZone(arg1)
```

## Description

Observed as a global function across 1 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a function or method reference. | Observed values: MapPin.Pins.WayPoint[ButtonNumber-9].Zone, MapPin.Pins.WayPoint[ClickedButton].Zone |

## Returns

- Not confidently inferable from addon-api docs alone.

## Side Effects

- No side effect is confidently inferable from addon-api docs alone.

## Seen In

- MapPin

## Examples

- MapPin: MapPin.LButtonUp -> EA_Window_WorldMap.ShowZone(MapPin.Pins.WayPoint[ButtonNumber-9].Zone)
- MapPin: MapPin.QuickWP -> EA_Window_WorldMap.ShowZone(MapPin.Pins.WayPoint[ClickedButton].Zone)

## Related APIs

- none

## Used With

- [DialogManager.MakeTextEntryDialog](global_DialogManager.MakeTextEntryDialog.md) (HIGH 100/100) - Global Function
- [EA_ChatWindow.InsertText](global_EA_ChatWindow.InsertText.md) (HIGH 100/100) - Global Function
- [EA_Window_WorldMap](../tables/table_EA_Window_WorldMap.md) (HIGH 100/100) - Global Table
- [OnLButtonUp](../../events/window_events/window_event_OnLButtonUp.md) (HIGH 100/100) - Window Event

## Triggered By

- [OnLButtonUp](../../xml/handlers/handler_OnLButtonUp.md) (HIGH 100/100) - XML Handler
- [OnLButtonUp](../../events/window_events/window_event_OnLButtonUp.md) (HIGH 100/100) - Window Event

## Affects

- [EA_Window_WorldMap](../tables/table_EA_Window_WorldMap.md) (HIGH 100/100) - Global Table
- [PartyUtils.GetWarbandLeader](global_PartyUtils.GetWarbandLeader.md) (HIGH 100/100) - Global Function
- [SystemData.MouseOverWindow.name](../../systemdata/fields/systemdata_SystemData.MouseOverWindow.name.md) (HIGH 100/100) - SystemData Field
- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Notes

- Only one addon surfaced this symbol in the current corpus.
- Canonical entry built from observed call sites, not from engine source or decompiled definitions.
