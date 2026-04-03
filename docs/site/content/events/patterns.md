# Event Patterns

## RegisterEventHandler

- Confidence: HIGH

- Description: Observed registering global runtime handlers against shared event identifiers.

- Evidence:

- AbilityAlert: RegisterEventHandler(SystemData.Events.WORLD_OBJ_COMBAT_EVENT, "AbilityAlert.CombatEvent")
  - ActionBarHide: RegisterEventHandler(SystemData.Events.LOADING_END, "ActionBarHide.OnLoad")
  - ActionBarHide: RegisterEventHandler(SystemData.Events.RELOAD_INTERFACE, "ActionBarHide.OnLoad")
  - ActionBarHide: RegisterEventHandler(SystemData.Events.PLAYER_COMBAT_FLAG_UPDATED, "ActionBarHide.Combat")
  - ActionFraction: RegisterEventHandler(SystemData.Events.LOADING_END, "ActionFractionWindow.UpdateVisibility")
  - ActionFraction: RegisterEventHandler(SystemData.Events.ENTER_WORLD, "ActionFraction.Initialize")

## UnregisterEventHandler

- Confidence: HIGH

- Description: Observed removing previously registered global runtime handlers.

- Evidence:

- Ace: UnregisterEventHandler(SystemData.Events.UPDATE_PROCESSED, "AceAddon_OnUpdate_DONOTTOUCH")
  - ActionFraction: UnregisterEventHandler(SystemData.Events.ENTER_WORLD, "ActionFraction.Initialize")
  - ActionFraction: UnregisterEventHandler(SystemData.Events.INTERFACE_RELOADED, "ActionFraction.Initialize")
  - ActionFraction: UnregisterEventHandler(SystemData.Events.LOADING_END, "ActionFractionWindow.UpdateVisibility")
  - AdvancedPetAssist: UnregisterEventHandler(SystemData.Events.LOADING_END, LOADING_END_HANDLER)
  - AdvancedRenownTrainer: UnregisterEventHandler(SystemData.Events.PLAYER_CAREER_CATEGORY_UPDATED, "AdvancedRenownTraining.CreateDataTable")

## BroadcastEvent

- Confidence: HIGH

- Description: Observed triggering a runtime event so existing handlers are notified.

- Evidence:

- AuctionStats: BroadcastEvent(SystemData.Events.EXIT_GAME)
  - AutoFocus: BroadcastEvent(SystemData.Events.SEND_CHAT_TEXT)
  - Brizio's Crappy Computer Medic: BroadcastEvent(SystemData.Events.USER_SETTINGS_CHANGED)
  - Brizio's Crappy Computer Medic: BroadcastEvent(SystemData.Events.USER_SETTINGS_CHANGED)
  - Brizio's Crappy Computer Medic: BroadcastEvent(SystemData.Events.USER_SETTINGS_CHANGED)
  - CNC: BroadcastEvent(SystemData.Events.USER_SETTINGS_CHANGED)

## Game event fan-in

- Confidence: HIGH

- Description: Observed multiple addons converging on a small set of runtime events such as LOADING_END, RELOAD_INTERFACE, and player update events.

- Evidence:

- BroadcastMessageInvite (MEDIUM)
  - CREATED_PREFIX..index (MEDIUM)
  - ChatTextArrived (MEDIUM)
  - ConfigDialogInitializeSections (MEDIUM)
  - Global (MEDIUM)
  - GroupsPlayerCombatUpdated (MEDIUM)
  - GroupsPlayerDistanceUpdated (MEDIUM)
  - GroupsPlayerEffectsUpdated (MEDIUM)

## Window event hooks

- Confidence: HIGH

- Description: Observed XML and window-scoped handlers using On* hooks to bridge engine UI events into Lua.

- Evidence:

- OnLButtonUp (MEDIUM)
