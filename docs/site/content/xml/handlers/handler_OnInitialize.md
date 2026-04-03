# OnInitialize

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
| Addons seen in | AdvancedRenownTrainer, AuctionStats, AutoBand, AutoSalvage, BlackBook, CM_ClosetGoblin, CaVES, Calling |
| Namespaces detected | OnInitialize |
| Source kinds | bindings, xml_handlers |
| Example locations | AdvancedRenownTrainer: .OnInitialize, AuctionStats: .OnInitialize, AutoBand: .OnInitialize, AutoSalvage: .OnInitialize, BlackBook: .OnInitialize, CM_ClosetGoblin: .OnInitialize |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 31 |
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

XML handler event observed across 25 addons.

## Expected Lua Binding

```lua
function()
```

## Element Types

- none

## Seen In

- AdvancedRenownTrainer
- AuctionStats
- AutoBand
- AutoSalvage
- BlackBook
- CM_ClosetGoblin
- CaVES
- Calling
- DPSMeter
- EA_ScenarioGroupWindow
- FastInteract
- LoyalPet
- MapMonster
- Minmap
- Miracle Grow Remix
- MiracleGrow
- RandomMount
- TastyButtons
- TokenMachine
- Tome Titan
- WSCT
- WhoHealedMe
- XpStatus+G
- alertMod
- wbLeadHelper

## Examples

- AdvancedRenownTrainer: .OnInitialize -> AdvancedRenownTraining.Initialize
- AuctionStats: .OnInitialize -> AuctionStats.InitOptionsWindow
- AutoBand: .OnInitialize -> AutoBand.OnCopyLinkInitialize
- AutoSalvage: .OnInitialize -> AutoSalvage.OptionsWindow.OnInitialize
- BlackBook: .OnInitialize -> BlackBookWindow.Initialize
- CM_ClosetGoblin: .OnInitialize -> ClosetGoblinOptionWindow.OnInitialize

## Related APIs

