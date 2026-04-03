# Function ClosetGoblinCharacterWindow.OnSubmitNewSetName

- Addon: CM_ClosetGoblin
- Kind: handler
- Module: ClosetGoblinCharacterWindow
- Local: no
- Source: `C:/Return of Reckoning/Interface/AddOns/WAR_API_Ref/.debug/source-subset/ClosetGoblin/ClosetGoblinCharacterWindow.lua:234`

## Parameters

- name

## Aliases

- none

## Calls

| Call | Line | Arguments |
| --- | --- | --- |
| ClosetGoblin.AddNewSet | 236 | name |
| ClosetGoblin.Alert | 239 | cgL["Set_name_empty"] |
| ClosetGoblin.Alert | 241 | cgL["Set_name_unique"] |
| DialogManager.MakeTextEntryDialog | 243 | cgL["Set name"], L "Enter set name :", name, ClosetGoblinCharacterWindow.OnSubmitNewSetName, nil |
| ClosetGoblin.UpdateListSets | 245 |  |
| table.getn | 246 | ClosetGoblin.GetProfileSets() |
| ClosetGoblin.GetProfileSets | 246 |  |
| ClosetGoblinCharacterWindow.RefreshList | 247 |  |

## Event Registrations

| Event | Scope | Handler |
| --- | --- | --- |

## State Writes

- ClosetGoblin.AddNewSet
- ClosetGoblinCharacterWindow.selectedSetDataIndex
