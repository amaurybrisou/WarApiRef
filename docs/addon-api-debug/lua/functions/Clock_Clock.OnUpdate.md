# Function Clock.OnUpdate

- Addon: Clock
- Kind: handler
- Module: Clock
- Local: no
- Source: `C:/Return of Reckoning/Interface/AddOns/WAR_API_Ref/.debug/source-subset/Clock/Clock.lua:31`

## Parameters

- elapsedTime

## Aliases

- none

## Calls

| Call | Line | Arguments |
| --- | --- | --- |
| GetComputerTime | 33 |  |
| LabelSetText | 51 | "ClockWindowText", wstring.format(fmt,12,Clock.min,Clock.sec) |
| wstring.format | 51 | fmt, 12, Clock.min, Clock.sec |
| LabelSetText | 53 | "ClockWindowText", wstring.format(fmt,Clock.hour%ClockSettings.Hours,Clock.min,Clock.sec) |
| wstring.format | 53 | fmt, Clock.hour%ClockSettings.Hours, Clock.min, Clock.sec |
| LabelSetTextColor | 55 | "ClockWindowText", ClockSettings.R, ClockSettings.G, ClockSettings.B |

## Event Registrations

| Event | Scope | Handler |
| --- | --- | --- |

## State Writes

- Clock.hour
- Clock.lastUpdate
- ClockSettings.Format