- [Button](../element_types/element_Button.md) (HIGH 100/100) - XML Element Type
- [ButtonSetCheckButtonFlag](../../window_api/functions/window_ButtonSetCheckButtonFlag.md) (HIGH 100/100) - Window Function
- [ButtonSetPressedFlag](../../window_api/functions/window_ButtonSetPressedFlag.md) (HIGH 100/100) - Window Function
- [ButtonSetStayDownFlag](../../window_api/functions/window_ButtonSetStayDownFlag.md) (HIGH 100/100) - Window Function
- [ButtonSetText](../../window_api/functions/window_ButtonSetText.md) (HIGH 100/100) - Window Function
- [ComboBoxAddMenuItem](../../window_api/functions/window_ComboBoxAddMenuItem.md) (HIGH 100/100) - Window Function
- [ComboBoxClearMenuItems](../../window_api/functions/window_ComboBoxClearMenuItems.md) (HIGH 100/100) - Window Function
- [ComboBoxSetSelectedMenuItem](../../window_api/functions/window_ComboBoxSetSelectedMenuItem.md) (HIGH 100/100) - Window Function
- [DynamicImageSetTexture](../../window_api/functions/window_DynamicImageSetTexture.md) (HIGH 100/100) - Window Function
- [EA_ChatWindow.Print](../../globals/functions/global_EA_ChatWindow.Print.md) (HIGH 100/100) - Global Function
- [InterfaceCore.GetScale](../../globals/functions/global_InterfaceCore.GetScale.md) (HIGH 100/100) - Global Function
- [LabelSetFont](../../window_api/functions/window_LabelSetFont.md) (HIGH 100/100) - Window Function
- [LabelSetText](../../window_api/functions/window_LabelSetText.md) (HIGH 100/100) - Window Function
- [LabelSetTextColor](../../window_api/functions/window_LabelSetTextColor.md) (HIGH 100/100) - Window Function
- [LayoutEditor.RegisterEditCallback](../../window_api/functions/window_LayoutEditor.RegisterEditCallback.md) (HIGH 100/100) - Window Function
- [LayoutEditor.RegisterWindow](../../window_api/functions/window_LayoutEditor.RegisterWindow.md) (HIGH 100/100) - Window Function
- [LibSlash.IsSlashCmdRegistered](../../globals/functions/global_LibSlash.IsSlashCmdRegistered.md) (HIGH 100/100) - Global Function
- [LibSlash.RegisterWSlashCmd](../../globals/functions/global_LibSlash.RegisterWSlashCmd.md) (HIGH 100/100) - Global Function
- [StatusBarSetMaximumValue](../../window_api/functions/window_StatusBarSetMaximumValue.md) (HIGH 100/100) - Window Function
- [Window](../element_types/element_Window.md) (HIGH 100/100) - XML Element Type
- [WindowAddAnchor](../../window_api/functions/window_WindowAddAnchor.md) (HIGH 100/100) - Window Function
- [WindowClearAnchors](../../window_api/functions/window_WindowClearAnchors.md) (HIGH 100/100) - Window Function
- [WindowGetAlpha](../../window_api/functions/window_WindowGetAlpha.md) (HIGH 100/100) - Window Function
- [WindowGetAnchorCount](../../window_api/functions/window_WindowGetAnchorCount.md) (HIGH 100/100) - Window Function
- [WindowGetScale](../../window_api/functions/window_WindowGetScale.md) (HIGH 100/100) - Window Function
- [WindowRegisterCoreEventHandler](../../window_api/functions/window_WindowRegisterCoreEventHandler.md) (HIGH 100/100) - Window Function
- [WindowRegisterEventHandler](../../window_api/functions/window_WindowRegisterEventHandler.md) (HIGH 100/100) - Window Function
- [WindowSetAlpha](../../window_api/functions/window_WindowSetAlpha.md) (HIGH 100/100) - Window Function
- [WindowSetDimensions](../../window_api/functions/window_WindowSetDimensions.md) (HIGH 100/100) - Window Function
- [WindowSetFontAlpha](../../window_api/functions/window_WindowSetFontAlpha.md) (HIGH 100/100) - Window Function
- [WindowSetParent](../../window_api/functions/window_WindowSetParent.md) (HIGH 100/100) - Window Function
- [WindowSetScale](../../window_api/functions/window_WindowSetScale.md) (HIGH 100/100) - Window Function
- [WindowSetShowing](../../window_api/functions/window_WindowSetShowing.md) (HIGH 100/100) - Window Function
- [WindowSetTintColor](../../window_api/functions/window_WindowSetTintColor.md) (HIGH 100/100) - Window Function
- [LibSlash.RegisterSlashCmd](../../globals/functions/global_LibSlash.RegisterSlashCmd.md) (HIGH 85/100) - Global Function
- [DoesWindowExist](../../globals/functions/global_DoesWindowExist.md) (HIGH 83/100) - Global Function
- [GetIconData](../../globals/functions/global_GetIconData.md) (HIGH 83/100) - Global Function
- [CreateWindow](../../globals/functions/global_CreateWindow.md) (HIGH 75/100) - Global Function
- [CreateWindowFromTemplate](../../globals/functions/global_CreateWindowFromTemplate.md) (HIGH 75/100) - Global Function
- [wstring.format](../../globals/functions/global_wstring.format.md) (HIGH 75/100) - Global Function
- [wstring.sub](../../globals/functions/global_wstring.sub.md) (HIGH 75/100) - Global Function
- [GuildWindow.IsPlayerInAGuild](../../globals/functions/global_GuildWindow.IsPlayerInAGuild.md) (HIGH 71/100) - Global Function
- [RegisterEventHandler](../../globals/functions/global_RegisterEventHandler.md) (MEDIUM 68/100) - Global Function
- [UnregisterEventHandler](../../globals/functions/global_UnregisterEventHandler.md) (MEDIUM 68/100) - Global Function

## Used With

