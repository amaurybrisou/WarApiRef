# SystemData.Events.BATTLEGROUP_UPDATED

- Category: SystemData Field
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 5 addons

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
- +20 Reinforced across multiple generated source types: Evidence comes from several independent addon-api source types.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | Aura, Enemy, Enemy, MegaphonePlusPlus, LibGroup, MegaphonePlusPlus |
| Files seen in | `/workspace/Aura/Source/AuraEngine.lua:312`, `/workspace/Enemy/Code/Core/Groups/Groups.lua:22`, `/workspace/LibGroup/LibGroup.lua:343`, `/workspace/megaphoneplusplus-1.0.4/MegaphonePlusPlus.lua:115`, `/workspace/megaphoneplusplus-1.0.4/MegaphonePlusPlus.lua:150` |
| Namespaces detected | SystemData |
| Source kinds | event_page, event_registration, flow, lua_call |
| Example locations | AuraEngine.JumpStartEventBasedAuras, Enemy.GroupsInitialize, Enemy.GroupsUpdateType, LibGroup.Initialize, Megaphone.GroupUpdate, Megaphone.Initialize |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 13 |
| Global usage count | 13 |
| Local definition count | 0 |
| Documentation references | 4 |
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

Observed SystemData field used by 5 addons through generated function calls, event pages, or lifecycle evidence.

## Seen In

- Aura
- Enemy
- Enemy, MegaphonePlusPlus
- LibGroup
- MegaphonePlusPlus

## Related APIs

- none

## Used With

- none

## Triggered By

- none

## Affects

- none

## Notes

- Observed in contexts: AuraEngine.JumpStartEventBasedAuras, Enemy.GroupsInitialize, Enemy.GroupsUpdateType, LibGroup.Initialize, Megaphone.GroupUpdate, Megaphone.Initialize
