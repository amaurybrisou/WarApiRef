# SystemData.Events.PLAYER_BEGIN_CAST

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
| Addons seen in | AdvancedPetAssist, Effigy, LibGuard, TurretRange |
| Files seen in | `/workspace/AdvancedPetAssist/AdvancedPetAssist.lua:60`, `/workspace/Effigy/States/EffigyStateCastbar.lua:41`, `/workspace/LibGuard/Source/LibGuard.lua:47`, `/workspace/TurrentRange/Core.lua:74` |
| Namespaces detected | SystemData |
| Source kinds | event_page, flows, lua_event_registration |
| Example locations | AdvancedPetAssist: AdvancedPetAssist.local.RegisterCoreEvents, AdvancedPetAssist: RegisterCoreEvents, Effigy: Effigy.RegisterStateInfoForCastbar, LibGuard: LibGuard.Init, TurretRange: TurretRange.Initialize |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 5 |
| Global usage count | 5 |
| Local definition count | 0 |
| Documentation references | 2 |
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

Observed as a shared SystemData runtime event used by 4 addons.

## Handler Pattern

Observed as a runtime event ID routed through RegisterEventHandler-style APIs.

## Payload

- Payload shape is not inferable from API_Ref alone; treat this as an engine event identifier.

## Seen In

- AdvancedPetAssist
- Effigy
- LibGuard
- TurretRange

## Registrars And Handlers

- AdvancedPetAssist.OnPlayerCast
- Effigy.Name..".StartCast"
- LibGuard.StartCast
- RegisterEventHandler
- TurretRange.OnBeginCast
- global

## Examples

- AdvancedPetAssist: AdvancedPetAssist.local.RegisterCoreEvents -> SystemData.Events.PLAYER_BEGIN_CAST -> AdvancedPetAssist.OnPlayerCast
- AdvancedPetAssist: RegisterCoreEvents -> SystemData.Events.PLAYER_BEGIN_CAST -> AdvancedPetAssist.OnPlayerCast
- Effigy: Effigy.RegisterStateInfoForCastbar -> SystemData.Events.PLAYER_BEGIN_CAST -> Effigy.Name..".StartCast"
- LibGuard: LibGuard.Init -> SystemData.Events.PLAYER_BEGIN_CAST -> LibGuard.StartCast
- TurretRange: TurretRange.Initialize -> SystemData.Events.PLAYER_BEGIN_CAST -> TurretRange.OnBeginCast
- AdvancedPetAssist: AdvancedPetAssist.OnPlayerCast -> RegisterEventHandler(SystemData.Events.PLAYER_BEGIN_CAST, AdvancedPetAssist.OnPlayerCast)

## Related APIs

- none

## Used With

- [RegisterEventHandler](../../globals/functions/global_RegisterEventHandler.md) (MEDIUM 68/100) - Global Function

## Triggered By

- [LPET](game_event_LPET.md) (HIGH 93/100) - Game Event

## Affects

- none

## Notes

- none
