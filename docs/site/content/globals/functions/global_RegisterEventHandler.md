# RegisterEventHandler

- Category: Global Function
- Confidence level: HIGH
- Confidence score: 93/100
- Seen in: 7 addons

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
| Addons seen in | Lib RuString, PartyCast, Soloq, TidyChat, TidyRoll, TimeToDie, ZCurse_Profiler |
| Files seen in | `/workspace/data/raw/CurseProfiler/CurseProfilerCompiled.lua:2134`, `/workspace/data/raw/PartyCast/PartyCast.lua:51`, `/workspace/data/raw/RuStringLib/RuStringLib.lua:300`, `/workspace/data/raw/Soloq/Soloq.lua:22`, `/workspace/data/raw/TidyChat/TidyChat.lua:1175`, `/workspace/data/raw/TidyChat/TidyChat.lua:144`, `/workspace/data/raw/TidyChat/TidyChat.lua:676`, `/workspace/data/raw/TidyRoll/TidyRoll.lua:227` |
| Namespaces detected | RegisterEventHandler |
| Source kinds | lua_calls |
| Example locations | Lib RuString: LibRuString.Init, PartyCast: PartyCast.Init, Soloq: Soloq.OnInitialize, TidyChat: TidyChat.Initialize, TidyChat: TidyChatHooks.SetupHooks, TidyChat: TidyChatLogs.Initialize |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 61 |
| Global usage count | 61 |
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
| eventId | Observed as a SystemData or runtime event identifier. | Observed values: SystemData.Events.ENTER_WORLD, SystemData.Events.EXIT_GAME, SystemData.Events.INTERACT_COMPLETE_QUEST |
| handlerName | Observed as a Lua handler function reference. | Observed values: "CurseProfiler.andnot", "CurseProfiler.breakelse", "CurseProfiler.breakfor" |

## Returns

- Not confidently inferable from addon-api docs alone.

## Side Effects

- Registers engine or library callbacks for later dispatch.

## Seen In

- Lib RuString
- PartyCast
- Soloq
- TidyChat
- TidyRoll
- TimeToDie
- ZCurse_Profiler

## Examples

- Lib RuString: LibRuString.Init -> RegisterEventHandler(SystemData.Events.LOADING_END, "LibRuString.OnLoad")
- Lib RuString: LibRuString.Init -> RegisterEventHandler(SystemData.Events.RELOAD_INTERFACE, "LibRuString.OnLoad")
- PartyCast: PartyCast.Init -> RegisterEventHandler(SystemData.Events.PLAYER_START_INTERACT_TIMER, "PartyCast.StartInteract")
- PartyCast: PartyCast.Init -> RegisterEventHandler(SystemData.Events.INTERACT_DONE, "PartyCast.EndCast")
- PartyCast: PartyCast.Init -> RegisterEventHandler(SystemData.Events.PLAYER_BEGIN_CAST, "PartyCast.StartCast")
- PartyCast: PartyCast.Init -> RegisterEventHandler(SystemData.Events.PLAYER_END_CAST, "PartyCast.EndCast")

## Related APIs

- none

## Used With

- [SystemData.Events.LOADING_END](../../systemdata/fields/systemdata_SystemData.Events.LOADING_END.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.RELOAD_INTERFACE](../../systemdata/fields/systemdata_SystemData.Events.RELOAD_INTERFACE.md) (HIGH 100/100) - SystemData Field

## Triggered By

- none

## Affects

- [SystemData.Events.ENTER_WORLD](../../systemdata/fields/systemdata_SystemData.Events.ENTER_WORLD.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.INTERACT_DONE](../../systemdata/fields/systemdata_SystemData.Events.INTERACT_DONE.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.INTERFACE_RELOADED](../../systemdata/fields/systemdata_SystemData.Events.INTERFACE_RELOADED.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.LOADING_END](../../systemdata/fields/systemdata_SystemData.Events.LOADING_END.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_BEGIN_CAST](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_BEGIN_CAST.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_CAST_TIMER_SETBACK](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_CAST_TIMER_SETBACK.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_DEATH](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_DEATH.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_END_CAST](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_END_CAST.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.PLAYER_START_INTERACT_TIMER](../../systemdata/fields/systemdata_SystemData.Events.PLAYER_START_INTERACT_TIMER.md) (HIGH 100/100) - SystemData Field
- [SystemData.Events.RELOAD_INTERFACE](../../systemdata/fields/systemdata_SystemData.Events.RELOAD_INTERFACE.md) (HIGH 100/100) - SystemData Field

## Notes

- Canonical entry built from observed call sites, not from engine source or decompiled definitions.
