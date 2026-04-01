# SystemData.Events.M_BUTTON_DOWN_PROCESSED

- Category: Game Event
- Confidence level: HIGH
- Confidence score: 100/100

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 153

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, matches a known engine namespace, referenced by generated docs or reference files.

## Evidence Signals

- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +18 Used in event registration or dispatch: Observed in event-driven engine hooks.
- +10 Referenced from initialization flow: Lifecycle reconstruction references this symbol.
- +25 Referenced by generated docs or reference files: The symbol is reinforced outside a single call page.
- +20 Reinforced across multiple generated source types: Evidence comes from several independent API_Ref source types.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | AdvancedPetAssist |
| Files seen in | `/workspace/AdvancedPetAssist/AdvancedPetAssist.lua:60` |
| Namespaces detected | SystemData |
| Source kinds | event_page, flows, lua_event_registration |
| Example locations | AdvancedPetAssist: AdvancedPetAssist.local.RegisterCoreEvents, AdvancedPetAssist: RegisterCoreEvents |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 2 |
| Global usage count | 2 |
| Local definition count | 0 |
| Documentation references | 2 |
| Initialization flow references | 1 |
| Known engine namespace | yes |
| Default UI presence | yes |
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

Observed as a shared SystemData runtime event used by 1 addons.

## Handler Pattern

Observed as a runtime event ID routed through RegisterEventHandler-style APIs.

## Payload

- Payload shape is not inferable from API_Ref alone; treat this as an engine event identifier.

## Seen In

- AdvancedPetAssist

## Registrars And Handlers

- AdvancedPetAssist.OnMButtonDown
- RegisterEventHandler
- global

## Examples

- AdvancedPetAssist: AdvancedPetAssist.local.RegisterCoreEvents -> SystemData.Events.M_BUTTON_DOWN_PROCESSED -> AdvancedPetAssist.OnMButtonDown
- AdvancedPetAssist: RegisterCoreEvents -> SystemData.Events.M_BUTTON_DOWN_PROCESSED -> AdvancedPetAssist.OnMButtonDown
- AdvancedPetAssist: AdvancedPetAssist.OnMButtonDown -> RegisterEventHandler(SystemData.Events.M_BUTTON_DOWN_PROCESSED, AdvancedPetAssist.OnMButtonDown)

## Related APIs

- none

## Used With

- [RegisterEventHandler](../../globals/functions/global_RegisterEventHandler.md) (MEDIUM 68/100) - Global Function

## Triggered By

- [LPET](game_event_LPET.md) (HIGH 93/100) - Game Event

## Affects

- none

## Notes

- Only one addon surfaced this event in the current API_Ref corpus.
