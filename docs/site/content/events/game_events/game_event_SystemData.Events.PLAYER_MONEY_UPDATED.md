# SystemData.Events.PLAYER_MONEY_UPDATED

- Category: Game Event
- Confidence level: MEDIUM
- Confidence score: 63/100

## Confidence Assessment

- Level: MEDIUM

- Score: 63/100

- Rationale: Promoted as MEDIUM confidence because matches default ui or extracted base ui surface, matches a known engine namespace, used in event registration or dispatch.

## Evidence Signals

- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +18 Used in event registration or dispatch: Observed in event-driven engine hooks.
- -15 Only referenced by one internal path: Structural evidence outside a single internal path is missing.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | Shinies |
| Namespaces detected | SystemData |
| Source kinds | event_page |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 0 |
| Global usage count | 0 |
| Local definition count | 0 |
| Documentation references | 0 |
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

Observed as a shared SystemData runtime event used by 1 addons.

## Handler Pattern

Observed as a runtime event ID routed through RegisterEventHandler-style APIs.

## Payload

- Payload shape is not inferable from addon-api docs alone; treat this as an engine event identifier.

## Seen In

- Shinies

## Registrars And Handlers

- RegisterEventHandler
- ShiniesPostUI.OnPlayerMoneyUpdated

## Examples

- Shinies: ShiniesPostUI.OnPlayerMoneyUpdated -> RegisterEventHandler(SystemData.Events.PLAYER_MONEY_UPDATED, ShiniesPostUI.OnPlayerMoneyUpdated)

## Related APIs

- none

## Used With

- none

## Triggered By

- none

## Affects

- none

## Notes

- Only one addon surfaced this event in the current addon-api corpus.
