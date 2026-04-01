# LibUnits.Events.FORMATION_MEMBER_UPDATED

- Category: Game Event
- Confidence level: HIGH
- Confidence score: 73/100

## Confidence Assessment

- Level: HIGH

- Score: 73/100

- Rationale: Promoted as HIGH confidence because referenced by generated docs or reference files, called globally with no local definition, reinforced across multiple generated source types.

## Evidence Signals

- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +18 Used in event registration or dispatch: Observed in event-driven engine hooks.
- +10 Referenced from initialization flow: Lifecycle reconstruction references this symbol.
- +25 Referenced by generated docs or reference files: The symbol is reinforced outside a single call page.
- +20 Reinforced across multiple generated source types: Evidence comes from several independent API_Ref source types.
- -20 Only one weak usage site: Evidence is too shallow to trust as platform API.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | Effigy |
| Files seen in | `/workspace/Effigy/States/EffigyStateFormation.lua:34` |
| Namespaces detected | LibUnits |
| Source kinds | event_page, flows, lua_event_registration |
| Example locations | Effigy: Effigy.RegisterEventsForFormation |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 1 |
| Global usage count | 1 |
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
| Weak usage only | yes |
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

- Effigy

## Registrars And Handlers

- RegisterEventHandler
- global
- nil

## Examples

- Effigy: Effigy.RegisterEventsForFormation -> LibUnits.Events.FORMATION_MEMBER_UPDATED -> nil
- Effigy: nil -> RegisterEventHandler(LibUnits.Events.FORMATION_MEMBER_UPDATED, nil)

## Related APIs

- none

## Used With

- none

## Triggered By

- none

## Affects

- none

## Notes

- Only one addon surfaced this event in the current API_Ref corpus.
