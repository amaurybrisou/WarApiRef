# Usage By Addon

## XML To Lua Examples

| Addon | Frame | Event | Lua Function |
| --- | --- | --- | --- |
| TidyChat | TChatCheckboxTemplate | OnLButtonUp | TidyChat.Options.OnCheckboxLBU |
| TidyChat | TChatTabButton | OnLButtonUp | TidyChat.Options.OnTabLBU |
| TidyChat | TChatTabWindowsTemplateSelectWindowCombo | OnSelChanged | TidyChat.Options.UpdateGroupTabs |
| TidyChat | TidyChatCopy | OnHidden | TidyChat.Copy.OnHidden |
| TidyChat | TidyChatCopy | OnMouseWheel | TidyChat.Copy.OnMouseWheel |
| TidyChat | TidyChatCopy | OnShown | TidyChat.Copy.OnShown |
| TidyChat | TidyChatCopyClose | OnLButtonUp | TidyChat.Copy.OnClose |
| TidyChat | TidyChatCopyNext | OnLButtonUp | TidyChat.Copy.CopyNext |
| TidyChat | TidyChatCopyPrev | OnLButtonUp | TidyChat.Copy.CopyPrev |
| TidyChat | TidyChatLootRoll | OnHidden | TidyChat.LootRoll.OnHidden |
| TidyChat | TidyChatLootRoll | OnShown | TidyChat.LootRoll.OnShown |
| TidyChat | TidyChatLootRollClose | OnLButtonUp | TidyChat.LootRoll.OnClose |

## Representative Function Usage

- DefaultColor.SetWindowTint: TidyRoll -> DefaultColor.SetWindowTint(rowName.."Background", color)
- DialogManager.MakeOneButtonDialog: TidyChat -> DialogManager.MakeOneButtonDialog(L "Log is empty\n", L "Ok")
- DoesWindowExist: TidyChat -> DoesWindowExist(c_TEXT_ENTRY_WINDOW.."EntryBoxLanguageButton")
- DoesWindowExist: TidyRoll -> DoesWindowExist("EA_Window_LootRoll")
- DoesWindowExist: TidyRoll -> DoesWindowExist(radioGroupName..index)
- DoesWindowExist: TidyRoll -> DoesWindowExist(radioGroupName..index)
- EA_ChatWindow.OnSettingsChanged: TidyChat -> EA_ChatWindow.OnSettingsChanged()
- EA_ChatWindow.Print: NPC Item Sale Price -> EA_ChatWindow.Print(L "Dump Table Cleared")
- EA_ChatWindow.Print: NPC Item Sale Price -> EA_ChatWindow.Print(L "Dumped item: "..itemData.name)
- EA_ChatWindow.Print: NPC Item Sale Price -> EA_ChatWindow.Print(L "Item Dump contains item: "..itemData.name)
- EA_ChatWindow.Print: NPC Item Sale Price -> EA_ChatWindow.Print(L "Nisp Installed and Enabled")
- EA_ChatWindow.Print: NPC Item Sale Price -> EA_ChatWindow.Print(L "Nisp Initialized and Enabled (/nisp for commands)")
- EA_ChatWindow.Print: NPC Item Sale Price -> EA_ChatWindow.Print(L "Nisp Initialized, but disabled (/nisp for commands)")
- EA_ChatWindow.ShowTabCycleButtons: TidyChat -> EA_ChatWindow.ShowTabCycleButtons(wndGroupName, false)
- EA_ChatWindow.UpdateTabScrollWindow: TidyChat -> EA_ChatWindow.UpdateTabScrollWindow(wndGroupId)
- EA_Window_ContextMenu.AddMenuItem: TidyChat -> EA_Window_ContextMenu.AddMenuItem(L "Copy...", TidyChatCopy.OnCopyButton, false, true)
- EA_Window_ContextMenu.AddMenuItem: TidyChat -> EA_Window_ContextMenu.AddMenuItem(L "To Bottom", TidyChatFrames.OnToBottomButton, false, true)
- EA_Window_ContextMenu.Finalize: TidyChat -> EA_Window_ContextMenu.Finalize()
- GetIconData: Moth -> GetIconData(Icons.GetCareerIconIDFromCareerLine(text))
- GetIconData: TidyRoll -> GetIconData(careerIconNum)
