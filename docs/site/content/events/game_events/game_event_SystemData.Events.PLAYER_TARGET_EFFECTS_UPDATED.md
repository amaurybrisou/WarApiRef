# SystemData.Events.PLAYER_TARGET_EFFECTS_UPDATED

- Category: Game Event
- Confidence level: HIGH
- Confidence score: 100/100

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 198

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, seen in 4 or more addons, matches a known engine namespace.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +18 Used in event registration or dispatch: Observed in event-driven engine hooks.
- +10 Referenced from initialization flow: Lifecycle reconstruction references this symbol.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +25 Referenced by generated docs or reference files: The symbol is reinforced outside a single call page.
- +20 Reinforced across multiple generated source types: Evidence comes from several independent API_Ref source types.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | BuffHead, DAoCBuff, Enemy, GCDsaver |
| Files seen in | `/workspace/BuffHead/Core.lua:152`, `/workspace/DAoCBuff/Source/DAoCBuff.lua:25`, `/workspace/Enemy/Code/Core/Groups/Groups.lua:22`, `/workspace/GCDsaver/GCDsaver.lua:253` |
| Namespaces detected | SystemData |
| Source kinds | event_page, flows, lua_event_registration |
| Example locations | BuffHead: BuffHead.Initialize, DAoCBuff: DAoCBuff.Initialize, Enemy: Enemy.GroupsInitialize, GCDsaver: GCDsaver.RegisterEvents |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 4 |
| Global usage count | 4 |
| Local definition count | 0 |
| Documentation references | 2 |
| Initialization flow references | 1 |
| Known engine namespace | yes |
| Default UI presence | yes |
| Event binding presence | yes |
| Observed in XML and Lua | no |
| Consistent role | yes |
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

Observed as a shared SystemData runtime event used by 4 addons.

## Handler Pattern

Observed as a runtime event ID routed through RegisterEventHandler-style APIs.

## Payload

- Payload shape is not inferable from API_Ref alone; treat this as an engine event identifier.

## Seen In

- BuffHead
- DAoCBuff
- Enemy
- GCDsaver

## Registrars And Handlers

- BuffHead.OnTargetEffectsUpdated
- DAoCBuff.OnEventH
- Enemy.Groups_OnPlayerTargetEffectsUpdated
- GCDsaver.PLAYER_TARGET_EFFECTS_UPDATED
- RegisterEventHandler
- global

## Examples

- BuffHead: BuffHead.Initialize -> SystemData.Events.PLAYER_TARGET_EFFECTS_UPDATED -> BuffHead.OnTargetEffectsUpdated
- DAoCBuff: DAoCBuff.Initialize -> SystemData.Events.PLAYER_TARGET_EFFECTS_UPDATED -> DAoCBuff.OnEventH
- Enemy: Enemy.GroupsInitialize -> SystemData.Events.PLAYER_TARGET_EFFECTS_UPDATED -> Enemy.Groups_OnPlayerTargetEffectsUpdated
- GCDsaver: GCDsaver.RegisterEvents -> SystemData.Events.PLAYER_TARGET_EFFECTS_UPDATED -> GCDsaver.PLAYER_TARGET_EFFECTS_UPDATED
- BuffHead: BuffHead.OnTargetEffectsUpdated -> RegisterEventHandler(SystemData.Events.PLAYER_TARGET_EFFECTS_UPDATED, BuffHead.OnTargetEffectsUpdated)
- DAoCBuff: DAoCBuff.OnEventH -> RegisterEventHandler(SystemData.Events.PLAYER_TARGET_EFFECTS_UPDATED, DAoCBuff.OnEventH)

## Related APIs

- none

## Used With

- none

## Triggered By

- none

## Affects

- none

## Notes

- none
