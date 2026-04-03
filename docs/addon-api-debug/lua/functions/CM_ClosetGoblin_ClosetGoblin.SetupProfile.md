# Function ClosetGoblin.SetupProfile

- Addon: CM_ClosetGoblin
- Kind: function
- Module: ClosetGoblin
- Local: no
- Source: `C:/Return of Reckoning/Interface/AddOns/WAR_API_Ref/.debug/source-subset/ClosetGoblin/ClosetGoblin.lua:413`

## Parameters

- none

## Aliases

- none

## Calls

| Call | Line | Arguments |
| --- | --- | --- |
| GameData.Account.ServerName:lower | 414 |  |
| GameData.Player.name:lower | 416 |  |
| match | 416 | L "([^^]*)" |
| ClosetGoblin.FindServerProfile | 418 | serverName |
| ClosetGoblin.CreateServerProfile | 420 | serverName |
| ClosetGoblin.FindCharacterProfile | 423 | server, name |
| ClosetGoblin.CreateCharacterProfile | 425 | server, name |

## Event Registrations

| Event | Scope | Handler |
| --- | --- | --- |

## State Writes

- ClosetGoblin.FindCharacterProfile
- ClosetGoblin.FindServerProfile
- ClosetGoblin.currentProfile
- GameData.Account.ServerName:lower
- GameData.Player.name:lower
