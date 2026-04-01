# OnScrollPosChanged

- Category: Window Event
- Confidence level: MEDIUM
- Confidence score: 53/100

## Confidence Assessment

- Level: MEDIUM

- Score: 53/100

- Rationale: Promoted as MEDIUM confidence because used directly in xml handler attributes, referenced by generated docs or reference files, used in event registration or dispatch.

## Evidence Signals

- +30 Used directly in XML handler attributes: XML exposure suggests an engine-level contract.
- +18 Used in event registration or dispatch: Observed in event-driven engine hooks.
- +25 Referenced by generated docs or reference files: The symbol is reinforced outside a single call page.
- -20 Only one weak usage site: Evidence is too shallow to trust as platform API.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | bigger_MacroWindow |
| Files seen in | `/workspace/bigger_macrowindow/Source/MacroIcons.xml:9` |
| Namespaces detected | OnScrollPosChanged |
| Source kinds | event_page, xml_handlers |
| Example locations | bigger_MacroWindow: MacroIconsScrollBar.OnScrollPosChanged |
| XML usage count | 1 |
| XML attribute usage count | 1 |
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
| Weak usage only | yes |
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

- bigger_MacroWindow

## Registrars And Handlers

- MacroIcons.ScrollPos

## Examples

- bigger_MacroWindow: MacroIconsScrollBar -> MacroIconsScrollBar.OnScrollPosChanged -> MacroIcons.ScrollPos

## Related APIs

- none

## Used With

- [VerticalScrollbar](../../xml/element_types/element_VerticalScrollbar.md) (HIGH 100/100) - XML Element Type
- [OnScrollPosChanged](../../xml/handlers/handler_OnScrollPosChanged.md) (HIGH 73/100) - XML Handler

## Triggered By

- none

## Affects

- none

## Notes

- Only one addon surfaced this event in the current addon-api corpus.
