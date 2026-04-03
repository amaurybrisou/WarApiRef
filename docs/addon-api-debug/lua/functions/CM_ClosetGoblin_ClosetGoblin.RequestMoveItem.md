# Function ClosetGoblin.RequestMoveItem

- Addon: CM_ClosetGoblin
- Kind: function
- Module: ClosetGoblin
- Local: no
- Source: `C:/Return of Reckoning/Interface/AddOns/WAR_API_Ref/.debug/source-subset/ClosetGoblin/ClosetGoblin.lua:913`

## Parameters

- source
- sourceSlot
- destination
- destinationSlot

## Aliases

- none

## Calls

| Call | Line | Arguments |
| --- | --- | --- |
| DEBUG | 915 | L "try to move from "..towstring(source)..L " slot "..towstring(sourceSlot)..L " to "..towstring(destination)..L " slot "..towstring(destinationSlot) |
| towstring | 915 | source |
| towstring | 915 | sourceSlot |
| towstring | 915 | destination |
| towstring | 915 | destinationSlot |
| ClosetGoblin.RequestRemoveItem | 921 | source, sourceSlot |
| RequestMoveItem | 922 | Cursor.SOURCE_INVENTORY, ClosetGoblin.removedLastItemSlot, destination, destinationSlot, 1 |
| RequestMoveItem | 925 | source, sourceSlot, destination, destinationSlot, 1 |

## Event Registrations

| Event | Scope | Handler |
| --- | --- | --- |

## State Writes

- ClosetGoblin.itemList
