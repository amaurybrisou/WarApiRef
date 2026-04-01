# OnSetMoving

- Type: XML Handler
- Confidence level: HIGH
- Confidence score: 93/100

## Confidence Assessment

- Level: HIGH

- Score: 93/100

- Rationale: Promoted as HIGH confidence because used directly in xml handler attributes, referenced by generated docs or reference files, observed in both xml and lua paths.

## Evidence Signals

- +30 Used directly in XML handler attributes: XML exposure suggests an engine-level contract.
- +18 Used in event registration or dispatch: Observed in event-driven engine hooks.
- +20 Observed in both XML and Lua paths: Cross-source linkage reinforces platform-level usage.
- +25 Referenced by generated docs or reference files: The symbol is reinforced outside a single call page.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | MapPin |
| Files seen in | `/workspace/MapPin/source/MapPin.xml:343`, `/workspace/MapPin/source/MapPin.xml:407` |
| Namespaces detected | OnSetMoving |
| Source kinds | bindings, xml_handlers |
| Example locations | MapPin: MapPinWBMarker.OnSetMoving, MapPin: MapPinWBMarker_NoBG.OnSetMoving |
| XML usage count | 2 |
| XML attribute usage count | 2 |
| Lua usage count | 2 |
| Global usage count | 0 |
| Local definition count | 0 |
| Documentation references | 1 |
| Initialization flow references | 0 |
| Known engine namespace | no |
| Default UI presence | no |
| Event binding presence | yes |
| Observed in XML and Lua | yes |
| Consistent role | no |
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

Observed as an XML handler hook bound by 1 addons through frame event handlers.

## Expected Lua Binding

```lua
function(...)
```

## Element Types

- Window

## Seen In

- MapPin

## Examples

- MapPin: MapPinWBMarker -> MapPinWBMarker.OnSetMoving -> MapPin.OnMoving
- MapPin: MapPinWBMarker_NoBG -> MapPinWBMarker_NoBG.OnSetMoving -> MapPin.OnMoving

## Related APIs

- none

## Used With

- [Window](../element_types/element_Window.md) (HIGH 100/100) - XML Element Type
- [OnSetMoving](../../events/window_events/window_event_OnSetMoving.md) (HIGH 73/100) - Window Event

## Triggered By

- none

## Affects

- none

## Notes

- Expected binding arguments remain uncertain because addon-api docs capture symbol linkage, not full handler signatures.
