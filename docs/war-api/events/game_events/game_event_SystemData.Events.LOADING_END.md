# SystemData.Events.LOADING_END

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
| Addons seen in | AdvancedPetAssist, AdvancedRenownTrainer, Aura, BagOMatic, BuffHead, CM_ClosetGoblin, Enemy, Killer |
| Files seen in | `/workspace/data/raw/AdvancedPetAssist/AdvancedPetAssist.lua:92`, `/workspace/data/raw/Aura/Source/AuraAddon.lua:70`, `/workspace/data/raw/BuffHead/Core.lua:152`, `/workspace/data/raw/ClosetGoblin/ClosetGoblin.lua:1300`, `/workspace/data/raw/ClosetGoblin/ClosetGoblin.lua:73`, `/workspace/data/raw/Enemy/Code/ScenarioInfo/ScenarioInfo.lua:303`, `/workspace/data/raw/Killer/KillerLifecycle.lua:4`, `/workspace/data/raw/LibWarBoardToggler/LibWBTogglerManager.lua:12` |
| Namespaces detected | SystemData |
| Source kinds | event_page, flows, lua_event_registration |
| Example locations | AdvancedPetAssist: AdvancedPetAssist.local.RegisterLoadingEnd, AdvancedPetAssist: RegisterLoadingEnd, AdvancedRenownTrainer: AdvancedRenownTraining.Initialize, Aura: AuraAddon.OnInitialize, BagOMatic: BagOMatic.init, BuffHead: BuffHead.Initialize |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 21 |
| Global usage count | 21 |
| Local definition count | 0 |
| Documentation references | 2 |
| Initialization flow references | 18 |
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

Observed as a shared SystemData runtime event used by 19 addons.

## Handler Pattern

Observed as a runtime event ID routed through RegisterEventHandler-style APIs.

## Payload

- Payload shape is not inferable from addon-api docs alone; treat this as an engine event identifier.

## Seen In

- AdvancedPetAssist
- AdvancedRenownTrainer
- Aura
- BagOMatic
- BuffHead
- CM_ClosetGoblin
- Enemy
- Killer
- LibWBToggler
- MiracleGrowLight
- PlanB
- Pocket Palette
- RoR_SoR
- TexturedButtons
- TidyChat
- TidyRoll
- TurretRange
- WarBoard
- WhoHealedMe

## Registrars And Handlers

- AdvancedPetAssist.OnLoadingEnd
- AdvancedRenownTraining.OnReload
- AuraAddon.OnLoad
- BagOMatic.restore_filters
- BuffHead.OnLoadingEnd
- ClosetGoblin.Initialize
- ClosetGoblin.LoadingEnd
- Enemy.ScenarioInfoUpdate
- Killer.OnLoadingEnd
- LibWBTogglerManager.CheckMods
- MiracleGrowLight.onZone
- PP.PreviewDyes
- PlanB.HandleLoading
- RegisterEventHandler
- RoR_SoR.OnLoadingEnd
- TexturedButtons.OnLoadComplete
- TidyChat.OnLoad
- TidyRoll.OnLoad
- TurretRange.OnLoadComplete
- WHMEvents.OnLoadingEnd
- WarBoard.SortMods
- global

## Examples

- AdvancedPetAssist: AdvancedPetAssist.local.RegisterLoadingEnd -> SystemData.Events.LOADING_END -> AdvancedPetAssist.OnLoadingEnd
- AdvancedPetAssist: RegisterLoadingEnd -> SystemData.Events.LOADING_END -> AdvancedPetAssist.OnLoadingEnd
- AdvancedRenownTrainer: AdvancedRenownTraining.Initialize -> SystemData.Events.LOADING_END -> AdvancedRenownTraining.OnReload
- Aura: AuraAddon.OnInitialize -> SystemData.Events.LOADING_END -> AuraAddon.OnLoad
- BagOMatic: BagOMatic.init -> SystemData.Events.LOADING_END -> BagOMatic.restore_filters
- BuffHead: BuffHead.Initialize -> SystemData.Events.LOADING_END -> BuffHead.OnLoadingEnd

## Related APIs

- [LibSlash.RegisterWSlashCmd](../../globals/functions/global_LibSlash.RegisterWSlashCmd.md) (HIGH 100/100) - Global Function
- [WindowSetParent](../../window_api/functions/window_WindowSetParent.md) (HIGH 100/100) - Window Function
- [WindowUnregisterCoreEventHandler](../../window_api/functions/window_WindowUnregisterCoreEventHandler.md) (HIGH 100/100) - Window Function
- [RegisterEventHandler](../../globals/functions/global_RegisterEventHandler.md) (HIGH 93/100) - Global Function
- [UnregisterEventHandler](../../globals/functions/global_UnregisterEventHandler.md) (HIGH 93/100) - Global Function

## Used With

- [LibSlash.RegisterWSlashCmd](../../globals/functions/global_LibSlash.RegisterWSlashCmd.md) (HIGH 100/100) - Global Function
- [SystemData.Events.RELOAD_INTERFACE](game_event_SystemData.Events.RELOAD_INTERFACE.md) (HIGH 100/100) - Game Event
- [DoesWindowExist](../../globals/functions/global_DoesWindowExist.md) (HIGH 83/100) - Global Function

## Triggered By

- none

## Affects

- none

## Notes

- none
