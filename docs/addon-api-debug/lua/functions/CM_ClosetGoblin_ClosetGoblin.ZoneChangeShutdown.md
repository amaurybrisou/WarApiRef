# Function ClosetGoblin.ZoneChangeShutdown

- Addon: CM_ClosetGoblin
- Kind: lifecycle
- Module: ClosetGoblin
- Local: no
- Source: `C:/Return of Reckoning/Interface/AddOns/WAR_API_Ref/.debug/source-subset/ClosetGoblin/ClosetGoblin.lua:1409`

## Parameters

- none

## Aliases

- none

## Calls

| Call | Line | Arguments |
| --- | --- | --- |
| UnregisterEventHandler | 1410 | SystemData.Events.PLAYER_ZONE_CHANGED, "ClosetGoblin.PlayerZoneChanged" |
| UnregisterEventHandler | 1411 | SystemData.Events.LOADING_BEGIN, "ClosetGoblin.LoadingBegin" |
| UnregisterEventHandler | 1412 | SystemData.Events.LOADING_END, "ClosetGoblin.LoadingEnd" |

## Event Registrations

| Event | Scope | Handler |
| --- | --- | --- |

## State Writes

- none
