# Function ClosetGoblin.Alert

- Addon: CM_ClosetGoblin
- Kind: function
- Module: ClosetGoblin
- Local: no
- Source: `C:/Return of Reckoning/Interface/AddOns/WAR_API_Ref/.debug/source-subset/ClosetGoblin/ClosetGoblin.lua:403`

## Parameters

- txt
- variable1
- variable2

## Aliases

- none

## Calls

| Call | Line | Arguments |
| --- | --- | --- |
| gsub | 405 | L "#1#", variable1 |
| gsub | 407 | L "#2#", variable2 |
| TextLogAddEntry | 410 | "Chat", ClosetGoblin.ALERT_LOG_FILTER, L "[ClosetGoblin]: "..towstring(txt) |
| towstring | 410 | txt |

## Event Registrations

| Event | Scope | Handler |
| --- | --- | --- |

## State Writes

- none
