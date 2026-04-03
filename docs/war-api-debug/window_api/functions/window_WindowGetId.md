# WindowGetId

- Category: Window Function
- Confidence level: HIGH
- Confidence score: 90/100
- Seen in: 1 addons

## Confidence Assessment

- Level: HIGH

- Score: 90/100

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, matches a known engine namespace, called globally with no local definition.

## Evidence Signals

- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +10 Argument pattern is consistent: Observed argument positions remain stable.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | CM_ClosetGoblin |
| Files seen in | ClosetGoblinCharacterWindow.lua, ClosetGoblinOptionWindow.lua, ClosetGoblinZoneWindow.lua |
| Namespaces detected | WindowGetId |
| Source kinds | lua_calls |
| Example locations | CM_ClosetGoblin: EquipmentLButtonUp, CM_ClosetGoblin: EquipmentMouseOver, CM_ClosetGoblin: EquipmentRButtonUp, CM_ClosetGoblin: HandleDrag, CM_ClosetGoblin: OnClickSetRow, CM_ClosetGoblin: OnClickZoneRow |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 11 |
| Global usage count | 11 |
| Local definition count | 0 |
| Documentation references | 0 |
| Initialization flow references | 0 |
| Known engine namespace | yes |
| Default UI presence | yes |
| Event binding presence | no |
| Observed in XML and Lua | no |
| Consistent role | no |
| Consistent arguments | yes |
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

## Signature (inferred)

```lua
WindowGetId(windowName)
```

## Description

Observed querying runtime window state or metadata.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| windowName | Observed as a target window name. | Observed values: SystemData.ActiveWindow.name, SystemData.MouseOverWindow.name |

## Returns

- Not confidently inferable from contract artifacts alone.

## Side Effects

- No side effect is confidently inferable from contract artifacts alone.

## Seen In

- CM_ClosetGoblin

## Examples

- CM_ClosetGoblin: EquipmentLButtonUp -> WindowGetId(SystemData.ActiveWindow.name)
- CM_ClosetGoblin: EquipmentMouseOver -> WindowGetId(SystemData.ActiveWindow.name)
- CM_ClosetGoblin: EquipmentRButtonUp -> WindowGetId(SystemData.ActiveWindow.name)
- CM_ClosetGoblin: HandleDrag -> WindowGetId(SystemData.ActiveWindow.name)
- CM_ClosetGoblin: OnClickSetRow -> WindowGetId(SystemData.ActiveWindow.name)
- CM_ClosetGoblin: OnClickZoneRow -> WindowGetId(SystemData.ActiveWindow.name)

## Used With

- [SystemData.ActiveWindow.name](../../systemdata/fields/systemdata_SystemData.ActiveWindow.name.md) (HIGH 100/100) - SystemData Field
- [EA_Window_ContextMenu.AddMenuItem](../../globals/functions/global_EA_Window_ContextMenu.AddMenuItem.md) (HIGH 90/100) - Global Function
- [EA_Window_ContextMenu.CreateContextMenu](../../globals/functions/global_EA_Window_ContextMenu.CreateContextMenu.md) (HIGH 90/100) - Global Function
- [EA_Window_ContextMenu.Finalize](../../globals/functions/global_EA_Window_ContextMenu.Finalize.md) (HIGH 90/100) - Global Function

## Affects

- [SystemData.ActiveWindow.name](../../systemdata/fields/systemdata_SystemData.ActiveWindow.name.md) (HIGH 100/100) - SystemData Field

## Notes

- Only one addon surfaced this symbol in the current corpus.
