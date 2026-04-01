# Conventions

## Initialization pattern

- Confidence: HIGH

- Description: Addons consistently move from manifest loading into initialize hooks, then into XML creation and runtime event registration.

- Evidence:

- initialize: AdvancedPetAssist: APAFollowTargetHUD, AdvancedPetAssist: APAInstantOnlyHUD
  - runtime: Ace: SystemData.Events.PLAYER_COMBAT_FLAG_UPDATED, Ace: SystemData.Events.UPDATE_PROCESSED
  - xml: AdvancedPetAssist: APAComboAttackBind, AdvancedPetAssist: APAComboAutoReattack

## Event registration pattern

- Confidence: HIGH

- Description: Runtime events are typically wired through RegisterEventHandler or window-scoped variants, with handler names staying module-qualified.

- Evidence:

- RegisterEventHandler: AdvancedPetAssist: RegisterEventHandler(SystemData.Events.LOADING_END, LOADING_END_HANDLER)
  - RegisterEventHandler: AdvancedPetAssist: RegisterEventHandler(SystemData.Events.LOADING_END, LOADING_END_HANDLER)
  - RegisterEventHandler: AdvancedRenownTrainer: RegisterEventHandler(SystemData.Events.LOADING_END, "AdvancedRenownTraining.OnReload")
  - RegisterEventHandler: AdvancedRenownTrainer: RegisterEventHandler(SystemData.Events.PLAYER_CAREER_CATEGORY_UPDATED, "AdvancedRenownTraining.CreateDataTable")
  - RegisterEventHandler: AdvancedRenownTrainer: RegisterEventHandler(SystemData.Events.RELOAD_INTERFACE, "AdvancedRenownTraining.OnReload")
  - RegisterEventHandler: AdvancedRenownTrainer: RegisterEventHandler(SystemData.Events.PLAYER_CAREER_CATEGORY_UPDATED, "AdvancedRenownTraining.CreateDataTable")
  - UnregisterEventHandler: Ace: UnregisterEventHandler(SystemData.Events.UPDATE_PROCESSED, "AceAddon_OnUpdate_DONOTTOUCH")
  - UnregisterEventHandler: AdvancedPetAssist: UnregisterEventHandler(SystemData.Events.LOADING_END, LOADING_END_HANDLER)

## UI creation pattern

- Confidence: HIGH

- Description: UI is commonly created from XML, then positioned in Lua through CreateWindow or CreateWindowFromTemplate followed by anchor calls.

- Evidence:

- Window creation: AdvancedPetAssist: CreateWindow("APAOptions", true)
  - Window creation: AdvancedRenownTrainer: CreateWindow("AdvancedRenownTrainingPresetsWindow", false)
  - Window creation: AdvancedRenownTrainer: CreateWindow(ImportWindowName, false)
  - Window creation: AdvancedRenownTrainer: CreateWindow(ImportNameInputWindowName, false)
  - Window creation: AdvancedRenownTrainer: CreateWindow(ExportWindowName, false)
  - Window creation: AdvancedRenownTrainer: CreateWindow(LinkWindowName, false)
  - Template instantiation: Ace: CreateWindowFromTemplate(w.name, base, w.parent)
  - Template instantiation: Ace: CreateWindowFromTemplate(w.name, base, w.parent)

## XML to Lua binding pattern

- Confidence: HIGH

- Description: XML handler names map directly to Lua functions and can be cross-checked through the bindings page.

- Evidence:

- AdvancedPetAssist: APAComboAttackBind.OnSelChanged -> APAGui.OnComboChanged
  - AdvancedPetAssist: APAComboAutoReattack.OnSelChanged -> APAGui.OnComboChanged
  - AdvancedPetAssist: APAComboAutoReattackDelay.OnSelChanged -> APAGui.OnComboChanged
  - AdvancedPetAssist: APAComboCastDelay.OnSelChanged -> APAGui.OnComboChanged
  - AdvancedPetAssist: APAComboCastOnAcquire.OnSelChanged -> APAGui.OnComboChanged
  - AdvancedPetAssist: APAComboCombatExitDelay.OnSelChanged -> APAGui.OnComboChanged
  - AdvancedPetAssist: APAComboDebug.OnSelChanged -> APAGui.OnComboChanged
  - AdvancedPetAssist: APAComboEnabled.OnSelChanged -> APAGui.OnComboChanged

## State management pattern

- Confidence: MEDIUM

- Description: Persistent state is typically rooted in addon-owned globals and saved variables, then initialized before runtime hooks are attached.

- Evidence:

- AdvancedPetAssist: APA_Settings
  - AdvancedRenownTrainer: AdvancedRenownTraining.Presets
  - AggroMeter: AggroMeter.Settings
  - AutoBand: AutoBand.saved
  - BagOMatic: BagOMatic.saved
  - BankArkel: BankArkel.db
  - BuffHead: BuffHead.Settings
  - CM_ClosetGoblin: ClosetGoblin.setData
