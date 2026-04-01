# Event Patterns

## RegisterEventHandler

- Confidence: HIGH

- Description: Observed registering global runtime handlers against shared event identifiers.

- Evidence:

- AdvancedPetAssist: RegisterEventHandler(SystemData.Events.LOADING_END, LOADING_END_HANDLER)
  - AdvancedPetAssist: RegisterEventHandler(SystemData.Events.LOADING_END, LOADING_END_HANDLER)
  - AdvancedRenownTrainer: RegisterEventHandler(SystemData.Events.LOADING_END, "AdvancedRenownTraining.OnReload")
  - AdvancedRenownTrainer: RegisterEventHandler(SystemData.Events.PLAYER_CAREER_CATEGORY_UPDATED, "AdvancedRenownTraining.CreateDataTable")
  - AdvancedRenownTrainer: RegisterEventHandler(SystemData.Events.RELOAD_INTERFACE, "AdvancedRenownTraining.OnReload")
  - AdvancedRenownTrainer: RegisterEventHandler(SystemData.Events.PLAYER_CAREER_CATEGORY_UPDATED, "AdvancedRenownTraining.CreateDataTable")

## UnregisterEventHandler

- Confidence: HIGH

- Description: Observed removing previously registered global runtime handlers.

- Evidence:

- Ace: UnregisterEventHandler(SystemData.Events.UPDATE_PROCESSED, "AceAddon_OnUpdate_DONOTTOUCH")
  - AdvancedPetAssist: UnregisterEventHandler(SystemData.Events.LOADING_END, LOADING_END_HANDLER)
  - AdvancedPetAssist: UnregisterEventHandler(SystemData.Events.LOADING_END, LOADING_END_HANDLER)
  - AdvancedRenownTrainer: UnregisterEventHandler(SystemData.Events.PLAYER_CAREER_CATEGORY_UPDATED, "AdvancedRenownTraining.CreateDataTable")
  - AggroMeter: UnregisterEventHandler(TextLogGetUpdateEventId("Chat"), "AggroMeter.OnChatLogUpdated")
  - Aura: UnregisterEventHandler(SystemData.Events.UPDATE_PROCESSED, "TargetInfoFix.APPLY_TARGETINFO_FIX_DONOTTOUCH")

## BroadcastEvent

- Confidence: HIGH

- Description: Observed triggering a runtime event so existing handlers are notified.

- Evidence:

- Effigy: BroadcastEvent(SystemData.Events.TARGET_PET)
  - Effigy: BroadcastEvent(SystemData.Events.GROUP_LEAVE)
  - Effigy: BroadcastEvent(SystemData.Events.GROUP_LEAVE)
  - Enemy: BroadcastEvent(custom_target_event)
  - Enemy: BroadcastEvent(SystemData.Events.SCENARIO_START_UPDATING_PLAYERS_STATS)
  - Enemy: BroadcastEvent(SystemData.Events.SCENARIO_STOP_UPDATING_PLAYERS_STATS)

## Game event fan-in

- Confidence: HIGH

- Description: Observed multiple addons converging on a small set of runtime events such as LOADING_END, RELOAD_INTERFACE, and player update events.

- Evidence:

- BroadcastMessageInvite (HIGH)
  - ChatTextArrived (HIGH)
  - CombatLogNewCombatEvent (HIGH)
  - CombatLogSessionsUpdated (HIGH)
  - CombatLogSettingsChanged (HIGH)
  - ConfigDialogInitializeSections (HIGH)
  - EA_CareerResourceWindow (HIGH)
  - Global (HIGH)

## Window event hooks

- Confidence: HIGH

- Description: Observed XML and window-scoped handlers using On* hooks to bridge engine UI events into Lua.

- Evidence:

- OnActionButtonLButtonDown (MEDIUM)
  - OnActionButtonLButtonUp (MEDIUM)
  - OnActionButtonMouseOver (MEDIUM)
  - OnActionButtonRButtonDown (MEDIUM)
  - OnHidden (HIGH)
  - OnHyperLinkLButtonUp (HIGH)
  - OnHyperLinkRButtonUp (HIGH)
  - OnInitialize (HIGH)
