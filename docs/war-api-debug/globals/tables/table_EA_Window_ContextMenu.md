# EA_Window_ContextMenu

- Category: Global Table
- Confidence level: HIGH
- Confidence score: 80/100

## Confidence Assessment

- Level: HIGH

- Score: 80/100

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, matches a known engine namespace, called globally with no local definition.

## Evidence Signals

- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | CM_ClosetGoblin |
| Namespaces detected | EA_Window_ContextMenu |
| Source kinds | lua_calls |
| Example locations | CM_ClosetGoblin: OnLButtonUp, CM_ClosetGoblin: OnSetRowContextMenu, CM_ClosetGoblin: OnSetZoneRowContextMenu |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 18 |
| Global usage count | 3 |
| Local definition count | 0 |
| Documentation references | 0 |
| Initialization flow references | 0 |
| Known engine namespace | yes |
| Default UI presence | yes |
| Event binding presence | no |
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

Shared function table with 3 member functions; the primary API surface for 1 addons.

## Functions

- EA_Window_ContextMenu.AddMenuItem
- EA_Window_ContextMenu.CreateContextMenu
- EA_Window_ContextMenu.Finalize

## Observed Members

- none

## Seen In

- CM_ClosetGoblin

## Examples

- CM_ClosetGoblin: OnLButtonUp -> EA_Window_ContextMenu.CreateContextMenu(SystemData.ActiveWindow.name)
- CM_ClosetGoblin: OnLButtonUp -> EA_Window_ContextMenu.Finalize()
- CM_ClosetGoblin: OnSetRowContextMenu -> EA_Window_ContextMenu.CreateContextMenu(SystemData.ActiveWindow.name)
- CM_ClosetGoblin: OnSetRowContextMenu -> EA_Window_ContextMenu.AddMenuItem(cgL["Rename_Set"], ClosetGoblinCharacterWindow.OnRowMenuRenameSetClick, false, true)
- CM_ClosetGoblin: OnSetRowContextMenu -> EA_Window_ContextMenu.AddMenuItem(cgL["Import_Current"], ClosetGoblinCharacterWindow.OnRowMenuImportCurrentClick, false, true)
- CM_ClosetGoblin: OnSetRowContextMenu -> EA_Window_ContextMenu.AddMenuItem(cgL["Copy"], ClosetGoblinCharacterWindow.OnRowMenuCopyClick, false, true)

## Notes

- none
