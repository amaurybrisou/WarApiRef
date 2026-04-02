# GroupsTypeUpdated

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
| Files seen in | `/workspace_addons/Enemy/Code/GroupIcons/GroupIcons.lua:62`, `/workspace_addons/Enemy/Code/UnitFrames/UnitFrames.lua:179` |
| Namespaces detected | GroupsTypeUpdated |
| Source kinds | event_page, lua_event_registration |
| Example locations | Enemy: Enemy._GroupIconsEnabledChanged, Enemy: Enemy._UnitFramesEnabledChanged |
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
- Enemy.GroupIcons_OnGroupsTypeUpdated
- Enemy.UnitFrames_GroupsTypeUpdated
- addon

## Examples

- Enemy: Enemy._GroupIconsEnabledChanged -> GroupsTypeUpdated -> Enemy.GroupIcons_OnGroupsTypeUpdated
- Enemy: Enemy._UnitFramesEnabledChanged -> GroupsTypeUpdated -> Enemy.UnitFrames_GroupsTypeUpdated
- Enemy: Enemy.GroupIcons_OnGroupsTypeUpdated -> Enemy.AddEventHandler(GroupsTypeUpdated, Enemy.GroupIcons_OnGroupsTypeUpdated)
- Enemy: Enemy.UnitFrames_GroupsTypeUpdated -> Enemy.AddEventHandler(GroupsTypeUpdated, Enemy.UnitFrames_GroupsTypeUpdated)

## Related APIs

- none

## Used With

- none

## Triggered By

- none

## Affects

- none

## Notes

- Triggered-by evidence: Enemy:Enemy.GroupsUpdateType
- Only one addon surfaced this event in the current addon-api corpus.
