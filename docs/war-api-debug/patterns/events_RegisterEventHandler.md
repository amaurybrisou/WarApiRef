# RegisterEventHandler

- Category: events
- Confidence: MEDIUM

## Description

Observed registering global runtime handlers against shared event identifiers.

## Involved APIs

- [SystemData.Events.ALL_MODULES_INITIALIZED](../systemdata/fields/systemdata_SystemData.Events.ALL_MODULES_INITIALIZED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.LOADING_END](../systemdata/fields/systemdata_SystemData.Events.LOADING_END.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_INVENTORY_SLOT_UPDATED](../systemdata/fields/systemdata_SystemData.Events.PLAYER_INVENTORY_SLOT_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.RELOAD_INTERFACE](../systemdata/fields/systemdata_SystemData.Events.RELOAD_INTERFACE.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.UPDATE_PROCESSED](../systemdata/fields/systemdata_SystemData.Events.UPDATE_PROCESSED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.VISIBLE_EQUIPMENT_UPDATED](../systemdata/fields/systemdata_SystemData.Events.VISIBLE_EQUIPMENT_UPDATED.md) (HIGH 100/100) - SystemData Field
- [EventHandler](../xml/element_types/element_EventHandler.md) (LOW 20/100) - XML Element Type

## Flow Diagram

```text
EventHandler <-> EventHandlers
```

## Example Code

```lua
CM_ClosetGoblin: RegisterEventHandler(SystemData.Events.LOADING_END, "ClosetGoblin.Initialize")
```

## Evidence

- CM_ClosetGoblin: RegisterEventHandler(SystemData.Events.LOADING_END, "ClosetGoblin.Initialize")
- CM_ClosetGoblin: RegisterEventHandler(SystemData.Events.RELOAD_INTERFACE, "ClosetGoblin.Initialize")
- CM_ClosetGoblin: RegisterEventHandler(SystemData.Events.ALL_MODULES_INITIALIZED, "ClosetGoblin.OnAllModulesInitialized")
- CM_ClosetGoblin: RegisterEventHandler(SystemData.Events.VISIBLE_EQUIPMENT_UPDATED, "ClosetGoblin.OnVisibleEquipmentUpdated")
- CM_ClosetGoblin: RegisterEventHandler(SystemData.Events.PLAYER_INVENTORY_SLOT_UPDATED, "ClosetGoblin.OnInventorySlotUpdated")
- CM_ClosetGoblin: RegisterEventHandler(SystemData.Events.UPDATE_PROCESSED, "ClosetGoblin.OnActivationWatchdog")
