# Function ClosetGoblin.ProcessSetActivation

- Addon: CM_ClosetGoblin
- Kind: function
- Module: ClosetGoblin
- Local: no
- Source: `C:/Return of Reckoning/Interface/AddOns/WAR_API_Ref/.debug/source-subset/ClosetGoblin/ClosetGoblin.lua:729`

## Parameters

- none

## Aliases

- none

## Calls

| Call | Line | Arguments |
| --- | --- | --- |
| ClosetGoblin.ResetActivationState | 732 |  |
| ClosetGoblin.DestroyItemList | 737 |  |
| ClosetGoblin.ResetActivationState | 738 |  |
| ClosetGoblin.Alert | 739 | cgL["player_in_combat"] |
| ClosetGoblin.BuildItemList | 743 |  |
| ClosetGoblin.GetEquipmentStateTokenFromList | 744 | ClosetGoblin.itemList[Cursor.SOURCE_EQUIPMENT] |
| ClosetGoblin.TryActivateSetStep | 745 | set |
| ClosetGoblin.DestroyItemList | 748 |  |
| ClosetGoblin.DestroyItemList | 752 |  |
| ClosetGoblin.FinalizeSetActivation | 753 | set |

## Event Registrations

| Event | Scope | Handler |
| --- | --- | --- |

## State Writes

- ClosetGoblin.GetEquipmentStateTokenFromList
- ClosetGoblin.activationLastState
- ClosetGoblin.activationSet
- ClosetGoblin.activationStallChecks
