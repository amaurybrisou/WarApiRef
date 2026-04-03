# RegisterEventHandler

- Category: Global Function
- Confidence level: MEDIUM
- Confidence score: 68/100
- Seen in: 216 addons

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
| Addons seen in | AbilityAlert, ActionBarHide, ActionFraction, ActionPointWatch, AdvancedPetAssist, AdvancedRenownTrainer, AggroMeter, ArmorGraphicToggle |
| Files seen in | AbilityAlert.lua, ActionBarHide.lua, ActionPointWatch.lua, AdvancedPetAssist.lua, AdvancedRenownTraining.lua, AggroMeter.lua, ArmorGraphicToggle.lua, ArsenalRank.lua |
| Namespaces detected | RegisterEventHandler |
| Source kinds | lua_calls |
| Example locations | AbilityAlert: Initialize, ActionBarHide: Initialize, ActionFraction: Initialize, ActionFraction: Start, ActionPointWatch: Initialize, AdvancedPetAssist: RegisterLoadingEnd |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 1060 |
| Global usage count | 1060 |
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
| eventId | Observed as a SystemData or runtime event identifier. | Observed values: "AuctionOnSearchResult", "EA_CareerResourceWindow", "LPET" |
| handlerName | Observed as a Lua handler function reference. | Observed values: "AbilityAlert.CombatEvent", "ActionBarHide.Combat", "ActionBarHide.OnLoad" |

## Returns

- Not confidently inferable from contract artifacts alone.

## Side Effects

- Registers engine or library callbacks for later dispatch.

## Seen In

- AbilityAlert
- ActionBarHide
- ActionFraction
- ActionPointWatch
- AdvancedPetAssist
- AdvancedRenownTrainer
- AggroMeter
- ArmorGraphicToggle
- Arsenal Rank
- Atlas
- AuctionStats
- Aura
- Auto Cape/Helm Show
- AutoBand
- AutoDismount
- AutoMark
- BBars - Mechanic Only
- BWMT
- BagOMatic
- BankArkel
- BlackBook
- Bloody Mess
- Brizio's Crappy Computer Medic
- BuffHead
- CDown
- CM_ClosetGoblin
- CNC
- CaVES
- CastSequence
- ChatAlert
- ChattyCathy
- CleanUnitFrames
- CleansingBuddy
- ClickSoundSuppressor
- CombatTextNames
- Compass3D
- CoolDownLine
- CraftingWillard
- Cram The Spam
- Crusher
- DAoCBuff
- DPSMeter
- Dascore
- DeMoNiCore
- Default Tactic Set
- DetauntHelper
- Ding
- DuffTimer
- EZGuard
- Effigy
- Emojii
- EmoteAlert
- Enemy
- EveryBodyGuard
- FastFriends
- FastInteract
- Fight Finder
- FixGit
- FlagCap
- GCDTracker
- GCDsaver
- GetStats
- Greedy
- GroupRange
- GroupSpotter
- GuardBot
- GuildWarden
- HammerTime
- HealAll
- HealGrid
- HealHoverAssist
- Hopper
- I HATE YOU THIS MUCH
- Info_DeathBlow
- Info_Loot
- Info_Points
- JunkDump
- KeyBar
- Keyset
- KeysetMonsterPlay
- KillTracker
- Killer
- Kwestor
- Lib RuString
- LibCooldown
- LibGuard
- LibSurveyor
- LibWBToggler
- LootAlert
- LoyalPet
- Map
- MapMonster
- Mech
- MegaphonePlus
- MegaphonePlusPlus
- MiniMapMonster
- Miracle Grow Remix
- MiracleGrow
- MiracleGrowLight
- MoraleBG
- Motion
- NAMBLA
- NaturalLog
- NerfedButtons
- NervAlert
- NoOverheal
- Obsidian
- OilTimer
- PMTB
- PartyAd
- PartyCast
- PeaceOut
- PetFixWindow
- Phantom
- PlanB
- PlayEffectsOn
- Pocket Palette
- PotionBar
- Pure
- Pure Careerbar
- Queue Queuer
- QueuePopTimer
- QuickNameActions+
- QuickTacticSwitch
- QuickWarReport
- RO-Style Combat Text
- RVAPI_ColorDialog
- RVAPI_Range
- RVMOD_3DPortrait
- RVMOD_Manager
- RVMOD_PlayerStatus
- RVMOD_SquaredDistances
- RVMOD_Targets
- RaidMeter
- Rangechecker
- RealmStatus
- Refer
- ReliquaryHunter
- ResHelp
- RetAlert
- RezEmote
- RoR_SoR
- RoR_debolster
- Rotation
- RvRContribution
- SNT_PANEL
- SOR
- ScenarioAlert
- ScenarioStats
- Sequencer
- SessionRPs
- Shinies
- ShowHealthPercent
- ShowMeTheBubbles
- SimpleCombatText
- Soloq
- SquaredClick
- SquaredHDIndicator
- SquaredHotIndicators
- Statdoll Remix
- StopRes
- Swinger
- TacticSetNames
- TargetInfoRing
- TargetRing
- Targets
- TastyButtons
- TexturedButtons
- ThankTheResser
- ThankTheTank
- TheSeeker
- TidyChat
- TidyQueue
- TidyRoll
- TimeToDie
- TokenMachine
- Tokens
- Tome Titan
- TomeTracker
- TortallDPSCore
- Trakario
- TurretRange
- TurretScrap
- Twister
- TwisterSet
- UnrealDBAnnouncer
- VPBreakdown
- Vectors
- WARCommander
- WSCT
- WarBoard
- WarBoard_Loc
- WarBoard_Session
- WarBoard_TaliMon
- War_RU
- WhatsCooking
- WhoHealedMe
- Wikki's Cooldown Bar
- Wikki's Cooldown Pulse
- WoH-Reticle
- ZCurse_Profiler
- ZonePOP
- compass
- followTheLeader
- hideInf
- nLootLink
- ovdeadnomore
- rorAutoInviter
- scenarioInfo
- scnoload
- talisman-monitor
- wbLeadHelper
- whatsPugSc
- whom
- xHUD
- zMailMod

