# Function ClosetGoblin.RequestRemoveItem

- Addon: CM_ClosetGoblin
- Kind: function
- Module: ClosetGoblin
- Local: no
- Source: `C:/Return of Reckoning/Interface/AddOns/WAR_API_Ref/.debug/source-subset/ClosetGoblin/ClosetGoblin.lua:900`

## Parameters

- source
- sourceSlot

## Aliases

- none

## Calls

| Call | Line | Arguments |
| --- | --- | --- |
| ClosetGoblin.FindPlaceInBackpack | 901 |  |
| ClosetGoblin.RequestMoveItem | 903 | source, sourceSlot, Cursor.SOURCE_INVENTORY, destinationSlot |
| ClosetGoblin.Alert | 907 | cgL["bag_full_set_problem"] |

## Event Registrations

| Event | Scope | Handler |
| --- | --- | --- |

## State Writes

- ClosetGoblin.FindPlaceInBackpack
- ClosetGoblin.removedLastItemSlot
