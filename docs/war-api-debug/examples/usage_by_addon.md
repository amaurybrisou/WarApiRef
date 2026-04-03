# Usage By Addon

## XML To Lua Examples

| Addon | Frame | Event | Lua Function |
| --- | --- | --- | --- |
| CM_ClosetGoblin |  | OnLButtonUp | ClosetGoblinZoneWindow.OnClickZoneRow |
| CM_ClosetGoblin |  | OnRButtonUp | ClosetGoblinZoneWindow.OnSetZoneRowContextMenu |
| CM_ClosetGoblin |  | OnLButtonUp | ClosetGoblinZoneWindow.OnClickZoneRow |
| CM_ClosetGoblin |  | OnRButtonUp | ClosetGoblinZoneWindow.OnSetZoneRowContextMenu |
| CM_ClosetGoblin |  | OnLButtonUp | ClosetGoblinCharacterWindow.HotbarPageUpProxy |
| CM_ClosetGoblin |  | OnLButtonUp | ClosetGoblinCharacterWindow.HotbarPageDownProxy |
| CM_ClosetGoblin |  | OnShown | ClosetGoblinCharacterWindow.OnShow |
| CM_ClosetGoblin |  | OnHidden | ClosetGoblinCharacterWindow.OnHide |
| CM_ClosetGoblin |  | OnShutdown | ClosetGoblinCharacterWindow.OnShutdown |
| CM_ClosetGoblin |  | OnLButtonUp | ClosetGoblinCharacterWindow.Hide |
| CM_ClosetGoblin |  | OnLButtonUp | ClosetGoblinCharacterWindow.OnClickNewSetButton |
| CM_ClosetGoblin |  | OnLButtonUp | ClosetGoblinCharacterWindow.OnClickDeleteSetButton |
| CM_ClosetGoblin |  | OnLButtonUp | ClosetGoblinCharacterWindow.OnClickZoneConfigButton |
| CM_ClosetGoblin |  | OnLButtonUp | ClosetGoblinCharacterWindow.EquipmentLButtonUp |
| CM_ClosetGoblin |  | OnLButtonUp | ClosetGoblinCharacterWindow.UseItemRackToggled |
| CM_ClosetGoblin |  | OnLButtonDown | ClosetGoblinCharacterWindow.EquipmentLButtonDown |
| CM_ClosetGoblin |  | OnRButtonUp | ClosetGoblinCharacterWindow.EquipmentRButtonUp |
| CM_ClosetGoblin |  | OnLButtonUp | ClosetGoblinCharacterWindow.OnClickSortNameButton |
| CM_ClosetGoblin |  | OnMouseOver | ClosetGoblinCharacterWindow.EquipmentMouseOver |
| CM_ClosetGoblin |  | OnLButtonUp | ClosetGoblinCharacterWindow.OnClickSortTacticsButton |
| CM_ClosetGoblin |  | OnMouseOverEnd | ClosetGoblinCharacterWindow.EquipmentMouseOverEnd |
| CM_ClosetGoblin |  | OnLButtonUp | ClosetGoblinCharacterWindow.HotbarChangeToggled1 |
| CM_ClosetGoblin |  | OnLButtonUp | ClosetGoblinCharacterWindow.HotbarChangeToggled2 |
| CM_ClosetGoblin |  | OnLButtonUp | ClosetGoblinCharacterWindow.HotbarChangeToggled3 |
| CM_ClosetGoblin |  | OnLButtonUp | ClosetGoblinCharacterWindow.HotbarChangeToggled4 |
| CM_ClosetGoblin |  | OnLButtonUp | ClosetGoblinCharacterWindow.HotbarChangeToggled5 |
| CM_ClosetGoblin |  | OnLButtonDown | ClosetGoblinCharacterWindow.EquipmentLButtonDown |
| CM_ClosetGoblin |  | OnLButtonUp | ClosetGoblinCharacterWindow.EquipmentLButtonUp |
| CM_ClosetGoblin |  | OnMouseOver | ClosetGoblinCharacterWindow.ShowShowHelm |
| CM_ClosetGoblin |  | OnMouseOverEnd | ClosetGoblinCharacterWindow.HideShowHelm |
| CM_ClosetGoblin |  | OnMouseOver | ClosetGoblinCharacterWindow.ShowShowHelmOnly |
| CM_ClosetGoblin |  | OnMouseOverEnd | ClosetGoblinCharacterWindow.HideShowHelm |
| CM_ClosetGoblin |  | OnLButtonUp | ClosetGoblinCharacterWindow.ShowHelm |
| CM_ClosetGoblin |  | OnLButtonDown | ClosetGoblinCharacterWindow.EquipmentLButtonDown |
| CM_ClosetGoblin |  | OnLButtonUp | ClosetGoblinCharacterWindow.EquipmentLButtonUp |
| CM_ClosetGoblin |  | OnMouseOver | ClosetGoblinCharacterWindow.ShowCloakOptions |
| CM_ClosetGoblin |  | OnMouseOverEnd | ClosetGoblinCharacterWindow.HideCloakOptions |
| CM_ClosetGoblin |  | OnMouseOver | ClosetGoblinCharacterWindow.ShowShowCloakOnly |
| CM_ClosetGoblin |  | OnMouseOverEnd | ClosetGoblinCharacterWindow.HideCloakOptions |
| CM_ClosetGoblin |  | OnLButtonUp | ClosetGoblinCharacterWindow.ShowCloak |
| CM_ClosetGoblin |  | OnMouseOver | ClosetGoblinCharacterWindow.ShowShowCloakHeraldryOnly |
| CM_ClosetGoblin |  | OnMouseOverEnd | ClosetGoblinCharacterWindow.HideCloakOptions |
| CM_ClosetGoblin |  | OnLButtonUp | ClosetGoblinCharacterWindow.ShowCloakHeraldry |
| CM_ClosetGoblin |  | OnLButtonUp | ClosetGoblinCharacterWindow.OnClickSetRow |
| CM_ClosetGoblin |  | OnRButtonUp | ClosetGoblinCharacterWindow.OnSetRowContextMenu |
| CM_ClosetGoblin |  | OnShown | ClosetGoblinZoneWindow.OnShow |
| CM_ClosetGoblin |  | OnHidden | ClosetGoblinZoneWindow.OnHide |
| CM_ClosetGoblin |  | OnShutdown | ClosetGoblinZoneWindow.OnShutdown |
| CM_ClosetGoblin |  | OnLButtonUp | ClosetGoblinZoneWindow.Hide |
| CM_ClosetGoblin |  | OnLButtonUp | ClosetGoblinZoneWindow.OnClickChangeZoneOptionButton |
| CM_ClosetGoblin |  | OnLButtonUp | ClosetGoblinCharacterWindow.OnClickSetRow |
| CM_ClosetGoblin |  | OnRButtonUp | ClosetGoblinCharacterWindow.OnSetRowContextMenu |
| CM_ClosetGoblin |  | OnInitialize | ClosetGoblinOptionWindow.OnInitialize |
| CM_ClosetGoblin |  | OnLButtonUp | ClosetGoblinOptionWindow.OnLButtonUp |
| CM_ClosetGoblin |  | OnRButtonUp | ClosetGoblinOptionWindow.OnRButtonUp |
| CM_ClosetGoblin |  | OnLButtonUp | ClosetGoblinOptionWindow.OnLButtonUp |
| CM_ClosetGoblin |  | OnRButtonUp | ClosetGoblinOptionWindow.OnRButtonUp |

