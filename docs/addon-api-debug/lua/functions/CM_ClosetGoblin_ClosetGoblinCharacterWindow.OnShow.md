# Function ClosetGoblinCharacterWindow.OnShow

- Addon: CM_ClosetGoblin
- Kind: handler
- Module: ClosetGoblinCharacterWindow
- Local: no
- Source: `C:/Return of Reckoning/Interface/AddOns/WAR_API_Ref/.debug/source-subset/ClosetGoblin/ClosetGoblinCharacterWindow.lua:69`

## Parameters

- none

## Aliases

- none

## Calls

| Call | Line | Arguments |
| --- | --- | --- |
| WindowUtils.OnShown | 72 | ClosetGoblinCharacterWindow.OnHide, WindowUtils.Cascade.MODE_AUTOMATIC |
| ClosetGoblin.HasSets | 74 |  |
| ButtonSetPressedFlag | 77 | "ClosetGoblinCharacterWindowContentsCheckboxUseItemRack", ClosetGoblin.settings.useItemRackOnCharacterWindow |
| ClosetGoblinCharacterWindow.RefreshList | 78 |  |

## Event Registrations

| Event | Scope | Handler |
| --- | --- | --- |

## State Writes

- ClosetGoblinCharacterWindow.selectedSetDataIndex
- ClosetGoblinCharacterWindow.show
