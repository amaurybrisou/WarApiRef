# Function ClosetGoblin.OnSlashCommand

- Addon: CM_ClosetGoblin
- Kind: handler
- Module: ClosetGoblin
- Local: no
- Source: `C:/Return of Reckoning/Interface/AddOns/WAR_API_Ref/.debug/source-subset/ClosetGoblin/ClosetGoblin.lua:130`

## Parameters

- what

## Aliases

- none

## Calls

| Call | Line | Arguments |
| --- | --- | --- |
| string.lower | 131 | what |
| ClosetGoblinCharacterWindow.Show | 133 |  |
| ClosetGoblinCharacterWindow.Hide | 136 |  |
| match | 140 | "activate[ ]?(.*)" |
| ClosetGoblin.ActivateSet | 142 | set |
| match | 146 | "set[ ]?([a-z0-9\-]+)[ ]?(.*)" |
| ClosetGoblin.Message | 151 | cgL["Setting_message_output"]..ClosetGoblin.ToEnabled(flag) |
| ClosetGoblin.ToEnabled | 151 | flag |
| ClosetGoblin.Message | 153 | cgL["Setting_message_output"]..ClosetGoblin.ToEnabled(flag) |
| ClosetGoblin.ToEnabled | 153 | flag |
| ClosetGoblin.Message | 156 | cgL["Setting_tooltips"]..ClosetGoblin.ToEnabled(flag) |
| ClosetGoblin.ToEnabled | 156 | flag |
| ClosetGoblin.Message | 157 | cgL["reloadui"] |
| ClosetGoblin.ZoneChangeEnabled | 162 |  |
| ClosetGoblin.ZoneChangeDisabled | 164 |  |
| ClosetGoblin.Message | 168 | cgL["Setting_remove_to_first"]..ClosetGoblin.ToEnabled(flag) |
| ClosetGoblin.ToEnabled | 168 | flag |
| ClosetGoblin.Message | 173 | L "Talisman checks: "..ClosetGoblin.ToEnabled(flag) |
| ClosetGoblin.ToEnabled | 173 | flag |
| match | 177 | "zone[ ]?([A-Za-z0-9\-]+)[ ]?(.*)" |
| ClosetGoblin.AssociateZoneToSet | 179 | zone, towstring(set) |
| towstring | 179 | set |

## Event Registrations

| Event | Scope | Handler |
| --- | --- | --- |

## State Writes

- ClosetGoblin.settings.checkTalismans
- ClosetGoblin.settings.outputMessages
- ClosetGoblin.settings.removeToFirst
- ClosetGoblin.settings.toolTips