- [ButtonSetText](../../window_api/functions/window_ButtonSetText.md) (HIGH 100/100) - Window Function
- [EA_ChatWindow.Print](../../globals/functions/global_EA_ChatWindow.Print.md) (HIGH 100/100) - Global Function
- [LayoutEditor.RegisterWindow](../../window_api/functions/window_LayoutEditor.RegisterWindow.md) (HIGH 100/100) - Window Function
- [SystemData.Events.LOADING_END](../../systemdata/fields/systemdata_SystemData.Events.LOADING_END.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.RELOAD_INTERFACE](../../systemdata/fields/systemdata_SystemData.Events.RELOAD_INTERFACE.md) (HIGH 100/100) - SystemData Field
- [WindowSetShowing](../../window_api/functions/window_WindowSetShowing.md) (HIGH 100/100) - Window Function
- [LibSlash.RegisterSlashCmd](../../globals/functions/global_LibSlash.RegisterSlashCmd.md) (HIGH 85/100) - Global Function
- [CreateWindow](../../globals/functions/global_CreateWindow.md) (HIGH 75/100) - Global Function
- [RegisterEventHandler](../../globals/functions/global_RegisterEventHandler.md) (MEDIUM 68/100) - Global Function

## Affects

- [GameData.CityId.CHAOS](../../gamedata/fields/gamedata_GameData.CityId.CHAOS.md) (HIGH 100/100) - GameData Field
- [GameData.CityId.EMPIRE](../../gamedata/fields/gamedata_GameData.CityId.EMPIRE.md) (HIGH 100/100) - GameData Field
- [GameData.Player.name](../../gamedata/fields/gamedata_GameData.Player.name.md) (HIGH 100/100) - GameData Field
- [SystemData.ChatLogFilters.SHOUT](../../systemdata/fields/systemdata_SystemData.ChatLogFilters.SHOUT.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.ALL_MODULES_INITIALIZED](../../systemdata/fields/systemdata_SystemData.Events.ALL_MODULES_INITIALIZED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.CAMPAIGN_PAIRING_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.CAMPAIGN_PAIRING_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.CAMPAIGN_ZONE_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.CAMPAIGN_ZONE_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.CHAT_TEXT_ARRIVED](../../systemdata/fields/systemdata_SystemData.Events.CHAT_TEXT_ARRIVED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.CONVERSATION_TEXT_ARRIVED](../../systemdata/fields/systemdata_SystemData.Events.CONVERSATION_TEXT_ARRIVED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.ENTER_WORLD](../../systemdata/fields/systemdata_SystemData.Events.ENTER_WORLD.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.GROUP_ACCEPT_INVITATION](../../systemdata/fields/systemdata_SystemData.Events.GROUP_ACCEPT_INVITATION.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.GROUP_ACCEPT_REFERRAL](../../systemdata/fields/systemdata_SystemData.Events.GROUP_ACCEPT_REFERRAL.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.GROUP_EFFECTS_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.GROUP_EFFECTS_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.GROUP_LEAVE](../../systemdata/fields/systemdata_SystemData.Events.GROUP_LEAVE.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.GROUP_STATUS_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.GROUP_STATUS_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.GROUP_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.GROUP_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.GUILD_EXP_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.GUILD_EXP_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.INTERACT_LEAVE_SCENARIO_QUEUE](../../systemdata/fields/systemdata_SystemData.Events.INTERACT_LEAVE_SCENARIO_QUEUE.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.INTERFACE_RELOADED](../../systemdata/fields/systemdata_SystemData.Events.INTERFACE_RELOADED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.LOADING_BEGIN](../../systemdata/fields/systemdata_SystemData.Events.LOADING_BEGIN.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.LOADING_END](../../systemdata/fields/systemdata_SystemData.Events.LOADING_END.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.L_BUTTON_DOWN_PROCESSED](../../systemdata/fields/systemdata_SystemData.Events.L_BUTTON_DOWN_PROCESSED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.L_BUTTON_UP_PROCESSED](../../systemdata/fields/systemdata_SystemData.Events.L_BUTTON_UP_PROCESSED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.MACRO_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.MACRO_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.OBJECTIVE_CONTROL_POINTS_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.OBJECTIVE_CONTROL_POINTS_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.OBJECTIVE_OWNER_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.OBJECTIVE_OWNER_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_ABILITY_TOGGLED](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_ABILITY_TOGGLED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_BEGIN_CAST](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_BEGIN_CAST.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_COMBAT_FLAG_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_COMBAT_FLAG_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_DEATH](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_DEATH.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_EFFECTS_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_EFFECTS_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_END_CAST](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_END_CAST.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_EXP_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_EXP_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_HOT_BAR_PAGE_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_HOT_BAR_PAGE_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_HOT_BAR_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_HOT_BAR_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_INVENTORY_SLOT_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_INVENTORY_SLOT_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_MONEY_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_MONEY_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_MORALE_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_MORALE_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_NEW_ABILITY_LEARNED](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_NEW_ABILITY_LEARNED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_POSITION_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_POSITION_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_RENOWN_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_RENOWN_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_TARGET_EFFECTS_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_TARGET_EFFECTS_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_TARGET_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_TARGET_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PUBLIC_QUEST_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.PUBLIC_QUEST_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.RELOAD_INTERFACE](../../systemdata/fields/systemdata_SystemData.Events.RELOAD_INTERFACE.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.R_BUTTON_UP_PROCESSED](../../systemdata/fields/systemdata_SystemData.Events.R_BUTTON_UP_PROCESSED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.SCENARIO_BEGIN](../../systemdata/fields/systemdata_SystemData.Events.SCENARIO_BEGIN.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.SCENARIO_END](../../systemdata/fields/systemdata_SystemData.Events.SCENARIO_END.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.SCENARIO_GROUP_JOIN](../../systemdata/fields/systemdata_SystemData.Events.SCENARIO_GROUP_JOIN.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.SCENARIO_GROUP_LEAVE](../../systemdata/fields/systemdata_SystemData.Events.SCENARIO_GROUP_LEAVE.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.SCENARIO_GROUP_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.SCENARIO_GROUP_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.SCENARIO_INSTANCE_CANCEL](../../systemdata/fields/systemdata_SystemData.Events.SCENARIO_INSTANCE_CANCEL.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.SCENARIO_INSTANCE_JOIN_NOW](../../systemdata/fields/systemdata_SystemData.Events.SCENARIO_INSTANCE_JOIN_NOW.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.SCENARIO_PLAYERS_LIST_GROUPS_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.SCENARIO_PLAYERS_LIST_GROUPS_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.SCENARIO_PLAYERS_LIST_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.SCENARIO_PLAYERS_LIST_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.SCENARIO_POST_MODE](../../systemdata/fields/systemdata_SystemData.Events.SCENARIO_POST_MODE.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.SCENARIO_SHOW_JOIN_PROMPT](../../systemdata/fields/systemdata_SystemData.Events.SCENARIO_SHOW_JOIN_PROMPT.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.SCENARIO_STARTING_SCENARIO_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.SCENARIO_STARTING_SCENARIO_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.SHOW_ALERT_TEXT](../../systemdata/fields/systemdata_SystemData.Events.SHOW_ALERT_TEXT.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.SPELL_CAST_CANCEL](../../systemdata/fields/systemdata_SystemData.Events.SPELL_CAST_CANCEL.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.UPDATE_PROCESSED](../../systemdata/fields/systemdata_SystemData.Events.UPDATE_PROCESSED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.USER_SETTINGS_CHANGED](../../systemdata/fields/systemdata_SystemData.Events.USER_SETTINGS_CHANGED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.VISIBLE_EQUIPMENT_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.VISIBLE_EQUIPMENT_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.WORLD_MAP_POINTS_LOADED](../../systemdata/fields/systemdata_SystemData.Events.WORLD_MAP_POINTS_LOADED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.WORLD_OBJ_COMBAT_EVENT](../../systemdata/fields/systemdata_SystemData.Events.WORLD_OBJ_COMBAT_EVENT.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.WORLD_OBJ_RENOWN_GAINED](../../systemdata/fields/systemdata_SystemData.Events.WORLD_OBJ_RENOWN_GAINED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.WORLD_OBJ_XP_GAINED](../../systemdata/fields/systemdata_SystemData.Events.WORLD_OBJ_XP_GAINED.md) (HIGH 100/100) - SystemData Field

## Notes

- Expected callback signature inferred from common WAR XML handler conventions (MEDIUM confidence).
