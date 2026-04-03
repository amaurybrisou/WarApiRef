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

- [SystemData.Events.ALL_MODULES_INITIALIZED](../../events/game_events/game_event_SystemData.Events.ALL_MODULES_INITIALIZED.md) (HIGH 100/100) - Game Event
- [SystemData.Events.AUCTION_BID_RESULT_RECEIVED](../../events/game_events/game_event_SystemData.Events.AUCTION_BID_RESULT_RECEIVED.md) (HIGH 100/100) - Game Event
- [SystemData.Events.AUCTION_INIT_RECEIVED](../../events/game_events/game_event_SystemData.Events.AUCTION_INIT_RECEIVED.md) (HIGH 100/100) - Game Event
- [SystemData.Events.AUCTION_SEARCH_RESULT_RECEIVED](../../events/game_events/game_event_SystemData.Events.AUCTION_SEARCH_RESULT_RECEIVED.md) (HIGH 100/100) - Game Event
- [SystemData.Events.BATTLEGROUP_ACCEPT_INVITATION](../../events/game_events/game_event_SystemData.Events.BATTLEGROUP_ACCEPT_INVITATION.md) (HIGH 100/100) - Game Event
- [SystemData.Events.BATTLEGROUP_MEMBER_UPDATED](../../events/game_events/game_event_SystemData.Events.BATTLEGROUP_MEMBER_UPDATED.md) (HIGH 100/100) - Game Event
- [SystemData.Events.BATTLEGROUP_UPDATED](../../events/game_events/game_event_SystemData.Events.BATTLEGROUP_UPDATED.md) (HIGH 100/100) - Game Event
- [SystemData.Events.CAMPAIGN_ZONE_UPDATED](../../events/game_events/game_event_SystemData.Events.CAMPAIGN_ZONE_UPDATED.md) (HIGH 100/100) - Game Event
- [SystemData.Events.CHAT_TEXT_ARRIVED](../../events/game_events/game_event_SystemData.Events.CHAT_TEXT_ARRIVED.md) (HIGH 100/100) - Game Event
- [SystemData.Events.CITY_SCENARIO_BEGIN](../../events/game_events/game_event_SystemData.Events.CITY_SCENARIO_BEGIN.md) (HIGH 100/100) - Game Event
- [SystemData.Events.CITY_SCENARIO_END](../../events/game_events/game_event_SystemData.Events.CITY_SCENARIO_END.md) (HIGH 100/100) - Game Event
- [SystemData.Events.CITY_SCENARIO_UPDATE_POINTS](../../events/game_events/game_event_SystemData.Events.CITY_SCENARIO_UPDATE_POINTS.md) (HIGH 100/100) - Game Event
- [SystemData.Events.CUSTOM_UI_SCALE_CHANGED](../../events/game_events/game_event_SystemData.Events.CUSTOM_UI_SCALE_CHANGED.md) (HIGH 100/100) - Game Event
- [SystemData.Events.END_ITEM_ENHANCEMENT](../../events/game_events/game_event_SystemData.Events.END_ITEM_ENHANCEMENT.md) (HIGH 100/100) - Game Event
- [SystemData.Events.ENTER_WORLD](../../events/game_events/game_event_SystemData.Events.ENTER_WORLD.md) (HIGH 100/100) - Game Event
- [SystemData.Events.GROUP_ACCEPT_INVITATION](../../events/game_events/game_event_SystemData.Events.GROUP_ACCEPT_INVITATION.md) (HIGH 100/100) - Game Event
- [SystemData.Events.GROUP_EFFECTS_UPDATED](../../events/game_events/game_event_SystemData.Events.GROUP_EFFECTS_UPDATED.md) (HIGH 100/100) - Game Event
- [SystemData.Events.GROUP_LEAVE](../../events/game_events/game_event_SystemData.Events.GROUP_LEAVE.md) (HIGH 100/100) - Game Event
- [SystemData.Events.GROUP_PLAYER_ADDED](../../events/game_events/game_event_SystemData.Events.GROUP_PLAYER_ADDED.md) (HIGH 100/100) - Game Event
- [SystemData.Events.GROUP_SETTINGS_UPDATED](../../events/game_events/game_event_SystemData.Events.GROUP_SETTINGS_UPDATED.md) (HIGH 100/100) - Game Event
- [SystemData.Events.GROUP_SET_LEADER](../../events/game_events/game_event_SystemData.Events.GROUP_SET_LEADER.md) (HIGH 100/100) - Game Event
- [SystemData.Events.GROUP_STATUS_UPDATED](../../events/game_events/game_event_SystemData.Events.GROUP_STATUS_UPDATED.md) (HIGH 100/100) - Game Event
- [SystemData.Events.GROUP_UPDATED](../../events/game_events/game_event_SystemData.Events.GROUP_UPDATED.md) (HIGH 100/100) - Game Event
- [SystemData.Events.INTERACT_DONE](../../events/game_events/game_event_SystemData.Events.INTERACT_DONE.md) (HIGH 100/100) - Game Event
- [SystemData.Events.INTERACT_LEAVE_SCENARIO_QUEUE](../../events/game_events/game_event_SystemData.Events.INTERACT_LEAVE_SCENARIO_QUEUE.md) (HIGH 100/100) - Game Event
- [SystemData.Events.INTERACT_LOOT_ROLL_FIRST_ITEM](../../events/game_events/game_event_SystemData.Events.INTERACT_LOOT_ROLL_FIRST_ITEM.md) (HIGH 100/100) - Game Event
- [SystemData.Events.INTERACT_OPEN_BANK](../../events/game_events/game_event_SystemData.Events.INTERACT_OPEN_BANK.md) (HIGH 100/100) - Game Event
- [SystemData.Events.INTERACT_SHOW_LOOT_ROLL_DATA](../../events/game_events/game_event_SystemData.Events.INTERACT_SHOW_LOOT_ROLL_DATA.md) (HIGH 100/100) - Game Event
- [SystemData.Events.INTERFACE_RELOADED](../../events/game_events/game_event_SystemData.Events.INTERFACE_RELOADED.md) (HIGH 100/100) - Game Event
- [SystemData.Events.ITEM_SET_DATA_ARRIVED](../../events/game_events/game_event_SystemData.Events.ITEM_SET_DATA_ARRIVED.md) (HIGH 100/100) - Game Event
- [SystemData.Events.LOADING_BEGIN](../../events/game_events/game_event_SystemData.Events.LOADING_BEGIN.md) (HIGH 100/100) - Game Event
- [SystemData.Events.LOADING_END](../../events/game_events/game_event_SystemData.Events.LOADING_END.md) (HIGH 100/100) - Game Event
- [SystemData.Events.LOG_OUT](../../events/game_events/game_event_SystemData.Events.LOG_OUT.md) (HIGH 100/100) - Game Event
- [SystemData.Events.L_BUTTON_DOWN_PROCESSED](../../events/game_events/game_event_SystemData.Events.L_BUTTON_DOWN_PROCESSED.md) (HIGH 100/100) - Game Event
- [SystemData.Events.L_BUTTON_UP_PROCESSED](../../events/game_events/game_event_SystemData.Events.L_BUTTON_UP_PROCESSED.md) (HIGH 100/100) - Game Event
- [SystemData.Events.MACROS_LOADED](../../events/game_events/game_event_SystemData.Events.MACROS_LOADED.md) (HIGH 100/100) - Game Event
- [SystemData.Events.MACRO_UPDATED](../../events/game_events/game_event_SystemData.Events.MACRO_UPDATED.md) (HIGH 100/100) - Game Event
- [SystemData.Events.M_BUTTON_DOWN_PROCESSED](../../events/game_events/game_event_SystemData.Events.M_BUTTON_DOWN_PROCESSED.md) (HIGH 100/100) - Game Event
- [SystemData.Events.PLAYER_BATTLE_LEVEL_UPDATED](../../events/game_events/game_event_SystemData.Events.PLAYER_BATTLE_LEVEL_UPDATED.md) (HIGH 100/100) - Game Event
- [SystemData.Events.PLAYER_BEGIN_CAST](../../events/game_events/game_event_SystemData.Events.PLAYER_BEGIN_CAST.md) (HIGH 100/100) - Game Event
- [SystemData.Events.PLAYER_CAREER_CATEGORY_UPDATED](../../events/game_events/game_event_SystemData.Events.PLAYER_CAREER_CATEGORY_UPDATED.md) (HIGH 100/100) - Game Event
- [SystemData.Events.PLAYER_CAREER_LINE_UPDATED](../../events/game_events/game_event_SystemData.Events.PLAYER_CAREER_LINE_UPDATED.md) (HIGH 100/100) - Game Event
- [SystemData.Events.PLAYER_CAREER_RANK_UPDATED](../../events/game_events/game_event_SystemData.Events.PLAYER_CAREER_RANK_UPDATED.md) (HIGH 100/100) - Game Event
- [SystemData.Events.PLAYER_CAREER_RESOURCE_UPDATED](../../events/game_events/game_event_SystemData.Events.PLAYER_CAREER_RESOURCE_UPDATED.md) (HIGH 100/100) - Game Event
- [SystemData.Events.PLAYER_CAST_TIMER_SETBACK](../../events/game_events/game_event_SystemData.Events.PLAYER_CAST_TIMER_SETBACK.md) (HIGH 100/100) - Game Event
- [SystemData.Events.PLAYER_COMBAT_FLAG_UPDATED](../../events/game_events/game_event_SystemData.Events.PLAYER_COMBAT_FLAG_UPDATED.md) (HIGH 100/100) - Game Event
- [SystemData.Events.PLAYER_CURRENCY_SLOT_UPDATED](../../events/game_events/game_event_SystemData.Events.PLAYER_CURRENCY_SLOT_UPDATED.md) (HIGH 100/100) - Game Event
- [SystemData.Events.PLAYER_CUR_ACTION_POINTS_UPDATED](../../events/game_events/game_event_SystemData.Events.PLAYER_CUR_ACTION_POINTS_UPDATED.md) (HIGH 100/100) - Game Event
- [SystemData.Events.PLAYER_CUR_HIT_POINTS_UPDATED](../../events/game_events/game_event_SystemData.Events.PLAYER_CUR_HIT_POINTS_UPDATED.md) (HIGH 100/100) - Game Event
- [SystemData.Events.PLAYER_DEATH](../../events/game_events/game_event_SystemData.Events.PLAYER_DEATH.md) (HIGH 100/100) - Game Event
- [SystemData.Events.PLAYER_DEATH_CLEARED](../../events/game_events/game_event_SystemData.Events.PLAYER_DEATH_CLEARED.md) (HIGH 100/100) - Game Event
- [SystemData.Events.PLAYER_EFFECTS_UPDATED](../../events/game_events/game_event_SystemData.Events.PLAYER_EFFECTS_UPDATED.md) (HIGH 100/100) - Game Event
- [SystemData.Events.PLAYER_END_CAST](../../events/game_events/game_event_SystemData.Events.PLAYER_END_CAST.md) (HIGH 100/100) - Game Event
- [SystemData.Events.PLAYER_EQUIPMENT_SLOT_UPDATED](../../events/game_events/game_event_SystemData.Events.PLAYER_EQUIPMENT_SLOT_UPDATED.md) (HIGH 100/100) - Game Event
- [SystemData.Events.PLAYER_GROUP_LEADER_STATUS_UPDATED](../../events/game_events/game_event_SystemData.Events.PLAYER_GROUP_LEADER_STATUS_UPDATED.md) (HIGH 100/100) - Game Event
- [SystemData.Events.PLAYER_HOT_BAR_UPDATED](../../events/game_events/game_event_SystemData.Events.PLAYER_HOT_BAR_UPDATED.md) (HIGH 100/100) - Game Event
- [SystemData.Events.PLAYER_INVENTORY_SLOT_UPDATED](../../events/game_events/game_event_SystemData.Events.PLAYER_INVENTORY_SLOT_UPDATED.md) (HIGH 100/100) - Game Event
- [SystemData.Events.PLAYER_MAX_ACTION_POINTS_UPDATED](../../events/game_events/game_event_SystemData.Events.PLAYER_MAX_ACTION_POINTS_UPDATED.md) (HIGH 100/100) - Game Event
- [SystemData.Events.PLAYER_MAX_HIT_POINTS_UPDATED](../../events/game_events/game_event_SystemData.Events.PLAYER_MAX_HIT_POINTS_UPDATED.md) (HIGH 100/100) - Game Event
- [SystemData.Events.PLAYER_MORALE_UPDATED](../../events/game_events/game_event_SystemData.Events.PLAYER_MORALE_UPDATED.md) (HIGH 100/100) - Game Event
- [SystemData.Events.PLAYER_PET_HEALTH_UPDATED](../../events/game_events/game_event_SystemData.Events.PLAYER_PET_HEALTH_UPDATED.md) (HIGH 100/100) - Game Event
- [SystemData.Events.PLAYER_PET_TARGET_UPDATED](../../events/game_events/game_event_SystemData.Events.PLAYER_PET_TARGET_UPDATED.md) (HIGH 100/100) - Game Event
- [SystemData.Events.PLAYER_PET_UPDATED](../../events/game_events/game_event_SystemData.Events.PLAYER_PET_UPDATED.md) (HIGH 100/100) - Game Event
- [SystemData.Events.PLAYER_POSITION_UPDATED](../../events/game_events/game_event_SystemData.Events.PLAYER_POSITION_UPDATED.md) (HIGH 100/100) - Game Event
- [SystemData.Events.PLAYER_RENOWN_RANK_UPDATED](../../events/game_events/game_event_SystemData.Events.PLAYER_RENOWN_RANK_UPDATED.md) (HIGH 100/100) - Game Event
- [SystemData.Events.PLAYER_RENOWN_UPDATED](../../events/game_events/game_event_SystemData.Events.PLAYER_RENOWN_UPDATED.md) (HIGH 100/100) - Game Event
- [SystemData.Events.PLAYER_RVR_FLAG_UPDATED](../../events/game_events/game_event_SystemData.Events.PLAYER_RVR_FLAG_UPDATED.md) (HIGH 100/100) - Game Event
- [SystemData.Events.PLAYER_START_INTERACT_TIMER](../../events/game_events/game_event_SystemData.Events.PLAYER_START_INTERACT_TIMER.md) (HIGH 100/100) - Game Event
- [SystemData.Events.PLAYER_TARGET_EFFECTS_UPDATED](../../events/game_events/game_event_SystemData.Events.PLAYER_TARGET_EFFECTS_UPDATED.md) (HIGH 100/100) - Game Event
- [SystemData.Events.PLAYER_TARGET_UPDATED](../../events/game_events/game_event_SystemData.Events.PLAYER_TARGET_UPDATED.md) (HIGH 100/100) - Game Event
- [SystemData.Events.PLAYER_ZONE_CHANGED](../../events/game_events/game_event_SystemData.Events.PLAYER_ZONE_CHANGED.md) (HIGH 100/100) - Game Event
- [SystemData.Events.PUBLIC_QUEST_ADDED](../../events/game_events/game_event_SystemData.Events.PUBLIC_QUEST_ADDED.md) (HIGH 100/100) - Game Event
- [SystemData.Events.PUBLIC_QUEST_INFO_UPDATED](../../events/game_events/game_event_SystemData.Events.PUBLIC_QUEST_INFO_UPDATED.md) (HIGH 100/100) - Game Event
- [SystemData.Events.PUBLIC_QUEST_REMOVED](../../events/game_events/game_event_SystemData.Events.PUBLIC_QUEST_REMOVED.md) (HIGH 100/100) - Game Event
- [SystemData.Events.PUBLIC_QUEST_UPDATED](../../events/game_events/game_event_SystemData.Events.PUBLIC_QUEST_UPDATED.md) (HIGH 100/100) - Game Event
- [SystemData.Events.RELEASE_CORPSE](../../events/game_events/game_event_SystemData.Events.RELEASE_CORPSE.md) (HIGH 100/100) - Game Event
- [SystemData.Events.RELOAD_INTERFACE](../../events/game_events/game_event_SystemData.Events.RELOAD_INTERFACE.md) (HIGH 100/100) - Game Event
- [SystemData.Events.R_BUTTON_DOWN_PROCESSED](../../events/game_events/game_event_SystemData.Events.R_BUTTON_DOWN_PROCESSED.md) (HIGH 100/100) - Game Event
- [SystemData.Events.SCENARIO_BEGIN](../../events/game_events/game_event_SystemData.Events.SCENARIO_BEGIN.md) (HIGH 100/100) - Game Event
- [SystemData.Events.SCENARIO_END](../../events/game_events/game_event_SystemData.Events.SCENARIO_END.md) (HIGH 100/100) - Game Event
- [SystemData.Events.SCENARIO_GROUP_JOIN](../../events/game_events/game_event_SystemData.Events.SCENARIO_GROUP_JOIN.md) (HIGH 100/100) - Game Event
- [SystemData.Events.SCENARIO_GROUP_LEAVE](../../events/game_events/game_event_SystemData.Events.SCENARIO_GROUP_LEAVE.md) (HIGH 100/100) - Game Event
- [SystemData.Events.SCENARIO_INSTANCE_JOIN_NOW](../../events/game_events/game_event_SystemData.Events.SCENARIO_INSTANCE_JOIN_NOW.md) (HIGH 100/100) - Game Event
- [SystemData.Events.SCENARIO_PLAYERS_LIST_GROUPS_UPDATED](../../events/game_events/game_event_SystemData.Events.SCENARIO_PLAYERS_LIST_GROUPS_UPDATED.md) (HIGH 100/100) - Game Event
- [SystemData.Events.SCENARIO_PLAYERS_LIST_RESERVATIONS_UPDATED](../../events/game_events/game_event_SystemData.Events.SCENARIO_PLAYERS_LIST_RESERVATIONS_UPDATED.md) (HIGH 100/100) - Game Event
- [SystemData.Events.SCENARIO_PLAYERS_LIST_STATS_UPDATED](../../events/game_events/game_event_SystemData.Events.SCENARIO_PLAYERS_LIST_STATS_UPDATED.md) (HIGH 100/100) - Game Event
- [SystemData.Events.SCENARIO_PLAYERS_LIST_UPDATED](../../events/game_events/game_event_SystemData.Events.SCENARIO_PLAYERS_LIST_UPDATED.md) (HIGH 100/100) - Game Event
- [SystemData.Events.SCENARIO_PLAYER_HITS_UPDATED](../../events/game_events/game_event_SystemData.Events.SCENARIO_PLAYER_HITS_UPDATED.md) (HIGH 100/100) - Game Event
- [SystemData.Events.SCENARIO_POST_MODE](../../events/game_events/game_event_SystemData.Events.SCENARIO_POST_MODE.md) (HIGH 100/100) - Game Event
- [SystemData.Events.SCENARIO_SHOW_JOIN_PROMPT](../../events/game_events/game_event_SystemData.Events.SCENARIO_SHOW_JOIN_PROMPT.md) (HIGH 100/100) - Game Event
- [SystemData.Events.SCENARIO_STOP_UPDATING_PLAYERS_STATS](../../events/game_events/game_event_SystemData.Events.SCENARIO_STOP_UPDATING_PLAYERS_STATS.md) (HIGH 100/100) - Game Event
- [SystemData.Events.SCENARIO_UPDATE_POINTS](../../events/game_events/game_event_SystemData.Events.SCENARIO_UPDATE_POINTS.md) (HIGH 100/100) - Game Event
- [SystemData.Events.SHOW_ALERT_TEXT](../../events/game_events/game_event_SystemData.Events.SHOW_ALERT_TEXT.md) (HIGH 100/100) - Game Event
- [SystemData.Events.UPDATE_ITEM_ENHANCEMENT](../../events/game_events/game_event_SystemData.Events.UPDATE_ITEM_ENHANCEMENT.md) (HIGH 100/100) - Game Event
- [SystemData.Events.UPDATE_PROCESSED](../../events/game_events/game_event_SystemData.Events.UPDATE_PROCESSED.md) (HIGH 100/100) - Game Event
- [SystemData.Events.USER_SETTINGS_CHANGED](../../events/game_events/game_event_SystemData.Events.USER_SETTINGS_CHANGED.md) (HIGH 100/100) - Game Event
- [SystemData.Events.VISIBLE_EQUIPMENT_UPDATED](../../events/game_events/game_event_SystemData.Events.VISIBLE_EQUIPMENT_UPDATED.md) (HIGH 100/100) - Game Event
- [SystemData.Events.WORLD_OBJ_COMBAT_EVENT](../../events/game_events/game_event_SystemData.Events.WORLD_OBJ_COMBAT_EVENT.md) (HIGH 100/100) - Game Event
- [SystemData.Events.WORLD_OBJ_INFLUENCE_GAINED](../../events/game_events/game_event_SystemData.Events.WORLD_OBJ_INFLUENCE_GAINED.md) (HIGH 100/100) - Game Event
- [SystemData.Events.WORLD_OBJ_RENOWN_GAINED](../../events/game_events/game_event_SystemData.Events.WORLD_OBJ_RENOWN_GAINED.md) (HIGH 100/100) - Game Event
- [SystemData.Events.WORLD_OBJ_XP_GAINED](../../events/game_events/game_event_SystemData.Events.WORLD_OBJ_XP_GAINED.md) (HIGH 100/100) - Game Event
- [SystemData.Events.PLAYER_BANK_SLOT_UPDATED](../../events/game_events/game_event_SystemData.Events.PLAYER_BANK_SLOT_UPDATED.md) (HIGH 98/100) - Game Event
- [SystemData.Events.PLAYER_CRAFTING_SLOT_UPDATED](../../events/game_events/game_event_SystemData.Events.PLAYER_CRAFTING_SLOT_UPDATED.md) (MEDIUM 63/100) - Game Event
- [SystemData.Events.PLAYER_MONEY_UPDATED](../../events/game_events/game_event_SystemData.Events.PLAYER_MONEY_UPDATED.md) (MEDIUM 63/100) - Game Event

## Used With

- [LibSlash.RegisterSlashCmd](global_LibSlash.RegisterSlashCmd.md) (HIGH 100/100) - Global Function
- [SystemData.Events.INTERFACE_RELOADED](../../events/game_events/game_event_SystemData.Events.INTERFACE_RELOADED.md) (HIGH 100/100) - Game Event
- [SystemData.Events.LOADING_END](../../systemdata/fields/systemdata_SystemData.Events.LOADING_END.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.LOADING_END](../../events/game_events/game_event_SystemData.Events.LOADING_END.md) (HIGH 100/100) - Game Event
- [SystemData.Events.RELOAD_INTERFACE](../../events/game_events/game_event_SystemData.Events.RELOAD_INTERFACE.md) (HIGH 100/100) - Game Event
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
