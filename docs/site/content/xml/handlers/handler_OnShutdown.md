# OnShutdown

- Type: XML Handler
- Confidence level: HIGH
- Confidence score: 88/100

## Confidence Assessment

- Level: HIGH

- Score: 88/100

- Rationale: Promoted as HIGH confidence because seen in 4 or more addons, referenced by generated docs or reference files, used in event registration or dispatch.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +18 Used in event registration or dispatch: Observed in event-driven engine hooks.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +25 Referenced by generated docs or reference files: The symbol is reinforced outside a single call page.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | CM_ClosetGoblin, Calling, Map, Tome Titan, VerticalMorale |
| Namespaces detected | OnShutdown |
| Source kinds | bindings, xml_handlers |
| Example locations | CM_ClosetGoblin: .OnShutdown, Calling: .OnShutdown, Map: .OnShutdown, Tome Titan: .OnShutdown, VerticalMorale: .OnShutdown |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 6 |
| Global usage count | 0 |
| Local definition count | 0 |
| Documentation references | 1 |
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
| Conflicting signatures | no |
| Conflicting roles | no |
| Wrapper likely | no |
| Never outside local graph | no |
| Local helper only | no |

## Description

XML handler event observed across 5 addons.

## Expected Lua Binding

```lua
function()
```

## Element Types

- none

## Seen In

- CM_ClosetGoblin
- Calling
- Map
- Tome Titan
- VerticalMorale

## Examples

- CM_ClosetGoblin: .OnShutdown -> ClosetGoblinCharacterWindow.OnShutdown
- CM_ClosetGoblin: .OnShutdown -> ClosetGoblinZoneWindow.OnShutdown
- Calling: .OnShutdown -> CallingSetup.Shutdown
- Map: .OnShutdown -> Map.Shutdown
- Tome Titan: .OnShutdown -> TTitan.UI.Shutdown
- VerticalMorale: .OnShutdown -> VerticalMorale.Shutdown

## Related APIs

- [MapDisplay](../element_types/element_MapDisplay.md) (HIGH 100/100) - XML Element Type
- [Window](../element_types/element_Window.md) (HIGH 100/100) - XML Element Type
- [WindowGetAnchor](../../window_api/functions/window_WindowGetAnchor.md) (HIGH 100/100) - Window Function
- [WindowGetOffsetFromParent](../../window_api/functions/window_WindowGetOffsetFromParent.md) (HIGH 100/100) - Window Function
- [WindowUnregisterEventHandler](../../window_api/functions/window_WindowUnregisterEventHandler.md) (HIGH 100/100) - Window Function
- [DoesWindowExist](../../globals/functions/global_DoesWindowExist.md) (HIGH 83/100) - Global Function
- [TextEditBoxGetHistory](../../window_api/functions/window_TextEditBoxGetHistory.md) (HIGH 80/100) - Window Function
- [DestroyWindow](../../globals/functions/global_DestroyWindow.md) (HIGH 75/100) - Global Function
- [UnregisterEventHandler](../../globals/functions/global_UnregisterEventHandler.md) (MEDIUM 68/100) - Global Function
- [Stop](../../events/game_events/game_event_Stop.md) (MEDIUM 43/100) - Game Event

## Used With

- [DestroyWindow](../../globals/functions/global_DestroyWindow.md) (HIGH 75/100) - Global Function
- [UnregisterEventHandler](../../globals/functions/global_UnregisterEventHandler.md) (MEDIUM 68/100) - Global Function

## Affects