## Representative Function Usage

- Cursor.Clear: CM_ClosetGoblin -> Cursor.Clear()
- Cursor.IconOnCursor: CM_ClosetGoblin -> Cursor.IconOnCursor()
- DialogManager.MakeTextEntryDialog: CM_ClosetGoblin -> DialogManager.MakeTextEntryDialog(cgL["Set_Name"], cgL["Enter_set_name"], L "", ClosetGoblinCharacterWindow.OnSubmitNewSetName, nil)
- DialogManager.MakeTextEntryDialog: CM_ClosetGoblin -> DialogManager.MakeTextEntryDialog(cgL["New_set_name"], cgL["Enter_set_name"], set.name, ClosetGoblinCharacterWindow.OnSubmitSetRename, nil)
- DialogManager.MakeTextEntryDialog: CM_ClosetGoblin -> DialogManager.MakeTextEntryDialog(cgL["Set name"], L "Enter set name :", name, ClosetGoblinCharacterWindow.OnSubmitNewSetName, nil)
- DialogManager.MakeTextEntryDialog: CM_ClosetGoblin -> DialogManager.MakeTextEntryDialog(cgL["New_set_name"], cgL["Enter_set_name"], name, ClosetGoblinCharacterWindow.OnSubmitSetRename, nil)
- DialogManager.MakeTwoButtonDialog: CM_ClosetGoblin -> DialogManager.MakeTwoButtonDialog(cgL["confirm_delete_set"]:gsub(L "#1#",set.name), confirmYes, ClosetGoblinCharacterWindow.OnConfirmDeleteSet, confirmNo, nil)
- EA_Window_ContextMenu.AddMenuItem: CM_ClosetGoblin -> EA_Window_ContextMenu.AddMenuItem(cgL["Rename_Set"], ClosetGoblinCharacterWindow.OnRowMenuRenameSetClick, false, true)
- EA_Window_ContextMenu.AddMenuItem: CM_ClosetGoblin -> EA_Window_ContextMenu.AddMenuItem(cgL["Import_Current"], ClosetGoblinCharacterWindow.OnRowMenuImportCurrentClick, false, true)
- EA_Window_ContextMenu.AddMenuItem: CM_ClosetGoblin -> EA_Window_ContextMenu.AddMenuItem(cgL["Copy"], ClosetGoblinCharacterWindow.OnRowMenuCopyClick, false, true)
- EA_Window_ContextMenu.AddMenuItem: CM_ClosetGoblin -> EA_Window_ContextMenu.AddMenuItem(cgL["Paste"], ClosetGoblinCharacterWindow.OnRowMenuPasteClick, not ClosetGoblinCharacterWindow.IsValidPasteTarget(dataIndex), true)
- EA_Window_ContextMenu.AddMenuItem: CM_ClosetGoblin -> EA_Window_ContextMenu.AddMenuItem(cgL["Associate_with_Zone"], ClosetGoblinZoneWindow.Show, false, true)
- EA_Window_ContextMenu.AddMenuItem: CM_ClosetGoblin -> EA_Window_ContextMenu.AddMenuItem(cgL["Link_with_tactics_set"]:gsub(L "#1#",L "1"), ClosetGoblinCharacterWindow.OnRowMenuLinkTactics, set.tactics==L "1", true)
- EA_Window_ContextMenu.CreateContextMenu: CM_ClosetGoblin -> EA_Window_ContextMenu.CreateContextMenu(SystemData.ActiveWindow.name)
- EA_Window_ContextMenu.CreateContextMenu: CM_ClosetGoblin -> EA_Window_ContextMenu.CreateContextMenu(SystemData.ActiveWindow.name)
- EA_Window_ContextMenu.CreateContextMenu: CM_ClosetGoblin -> EA_Window_ContextMenu.CreateContextMenu(SystemData.ActiveWindow.name)
- EA_Window_ContextMenu.Finalize: CM_ClosetGoblin -> EA_Window_ContextMenu.Finalize()
- EA_Window_ContextMenu.Finalize: CM_ClosetGoblin -> EA_Window_ContextMenu.Finalize()
- EA_Window_ContextMenu.Finalize: CM_ClosetGoblin -> EA_Window_ContextMenu.Finalize()
- LibSlash.RegisterSlashCmd: CM_ClosetGoblin -> LibSlash.RegisterSlashCmd("cg", ClosetGoblin.OnSlashCommand)
