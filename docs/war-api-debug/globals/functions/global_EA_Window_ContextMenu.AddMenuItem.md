# EA_Window_ContextMenu.AddMenuItem

- Category: Global Function
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
| Files seen in | ClosetGoblinCharacterWindow.lua, ClosetGoblinZoneWindow.lua |
| Namespaces detected | EA_Window_ContextMenu |
| Source kinds | lua_calls |
| Example locations | CM_ClosetGoblin: OnSetRowContextMenu, CM_ClosetGoblin: OnSetZoneRowContextMenu |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 12 |
| Global usage count | 12 |
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
EA_Window_ContextMenu.AddMenuItem(arg1, arg2, arg3, arg4)
```

## Description

Observed as a global function across 1 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a runtime window or control identifier. | Observed values: cgL["Associate_with_Zone"], cgL["Copy"], cgL["Import_Current"] |
| arg2 | Observed as a function or method reference. | Observed values: ClosetGoblinCharacterWindow.OnRowMenuCopyClick, ClosetGoblinCharacterWindow.OnRowMenuImportCurrentClick, ClosetGoblinCharacterWindow.OnRowMenuLinkTactics |
| arg3 | Observed as a function or method reference. | Observed values: false, not ClosetGoblinCharacterWindow.IsValidPasteTarget(dataIndex), set.tactics==L "1" |
| arg4 | Observed as a boolean toggle. | Observed values: true |

## Returns

- Not confidently inferable from contract artifacts alone.

## Side Effects

- No side effect is confidently inferable from contract artifacts alone.

## Seen In

- CM_ClosetGoblin

## Examples

- CM_ClosetGoblin: OnSetRowContextMenu -> EA_Window_ContextMenu.AddMenuItem(cgL["Rename_Set"], ClosetGoblinCharacterWindow.OnRowMenuRenameSetClick, false, true)
- CM_ClosetGoblin: OnSetRowContextMenu -> EA_Window_ContextMenu.AddMenuItem(cgL["Import_Current"], ClosetGoblinCharacterWindow.OnRowMenuImportCurrentClick, false, true)
- CM_ClosetGoblin: OnSetRowContextMenu -> EA_Window_ContextMenu.AddMenuItem(cgL["Copy"], ClosetGoblinCharacterWindow.OnRowMenuCopyClick, false, true)
- CM_ClosetGoblin: OnSetRowContextMenu -> EA_Window_ContextMenu.AddMenuItem(cgL["Paste"], ClosetGoblinCharacterWindow.OnRowMenuPasteClick, not ClosetGoblinCharacterWindow.IsValidPasteTarget(dataIndex), true)
- CM_ClosetGoblin: OnSetRowContextMenu -> EA_Window_ContextMenu.AddMenuItem(cgL["Associate_with_Zone"], ClosetGoblinZoneWindow.Show, false, true)
- CM_ClosetGoblin: OnSetRowContextMenu -> EA_Window_ContextMenu.AddMenuItem(cgL["Link_with_tactics_set"]:gsub(L "#1#",L "1"), ClosetGoblinCharacterWindow.OnRowMenuLinkTactics, set.tactics==L "1", true)

## Used With

- [SystemData.ActiveWindow.name](../../systemdata/fields/systemdata_SystemData.ActiveWindow.name.md) (HIGH 100/100) - SystemData Field
- [EA_Window_ContextMenu.CreateContextMenu](global_EA_Window_ContextMenu.CreateContextMenu.md) (HIGH 90/100) - Global Function
- [EA_Window_ContextMenu.Finalize](global_EA_Window_ContextMenu.Finalize.md) (HIGH 90/100) - Global Function
- [WindowGetId](../../window_api/functions/window_WindowGetId.md) (HIGH 90/100) - Window Function

## Affects

- [SystemData.ActiveWindow.name](../../systemdata/fields/systemdata_SystemData.ActiveWindow.name.md) (HIGH 100/100) - SystemData Field

## Notes

- Only one addon surfaced this symbol in the current corpus.
- Canonical entry built from observed call sites, not from engine source or decompiled definitions.
