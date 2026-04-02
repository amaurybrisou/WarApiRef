# ConfigDialogInitializeSections

- Category: Game Event
- Confidence level: HIGH
- Confidence score: 93/100

## Confidence Assessment

- Level: HIGH

- Score: 93/100

- Rationale: Promoted as HIGH confidence because referenced by generated docs or reference files, called globally with no local definition, reinforced across multiple generated source types.

## Evidence Signals

- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +18 Used in event registration or dispatch: Observed in event-driven engine hooks.
- +10 Referenced from initialization flow: Lifecycle reconstruction references this symbol.
- +25 Referenced by generated docs or reference files: The symbol is reinforced outside a single call page.
- +20 Reinforced across multiple generated source types: Evidence comes from several independent addon-api source types.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | Enemy |
| Files seen in | `/workspace_addons/Enemy/Code/Assist/Assist.lua:7`, `/workspace_addons/Enemy/Code/CombatLog/CombatLog.lua:63`, `/workspace_addons/Enemy/Code/Guard/Guard.lua:6`, `/workspace_addons/Enemy/Code/Intercom/Intercom.lua:4`, `/workspace_addons/Enemy/Code/ScenarioAlerter/ScenarioAlerter.lua:4`, `/workspace_addons/Enemy/Code/ScenarioInfo/ScenarioInfo.lua:21`, `/workspace_addons/Enemy/Code/TalismanAlerter/TalismanAlerter.lua:7`, `/workspace_addons/Enemy/Code/Timer/Timer.lua:7` |
| Namespaces detected | ConfigDialogInitializeSections |
| Source kinds | event_page, flows, lua_event_registration |
| Example locations | Enemy: Enemy.AssistInitialize, Enemy: Enemy.CombatLogInitialize, Enemy: Enemy.GuardInitialize, Enemy: Enemy.IntercomInitialize, Enemy: Enemy.ScenarioAlerterInitialize, Enemy: Enemy.ScenarioInfoInitialize |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 8 |
| Global usage count | 8 |
| Local definition count | 0 |
| Documentation references | 2 |
| Initialization flow references | 1 |
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
- function(sections)table.insert(sections,g.configDlgSection)end
- function(sections)table.insert(sections,{name="Assist",title=L "Assist",templateName="EnemyAssistConfiguration",onInitialize=Enemy.AssistUI_ConfigDialog_OnInitialize,onLoad=Enemy.AssistUI_ConfigDialog_OnLoad,onSave=Enemy.AssistUI_ConfigDialog_OnSave,onReset=Enemy.AssistUI_ConfigDialog_OnReset})end
- function(sections)table.insert(sections,{name="Guard",title=L "Guard",templateName="EnemyGuardConfiguration",onInitialize=Enemy.GuardUI_ConfigDialog_OnInitialize,onLoad=Enemy.GuardUI_ConfigDialog_OnLoad,onSave=Enemy.GuardUI_ConfigDialog_OnSave,onReset=Enemy.GuardUI_ConfigDialog_OnReset})end
- function(sections)table.insert(sections,{name="Intercom",title=L "Intercom channel",templateName="EnemyIntercomConfiguration",onInitialize=Enemy.IntercomUI_ConfigDialog_OnInitialize,onLoad=Enemy.IntercomUI_ConfigDialog_OnLoad,onSave=Enemy.IntercomUI_ConfigDialog_OnSave})end
- function(sections)table.insert(sections,{name="ScenarioAlerter",title=L "Scenarion alerter",templateName="EnemyScenarioAlerterConfiguration",onInitialize=Enemy.ScenarioAlerterUI_ConfigDialog_OnInitialize,onLoad=Enemy.ScenarioAlerterUI_ConfigDialog_OnLoad,onSave=Enemy.ScenarioAlerterUI_ConfigDialog_OnSave})end
- function(sections)table.insert(sections,{name="ScenarioInfo",title=L "Scenario info",templateName="EnemyScenarioInfoConfiguration",onInitialize=Enemy.ScenarioInfoUI_ConfigDialog_OnInitialize,onLoad=Enemy.ScenarioInfoUI_ConfigDialog_OnLoad,onSave=Enemy.ScenarioInfoUI_ConfigDialog_OnSave})end

## Examples

- Enemy: Enemy.AssistInitialize -> ConfigDialogInitializeSections -> function(sections)table.insert(sections,{name="Assist",title=L "Assist",templateName="EnemyAssistConfiguration",onInitialize=Enemy.AssistUI_ConfigDialog_OnInitialize,onLoad=Enemy.AssistUI_ConfigDialog_OnLoad,onSave=Enemy.AssistUI_ConfigDialog_OnSave,onReset=Enemy.AssistUI_ConfigDialog_OnReset})end
- Enemy: Enemy.CombatLogInitialize -> ConfigDialogInitializeSections -> function(sections)table.insert(sections,g.configDlgSection)end
- Enemy: Enemy.GuardInitialize -> ConfigDialogInitializeSections -> function(sections)table.insert(sections,{name="Guard",title=L "Guard",templateName="EnemyGuardConfiguration",onInitialize=Enemy.GuardUI_ConfigDialog_OnInitialize,onLoad=Enemy.GuardUI_ConfigDialog_OnLoad,onSave=Enemy.GuardUI_ConfigDialog_OnSave,onReset=Enemy.GuardUI_ConfigDialog_OnReset})end
- Enemy: Enemy.IntercomInitialize -> ConfigDialogInitializeSections -> function(sections)table.insert(sections,{name="Intercom",title=L "Intercom channel",templateName="EnemyIntercomConfiguration",onInitialize=Enemy.IntercomUI_ConfigDialog_OnInitialize,onLoad=Enemy.IntercomUI_ConfigDialog_OnLoad,onSave=Enemy.IntercomUI_ConfigDialog_OnSave})end
- Enemy: Enemy.ScenarioAlerterInitialize -> ConfigDialogInitializeSections -> function(sections)table.insert(sections,{name="ScenarioAlerter",title=L "Scenarion alerter",templateName="EnemyScenarioAlerterConfiguration",onInitialize=Enemy.ScenarioAlerterUI_ConfigDialog_OnInitialize,onLoad=Enemy.ScenarioAlerterUI_ConfigDialog_OnLoad,onSave=Enemy.ScenarioAlerterUI_ConfigDialog_OnSave})end
- Enemy: Enemy.ScenarioInfoInitialize -> ConfigDialogInitializeSections -> function(sections)table.insert(sections,{name="ScenarioInfo",title=L "Scenario info",templateName="EnemyScenarioInfoConfiguration",onInitialize=Enemy.ScenarioInfoUI_ConfigDialog_OnInitialize,onLoad=Enemy.ScenarioInfoUI_ConfigDialog_OnLoad,onSave=Enemy.ScenarioInfoUI_ConfigDialog_OnSave})end

## Related APIs

- none

## Used With

- none

## Triggered By

- none

## Affects

- none

## Notes

- Triggered-by evidence: Enemy:Enemy.UI_ConfigDialog_Open
- Only one addon surfaced this event in the current addon-api corpus.
