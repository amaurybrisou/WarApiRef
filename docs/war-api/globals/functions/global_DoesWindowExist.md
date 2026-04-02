# DoesWindowExist

- Category: Global Function
- Confidence level: HIGH
- Confidence score: 71/100
- Seen in: 2 addons

## Confidence Assessment

- Level: HIGH

- Score: 71/100

- Rationale: Promoted as HIGH confidence because called globally with no local definition, seen in 2 to 3 addons, role is consistent across addons.

## Evidence Signals

- +18 Seen in 2 to 3 addons: Cross-addon spread is present but limited.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +10 Argument pattern is consistent: Observed argument positions remain stable.
- +8 Return usage is consistent: Observed as a stable query-style API.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | TidyChat, TidyRoll |
| Files seen in | `/workspace/data/raw/TidyChat/TidyChat.lua:930`, `/workspace/data/raw/TidyRoll/TidyRoll.lua:265`, `/workspace/data/raw/TidyRoll/TidyRollOptions.lua:853`, `/workspace/data/raw/TidyRoll/TidyRollOptions.lua:865` |
| Namespaces detected | DoesWindowExist |
| Source kinds | lua_calls |
| Example locations | TidyChat: TidyChatFrames.Initialize, TidyRoll: TidyRoll.OnLoad, TidyRoll: TidyRollOptions.RadioGetId, TidyRoll: TidyRollOptions.RadioSetId |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 4 |
| Global usage count | 4 |
| Local definition count | 0 |
| Documentation references | 0 |
| Initialization flow references | 0 |
| Known engine namespace | no |
| Default UI presence | no |
| Event binding presence | no |
| Observed in XML and Lua | no |
| Consistent role | yes |
| Consistent arguments | yes |
| Consistent returns | yes |
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
DoesWindowExist(windowName)
```

## Description

Observed guarding runtime window creation and cleanup by checking whether a named window exists.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| windowName | Observed as a target window name. | Observed values: "EA_Window_LootRoll", c_TEXT_ENTRY_WINDOW.."EntryBoxLanguageButton", radioGroupName..index |

## Returns

- Observed as a query-style API. The concrete return shape is not inferable from addon-api docs alone.

## Side Effects

- No side effect is confidently inferable from addon-api docs alone.

## Seen In

- TidyChat
- TidyRoll

## Examples

- TidyChat: TidyChatFrames.Initialize -> DoesWindowExist(c_TEXT_ENTRY_WINDOW.."EntryBoxLanguageButton")
- TidyRoll: TidyRoll.OnLoad -> DoesWindowExist("EA_Window_LootRoll")
- TidyRoll: TidyRollOptions.RadioGetId -> DoesWindowExist(radioGroupName..index)
- TidyRoll: TidyRollOptions.RadioSetId -> DoesWindowExist(radioGroupName..index)

## Related APIs

- none

## Used With

- [LayoutEditor.RegisterWindow](../../window_api/functions/window_LayoutEditor.RegisterWindow.md) (HIGH 100/100) - Window Function
- [LibSlash.RegisterWSlashCmd](global_LibSlash.RegisterWSlashCmd.md) (HIGH 100/100) - Global Function
- [SystemData.Events.LOADING_END](../../events/game_events/game_event_SystemData.Events.LOADING_END.md) (HIGH 100/100) - Game Event
- [SystemData.Events.RELOAD_INTERFACE](../../events/game_events/game_event_SystemData.Events.RELOAD_INTERFACE.md) (HIGH 100/100) - Game Event
- [WindowAddAnchor](../../window_api/functions/window_WindowAddAnchor.md) (HIGH 100/100) - Window Function
- [WindowClearAnchors](../../window_api/functions/window_WindowClearAnchors.md) (HIGH 100/100) - Window Function
- [WindowGetShowing](../../window_api/functions/window_WindowGetShowing.md) (HIGH 100/100) - Window Function
- [WindowRegisterCoreEventHandler](../../window_api/functions/window_WindowRegisterCoreEventHandler.md) (HIGH 100/100) - Window Function
- [WindowSetShowing](../../window_api/functions/window_WindowSetShowing.md) (HIGH 100/100) - Window Function
- [WindowUnregisterCoreEventHandler](../../window_api/functions/window_WindowUnregisterCoreEventHandler.md) (HIGH 90/100) - Window Function

## Triggered By

- [SystemData.Events.LOADING_END](../../events/game_events/game_event_SystemData.Events.LOADING_END.md) (HIGH 100/100) - Game Event
- [SystemData.Events.RELOAD_INTERFACE](../../events/game_events/game_event_SystemData.Events.RELOAD_INTERFACE.md) (HIGH 100/100) - Game Event

## Affects

- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Notes

- Canonical entry built from observed call sites, not from engine source or decompiled definitions.
