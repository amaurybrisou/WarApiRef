# LPET

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
| Addons seen in | LoyalPet |
| Files seen in | `/workspace/LoyalPet/LoyalPet.lua:134` |
| Namespaces detected | LPET |
| Source kinds | event_page, flows, lua_event_registration |
| Example locations | LoyalPet: LPET.Initialize |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 10 |
| Global usage count | 10 |
| Local definition count | 0 |
| Documentation references | 2 |
| Initialization flow references | 8 |
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

- LoyalPet

## Registrars And Handlers

- RegisterEventHandler
- SystemData.Events.L_BUTTON_DOWN_PROCESSED
- SystemData.Events.M_BUTTON_DOWN_PROCESSED
- SystemData.Events.PLAYER_BEGIN_CAST
- SystemData.Events.PLAYER_CAREER_RESOURCE_UPDATED
- SystemData.Events.PLAYER_PET_STATE_UPDATED
- SystemData.Events.PLAYER_PET_TARGET_UPDATED
- SystemData.Events.PLAYER_PET_UPDATED
- SystemData.Events.PLAYER_TARGET_UPDATED
- SystemData.Events.R_BUTTON_DOWN_PROCESSED
- SystemData.Events.WORLD_OBJ_COMBAT_EVENT
- global

## Examples

- LoyalPet: LPET.Initialize -> LPET -> SystemData.Events.L_BUTTON_DOWN_PROCESSED
- LoyalPet: LPET.Initialize -> LPET -> SystemData.Events.M_BUTTON_DOWN_PROCESSED
- LoyalPet: LPET.Initialize -> LPET -> SystemData.Events.PLAYER_BEGIN_CAST
- LoyalPet: LPET.Initialize -> LPET -> SystemData.Events.PLAYER_CAREER_RESOURCE_UPDATED
- LoyalPet: LPET.Initialize -> LPET -> SystemData.Events.PLAYER_PET_STATE_UPDATED
- LoyalPet: LPET.Initialize -> LPET -> SystemData.Events.PLAYER_PET_TARGET_UPDATED

## Related APIs

- [SystemData.Events.L_BUTTON_DOWN_PROCESSED](game_event_SystemData.Events.L_BUTTON_DOWN_PROCESSED.md) (HIGH 100/100) - Game Event
- [SystemData.Events.M_BUTTON_DOWN_PROCESSED](game_event_SystemData.Events.M_BUTTON_DOWN_PROCESSED.md) (HIGH 100/100) - Game Event
- [SystemData.Events.PLAYER_BEGIN_CAST](game_event_SystemData.Events.PLAYER_BEGIN_CAST.md) (HIGH 100/100) - Game Event
- [SystemData.Events.PLAYER_CAREER_RESOURCE_UPDATED](game_event_SystemData.Events.PLAYER_CAREER_RESOURCE_UPDATED.md) (HIGH 100/100) - Game Event
- [SystemData.Events.PLAYER_PET_STATE_UPDATED](game_event_SystemData.Events.PLAYER_PET_STATE_UPDATED.md) (HIGH 100/100) - Game Event
- [SystemData.Events.PLAYER_PET_TARGET_UPDATED](game_event_SystemData.Events.PLAYER_PET_TARGET_UPDATED.md) (HIGH 100/100) - Game Event
- [SystemData.Events.PLAYER_PET_UPDATED](game_event_SystemData.Events.PLAYER_PET_UPDATED.md) (HIGH 100/100) - Game Event
- [SystemData.Events.PLAYER_TARGET_UPDATED](game_event_SystemData.Events.PLAYER_TARGET_UPDATED.md) (HIGH 100/100) - Game Event
- [SystemData.Events.R_BUTTON_DOWN_PROCESSED](game_event_SystemData.Events.R_BUTTON_DOWN_PROCESSED.md) (HIGH 100/100) - Game Event
- [SystemData.Events.WORLD_OBJ_COMBAT_EVENT](game_event_SystemData.Events.WORLD_OBJ_COMBAT_EVENT.md) (HIGH 100/100) - Game Event

## Used With

- none

## Triggered By

- none

## Affects

- none

## Notes

- Only one addon surfaced this event in the current API_Ref corpus.
