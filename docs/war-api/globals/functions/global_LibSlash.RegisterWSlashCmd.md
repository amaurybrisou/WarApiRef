# LibSlash.RegisterWSlashCmd

- Category: Global Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 4 addons

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 135

- Rationale: Promoted as HIGH confidence because seen in 4 or more addons, matches a known engine namespace, referenced by generated docs or reference files.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +10 Argument pattern is consistent: Observed argument positions remain stable.
- +10 Appears in slash command registration patterns: Observed in shared command registration flows.
- +25 Referenced by generated docs or reference files: The symbol is reinforced outside a single call page.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | Lib RuString, TidyChat, TidyRoll, TimeToDie |
| Files seen in | `/workspace/data/raw/RuStringLib/RuStringLib.lua:282`, `/workspace/data/raw/TidyChat/TidyChat.lua:189`, `/workspace/data/raw/TidyRoll/TidyRoll.lua:265`, `/workspace/data/raw/TimeToDie/TimeToDie.lua:229` |
| Namespaces detected | LibSlash |
| Source kinds | globals, lua_calls |
| Example locations | Lib RuString: LibRuString.OnLoad, TidyChat: TidyChat.OnLoad, TidyRoll: TidyRoll.OnLoad, TimeToDie: TimeToDie.LoadingEnd |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 7 |
| Global usage count | 7 |
| Local definition count | 0 |
| Documentation references | 1 |
| Initialization flow references | 0 |
| Known engine namespace | yes |
| Default UI presence | no |
| Event binding presence | no |
| Observed in XML and Lua | no |
| Consistent role | yes |
| Consistent arguments | yes |
| Consistent returns | no |
| Slash command presence | yes |
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
LibSlash.RegisterWSlashCmd(arg1, arg2)
```

## Description

Observed wiring slash commands through a shared command-registration table.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a text or wstring payload. | Observed values: "forcedrustrings", "tchat", "timetodie" |
| arg2 | Observed as a function or method reference. | Observed values: TidyChat.ToggleOptions, TidyRoll.ToggleOptions, TimeToDie.AverageReport |

## Returns

- Not confidently inferable from addon-api docs alone.

## Side Effects

- Registers engine or library callbacks for later dispatch.

## Seen In

- Lib RuString
- TidyChat
- TidyRoll
- TimeToDie

## Examples

- Lib RuString: LibRuString.OnLoad -> LibSlash.RegisterWSlashCmd("forcedrustrings", function(input)if(input==L "true")then LibRuString.ToggleHook(true)elseif(input==L "false")then LibRuString.ToggleHook(false)end end)
- TidyChat: TidyChat.OnLoad -> LibSlash.RegisterWSlashCmd("tchat", TidyChat.ToggleOptions)
- TidyRoll: TidyRoll.OnLoad -> LibSlash.RegisterWSlashCmd("troll", TidyRoll.ToggleOptions)
- TimeToDie: TimeToDie.LoadingEnd -> LibSlash.RegisterWSlashCmd("ttd", TimeToDie.AverageReport)
- TimeToDie: TimeToDie.LoadingEnd -> LibSlash.RegisterWSlashCmd("timetodie", TimeToDie.AverageReport)
- TimeToDie: TimeToDie.LoadingEnd -> LibSlash.RegisterWSlashCmd("ttdlast", TimeToDie.LastDeathReport)

## Related APIs

- none

## Used With

- [SystemData.Events.LOADING_END](../../events/game_events/game_event_SystemData.Events.LOADING_END.md) (HIGH 100/100) - Game Event
- [SystemData.Events.RELOAD_INTERFACE](../../events/game_events/game_event_SystemData.Events.RELOAD_INTERFACE.md) (HIGH 100/100) - Game Event
- [WindowSetShowing](../../window_api/functions/window_WindowSetShowing.md) (HIGH 100/100) - Window Function
- [WindowUnregisterCoreEventHandler](../../window_api/functions/window_WindowUnregisterCoreEventHandler.md) (HIGH 100/100) - Window Function

## Triggered By

- [SystemData.Events.LOADING_END](../../events/game_events/game_event_SystemData.Events.LOADING_END.md) (HIGH 100/100) - Game Event
- [SystemData.Events.RELOAD_INTERFACE](../../events/game_events/game_event_SystemData.Events.RELOAD_INTERFACE.md) (HIGH 100/100) - Game Event

## Affects

- none

## Notes

- Canonical entry built from observed call sites, not from engine source or decompiled definitions.
