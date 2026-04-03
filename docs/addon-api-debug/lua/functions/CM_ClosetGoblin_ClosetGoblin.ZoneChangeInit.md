# Function ClosetGoblin.ZoneChangeInit

- Addon: CM_ClosetGoblin
- Kind: lifecycle
- Module: ClosetGoblin
- Local: no
- Source: `C:/Return of Reckoning/Interface/AddOns/WAR_API_Ref/.debug/source-subset/ClosetGoblin/ClosetGoblin.lua:1300`

## Parameters

- none

## Aliases

- none

## Calls

| Call | Line | Arguments |
| --- | --- | --- |
| RegisterEventHandler | 1404 | SystemData.Events.PLAYER_ZONE_CHANGED, "ClosetGoblin.PlayerZoneChanged" |
| RegisterEventHandler | 1405 | SystemData.Events.LOADING_BEGIN, "ClosetGoblin.LoadingBegin" |
| RegisterEventHandler | 1406 | SystemData.Events.LOADING_END, "ClosetGoblin.LoadingEnd" |

## Event Registrations

| Event | Scope | Handler |
| --- | --- | --- |
| SystemData.Events.LOADING_BEGIN | global | ClosetGoblin.LoadingBegin |
| SystemData.Events.LOADING_END | global | ClosetGoblin.LoadingEnd |
| SystemData.Events.PLAYER_ZONE_CHANGED | global | ClosetGoblin.PlayerZoneChanged |

## State Writes

- ClosetGoblin.zoneType
