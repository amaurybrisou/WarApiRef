# Function ClosetGoblin.OnActivationWatchdog

- Addon: CM_ClosetGoblin
- Kind: handler
- Module: ClosetGoblin
- Local: no
- Source: `C:/Return of Reckoning/Interface/AddOns/WAR_API_Ref/.debug/source-subset/ClosetGoblin/ClosetGoblin.lua:222`

## Parameters

- none

## Aliases

- none

## Calls

| Call | Line | Arguments |
| --- | --- | --- |
| ClosetGoblin.GetCurrentEquipmentStateToken | 233 |  |
| ClosetGoblin.ProcessSetActivation | 236 |  |
| ClosetGoblin.DestroyItemList | 242 |  |
| ClosetGoblin.ResetActivationState | 243 |  |
| ClosetGoblin.Alert | 244 | L "Set switch was reset after no progress. Try again." |
| ClosetGoblin.ProcessSetActivation | 249 |  |

## Event Registrations

| Event | Scope | Handler |
| --- | --- | --- |

## State Writes

- ClosetGoblin.GetCurrentEquipmentStateToken
- ClosetGoblin.activationStallChecks
- ClosetGoblin.activationWatchdogTick
