# Function ClosetGoblin.Initialize

- Addon: CM_ClosetGoblin
- Kind: lifecycle
- Module: ClosetGoblin
- Local: no
- Source: `C:/Return of Reckoning/Interface/AddOns/WAR_API_Ref/.debug/source-subset/ClosetGoblin/ClosetGoblin.lua:87`

## Parameters

- none

## Aliases

- none

## Calls

| Call | Line | Arguments |
| --- | --- | --- |
| UnregisterEventHandler | 88 | SystemData.Events.LOADING_END, "ClosetGoblin.Initialize" |
| LibToolkit.NewChatFilter | 89 | ClosetGoblin.CHAT_LOG_FILTER, ClosetGoblin.CHAT_LOG_COLOR |
| LibToolkit.NewChatFilter | 90 | ClosetGoblin.ALERT_LOG_FILTER, ClosetGoblin.ALERT_LOG_COLOR |
| DoesWindowExist | 92 | "yClosetGoblinButton" |
| CreateWindow | 93 | "yClosetGoblinButton", true |
| WindowSetParent | 94 | "yClosetGoblinButton", "CharacterWindow" |
| ClosetGoblin.LoadLanguage | 96 |  |
| ClosetGoblin.SetupProfile | 97 |  |
| ClosetGoblin.UpdateSettings | 98 |  |
| d | 99 | L "Initialize" |
| ClosetGoblin.UpdateListSets | 101 |  |
| ClosetGoblinCharacterWindow.OnInitialize | 102 |  |
| ClosetGoblinZoneWindow.OnInitialize | 103 |  |
| ClosetGoblin.ZoneChangeInit | 106 |  |
| ClosetGoblin.Message | 113 | cgL["loaded"], towstring(ClosetGoblin.version) |
| towstring | 113 | ClosetGoblin.version |
| LibSlash.RegisterSlashCmd | 114 | "cg", ClosetGoblin.OnSlashCommand |

## Event Registrations

| Event | Scope | Handler |
| --- | --- | --- |

## State Writes

- ClosetGoblin.CreateItemTooltipOLD
- Tooltips.CreateItemTooltip
