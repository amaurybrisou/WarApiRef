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
| Addons seen in | BuffHead, CMap, EA_UiDebugTools, Queue Queuer, TidyQueue, TidyRoll |
| Files seen in | `/workspace/BuffHead/Setup/SetupLayout.xml:249`, `/workspace/QueueQueuer/QueueQueuer_GUI.xml:102`, `/workspace/QueueQueuer/QueueQueuer_GUI.xml:125`, `/workspace/QueueQueuer/QueueQueuer_GUI.xml:167`, `/workspace/QueueQueuer/QueueQueuer_GUI.xml:56`, `/workspace/QueueQueuer/QueueQueuer_GUI.xml:79`, `/workspace/TidyQueue/TidyQueue.xml:126`, `/workspace/TidyRoll/TidyRoll.xml:275` |
| Namespaces detected | OnUpdate |
| Source kinds | bindings, xml_handlers |
| Example locations | BuffHead: BuffHeadSetupLayoutWindow.OnUpdate, CMap: CMapWindowWMap.OnUpdate, EA_UiDebugTools: DebugWindow.OnUpdate, Queue Queuer: QueueQueuer_GUI.OnUpdate, Queue Queuer: QueueQueuer_GUI_MapButton.OnUpdate, Queue Queuer: QueueQueuer_GUI_MapButton_COOLDOWN.OnUpdate |
| XML usage count | 10 |
| XML attribute usage count | 10 |
| Lua usage count | 10 |
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

- Button
- MapDisplay
- Window

## Seen In

- BuffHead
- CMap
- EA_UiDebugTools
- Queue Queuer
- TidyQueue
- TidyRoll

## Examples

- BuffHead: BuffHeadSetupLayoutWindow -> BuffHeadSetupLayoutWindow.OnUpdate -> BuffHead.Setup.Layout.OnUpdate
- CMap: CMapWindowWMap -> CMapWindowWMap.OnUpdate -> CMapWindow.UpdateCoordinatesWMap
- EA_UiDebugTools: DebugWindow -> DebugWindow.OnUpdate -> DebugWindow.Update
- Queue Queuer: QueueQueuer_GUI -> QueueQueuer_GUI.OnUpdate -> QueueQueuer_GUI.OnUpdate
- Queue Queuer: QueueQueuer_GUI_MapButton -> QueueQueuer_GUI_MapButton.OnUpdate -> QueueQueuer_GUI.MapButton_OnUpdate
- Queue Queuer: QueueQueuer_GUI_MapButton_COOLDOWN -> QueueQueuer_GUI_MapButton_COOLDOWN.OnUpdate -> QueueQueuer_GUI.MapButton_OnUpdate

## Related APIs

- [EA_Window_ScenarioLobby.OnJoinInstanceWait](../../globals/functions/global_EA_Window_ScenarioLobby.OnJoinInstanceWait.md) (HIGH 100/100) - Global Function
- [WindowSetShowing](../../window_api/functions/window_WindowSetShowing.md) (HIGH 100/100) - Window Function

## Used With

- [Button](../element_types/element_Button.md) (HIGH 100/100) - XML Element Type
- [MapDisplay](../element_types/element_MapDisplay.md) (HIGH 100/100) - XML Element Type
- [OnUpdate](../../events/window_events/window_event_OnUpdate.md) (HIGH 100/100) - Window Event
- [Window](../element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Triggered By

- none

## Affects

- none

## Notes

- Expected binding arguments remain uncertain because API_Ref captures symbol linkage, not full handler signatures.
