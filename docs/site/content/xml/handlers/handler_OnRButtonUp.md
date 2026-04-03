# OnRButtonUp

- Type: XML Handler
- Confidence level: HIGH
- Confidence score: 73/100

## Confidence Assessment

- Level: HIGH

- Score: 73/100

- Rationale: Promoted as HIGH confidence because used directly in xml handler attributes, referenced by generated docs or reference files, observed in both xml and lua paths.

## Evidence Signals

- +30 Used directly in XML handler attributes: XML exposure suggests an engine-level contract.
- +18 Used in event registration or dispatch: Observed in event-driven engine hooks.
- +20 Observed in both XML and Lua paths: Cross-source linkage reinforces platform-level usage.
- +25 Referenced by generated docs or reference files: The symbol is reinforced outside a single call page.
- -20 Only one weak usage site: Evidence is too shallow to trust as platform API.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | TidyRoll |
| Files seen in | `/workspace/data/raw/TidyRoll/TidyRoll.xml:201` |
| Namespaces detected | OnRButtonUp |
| Source kinds | bindings, xml_handlers |
| Example locations | TidyRoll: TidyRollFrame.OnRButtonUp |
| XML usage count | 1 |
| XML attribute usage count | 1 |
| Lua usage count | 1 |
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
| Weak usage only | yes |
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

- TidyRoll

## Examples

- TidyRoll: TidyRollFrame -> TidyRollFrame.OnRButtonUp -> FrameManager.OnRButtonUp

## Related APIs

- [Window](../element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Used With

- none

## Triggered By

- none

## Affects

- none

## Notes

- Expected binding arguments remain uncertain because addon-api docs capture symbol linkage, not full handler signatures.
