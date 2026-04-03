# Function ClosetGoblin.Message

- Addon: CM_ClosetGoblin
- Kind: function
- Module: ClosetGoblin
- Local: no
- Source: `C:/Return of Reckoning/Interface/AddOns/WAR_API_Ref/.debug/source-subset/ClosetGoblin/ClosetGoblin.lua:391`

## Parameters

- txt
- variable1
- variable2

## Aliases

- none

## Calls

| Call | Line | Arguments |
| --- | --- | --- |
| gsub | 394 | L "#1#", variable1 |
| gsub | 396 | L "#2#", variable2 |
| TextLogAddEntry | 399 | "Chat", ClosetGoblin.CHAT_LOG_FILTER, L "[ClosetGoblin]: "..towstring(txt) |
| towstring | 399 | txt |

## Event Registrations

| Event | Scope | Handler |
| --- | --- | --- |

## State Writes

- none