## Examples

- AbilityAlert: Initialize -> RegisterEventHandler(SystemData.Events.WORLD_OBJ_COMBAT_EVENT, "AbilityAlert.CombatEvent")
- ActionBarHide: Initialize -> RegisterEventHandler(SystemData.Events.LOADING_END, "ActionBarHide.OnLoad")
- ActionBarHide: Initialize -> RegisterEventHandler(SystemData.Events.RELOAD_INTERFACE, "ActionBarHide.OnLoad")
- ActionBarHide: Initialize -> RegisterEventHandler(SystemData.Events.PLAYER_COMBAT_FLAG_UPDATED, "ActionBarHide.Combat")
- ActionFraction: Initialize -> RegisterEventHandler(SystemData.Events.LOADING_END, "ActionFractionWindow.UpdateVisibility")
- ActionFraction: Start -> RegisterEventHandler(SystemData.Events.ENTER_WORLD, "ActionFraction.Initialize")

## Related APIs

- [SystemData.Events.BATTLEGROUP_MEMBER_UPDATED](../../events/game_events/game_event_SystemData.Events.BATTLEGROUP_MEMBER_UPDATED.md) (HIGH 100/100) - Game Event
- [SystemData.Events.BATTLEGROUP_UPDATED](../../events/game_events/game_event_SystemData.Events.BATTLEGROUP_UPDATED.md) (HIGH 100/100) - Game Event
- [SystemData.Events.CITY_SCENARIO_BEGIN](../../events/game_events/game_event_SystemData.Events.CITY_SCENARIO_BEGIN.md) (HIGH 100/100) - Game Event
- [SystemData.Events.CITY_SCENARIO_END](../../events/game_events/game_event_SystemData.Events.CITY_SCENARIO_END.md) (HIGH 100/100) - Game Event
- [SystemData.Events.GROUP_STATUS_UPDATED](../../events/game_events/game_event_SystemData.Events.GROUP_STATUS_UPDATED.md) (HIGH 100/100) - Game Event
- [SystemData.Events.GROUP_UPDATED](../../events/game_events/game_event_SystemData.Events.GROUP_UPDATED.md) (HIGH 100/100) - Game Event
- [SystemData.Events.LOADING_END](../../events/game_events/game_event_SystemData.Events.LOADING_END.md) (HIGH 100/100) - Game Event
- [SystemData.Events.MACROS_LOADED](../../events/game_events/game_event_SystemData.Events.MACROS_LOADED.md) (HIGH 100/100) - Game Event
- [SystemData.Events.MACRO_UPDATED](../../events/game_events/game_event_SystemData.Events.MACRO_UPDATED.md) (HIGH 100/100) - Game Event
- [SystemData.Events.PLAYER_ACTIVE_TACTICS_UPDATED](../../events/game_events/game_event_SystemData.Events.PLAYER_ACTIVE_TACTICS_UPDATED.md) (HIGH 100/100) - Game Event
- [SystemData.Events.PLAYER_BEGIN_CAST](../../events/game_events/game_event_SystemData.Events.PLAYER_BEGIN_CAST.md) (HIGH 100/100) - Game Event
- [SystemData.Events.PLAYER_CAREER_LINE_UPDATED](../../events/game_events/game_event_SystemData.Events.PLAYER_CAREER_LINE_UPDATED.md) (HIGH 100/100) - Game Event
- [SystemData.Events.PLAYER_CAREER_RANK_UPDATED](../../events/game_events/game_event_SystemData.Events.PLAYER_CAREER_RANK_UPDATED.md) (HIGH 100/100) - Game Event
- [SystemData.Events.PLAYER_CUR_ACTION_POINTS_UPDATED](../../events/game_events/game_event_SystemData.Events.PLAYER_CUR_ACTION_POINTS_UPDATED.md) (HIGH 100/100) - Game Event
- [SystemData.Events.PLAYER_CUR_HIT_POINTS_UPDATED](../../events/game_events/game_event_SystemData.Events.PLAYER_CUR_HIT_POINTS_UPDATED.md) (HIGH 100/100) - Game Event
- [SystemData.Events.PLAYER_EFFECTS_UPDATED](../../events/game_events/game_event_SystemData.Events.PLAYER_EFFECTS_UPDATED.md) (HIGH 100/100) - Game Event
- [SystemData.Events.PLAYER_EQUIPMENT_SLOT_UPDATED](../../events/game_events/game_event_SystemData.Events.PLAYER_EQUIPMENT_SLOT_UPDATED.md) (HIGH 100/100) - Game Event
- [SystemData.Events.PLAYER_HOT_BAR_ENABLED_STATE_CHANGED](../../events/game_events/game_event_SystemData.Events.PLAYER_HOT_BAR_ENABLED_STATE_CHANGED.md) (HIGH 100/100) - Game Event
- [SystemData.Events.PLAYER_MAX_HIT_POINTS_UPDATED](../../events/game_events/game_event_SystemData.Events.PLAYER_MAX_HIT_POINTS_UPDATED.md) (HIGH 100/100) - Game Event
- [SystemData.Events.PLAYER_STATS_UPDATED](../../events/game_events/game_event_SystemData.Events.PLAYER_STATS_UPDATED.md) (HIGH 100/100) - Game Event
- [SystemData.Events.PLAYER_TARGET_UPDATED](../../events/game_events/game_event_SystemData.Events.PLAYER_TARGET_UPDATED.md) (HIGH 100/100) - Game Event
- [SystemData.Events.RELOAD_INTERFACE](../../events/game_events/game_event_SystemData.Events.RELOAD_INTERFACE.md) (HIGH 100/100) - Game Event
- [SystemData.Events.R_BUTTON_UP_PROCESSED](../../events/game_events/game_event_SystemData.Events.R_BUTTON_UP_PROCESSED.md) (HIGH 100/100) - Game Event
- [SystemData.Events.SCENARIO_BEGIN](../../events/game_events/game_event_SystemData.Events.SCENARIO_BEGIN.md) (HIGH 100/100) - Game Event
- [SystemData.Events.SCENARIO_END](../../events/game_events/game_event_SystemData.Events.SCENARIO_END.md) (HIGH 100/100) - Game Event
- [SystemData.Events.SCENARIO_GROUP_JOIN](../../events/game_events/game_event_SystemData.Events.SCENARIO_GROUP_JOIN.md) (HIGH 100/100) - Game Event
- [SystemData.Events.SCENARIO_GROUP_LEAVE](../../events/game_events/game_event_SystemData.Events.SCENARIO_GROUP_LEAVE.md) (HIGH 100/100) - Game Event
- [SystemData.Events.SCENARIO_PLAYERS_LIST_GROUPS_UPDATED](../../events/game_events/game_event_SystemData.Events.SCENARIO_PLAYERS_LIST_GROUPS_UPDATED.md) (HIGH 100/100) - Game Event
- [SystemData.Events.SCENARIO_PLAYERS_LIST_RESERVATIONS_UPDATED](../../events/game_events/game_event_SystemData.Events.SCENARIO_PLAYERS_LIST_RESERVATIONS_UPDATED.md) (HIGH 100/100) - Game Event
- [SystemData.Events.SCENARIO_PLAYER_HITS_UPDATED](../../events/game_events/game_event_SystemData.Events.SCENARIO_PLAYER_HITS_UPDATED.md) (HIGH 100/100) - Game Event
- [SystemData.Events.SHOW_ALERT_TEXT](../../events/game_events/game_event_SystemData.Events.SHOW_ALERT_TEXT.md) (HIGH 100/100) - Game Event
- [SystemData.Events.TOME_STAT_PLAYED_TIME_UPDATED](../../events/game_events/game_event_SystemData.Events.TOME_STAT_PLAYED_TIME_UPDATED.md) (HIGH 100/100) - Game Event
- [SystemData.Events.UPDATE_PROCESSED](../../events/game_events/game_event_SystemData.Events.UPDATE_PROCESSED.md) (HIGH 100/100) - Game Event
- [SystemData.Events.WORLD_OBJ_COMBAT_EVENT](../../events/game_events/game_event_SystemData.Events.WORLD_OBJ_COMBAT_EVENT.md) (HIGH 100/100) - Game Event
- [WindowRegisterCoreEventHandler](../../window_api/functions/window_WindowRegisterCoreEventHandler.md) (HIGH 100/100) - Window Function
- [OnInitialize](../../xml/handlers/handler_OnInitialize.md) (HIGH 88/100) - XML Event
- [OnShown](../../xml/handlers/handler_OnShown.md) (HIGH 88/100) - XML Event
- [CreateWindow](global_CreateWindow.md) (HIGH 75/100) - Global Function
- [Start](../../events/game_events/game_event_Start.md) (MEDIUM 43/100) - Game Event

