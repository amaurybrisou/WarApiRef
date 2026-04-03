# SystemData.Events.PLAYER_PET_UPDATED

- Category: SystemData Field
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 11 addons

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 168

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, seen in 4 or more addons, matches a known engine namespace.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +18 Used in event registration or dispatch: Observed in event-driven engine hooks.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +25 Referenced by generated docs or reference files: The symbol is reinforced outside a single call page.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | BBars - Mechanic Only, Effigy, HealGrid, LoyalPet, NaturalLog, PetFixWindow, Phantom, Pure |
| Files seen in | BBars.lua, Core.lua, Handler.lua, LoyalPet.lua, PetFixWindow.lua, PhantomLib/PhantomLib.lua, Rangechecker.lua, Source/PurePlayerPet.lua |
| Namespaces detected | SystemData |
| Source kinds | lua_call |
| Example locations | EnforceWindowStates, Init, Initialize, LoadUnitFrame, OnInit, OnUpdate |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 13 |
| Global usage count | 13 |
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

SystemData.SystemData.Events.PLAYER_PET_UPDATED field accessed by 11 addons; commonly found in EnforceWindowStates and Init, Initialize, LoadUnitFrame, OnInit, OnUpdate, RegisterEventHandlers, RegisterStateInfoForPets, Shutdown, Unload, UnloadUnitFrame, lua_call, onInit contexts.

## Seen In

- BBars - Mechanic Only
- Effigy
- HealGrid
- LoyalPet
- NaturalLog
- PetFixWindow
- Phantom
- Pure
- Rangechecker
- TastyButtons
- TurretRange

## Related APIs

- [OnUpdate](../../xml/handlers/handler_OnUpdate.md) (HIGH 88/100) - XML Event

## Notes

- Observed in contexts: EnforceWindowStates, Init, Initialize, LoadUnitFrame, OnInit, OnUpdate
