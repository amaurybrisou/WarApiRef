# Window Patterns

## Window creation

- Confidence: MEDIUM

- Description: Observed top-level UI windows being created from XML definitions at initialization time.

- Evidence:

- Moth: CreateWindow("Moth", true)
  - TidyChat: CreateWindow(c_TEXT_ENTRY_ANCHOR, false)
  - TidyRoll: CreateWindow(c_TROLL_AUTO_ROLL_WINDOW, false)
  - TidyRoll: CreateWindow(c_TIDY_ROLL_ANCHOR, false)
  - TidyRoll: CreateWindow(c_TIDY_ROLL_TIMER, false)
  - TidyRoll: CreateWindow(c_TIDY_ROLL_ESC, false)

## Template instantiation

- Confidence: MEDIUM

- Description: Observed repeated UI elements being instantiated from XML templates at runtime.

- Evidence:

- TidyChat: CreateWindowFromTemplate(c_CHANNEL_MENU.."AllianceButton", "ChannelMenuButton", "ChatChannelSelectionWindow")
  - TidyChat: CreateWindowFromTemplate(c_CHANNEL_MENU.."AdviceButton", "ChannelMenuButton", "ChatChannelSelectionWindow")
  - TidyChat: CreateWindowFromTemplate(c_CHANNEL_MENU.."ScenarioButton", "ChannelMenuButton", "ChatChannelSelectionWindow")
  - TidyRoll: CreateWindowFromTemplate(c_TROLL_TITLE, "EA_Label_DefaultText", c_TIDY_ROLL_OPTIONS)
  - TidyRoll: CreateWindowFromTemplate(c_TROLL_VERSION, "EA_Label_DefaultText", c_TIDY_ROLL_OPTIONS)
  - TidyRoll: CreateWindowFromTemplate(c_TROLL_CLOSE, "EA_Button_DefaultWindowClose", c_TIDY_ROLL_OPTIONS)

## Runtime anchoring

- Confidence: MEDIUM

- Description: Observed layouts being finalized in Lua by clearing and re-adding anchors after window creation.

- Evidence:

- Moth: WindowAddAnchor("Moth", "bottomleft", "CursorWindow", "topleft", 0, 0)
  - Moth: WindowAddAnchor("MothHealthBar", "bottomleft", "MothBackground", "topleft", 0, yOffset)
  - TidyChat: WindowAddAnchor(c_TEXT_ENTRY_WINDOW.."EntryBox", "topleft", c_TEXT_ENTRY_WINDOW, "topleft", EntryBoxAnchorXOffset, 0)
  - TidyChat: WindowAddAnchor(c_TEXT_ENTRY_WINDOW.."EntryBox", "bottomright", c_TEXT_ENTRY_WINDOW, "bottomright", 0, 0)
  - TidyChat: WindowAddAnchor(c_TEXT_ENTRY_WINDOW, TEAnchorPoint.."left", TEAnchorRelativeWindow, TEAnchorRelativePoint.."left", 0, 0)
  - TidyChat: WindowAddAnchor(c_TEXT_ENTRY_WINDOW, TEAnchorPoint.."right", TEAnchorRelativeWindow, TEAnchorRelativePoint.."right", 0, 0)

## XML to Lua binding

- Confidence: HIGH

- Description: Observed XML handler events routing directly into addon Lua functions through shared engine bindings.

- Evidence:

- TidyChat: TChatCheckboxTemplate.OnLButtonUp -> TidyChat.Options.OnCheckboxLBU
  - TidyChat: TChatTabButton.OnLButtonUp -> TidyChat.Options.OnTabLBU
  - TidyChat: TChatTabWindowsTemplateSelectWindowCombo.OnSelChanged -> TidyChat.Options.UpdateGroupTabs
  - TidyChat: TidyChatCopy.OnHidden -> TidyChat.Copy.OnHidden
  - TidyChat: TidyChatCopy.OnMouseWheel -> TidyChat.Copy.OnMouseWheel
  - TidyChat: TidyChatCopy.OnShown -> TidyChat.Copy.OnShown
