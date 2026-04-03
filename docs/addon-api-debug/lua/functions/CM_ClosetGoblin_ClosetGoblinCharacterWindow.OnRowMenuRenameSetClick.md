# Function ClosetGoblinCharacterWindow.OnRowMenuRenameSetClick

- Addon: CM_ClosetGoblin
- Kind: handler
- Module: ClosetGoblinCharacterWindow
- Local: no
- Source: `C:/Return of Reckoning/Interface/AddOns/WAR_API_Ref/.debug/source-subset/ClosetGoblin/ClosetGoblinCharacterWindow.lua:147`

## Parameters

- none

## Aliases

- none

## Calls

| Call | Line | Arguments |
| --- | --- | --- |
| ClosetGoblin.GetSet | 150 | dataIndex |
| DialogManager.MakeTextEntryDialog | 151 | cgL["New_set_name"], cgL["Enter_set_name"], set.name, ClosetGoblinCharacterWindow.OnSubmitSetRename, nil |

## Event Registrations

| Event | Scope | Handler |
| --- | --- | --- |

## State Writes

- ClosetGoblin.GetSet
- ClosetGoblinCharacterWindow.activeMenuRow
- ClosetGoblinCharacterWindowContentsSetList.PopulatorIndices
