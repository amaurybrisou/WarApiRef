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
| Addons seen in | AdvancedPetAssist, AdvancedRenownTrainer, Aura, BagOMatic, BuffHead, CM_ClosetGoblin, Effigy, Enemy |
| Files seen in | `/workspace_addons/AdvancedPetAssist/AdvancedPetAssist.lua:92`, `/workspace_addons/Aura/Source/AuraAddon.lua:70`, `/workspace_addons/BuffHead/Core.lua:152`, `/workspace_addons/ClosetGoblin/ClosetGoblin.lua:1240`, `/workspace_addons/ClosetGoblin/ClosetGoblin.lua:72`, `/workspace_addons/Effigy/Effigy.lua:111`, `/workspace_addons/Effigy/States/EffigyStateCastbar.lua:41`, `/workspace_addons/Effigy/States/EffigyStatePlayer.lua:45` |
| Namespaces detected | SystemData |
| Source kinds | event_page, flows, lua_event_registration |
| Example locations | AdvancedPetAssist: AdvancedPetAssist.local.RegisterLoadingEnd, AdvancedPetAssist: RegisterLoadingEnd, AdvancedRenownTrainer: AdvancedRenownTraining.Initialize, Aura: AuraAddon.OnInitialize, BagOMatic: BagOMatic.init, BuffHead: BuffHead.Initialize |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 30 |
| Global usage count | 30 |
| Local definition count | 0 |
| Documentation references | 2 |
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

Observed as a shared SystemData runtime event used by 25 addons.

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
- Effigy
- Enemy
- Killer
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

## Registrars And Handlers

- AdvancedPetAssist.OnLoadingEnd
- AdvancedRenownTraining.OnReload
- AuraAddon.OnLoad
- BagOMatic.restore_filters
- BuffHead.OnLoadingEnd
- ClosetGoblin.Initialize
- ClosetGoblin.LoadingEnd
- Effigy.Name..".OnLoad"
- Effigy.Name..".OnLoadCastbar"
- Effigy.ResetFMOT
- Effigy.UpdatePlayerAP
- Enemy.ScenarioInfoUpdate
- Killer.OnLoadingEnd
- LibWBTogglerManager.CheckMods
- MapMonster.OnLoadingEnd
- MiracleGrow.onZone
- MiracleGrow2.onZone
- MiracleGrowLight.onZone
- PP.PreviewDyes
- PetFixWindow.PetFixes
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
- wbLeadHelper.LOADING_END
- wbLeadHelper._ApplyMiscWhitelistGate

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
- [RegisterEventHandler](../../globals/functions/global_RegisterEventHandler.md) (MEDIUM 68/100) - Global Function
- [UnregisterEventHandler](../../globals/functions/global_UnregisterEventHandler.md) (MEDIUM 68/100) - Global Function

## Used With

- [SystemData.Events.RELOAD_INTERFACE](game_event_SystemData.Events.RELOAD_INTERFACE.md) (HIGH 100/100) - Game Event

## Triggered By

- none

## Affects

- none

## Notes

- none
