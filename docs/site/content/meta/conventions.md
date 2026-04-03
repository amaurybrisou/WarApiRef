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

- RegisterEventHandler: AbilityAlert: RegisterEventHandler(SystemData.Events.WORLD_OBJ_COMBAT_EVENT, "AbilityAlert.CombatEvent")
  - RegisterEventHandler: ActionBarHide: RegisterEventHandler(SystemData.Events.LOADING_END, "ActionBarHide.OnLoad")
  - RegisterEventHandler: ActionBarHide: RegisterEventHandler(SystemData.Events.RELOAD_INTERFACE, "ActionBarHide.OnLoad")
  - RegisterEventHandler: ActionBarHide: RegisterEventHandler(SystemData.Events.PLAYER_COMBAT_FLAG_UPDATED, "ActionBarHide.Combat")
  - RegisterEventHandler: ActionFraction: RegisterEventHandler(SystemData.Events.LOADING_END, "ActionFractionWindow.UpdateVisibility")
  - RegisterEventHandler: ActionFraction: RegisterEventHandler(SystemData.Events.ENTER_WORLD, "ActionFraction.Initialize")
  - UnregisterEventHandler: Ace: UnregisterEventHandler(SystemData.Events.UPDATE_PROCESSED, "AceAddon_OnUpdate_DONOTTOUCH")
  - UnregisterEventHandler: ActionFraction: UnregisterEventHandler(SystemData.Events.ENTER_WORLD, "ActionFraction.Initialize")

## UI creation pattern

- Confidence: HIGH

- Description: UI is commonly created from XML, then positioned in Lua through CreateWindow or CreateWindowFromTemplate followed by anchor calls.

- Evidence:

- Window creation: AbilityAlert: CreateWindow("AAWindow", true)
  - Window creation: ActionFraction: CreateWindow(windowName, true)
  - Window creation: AdvancedPetAssist: CreateWindow("APAOptions", true)
  - Window creation: AdvancedRenownTrainer: CreateWindow("AdvancedRenownTrainingPresetsWindow", false)
  - Window creation: AdvancedRenownTrainer: CreateWindow(ImportWindowName, false)
  - Window creation: AdvancedRenownTrainer: CreateWindow(ImportNameInputWindowName, false)
  - Template instantiation: Ace: CreateWindowFromTemplate(w.name, base, w.parent)
  - Template instantiation: Ace: CreateWindowFromTemplate(w.name, "EA_Button_DefaultWindowClose", w.parent)

## XML to Lua binding pattern

- Confidence: HIGH

- Description: XML handler names map directly to Lua functions and can be cross-checked through the bindings page.

- Evidence:

- AdvancedPetAssist: .OnLButtonUp -> APAGui.OnTabButtonUp
  - AdvancedPetAssist: .OnLButtonUp -> APAGui.OnTabButtonUp
  - AdvancedPetAssist: .OnLButtonUp -> APAGui.OnTabButtonUp
  - AdvancedPetAssist: .OnLButtonUp -> APAGui.OnTabButtonUp
  - AdvancedPetAssist: .OnLButtonUp -> APAGui.OnTabButtonUp
  - AdvancedPetAssist: .OnLButtonUp -> APAGui.OnTabButtonUp
  - AdvancedPetAssist: .OnSelChanged -> APAGui.OnComboChanged
  - AdvancedPetAssist: .OnSelChanged -> APAGui.OnComboChanged

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
