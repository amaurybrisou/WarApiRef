# SystemData.Events.ALL_MODULES_INITIALIZED

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
- +20 Reinforced across multiple generated source types: Evidence comes from several independent addon-api source types.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | CM_ClosetGoblin, Effigy, RVAPI_ColorDialog, RVMOD_Manager, RoR_SoR |
| Files seen in | `/workspace_addons/ClosetGoblin/ClosetGoblin.lua:72`, `/workspace_addons/Effigy/Effigy.lua:111`, `/workspace_addons/RVAPI_ColorDialog/RVAPI_ColorDialog.lua:83`, `/workspace_addons/RVMOD_Manager/RVMOD_Manager.lua:210`, `/workspace_addons/RoR_SoR/RoR_SoR.lua:178` |
| Namespaces detected | SystemData |
| Source kinds | event_page, flows, lua_event_registration |
| Example locations | CM_ClosetGoblin: ClosetGoblin.OnInitialize, Effigy: Effigy.Initialize, RVAPI_ColorDialog: RVAPI_ColorDialog.Initialize, RVMOD_Manager: RVMOD_Manager.Initialize, RoR_SoR: RoR_SoR:RegisterSelfEvents |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 5 |
| Global usage count | 5 |
| Local definition count | 0 |
| Documentation references | 2 |
| Initialization flow references | 4 |
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

Observed as a shared SystemData runtime event used by 5 addons.

## Handler Pattern

Observed as a runtime event ID routed through RegisterEventHandler-style APIs.

## Payload

- Payload shape is not inferable from addon-api docs alone; treat this as an engine event identifier.

## Seen In

- CM_ClosetGoblin
- Effigy
- RVAPI_ColorDialog
- RVMOD_Manager
- RoR_SoR

## Registrars And Handlers

- ClosetGoblin.OnAllModulesInitialized
- Effigy.Name..".OnAllModulesInitialized"
- RVAPI_ColorDialog.OnAllModulesInitialized
- RVMOD_Manager.OnAllModulesInitialized
- RegisterEventHandler
- RoR_SoR.Enable
- global

## Examples

- CM_ClosetGoblin: ClosetGoblin.OnInitialize -> SystemData.Events.ALL_MODULES_INITIALIZED -> ClosetGoblin.OnAllModulesInitialized
- Effigy: Effigy.Initialize -> SystemData.Events.ALL_MODULES_INITIALIZED -> Effigy.Name..".OnAllModulesInitialized"
- RVAPI_ColorDialog: RVAPI_ColorDialog.Initialize -> SystemData.Events.ALL_MODULES_INITIALIZED -> RVAPI_ColorDialog.OnAllModulesInitialized
- RVMOD_Manager: RVMOD_Manager.Initialize -> SystemData.Events.ALL_MODULES_INITIALIZED -> RVMOD_Manager.OnAllModulesInitialized
- RoR_SoR: RoR_SoR:RegisterSelfEvents -> SystemData.Events.ALL_MODULES_INITIALIZED -> RoR_SoR.Enable
- CM_ClosetGoblin: ClosetGoblin.OnAllModulesInitialized -> RegisterEventHandler(SystemData.Events.ALL_MODULES_INITIALIZED, ClosetGoblin.OnAllModulesInitialized)

## Related APIs

- none

## Used With

- none

## Triggered By

- none

## Affects

- none

## Notes

- none
