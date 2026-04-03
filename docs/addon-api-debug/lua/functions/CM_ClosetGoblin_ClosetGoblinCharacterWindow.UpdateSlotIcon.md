# Function ClosetGoblinCharacterWindow.UpdateSlotIcon

- Addon: CM_ClosetGoblin
- Kind: update
- Module: ClosetGoblinCharacterWindow
- Local: no
- Source: `C:/Return of Reckoning/Interface/AddOns/WAR_API_Ref/.debug/source-subset/ClosetGoblin/ClosetGoblinCharacterWindow.lua:384`

## Parameters

- set
- slot
- defaultIcon

## Aliases

- none

## Calls

| Call | Line | Arguments |
| --- | --- | --- |
| ClosetGoblin.FindItem | 390 | set.slots[slot].id |
| GetIconData | 400 | iconNum |
| type | 411 | defaultIcon |
| GetIconData | 412 | defaultIcon |
| DynamicImageSetTexture | 419 | "ClosetGoblinCharacterWindowContentsEquipmentSlot"..slot.."IconBase", texture, x, y |
| WindowSetTintColor | 422 | "ClosetGoblinCharacterWindowContentsEquipmentSlot"..slot, tint.r, tint.g, tint.b |
| LabelSetText | 424 | "ClosetGoblinCharacterWindowContentsEquipmentSlot"..tostring(slot).."Talisman1", L "" |
| tostring | 424 | slot |
| LabelSetText | 425 | "ClosetGoblinCharacterWindowContentsEquipmentSlot"..tostring(slot).."Talisman2", L "" |
| tostring | 425 | slot |

## Event Registrations

| Event | Scope | Handler |
| --- | --- | --- |

## State Writes

- ClosetGoblin.FindItem
- DefaultColor.ZERO_TINT
