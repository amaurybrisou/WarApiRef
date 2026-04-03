# GameData.AbilityType.PET

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
| Addons seen in | AdvancedPetAssist, DaemonAssist, HealGrid, LoyalPet, NerfedButtons, TastyButtons |
| Files seen in | APACommands.lua, DAAbilities.lua, Gui/HealGridGuiTabMouseClick.lua, LoyalPet.lua, Modules/HealGridMouseClick.lua, NerfedMatchMaking.lua, NerfedMemory.lua, TastyButtonsOptions.lua |
| Namespaces detected | GameData |
| Source kinds | lua_call |
| Example locations | ActionChanged, ActionDetailToActionDetailLabel, GetPetAbilitiesById, GetPetAbilitiesByName, OnPlayerAbilitiesListUpdated, OnRButtonUp |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 10 |
| Global usage count | 10 |
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

GameData.GameData.AbilityType.PET field accessed by 6 addons; commonly found in ActionChanged and ActionDetailToActionDetailLabel, GetPetAbilitiesById, GetPetAbilitiesByName, OnPlayerAbilitiesListUpdated, OnRButtonUp, TryPetAbilities, abilities, buildActionDataCache, lua_call contexts.

## Seen In

- AdvancedPetAssist
- DaemonAssist
- HealGrid
- LoyalPet
- NerfedButtons
- TastyButtons

## Related APIs

- [OnRButtonUp](../../xml/handlers/handler_OnRButtonUp.md) (HIGH 88/100) - XML Event

## Notes

- Observed in contexts: ActionChanged, ActionDetailToActionDetailLabel, GetPetAbilitiesById, GetPetAbilitiesByName, OnPlayerAbilitiesListUpdated, OnRButtonUp
