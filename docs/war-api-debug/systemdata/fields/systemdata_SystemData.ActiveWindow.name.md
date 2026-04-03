# SystemData.ActiveWindow.name

- Category: SystemData Field
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 1 addons

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 105

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, matches a known engine namespace, referenced by generated docs or reference files.

## Evidence Signals

- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +25 Referenced by generated docs or reference files: The symbol is reinforced outside a single call page.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | CM_ClosetGoblin |
| Files seen in | ClosetGoblinCharacterWindow.lua, ClosetGoblinOptionWindow.lua, ClosetGoblinZoneWindow.lua |
| Namespaces detected | SystemData |
| Source kinds | lua_call |
| Example locations | EquipmentLButtonUp, EquipmentMouseOver, EquipmentRButtonUp, HandleDrag, OnClickSetRow, OnClickZoneRow |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 10 |
| Global usage count | 10 |
| Local definition count | 0 |
| Documentation references | 1 |
| Initialization flow references | 0 |
| Known engine namespace | yes |
| Default UI presence | yes |
| Event binding presence | no |
| Observed in XML and Lua | no |
| Consistent role | no |
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

SystemData.SystemData.ActiveWindow.name field accessed by 1 addons; commonly found in EquipmentLButtonUp and EquipmentMouseOver, EquipmentRButtonUp, HandleDrag, OnClickSetRow, OnClickZoneRow, OnLButtonUp, OnSetRowContextMenu, OnSetZoneRowContextMenu, lua_call contexts.

## Seen In

- CM_ClosetGoblin

## Related APIs

- [EA_Window_ContextMenu.AddMenuItem](../../globals/functions/global_EA_Window_ContextMenu.AddMenuItem.md) (HIGH 90/100) - Global Function
- [EA_Window_ContextMenu.CreateContextMenu](../../globals/functions/global_EA_Window_ContextMenu.CreateContextMenu.md) (HIGH 90/100) - Global Function
- [EA_Window_ContextMenu.Finalize](../../globals/functions/global_EA_Window_ContextMenu.Finalize.md) (HIGH 90/100) - Global Function
- [WindowGetId](../../window_api/functions/window_WindowGetId.md) (HIGH 90/100) - Window Function
- [Cursor.Clear](../../globals/functions/global_Cursor.Clear.md) (MEDIUM 45/100) - Global Function
- [Cursor.IconOnCursor](../../globals/functions/global_Cursor.IconOnCursor.md) (MEDIUM 45/100) - Global Function

## Used With

- [EA_Window_ContextMenu.AddMenuItem](../../globals/functions/global_EA_Window_ContextMenu.AddMenuItem.md) (HIGH 90/100) - Global Function
- [EA_Window_ContextMenu.CreateContextMenu](../../globals/functions/global_EA_Window_ContextMenu.CreateContextMenu.md) (HIGH 90/100) - Global Function
- [EA_Window_ContextMenu.Finalize](../../globals/functions/global_EA_Window_ContextMenu.Finalize.md) (HIGH 90/100) - Global Function
- [WindowGetId](../../window_api/functions/window_WindowGetId.md) (HIGH 90/100) - Window Function

## Notes

- Observed in contexts: EquipmentLButtonUp, EquipmentMouseOver, EquipmentRButtonUp, HandleDrag, OnClickSetRow, OnClickZoneRow
