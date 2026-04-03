# GroupsPlayerWorldObjectIdUpdated

- Category: Game Event
- Confidence level: MEDIUM
- Confidence score: 63/100

## Confidence Assessment

- Level: MEDIUM

- Score: 63/100

- Rationale: Promoted as MEDIUM confidence because referenced by generated docs or reference files, called globally with no local definition, used in event registration or dispatch.

## Evidence Signals

- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +18 Used in event registration or dispatch: Observed in event-driven engine hooks.
- +25 Referenced by generated docs or reference files: The symbol is reinforced outside a single call page.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | Enemy |
| Files seen in | `/workspace/data/raw/Enemy/Code/Core/ObjectWindows.lua:237`, `/workspace/data/raw/Enemy/Code/Marks/Marks.lua:8` |
| Namespaces detected | GroupsPlayerWorldObjectIdUpdated |
| Source kinds | event_page, lua_event_registration |
| Example locations | Enemy: Enemy.MarksInitialize, Enemy: Enemy.ObjectWindowsInitialize |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 2 |
| Global usage count | 2 |
| Local definition count | 0 |
| Documentation references | 1 |
| Initialization flow references | 0 |
| Known engine namespace | no |
| Default UI presence | no |
| Event binding presence | yes |
| Observed in XML and Lua | no |
| Consistent role | no |
| Consistent arguments | no |
| Consistent returns | no |
| Slash command presence | no |
| Weak usage only | no |
| Project-specific name | no |
| Placeholder or computed name | no |
| Conflicting signatures | no |
| Conflicting roles | no |
| Wrapper likely | no |
| Never outside local graph | no |
| Local helper only | no |

## Description

Observed as a runtime event or event-like identifier used by 1 addons.

## Handler Pattern

Observed as a runtime event ID routed through RegisterEventHandler-style APIs.

## Payload

- Payload not inferable from addon-level documentation alone.

## Seen In

- Enemy

## Registrars And Handlers

- Enemy.AddEventHandler
- addon
- function(player)Enemy._MarksCheckPlayerNewId(player.name,player.worldObjectId,true,player.career)end
- function(player)Enemy._ObjectWindowsCheckPlayerNewId(player.name,player.worldObjectId)end

## Examples

- Enemy: Enemy.MarksInitialize -> GroupsPlayerWorldObjectIdUpdated -> function(player)Enemy._MarksCheckPlayerNewId(player.name,player.worldObjectId,true,player.career)end
- Enemy: Enemy.ObjectWindowsInitialize -> GroupsPlayerWorldObjectIdUpdated -> function(player)Enemy._ObjectWindowsCheckPlayerNewId(player.name,player.worldObjectId)end
- Enemy: function(player)Enemy._MarksCheckPlayerNewId(player.name,player.worldObjectId,true,player.career)end -> Enemy.AddEventHandler(GroupsPlayerWorldObjectIdUpdated, function(player)Enemy._MarksCheckPlayerNewId(player.name,player.worldObjectId,true,player.career)end)
- Enemy: function(player)Enemy._ObjectWindowsCheckPlayerNewId(player.name,player.worldObjectId)end -> Enemy.AddEventHandler(GroupsPlayerWorldObjectIdUpdated, function(player)Enemy._ObjectWindowsCheckPlayerNewId(player.name,player.worldObjectId)end)

## Related APIs

- none

## Used With

- none

## Triggered By

- none

## Affects

- none

## Notes

- Triggered-by evidence: Enemy:EnemyPlayer:SetObjectWorldId
- Only one addon surfaced this event in the current addon-api corpus.
