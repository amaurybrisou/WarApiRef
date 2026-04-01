# TextLogGetUpdateEventId

- Category: Game Event
- Confidence level: HIGH
- Confidence score: 100/100

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 126

- Rationale: Promoted as HIGH confidence because referenced by generated docs or reference files, called globally with no local definition, reinforced across multiple generated source types.

## Evidence Signals

- +18 Seen in 2 to 3 addons: Cross-addon spread is present but limited.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +18 Used in event registration or dispatch: Observed in event-driven engine hooks.
- +10 Referenced from initialization flow: Lifecycle reconstruction references this symbol.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +25 Referenced by generated docs or reference files: The symbol is reinforced outside a single call page.
- +20 Reinforced across multiple generated source types: Evidence comes from several independent API_Ref source types.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | Busted, EA_UiDebugTools, TidyChat |
| Files seen in | `/workspace/Busted/Busted.lua:324`, `/workspace/TidyChat/TidyChat.lua:1175`, `/workspace/ea_uidebugtools/Source/bust/Busted.lua:331` |
| Namespaces detected | TextLogGetUpdateEventId |
| Source kinds | event_page, flows, lua_event_registration |
| Example locations | Busted: Busted.Initialize, EA_UiDebugTools: Busted.Initialize, TidyChat: TidyChatLogs.Initialize |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 5 |
| Global usage count | 5 |
| Local definition count | 0 |
| Documentation references | 2 |
| Initialization flow references | 9 |
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
| Conflicting signatures | no |
| Conflicting roles | no |
| Wrapper likely | no |
| Never outside local graph | no |
| Local helper only | no |

## Description

Observed as a runtime event or event-like identifier used by 3 addons.

## Handler Pattern

Observed as a runtime event ID routed through RegisterEventHandler-style APIs.

## Payload

- Payload not inferable from addon-level documentation alone.

## Seen In

- Busted
- EA_UiDebugTools
- TidyChat

## Registrars And Handlers

- Busted.ProcessUiLogUpdateProxy
- RegisterEventHandler
- TidyChat.OnChatLogUpdated
- TidyChat.OnCombatLogUpdated
- TidyChat.OnSystemLogUpdated
- WindowRegisterEventHandler
- global
- window

## Examples

- Busted: Busted.Initialize -> TextLogGetUpdateEventId -> Busted.ProcessUiLogUpdateProxy
- EA_UiDebugTools: Busted.Initialize -> TextLogGetUpdateEventId -> Busted.ProcessUiLogUpdateProxy
- TidyChat: TidyChatLogs.Initialize -> TextLogGetUpdateEventId -> TidyChat.OnChatLogUpdated
- TidyChat: TidyChatLogs.Initialize -> TextLogGetUpdateEventId -> TidyChat.OnCombatLogUpdated
- TidyChat: TidyChatLogs.Initialize -> TextLogGetUpdateEventId -> TidyChat.OnSystemLogUpdated
- Busted: Busted.ProcessUiLogUpdateProxy -> WindowRegisterEventHandler(TextLogGetUpdateEventId, Busted.ProcessUiLogUpdateProxy)

## Related APIs

- none

## Used With

- [WindowRegisterEventHandler](../../window_api/functions/window_WindowRegisterEventHandler.md) (HIGH 100/100) - Window Function
- [WindowSetShowing](../../window_api/functions/window_WindowSetShowing.md) (HIGH 100/100) - Window Function
- [CreateWindow](../../globals/functions/global_CreateWindow.md) (HIGH 75/100) - Global Function
- [RegisterEventHandler](../../globals/functions/global_RegisterEventHandler.md) (MEDIUM 68/100) - Global Function

## Triggered By

- none

## Affects

- none

## Notes

- none
