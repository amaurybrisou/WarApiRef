# SystemData.Events.WORLD_OBJ_COMBAT_EVENT

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
| Addons seen in | CombatTextNames, Effigy, RetAlert, WSCT |
| Files seen in | `/workspace/Effigy/States/EffigyStateCastbar.lua:41`, `/workspace/RetAlert/Source/RetAlert.lua:61`, `/workspace/combattextnames/cmd.lua:55`, `/workspace/wsct/wsct.lua:117`, `/workspace/wsct/wsct.lua:137` |
| Namespaces detected | SystemData |
| Source kinds | event_page, flows, lua_event_registration |
| Example locations | CombatTextNames: CombatTextNames.Enable, Effigy: Effigy.RegisterStateInfoForCastbar, RetAlert: RetAlert.local.RetAlert_Enable, RetAlert: RetAlert_Enable, WSCT: WSCT:RegisterSelfEvents, WSCT: WSCT:UnregisterSelfEvents |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 6 |
| Global usage count | 6 |
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

- CombatTextNames
- Effigy
- RetAlert
- WSCT

## Registrars And Handlers

- EA_System_EventText.AddCombatEventText
- Effigy.Name..".OnWorldObjCombatEvent"
- RegisterEventHandler
- RetAlert_OnWorldObjCombatEvent
- WSCT.OnCombatEvent
- global

## Examples

- CombatTextNames: CombatTextNames.Enable -> SystemData.Events.WORLD_OBJ_COMBAT_EVENT -> EA_System_EventText.AddCombatEventText
- Effigy: Effigy.RegisterStateInfoForCastbar -> SystemData.Events.WORLD_OBJ_COMBAT_EVENT -> Effigy.Name..".OnWorldObjCombatEvent"
- RetAlert: RetAlert.local.RetAlert_Enable -> SystemData.Events.WORLD_OBJ_COMBAT_EVENT -> RetAlert_OnWorldObjCombatEvent
- RetAlert: RetAlert_Enable -> SystemData.Events.WORLD_OBJ_COMBAT_EVENT -> RetAlert_OnWorldObjCombatEvent
- WSCT: WSCT:RegisterSelfEvents -> SystemData.Events.WORLD_OBJ_COMBAT_EVENT -> WSCT.OnCombatEvent
- WSCT: WSCT:UnregisterSelfEvents -> SystemData.Events.WORLD_OBJ_COMBAT_EVENT -> EA_System_EventText.AddCombatEventText

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
