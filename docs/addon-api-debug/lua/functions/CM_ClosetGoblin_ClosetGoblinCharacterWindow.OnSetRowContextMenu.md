# Function ClosetGoblinCharacterWindow.OnSetRowContextMenu

- Addon: CM_ClosetGoblin
- Kind: handler
- Module: ClosetGoblinCharacterWindow
- Local: no
- Source: `C:/Return of Reckoning/Interface/AddOns/WAR_API_Ref/.debug/source-subset/ClosetGoblin/ClosetGoblinCharacterWindow.lua:123`

## Parameters

- none

## Aliases

- none

## Calls

| Call | Line | Arguments |
| --- | --- | --- |
| WindowGetId | 124 | SystemData.ActiveWindow.name |
| ClosetGoblinCharacterWindow.GetCurrentSet | 127 |  |
| ClosetGoblinCharacterWindow.UpdateSetList | 129 |  |
| ClosetGoblinCharacterWindow.UpdateSlotIcons | 130 |  |
| EA_Window_ContextMenu.CreateContextMenu | 132 | SystemData.ActiveWindow.name |
| EA_Window_ContextMenu.AddMenuItem | 133 | cgL["Rename_Set"], ClosetGoblinCharacterWindow.OnRowMenuRenameSetClick, false, true |
| EA_Window_ContextMenu.AddMenuItem | 134 | cgL["Import_Current"], ClosetGoblinCharacterWindow.OnRowMenuImportCurrentClick, false, true |
| EA_Window_ContextMenu.AddMenuItem | 135 | cgL["Copy"], ClosetGoblinCharacterWindow.OnRowMenuCopyClick, false, true |
| EA_Window_ContextMenu.AddMenuItem | 136 | cgL["Paste"], ClosetGoblinCharacterWindow.OnRowMenuPasteClick, not ClosetGoblinCharacterWindow.IsValidPasteTarget(dataIndex), true |
| ClosetGoblinCharacterWindow.IsValidPasteTarget | 136 | dataIndex |
| EA_Window_ContextMenu.AddMenuItem | 137 | cgL["Associate_with_Zone"], ClosetGoblinZoneWindow.Show, false, true |
| EA_Window_ContextMenu.AddMenuItem | 138 | cgL["Link_with_tactics_set"]:gsub(L "#1#",L "1"), ClosetGoblinCharacterWindow.OnRowMenuLinkTactics, set.tactics==L "1", true |
| cgL.Link_with_tactics_set:gsub | 138 | L "#1#", L "1" |
| EA_Window_ContextMenu.AddMenuItem | 139 | cgL["Link_with_tactics_set"]:gsub(L "#1#",L "2"), ClosetGoblinCharacterWindow.OnRowMenuLinkTactics, set.tactics==L "2", true |
| cgL.Link_with_tactics_set:gsub | 139 | L "#1#", L "2" |
| EA_Window_ContextMenu.AddMenuItem | 140 | cgL["Link_with_tactics_set"]:gsub(L "#1#",L "3"), ClosetGoblinCharacterWindow.OnRowMenuLinkTactics, set.tactics==L "3", true |
| cgL.Link_with_tactics_set:gsub | 140 | L "#1#", L "3" |
| EA_Window_ContextMenu.AddMenuItem | 141 | cgL["Link_with_tactics_set"]:gsub(L "#1#",L "4"), ClosetGoblinCharacterWindow.OnRowMenuLinkTactics, set.tactics==L "4", true |
| cgL.Link_with_tactics_set:gsub | 141 | L "#1#", L "4" |
| EA_Window_ContextMenu.AddMenuItem | 142 | cgL["Link_with_tactics_set"]:gsub(L "#1#",L "5"), ClosetGoblinCharacterWindow.OnRowMenuLinkTactics, set.tactics==L "5", true |
| cgL.Link_with_tactics_set:gsub | 142 | L "#1#", L "5" |
| EA_Window_ContextMenu.AddMenuItem | 143 | cgL["Unlink_with_tactics_set"], ClosetGoblinCharacterWindow.OnRowMenuLinkTactics, set.tactics==L "x", true |
| EA_Window_ContextMenu.Finalize | 144 |  |

## Event Registrations

| Event | Scope | Handler |
| --- | --- | --- |

## State Writes

- ClosetGoblinCharacterWindow.GetCurrentSet
- ClosetGoblinCharacterWindow.activeMenuRow
- ClosetGoblinCharacterWindow.selectedSetDataIndex
- ClosetGoblinCharacterWindowContentsSetList.PopulatorIndices
