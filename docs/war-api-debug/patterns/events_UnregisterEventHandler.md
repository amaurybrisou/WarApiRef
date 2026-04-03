# UnregisterEventHandler

- Category: events
- Confidence: MEDIUM

## Description

Observed removing previously registered global runtime handlers.

## Involved APIs

- [SystemData.Events.LOADING_BEGIN](../systemdata/fields/systemdata_SystemData.Events.LOADING_BEGIN.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.LOADING_END](../systemdata/fields/systemdata_SystemData.Events.LOADING_END.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_INVENTORY_SLOT_UPDATED](../systemdata/fields/systemdata_SystemData.Events.PLAYER_INVENTORY_SLOT_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_ZONE_CHANGED](../systemdata/fields/systemdata_SystemData.Events.PLAYER_ZONE_CHANGED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.RELOAD_INTERFACE](../systemdata/fields/systemdata_SystemData.Events.RELOAD_INTERFACE.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.UPDATE_PROCESSED](../systemdata/fields/systemdata_SystemData.Events.UPDATE_PROCESSED.md) (HIGH 100/100) - SystemData Field
- [EventHandler](../xml/element_types/element_EventHandler.md) (LOW 20/100) - XML Element Type

## Flow Diagram

```text
EventHandler <-> EventHandlers
```

## Example Code

```lua
CM_ClosetGoblin: UnregisterEventHandler(SystemData.Events.LOADING_END, "ClosetGoblin.Initialize")
```

## Evidence

- CM_ClosetGoblin: UnregisterEventHandler(SystemData.Events.LOADING_END, "ClosetGoblin.Initialize")
- CM_ClosetGoblin: UnregisterEventHandler(SystemData.Events.RELOAD_INTERFACE, "ClosetGoblin.Initialize")
- CM_ClosetGoblin: UnregisterEventHandler(SystemData.Events.PLAYER_INVENTORY_SLOT_UPDATED, "ClosetGoblin.OnInventorySlotUpdated")
- CM_ClosetGoblin: UnregisterEventHandler(SystemData.Events.UPDATE_PROCESSED, "ClosetGoblin.OnActivationWatchdog")
- CM_ClosetGoblin: UnregisterEventHandler(SystemData.Events.PLAYER_ZONE_CHANGED, "ClosetGoblin.PlayerZoneChanged")
- CM_ClosetGoblin: UnregisterEventHandler(SystemData.Events.LOADING_BEGIN, "ClosetGoblin.LoadingBegin")
