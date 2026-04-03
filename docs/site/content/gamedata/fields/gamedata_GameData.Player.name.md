# GameData.Player.name

- Category: GameData Field
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 51 addons

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 150

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, seen in 4 or more addons, matches a known engine namespace.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +25 Referenced by generated docs or reference files: The symbol is reinforced outside a single call page.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | Arsenal Rank, Aura, AutoBand, CDown, Calling, DammazKron, Dascore, DetauntHelper |
| Files seen in | ArsenalRank.lua, AutoBand.lua, CDown.lua, CDownSettings.lua, Calling.lua, Code/Core/Main.lua, Core/DK_Core.lua, DascorePars.lua |
| Namespaces detected | GameData |
| Source kinds | lua_call |
| Example locations | AllModulesInit, Check, CheckSettings, CheckSettingsInit, Convert03To04, FetchStats |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 76 |
| Global usage count | 76 |
| Local definition count | 0 |
| Documentation references | 1 |
| Initialization flow references | 0 |
| Known engine namespace | yes |
| Default UI presence | yes |
| Event binding presence | no |
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

GameData.GameData.Player.name field accessed by 51 addons; commonly found in AllModulesInit and Check, CheckSettings, CheckSettingsInit, Convert03To04, FetchStats, FollowCurrentTarget, GROUP_UPDATED, GetCurProfileNum, GetCurrentCharacter, GetCurrentLeaderName, GetDynamicLeaderName, GetIDs, GetItemPrice, GetLeaderEffectiveRoleCategory, GetMainAssist, GetPlayerName, GetRealPlayerName, GetSetting, InitProfile, InitUI, Initialize, InitializeLocal, InitializeSettings, IsUnitInRange, LevelUp, ON_DEATH, OnBattleGroupDataChanged, OnChat, OnGroupChange, OnInitialize, OnLoad, OnMButtonUp, OnMessageReceived, OnMouseOver, OnRButtonUp, OnRButtonUp_Results_ListItem, OnUpdate, ParsePlayer, ParserCreateExport, ProcessInboundChatMessage, ProcessSlashCmd, SetPlayerDefaults, SetSetting, SettingsInitialize, ShowDropdownMenu, ShowMenu, ShowTab, ShowWindowByName, StartTracker, SwitchProfile, UnitRRelease, UpdateBySpellsRange, UpdateGroup, UpdateHealthArray, UpdateInformation, UpdatePetHealthProxy, UpdateQueuerList, UpdateStateMachine, UpdateWindow, addInteractive, auto_kick, build_arrival_seed_from_open_party, cmd_custom_role, filter, getSelfName, get_party_counts, is_wb_leader, lua_call, pa_collect_group_leader_name_keys, pa_get_live_partynote_from_members, setStaticLabels, slash, update, updateRvRTab, update_wb_arrival_tracking contexts.

## Seen In

- Arsenal Rank
- Aura
- AutoBand
- CDown
- Calling
- DammazKron
- Dascore
- DetauntHelper
- Ding
- EA_ScenarioGroupWindow
- Effigy
- EmoteAlert
- Enemy
- FastFriends
- GuardLine
- GuildWarden
- HealGrid
- HealHoverAssist
- I HATE YOU THIS MUCH
- Info_DeathBlow
- Info_Points
- KeyBar
- Kwestor
- LibGuard
- MegaphonePlusPlus
- NerfedButtons
- NoOverheal
- PartyAd
- PartyCast
- Pure
- Queue Queuer
- QuickWarReport
- RVAPI_Range
- Rotation
- RvRStats
- RvRStatsTab
- ScenarioStats
- SelfTarget
- Sequencer
- SessionRPs
- Shinies
- Soloq
- Squared
- SquaredClick
- TacticSetNames
- ThinkOutLoud
- TortallDPSCore
- War_RU
- followTheLeader
- rorAutoInviter
- zMailMod

## Related APIs

- [GameData.GetScenarioPlayers](../../globals/functions/global_GameData.GetScenarioPlayers.md) (HIGH 100/100) - Global Function
- [LayoutEditor.RegisterEditCallback](../../window_api/functions/window_LayoutEditor.RegisterEditCallback.md) (HIGH 100/100) - Window Function
- [PartyUtils.IsPartyActive](../../globals/functions/global_PartyUtils.IsPartyActive.md) (HIGH 100/100) - Global Function
- [PartyUtils.IsPlayerInWarband](../../globals/functions/global_PartyUtils.IsPlayerInWarband.md) (HIGH 100/100) - Global Function
- [OnInitialize](../../xml/handlers/handler_OnInitialize.md) (HIGH 88/100) - XML Event
- [OnMButtonUp](../../xml/handlers/handler_OnMButtonUp.md) (HIGH 88/100) - XML Event
- [OnMouseOver](../../xml/handlers/handler_OnMouseOver.md) (HIGH 88/100) - XML Event
- [OnRButtonUp](../../xml/handlers/handler_OnRButtonUp.md) (HIGH 88/100) - XML Event
- [OnUpdate](../../xml/handlers/handler_OnUpdate.md) (HIGH 88/100) - XML Event
- [PartyUtils.IsWarbandFull](../../globals/functions/global_PartyUtils.IsWarbandFull.md) (MEDIUM 45/100) - Global Function

## Notes

- Observed in contexts: AllModulesInit, Check, CheckSettings, CheckSettingsInit, Convert03To04, FetchStats
