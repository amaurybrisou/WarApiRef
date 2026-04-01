# SettingsChanged

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
| Files seen in | `/workspace/Enemy/Code/Assist/Assist.lua:7`, `/workspace/Enemy/Code/CombatLog/CombatLog.lua:63`, `/workspace/Enemy/Code/Guard/Guard.lua:6`, `/workspace/Enemy/Code/ScenarioAlerter/ScenarioAlerter.lua:4`, `/workspace/Enemy/Code/ScenarioInfo/ScenarioInfo.lua:21`, `/workspace/Enemy/Code/TalismanAlerter/TalismanAlerter.lua:7`, `/workspace/Enemy/Code/Timer/Timer.lua:7` |
| Namespaces detected | SettingsChanged |
| Source kinds | event_page, flows, lua_event_registration |
| Example locations | Enemy: Enemy.AssistInitialize, Enemy: Enemy.CombatLogInitialize, Enemy: Enemy.GuardInitialize, Enemy: Enemy.ScenarioAlerterInitialize, Enemy: Enemy.ScenarioInfoInitialize, Enemy: Enemy.TalismanAlerterInitialize |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 7 |
| Global usage count | 7 |
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
- Enemy.Assist_OnSettingsChanged
- Enemy.CombatLog_OnSettingsChanged
- Enemy.Guard_OnSettingsChanged
- Enemy.ScenarioAlerter_OnSettingsChanged
- Enemy.ScenarioInfo_OnSettingsChanged
- Enemy.TalismanAlerter_OnSettingsChanged
- Enemy.Timer_OnSettingsChanged
- addon

## Examples

- Enemy: Enemy.AssistInitialize -> SettingsChanged -> Enemy.Assist_OnSettingsChanged
- Enemy: Enemy.CombatLogInitialize -> SettingsChanged -> Enemy.CombatLog_OnSettingsChanged
- Enemy: Enemy.GuardInitialize -> SettingsChanged -> Enemy.Guard_OnSettingsChanged
- Enemy: Enemy.ScenarioAlerterInitialize -> SettingsChanged -> Enemy.ScenarioAlerter_OnSettingsChanged
- Enemy: Enemy.ScenarioInfoInitialize -> SettingsChanged -> Enemy.ScenarioInfo_OnSettingsChanged
- Enemy: Enemy.TalismanAlerterInitialize -> SettingsChanged -> Enemy.TalismanAlerter_OnSettingsChanged

## Related APIs

- none

## Used With

- none

## Triggered By

- none

## Affects

- none

## Notes

- Triggered-by evidence: Enemy:Enemy.UI_ConfigDialog_Hide, Enemy:Enemy._Initialize
- Only one addon surfaced this event in the current addon-api corpus.
