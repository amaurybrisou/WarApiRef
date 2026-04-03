# SystemData.Events.RELOAD_INTERFACE

- Category: SystemData Field
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 62 addons

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 188

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, seen in 4 or more addons, matches a known engine namespace.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +18 Used in event registration or dispatch: Observed in event-driven engine hooks.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +25 Referenced by generated docs or reference files: The symbol is reinforced outside a single call page.
- +20 Reinforced across multiple generated source types: Evidence comes from several independent addon-api source types.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | ActionBarHide, AdvancedRenownTrainer, Atlas, Aura, BagOMatic, BuffHead, CM_ClosetGoblin, CleanUnitFrames |
| Files seen in | ActionBarHide.lua, AdvancedRenownTraining.lua, BagOMatic.lua, CleanTargetWindow.lua, ClosetGoblin.lua, Core.lua, DuffTimer.lua, Effigy.lua |
| Namespaces detected | SystemData |
| Source kinds | event_page, event_registration, lua_call |
| Example locations | Init, Initialize, LOADING_END, LoadUnitFrame, NAMBLA_Initialize, OnClickReloadUI |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 23 |
| Global usage count | 23 |
| Local definition count | 0 |
| Documentation references | 3 |
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

SystemData.SystemData.Events.RELOAD_INTERFACE field accessed by 62 addons; commonly found in Init and Initialize, LOADING_END, LoadUnitFrame, NAMBLA_Initialize, OnClickReloadUI, OnInitialize, OnShutdown, RegisterCoreEvents, RegisterEventHandlers, Shutdown, Squared.OnLoad, SystemData.Events.RELOAD_INTERFACE, ToggleHook, Unload, UnloadUnitFrame, UnregisterCoreEvents, entry_point, event_page, event_registration, init, lua_call, onInitialize contexts.

## Seen In

- ActionBarHide
- AdvancedRenownTrainer
- Atlas
- Aura
- BagOMatic
- BuffHead
- CM_ClosetGoblin
- CleanUnitFrames
- Crusher
- Default Tactic Set
- DuffTimer
- Effigy
- GroupSpotter
- HammerTime
- HealGrid
- HealHoverAssist
- Hopper
- Keyset
- KeysetMonsterPlay
- Kwestor
- Lib RuString
- LibGroup
- LibWBToggler
- MarkBuff
- MoraleBG
- Motion
- NAMBLA
- NaturalLog
- NerfedButtons
- Obsidian
- Phantom
- PlanB
- Pure
- Pure Careerbar
- QuickNameActions+
- RO-Style Combat Text
- RVMOD_Manager
- Rangechecker
- RealmStatus
- SNT_PANEL
- SOR
- SessionRPs
- Squared
- SquaredHotIndicators
- Statdoll Remix
- TexturedButtons
- TidyChat
- TidyQueue
- TidyRoll
- TimeToDie
- TokenMachine
- Tome Titan
- TurretRange
- Twister
- TwisterSet
- WarBoard
- WhoHealedMe
- Wikki's Cooldown Bar
- Wikki's Cooldown Pulse
- nLootLink
- scnoload
- xHUD

## Related APIs

- [LayoutEditor.RegisterEditCallback](../../window_api/functions/window_LayoutEditor.RegisterEditCallback.md) (HIGH 100/100) - Window Function
- [LibSlash.IsSlashCmdRegistered](../../globals/functions/global_LibSlash.IsSlashCmdRegistered.md) (HIGH 100/100) - Global Function
- [WindowRegisterEventHandler](../../window_api/functions/window_WindowRegisterEventHandler.md) (HIGH 100/100) - Window Function
- [EA_Window_InteractionRenownTraining.GetPointsAvailable](../../globals/functions/global_EA_Window_InteractionRenownTraining.GetPointsAvailable.md) (HIGH 98/100) - Global Function
- [EA_Window_InteractionRenownTraining.GetPointsSpent](../../globals/functions/global_EA_Window_InteractionRenownTraining.GetPointsSpent.md) (HIGH 98/100) - Global Function
- [OnInitialize](../../xml/handlers/handler_OnInitialize.md) (HIGH 88/100) - XML Event
- [OnShutdown](../../xml/handlers/handler_OnShutdown.md) (HIGH 88/100) - XML Event
- [CreateWindow](../../globals/functions/global_CreateWindow.md) (HIGH 75/100) - Global Function
- [RegisterEventHandler](../../globals/functions/global_RegisterEventHandler.md) (MEDIUM 68/100) - Global Function

## Used With

- [SystemData.Events.ALL_MODULES_INITIALIZED](systemdata_SystemData.Events.ALL_MODULES_INITIALIZED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.LOADING_END](systemdata_SystemData.Events.LOADING_END.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_CAREER_CATEGORY_UPDATED](systemdata_SystemData.Events.PLAYER_CAREER_CATEGORY_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_COMBAT_FLAG_UPDATED](systemdata_SystemData.Events.PLAYER_COMBAT_FLAG_UPDATED.md) (HIGH 100/100) - SystemData Field
- [OnInitialize](../../xml/handlers/handler_OnInitialize.md) (HIGH 88/100) - XML Event
- [CreateWindow](../../globals/functions/global_CreateWindow.md) (HIGH 75/100) - Global Function
- [RegisterEventHandler](../../globals/functions/global_RegisterEventHandler.md) (MEDIUM 68/100) - Global Function

## Notes

- Observed in contexts: Init, Initialize, LOADING_END, LoadUnitFrame, NAMBLA_Initialize, OnClickReloadUI
