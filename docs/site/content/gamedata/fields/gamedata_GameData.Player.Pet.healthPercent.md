# GameData.Player.Pet.healthPercent

- Category: GameData Field
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 4 addons

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 150

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, seen in 4 or more addons, matches a known engine namespace.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +25 Referenced by generated docs or reference files: The symbol is reinforced outside a single call page.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | Effigy, GuardList, GuardRange, Pure |
| Files seen in | GuardList.lua, GuardRange.lua, Source/PureGroupPet.lua, Source/PurePlayerPet.lua, States/EffigyStatePets.lua |
| Namespaces detected | GameData |
| Source kinds | lua_call |
| Example locations | OnPlayerPetHealthUpdated, UpdatePlayerPet, UpdateStateMachine, lua_call |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 4 |
| Global usage count | 4 |
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

GameData.GameData.Player.Pet.healthPercent field accessed by 4 addons; commonly found in OnPlayerPetHealthUpdated and UpdatePlayerPet, UpdateStateMachine, lua_call contexts.

## Seen In

- Effigy
- GuardList
- GuardRange
- Pure

## Notes

- Observed in contexts: OnPlayerPetHealthUpdated, UpdatePlayerPet, UpdateStateMachine, lua_call
