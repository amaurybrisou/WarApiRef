# OnMouseOut

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
| Addons seen in | Killer |
| Files seen in | `/workspace/data/raw/Killer/Killer.xml:0` |
| Namespaces detected | OnMouseOut |
| Source kinds | event_page, xml_handlers |
| Example locations | Killer: KillerFeedRowTemplate.OnMouseOut, Killer: KillerPersonalCounter.OnMouseOut, Killer: KillerWindowAllTimeHover.OnMouseOut, Killer: KillerWindowRecentHover.OnMouseOut |
| XML usage count | 4 |
| XML attribute usage count | 4 |
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

- Killer

## Registrars And Handlers

- Killer.HideRowTooltip

## Examples

- Killer: KillerFeedRowTemplate -> KillerFeedRowTemplate.OnMouseOut -> Killer.HideRowTooltip
- Killer: KillerPersonalCounter -> KillerPersonalCounter.OnMouseOut -> Killer.HideRowTooltip
- Killer: KillerWindowAllTimeHover -> KillerWindowAllTimeHover.OnMouseOut -> Killer.HideRowTooltip
- Killer: KillerWindowRecentHover -> KillerWindowRecentHover.OnMouseOut -> Killer.HideRowTooltip

## Related APIs

- none

## Used With

- none

## Triggered By

- none

## Affects

- none

## Notes

- Only one addon surfaced this event in the current addon-api corpus.
