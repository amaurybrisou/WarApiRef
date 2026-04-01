# RegisterEventHandler

- Category: Global Function
- Confidence level: MEDIUM
- Confidence score: 68/100
- Seen in: 58 addons

## Confidence Assessment

- Level: MEDIUM

- Score: 68/100

- Rationale: Promoted as MEDIUM confidence because seen in 4 or more addons, called globally with no local definition, used in event registration or dispatch.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +18 Used in event registration or dispatch: Observed in event-driven engine hooks.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- -15 Conflicting signatures across usages: Observed arity or argument shape conflicts across usages.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | AdvancedPetAssist, AdvancedRenownTrainer, AggroMeter, Aura, AutoMark, Autolooter, BagOMatic, BankArkel |
| Files seen in | `/workspace/AdvancedPetAssist/AdvancedPetAssist.lua:92`, `/workspace/AggroMeter/AggroMeter.lua:5`, `/workspace/Aura/Source/AuraAddon.lua:70`, `/workspace/AutoMark/Source/AutoMark.lua:33`, `/workspace/AutoMark/Source/AutoMark.lua:78`, `/workspace/Autolooter/Autolooter.lua:3`, `/workspace/BankArkel/BankArkel.lua:95`, `/workspace/BuffHead/Core.lua:152` |
| Namespaces detected | RegisterEventHandler |
| Source kinds | lua_calls |
| Example locations | AdvancedPetAssist: AdvancedPetAssist.local.RegisterLoadingEnd, AdvancedPetAssist: RegisterLoadingEnd, AdvancedRenownTrainer: AdvancedRenownTraining.Initialize, AdvancedRenownTrainer: AdvancedRenownTraining.OnReload, AggroMeter: AggroMeter.Initialize, Aura: AuraAddon.OnInitialize |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 321 |
| Global usage count | 321 |
| Local definition count | 0 |
| Documentation references | 0 |
| Initialization flow references | 0 |
| Known engine namespace | no |
| Default UI presence | no |
| Event binding presence | yes |
| Observed in XML and Lua | no |
| Consistent role | yes |
| Consistent arguments | no |
| Consistent returns | no |
| Slash command presence | no |
| Weak usage only | no |
| Project-specific name | no |
| Placeholder or computed name | no |
| Conflicting signatures | yes |
| Conflicting roles | no |
| Wrapper likely | no |
| Never outside local graph | no |
| Local helper only | no |

## Signature (inferred)

```lua
RegisterEventHandler(eventId, handlerName)
```

## Description

Observed registering global runtime handlers against shared event identifiers.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| eventId | Observed as a SystemData or runtime event identifier. | Observed values: "EA_CareerResourceWindow", "LPET", S.combatLogEventId |
| handlerName | Observed as a Lua handler function reference. | Observed values: "AdvancedRenownTraining.CreateDataTable", "AdvancedRenownTraining.OnReload", "AggroMeter.OnChatLogUpdated" |

## Returns

- Not confidently inferable from addon-api docs alone.

## Side Effects

- Registers engine or library callbacks for later dispatch.

## Seen In

- AdvancedPetAssist
- AdvancedRenownTrainer
- AggroMeter
- Aura
- AutoMark
- Autolooter
- BagOMatic
- BankArkel
- BuffHead
- CM_ClosetGoblin
- CombatTextNames
- DAoCBuff
- Effigy
- Enemy
- FastInteract
- GCDsaver
- GetStats
- HealAll
- JunkDump
- Killer
- LibGuard
- LibWBToggler
- LoyalPet
- MapMonster
- MegaphonePlusPlus
- Miracle Grow Remix
- MiracleGrow
- MiracleGrowLight
- PeaceOut
- PetFixWindow
- PlanB
- Pocket Palette
- PotionBar
- Queue Queuer
- QuickTacticSwitch
- QuickWarReport
- RRCount
- RVAPI_ColorDialog
- RVMOD_Manager
- RetAlert
- RoR_SoR
- ScenarioAlert
- Shinies
- Targets
- TexturedButtons
- TidyChat
- TidyQueue
- TidyRoll
- TimeInQueue
- TortallDPSCore
- TurretRange
- Twister
- WSCT
- WarBoard
- WhoHealedMe
- WoH-Reticle
- followTheLeader
- wbLeadHelper

## Examples

