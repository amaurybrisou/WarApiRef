# Conventions

## Initialization pattern

- Confidence: HIGH

- Description: Addons consistently move from manifest loading into initialize hooks, then into XML creation and runtime event registration.

- Evidence:

- none

## Event registration pattern

- Confidence: HIGH

- Description: Runtime events are typically wired through RegisterEventHandler or window-scoped variants, with handler names staying module-qualified.

- Evidence:

- RegisterEventHandler: CM_ClosetGoblin: RegisterEventHandler(SystemData.Events.LOADING_END, "ClosetGoblin.Initialize")
  - RegisterEventHandler: CM_ClosetGoblin: RegisterEventHandler(SystemData.Events.RELOAD_INTERFACE, "ClosetGoblin.Initialize")
  - RegisterEventHandler: CM_ClosetGoblin: RegisterEventHandler(SystemData.Events.ALL_MODULES_INITIALIZED, "ClosetGoblin.OnAllModulesInitialized")
  - RegisterEventHandler: CM_ClosetGoblin: RegisterEventHandler(SystemData.Events.VISIBLE_EQUIPMENT_UPDATED, "ClosetGoblin.OnVisibleEquipmentUpdated")
  - RegisterEventHandler: CM_ClosetGoblin: RegisterEventHandler(SystemData.Events.PLAYER_INVENTORY_SLOT_UPDATED, "ClosetGoblin.OnInventorySlotUpdated")
  - RegisterEventHandler: CM_ClosetGoblin: RegisterEventHandler(SystemData.Events.UPDATE_PROCESSED, "ClosetGoblin.OnActivationWatchdog")
  - UnregisterEventHandler: CM_ClosetGoblin: UnregisterEventHandler(SystemData.Events.LOADING_END, "ClosetGoblin.Initialize")
  - UnregisterEventHandler: CM_ClosetGoblin: UnregisterEventHandler(SystemData.Events.RELOAD_INTERFACE, "ClosetGoblin.Initialize")

## UI creation pattern

- Confidence: HIGH

- Description: UI is commonly created from XML, then positioned in Lua through CreateWindow or CreateWindowFromTemplate followed by anchor calls.

- Evidence:

- Window creation: CM_ClosetGoblin: CreateWindow("yClosetGoblinButton", true)
  - Window creation: Clock: CreateWindow("ClockWindow", true)
  - XML to Lua binding: CM_ClosetGoblin: .OnLButtonUp -> ClosetGoblinZoneWindow.OnClickZoneRow
  - XML to Lua binding: CM_ClosetGoblin: .OnRButtonUp -> ClosetGoblinZoneWindow.OnSetZoneRowContextMenu
  - XML to Lua binding: CM_ClosetGoblin: .OnLButtonUp -> ClosetGoblinZoneWindow.OnClickZoneRow
  - XML to Lua binding: CM_ClosetGoblin: .OnRButtonUp -> ClosetGoblinZoneWindow.OnSetZoneRowContextMenu
  - XML to Lua binding: CM_ClosetGoblin: .OnLButtonUp -> ClosetGoblinCharacterWindow.HotbarPageUpProxy
  - XML to Lua binding: CM_ClosetGoblin: .OnLButtonUp -> ClosetGoblinCharacterWindow.HotbarPageDownProxy

## XML to Lua binding pattern

- Confidence: HIGH

- Description: XML handler names map directly to Lua functions and can be cross-checked through the bindings page.

- Evidence:

- CM_ClosetGoblin: .OnLButtonUp -> ClosetGoblinZoneWindow.OnClickZoneRow
  - CM_ClosetGoblin: .OnRButtonUp -> ClosetGoblinZoneWindow.OnSetZoneRowContextMenu
  - CM_ClosetGoblin: .OnLButtonUp -> ClosetGoblinZoneWindow.OnClickZoneRow
  - CM_ClosetGoblin: .OnRButtonUp -> ClosetGoblinZoneWindow.OnSetZoneRowContextMenu
  - CM_ClosetGoblin: .OnLButtonUp -> ClosetGoblinCharacterWindow.HotbarPageUpProxy
  - CM_ClosetGoblin: .OnLButtonUp -> ClosetGoblinCharacterWindow.HotbarPageDownProxy
  - CM_ClosetGoblin: .OnShown -> ClosetGoblinCharacterWindow.OnShow
  - CM_ClosetGoblin: .OnHidden -> ClosetGoblinCharacterWindow.OnHide

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
