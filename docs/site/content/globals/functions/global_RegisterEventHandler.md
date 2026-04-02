# RegisterEventHandler

- Category: Global Function
- Confidence level: HIGH
- Confidence score: 81/100
- Seen in: 2 addons

## Confidence Assessment

- Level: HIGH

- Score: 81/100

- Rationale: Promoted as HIGH confidence because called globally with no local definition, seen in 2 to 3 addons, used in event registration or dispatch.

## Evidence Signals

- +18 Seen in 2 to 3 addons: Cross-addon spread is present but limited.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +18 Used in event registration or dispatch: Observed in event-driven engine hooks.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +10 Argument pattern is consistent: Observed argument positions remain stable.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | TidyChat, TidyRoll |
| Files seen in | `/workspace/data/raw/TidyChat/TidyChat.lua:1175`, `/workspace/data/raw/TidyChat/TidyChat.lua:144`, `/workspace/data/raw/TidyChat/TidyChat.lua:676`, `/workspace/data/raw/TidyRoll/TidyRoll.lua:227` |
| Namespaces detected | RegisterEventHandler |
| Source kinds | lua_calls |
| Example locations | TidyChat: TidyChat.Initialize, TidyChat: TidyChatHooks.SetupHooks, TidyChat: TidyChatLogs.Initialize, TidyRoll: TidyRoll.Initialize |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 10 |
| Global usage count | 10 |
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
| eventId | Observed as a SystemData or runtime event identifier. | Observed values: SystemData.Events.INTERACT_LOOT_ROLL_FIRST_ITEM, SystemData.Events.INTERACT_SHOW_LOOT_ROLL_DATA, SystemData.Events.LOADING_END |
| handlerName | Observed as a Lua handler function reference. | Observed values: "TidyChat.OnChatLogUpdated", "TidyChat.OnCombatLogUpdated", "TidyChat.OnLBU" |

## Returns

- Not confidently inferable from addon-api docs alone.

## Side Effects

- Registers engine or library callbacks for later dispatch.

## Seen In

- TidyChat
- TidyRoll

## Examples

- TidyChat: TidyChat.Initialize -> RegisterEventHandler(SystemData.Events.LOADING_END, "TidyChat.OnLoad")
- TidyChat: TidyChat.Initialize -> RegisterEventHandler(SystemData.Events.RELOAD_INTERFACE, "TidyChat.OnLoad")
- TidyChat: TidyChatHooks.SetupHooks -> RegisterEventHandler(SystemData.Events.L_BUTTON_UP_PROCESSED, "TidyChat.OnLBU")
- TidyChat: TidyChatLogs.Initialize -> RegisterEventHandler(chatLogEventId, "TidyChat.OnChatLogUpdated")
- TidyChat: TidyChatLogs.Initialize -> RegisterEventHandler(combatLogEventId, "TidyChat.OnCombatLogUpdated")
- TidyChat: TidyChatLogs.Initialize -> RegisterEventHandler(systemLogEventId, "TidyChat.OnSystemLogUpdated")

## Related APIs

- none

## Used With

- [TextLogGetUpdateEventId](../../events/game_events/game_event_TextLogGetUpdateEventId.md) (MEDIUM 63/100) - Game Event

## Triggered By

- none

## Affects

- [EA_ChatWindow](../tables/table_EA_ChatWindow.md) (HIGH 100/100) - Global Table
- [SystemData.Events.LOADING_END](../../systemdata/fields/systemdata_SystemData.Events.LOADING_END.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.L_BUTTON_UP_PROCESSED](../../systemdata/fields/systemdata_SystemData.Events.L_BUTTON_UP_PROCESSED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.RELOAD_INTERFACE](../../systemdata/fields/systemdata_SystemData.Events.RELOAD_INTERFACE.md) (HIGH 100/100) - SystemData Field

## Notes

- Canonical entry built from observed call sites, not from engine source or decompiled definitions.
