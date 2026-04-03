# UnregisterEventHandler

- Category: Global Function
- Confidence level: MEDIUM
- Confidence score: 68/100
- Seen in: 104 addons

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
| Addons seen in | Ace, ActionFraction, AdvancedPetAssist, AdvancedRenownTrainer, AggroMeter, ArmorGraphicToggle, AuctionStats, Aura |
| Files seen in | AceAddon-3.0.lua, AdvancedPetAssist.lua, AdvancedRenownTraining.lua, AggroMeter.lua, ArmorGraphicToggle.lua, AuctionAssist.lua, AuctionStats.lua, AutoDismount.lua |
| Namespaces detected | UnregisterEventHandler |
| Source kinds | lua_calls |
| Example locations | Ace: AceAddon_OnUpdate_DONOTTOUCH, ActionFraction: Initialize, ActionFraction: Shutdown, AdvancedPetAssist: UnregisterLoadingEnd, AdvancedRenownTrainer: OnReload, AggroMeter: Shutdown |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 483 |
| Global usage count | 483 |
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
UnregisterEventHandler(eventId, handlerName)
```

## Description

Observed removing previously registered global runtime handlers.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| eventId | Observed as a SystemData or runtime event identifier. | Observed values: "EA_CareerResourceWindow", "LPET", "Root" |
| handlerName | Observed as a Lua handler function reference. | Observed values: "AceAddon_OnUpdate_DONOTTOUCH", "ActionFraction.Initialize", "ActionFractionWindow.UpdateVisibility" |

## Returns

- Not confidently inferable from contract artifacts alone.

## Side Effects

- Removes a previously registered callback or command binding.

## Seen In

- Ace
- ActionFraction
- AdvancedPetAssist
- AdvancedRenownTrainer
- AggroMeter
- ArmorGraphicToggle
- AuctionStats
- Aura
- AutoDismount
- AutoMark
- BBars - Mechanic Only
- BWMT
- BlackBook
- Bloody Mess
- Brizio's Crappy Computer Medic
- BuffHead
- CCTV
- CDown
- CM_ClosetGoblin
- CaVES
- ChatAlert
- CleansingBuddy
- CombatTextNames
- Compass3D
- Cram The Spam
- DAoCBuff
- DPSMeter
- DeMoNiCore
- DetauntHelper
- Ding
- EZGuard
- Effigy
- Enemy
- EveryBodyGuard
- FastFriends
- FastInteract
- FlagCap
- GCDsaver
- Greedy
- GuardBot
- GuildWarden
- HealGrid
- Hopper
- I HATE YOU THIS MUCH
- Info_DeathBlow
- Info_Points
- KillTracker
- Killer
- LibGuard
- LoyalPet
- MapMonster
- Mech
- MegaphonePlusPlus
- NerfedButtons
- NoOverheal
- Obsidian
- OilTimer
- PartyAd
- PartyCast
- PeaceOut
- Phantom
- Pure
- Queue Queuer
- QuickNameActions+
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
- Refer
- ReliquaryHunter
- ResHelp
- RetAlert
- RezEmote
- RoR_SoR
- SNT_BUTTONS
- ScenarioStats
- SessionRPs
- Shinies
- SquaredClick
- Statdoll Remix
- StopRes
- TastyButtons
- TexturedButtons
- ThankTheResser
- TheSeeker
- TidyRoll
- TomeTracker
- Trakario
- TurretRange
- Vectors
- WSCT
- WhoHealedMe
- followTheLeader
- hideInf
- nLootLink
- rorAutoInviter
- scenarioInfo
- whatsPugSc

## Examples

- Ace: AceAddon_OnUpdate_DONOTTOUCH -> UnregisterEventHandler(SystemData.Events.UPDATE_PROCESSED, "AceAddon_OnUpdate_DONOTTOUCH")
- ActionFraction: Initialize -> UnregisterEventHandler(SystemData.Events.ENTER_WORLD, "ActionFraction.Initialize")
- ActionFraction: Initialize -> UnregisterEventHandler(SystemData.Events.INTERFACE_RELOADED, "ActionFraction.Initialize")
- ActionFraction: Shutdown -> UnregisterEventHandler(SystemData.Events.LOADING_END, "ActionFractionWindow.UpdateVisibility")
- AdvancedPetAssist: UnregisterLoadingEnd -> UnregisterEventHandler(SystemData.Events.LOADING_END, LOADING_END_HANDLER)
- AdvancedRenownTrainer: OnReload -> UnregisterEventHandler(SystemData.Events.PLAYER_CAREER_CATEGORY_UPDATED, "AdvancedRenownTraining.CreateDataTable")

## Related APIs

- [WindowUnregisterCoreEventHandler](../../window_api/functions/window_WindowUnregisterCoreEventHandler.md) (HIGH 100/100) - Window Function
- [OnHidden](../../xml/handlers/handler_OnHidden.md) (HIGH 88/100) - XML Event
- [OnInitialize](../../xml/handlers/handler_OnInitialize.md) (HIGH 88/100) - XML Event
- [OnShutdown](../../xml/handlers/handler_OnShutdown.md) (HIGH 88/100) - XML Event
- [DestroyWindow](global_DestroyWindow.md) (HIGH 75/100) - Global Function
- [Stop](../../events/game_events/game_event_Stop.md) (MEDIUM 43/100) - Game Event

## Used With

- [SystemData.Events.PLAYER_TARGET_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_TARGET_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.UPDATE_PROCESSED](../../systemdata/fields/systemdata_SystemData.Events.UPDATE_PROCESSED.md) (HIGH 100/100) - SystemData Field
- [OnShutdown](../../xml/handlers/handler_OnShutdown.md) (HIGH 88/100) - XML Event
- [DestroyWindow](global_DestroyWindow.md) (HIGH 75/100) - Global Function

## Affects

- [SystemData.Events.LOADING_END](../../systemdata/fields/systemdata_SystemData.Events.LOADING_END.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_ACTIVE_TACTICS_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_ACTIVE_TACTICS_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_AGRO_MODE_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_AGRO_MODE_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_CAREER_CATEGORY_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_CAREER_CATEGORY_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_CUR_ACTION_POINTS_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_CUR_ACTION_POINTS_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_CUR_HIT_POINTS_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_CUR_HIT_POINTS_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_EFFECTS_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_EFFECTS_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_EQUIPMENT_SLOT_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_EQUIPMENT_SLOT_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_HEALTH_FADE_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_HEALTH_FADE_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_HOT_BAR_ENABLED_STATE_CHANGED](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_HOT_BAR_ENABLED_STATE_CHANGED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_MAX_ACTION_POINTS_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_MAX_ACTION_POINTS_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_STATS_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_STATS_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_TARGET_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_TARGET_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.TOME_STAT_PLAYED_TIME_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.TOME_STAT_PLAYED_TIME_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.UPDATE_PROCESSED](../../systemdata/fields/systemdata_SystemData.Events.UPDATE_PROCESSED.md) (HIGH 100/100) - SystemData Field

## Notes

- Canonical entry built from observed call sites, not from engine source or decompiled definitions.
- Advanced return analysis: No strong return evidence observed
