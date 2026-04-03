# Function ClosetGoblinCharacterWindow.HandleDrag

- Addon: CM_ClosetGoblin
- Kind: function
- Module: ClosetGoblinCharacterWindow
- Local: no
- Source: `C:/Return of Reckoning/Interface/AddOns/WAR_API_Ref/.debug/source-subset/ClosetGoblin/ClosetGoblinCharacterWindow.lua:512`

## Parameters

- none

## Aliases

- none

## Calls

| Call | Line | Arguments |
| --- | --- | --- |
| WindowGetId | 513 | SystemData.ActiveWindow.name |
| Cursor.IconOnCursor | 514 |  |
| ClosetGoblinCharacterWindow.FindCursorItem | 515 | Cursor.Data |
| ClosetGoblin.CanEquip | 520 | item, slot |
| ClosetGoblin.SetSlot | 522 | ClosetGoblin.GetSet(ClosetGoblinCharacterWindow.selectedSetDataIndex), slot, Cursor.Data.Source, Cursor.Data.SourceSlot |
| ClosetGoblin.GetSet | 522 | ClosetGoblinCharacterWindow.selectedSetDataIndex |
| ClosetGoblinCharacterWindow.UpdateSlotIcons | 523 |  |
| Tooltips.ClearTooltip | 524 |  |
| ClosetGoblinCharacterWindow.EquipmentMouseOver | 525 |  |
| Cursor.Clear | 528 |  |

## Event Registrations

| Event | Scope | Handler |
| --- | --- | --- |

## State Writes

- none
