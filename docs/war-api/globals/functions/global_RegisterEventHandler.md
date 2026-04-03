# RegisterEventHandler

- Category: Global Function
- Confidence level: HIGH
- Confidence score: 93/100
- Seen in: 31 addons

## Confidence Assessment

- Level: HIGH

- Score: 93/100

- Rationale: Promoted as HIGH confidence because seen in 4 or more addons, called globally with no local definition, used in event registration or dispatch.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +18 Used in event registration or dispatch: Observed in event-driven engine hooks.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +10 Argument pattern is consistent: Observed argument positions remain stable.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | AdvancedPetAssist, AdvancedRenownTrainer, AggroMeter, Aura, AutoMark, BagOMatic, BankArkel, BuffHead |
| Files seen in | `/workspace/data/raw/AdvancedPetAssist/AdvancedPetAssist.lua:92`, `/workspace/data/raw/AggroMeter/AggroMeter.lua:5`, `/workspace/data/raw/Aura/Source/AuraAddon.lua:70`, `/workspace/data/raw/AutoMark/Source/AutoMark.lua:33`, `/workspace/data/raw/AutoMark/Source/AutoMark.lua:78`, `/workspace/data/raw/BankArkel/BankArkel.lua:95`, `/workspace/data/raw/BuffHead/Core.lua:152`, `/workspace/data/raw/BuffHead/Core.lua:207` |
| Namespaces detected | RegisterEventHandler |
| Source kinds | lua_calls |
| Example locations | AdvancedPetAssist: AdvancedPetAssist.local.RegisterLoadingEnd, AdvancedPetAssist: RegisterLoadingEnd, AdvancedRenownTrainer: AdvancedRenownTraining.Initialize, AdvancedRenownTrainer: AdvancedRenownTraining.OnReload, AggroMeter: AggroMeter.Initialize, Aura: AuraAddon.OnInitialize |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 214 |
| Global usage count | 214 |
| Local definition count | 0 |
| Documentation references | 0 |
| Initialization flow references | 0 |
| Known engine namespace | no |
| Default UI presence | no |
| Event binding presence | yes |
| Observed in XML and Lua | no |
| Consistent role | yes |
| Consistent arguments | yes |
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

## Signature (inferred)

```lua
RegisterEventHandler(eventId, handlerName)
```

## Description

Observed registering global runtime handlers against shared event identifiers.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| eventId | Observed as a SystemData or runtime event identifier. | Observed values: S.combatLogEventId, SystemData.Events.ALL_MODULES_INITIALIZED, SystemData.Events.AUCTION_INIT_RECEIVED |
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
- Enemy
- Killer
- LibGuard
- LibWBToggler
- MiracleGrowLight
- PartyCast
- PlanB
- Pocket Palette
- PotionBar
- RoR_SoR
- Shinies
- TexturedButtons
- TidyChat
- TidyRoll
- TurretRange
- WSCT
- WarBoard
- WhoHealedMe
- WoH-Reticle
- followTheLeader

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
- [WindowRegisterEventHandler](../../window_api/functions/window_WindowRegisterEventHandler.md) (HIGH 100/100) - Window Function
- [CreateWindow](global_CreateWindow.md) (HIGH 75/100) - Global Function

## Triggered By

- [OnInitialize](../../xml/handlers/handler_OnInitialize.md) (HIGH 100/100) - XML Event
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
