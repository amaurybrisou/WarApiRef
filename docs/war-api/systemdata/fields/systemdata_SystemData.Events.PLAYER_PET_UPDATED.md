# SystemData.Events.PLAYER_PET_UPDATED

- Category: SystemData Field
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 6 addons

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
| Addons seen in | AdvancedPetAssist, AdvancedPetAssist, PetFixWindow, TurretRange, Effigy, LoyalPet, PetFixWindow, TurretRange |
| Files seen in | `/workspace_addons/AdvancedPetAssist/AdvancedPetAssist.lua:60`, `/workspace_addons/Effigy/States/EffigyStatePets.lua:5`, `/workspace_addons/LoyalPet/LoyalPet.lua:134`, `/workspace_addons/LoyalPet/LoyalPet.lua:203`, `/workspace_addons/LoyalPet/LoyalPet.lua:371`, `/workspace_addons/PetFixWindow/PetFixWindow.lua:3`, `/workspace_addons/TurrentRange/Core.lua:74`, `/workspace_addons/TurrentRange/Core.lua:95` |
| Namespaces detected | SystemData |
| Source kinds | event_page, event_registration, flow, lua_call |
| Example locations | AdvancedPetAssist.OnPetUpdated, Effigy.RegisterStateInfoForPets, LPET.Initialize, LPET.OnUpdate, LPET.Shutdown, PetFixWindow.PetFixes |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 16 |
| Global usage count | 16 |
| Local definition count | 0 |
| Documentation references | 4 |
| Initialization flow references | 3 |
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

Observed SystemData field used by 6 addons through generated function calls, event pages, or lifecycle evidence.

## Seen In

- AdvancedPetAssist
- AdvancedPetAssist, PetFixWindow, TurretRange
- Effigy
- LoyalPet
- PetFixWindow
- TurretRange

## Related APIs

- none

## Used With

- none

## Triggered By

- none

## Affects

- none

## Notes

- Observed in contexts: AdvancedPetAssist.OnPetUpdated, Effigy.RegisterStateInfoForPets, LPET.Initialize, LPET.OnUpdate, LPET.Shutdown, PetFixWindow.PetFixes
