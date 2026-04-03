# Function ClosetGoblin.ActivateSet

- Addon: CM_ClosetGoblin
- Kind: function
- Module: ClosetGoblin
- Local: no
- Source: `C:/Return of Reckoning/Interface/AddOns/WAR_API_Ref/.debug/source-subset/ClosetGoblin/ClosetGoblin.lua:782`

## Parameters

- name

## Aliases

- none

## Calls

| Call | Line | Arguments |
| --- | --- | --- |
| ClosetGoblin.GetSetByName | 783 | name |
| ClosetGoblin.TryRecoverStuckActivation | 786 |  |
| ClosetGoblin.Alert | 787 | cgL["already_activating_set"] |
| ClosetGoblin.Message | 795 | cgL["Equiping_set"], set.name |
| ClosetGoblin.ProcessSetActivation | 796 |  |
| ClosetGoblin.Alert | 798 | cgL["player_in_combat"] |
| ClosetGoblin.Alert | 801 | cgL["setnotfound"], name |

## Event Registrations

| Event | Scope | Handler |
| --- | --- | --- |

## State Writes

- ClosetGoblin.GetSetByName
- ClosetGoblin.activatingSet
- ClosetGoblin.activationLastState
- ClosetGoblin.activationSet
