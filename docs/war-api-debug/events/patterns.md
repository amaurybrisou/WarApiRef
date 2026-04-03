# Event Patterns

## RegisterEventHandler

- Confidence: MEDIUM

- Description: Observed registering global runtime handlers against shared event identifiers.

- Evidence:

- CM_ClosetGoblin: RegisterEventHandler(SystemData.Events.LOADING_END, "ClosetGoblin.Initialize")
  - CM_ClosetGoblin: RegisterEventHandler(SystemData.Events.RELOAD_INTERFACE, "ClosetGoblin.Initialize")
  - CM_ClosetGoblin: RegisterEventHandler(SystemData.Events.ALL_MODULES_INITIALIZED, "ClosetGoblin.OnAllModulesInitialized")
  - CM_ClosetGoblin: RegisterEventHandler(SystemData.Events.VISIBLE_EQUIPMENT_UPDATED, "ClosetGoblin.OnVisibleEquipmentUpdated")
  - CM_ClosetGoblin: RegisterEventHandler(SystemData.Events.PLAYER_INVENTORY_SLOT_UPDATED, "ClosetGoblin.OnInventorySlotUpdated")
  - CM_ClosetGoblin: RegisterEventHandler(SystemData.Events.UPDATE_PROCESSED, "ClosetGoblin.OnActivationWatchdog")

## UnregisterEventHandler

- Confidence: MEDIUM

- Description: Observed removing previously registered global runtime handlers.

- Evidence:

- CM_ClosetGoblin: UnregisterEventHandler(SystemData.Events.LOADING_END, "ClosetGoblin.Initialize")
  - CM_ClosetGoblin: UnregisterEventHandler(SystemData.Events.RELOAD_INTERFACE, "ClosetGoblin.Initialize")
  - CM_ClosetGoblin: UnregisterEventHandler(SystemData.Events.PLAYER_INVENTORY_SLOT_UPDATED, "ClosetGoblin.OnInventorySlotUpdated")
  - CM_ClosetGoblin: UnregisterEventHandler(SystemData.Events.UPDATE_PROCESSED, "ClosetGoblin.OnActivationWatchdog")
  - CM_ClosetGoblin: UnregisterEventHandler(SystemData.Events.PLAYER_ZONE_CHANGED, "ClosetGoblin.PlayerZoneChanged")
  - CM_ClosetGoblin: UnregisterEventHandler(SystemData.Events.LOADING_BEGIN, "ClosetGoblin.LoadingBegin")
