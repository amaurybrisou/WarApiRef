# OnUpdate

- Category: Window Event
- Confidence level: HIGH
- Confidence score: 100/100

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 188

- Rationale: Promoted as HIGH confidence because seen in 4 or more addons, used directly in xml handler attributes, referenced by generated docs or reference files.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +30 Used directly in XML handler attributes: XML exposure suggests an engine-level contract.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +18 Used in event registration or dispatch: Observed in event-driven engine hooks.
- +10 Referenced from initialization flow: Lifecycle reconstruction references this symbol.
- +20 Observed in both XML and Lua paths: Cross-source linkage reinforces platform-level usage.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +25 Referenced by generated docs or reference files: The symbol is reinforced outside a single call page.
- +20 Reinforced across multiple generated source types: Evidence comes from several independent API_Ref source types.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | BuffHead, CMap, EA_UiDebugTools, Effigy, Queue Queuer, TidyQueue, TidyRoll |
| Files seen in | `/workspace/BuffHead/Setup/SetupLayout.xml:249`, `/workspace/Effigy/States/EffigyStateCastbar.lua:41`, `/workspace/Effigy/States/EffigyStatePlayer.lua:45`, `/workspace/QueueQueuer/QueueQueuer_GUI.xml:102`, `/workspace/QueueQueuer/QueueQueuer_GUI.xml:125`, `/workspace/QueueQueuer/QueueQueuer_GUI.xml:167`, `/workspace/QueueQueuer/QueueQueuer_GUI.xml:56`, `/workspace/QueueQueuer/QueueQueuer_GUI.xml:79` |
| Namespaces detected | OnUpdate |
| Source kinds | event_page, flows, lua_event_registration, xml_handlers |
| Example locations | BuffHead: BuffHeadSetupLayoutWindow.OnUpdate, CMap: CMapWindowWMap.OnUpdate, EA_UiDebugTools: DebugWindow.OnUpdate, Effigy: Effigy.RegisterStateInfoForCastbar, Effigy: Effigy.RegisterStateInfoForPlayer, Queue Queuer: QueueQueuer_GUI.OnUpdate |
| XML usage count | 10 |
| XML attribute usage count | 10 |
| Lua usage count | 2 |
| Global usage count | 2 |
| Local definition count | 0 |
| Documentation references | 3 |
| Initialization flow references | 21 |
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

Observed as an engine-supplied UI event hook used by 7 addons.

## Handler Pattern

Observed as an On* callback routed into a module-qualified Lua function.

## Payload

- Window callback arguments are not fully inferable from API_Ref alone.

## Seen In

- BuffHead
- CMap
- EA_UiDebugTools
- Effigy
- Queue Queuer
- TidyQueue
- TidyRoll

## Registrars And Handlers

- BuffHead.Setup.Layout.OnUpdate
- CMapWindow.UpdateCoordinatesWMap
- DebugWindow.Update
- Effigy.Name..".CastbarUpdate"
- Effigy.Name..".SwingTimerUpdate"
- QueueQueuer_GUI.MapButton_OnUpdate
- QueueQueuer_GUI.OnUpdate
- TidyQueue.OnUpdate
- TidyRoll.OnUpdate
- WindowRegisterCoreEventHandler
- core

## Examples

- Effigy: Effigy.RegisterStateInfoForCastbar -> OnUpdate -> Effigy.Name..".CastbarUpdate"
- Effigy: Effigy.RegisterStateInfoForPlayer -> OnUpdate -> Effigy.Name..".SwingTimerUpdate"
- BuffHead: BuffHeadSetupLayoutWindow -> BuffHeadSetupLayoutWindow.OnUpdate -> BuffHead.Setup.Layout.OnUpdate
- CMap: CMapWindowWMap -> CMapWindowWMap.OnUpdate -> CMapWindow.UpdateCoordinatesWMap
- EA_UiDebugTools: DebugWindow -> DebugWindow.OnUpdate -> DebugWindow.Update
- Queue Queuer: QueueQueuer_GUI -> QueueQueuer_GUI.OnUpdate -> QueueQueuer_GUI.OnUpdate

## Related APIs

- [EA_Window_ScenarioLobby.OnJoinInstanceWait](../../globals/functions/global_EA_Window_ScenarioLobby.OnJoinInstanceWait.md) (HIGH 100/100) - Global Function

## Used With

- [Button](../../xml/element_types/element_Button.md) (HIGH 100/100) - XML Element Type
- [MapDisplay](../../xml/element_types/element_MapDisplay.md) (HIGH 100/100) - XML Element Type
- [OnUpdate](../../xml/handlers/handler_OnUpdate.md) (HIGH 100/100) - XML Handler
- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Triggered By

- none

## Affects

- none

## Notes

- none
