# OnShutdown

- Category: Window Event
- Confidence level: HIGH
- Confidence score: 100/100

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 148

- Rationale: Promoted as HIGH confidence because seen in 4 or more addons, used directly in xml handler attributes, referenced by generated docs or reference files.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +30 Used directly in XML handler attributes: XML exposure suggests an engine-level contract.
- +18 Used in event registration or dispatch: Observed in event-driven engine hooks.
- +10 Referenced from initialization flow: Lifecycle reconstruction references this symbol.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +25 Referenced by generated docs or reference files: The symbol is reinforced outside a single call page.
- +20 Reinforced across multiple generated source types: Evidence comes from several independent API_Ref source types.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | AutoBand, CM_ClosetGoblin, CMap, EA_UiDebugTools, Queue Queuer, bigger_MacroWindow |
| Files seen in | `/workspace/Autoband/AutoBandWindowConfig.xml:21`, `/workspace/Autoband/AutoBandWindowTemplate.xml:22`, `/workspace/Autoband/AutoBandWindowTools.xml:21`, `/workspace/ClosetGoblin/ClosetGoblin.xml:1228`, `/workspace/ClosetGoblin/ClosetGoblin.xml:264`, `/workspace/QueueQueuer/QueueQueuer_GUI.xml:168`, `/workspace/QueueQueuer/QueueQueuer_GUI.xml:40`, `/workspace/QueueQueuer/QueueQueuer_GUI_TabAbout.xml:40` |
| Namespaces detected | OnShutdown |
| Source kinds | event_page, flows, xml_handlers |
| Example locations | AutoBand: AutoBandWindowConfig.OnShutdown, AutoBand: AutoBandWindowTemplate.OnShutdown, AutoBand: AutoBandWindowTools.OnShutdown, CM_ClosetGoblin: ClosetGoblinCharacterWindow.OnShutdown, CM_ClosetGoblin: ClosetGoblinZoneWindow.OnShutdown, CMap: CMapWindow.OnShutdown |
| XML usage count | 17 |
| XML attribute usage count | 17 |
| Lua usage count | 0 |
| Global usage count | 0 |
| Local definition count | 0 |
| Documentation references | 2 |
| Initialization flow references | 16 |
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

Observed as an engine-supplied UI event hook used by 6 addons.

## Handler Pattern

Observed as an On* callback routed into a module-qualified Lua function.

## Payload

- Window callback arguments are not fully inferable from API_Ref alone.

## Seen In

- AutoBand
- CM_ClosetGoblin
- CMap
- EA_UiDebugTools
- Queue Queuer
- bigger_MacroWindow

## Registrars And Handlers

- AutoBandWindowConfig.Shutdown
- AutoBandWindowTemplate.Shutdown
- AutoBandWindowTools.Shutdown
- CMapWindow.Shutdown
- ClosetGoblinCharacterWindow.OnShutdown
- ClosetGoblinZoneWindow.OnShutdown
- DebugWindow.Shutdown
- EA_Window_ContextMenu.Shutdown
- EA_Window_Macro.Shutdown
- QueueQueuer_GUI.OnShutdown
- QueueQueuer_GUI_MapButtons.OnShutdown
- QueueQueuer_GUI_TabAbout.OnShutdown
- QueueQueuer_GUI_TabHelp.OnShutdown
- QueueQueuer_GUI_TabTier1.OnShutdown
- QueueQueuer_GUI_TabTier2.OnShutdown
- QueueQueuer_GUI_TabTier3.OnShutdown
- QueueQueuer_GUI_TabTier4.OnShutdown

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

- [OnShutdown](../../xml/handlers/handler_OnShutdown.md) (HIGH 100/100) - XML Handler
- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Triggered By

- none

## Affects

- none

## Notes

- none
