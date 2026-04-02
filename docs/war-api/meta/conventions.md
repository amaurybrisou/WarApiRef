# Conventions

## Initialization pattern

- Confidence: HIGH

- Description: Addons consistently move from manifest loading into initialize hooks, then into XML creation and runtime event registration.

- Evidence:

- initialize: Creates 0 windows and calls 1 initialize hooks, Creates 3 windows and calls 1 initialize hooks
  - runtime: Moth: SystemData.Events.PLAYER_TARGET_UPDATED, Registers 1 runtime events
  - xml: Defines 20 XML frames and 0 bound handlers, Defines 47 XML frames and 22 bound handlers

## Event registration pattern

- Confidence: HIGH

- Description: Runtime events are typically wired through RegisterEventHandler or window-scoped variants, with handler names staying module-qualified.

- Evidence:

- RegisterEventHandler: TidyChat: RegisterEventHandler(SystemData.Events.LOADING_END, "TidyChat.OnLoad")
  - RegisterEventHandler: TidyChat: RegisterEventHandler(SystemData.Events.RELOAD_INTERFACE, "TidyChat.OnLoad")
  - RegisterEventHandler: TidyChat: RegisterEventHandler(SystemData.Events.L_BUTTON_UP_PROCESSED, "TidyChat.OnLBU")
  - RegisterEventHandler: TidyChat: RegisterEventHandler(chatLogEventId, "TidyChat.OnChatLogUpdated")
  - RegisterEventHandler: TidyChat: RegisterEventHandler(combatLogEventId, "TidyChat.OnCombatLogUpdated")
  - RegisterEventHandler: TidyChat: RegisterEventHandler(systemLogEventId, "TidyChat.OnSystemLogUpdated")
  - UnregisterEventHandler: TidyRoll: UnregisterEventHandler(SystemData.Events.LOADING_END, "TidyRoll.OnLoad")
  - UnregisterEventHandler: TidyRoll: UnregisterEventHandler(SystemData.Events.RELOAD_INTERFACE, "TidyRoll.OnLoad")

## UI creation pattern

- Confidence: HIGH

- Description: UI is commonly created from XML, then positioned in Lua through CreateWindow or CreateWindowFromTemplate followed by anchor calls.

- Evidence:

- Window creation: Moth: CreateWindow("Moth", true)
  - Window creation: TidyChat: CreateWindow(c_TEXT_ENTRY_ANCHOR, false)
  - Window creation: TidyRoll: CreateWindow(c_TROLL_AUTO_ROLL_WINDOW, false)
  - Window creation: TidyRoll: CreateWindow(c_TIDY_ROLL_ANCHOR, false)
  - Window creation: TidyRoll: CreateWindow(c_TIDY_ROLL_TIMER, false)
  - Window creation: TidyRoll: CreateWindow(c_TIDY_ROLL_ESC, false)
  - Template instantiation: TidyChat: CreateWindowFromTemplate(c_CHANNEL_MENU.."AllianceButton", "ChannelMenuButton", "ChatChannelSelectionWindow")
  - Template instantiation: TidyChat: CreateWindowFromTemplate(c_CHANNEL_MENU.."AdviceButton", "ChannelMenuButton", "ChatChannelSelectionWindow")

## XML to Lua binding pattern

- Confidence: HIGH

- Description: XML handler names map directly to Lua functions and can be cross-checked through the bindings page.

- Evidence:

- TidyChat: TChatCheckboxTemplate.OnLButtonUp -> TidyChat.Options.OnCheckboxLBU
  - TidyChat: TChatTabButton.OnLButtonUp -> TidyChat.Options.OnTabLBU
  - TidyChat: TChatTabWindowsTemplateSelectWindowCombo.OnSelChanged -> TidyChat.Options.UpdateGroupTabs
  - TidyChat: TidyChatCopy.OnHidden -> TidyChat.Copy.OnHidden
  - TidyChat: TidyChatCopy.OnMouseWheel -> TidyChat.Copy.OnMouseWheel
  - TidyChat: TidyChatCopy.OnShown -> TidyChat.Copy.OnShown
  - TidyChat: TidyChatCopyClose.OnLButtonUp -> TidyChat.Copy.OnClose
  - TidyChat: TidyChatCopyNext.OnLButtonUp -> TidyChat.Copy.CopyNext

## XML runtime caveats

- Confidence: MEDIUM

- Description: Implementation-validated findings show that XML input and scroll layout behavior can depend on ancestor state and on outer-window sizing, even when child nodes appear correctly configured.

- Evidence:

- WhoHealedMe: a child `OnLButtonUp` target remained inert until the parent or template input chain was made input-enabled.
  - guidance: validate ancestor `handleinput` state across the clickable chain, not only on the child node.
  - caveat: treat this as a reusable runtime warning, not a guaranteed engine contract.
  - WhoHealedMe: nested scroll content dimensions initially under-reported usable space during early layout.
  - guidance: compute size from the outer parent first, then resize child content and call `ScrollWindowUpdateScrollRect`.

## XML list binding pattern

- Confidence: MEDIUM

- Description: ListBox rows are commonly bound through ListData-backed Lua tables, with ListColumns supplying text fields and Lua population callbacks handling extra row setup such as icons or reordered display.

- Evidence:

- QuickTacticSwitch: `ListData table="QTS.listDisplayData" populationfunction="QTS.PopulateDisplay"` binds a ListBox to Lua-backed row data.
  - QuickTacticSwitch: `ListColumns` binds `Name` and `Enemy`, while `QTS.PopulateDisplay` uses `QuickTacticSwitchWindowList.PopulatorIndices` to populate row icons.
  - QuickTacticSwitch: `ListBoxSetDisplayOrder` and `ListBoxGetDataIndex` are used to manage visible ordering and row-to-data mapping.
  - AggroMeter: `ListData table="AggroMeter.Listdata" populationfunction=""` suggests column-only text binding works without a custom population callback.

## State management pattern

- Confidence: MEDIUM

- Description: Persistent state is typically rooted in addon-owned globals and saved variables, then initialized before runtime hooks are attached.

- Evidence:

- NPC Item Sale Price: Nisp.DebugEnabled
  - NPC Item Sale Price: Nisp.DumpItemsTable
  - NPC Item Sale Price: Nisp.Enabled
  - TidyChat: TidyChat.Settings
