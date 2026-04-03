# OnShutdown

- Category: Window Event
- Confidence level: HIGH
- Confidence score: 100/100

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 136

- Rationale: Promoted as HIGH confidence because used directly in xml handler attributes, referenced by generated docs or reference files, reinforced across multiple generated source types.

## Evidence Signals

- +18 Seen in 2 to 3 addons: Cross-addon spread is present but limited.
- +30 Used directly in XML handler attributes: XML exposure suggests an engine-level contract.
- +18 Used in event registration or dispatch: Observed in event-driven engine hooks.
- +10 Referenced from initialization flow: Lifecycle reconstruction references this symbol.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +25 Referenced by generated docs or reference files: The symbol is reinforced outside a single call page.
- +20 Reinforced across multiple generated source types: Evidence comes from several independent addon-api source types.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | CM_ClosetGoblin, bigger_MacroWindow |
| Files seen in | `/workspace/data/raw/ClosetGoblin/ClosetGoblin.xml:0`, `/workspace/data/raw/bigger_macrowindow/Source/MacroWindow.xml:0` |
| Namespaces detected | OnShutdown |
| Source kinds | event_page, flows, xml_handlers |
| Example locations | CM_ClosetGoblin: ClosetGoblinCharacterWindow.OnShutdown, CM_ClosetGoblin: ClosetGoblinZoneWindow.OnShutdown, bigger_MacroWindow: EA_Window_Macro.OnShutdown |
| XML usage count | 3 |
| XML attribute usage count | 3 |
| Lua usage count | 0 |
| Global usage count | 0 |
| Local definition count | 0 |
| Documentation references | 2 |
| Initialization flow references | 8 |
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

Observed as an engine-supplied UI event hook used by 2 addons.

## Handler Pattern

Observed as an On* callback routed into a module-qualified Lua function.

## Payload

- Window callback arguments are not fully inferable from addon-api docs alone.

## Seen In

- CM_ClosetGoblin
- bigger_MacroWindow

## Registrars And Handlers

- ClosetGoblinCharacterWindow.OnShutdown
- ClosetGoblinZoneWindow.OnShutdown
- EA_Window_Macro.Shutdown

## Examples

- CM_ClosetGoblin: ClosetGoblinCharacterWindow -> ClosetGoblinCharacterWindow.OnShutdown -> ClosetGoblinCharacterWindow.OnShutdown
- CM_ClosetGoblin: ClosetGoblinZoneWindow -> ClosetGoblinZoneWindow.OnShutdown -> ClosetGoblinZoneWindow.OnShutdown
- bigger_MacroWindow: EA_Window_Macro -> EA_Window_Macro.OnShutdown -> EA_Window_Macro.Shutdown

## Related APIs

- none

## Used With

- none

## Triggered By

- none

## Affects

- none

## Notes

- none
