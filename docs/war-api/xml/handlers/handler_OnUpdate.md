# OnUpdate

- Type: XML Handler
- Confidence level: HIGH
- Confidence score: 100/100

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 138

- Rationale: Promoted as HIGH confidence because seen in 4 or more addons, used directly in xml handler attributes, referenced by generated docs or reference files.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +30 Used directly in XML handler attributes: XML exposure suggests an engine-level contract.
- +18 Used in event registration or dispatch: Observed in event-driven engine hooks.
- +20 Observed in both XML and Lua paths: Cross-source linkage reinforces platform-level usage.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +25 Referenced by generated docs or reference files: The symbol is reinforced outside a single call page.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | BuffHead, CMap, EA_UiDebugTools, TidyRoll |
| Files seen in | `/workspace_addons/BuffHead/Setup/SetupLayout.xml:249`, `/workspace_addons/TidyRoll/TidyRoll.xml:275`, `/workspace_addons/cmap/CMap.xml:104`, `/workspace_addons/ea_uidebugtools/Source/DebugWindow.xml:29` |
| Namespaces detected | OnUpdate |
| Source kinds | bindings, xml_handlers |
| Example locations | BuffHead: BuffHeadSetupLayoutWindow.OnUpdate, CMap: CMapWindowWMap.OnUpdate, EA_UiDebugTools: DebugWindow.OnUpdate, TidyRoll: TidyRollTimer.OnUpdate |
| XML usage count | 4 |
| XML attribute usage count | 4 |
| Lua usage count | 4 |
| Global usage count | 0 |
| Local definition count | 0 |
| Documentation references | 1 |
| Initialization flow references | 0 |
| Known engine namespace | no |
| Default UI presence | no |
| Event binding presence | yes |
| Observed in XML and Lua | yes |
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

Observed as an XML handler hook bound by 4 addons through frame event handlers.

## Expected Lua Binding

```lua
function(...)
```

## Element Types

- MapDisplay
- Window

## Seen In

- BuffHead
- CMap
- EA_UiDebugTools
- TidyRoll

## Examples

- BuffHead: BuffHeadSetupLayoutWindow -> BuffHeadSetupLayoutWindow.OnUpdate -> BuffHead.Setup.Layout.OnUpdate
- CMap: CMapWindowWMap -> CMapWindowWMap.OnUpdate -> CMapWindow.UpdateCoordinatesWMap
- EA_UiDebugTools: DebugWindow -> DebugWindow.OnUpdate -> DebugWindow.Update
- TidyRoll: TidyRollTimer -> TidyRollTimer.OnUpdate -> TidyRoll.OnUpdate

## Related APIs

- [WindowSetShowing](../../window_api/functions/window_WindowSetShowing.md) (HIGH 100/100) - Window Function

## Used With

- [OnUpdate](../../events/window_events/window_event_OnUpdate.md) (HIGH 100/100) - Window Event

## Triggered By

- none

## Affects

- none

## Notes

- Expected binding arguments remain uncertain because addon-api docs capture symbol linkage, not full handler signatures.
