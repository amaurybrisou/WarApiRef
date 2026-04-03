# GameData.Player.Pet.name

- Category: GameData Field
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 6 addons

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
| Addons seen in | Effigy, Enemy, FixGit, GuardLine, LibGuard, Pure |
| Files seen in | Code/Core/Groups/EnemyPlayer.lua, FixPet.lua, GuardLine.lua, Source/LibGuard.lua, Source/PurePlayerPet.lua, States/EffigyStatePets.lua |
| Namespaces detected | GameData |
| Source kinds | lua_call |
| Example locations | GetIDs, Initialize, LoadPetFromCurrentPlayer, UpdateInformation, UpdatePet, UpdatePetHealthProxy |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 9 |
| Global usage count | 9 |
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

GameData.GameData.Player.Pet.name field accessed by 6 addons; commonly found in GetIDs and Initialize, LoadPetFromCurrentPlayer, UpdateInformation, UpdatePet, UpdatePetHealthProxy, UpdatePlayerPet, UpdateStateMachine, lua_call contexts.

## Seen In

- Effigy
- Enemy
- FixGit
- GuardLine
- LibGuard
- Pure

## Related APIs

- [PartyUtils.IsPartyActive](../../globals/functions/global_PartyUtils.IsPartyActive.md) (HIGH 100/100) - Global Function

## Notes

- Observed in contexts: GetIDs, Initialize, LoadPetFromCurrentPlayer, UpdateInformation, UpdatePet, UpdatePetHealthProxy
