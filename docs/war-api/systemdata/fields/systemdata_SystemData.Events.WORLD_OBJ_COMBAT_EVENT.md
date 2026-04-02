# SystemData.Events.WORLD_OBJ_COMBAT_EVENT

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
| Addons seen in | CombatTextNames, CombatTextNames, Effigy, WSCT, Effigy, LoyalPet, WSCT |
| Files seen in | `/workspace_addons/Effigy/States/EffigyStateCastbar.lua:41`, `/workspace_addons/LoyalPet/LoyalPet.lua:134`, `/workspace_addons/LoyalPet/LoyalPet.lua:203`, `/workspace_addons/combattextnames/cmd.lua:49`, `/workspace_addons/combattextnames/cmd.lua:55`, `/workspace_addons/wsct/wsct.lua:117`, `/workspace_addons/wsct/wsct.lua:137` |
| Namespaces detected | SystemData |
| Source kinds | event_page, event_registration, flow, lua_call |
| Example locations | CombatTextNames.Disable, CombatTextNames.Enable, EA_System_EventText.AddCombatEventText, Effigy.Name..".OnWorldObjCombatEvent", Effigy.RegisterStateInfoForCastbar, LPET.Initialize |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 16 |
| Global usage count | 16 |
| Local definition count | 0 |
| Documentation references | 4 |
| Initialization flow references | 2 |
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

- CombatTextNames
- CombatTextNames, Effigy, WSCT
- Effigy
- LoyalPet
- WSCT

## Related APIs

- [WindowRegisterCoreEventHandler](../../window_api/functions/window_WindowRegisterCoreEventHandler.md) (HIGH 100/100) - Window Function

## Used With

- none

## Triggered By

- none

## Affects

- none

## Notes

- Observed in contexts: CombatTextNames.Disable, CombatTextNames.Enable, EA_System_EventText.AddCombatEventText, Effigy.Name..".OnWorldObjCombatEvent", Effigy.RegisterStateInfoForCastbar, LPET.Initialize
