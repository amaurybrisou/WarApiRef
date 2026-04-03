# Function ClosetGoblin.FinalizeSetActivation

- Addon: CM_ClosetGoblin
- Kind: function
- Module: ClosetGoblin
- Local: no
- Source: `C:/Return of Reckoning/Interface/AddOns/WAR_API_Ref/.debug/source-subset/ClosetGoblin/ClosetGoblin.lua:700`

## Parameters

- set

## Aliases

- none

## Calls

| Call | Line | Arguments |
| --- | --- | --- |
| d | 702 | L "Showing helmet: ", set.showHelm |
| SetEquippedItemVisible | 703 | GameData.EquipSlots.HELM, set.showHelm |
| d | 706 | L "Showing cloak: ", set.showCloak |
| SetEquippedItemVisible | 707 | GameData.EquipSlots.BACK, set.showCloak |
| d | 710 | L "Showing cloak heraldry: ", set.showCloakHeraldry |
| SendChatText | 711 | L "/togglecloakheraldry", L "" |
| tonumber | 714 | set.tactics |
| TacticsEditor.OnSetMenuSelectionChanged | 716 | setTactics |

## Event Registrations

| Event | Scope | Handler |
| --- | --- | --- |

## State Writes

- none