## Used With

- [EA_ChatWindow.Print](global_EA_ChatWindow.Print.md) (HIGH 100/100) - Global Function
- [LabelSetText](../../window_api/functions/window_LabelSetText.md) (HIGH 100/100) - Window Function
- [LayoutEditor.RegisterWindow](../../window_api/functions/window_LayoutEditor.RegisterWindow.md) (HIGH 100/100) - Window Function
- [LibSlash.IsSlashCmdRegistered](global_LibSlash.IsSlashCmdRegistered.md) (HIGH 100/100) - Global Function
- [LibSlash.RegisterWSlashCmd](global_LibSlash.RegisterWSlashCmd.md) (HIGH 100/100) - Global Function
- [SystemData.Events.ENTER_WORLD](../../systemdata/fields/systemdata_SystemData.Events.ENTER_WORLD.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.INTERFACE_RELOADED](../../systemdata/fields/systemdata_SystemData.Events.INTERFACE_RELOADED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.LOADING_END](../../systemdata/fields/systemdata_SystemData.Events.LOADING_END.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_TARGET_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_TARGET_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.RELOAD_INTERFACE](../../systemdata/fields/systemdata_SystemData.Events.RELOAD_INTERFACE.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.UPDATE_PROCESSED](../../systemdata/fields/systemdata_SystemData.Events.UPDATE_PROCESSED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.WORLD_OBJ_COMBAT_EVENT](../../systemdata/fields/systemdata_SystemData.Events.WORLD_OBJ_COMBAT_EVENT.md) (HIGH 100/100) - SystemData Field
- [WindowSetShowing](../../window_api/functions/window_WindowSetShowing.md) (HIGH 100/100) - Window Function
- [OnInitialize](../../xml/handlers/handler_OnInitialize.md) (HIGH 88/100) - XML Event
- [LibSlash.RegisterSlashCmd](global_LibSlash.RegisterSlashCmd.md) (HIGH 85/100) - Global Function
- [CreateWindow](global_CreateWindow.md) (HIGH 75/100) - Global Function
- [Start](../../events/game_events/game_event_Start.md) (MEDIUM 43/100) - Game Event

