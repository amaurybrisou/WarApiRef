# OnPointMouseOver

- Type: XML Handler
- Confidence level: HIGH
- Confidence score: 88/100

## Confidence Assessment

- Level: HIGH

- Score: 88/100

- Rationale: Promoted as HIGH confidence because seen in 4 or more addons, referenced by generated docs or reference files, used in event registration or dispatch.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +18 Used in event registration or dispatch: Observed in event-driven engine hooks.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +25 Referenced by generated docs or reference files: The symbol is reinforced outside a single call page.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | Atlas, CMap, Map, Minmap |
| Namespaces detected | OnPointMouseOver |
| Source kinds | bindings, xml_handlers |
| Example locations | Atlas: .OnPointMouseOver, CMap: .OnPointMouseOver, Map: .OnPointMouseOver, Minmap: .OnPointMouseOver |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 5 |
| Global usage count | 0 |
| Local definition count | 0 |
| Documentation references | 1 |
| Initialization flow references | 0 |
| Known engine namespace | no |
| Default UI presence | no |
| Event binding presence | yes |
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

XML handler event observed across 4 addons.

## Expected Lua Binding

```lua
function(...)
```

## Element Types

- none

## Seen In

- Atlas
- CMap
- Map
- Minmap

## Examples

- Atlas: .OnPointMouseOver -> AtlasMap.OnMouseOverPoint
- CMap: .OnPointMouseOver -> CMapWindow.OnMouseOverPoint
- CMap: .OnPointMouseOver -> CMapWindow.WmapOnMouseOverPoint
- Map: .OnPointMouseOver -> Map.OnMouseOverPoint
- Minmap: .OnPointMouseOver -> Minmap.OnMouseOverPoint

## Related APIs

- [MapDisplay](../element_types/element_MapDisplay.md) (HIGH 100/100) - XML Element Type

## Notes

- Expected binding arguments remain uncertain because contract artifacts capture symbol linkage, not full handler signatures.
