# Function ClosetGoblinCharacterWindow.OnSubmitSetRename

- Addon: CM_ClosetGoblin
- Kind: handler
- Module: ClosetGoblinCharacterWindow
- Local: no
- Source: `C:/Return of Reckoning/Interface/AddOns/WAR_API_Ref/.debug/source-subset/ClosetGoblin/ClosetGoblinCharacterWindow.lua:154`

## Parameters

- name

## Aliases

- none

## Calls

| Call | Line | Arguments |
| --- | --- | --- |
| ClosetGoblin.RenameSet | 157 | ClosetGoblin.GetSet(dataIndex), name |
| ClosetGoblin.GetSet | 157 | dataIndex |
| ClosetGoblin.Alert | 160 | cgL["Set_name_empty"] |
| ClosetGoblin.Alert | 162 | cgL["Set_name_unique"] |
| DialogManager.MakeTextEntryDialog | 164 | cgL["New_set_name"], cgL["Enter_set_name"], name, ClosetGoblinCharacterWindow.OnSubmitSetRename, nil |
| ClosetGoblin.UpdateListSets | 166 |  |
| ClosetGoblinCharacterWindow.RefreshList | 167 |  |

## Event Registrations

| Event | Scope | Handler |
| --- | --- | --- |

## State Writes

- ClosetGoblin.RenameSet
- ClosetGoblinCharacterWindow.activeMenuRow
- ClosetGoblinCharacterWindowContentsSetList.PopulatorIndices
