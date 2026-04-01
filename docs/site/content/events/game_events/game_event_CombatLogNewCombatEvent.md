# CombatLogNewCombatEvent

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
- +20 Reinforced across multiple generated source types: Evidence comes from several independent API_Ref source types.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | Enemy |
| Files seen in | `/workspace/Enemy/Code/CombatLog/CombatLogEpsWindow.lua:12`, `/workspace/Enemy/Code/CombatLog/CombatLogIDS.lua:55`, `/workspace/Enemy/Code/CombatLog/CombatLogStatsWindow.lua:107`, `/workspace/Enemy/Code/CombatLog/CombatLogTargetDefenseWindow.lua:63` |
| Namespaces detected | CombatLogNewCombatEvent |
| Source kinds | event_page, flows, lua_event_registration |
| Example locations | Enemy: Enemy.CombatLogUI_EpsWindow_Initialize, Enemy: Enemy.CombatLogUI_IDS_OnSettingsChanged, Enemy: Enemy.CombatLogUI_StatsWindow_Open, Enemy: Enemy.CombatLogUI_TargetDefenseWindow_OnSettingsChanged |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 4 |
| Global usage count | 4 |
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
- Enemy.CombatLogUI_EpsWindow_Update
- Enemy.CombatLogUI_StatsWindow_OnNewCombatEvent
- Enemy.CombatLogUI_TargetDefenseWindow_OnCombatLogEvent
- Enemy.CombatLog_IDS_OnNewCombatEvent
- addon

## Examples

- Enemy: Enemy.CombatLogUI_EpsWindow_Initialize -> CombatLogNewCombatEvent -> Enemy.CombatLogUI_EpsWindow_Update
- Enemy: Enemy.CombatLogUI_IDS_OnSettingsChanged -> CombatLogNewCombatEvent -> Enemy.CombatLog_IDS_OnNewCombatEvent
- Enemy: Enemy.CombatLogUI_StatsWindow_Open -> CombatLogNewCombatEvent -> Enemy.CombatLogUI_StatsWindow_OnNewCombatEvent
- Enemy: Enemy.CombatLogUI_TargetDefenseWindow_OnSettingsChanged -> CombatLogNewCombatEvent -> Enemy.CombatLogUI_TargetDefenseWindow_OnCombatLogEvent
- Enemy: Enemy.CombatLogUI_EpsWindow_Update -> Enemy.AddEventHandler(CombatLogNewCombatEvent, Enemy.CombatLogUI_EpsWindow_Update)
- Enemy: Enemy.CombatLogUI_StatsWindow_OnNewCombatEvent -> Enemy.AddEventHandler(CombatLogNewCombatEvent, Enemy.CombatLogUI_StatsWindow_OnNewCombatEvent)

## Related APIs

- none

## Used With

- none

## Triggered By

- none

## Affects

- none

## Notes

- Triggered-by evidence: Enemy:Enemy.CombatLog_OnCombatLog
- Only one addon surfaced this event in the current API_Ref corpus.
