# GameData.PetCommand.ATTACK

- Category: GameData Field
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 1 addons

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 105

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, matches a known engine namespace, referenced by generated docs or reference files.

## Evidence Signals

- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +25 Referenced by generated docs or reference files: The symbol is reinforced outside a single call page.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | AdvancedPetAssist |
| Files seen in | `/workspace/data/raw/AdvancedPetAssist/APAAbilityQueue.lua:334`, `/workspace/data/raw/AdvancedPetAssist/APAAbilityQueue.lua:418`, `/workspace/data/raw/AdvancedPetAssist/APAAbilityQueue.lua:462`, `/workspace/data/raw/AdvancedPetAssist/APAEvents.lua:165`, `/workspace/data/raw/AdvancedPetAssist/APAUpdateKiting.lua:137`, `/workspace/data/raw/AdvancedPetAssist/APAUpdateKiting.lua:209`, `/workspace/data/raw/AdvancedPetAssist/APAUpdateKiting.lua:61`, `/workspace/data/raw/AdvancedPetAssist/APAUpdateKiting.lua:81` |
| Namespaces detected | GameData |
| Source kinds | lua_call |
| Example locations | AdvancedPetAssist.OnPetUpdated, AdvancedPetAssist.QueueCastNext, AdvancedPetAssist.SendAttackCommands, AdvancedPetAssist.TickAbilityQueue, AdvancedPetAssist.local.HandleKitingInterrupt, AdvancedPetAssist.local.InitializeKitingCycle |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 13 |
| Global usage count | 13 |
| Local definition count | 0 |
| Documentation references | 1 |
| Initialization flow references | 0 |
| Known engine namespace | yes |
| Default UI presence | yes |
| Event binding presence | no |
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

Observed GameData field used by 1 addons through generated function calls, event pages, or lifecycle evidence.

## Seen In

- AdvancedPetAssist

## Related APIs

- none

## Used With

- none

## Triggered By

- none

## Affects

- none

## Notes

- Observed in contexts: AdvancedPetAssist.OnPetUpdated, AdvancedPetAssist.QueueCastNext, AdvancedPetAssist.SendAttackCommands, AdvancedPetAssist.TickAbilityQueue, AdvancedPetAssist.local.HandleKitingInterrupt, AdvancedPetAssist.local.InitializeKitingCycle
