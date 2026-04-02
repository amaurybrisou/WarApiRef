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
| Addons seen in | CM_ClosetGoblin, CMap, EA_UiDebugTools, bigger_MacroWindow |
| Files seen in | `/workspace_addons/ClosetGoblin/ClosetGoblin.xml:1228`, `/workspace_addons/ClosetGoblin/ClosetGoblin.xml:264`, `/workspace_addons/bigger_macrowindow/Source/MacroWindow.xml:358`, `/workspace_addons/cmap/CMap.xml:40`, `/workspace_addons/cmap/CMap.xml:430`, `/workspace_addons/ea_uidebugtools/Source/DebugWindow.xml:30` |
| Namespaces detected | OnShutdown |
| Source kinds | bindings, xml_handlers |
| Example locations | CM_ClosetGoblin: ClosetGoblinCharacterWindow.OnShutdown, CM_ClosetGoblin: ClosetGoblinZoneWindow.OnShutdown, CMap: CMapWindow.OnShutdown, CMap: CMapWindowPinFilterMenu.OnShutdown, EA_UiDebugTools: DebugWindow.OnShutdown, bigger_MacroWindow: EA_Window_Macro.OnShutdown |
| XML usage count | 6 |
| XML attribute usage count | 6 |
| Lua usage count | 6 |
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

- Window

## Seen In

- CM_ClosetGoblin
- CMap
- EA_UiDebugTools
- bigger_MacroWindow

## Examples

- CM_ClosetGoblin: ClosetGoblinCharacterWindow -> ClosetGoblinCharacterWindow.OnShutdown -> ClosetGoblinCharacterWindow.OnShutdown
- CM_ClosetGoblin: ClosetGoblinZoneWindow -> ClosetGoblinZoneWindow.OnShutdown -> ClosetGoblinZoneWindow.OnShutdown
- CMap: CMapWindow -> CMapWindow.OnShutdown -> CMapWindow.Shutdown
- CMap: CMapWindowPinFilterMenu -> CMapWindowPinFilterMenu.OnShutdown -> EA_Window_ContextMenu.Shutdown
- EA_UiDebugTools: DebugWindow -> DebugWindow.OnShutdown -> DebugWindow.Shutdown
- bigger_MacroWindow: EA_Window_Macro -> EA_Window_Macro.OnShutdown -> EA_Window_Macro.Shutdown

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

- Expected binding arguments remain uncertain because addon-api docs capture symbol linkage, not full handler signatures.
