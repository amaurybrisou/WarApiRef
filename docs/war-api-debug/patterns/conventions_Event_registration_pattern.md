# Event registration pattern

- Category: conventions
- Confidence: HIGH

## Description

Runtime events are typically wired through RegisterEventHandler or window-scoped variants, with handler names staying module-qualified.

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
RegisterEventHandler: CM_ClosetGoblin: RegisterEventHandler(SystemData.Events.LOADING_END, "ClosetGoblin.Initialize")
```

## Evidence

- RegisterEventHandler: CM_ClosetGoblin: RegisterEventHandler(SystemData.Events.LOADING_END, "ClosetGoblin.Initialize")
- RegisterEventHandler: CM_ClosetGoblin: RegisterEventHandler(SystemData.Events.RELOAD_INTERFACE, "ClosetGoblin.Initialize")
- RegisterEventHandler: CM_ClosetGoblin: RegisterEventHandler(SystemData.Events.ALL_MODULES_INITIALIZED, "ClosetGoblin.OnAllModulesInitialized")
- RegisterEventHandler: CM_ClosetGoblin: RegisterEventHandler(SystemData.Events.VISIBLE_EQUIPMENT_UPDATED, "ClosetGoblin.OnVisibleEquipmentUpdated")
- RegisterEventHandler: CM_ClosetGoblin: RegisterEventHandler(SystemData.Events.PLAYER_INVENTORY_SLOT_UPDATED, "ClosetGoblin.OnInventorySlotUpdated")
- RegisterEventHandler: CM_ClosetGoblin: RegisterEventHandler(SystemData.Events.UPDATE_PROCESSED, "ClosetGoblin.OnActivationWatchdog")
- UnregisterEventHandler: CM_ClosetGoblin: UnregisterEventHandler(SystemData.Events.LOADING_END, "ClosetGoblin.Initialize")
- UnregisterEventHandler: CM_ClosetGoblin: UnregisterEventHandler(SystemData.Events.RELOAD_INTERFACE, "ClosetGoblin.Initialize")
