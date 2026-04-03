# Function ClosetGoblinCharacterWindow.EquipmentMouseOver

- Addon: CM_ClosetGoblin
- Kind: function
- Module: ClosetGoblinCharacterWindow
- Local: no
- Source: `C:/Return of Reckoning/Interface/AddOns/WAR_API_Ref/.debug/source-subset/ClosetGoblin/ClosetGoblinCharacterWindow.lua:297`

## Parameters

- none

## Aliases

- none

## Calls

| Call | Line | Arguments |
| --- | --- | --- |
| WindowGetId | 298 | SystemData.ActiveWindow.name |
| ClosetGoblinCharacterWindow.GetCurrentSet | 299 |  |
| ClosetGoblin.FindItem | 303 | set.slots[slot].id |
| ClosetGoblinCharacterWindow.CreateTextOnlyTooltip | 307 | "ClosetGoblinCharacterWindowContentsEquipmentSlot"..slot, CharacterWindow.EquipmentSlotInfo[slot].name, nil, {r=123,g=172,b=220} |
| Tooltips.CreateItemTooltip | 309 | item, "ClosetGoblinCharacterWindowContentsEquipmentSlot"..slot, Tooltips.ANCHOR_WINDOW_RIGHT, true |
| DataUtils.ItemHasEnhancementTimer | 310 | item |
| Tooltips.SetUpdateCallback | 311 | Tooltips.TemporaryItemEnhancementCallback |

## Event Registrations

| Event | Scope | Handler |
| --- | --- | --- |

## State Writes

- ClosetGoblinCharacterWindow.GetCurrentSet
