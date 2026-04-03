# Event Patterns

## RegisterEventHandler

- Confidence: HIGH

- Description: Observed registering global runtime handlers against shared event identifiers.

- Evidence:

- Lib RuString: RegisterEventHandler(SystemData.Events.LOADING_END, "LibRuString.OnLoad")
  - Lib RuString: RegisterEventHandler(SystemData.Events.RELOAD_INTERFACE, "LibRuString.OnLoad")
  - PartyCast: RegisterEventHandler(SystemData.Events.PLAYER_START_INTERACT_TIMER, "PartyCast.StartInteract")
  - PartyCast: RegisterEventHandler(SystemData.Events.INTERACT_DONE, "PartyCast.EndCast")
  - PartyCast: RegisterEventHandler(SystemData.Events.PLAYER_BEGIN_CAST, "PartyCast.StartCast")
  - PartyCast: RegisterEventHandler(SystemData.Events.PLAYER_END_CAST, "PartyCast.EndCast")

## UnregisterEventHandler

- Confidence: MEDIUM

- Description: Observed removing previously registered global runtime handlers.

- Evidence:

- PartyCast: UnregisterEventHandler(SystemData.Events.PLAYER_START_INTERACT_TIMER, "PartyCast.StartInteract")
  - PartyCast: UnregisterEventHandler(SystemData.Events.INTERACT_DONE, "PartyCast.EndCast")
  - PartyCast: UnregisterEventHandler(SystemData.Events.PLAYER_BEGIN_CAST, "PartyCast.StartCast")
  - PartyCast: UnregisterEventHandler(SystemData.Events.PLAYER_END_CAST, "PartyCast.EndCast")
  - PartyCast: UnregisterEventHandler(SystemData.Events.PLAYER_CAST_TIMER_SETBACK, "PartyCast.SetbackCast")
  - PartyCast: UnregisterEventHandler(SystemData.Events.PLAYER_DEATH, "PartyCast.ON_DEATH")

## BroadcastEvent

- Confidence: MEDIUM

- Description: Observed triggering a runtime event so existing handlers are notified.

- Evidence:

- Lib RuString: BroadcastEvent(SystemData.Events.RELOAD_INTERFACE)

## Game event fan-in

- Confidence: HIGH

- Description: Observed multiple addons converging on a small set of runtime events such as LOADING_END, RELOAD_INTERFACE, and player update events.

- Evidence:

- SystemData.Events.ENTER_WORLD (HIGH)
  - SystemData.Events.EXIT_GAME (HIGH)
  - SystemData.Events.INTERACT_COMPLETE_QUEST (HIGH)
  - SystemData.Events.INTERACT_DEFAULT (HIGH)
  - SystemData.Events.INTERACT_DONE (HIGH)
  - SystemData.Events.INTERACT_LOOT_ROLL_FIRST_ITEM (HIGH)
  - SystemData.Events.INTERACT_SHOW_INFLUENCE_REWARDS (HIGH)
  - SystemData.Events.INTERACT_SHOW_LOOT (HIGH)

## Window event hooks

- Confidence: HIGH

- Description: Observed XML and window-scoped handlers using On* hooks to bridge engine UI events into Lua.

- Evidence:

- OnHidden (HIGH)
  - OnHyperLinkLButtonUp (HIGH)
  - OnHyperLinkRButtonUp (HIGH)
  - OnInitialize (HIGH)
  - OnLButtonDown (HIGH)
  - OnLButtonUp (HIGH)
  - OnMButtonUp (HIGH)
  - OnMouseOver (MEDIUM)
