# Function ClosetGoblinCharacterWindow.OnInitialize

- Addon: CM_ClosetGoblin
- Kind: handler
- Module: ClosetGoblinCharacterWindow
- Local: no
- Source: `C:/Return of Reckoning/Interface/AddOns/WAR_API_Ref/.debug/source-subset/ClosetGoblin/ClosetGoblinCharacterWindow.lua:38`

## Parameters

- none

## Aliases

- none

## Calls

| Call | Line | Arguments |
| --- | --- | --- |
| WindowSetShowing | 39 | "ClosetGoblinCharacterWindow", false |
| ButtonSetText | 42 | "ClosetGoblinCharacterWindowContentsSortNameButton", cgL["Set_Name"] |
| ButtonSetText | 43 | "ClosetGoblinCharacterWindowContentsSortTacticsButton", cgL["Tactics"] |
| ButtonSetText | 44 | "ClosetGoblinCharacterWindowContentsNewSet", cgL["New_Set"] |
| ButtonSetText | 45 | "ClosetGoblinCharacterWindowContentsDeleteSet", cgL["Delete_Set"] |
| ButtonSetText | 46 | "ClosetGoblinCharacterWindowContentsZoneConfig", cgL["Zone_Configuration"] |
| LabelSetText | 47 | "ClosetGoblinCharacterWindowContentsLabelUseItemRack", L "ItemRack:" |
| LabelSetText | 49 | "ClosetGoblinCharacterWindowTitleBarText", L "ClosetGoblin "..towstring(ClosetGoblin.version) |
| towstring | 49 | ClosetGoblin.version |
| LabelSetText | 54 | "ClosetGoblinCharacterWindowContentsActionBarSettingsLabel", L "Change ActionBar" |
| LabelSetText | 55 | "ClosetGoblinCharacterWindowContentsActionBarSettingsLabelAB1", L "1" |
| LabelSetText | 56 | "ClosetGoblinCharacterWindowContentsActionBarSettingsLabelAB2", L "2" |
| LabelSetText | 57 | "ClosetGoblinCharacterWindowContentsActionBarSettingsLabelAB3", L "3" |
| LabelSetText | 58 | "ClosetGoblinCharacterWindowContentsActionBarSettingsLabelAB4", L "4" |
| LabelSetText | 59 | "ClosetGoblinCharacterWindowContentsActionBarSettingsLabelAB5", L "5" |
| WindowSetShowing | 61 | "ClosetGoblinCharacterWindowContentsSortTacticsUpArrow", false |
| WindowSetShowing | 62 | "ClosetGoblinCharacterWindowContentsSortTacticsDownArrow", false |
| WindowSetShowing | 64 | "ClosetGoblinCharacterWindowContentsEquipmentShowHelm", false |
| WindowSetShowing | 65 | "ClosetGoblinCharacterWindowContentsEquipmentShowCloak", false |
| WindowSetShowing | 66 | "ClosetGoblinCharacterWindowContentsEquipmentShowCloakHeraldry", false |

## Event Registrations

| Event | Scope | Handler |
| --- | --- | --- |

## State Writes

- ClosetGoblinCharacterWindow.show
