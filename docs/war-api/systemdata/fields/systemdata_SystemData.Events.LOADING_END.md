# SystemData.Events.LOADING_END

- Category: SystemData Field
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 27 addons

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
| Addons seen in | AdvancedPetAssist, AdvancedPetAssist, AdvancedRenownTrainer, Aura, BagOMatic, BuffHead, CM_ClosetGoblin, Effigy, Enemy, Killer, LibWBToggler, MapMonster, Miracle Grow Remix, MiracleGrow, MiracleGrowLight, PetFixWindow, PlanB, Pocket Palette, RoR_SoR, TexturedButtons, TidyChat, TidyRoll, TurretRange, WarBoard, WhoHealedMe, wbLeadHelper, AdvancedRenownTrainer, Aura, BagOMatic, BuffHead, CM_ClosetGoblin, Effigy |
| Files seen in | `/workspace_addons/AdvancedPetAssist/AdvancedPetAssist.lua:92`, `/workspace_addons/AdvancedPetAssist/AdvancedPetAssist.lua:98`, `/workspace_addons/Aura/Source/AuraAddon.lua:70`, `/workspace_addons/BuffHead/Core.lua:152`, `/workspace_addons/ClosetGoblin/ClosetGoblin.lua:1300`, `/workspace_addons/ClosetGoblin/ClosetGoblin.lua:1409`, `/workspace_addons/ClosetGoblin/ClosetGoblin.lua:73`, `/workspace_addons/ClosetGoblin/ClosetGoblin.lua:87` |
| Namespaces detected | SystemData |
| Source kinds | event_page, event_registration, flow, lua_call |
| Example locations | AdvancedPetAssist.OnLoadingEnd, AdvancedPetAssist.local.RegisterLoadingEnd, AdvancedPetAssist.local.UnregisterLoadingEnd, AdvancedRenownTraining.Initialize, AdvancedRenownTraining.OnReload, AuraAddon.OnInitialize |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 77 |
| Global usage count | 77 |
| Local definition count | 0 |
| Documentation references | 4 |
| Initialization flow references | 23 |
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

Observed SystemData field used by 27 addons through generated function calls, event pages, or lifecycle evidence.

## Seen In

- AdvancedPetAssist
- AdvancedPetAssist, AdvancedRenownTrainer, Aura, BagOMatic, BuffHead, CM_ClosetGoblin, Effigy, Enemy, Killer, LibWBToggler, MapMonster, Miracle Grow Remix, MiracleGrow, MiracleGrowLight, PetFixWindow, PlanB, Pocket Palette, RoR_SoR, TexturedButtons, TidyChat, TidyRoll, TurretRange, WarBoard, WhoHealedMe, wbLeadHelper
- AdvancedRenownTrainer
- Aura
- BagOMatic
- BuffHead
- CM_ClosetGoblin
- Effigy
- Enemy
- Killer
- LibGroup
- LibWBToggler
- MapMonster
- Miracle Grow Remix
- MiracleGrow
- MiracleGrowLight
- PetFixWindow
- PlanB
- Pocket Palette
- RoR_SoR
- TexturedButtons
- TidyChat
- TidyRoll
- TurretRange
- WarBoard
- WhoHealedMe
- wbLeadHelper

## Related APIs

- [LayoutEditor.RegisterEditCallback](../../window_api/functions/window_LayoutEditor.RegisterEditCallback.md) (HIGH 100/100) - Window Function
- [LayoutEditor.UnregisterWindow](../../window_api/functions/window_LayoutEditor.UnregisterWindow.md) (HIGH 100/100) - Window Function
- [LibSlash.IsSlashCmdRegistered](../../globals/functions/global_LibSlash.IsSlashCmdRegistered.md) (HIGH 100/100) - Global Function
- [LibSlash.RegisterSlashCmd](../../globals/functions/global_LibSlash.RegisterSlashCmd.md) (HIGH 100/100) - Global Function
- [LibSlash.UnregisterSlashCmd](../../globals/functions/global_LibSlash.UnregisterSlashCmd.md) (HIGH 100/100) - Global Function
- [WindowRegisterCoreEventHandler](../../window_api/functions/window_WindowRegisterCoreEventHandler.md) (HIGH 100/100) - Window Function
- [WindowRegisterEventHandler](../../window_api/functions/window_WindowRegisterEventHandler.md) (HIGH 100/100) - Window Function
- [WindowSetParent](../../window_api/functions/window_WindowSetParent.md) (HIGH 100/100) - Window Function
- [CreateWindow](../../globals/functions/global_CreateWindow.md) (HIGH 75/100) - Global Function
- [RegisterEventHandler](../../globals/functions/global_RegisterEventHandler.md) (MEDIUM 68/100) - Global Function
- [UnregisterEventHandler](../../globals/functions/global_UnregisterEventHandler.md) (MEDIUM 68/100) - Global Function

## Used With

- [EA_Window_InteractionRenownTraining.Hide](../../globals/functions/global_EA_Window_InteractionRenownTraining.Hide.md) (HIGH 100/100) - Global Function
- [LibSlash.IsSlashCmdRegistered](../../globals/functions/global_LibSlash.IsSlashCmdRegistered.md) (HIGH 100/100) - Global Function
- [LibSlash.RegisterSlashCmd](../../globals/functions/global_LibSlash.RegisterSlashCmd.md) (HIGH 100/100) - Global Function
- [OnInitialize](../../events/window_events/window_event_OnInitialize.md) (HIGH 100/100) - Window Event
- [SystemData.Events.RELOAD_INTERFACE](systemdata_SystemData.Events.RELOAD_INTERFACE.md) (HIGH 100/100) - SystemData Field
- [DoesWindowExist](../../globals/functions/global_DoesWindowExist.md) (HIGH 83/100) - Global Function
- [CreateWindow](../../globals/functions/global_CreateWindow.md) (HIGH 75/100) - Global Function
- [RegisterEventHandler](../../globals/functions/global_RegisterEventHandler.md) (MEDIUM 68/100) - Global Function

## Triggered By

- none

## Affects

- none

## Notes

- Observed in contexts: AdvancedPetAssist.OnLoadingEnd, AdvancedPetAssist.local.RegisterLoadingEnd, AdvancedPetAssist.local.UnregisterLoadingEnd, AdvancedRenownTraining.Initialize, AdvancedRenownTraining.OnReload, AuraAddon.OnInitialize
