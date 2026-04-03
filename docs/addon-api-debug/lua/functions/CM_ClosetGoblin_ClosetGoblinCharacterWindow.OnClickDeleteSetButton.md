# Function ClosetGoblinCharacterWindow.OnClickDeleteSetButton

- Addon: CM_ClosetGoblin
- Kind: handler
- Module: ClosetGoblinCharacterWindow
- Local: no
- Source: `C:/Return of Reckoning/Interface/AddOns/WAR_API_Ref/.debug/source-subset/ClosetGoblin/ClosetGoblinCharacterWindow.lua:216`

## Parameters

- none

## Aliases

- none

## Calls

| Call | Line | Arguments |
| --- | --- | --- |
| ClosetGoblinCharacterWindow.GetCurrentSet | 217 |  |
| GetGuildString | 219 | StringTables.Guild.BUTTON_CONFIRM_YES |
| GetGuildString | 220 | StringTables.Guild.BUTTON_CONFIRM_NO |
| DialogManager.MakeTwoButtonDialog | 222 | cgL["confirm_delete_set"]:gsub(L "#1#",set.name), confirmYes, ClosetGoblinCharacterWindow.OnConfirmDeleteSet, confirmNo, nil |
| cgL.confirm_delete_set:gsub | 223 | L "#1#", set.name |

## Event Registrations

| Event | Scope | Handler |
| --- | --- | --- |

## State Writes

- ClosetGoblinCharacterWindow.GetCurrentSet
- ClosetGoblinCharacterWindow.pendingDeleteSet
