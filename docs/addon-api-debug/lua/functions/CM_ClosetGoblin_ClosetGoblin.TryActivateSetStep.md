# Function ClosetGoblin.TryActivateSetStep

- Addon: CM_ClosetGoblin
- Kind: function
- Module: ClosetGoblin
- Local: no
- Source: `C:/Return of Reckoning/Interface/AddOns/WAR_API_Ref/.debug/source-subset/ClosetGoblin/ClosetGoblin.lua:639`

## Parameters

- set

## Aliases

- none

## Calls

| Call | Line | Arguments |
| --- | --- | --- |
| ClosetGoblin.RequestRemoveItem | 641 | Cursor.SOURCE_EQUIPMENT, 2 |
| ClosetGoblin.FindBestMatchingItem | 687 |  |
| ClosetGoblin.RequestActivationMove | 689 | matchingItem.list, matchingItem.slot, Cursor.SOURCE_EQUIPMENT, k |
| ClosetGoblin.RequestRemoveItem | 693 | Cursor.SOURCE_EQUIPMENT, k |

## Event Registrations

| Event | Scope | Handler |
| --- | --- | --- |

## State Writes

- ClosetGoblin.FindBestMatchingItem
