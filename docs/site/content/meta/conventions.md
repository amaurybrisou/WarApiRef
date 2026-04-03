# Conventions

## Initialization pattern

- Confidence: HIGH

- Description: Addons consistently move from manifest loading into initialize hooks, then into XML creation and runtime event registration.

- Evidence:

- initialize: Creates 0 windows and calls 1 initialize hooks, Creates 3 windows and calls 1 initialize hooks
  - runtime: InfoScroller: SystemData.Events.UPDATE_PROCESSED, InfoScroller: e
  - xml: Defines 20 XML frames and 0 bound handlers, Defines 39 XML frames and 4 bound handlers

## Event registration pattern

- Confidence: HIGH

- Description: Runtime events are typically wired through RegisterEventHandler or window-scoped variants, with handler names staying module-qualified.

- Evidence:

- RegisterEventHandler: Lib RuString: RegisterEventHandler(SystemData.Events.LOADING_END, "LibRuString.OnLoad")
  - RegisterEventHandler: Lib RuString: RegisterEventHandler(SystemData.Events.RELOAD_INTERFACE, "LibRuString.OnLoad")
  - RegisterEventHandler: PartyCast: RegisterEventHandler(SystemData.Events.PLAYER_START_INTERACT_TIMER, "PartyCast.StartInteract")
  - RegisterEventHandler: PartyCast: RegisterEventHandler(SystemData.Events.INTERACT_DONE, "PartyCast.EndCast")
  - RegisterEventHandler: PartyCast: RegisterEventHandler(SystemData.Events.PLAYER_BEGIN_CAST, "PartyCast.StartCast")
  - RegisterEventHandler: PartyCast: RegisterEventHandler(SystemData.Events.PLAYER_END_CAST, "PartyCast.EndCast")
  - UnregisterEventHandler: PartyCast: UnregisterEventHandler(SystemData.Events.PLAYER_START_INTERACT_TIMER, "PartyCast.StartInteract")
  - UnregisterEventHandler: PartyCast: UnregisterEventHandler(SystemData.Events.INTERACT_DONE, "PartyCast.EndCast")

## UI creation pattern

- Confidence: HIGH

- Description: UI is commonly created from XML, then positioned in Lua through CreateWindow or CreateWindowFromTemplate followed by anchor calls.

- Evidence:

- Window creation: InfoScroller: CreateWindow("InfoScrollerMainWindow", true)
  - Window creation: Moth: CreateWindow("Moth", true)
  - Window creation: Soloq: CreateWindow(overviewWindowName, false)
  - Window creation: TidyChat: CreateWindow(c_TEXT_ENTRY_ANCHOR, false)
  - Window creation: TidyRoll: CreateWindow(c_TROLL_AUTO_ROLL_WINDOW, false)
  - Window creation: TidyRoll: CreateWindow(c_TIDY_ROLL_ANCHOR, false)
  - Template instantiation: InfoScroller: CreateWindowFromTemplate(WindowName, "InfoScrollerTemplate", "InfoScrollerMainWindow")
  - Template instantiation: InfoScroller: CreateWindowFromTemplate(w.name, base, w.parent)

## XML to Lua binding pattern

- Confidence: HIGH

- Description: XML handler names map directly to Lua functions and can be cross-checked through the bindings page.

- Evidence:

- InfoScroller: InfoScrollerTemplateLabel1.OnHyperLinkLButtonUp -> EA_ChatWindow.OnHyperLinkLButtonUp
  - InfoScroller: InfoScrollerTemplateLabel1.OnHyperLinkRButtonUp -> EA_ChatWindow.OnHyperLinkRButtonUp
  - InfoScroller: InfoScrollerTemplateLabel2.OnHyperLinkLButtonUp -> EA_ChatWindow.OnHyperLinkLButtonUp
  - InfoScroller: InfoScrollerTemplateLabel2.OnHyperLinkRButtonUp -> EA_ChatWindow.OnHyperLinkRButtonUp
  - InfoScroller: InfoScrollerTemplateLabel3.OnHyperLinkLButtonUp -> EA_ChatWindow.OnHyperLinkLButtonUp
  - InfoScroller: InfoScrollerTemplateLabel3.OnHyperLinkRButtonUp -> EA_ChatWindow.OnHyperLinkRButtonUp
  - InfoScroller: InfoScrollerTemplateLabel4.OnHyperLinkLButtonUp -> EA_ChatWindow.OnHyperLinkLButtonUp
  - InfoScroller: InfoScrollerTemplateLabel4.OnHyperLinkRButtonUp -> EA_ChatWindow.OnHyperLinkRButtonUp

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

- InfoScroller: InfoScroller.Settings
  - Lib RuString: LibRuString.Settings
  - NPC Item Sale Price: Nisp.DebugEnabled
  - NPC Item Sale Price: Nisp.DumpItemsTable
  - NPC Item Sale Price: Nisp.Enabled
  - PartyCast: PartyCast.Settings
  - Soloq: Soloq.db
  - TidyChat: TidyChat.Settings
