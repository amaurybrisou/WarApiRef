# OnSetMoving

- Category: Window Event
- Confidence level: HIGH
- Confidence score: 73/100

## Confidence Assessment

- Level: HIGH

- Score: 73/100

- Rationale: Promoted as HIGH confidence because used directly in xml handler attributes, referenced by generated docs or reference files, used in event registration or dispatch.

## Evidence Signals

- +30 Used directly in XML handler attributes: XML exposure suggests an engine-level contract.
- +18 Used in event registration or dispatch: Observed in event-driven engine hooks.
- +25 Referenced by generated docs or reference files: The symbol is reinforced outside a single call page.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | MapPin |
| Files seen in | `/workspace_addons/MapPin/source/MapPin.xml:343`, `/workspace_addons/MapPin/source/MapPin.xml:407` |
| Namespaces detected | OnSetMoving |
| Source kinds | event_page, xml_handlers |
| Example locations | MapPin: MapPinWBMarker.OnSetMoving, MapPin: MapPinWBMarker_NoBG.OnSetMoving |
| XML usage count | 2 |
| XML attribute usage count | 2 |
| Lua usage count | 0 |
| Global usage count | 0 |
| Local definition count | 0 |
| Documentation references | 1 |
| Initialization flow references | 0 |
| Known engine namespace | no |
| Default UI presence | no |
| Event binding presence | yes |
| Observed in XML and Lua | no |
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

Observed as an engine-supplied UI event hook used by 1 addons.

## Handler Pattern

Observed as an On* callback routed into a module-qualified Lua function.

## Payload

- Window callback arguments are not fully inferable from addon-api docs alone.

## Seen In

- MapPin

## Registrars And Handlers

- MapPin.OnMoving

## Examples

- MapPin: MapPinWBMarker -> MapPinWBMarker.OnSetMoving -> MapPin.OnMoving
- MapPin: MapPinWBMarker_NoBG -> MapPinWBMarker_NoBG.OnSetMoving -> MapPin.OnMoving

## Related APIs

- none

## Used With

- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type
- [OnSetMoving](../../xml/handlers/handler_OnSetMoving.md) (HIGH 93/100) - XML Handler

## Triggered By

- none

## Affects

- none

## Notes

- Only one addon surfaced this event in the current addon-api corpus.
