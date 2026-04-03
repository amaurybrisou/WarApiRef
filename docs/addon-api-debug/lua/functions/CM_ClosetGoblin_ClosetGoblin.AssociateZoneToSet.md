# Function ClosetGoblin.AssociateZoneToSet

- Addon: CM_ClosetGoblin
- Kind: function
- Module: ClosetGoblin
- Local: no
- Source: `C:/Return of Reckoning/Interface/AddOns/WAR_API_Ref/.debug/source-subset/ClosetGoblin/ClosetGoblin.lua:1429`

## Parameters

- zone
- set

## Aliases

- none

## Calls

| Call | Line | Arguments |
| --- | --- | --- |
| ClosetGoblin.Message | 1432 | cgL["Unassociating_zone_with_set"], towstring(zone) |
| towstring | 1432 | zone |
| ClosetGoblinZoneWindow.RefreshList | 1434 |  |
| ClosetGoblin.GetSetByName | 1436 | set |
| ClosetGoblin.Alert | 1437 | cgL["setnotfound"], set |
| ClosetGoblin.Message | 1439 | cgL["Associating_zone_with_set"], towstring(zone), set |
| towstring | 1439 | zone |
| ClosetGoblinZoneWindow.RefreshList | 1441 |  |
| ClosetGoblin.Message | 1445 | cgL["Setting_zones_to_sets"], ClosetGoblin.ToEnabled(ClosetGoblin.settings.zoneChange) |
| ClosetGoblin.ToEnabled | 1445 | ClosetGoblin.settings.zoneChange |

## Event Registrations

| Event | Scope | Handler |
| --- | --- | --- |

## State Writes

- none
