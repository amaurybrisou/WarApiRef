# Addon Lifecycle

## Manifest

- Seen in addons: 94

- Description: Addon manifests declare file load order and bootstrap hooks before Lua runtime logic begins.

- Evidence:

- Ace: AceAddon-3.0.lua
  - Ace: AceDB-3.0.lua
  - Ace: AceDBOptions-3.0.lua
  - Ace: AceLocale-3.0.lua
  - Ace: AceTimer-3.0.lua
  - Ace: CallbackHandler-1.0.lua
  - Ace: LibGUI.lua
  - Ace: LibStub.lua
  - AdvancedPetAssist: APAAbilityCandidates.lua
  - AdvancedPetAssist: APAAbilityQueue.lua
  - AdvancedPetAssist: APAConstants.lua
  - AdvancedPetAssist: APACore.lua

## Initialize

- Seen in addons: 92

- Description: Initialization hooks create windows, hydrate settings, and bind runtime callbacks.

- Evidence:

- AdvancedPetAssist: APAFollowTargetHUD
  - AdvancedPetAssist: APAInstantOnlyHUD
  - AdvancedPetAssist: APAKitingHUD
  - AdvancedPetAssist: APAPetTargetHUD
  - AdvancedPetAssist: AdvancedPetAssist.Initialize
  - AdvancedRenownTrainer: AdvancedRenownTrainingWindow
  - AggroMeter: AggroMeter.Initialize
  - AnywhereTrainer: AnywhereTrainer.Initialize
  - AnywhereTrainerAdditions: AnywhereTrainerAdditions.Initialize
  - Aura: AuraAddon.OnInitialize
  - Aura: AuraColorPicker
  - Aura: AuraConfig

## Saved-Variables

- Seen in addons: 67

- Description: Saved-variable roots appear before normal runtime hooks and provide persistent addon state.

- Evidence:

- AdvancedPetAssist: APA_Settings
  - AdvancedRenownTrainer: AdvancedRenownTraining.Presets
  - AggroMeter: AggroMeter.Settings
  - AutoBand: AutoBand.saved
  - BagOMatic: BagOMatic.saved
  - BankArkel: BankArkel.db
  - BuffHead: BuffHead.Settings
  - CM_ClosetGoblin: ClosetGoblin.setData
  - CM_ClosetGoblin: ClosetGoblin.settings
  - CMap: CMapWindow.Settings
  - CMap: CMapWindow.VisSettings
  - Clock: ClockSettings

## Xml

- Seen in addons: 77

- Description: XML windows, templates, and handlers become available as the UI layer is created.

- Evidence:

- AdvancedPetAssist: APAComboAttackBind
  - AdvancedPetAssist: APAComboAutoReattack
  - AdvancedPetAssist: APAComboAutoReattackDelay
  - AdvancedPetAssist: APAComboCastDelay
  - AdvancedPetAssist: APAComboCastOnAcquire
  - AdvancedPetAssist: APAComboCombatExitDelay
  - AdvancedPetAssist: APAComboDebug
  - AdvancedPetAssist: APAComboEnabled
  - AdvancedRenownTrainer: AbilityButtonTemplate
  - AdvancedRenownTrainer: AbilityButtonTemplateSquare
  - AdvancedRenownTrainer: AbilityButtonTemplateSquareFrame
  - AdvancedRenownTrainer: AdvancedRenownTrainingExportWindow

## Runtime

- Seen in addons: 66

- Description: Runtime logic pivots into event-driven updates wired through shared event registration APIs.

- Evidence:

- Ace: SystemData.Events.PLAYER_COMBAT_FLAG_UPDATED
  - Ace: SystemData.Events.UPDATE_PROCESSED
  - Ace: e
  - AdvancedPetAssist: SystemData.Events.LOADING_END
  - AdvancedPetAssist: SystemData.Events.L_BUTTON_DOWN_PROCESSED
  - AdvancedPetAssist: SystemData.Events.M_BUTTON_DOWN_PROCESSED
  - AdvancedPetAssist: SystemData.Events.PLAYER_BEGIN_CAST
  - AdvancedPetAssist: SystemData.Events.PLAYER_COMBAT_FLAG_UPDATED
  - AdvancedPetAssist: SystemData.Events.PLAYER_PET_TARGET_UPDATED
  - AdvancedPetAssist: SystemData.Events.PLAYER_PET_UPDATED
  - AdvancedPetAssist: SystemData.Events.PLAYER_TARGET_UPDATED
  - AdvancedRenownTrainer: SystemData.Events.LOADING_END

## Update

- Seen in addons: 39

- Description: Optional update hooks provide repeated processing after initialization has completed.

- Evidence:

- AdvancedPetAssist: AdvancedPetAssist.OnUpdate
  - AggroMeter: AggroMeter.OnUpdate
  - Aura: AuraEngine.Event_UPDATE_PROCESSED
  - AutoBand: AutoBand.update
  - AutoMark: AutoMark.OnUpdate
  - BagOMatic: BagOMatic.update
  - BuffHead: BuffHead.OnUpdate
  - Cheeseboard: Cheeseboard.OnUpdate
  - Clock: Clock.OnUpdate
  - DAoCBuff: DAoCBuff.UpdateWindow
  - Enemy: Enemy.Update
  - GCDsaver: GCDsaver.OnUpdate

## Shutdown

- Seen in addons: 39

- Description: Shutdown hooks unregister commands or handlers and persist addon-owned state.

- Evidence:

- AdvancedPetAssist: AdvancedPetAssist.Shutdown
  - AggroMeter: AggroMeter.Shutdown
  - AnywhereTrainer: AnywhereTrainer.Shutdown
  - AnywhereTrainerAdditions: AnywhereTrainerAdditions.Shutdown
  - Aura: AuraAddon.OnShutdown
  - Aura: AuraSettings.OnShutdown
  - CM_ClosetGoblin: ClosetGoblin.OnShutdown
  - Crafting Info Tooltip: CraftValueTip.ShutDown
  - DAoCBuff: DAoCBuff.Shutdown
  - DaemonAssist: DaemonAssist.Shutdown
  - DeepSleep: DeepSleep.Shutdown
  - EA_ThreePartBar: ThreePartBar.Shutdown
