# Event Patterns

## RegisterEventHandler

- Confidence: MEDIUM

- Description: Observed registering global runtime handlers against shared event identifiers.

- Evidence:

- TidyChat: RegisterEventHandler(SystemData.Events.LOADING_END, "TidyChat.OnLoad")
  - TidyChat: RegisterEventHandler(SystemData.Events.RELOAD_INTERFACE, "TidyChat.OnLoad")
  - TidyChat: RegisterEventHandler(SystemData.Events.L_BUTTON_UP_PROCESSED, "TidyChat.OnLBU")
  - TidyChat: RegisterEventHandler(chatLogEventId, "TidyChat.OnChatLogUpdated")
  - TidyChat: RegisterEventHandler(combatLogEventId, "TidyChat.OnCombatLogUpdated")
  - TidyChat: RegisterEventHandler(systemLogEventId, "TidyChat.OnSystemLogUpdated")

## UnregisterEventHandler

- Confidence: MEDIUM

- Description: Observed removing previously registered global runtime handlers.

- Evidence:

- TidyRoll: UnregisterEventHandler(SystemData.Events.LOADING_END, "TidyRoll.OnLoad")
  - TidyRoll: UnregisterEventHandler(SystemData.Events.RELOAD_INTERFACE, "TidyRoll.OnLoad")
  - TidyRoll: UnregisterEventHandler(SystemData.Events.INTERACT_SHOW_LOOT_ROLL_DATA, "TidyRoll.UpdateLootRollData")

## Game event fan-in

- Confidence: HIGH

- Description: Observed multiple addons converging on a small set of runtime events such as LOADING_END, RELOAD_INTERFACE, and player update events.

- Evidence:

- SystemData.Events.INTERACT_LOOT_ROLL_FIRST_ITEM (HIGH)
  - SystemData.Events.INTERACT_SHOW_LOOT_ROLL_DATA (HIGH)
  - SystemData.Events.LOADING_END (HIGH)
  - SystemData.Events.L_BUTTON_UP_PROCESSED (HIGH)
  - SystemData.Events.PLAYER_TARGET_UPDATED (HIGH)
  - SystemData.Events.RELOAD_INTERFACE (HIGH)
  - TextLogGetUpdateEventId (MEDIUM)

## Window event hooks

- Confidence: HIGH

- Description: Observed XML and window-scoped handlers using On* hooks to bridge engine UI events into Lua.

- Evidence:

- OnHidden (HIGH)
  - OnInitialize (MEDIUM)
  - OnLButtonDown (HIGH)
  - OnLButtonUp (HIGH)
  - OnMButtonUp (HIGH)
  - OnMouseOver (MEDIUM)
  - OnMouseWheel (HIGH)
  - OnRButtonUp (HIGH)