- AdvancedPetAssist: AdvancedPetAssist.local.RegisterLoadingEnd -> RegisterEventHandler(SystemData.Events.LOADING_END, LOADING_END_HANDLER)
- AdvancedPetAssist: RegisterLoadingEnd -> RegisterEventHandler(SystemData.Events.LOADING_END, LOADING_END_HANDLER)
- AdvancedRenownTrainer: AdvancedRenownTraining.Initialize -> RegisterEventHandler(SystemData.Events.LOADING_END, "AdvancedRenownTraining.OnReload")
- AdvancedRenownTrainer: AdvancedRenownTraining.Initialize -> RegisterEventHandler(SystemData.Events.PLAYER_CAREER_CATEGORY_UPDATED, "AdvancedRenownTraining.CreateDataTable")
- AdvancedRenownTrainer: AdvancedRenownTraining.Initialize -> RegisterEventHandler(SystemData.Events.RELOAD_INTERFACE, "AdvancedRenownTraining.OnReload")
- AdvancedRenownTrainer: AdvancedRenownTraining.OnReload -> RegisterEventHandler(SystemData.Events.PLAYER_CAREER_CATEGORY_UPDATED, "AdvancedRenownTraining.CreateDataTable")

## Related APIs

- none

## Used With

- [BankWindow](../tables/table_BankWindow.md) (HIGH 100/100) - Global Table
- [DialogManager.MakeOneButtonDialog](global_DialogManager.MakeOneButtonDialog.md) (HIGH 100/100) - Global Function
- [EA_ChatWindow.OnKeyEnter](global_EA_ChatWindow.OnKeyEnter.md) (HIGH 100/100) - Global Function
- [EA_ChatWindow.Print](global_EA_ChatWindow.Print.md) (HIGH 100/100) - Global Function
- [EA_Window_Backpack](../tables/table_EA_Window_Backpack.md) (HIGH 100/100) - Global Table
- [EA_Window_InteractionRenownTraining.Hide](global_EA_Window_InteractionRenownTraining.Hide.md) (HIGH 100/100) - Global Function
- [EA_Window_InteractionRenownTraining.Show](global_EA_Window_InteractionRenownTraining.Show.md) (HIGH 100/100) - Global Function
- [LayoutEditor.RegisterEditCallback](../../window_api/functions/window_LayoutEditor.RegisterEditCallback.md) (HIGH 100/100) - Window Function
- [LayoutEditor.RegisterWindow](../../window_api/functions/window_LayoutEditor.RegisterWindow.md) (HIGH 100/100) - Window Function
- [LibSlash.IsSlashCmdRegistered](global_LibSlash.IsSlashCmdRegistered.md) (HIGH 100/100) - Global Function
- [LibSlash.RegisterSlashCmd](global_LibSlash.RegisterSlashCmd.md) (HIGH 100/100) - Global Function
- [OnInitialize](../../events/window_events/window_event_OnInitialize.md) (HIGH 100/100) - Window Event
- [OnInitialize](../../xml/handlers/handler_OnInitialize.md) (HIGH 100/100) - XML Handler
- [SystemData.Events.INTERACT_OPEN_BANK](../../systemdata/fields/systemdata_SystemData.Events.INTERACT_OPEN_BANK.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.INTERFACE_RELOADED](../../events/game_events/game_event_SystemData.Events.INTERFACE_RELOADED.md) (HIGH 100/100) - Game Event
- [SystemData.Events.LOADING_END](../../events/game_events/game_event_SystemData.Events.LOADING_END.md) (HIGH 100/100) - Game Event
- [SystemData.Events.LOADING_END](../../systemdata/fields/systemdata_SystemData.Events.LOADING_END.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.LOG_OUT](../../systemdata/fields/systemdata_SystemData.Events.LOG_OUT.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.L_BUTTON_DOWN_PROCESSED](../../events/game_events/game_event_SystemData.Events.L_BUTTON_DOWN_PROCESSED.md) (HIGH 100/100) - Game Event
- [SystemData.Events.M_BUTTON_DOWN_PROCESSED](../../events/game_events/game_event_SystemData.Events.M_BUTTON_DOWN_PROCESSED.md) (HIGH 100/100) - Game Event
- [SystemData.Events.PLAYER_BEGIN_CAST](../../events/game_events/game_event_SystemData.Events.PLAYER_BEGIN_CAST.md) (HIGH 100/100) - Game Event
- [SystemData.Events.PLAYER_CAREER_CATEGORY_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_CAREER_CATEGORY_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_CAREER_RESOURCE_UPDATED](../../events/game_events/game_event_SystemData.Events.PLAYER_CAREER_RESOURCE_UPDATED.md) (HIGH 100/100) - Game Event
- [SystemData.Events.PLAYER_CAREER_RESOURCE_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_CAREER_RESOURCE_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_COMBAT_FLAG_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_COMBAT_FLAG_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_DEATH](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_DEATH.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_EFFECTS_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_EFFECTS_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_PET_STATE_UPDATED](../../events/game_events/game_event_SystemData.Events.PLAYER_PET_STATE_UPDATED.md) (HIGH 100/100) - Game Event
- [SystemData.Events.PLAYER_PET_TARGET_UPDATED](../../events/game_events/game_event_SystemData.Events.PLAYER_PET_TARGET_UPDATED.md) (HIGH 100/100) - Game Event
- [SystemData.Events.PLAYER_PET_UPDATED](../../events/game_events/game_event_SystemData.Events.PLAYER_PET_UPDATED.md) (HIGH 100/100) - Game Event
- [SystemData.Events.PLAYER_TARGET_UPDATED](../../events/game_events/game_event_SystemData.Events.PLAYER_TARGET_UPDATED.md) (HIGH 100/100) - Game Event
- [SystemData.Events.PLAYER_TARGET_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_TARGET_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.RELOAD_INTERFACE](../../events/game_events/game_event_SystemData.Events.RELOAD_INTERFACE.md) (HIGH 100/100) - Game Event
- [SystemData.Events.RELOAD_INTERFACE](../../systemdata/fields/systemdata_SystemData.Events.RELOAD_INTERFACE.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.R_BUTTON_DOWN_PROCESSED](../../events/game_events/game_event_SystemData.Events.R_BUTTON_DOWN_PROCESSED.md) (HIGH 100/100) - Game Event
- [SystemData.Events.WORLD_OBJ_COMBAT_EVENT](../../events/game_events/game_event_SystemData.Events.WORLD_OBJ_COMBAT_EVENT.md) (HIGH 100/100) - Game Event
- [TextLogGetUpdateEventId](../../events/game_events/game_event_TextLogGetUpdateEventId.md) (HIGH 100/100) - Game Event
- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type
- [WindowAddAnchor](../../window_api/functions/window_WindowAddAnchor.md) (HIGH 100/100) - Window Function
- [WindowRegisterCoreEventHandler](../../window_api/functions/window_WindowRegisterCoreEventHandler.md) (HIGH 100/100) - Window Function
- [WindowRegisterEventHandler](../../window_api/functions/window_WindowRegisterEventHandler.md) (HIGH 100/100) - Window Function
- [WindowSetLayer](../../window_api/functions/window_WindowSetLayer.md) (HIGH 100/100) - Window Function
- [WindowSetParent](../../window_api/functions/window_WindowSetParent.md) (HIGH 100/100) - Window Function
- [WindowSetShowing](../../window_api/functions/window_WindowSetShowing.md) (HIGH 100/100) - Window Function
- [BankWindow.Hide](global_BankWindow.Hide.md) (HIGH 88/100) - Global Function
- [DoesWindowExist](global_DoesWindowExist.md) (HIGH 83/100) - Global Function
- [CreateWindow](global_CreateWindow.md) (HIGH 75/100) - Global Function
- [CreateWindowFromTemplate](global_CreateWindowFromTemplate.md) (HIGH 75/100) - Global Function
- [UnregisterEventHandler](global_UnregisterEventHandler.md) (MEDIUM 68/100) - Global Function
- [selffriendlytarget](../../events/game_events/game_event_selffriendlytarget.md) (MEDIUM 43/100) - Game Event
- [selfhostiletarget](../../events/game_events/game_event_selfhostiletarget.md) (MEDIUM 43/100) - Game Event

