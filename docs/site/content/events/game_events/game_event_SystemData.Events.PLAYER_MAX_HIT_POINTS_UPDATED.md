# SystemData.Events.PLAYER_MAX_HIT_POINTS_UPDATED

- Category: Game Event
- Confidence level: HIGH
- Confidence score: 100/100

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 123

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, matches a known engine namespace, referenced by generated docs or reference files.

## Evidence Signals

- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +18 Used in event registration or dispatch: Observed in event-driven engine hooks.
- +25 Referenced by generated docs or reference files: The symbol is reinforced outside a single call page.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | Squared |
| Files seen in | Squared.lua |
| Namespaces detected | SystemData |
| Source kinds | event_page, lua_event_registration |
| Example locations | Squared: RegisterEventHandler |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 2 |
| Global usage count | 2 |
| Local definition count | 0 |
| Documentation references | 1 |
| Initialization flow references | 0 |
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

Runtime event with 2 handler registrations observed across 1 addons.

## Handler Pattern

Observed as a runtime event ID routed through RegisterEventHandler-style APIs.

## Payload

- Payload shape is not inferable from contract artifacts alone; treat this as an engine event identifier.

## Seen In

- Squared

## Registrars And Handlers

- RegisterEventHandler
- SquaredGroup.UpdateMaxHP
- SquaredPlayer.UpdateMaxHP
- global

## Examples

- Squared: RegisterEventHandler -> SystemData.Events.PLAYER_MAX_HIT_POINTS_UPDATED -> SquaredGroup.UpdateMaxHP
- Squared: RegisterEventHandler -> SystemData.Events.PLAYER_MAX_HIT_POINTS_UPDATED -> SquaredPlayer.UpdateMaxHP
- Squared: SquaredGroup.UpdateMaxHP -> RegisterEventHandler(SystemData.Events.PLAYER_MAX_HIT_POINTS_UPDATED, SquaredGroup.UpdateMaxHP)
- Squared: SquaredPlayer.UpdateMaxHP -> RegisterEventHandler(SystemData.Events.PLAYER_MAX_HIT_POINTS_UPDATED, SquaredPlayer.UpdateMaxHP)

## Related APIs

- [RegisterEventHandler](../../globals/functions/global_RegisterEventHandler.md) (MEDIUM 68/100) - Global Function

## Notes

- Only one addon surfaced this event in the current addon-api corpus.
