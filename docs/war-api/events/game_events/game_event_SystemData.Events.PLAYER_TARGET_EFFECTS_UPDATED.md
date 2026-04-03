# SystemData.Events.PLAYER_TARGET_EFFECTS_UPDATED

- Category: Game Event
- Confidence level: HIGH
- Confidence score: 100/100

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 156

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, matches a known engine namespace, referenced by generated docs or reference files.

## Evidence Signals

- +18 Seen in 2 to 3 addons: Cross-addon spread is present but limited.
- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +18 Used in event registration or dispatch: Observed in event-driven engine hooks.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +25 Referenced by generated docs or reference files: The symbol is reinforced outside a single call page.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | BuffHead, DAoCBuff, Enemy |
| Files seen in | `/workspace/data/raw/BuffHead/Core.lua:152`, `/workspace/data/raw/DAoCBuff/Source/DAoCBuff.lua:25`, `/workspace/data/raw/Enemy/Code/Core/Groups/Groups.lua:22` |
| Namespaces detected | SystemData |
| Source kinds | event_page, lua_event_registration |
| Example locations | BuffHead: BuffHead.Initialize, DAoCBuff: DAoCBuff.Initialize, Enemy: Enemy.GroupsInitialize |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 3 |
| Global usage count | 3 |
| Local definition count | 0 |
| Documentation references | 1 |
| Initialization flow references | 0 |
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

Observed as a shared SystemData runtime event used by 3 addons.

## Handler Pattern

Observed as a runtime event ID routed through RegisterEventHandler-style APIs.

## Payload

- Payload shape is not inferable from addon-api docs alone; treat this as an engine event identifier.

## Seen In

- BuffHead
- DAoCBuff
- Enemy

## Registrars And Handlers

- BuffHead.OnTargetEffectsUpdated
- DAoCBuff.OnEventH
- Enemy.Groups_OnPlayerTargetEffectsUpdated
- RegisterEventHandler
- global

## Examples

- BuffHead: BuffHead.Initialize -> SystemData.Events.PLAYER_TARGET_EFFECTS_UPDATED -> BuffHead.OnTargetEffectsUpdated
- DAoCBuff: DAoCBuff.Initialize -> SystemData.Events.PLAYER_TARGET_EFFECTS_UPDATED -> DAoCBuff.OnEventH
- Enemy: Enemy.GroupsInitialize -> SystemData.Events.PLAYER_TARGET_EFFECTS_UPDATED -> Enemy.Groups_OnPlayerTargetEffectsUpdated
- BuffHead: BuffHead.OnTargetEffectsUpdated -> RegisterEventHandler(SystemData.Events.PLAYER_TARGET_EFFECTS_UPDATED, BuffHead.OnTargetEffectsUpdated)
- DAoCBuff: DAoCBuff.OnEventH -> RegisterEventHandler(SystemData.Events.PLAYER_TARGET_EFFECTS_UPDATED, DAoCBuff.OnEventH)
- Enemy: Enemy.Groups_OnPlayerTargetEffectsUpdated -> RegisterEventHandler(SystemData.Events.PLAYER_TARGET_EFFECTS_UPDATED, Enemy.Groups_OnPlayerTargetEffectsUpdated)

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