- [SystemData.Events.ALL_MODULES_INITIALIZED](../../systemdata/fields/systemdata_SystemData.Events.ALL_MODULES_INITIALIZED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.BATTLEGROUP_MEMBER_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.BATTLEGROUP_MEMBER_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.BATTLEGROUP_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.BATTLEGROUP_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.CHAT_TEXT_ARRIVED](../../systemdata/fields/systemdata_SystemData.Events.CHAT_TEXT_ARRIVED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.CONVERSATION_TEXT_ARRIVED](../../systemdata/fields/systemdata_SystemData.Events.CONVERSATION_TEXT_ARRIVED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.ENTER_WORLD](../../systemdata/fields/systemdata_SystemData.Events.ENTER_WORLD.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.GROUP_ACCEPT_INVITATION](../../systemdata/fields/systemdata_SystemData.Events.GROUP_ACCEPT_INVITATION.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.GROUP_ACCEPT_REFERRAL](../../systemdata/fields/systemdata_SystemData.Events.GROUP_ACCEPT_REFERRAL.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.GROUP_LEAVE](../../systemdata/fields/systemdata_SystemData.Events.GROUP_LEAVE.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.GROUP_PLAYER_ADDED](../../systemdata/fields/systemdata_SystemData.Events.GROUP_PLAYER_ADDED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.GROUP_STATUS_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.GROUP_STATUS_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.GROUP_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.GROUP_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.INTERACT_DEFAULT](../../systemdata/fields/systemdata_SystemData.Events.INTERACT_DEFAULT.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.INTERACT_DONE](../../systemdata/fields/systemdata_SystemData.Events.INTERACT_DONE.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.INTERACT_LEAVE_SCENARIO_QUEUE](../../systemdata/fields/systemdata_SystemData.Events.INTERACT_LEAVE_SCENARIO_QUEUE.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.INTERACT_SHOW_HEALER](../../systemdata/fields/systemdata_SystemData.Events.INTERACT_SHOW_HEALER.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.INTERACT_SHOW_QUEST](../../systemdata/fields/systemdata_SystemData.Events.INTERACT_SHOW_QUEST.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.INTERFACE_RELOADED](../../systemdata/fields/systemdata_SystemData.Events.INTERFACE_RELOADED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.LOADING_END](../../systemdata/fields/systemdata_SystemData.Events.LOADING_END.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.OBJECTIVE_CONTROL_POINTS_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.OBJECTIVE_CONTROL_POINTS_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.OBJECTIVE_OWNER_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.OBJECTIVE_OWNER_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_ABILITY_TOGGLED](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_ABILITY_TOGGLED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_BEGIN_CAST](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_BEGIN_CAST.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_CAREER_CATEGORY_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_CAREER_CATEGORY_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_CAREER_RANK_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_CAREER_RANK_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_CAREER_RESOURCE_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_CAREER_RESOURCE_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_CAST_TIMER_SETBACK](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_CAST_TIMER_SETBACK.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_COMBAT_FLAG_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_COMBAT_FLAG_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_CUR_ACTION_POINTS_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_CUR_ACTION_POINTS_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_DEATH](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_DEATH.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_EFFECTS_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_EFFECTS_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_END_CAST](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_END_CAST.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_HOT_BAR_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_HOT_BAR_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_INVENTORY_SLOT_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_INVENTORY_SLOT_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_MORALE_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_MORALE_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_PET_HEALTH_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_PET_HEALTH_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_POSITION_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_POSITION_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_RENOWN_RANK_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_RENOWN_RANK_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_START_INTERACT_TIMER](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_START_INTERACT_TIMER.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_TARGET_EFFECTS_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_TARGET_EFFECTS_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_TARGET_HIT_POINTS_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_TARGET_HIT_POINTS_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_TARGET_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_TARGET_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_ZONE_CHANGED](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_ZONE_CHANGED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PUBLIC_QUEST_ADDED](../../systemdata/fields/systemdata_SystemData.Events.PUBLIC_QUEST_ADDED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PUBLIC_QUEST_INFO_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.PUBLIC_QUEST_INFO_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PUBLIC_QUEST_REMOVED](../../systemdata/fields/systemdata_SystemData.Events.PUBLIC_QUEST_REMOVED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PUBLIC_QUEST_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.PUBLIC_QUEST_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.RELOAD_INTERFACE](../../systemdata/fields/systemdata_SystemData.Events.RELOAD_INTERFACE.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.R_BUTTON_DOWN_PROCESSED](../../systemdata/fields/systemdata_SystemData.Events.R_BUTTON_DOWN_PROCESSED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.SCENARIO_BEGIN](../../systemdata/fields/systemdata_SystemData.Events.SCENARIO_BEGIN.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.SCENARIO_END](../../systemdata/fields/systemdata_SystemData.Events.SCENARIO_END.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.SCENARIO_GROUP_JOIN](../../systemdata/fields/systemdata_SystemData.Events.SCENARIO_GROUP_JOIN.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.SCENARIO_GROUP_LEAVE](../../systemdata/fields/systemdata_SystemData.Events.SCENARIO_GROUP_LEAVE.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.SCENARIO_INSTANCE_CANCEL](../../systemdata/fields/systemdata_SystemData.Events.SCENARIO_INSTANCE_CANCEL.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.SCENARIO_INSTANCE_JOIN_NOW](../../systemdata/fields/systemdata_SystemData.Events.SCENARIO_INSTANCE_JOIN_NOW.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.SCENARIO_PLAYERS_LIST_GROUPS_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.SCENARIO_PLAYERS_LIST_GROUPS_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.SHOW_ALERT_TEXT](../../systemdata/fields/systemdata_SystemData.Events.SHOW_ALERT_TEXT.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.SPELL_CAST_CANCEL](../../systemdata/fields/systemdata_SystemData.Events.SPELL_CAST_CANCEL.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.UPDATE_PROCESSED](../../systemdata/fields/systemdata_SystemData.Events.UPDATE_PROCESSED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.USER_SETTINGS_CHANGED](../../systemdata/fields/systemdata_SystemData.Events.USER_SETTINGS_CHANGED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.WORLD_OBJ_COMBAT_EVENT](../../systemdata/fields/systemdata_SystemData.Events.WORLD_OBJ_COMBAT_EVENT.md) (HIGH 100/100) - SystemData Field

## Notes

- Expected callback signature inferred from common WAR XML handler conventions (MEDIUM confidence).
