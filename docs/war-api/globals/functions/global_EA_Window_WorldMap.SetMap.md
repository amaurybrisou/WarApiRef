# EA_Window_WorldMap.SetMap

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
| Addons seen in | MapMonster |
| Files seen in | `/workspace/MapMonster/Source/MapMonster_EditorWindow.lua:822`, `/workspace/MapMonster/Source/MapMonster_Pins.lua:2492` |
| Namespaces detected | EA_Window_WorldMap |
| Source kinds | globals, lua_calls |
| Example locations | MapMonster: MapMonster.Editor.OnZoneNameChange, MapMonster: MapMonsterAPI.HighlightMapPin |
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
EA_Window_WorldMap.SetMap(arg1, arg2)
```

## Description

Observed as a global function across 1 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a function or method reference. | Observed values: GameDefs.MapLevel.ZONE_MAP |
| arg2 | Observed as a function or method reference. | Observed values: pinData.zoneId, zoneComboOrder[selectedZone] |

## Returns

- Not confidently inferable from addon-api docs alone.

## Side Effects

- Mutates engine-visible UI state or presentation.

## Seen In

- MapMonster

## Examples

- MapMonster: MapMonster.Editor.OnZoneNameChange -> EA_Window_WorldMap.SetMap(GameDefs.MapLevel.ZONE_MAP, zoneComboOrder[selectedZone])
- MapMonster: MapMonsterAPI.HighlightMapPin -> EA_Window_WorldMap.SetMap(GameDefs.MapLevel.ZONE_MAP, pinData.zoneId)

## Related APIs

- none

## Used With

- none

## Triggered By

- [OnSelChanged](../../xml/handlers/handler_OnSelChanged.md) (HIGH 100/100) - XML Handler
- [OnSelChanged](../../events/window_events/window_event_OnSelChanged.md) (HIGH 100/100) - Window Event

## Affects

- [ComboBox](../../xml/element_types/element_ComboBox.md) (HIGH 100/100) - XML Element Type

## Notes

- Only one addon surfaced this symbol in the current corpus.
- Canonical entry built from observed call sites, not from engine source or decompiled definitions.
