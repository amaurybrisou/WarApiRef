# UnregisterEventHandler

- Category: Global Function
- Confidence level: HIGH
- Confidence score: 93/100
- Seen in: 22 addons

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
| Addons seen in | Ace, AdvancedPetAssist, AdvancedRenownTrainer, AggroMeter, Aura, AutoMark, BuffHead, CM_ClosetGoblin |
| Files seen in | `/workspace/data/raw/Ace/AceAddon-3.0.lua:591`, `/workspace/data/raw/AdvancedPetAssist/AdvancedPetAssist.lua:98`, `/workspace/data/raw/AggroMeter/AggroMeter.lua:49`, `/workspace/data/raw/Aura/Source/TargetInfoFix.lua:54`, `/workspace/data/raw/AutoMark/Source/AutoMark.lua:78`, `/workspace/data/raw/BuffHead/Core.lua:207`, `/workspace/data/raw/BuffHead/Core.lua:229`, `/workspace/data/raw/BuffHead/Setup/ContainerDemo.lua:223` |
| Namespaces detected | UnregisterEventHandler |
| Source kinds | lua_calls |
| Example locations | Ace: AceAddon_OnUpdate_DONOTTOUCH, AdvancedPetAssist: AdvancedPetAssist.local.UnregisterLoadingEnd, AdvancedPetAssist: UnregisterLoadingEnd, AdvancedRenownTrainer: AdvancedRenownTraining.OnReload, AggroMeter: AggroMeter.Shutdown, Aura: TargetInfoFix.SET_TARGETINFO_FIX_UPDATE_FLAG_DONOTTOUCH |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 129 |
| Global usage count | 129 |
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
UnregisterEventHandler(eventId, handlerName)
```

## Description

Observed removing previously registered global runtime handlers.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| eventId | Observed as a SystemData or runtime event identifier. | Observed values: S.combatLogEventId, SystemData.Events.AUCTION_INIT_RECEIVED, SystemData.Events.CAMPAIGN_ZONE_UPDATED |
| handlerName | Observed as a Lua handler function reference. | Observed values: "AceAddon_OnUpdate_DONOTTOUCH", "AdvancedRenownTraining.CreateDataTable", "AggroMeter.OnChatLogUpdated" |

## Returns

- Not confidently inferable from addon-api docs alone.

## Side Effects

- Removes a previously registered callback or command binding.

## Seen In

- Ace
- AdvancedPetAssist
- AdvancedRenownTrainer
- AggroMeter
- Aura
- AutoMark
- BuffHead
- CM_ClosetGoblin
- CombatTextNames
- DAoCBuff
- Enemy
- Killer
- LibGuard
- PartyCast
- RoR_SoR
- Shinies
- TexturedButtons
- TidyRoll
- TurretRange
- WSCT
- WhoHealedMe
- followTheLeader

## Examples

- Ace: AceAddon_OnUpdate_DONOTTOUCH -> UnregisterEventHandler(SystemData.Events.UPDATE_PROCESSED, "AceAddon_OnUpdate_DONOTTOUCH")
- AdvancedPetAssist: AdvancedPetAssist.local.UnregisterLoadingEnd -> UnregisterEventHandler(SystemData.Events.LOADING_END, LOADING_END_HANDLER)
- AdvancedPetAssist: UnregisterLoadingEnd -> UnregisterEventHandler(SystemData.Events.LOADING_END, LOADING_END_HANDLER)
- AdvancedRenownTrainer: AdvancedRenownTraining.OnReload -> UnregisterEventHandler(SystemData.Events.PLAYER_CAREER_CATEGORY_UPDATED, "AdvancedRenownTraining.CreateDataTable")
- AggroMeter: AggroMeter.Shutdown -> UnregisterEventHandler(TextLogGetUpdateEventId("Chat"), "AggroMeter.OnChatLogUpdated")
- Aura: TargetInfoFix.SET_TARGETINFO_FIX_UPDATE_FLAG_DONOTTOUCH -> UnregisterEventHandler(SystemData.Events.UPDATE_PROCESSED, "TargetInfoFix.APPLY_TARGETINFO_FIX_DONOTTOUCH")

## Related APIs

- none

## Used With

- [SystemData.Events.UPDATE_PROCESSED](../../systemdata/fields/systemdata_SystemData.Events.UPDATE_PROCESSED.md) (HIGH 100/100) - SystemData Field
- [WindowGetAlpha](../../window_api/functions/window_WindowGetAlpha.md) (HIGH 100/100) - Window Function

## Triggered By

- [SystemData.Events.LOADING_END](../../events/game_events/game_event_SystemData.Events.LOADING_END.md) (HIGH 100/100) - Game Event
- [SystemData.Events.PLAYER_TARGET_UPDATED](../../events/game_events/game_event_SystemData.Events.PLAYER_TARGET_UPDATED.md) (HIGH 100/100) - Game Event
- [SystemData.Events.RELOAD_INTERFACE](../../events/game_events/game_event_SystemData.Events.RELOAD_INTERFACE.md) (HIGH 100/100) - Game Event
- [SystemData.Events.UPDATE_PROCESSED](../../events/game_events/game_event_SystemData.Events.UPDATE_PROCESSED.md) (HIGH 100/100) - Game Event

## Affects

- [SystemData.Events.LOADING_END](../../systemdata/fields/systemdata_SystemData.Events.LOADING_END.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_CAREER_CATEGORY_UPDATED](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_CAREER_CATEGORY_UPDATED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.UPDATE_PROCESSED](../../systemdata/fields/systemdata_SystemData.Events.UPDATE_PROCESSED.md) (HIGH 100/100) - SystemData Field

## Notes

- Canonical entry built from observed call sites, not from engine source or decompiled definitions.
- Advanced return analysis: No strong return evidence observed
