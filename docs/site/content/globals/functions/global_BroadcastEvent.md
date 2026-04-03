# BroadcastEvent

- Category: Global Function
- Confidence level: HIGH
- Confidence score: 93/100
- Seen in: 45 addons

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
| Addons seen in | AuctionStats, AutoFocus, Brizio's Crappy Computer Medic, CNC, Calling, DasBoot, DetauntHelper, EA_ScenarioGroupWindow |
| Files seen in | AuctionStats.lua, AutoFocus.lua, CCM.lua, CNC.lua, Calling.lua, CallingChat.lua, CallingKeybinding.lua, Code/GroupIcons/GroupIcons.lua |
| Namespaces detected | BroadcastEvent |
| Source kinds | lua_calls |
| Example locations | AuctionStats: OnUpdate, AutoFocus: SlashExec, Brizio's Crappy Computer Medic: OnLButtonUpB1, Brizio's Crappy Computer Medic: OnLButtonUpB2, Brizio's Crappy Computer Medic: Switch, CNC: CombatCheck |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 125 |
| Global usage count | 125 |
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
BroadcastEvent(eventId)
```

## Description

Observed triggering a runtime event so existing handlers are notified.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| eventId | Observed as a runtime event identifier dispatched to listeners. | Observed values: EA_Window_ScenarioLobby.joinModes[EA_Window_ScenarioLobby.joinMode].joinSingleEvent, PartyTargetEvent[EZGuard.NewGuardTarget.index], SystemData.Events.CHANNEL_NAMES_UPDATED |

## Returns

- Not confidently inferable from contract artifacts alone.

## Side Effects

- Dispatches an event into the runtime event system.

## Seen In

- AuctionStats
- AutoFocus
- Brizio's Crappy Computer Medic
- CNC
- Calling
- DasBoot
- DetauntHelper
- EA_ScenarioGroupWindow
- EZGuard
- Effigy
- Enemy
- EveryBodyGuard
- FastInteract
- GCDsaver
- GuardBot
- GuildWarden
- HealGrid
- I HATE YOU THIS MUCH
- IdentityFound
- JunkDump
- Keyset
- Lib RuString
- LoyalPet
- Minmap
- PeaceOut
- Phantom
- PlayEffectsOn
- Pure
- Queue Queuer
- Quick Performance Toggle
- RVMOD_Manager
- SOR
- Squared
- SquaredClick
- Targets
- TastyButtons
- TidyQueue
- Trakario
- WARCommander
- compass
- nLootLink
- scenarioInfo
- scnoload
- whatsPugSc
- zMailMod

## Examples

- AuctionStats: OnUpdate -> BroadcastEvent(SystemData.Events.EXIT_GAME)
- AutoFocus: SlashExec -> BroadcastEvent(SystemData.Events.SEND_CHAT_TEXT)
- Brizio's Crappy Computer Medic: OnLButtonUpB1 -> BroadcastEvent(SystemData.Events.USER_SETTINGS_CHANGED)
- Brizio's Crappy Computer Medic: OnLButtonUpB2 -> BroadcastEvent(SystemData.Events.USER_SETTINGS_CHANGED)
- Brizio's Crappy Computer Medic: Switch -> BroadcastEvent(SystemData.Events.USER_SETTINGS_CHANGED)
- CNC: CombatCheck -> BroadcastEvent(SystemData.Events.USER_SETTINGS_CHANGED)

## Related APIs

- [OnLButtonDown](../../xml/handlers/handler_OnLButtonDown.md) (HIGH 88/100) - XML Event
- [OnUpdate](../../xml/handlers/handler_OnUpdate.md) (HIGH 88/100) - XML Event

## Used With

- [EA_ChatWindow.Print](global_EA_ChatWindow.Print.md) (HIGH 100/100) - Global Function
- [SystemData.Events.USER_SETTINGS_CHANGED](../../systemdata/fields/systemdata_SystemData.Events.USER_SETTINGS_CHANGED.md) (HIGH 100/100) - SystemData Field

## Affects

- [SystemData.Events.EXIT_GAME](../../systemdata/fields/systemdata_SystemData.Events.EXIT_GAME.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.SEND_CHAT_TEXT](../../systemdata/fields/systemdata_SystemData.Events.SEND_CHAT_TEXT.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.USER_SETTINGS_CHANGED](../../systemdata/fields/systemdata_SystemData.Events.USER_SETTINGS_CHANGED.md) (HIGH 100/100) - SystemData Field

## Notes

- Canonical entry built from observed call sites, not from engine source or decompiled definitions.
