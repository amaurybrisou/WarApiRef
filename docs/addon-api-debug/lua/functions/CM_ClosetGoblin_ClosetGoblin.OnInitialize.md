# Function ClosetGoblin.OnInitialize

- Addon: CM_ClosetGoblin
- Kind: handler
- Module: ClosetGoblin
- Local: no
- Source: `C:/Return of Reckoning/Interface/AddOns/WAR_API_Ref/.debug/source-subset/ClosetGoblin/ClosetGoblin.lua:73`

## Parameters

- none

## Aliases

- none

## Calls

| Call | Line | Arguments |
| --- | --- | --- |
| RegisterEventHandler | 74 | SystemData.Events.LOADING_END, "ClosetGoblin.Initialize" |
| UnregisterEventHandler | 78 | SystemData.Events.RELOAD_INTERFACE, "ClosetGoblin.Initialize" |
| RegisterEventHandler | 80 | SystemData.Events.RELOAD_INTERFACE, "ClosetGoblin.Initialize" |
| RegisterEventHandler | 81 | SystemData.Events.ALL_MODULES_INITIALIZED, "ClosetGoblin.OnAllModulesInitialized" |
| RegisterEventHandler | 82 | SystemData.Events.VISIBLE_EQUIPMENT_UPDATED, "ClosetGoblin.OnVisibleEquipmentUpdated" |
| RegisterEventHandler | 83 | SystemData.Events.PLAYER_INVENTORY_SLOT_UPDATED, "ClosetGoblin.OnInventorySlotUpdated" |
| RegisterEventHandler | 84 | SystemData.Events.UPDATE_PROCESSED, "ClosetGoblin.OnActivationWatchdog" |

## Event Registrations

| Event | Scope | Handler |
| --- | --- | --- |
| SystemData.Events.ALL_MODULES_INITIALIZED | global | ClosetGoblin.OnAllModulesInitialized |
| SystemData.Events.LOADING_END | global | ClosetGoblin.Initialize |
| SystemData.Events.PLAYER_INVENTORY_SLOT_UPDATED | global | ClosetGoblin.OnInventorySlotUpdated |
| SystemData.Events.RELOAD_INTERFACE | global | ClosetGoblin.Initialize |
| SystemData.Events.UPDATE_PROCESSED | global | ClosetGoblin.OnActivationWatchdog |
| SystemData.Events.VISIBLE_EQUIPMENT_UPDATED | global | ClosetGoblin.OnVisibleEquipmentUpdated |

## State Writes

- ClosetGoblin.firstloadingGame