## Triggered By

- [OnInitialize](../../xml/handlers/handler_OnInitialize.md) (HIGH 100/100) - XML Handler
- [OnInitialize](../../events/window_events/window_event_OnInitialize.md) (HIGH 100/100) - Window Event
- [SystemData.Events.LOADING_END](../../events/game_events/game_event_SystemData.Events.LOADING_END.md) (HIGH 100/100) - Game Event
- [SystemData.Events.RELOAD_INTERFACE](../../events/game_events/game_event_SystemData.Events.RELOAD_INTERFACE.md) (HIGH 100/100) - Game Event

## Affects

- [EA_Window_InteractionRenownTraining.Hide](global_EA_Window_InteractionRenownTraining.Hide.md) (HIGH 100/100) - Global Function
- [EA_Window_InteractionRenownTraining.Show](global_EA_Window_InteractionRenownTraining.Show.md) (HIGH 100/100) - Global Function
- [SystemData.Events.LOADING_END](../../systemdata/fields/systemdata_SystemData.Events.LOADING_END.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_CAREER_CATEGORY_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_CAREER_CATEGORY_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.RELOAD_INTERFACE](../../systemdata/fields/systemdata_SystemData.Events.RELOAD_INTERFACE.md) (HIGH 100/100) - SystemData Field
- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Notes

- Canonical entry built from observed call sites, not from engine source or decompiled definitions.
