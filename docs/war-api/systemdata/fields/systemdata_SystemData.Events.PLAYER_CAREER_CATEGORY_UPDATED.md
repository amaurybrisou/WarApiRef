# SystemData.Events.PLAYER_CAREER_CATEGORY_UPDATED

- Category: SystemData Field
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 3 addons

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 156

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, matches a known engine namespace, referenced by generated docs or reference files.

## Evidence Signals

- +18 Seen in 2 to 3 addons: Cross-addon spread is present but limited.
- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +18 Used in event registration or dispatch: Observed in event-driven engine hooks.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +25 Referenced by generated docs or reference files: The symbol is reinforced outside a single call page.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | AdvancedRenownTrainer, QuickNameActions+, RoR_debolster |
| Files seen in | AdvancedRenownTraining.lua, QuickNameActionsRessurected.lua, RoR_debolster.lua |
| Namespaces detected | SystemData |
| Source kinds | lua_call |
| Example locations | Initialize, OnReload, OnShutdown, lua_call |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 4 |
| Global usage count | 4 |
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

SystemData.SystemData.Events.PLAYER_CAREER_CATEGORY_UPDATED field accessed by 3 addons; commonly found in Initialize and OnReload, OnShutdown, lua_call contexts.

## Seen In

- AdvancedRenownTrainer
- QuickNameActions+
- RoR_debolster

## Related APIs

- [EA_Window_InteractionRenownTraining.GetPointsAvailable](../../globals/functions/global_EA_Window_InteractionRenownTraining.GetPointsAvailable.md) (HIGH 98/100) - Global Function
- [EA_Window_InteractionRenownTraining.GetPointsSpent](../../globals/functions/global_EA_Window_InteractionRenownTraining.GetPointsSpent.md) (HIGH 98/100) - Global Function
- [OnShutdown](../../xml/handlers/handler_OnShutdown.md) (HIGH 88/100) - XML Event
- [CreateWindow](../../globals/functions/global_CreateWindow.md) (HIGH 75/100) - Global Function
- [UnregisterEventHandler](../../globals/functions/global_UnregisterEventHandler.md) (MEDIUM 68/100) - Global Function

## Used With

- [SystemData.Events.RELOAD_INTERFACE](systemdata_SystemData.Events.RELOAD_INTERFACE.md) (HIGH 100/100) - SystemData Field

## Notes

- Observed in contexts: Initialize, OnReload, OnShutdown, lua_call
