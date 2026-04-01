# GroupsPlayerUpdated

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
| Files seen in | `/workspace/Enemy/Code/GroupIcons/GroupIcons.lua:62`, `/workspace/Enemy/Code/Guard/Guard.lua:57`, `/workspace/Enemy/Code/UnitFrames/UnitFrames.lua:179` |
| Namespaces detected | GroupsPlayerUpdated |
| Source kinds | event_page, lua_event_registration |
| Example locations | Enemy: Enemy._GroupIconsEnabledChanged, Enemy: Enemy._GuardEnabledChanged, Enemy: Enemy._UnitFramesEnabledChanged |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 3 |
| Global usage count | 3 |
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
- Enemy.GroupIcons_OnGroupsPlayerUpdated
- Enemy.Guard_OnGroupsPlayerUpdated
- Enemy.UnitFrames_OnGroupsPlayerUpdated
- addon

## Examples

- Enemy: Enemy._GroupIconsEnabledChanged -> GroupsPlayerUpdated -> Enemy.GroupIcons_OnGroupsPlayerUpdated
- Enemy: Enemy._GuardEnabledChanged -> GroupsPlayerUpdated -> Enemy.Guard_OnGroupsPlayerUpdated
- Enemy: Enemy._UnitFramesEnabledChanged -> GroupsPlayerUpdated -> Enemy.UnitFrames_OnGroupsPlayerUpdated
- Enemy: Enemy.GroupIcons_OnGroupsPlayerUpdated -> Enemy.AddEventHandler(GroupsPlayerUpdated, Enemy.GroupIcons_OnGroupsPlayerUpdated)
- Enemy: Enemy.Guard_OnGroupsPlayerUpdated -> Enemy.AddEventHandler(GroupsPlayerUpdated, Enemy.Guard_OnGroupsPlayerUpdated)
- Enemy: Enemy.UnitFrames_OnGroupsPlayerUpdated -> Enemy.AddEventHandler(GroupsPlayerUpdated, Enemy.UnitFrames_OnGroupsPlayerUpdated)

## Related APIs

- none

## Used With

- none

## Triggered By

- none

## Affects

- none

## Notes

- Triggered-by evidence: Enemy:Enemy.Groups_OnBattlegroupMemberUpdated, Enemy:Enemy.Groups_OnCurrentPlayerUpdated, Enemy:Enemy.Groups_OnGroupStatusUpdated, Enemy:Enemy.Groups_OnPlayerTargetUpdated, Enemy:Enemy.Groups_OnScenarioPlayerHitsUpdated, Enemy:EnemyPlayer:SetDistant, Enemy:EnemyPlayer:SetLOS
- Only one addon surfaced this event in the current API_Ref corpus.
