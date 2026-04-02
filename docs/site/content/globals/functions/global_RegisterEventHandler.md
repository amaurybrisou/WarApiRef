# RegisterEventHandler

- Category: Global Function
- Confidence level: MEDIUM
- Confidence score: 68/100
- Seen in: 48 addons

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
| Addons seen in | AdvancedPetAssist, AdvancedRenownTrainer, AggroMeter, Aura, AutoMark, BagOMatic, BankArkel, BuffHead |
| Files seen in | `/workspace_addons/AdvancedPetAssist/AdvancedPetAssist.lua:92`, `/workspace_addons/AggroMeter/AggroMeter.lua:5`, `/workspace_addons/Aura/Source/AuraAddon.lua:70`, `/workspace_addons/AutoMark/Source/AutoMark.lua:33`, `/workspace_addons/AutoMark/Source/AutoMark.lua:78`, `/workspace_addons/BankArkel/BankArkel.lua:95`, `/workspace_addons/BuffHead/Core.lua:152`, `/workspace_addons/BuffHead/Core.lua:207` |
| Namespaces detected | RegisterEventHandler |
| Source kinds | lua_calls |
| Example locations | AdvancedPetAssist: AdvancedPetAssist.local.RegisterLoadingEnd, AdvancedPetAssist: RegisterLoadingEnd, AdvancedRenownTrainer: AdvancedRenownTraining.Initialize, AdvancedRenownTrainer: AdvancedRenownTraining.OnReload, AggroMeter: AggroMeter.Initialize, Aura: AuraAddon.OnInitialize |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 281 |
| Global usage count | 281 |
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
- Miracle Grow Remix
- MiracleGrow
- MiracleGrowLight
- PetFixWindow
- PlanB
- Pocket Palette
- PotionBar
- QuickTacticSwitch
- QuickWarReport
- RVAPI_ColorDialog
- RVMOD_Manager
- RoR_SoR
- Shinies
- Targets
- TexturedButtons
- TidyChat
- TidyRoll
- TortallDPSCore
- TurretRange
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

- [LibSlash.RegisterSlashCmd](global_LibSlash.RegisterSlashCmd.md) (HIGH 100/100) - Global Function
- [SystemData.Events.LOADING_END](../../systemdata/fields/systemdata_SystemData.Events.LOADING_END.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.RELOAD_INTERFACE](../../systemdata/fields/systemdata_SystemData.Events.RELOAD_INTERFACE.md) (HIGH 100/100) - SystemData Field
- [WindowRegisterEventHandler](../../window_api/functions/window_WindowRegisterEventHandler.md) (HIGH 100/100) - Window Function

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
- Advanced return analysis: No strong return evidence observed
