# SystemData.Events.RELOAD_INTERFACE

- Category: SystemData Field
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 18 addons

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
| Addons seen in | AdvancedRenownTrainer, AdvancedRenownTrainer, Aura, BagOMatic, BuffHead, CM_ClosetGoblin, Effigy, LibWBToggler, PlanB, TexturedButtons, TidyChat, TidyRoll, TurretRange, WarBoard, WhoHealedMe, wbLeadHelper, Aura, BagOMatic, BuffHead, CM_ClosetGoblin, Effigy, LibGroup |
| Files seen in | `/workspace_addons/Aura/Source/AuraAddon.lua:70`, `/workspace_addons/BuffHead/Core.lua:152`, `/workspace_addons/ClosetGoblin/ClosetGoblin.lua:72`, `/workspace_addons/Effigy/Effigy.lua:111`, `/workspace_addons/LibGroup/LibGroup.lua:343`, `/workspace_addons/LibWarBoardToggler/LibWBTogglerManager.lua:12`, `/workspace_addons/PlanB/PlanB.lua:35`, `/workspace_addons/RVMOD_Manager/RVMOD_Manager.lua:1199` |
| Namespaces detected | SystemData |
| Source kinds | event_page, event_registration, flow, lua_call |
| Example locations | AdvancedRenownTraining.Initialize, AdvancedRenownTraining.OnReload, AuraAddon.OnInitialize, AuraAddon.OnLoad, BagOMatic.init, BagOMatic.restore_filters |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 40 |
| Global usage count | 40 |
| Local definition count | 0 |
| Documentation references | 4 |
| Initialization flow references | 13 |
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

Observed SystemData field used by 18 addons through generated function calls, event pages, or lifecycle evidence.

## Seen In

- AdvancedRenownTrainer
- AdvancedRenownTrainer, Aura, BagOMatic, BuffHead, CM_ClosetGoblin, Effigy, LibWBToggler, PlanB, TexturedButtons, TidyChat, TidyRoll, TurretRange, WarBoard, WhoHealedMe, wbLeadHelper
- Aura
- BagOMatic
- BuffHead
- CM_ClosetGoblin
- Effigy
- LibGroup
- LibWBToggler
- PlanB
- RVMOD_Manager
- TexturedButtons
- TidyChat
- TidyRoll
- TurretRange
- WarBoard
- WhoHealedMe
- wbLeadHelper

## Related APIs

- [LayoutEditor.RegisterEditCallback](../../window_api/functions/window_LayoutEditor.RegisterEditCallback.md) (HIGH 100/100) - Window Function
- [LibSlash.IsSlashCmdRegistered](../../globals/functions/global_LibSlash.IsSlashCmdRegistered.md) (HIGH 100/100) - Global Function
- [LibSlash.RegisterSlashCmd](../../globals/functions/global_LibSlash.RegisterSlashCmd.md) (HIGH 100/100) - Global Function
- [WindowSetParent](../../window_api/functions/window_WindowSetParent.md) (HIGH 100/100) - Window Function
- [CreateWindow](../../globals/functions/global_CreateWindow.md) (HIGH 75/100) - Global Function
- [RegisterEventHandler](../../globals/functions/global_RegisterEventHandler.md) (MEDIUM 68/100) - Global Function

## Used With

- [EA_Window_Backpack](../../globals/tables/table_EA_Window_Backpack.md) (HIGH 100/100) - Global Table
- [EA_Window_InteractionRenownTraining.Hide](../../globals/functions/global_EA_Window_InteractionRenownTraining.Hide.md) (HIGH 100/100) - Global Function
- [EA_Window_InteractionRenownTraining.Show](../../globals/functions/global_EA_Window_InteractionRenownTraining.Show.md) (HIGH 100/100) - Global Function
- [LibSlash.IsSlashCmdRegistered](../../globals/functions/global_LibSlash.IsSlashCmdRegistered.md) (HIGH 100/100) - Global Function
- [LibSlash.RegisterSlashCmd](../../globals/functions/global_LibSlash.RegisterSlashCmd.md) (HIGH 100/100) - Global Function
- [OnInitialize](../../events/window_events/window_event_OnInitialize.md) (HIGH 100/100) - Window Event
- [OnInitialize](../../xml/handlers/handler_OnInitialize.md) (HIGH 100/100) - XML Handler
- [SystemData.Events.LOADING_END](systemdata_SystemData.Events.LOADING_END.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_CAREER_CATEGORY_UPDATED](systemdata_SystemData.Events.PLAYER_CAREER_CATEGORY_UPDATED.md) (HIGH 100/100) - SystemData Field
- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type
- [WindowAddAnchor](../../window_api/functions/window_WindowAddAnchor.md) (HIGH 100/100) - Window Function
- [WindowSetParent](../../window_api/functions/window_WindowSetParent.md) (HIGH 100/100) - Window Function
- [DoesWindowExist](../../globals/functions/global_DoesWindowExist.md) (HIGH 83/100) - Global Function
- [CreateWindow](../../globals/functions/global_CreateWindow.md) (HIGH 75/100) - Global Function
- [RegisterEventHandler](../../globals/functions/global_RegisterEventHandler.md) (MEDIUM 68/100) - Global Function

## Triggered By

- none

## Affects

- none

## Notes

- Observed in contexts: AdvancedRenownTraining.Initialize, AdvancedRenownTraining.OnReload, AuraAddon.OnInitialize, AuraAddon.OnLoad, BagOMatic.init, BagOMatic.restore_filters
