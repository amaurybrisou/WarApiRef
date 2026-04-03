# SystemData.Events.SEND_CHAT_TEXT

- Category: SystemData Field
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 15 addons

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 168

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, seen in 4 or more addons, matches a known engine namespace.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +18 Used in event registration or dispatch: Observed in event-driven engine hooks.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +25 Referenced by generated docs or reference files: The symbol is reinforced outside a single call page.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | AutoFocus, Calling, DasBoot, DetauntHelper, EveryBodyGuard, GuardBot, GuildWarden, I HATE YOU THIS MUCH |
| Files seen in | AutoFocus.lua, Calling.lua, CallingChat.lua, DasBoot.lua, EveryBodyGuard.lua, GuardBot.lua, GuildWarden.lua, IHYTM.lua |
| Namespaces detected | SystemData |
| Source kinds | lua_call |
| Example locations | Assist, AssistOnCall, Broadcast, Broadcast1, Broadcast2, BroadcastC |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 40 |
| Global usage count | 40 |
| Local definition count | 0 |
| Documentation references | 1 |
| Initialization flow references | 0 |
| Known engine namespace | yes |
| Default UI presence | yes |
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

SystemData.SystemData.Events.SEND_CHAT_TEXT field accessed by 15 addons; commonly found in Assist and AssistOnCall, Broadcast, Broadcast1, Broadcast2, BroadcastC, BroadcastCity1, BroadcastCity2, BroadcastCityC, BroadcastCityP, BroadcastCityW, BroadcastD1, BroadcastD2, BroadcastDC, BroadcastDP, BroadcastDW, BroadcastInfo1, BroadcastInfo2, BroadcastInfoC, BroadcastInfoP, BroadcastInfoW, BroadcastP, BroadcastW, ChatPrint, DelayStart, LeaveGuild, LogChannel, ManageChat, OnPlayerEndCast, OnTargetPlayer, PurgeList, SendChatText, SendDetails, Sendtext, SlashExec, TargetPlayer, UnLogChannel, UnitLClick, lua_call, sendChatMessage contexts.

## Seen In

- AutoFocus
- Calling
- DasBoot
- DetauntHelper
- EveryBodyGuard
- GuardBot
- GuildWarden
- I HATE YOU THIS MUCH
- Queue Queuer
- SOR
- SquaredClick
- Targets
- Trakario
- WARCommander
- nLootLink

## Related APIs

- [BroadcastEvent](../../globals/functions/global_BroadcastEvent.md) (HIGH 93/100) - Global Function

## Notes

- Observed in contexts: Assist, AssistOnCall, Broadcast, Broadcast1, Broadcast2, BroadcastC
