# IconCreateContextMenu

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
| Files seen in | `/workspace/Enemy/Code/CombatLog/CombatLogEpsWindow.lua:12`, `/workspace/Enemy/Code/CombatLog/CombatLogStatsWindow.lua:16`, `/workspace/Enemy/Code/Intercom/Intercom.lua:4`, `/workspace/Enemy/Code/Marks/Marks.lua:8`, `/workspace/Enemy/Code/ScenarioInfo/ScenarioInfo.lua:21` |
| Namespaces detected | IconCreateContextMenu |
| Source kinds | event_page, lua_event_registration |
| Example locations | Enemy: Enemy.CombatLogUI_EpsWindow_Initialize, Enemy: Enemy.CombatLogUI_StatsWindow_Initialize, Enemy: Enemy.IntercomInitialize, Enemy: Enemy.MarksInitialize, Enemy: Enemy.ScenarioInfoInitialize |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 5 |
| Global usage count | 5 |
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
- function(data)if(Enemy.CanSendIntercomMessage())then table.insert(data,{text=L "Leave intercom channel '"..g.name..L "'",callback=Enemy.OnLeaveButton})table.insert(data,{text=L "Invite others to your intercom channel",callback=Enemy.IntercomUI_IntercomDialog_OnInviteButton})else table.insert(data,{text=L "Join Party intercom channel",callback=Enemy.IntercomUI_IntercomDialog_OnCreateButton})table.insert(data,{text=L "Join Warband intercom channel",callback=Enemy.IntercomUI_IntercomDialog_OnCreateButton4})table.insert(data,{text=L "Join Guild intercom channel",callback=Enemy.IntercomUI_IntercomDialog_OnCreateButton5})table.insert(data,{text=L "Join Alliance intercom channel",callback=Enemy.IntercomUI_IntercomDialog_OnCreateButton2})table.insert(data,{text=L "Join Scenario intercom channel",callback=Enemy.IntercomUI_IntercomDialog_OnCreateButton3})table.insert(data,{text=L "Join Scenario party intercom channel",callback=Enemy.IntercomUI_IntercomDialog_OnCreateButton7})end end
- function(data)if(Enemy.CombatLogUI_EpsWindow_IsOpen())then table.insert(data,{text=L "Hide combat log DPS window",callback=Enemy.CombatLogUI_EpsWindow_Hide})else table.insert(data,{text=L "Show combat log DPS window",callback=Enemy.CombatLogUI_EpsWindow_Open,disabled=(not Enemy.Settings.combatLogEnabled or not Enemy.Settings.combatLogStatisticsEnabled)})end end
- function(data)if(Enemy.CombatLogUI_StatsWindow_IsOpen())then table.insert(data,{text=L "Hide combat log stats window",callback=Enemy.CombatLogUI_StatsWindow_Hide})else table.insert(data,{text=L "Show combat log stats window",callback=Enemy.CombatLogUI_StatsWindow_Open,disabled=(not Enemy.Settings.combatLogEnabled or not Enemy.Settings.combatLogStatisticsEnabled)})end end
- function(data)table.insert(data,{text=L "",callback=nil})if(Enemy.ScenarioInfoUI_ScenarioInfoDialog_IsOpen())then table.insert(data,{text=L "Hide scenario info",callback=Enemy.ScenarioInfoUI_ScenarioInfoDialog_Hide})else Enemy.ScenarioInfoUpdateData()table.insert(data,{text=L "Show scenario info",callback=Enemy.ScenarioInfoUI_ScenarioInfoDialog_Open,disabled=(not g.isPluginEnabled or#g.players==0)})end end
- function(data)tinsert(data,{text=L ""})if(Enemy.MarksUI_EnemyMarksWindow_IsOpen())then tinsert(data,{text=L "Hide marks window",callback=Enemy.MarksUI_EnemyMarksWindow_Hide})else tinsert(data,{text=L "Show marks window",callback=Enemy.MarksUI_EnemyMarksWindow_Open})end tinsert(data,{text=L "Clear all active marks",callback=function()for _,template in ipairs(g.templates)do template:ClearMarks()end end})end

## Examples

- Enemy: Enemy.CombatLogUI_EpsWindow_Initialize -> IconCreateContextMenu -> function(data)if(Enemy.CombatLogUI_EpsWindow_IsOpen())then table.insert(data,{text=L "Hide combat log DPS window",callback=Enemy.CombatLogUI_EpsWindow_Hide})else table.insert(data,{text=L "Show combat log DPS window",callback=Enemy.CombatLogUI_EpsWindow_Open,disabled=(not Enemy.Settings.combatLogEnabled or not Enemy.Settings.combatLogStatisticsEnabled)})end end
- Enemy: Enemy.CombatLogUI_StatsWindow_Initialize -> IconCreateContextMenu -> function(data)if(Enemy.CombatLogUI_StatsWindow_IsOpen())then table.insert(data,{text=L "Hide combat log stats window",callback=Enemy.CombatLogUI_StatsWindow_Hide})else table.insert(data,{text=L "Show combat log stats window",callback=Enemy.CombatLogUI_StatsWindow_Open,disabled=(not Enemy.Settings.combatLogEnabled or not Enemy.Settings.combatLogStatisticsEnabled)})end end
- Enemy: Enemy.IntercomInitialize -> IconCreateContextMenu -> function(data)if(Enemy.CanSendIntercomMessage())then table.insert(data,{text=L "Leave intercom channel '"..g.name..L "'",callback=Enemy.OnLeaveButton})table.insert(data,{text=L "Invite others to your intercom channel",callback=Enemy.IntercomUI_IntercomDialog_OnInviteButton})else table.insert(data,{text=L "Join Party intercom channel",callback=Enemy.IntercomUI_IntercomDialog_OnCreateButton})table.insert(data,{text=L "Join Warband intercom channel",callback=Enemy.IntercomUI_IntercomDialog_OnCreateButton4})table.insert(data,{text=L "Join Guild intercom channel",callback=Enemy.IntercomUI_IntercomDialog_OnCreateButton5})table.insert(data,{text=L "Join Alliance intercom channel",callback=Enemy.IntercomUI_IntercomDialog_OnCreateButton2})table.insert(data,{text=L "Join Scenario intercom channel",callback=Enemy.IntercomUI_IntercomDialog_OnCreateButton3})table.insert(data,{text=L "Join Scenario party intercom channel",callback=Enemy.IntercomUI_IntercomDialog_OnCreateButton7})end end
- Enemy: Enemy.MarksInitialize -> IconCreateContextMenu -> function(data)tinsert(data,{text=L ""})if(Enemy.MarksUI_EnemyMarksWindow_IsOpen())then tinsert(data,{text=L "Hide marks window",callback=Enemy.MarksUI_EnemyMarksWindow_Hide})else tinsert(data,{text=L "Show marks window",callback=Enemy.MarksUI_EnemyMarksWindow_Open})end tinsert(data,{text=L "Clear all active marks",callback=function()for _,template in ipairs(g.templates)do template:ClearMarks()end end})end
- Enemy: Enemy.ScenarioInfoInitialize -> IconCreateContextMenu -> function(data)table.insert(data,{text=L "",callback=nil})if(Enemy.ScenarioInfoUI_ScenarioInfoDialog_IsOpen())then table.insert(data,{text=L "Hide scenario info",callback=Enemy.ScenarioInfoUI_ScenarioInfoDialog_Hide})else Enemy.ScenarioInfoUpdateData()table.insert(data,{text=L "Show scenario info",callback=Enemy.ScenarioInfoUI_ScenarioInfoDialog_Open,disabled=(not g.isPluginEnabled or#g.players==0)})end end
- Enemy: function(data)if(Enemy.CanSendIntercomMessage())then table.insert(data,{text=L "Leave intercom channel '"..g.name..L "'",callback=Enemy.OnLeaveButton})table.insert(data,{text=L "Invite others to your intercom channel",callback=Enemy.IntercomUI_IntercomDialog_OnInviteButton})else table.insert(data,{text=L "Join Party intercom channel",callback=Enemy.IntercomUI_IntercomDialog_OnCreateButton})table.insert(data,{text=L "Join Warband intercom channel",callback=Enemy.IntercomUI_IntercomDialog_OnCreateButton4})table.insert(data,{text=L "Join Guild intercom channel",callback=Enemy.IntercomUI_IntercomDialog_OnCreateButton5})table.insert(data,{text=L "Join Alliance intercom channel",callback=Enemy.IntercomUI_IntercomDialog_OnCreateButton2})table.insert(data,{text=L "Join Scenario intercom channel",callback=Enemy.IntercomUI_IntercomDialog_OnCreateButton3})table.insert(data,{text=L "Join Scenario party intercom channel",callback=Enemy.IntercomUI_IntercomDialog_OnCreateButton7})end end -> Enemy.AddEventHandler(IconCreateContextMenu, function(data)if(Enemy.CanSendIntercomMessage())then table.insert(data,{text=L "Leave intercom channel '"..g.name..L "'",callback=Enemy.OnLeaveButton})table.insert(data,{text=L "Invite others to your intercom channel",callback=Enemy.IntercomUI_IntercomDialog_OnInviteButton})else table.insert(data,{text=L "Join Party intercom channel",callback=Enemy.IntercomUI_IntercomDialog_OnCreateButton})table.insert(data,{text=L "Join Warband intercom channel",callback=Enemy.IntercomUI_IntercomDialog_OnCreateButton4})table.insert(data,{text=L "Join Guild intercom channel",callback=Enemy.IntercomUI_IntercomDialog_OnCreateButton5})table.insert(data,{text=L "Join Alliance intercom channel",callback=Enemy.IntercomUI_IntercomDialog_OnCreateButton2})table.insert(data,{text=L "Join Scenario intercom channel",callback=Enemy.IntercomUI_IntercomDialog_OnCreateButton3})table.insert(data,{text=L "Join Scenario party intercom channel",callback=Enemy.IntercomUI_IntercomDialog_OnCreateButton7})end end)

## Related APIs

- none

## Used With

- none

## Triggered By

- none

## Affects

- none

## Notes

- Triggered-by evidence: Enemy:Enemy.UI_Icon_OnRButtonUp
- Only one addon surfaced this event in the current API_Ref corpus.
