# Function ClosetGoblinZoneWindow.OnSetZoneRowContextMenu

- Addon: CM_ClosetGoblin
- Kind: handler
- Module: ClosetGoblinZoneWindow
- Local: no
- Source: `C:/Return of Reckoning/Interface/AddOns/WAR_API_Ref/.debug/source-subset/ClosetGoblin/ClosetGoblinZoneWindow.lua:155`

## Parameters

- none

## Aliases

- none

## Calls

| Call | Line | Arguments |
| --- | --- | --- |
| WindowGetId | 156 | SystemData.ActiveWindow.name |
| ClosetGoblinZoneWindow.UpdateSetList | 160 |  |
| EA_Window_ContextMenu.CreateContextMenu | 163 | SystemData.ActiveWindow.name |
| EA_Window_ContextMenu.AddMenuItem | 167 | cgL["Unassociate_with_set"], ClosetGoblinZoneWindow.OnRowMenuUnassoiacteClick, false, true |
| EA_Window_ContextMenu.Finalize | 168 |  |

## Event Registrations

| Event | Scope | Handler |
| --- | --- | --- |

## State Writes

- ClosetGoblinZoneWindow.activeMenuRow
- ClosetGoblinZoneWindow.selectedSetDataIndex
- ClosetGoblinZoneWindowContentsSetList.PopulatorIndices
