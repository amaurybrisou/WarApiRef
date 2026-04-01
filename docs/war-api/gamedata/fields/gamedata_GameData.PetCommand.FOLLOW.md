# GameData.PetCommand.FOLLOW

- Category: GameData Field
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 3 addons

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 138

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, matches a known engine namespace, referenced by generated docs or reference files.

## Evidence Signals

- +18 Seen in 2 to 3 addons: Cross-addon spread is present but limited.
- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +25 Referenced by generated docs or reference files: The symbol is reinforced outside a single call page.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | AdvancedPetAssist, DaemonAssist, LoyalPet |
| Files seen in | `/workspace/AdvancedPetAssist/APAAbilityQueue.lua:462`, `/workspace/AdvancedPetAssist/APACore.lua:269`, `/workspace/AdvancedPetAssist/APACore.lua:283`, `/workspace/AdvancedPetAssist/APAEvents.lua:285`, `/workspace/AdvancedPetAssist/APALosDetection.lua:45`, `/workspace/AdvancedPetAssist/APAUpdateKiting.lua:102`, `/workspace/AdvancedPetAssist/APAUpdateKiting.lua:180`, `/workspace/DaemonAssist/DACore.lua:23` |
| Namespaces detected | GameData |
| Source kinds | lua_call |
| Example locations | AdvancedPetAssist.ApplyDefaultStance, AdvancedPetAssist.SetPassiveFollow, AdvancedPetAssist.TickAbilityQueue, AdvancedPetAssist.TickLosDetection, AdvancedPetAssist.local.HandleMountCast, AdvancedPetAssist.local.TickKitingWait |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 14 |
| Global usage count | 14 |
| Local definition count | 0 |
| Documentation references | 1 |
| Initialization flow references | 0 |
| Known engine namespace | yes |
| Default UI presence | yes |
| Event binding presence | no |
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

Observed GameData field used by 3 addons through generated function calls, event pages, or lifecycle evidence.

## Seen In

- AdvancedPetAssist
- DaemonAssist
- LoyalPet

## Related APIs

- [PetWindow.SwitchToPassiveStance](../../globals/functions/global_PetWindow.SwitchToPassiveStance.md) (HIGH 88/100) - Global Function

## Used With

- [PetWindow.SwitchToPassiveStance](../../globals/functions/global_PetWindow.SwitchToPassiveStance.md) (HIGH 88/100) - Global Function

## Triggered By

- none

## Affects

- none

## Notes

- Observed in contexts: AdvancedPetAssist.ApplyDefaultStance, AdvancedPetAssist.SetPassiveFollow, AdvancedPetAssist.TickAbilityQueue, AdvancedPetAssist.TickLosDetection, AdvancedPetAssist.local.HandleMountCast, AdvancedPetAssist.local.TickKitingWait
