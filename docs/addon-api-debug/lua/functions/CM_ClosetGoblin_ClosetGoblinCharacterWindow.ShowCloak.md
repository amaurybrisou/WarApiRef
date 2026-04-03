# Function ClosetGoblinCharacterWindow.ShowCloak

- Addon: CM_ClosetGoblin
- Kind: function
- Module: ClosetGoblinCharacterWindow
- Local: no
- Source: `C:/Return of Reckoning/Interface/AddOns/WAR_API_Ref/.debug/source-subset/ClosetGoblin/ClosetGoblinCharacterWindow.lua:671`

## Parameters

- none

## Aliases

- none

## Calls

| Call | Line | Arguments |
| --- | --- | --- |
| ClosetGoblinCharacterWindow.GetCurrentSet | 674 |  |
| ButtonSetPressedFlag | 676 | "ClosetGoblinCharacterWindowContentsEquipmentShowCloak", setting |
| ButtonSetDisabledFlag | 679 | "ClosetGoblinCharacterWindowContentsEquipmentShowCloakHeraldry", false |
| ButtonSetDisabledFlag | 681 | "ClosetGoblinCharacterWindowContentsEquipmentShowCloakHeraldry", true |
| d | 684 | L "Toggled Show Cloak setting for "..towstring(set.name)..L " to:", setting |
| towstring | 684 | set.name |

## Event Registrations

| Event | Scope | Handler |
| --- | --- | --- |

## State Writes

- ClosetGoblinCharacterWindow.GetCurrentSet
- ClosetGoblinCharacterWindow.GetCurrentSet.showCloak
