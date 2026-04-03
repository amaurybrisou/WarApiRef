# SystemData.Events.RELOAD_INTERFACE

- Category: Game Event
- Confidence level: HIGH
- Confidence score: 100/100

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 198

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, seen in 4 or more addons, matches a known engine namespace.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +18 Used in event registration or dispatch: Observed in event-driven engine hooks.
- +10 Referenced from initialization flow: Lifecycle reconstruction references this symbol.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +25 Referenced by generated docs or reference files: The symbol is reinforced outside a single call page.
- +20 Reinforced across multiple generated source types: Evidence comes from several independent addon-api source types.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | Lib RuString, TidyChat, TidyRoll, TimeToDie |
| Files seen in | `/workspace/data/raw/RuStringLib/RuStringLib.lua:300`, `/workspace/data/raw/TidyChat/TidyChat.lua:144`, `/workspace/data/raw/TidyRoll/TidyRoll.lua:227`, `/workspace/data/raw/TimeToDie/TimeToDie.lua:244` |
| Namespaces detected | SystemData |
| Source kinds | event_page, flows, lua_event_registration |
| Example locations | Lib RuString: LibRuString.Init, TidyChat: TidyChat.Initialize, TidyRoll: TidyRoll.Initialize, TimeToDie: TimeToDie.Initialize |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 4 |
| Global usage count | 4 |
| Local definition count | 0 |
| Documentation references | 2 |
| Initialization flow references | 4 |
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

Observed as a shared SystemData runtime event used by 4 addons.

## Handler Pattern

Observed as a runtime event ID routed through RegisterEventHandler-style APIs.

## Payload

- Payload shape is not inferable from addon-api docs alone; treat this as an engine event identifier.

## Seen In

- Lib RuString
- TidyChat
- TidyRoll
- TimeToDie

## Registrars And Handlers

- LibRuString.OnLoad
- RegisterEventHandler
- TidyChat.OnLoad
- TidyRoll.OnLoad
- TimeToDie.LoadingEnd
- global

## Examples

- Lib RuString: LibRuString.Init -> SystemData.Events.RELOAD_INTERFACE -> LibRuString.OnLoad
- TidyChat: TidyChat.Initialize -> SystemData.Events.RELOAD_INTERFACE -> TidyChat.OnLoad
- TidyRoll: TidyRoll.Initialize -> SystemData.Events.RELOAD_INTERFACE -> TidyRoll.OnLoad
- TimeToDie: TimeToDie.Initialize -> SystemData.Events.RELOAD_INTERFACE -> TimeToDie.LoadingEnd
- Lib RuString: LibRuString.OnLoad -> RegisterEventHandler(SystemData.Events.RELOAD_INTERFACE, LibRuString.OnLoad)
- TidyChat: TidyChat.OnLoad -> RegisterEventHandler(SystemData.Events.RELOAD_INTERFACE, TidyChat.OnLoad)

## Related APIs

- [LibSlash.RegisterWSlashCmd](../../globals/functions/global_LibSlash.RegisterWSlashCmd.md) (HIGH 100/100) - Global Function
- [WindowUnregisterCoreEventHandler](../../window_api/functions/window_WindowUnregisterCoreEventHandler.md) (HIGH 100/100) - Window Function
- [DoesWindowExist](../../globals/functions/global_DoesWindowExist.md) (HIGH 83/100) - Global Function

## Used With

- [LibSlash.RegisterWSlashCmd](../../globals/functions/global_LibSlash.RegisterWSlashCmd.md) (HIGH 100/100) - Global Function
- [SystemData.Events.LOADING_END](game_event_SystemData.Events.LOADING_END.md) (HIGH 100/100) - Game Event
- [WindowSetShowing](../../window_api/functions/window_WindowSetShowing.md) (HIGH 100/100) - Window Function
- [WindowUnregisterCoreEventHandler](../../window_api/functions/window_WindowUnregisterCoreEventHandler.md) (HIGH 100/100) - Window Function
- [DoesWindowExist](../../globals/functions/global_DoesWindowExist.md) (HIGH 83/100) - Global Function

## Triggered By

- none

## Affects

- none

## Notes

- none
