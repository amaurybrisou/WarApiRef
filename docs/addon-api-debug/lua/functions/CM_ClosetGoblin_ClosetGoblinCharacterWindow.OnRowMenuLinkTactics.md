# Function ClosetGoblinCharacterWindow.OnRowMenuLinkTactics

- Addon: CM_ClosetGoblin
- Kind: handler
- Module: ClosetGoblinCharacterWindow
- Local: no
- Source: `C:/Return of Reckoning/Interface/AddOns/WAR_API_Ref/.debug/source-subset/ClosetGoblin/ClosetGoblinCharacterWindow.lua:200`

## Parameters

- none

## Aliases

- none

## Calls

| Call | Line | Arguments |
| --- | --- | --- |
| ClosetGoblinCharacterWindow.GetCurrentSet | 201 |  |
| towstring | 202 | WindowGetId(SystemData.MouseOverWindow.name)-5 |
| WindowGetId | 202 | SystemData.MouseOverWindow.name |
| ClosetGoblin.LinkTactics | 208 | idContextMenu, set.name |
| ClosetGoblinCharacterWindow.RefreshList | 209 |  |

## Event Registrations

| Event | Scope | Handler |
| --- | --- | --- |

## State Writes

- ClosetGoblinCharacterWindow.GetCurrentSet
