# Usage By Addon

## XML To Lua Examples

| Addon | Frame | Event | Lua Function |
| --- | --- | --- | --- |
| InfoScroller | InfoScrollerTemplateLabel1 | OnHyperLinkLButtonUp | EA_ChatWindow.OnHyperLinkLButtonUp |
| InfoScroller | InfoScrollerTemplateLabel1 | OnHyperLinkRButtonUp | EA_ChatWindow.OnHyperLinkRButtonUp |
| InfoScroller | InfoScrollerTemplateLabel2 | OnHyperLinkLButtonUp | EA_ChatWindow.OnHyperLinkLButtonUp |
| InfoScroller | InfoScrollerTemplateLabel2 | OnHyperLinkRButtonUp | EA_ChatWindow.OnHyperLinkRButtonUp |
| InfoScroller | InfoScrollerTemplateLabel3 | OnHyperLinkLButtonUp | EA_ChatWindow.OnHyperLinkLButtonUp |
| InfoScroller | InfoScrollerTemplateLabel3 | OnHyperLinkRButtonUp | EA_ChatWindow.OnHyperLinkRButtonUp |
| InfoScroller | InfoScrollerTemplateLabel4 | OnHyperLinkLButtonUp | EA_ChatWindow.OnHyperLinkLButtonUp |
| InfoScroller | InfoScrollerTemplateLabel4 | OnHyperLinkRButtonUp | EA_ChatWindow.OnHyperLinkRButtonUp |
| InfoScroller | InfoScrollerTemplateLabel5 | OnHyperLinkLButtonUp | EA_ChatWindow.OnHyperLinkLButtonUp |
| InfoScroller | InfoScrollerTemplateLabel5 | OnHyperLinkRButtonUp | EA_ChatWindow.OnHyperLinkRButtonUp |
| PartyCast | PartyCastWindow_TemplateTargetWindowLabel | OnHyperLinkLButtonUp | EA_ChatWindow.OnHyperLinkLButtonUp |
| PartyCast | PartyCastWindow_TemplateTargetWindowLabel | OnHyperLinkRButtonUp | EA_ChatWindow.OnHyperLinkRButtonUp |

## Representative Function Usage

- CreateWindow: InfoScroller -> CreateWindow("InfoScrollerMainWindow", true)
- CreateWindow: Moth -> CreateWindow("Moth", true)
- CreateWindow: Soloq -> CreateWindow(overviewWindowName, false)
- CreateWindow: TidyChat -> CreateWindow(c_TEXT_ENTRY_ANCHOR, false)
- CreateWindow: TidyRoll -> CreateWindow(c_TROLL_AUTO_ROLL_WINDOW, false)
- CreateWindow: TidyRoll -> CreateWindow(c_TIDY_ROLL_ANCHOR, false)
- CreateWindowFromTemplate: InfoScroller -> CreateWindowFromTemplate(WindowName, "InfoScrollerTemplate", "InfoScrollerMainWindow")
- CreateWindowFromTemplate: InfoScroller -> CreateWindowFromTemplate(w.name, base, w.parent)
- CreateWindowFromTemplate: InfoScroller -> CreateWindowFromTemplate(w.name, base, w.parent)
- CreateWindowFromTemplate: InfoScroller -> CreateWindowFromTemplate(w.name, "EA_Button_DefaultWindowClose", w.parent)
- CreateWindowFromTemplate: InfoScroller -> CreateWindowFromTemplate(w.name, base, w.parent)
- CreateWindowFromTemplate: InfoScroller -> CreateWindowFromTemplate(w.name, base, w.parent)
- DefaultColor.SetWindowTint: TidyRoll -> DefaultColor.SetWindowTint(rowName.."Background", color)
- DestroyWindow: InfoScroller -> DestroyWindow(self.name)
- DestroyWindow: PartyCast -> DestroyWindow(self.name)
- DestroyWindow: TidyRoll -> DestroyWindow(self:GetName())
- DestroyWindow: minesweep -> DestroyWindow("MineSweepWindow")
- DestroyWindow: minesweep -> DestroyWindow("MineSweepWindow")
- DialogManager.MakeOneButtonDialog: TidyChat -> DialogManager.MakeOneButtonDialog(L "Log is empty\n", L "Ok")
- DialogManager.MakeTwoButtonDialog: Lib RuString -> DialogManager.MakeTwoButtonDialog(GetStringFromTable("CustomizeUiStrings",StringTables.CustomizeUi.TEXT_UI_MOD_SETTINGS_CHANGED_DIALOG), GetString(StringTables.Default.LABEL_YES), function()BroadcastEvent(SystemData.Events.RELOAD_INTERFACE)end, GetString(StringTables.Default.LABEL_NO), nil)
