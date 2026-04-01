# OnShutdown

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
| Addons seen in | AutoBand, CM_ClosetGoblin, CMap, EA_UiDebugTools, Queue Queuer, bigger_MacroWindow |
| Files seen in | `/workspace/Autoband/AutoBandWindowConfig.xml:21`, `/workspace/Autoband/AutoBandWindowTemplate.xml:22`, `/workspace/Autoband/AutoBandWindowTools.xml:21`, `/workspace/ClosetGoblin/ClosetGoblin.xml:1228`, `/workspace/ClosetGoblin/ClosetGoblin.xml:264`, `/workspace/QueueQueuer/QueueQueuer_GUI.xml:168`, `/workspace/QueueQueuer/QueueQueuer_GUI.xml:40`, `/workspace/QueueQueuer/QueueQueuer_GUI_TabAbout.xml:40` |
| Namespaces detected | OnShutdown |
| Source kinds | bindings, xml_handlers |
| Example locations | AutoBand: AutoBandWindowConfig.OnShutdown, AutoBand: AutoBandWindowTemplate.OnShutdown, AutoBand: AutoBandWindowTools.OnShutdown, CM_ClosetGoblin: ClosetGoblinCharacterWindow.OnShutdown, CM_ClosetGoblin: ClosetGoblinZoneWindow.OnShutdown, CMap: CMapWindow.OnShutdown |
| XML usage count | 17 |
| XML attribute usage count | 17 |
| Lua usage count | 17 |
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

Observed as an XML handler hook bound by 6 addons through frame event handlers.

## Expected Lua Binding

```lua
function(...)
```

## Element Types

- Window

## Seen In

- AutoBand
- CM_ClosetGoblin
- CMap
- EA_UiDebugTools
- Queue Queuer
- bigger_MacroWindow

## Examples

- AutoBand: AutoBandWindowConfig -> AutoBandWindowConfig.OnShutdown -> AutoBandWindowConfig.Shutdown
- AutoBand: AutoBandWindowTemplate -> AutoBandWindowTemplate.OnShutdown -> AutoBandWindowTemplate.Shutdown
- AutoBand: AutoBandWindowTools -> AutoBandWindowTools.OnShutdown -> AutoBandWindowTools.Shutdown
- CM_ClosetGoblin: ClosetGoblinCharacterWindow -> ClosetGoblinCharacterWindow.OnShutdown -> ClosetGoblinCharacterWindow.OnShutdown
- CM_ClosetGoblin: ClosetGoblinZoneWindow -> ClosetGoblinZoneWindow.OnShutdown -> ClosetGoblinZoneWindow.OnShutdown
- CMap: CMapWindow -> CMapWindow.OnShutdown -> CMapWindow.Shutdown

## Related APIs

- none

## Used With

- [OnShutdown](../../events/window_events/window_event_OnShutdown.md) (HIGH 100/100) - Window Event
- [Window](../element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Triggered By

- none

## Affects

- none

## Notes

- Expected binding arguments remain uncertain because API_Ref captures symbol linkage, not full handler signatures.