## Affects

- [SystemData.Events.BATTLEGROUP_MEMBER_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.BATTLEGROUP_MEMBER_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.BATTLEGROUP_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.BATTLEGROUP_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.CITY_SCENARIO_BEGIN](../../systemdata/fields/systemdata_SystemData.Events.CITY_SCENARIO_BEGIN.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.CITY_SCENARIO_END](../../systemdata/fields/systemdata_SystemData.Events.CITY_SCENARIO_END.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.ENTER_WORLD](../../systemdata/fields/systemdata_SystemData.Events.ENTER_WORLD.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.GROUP_STATUS_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.GROUP_STATUS_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.GROUP_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.GROUP_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.INTERFACE_RELOADED](../../systemdata/fields/systemdata_SystemData.Events.INTERFACE_RELOADED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.LOADING_END](../../systemdata/fields/systemdata_SystemData.Events.LOADING_END.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.MACROS_LOADED](../../systemdata/fields/systemdata_SystemData.Events.MACROS_LOADED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.MACRO_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.MACRO_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_ACTIVE_TACTICS_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_ACTIVE_TACTICS_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_AGRO_MODE_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_AGRO_MODE_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_BEGIN_CAST](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_BEGIN_CAST.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_CAREER_LINE_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_CAREER_LINE_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_CAREER_RANK_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_CAREER_RANK_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_COMBAT_FLAG_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_COMBAT_FLAG_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_CUR_ACTION_POINTS_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_CUR_ACTION_POINTS_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_CUR_HIT_POINTS_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_CUR_HIT_POINTS_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_EFFECTS_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_EFFECTS_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_EQUIPMENT_SLOT_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_EQUIPMENT_SLOT_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_HEALTH_FADE_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_HEALTH_FADE_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_HOT_BAR_ENABLED_STATE_CHANGED](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_HOT_BAR_ENABLED_STATE_CHANGED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_MAX_ACTION_POINTS_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_MAX_ACTION_POINTS_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_MAX_HIT_POINTS_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_MAX_HIT_POINTS_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_STATS_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_STATS_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_TARGET_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_TARGET_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.RELOAD_INTERFACE](../../systemdata/fields/systemdata_SystemData.Events.RELOAD_INTERFACE.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.R_BUTTON_UP_PROCESSED](../../systemdata/fields/systemdata_SystemData.Events.R_BUTTON_UP_PROCESSED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.SCENARIO_BEGIN](../../systemdata/fields/systemdata_SystemData.Events.SCENARIO_BEGIN.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.SCENARIO_END](../../systemdata/fields/systemdata_SystemData.Events.SCENARIO_END.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.SCENARIO_GROUP_JOIN](../../systemdata/fields/systemdata_SystemData.Events.SCENARIO_GROUP_JOIN.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.SCENARIO_GROUP_LEAVE](../../systemdata/fields/systemdata_SystemData.Events.SCENARIO_GROUP_LEAVE.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.SCENARIO_PLAYERS_LIST_GROUPS_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.SCENARIO_PLAYERS_LIST_GROUPS_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.SCENARIO_PLAYERS_LIST_RESERVATIONS_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.SCENARIO_PLAYERS_LIST_RESERVATIONS_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.SCENARIO_PLAYER_HITS_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.SCENARIO_PLAYER_HITS_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.SHOW_ALERT_TEXT](../../systemdata/fields/systemdata_SystemData.Events.SHOW_ALERT_TEXT.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.TOME_STAT_PLAYED_TIME_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.TOME_STAT_PLAYED_TIME_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.UPDATE_PROCESSED](../../systemdata/fields/systemdata_SystemData.Events.UPDATE_PROCESSED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.WORLD_OBJ_COMBAT_EVENT](../../systemdata/fields/systemdata_SystemData.Events.WORLD_OBJ_COMBAT_EVENT.md) (HIGH 100/100) - SystemData Field

## Notes

- Canonical entry built from observed call sites, not from engine source or decompiled definitions.
- Advanced return analysis: No strong return evidence observed
