# WindowGetShowing

- Category: Window Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 2 addons

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 123

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, matches a known engine namespace, called globally with no local definition.

## Evidence Signals

- +18 Seen in 2 to 3 addons: Cross-addon spread is present but limited.
- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +10 Argument pattern is consistent: Observed argument positions remain stable.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | TidyChat, TidyRoll |
| Files seen in | `/workspace/data/raw/TidyChat/TidyChat.lua:1089`, `/workspace/data/raw/TidyChat/TidyChat.lua:218`, `/workspace/data/raw/TidyChat/TidyChat.lua:808`, `/workspace/data/raw/TidyChat/TidyChat.lua:825`, `/workspace/data/raw/TidyChat/TidyChat.lua:868`, `/workspace/data/raw/TidyChat/TidyChat.lua:904`, `/workspace/data/raw/TidyChat/TidyChat.lua:930`, `/workspace/data/raw/TidyRoll/CustomAutoRoll.lua:140` |
| Namespaces detected | WindowGetShowing |
| Source kinds | lua_calls |
| Example locations | TidyChat: TidyChat.OnEntryBoxUpdateShowing, TidyChat: TidyChat.OnLBU, TidyChat: TidyChat.ToggleOptions, TidyChat: TidyChatFrames.Initialize, TidyChat: TidyChatHooks.DestroyWindowGroupHook, TidyChat: TidyChatHooks.OnEnterChatTextHook |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 10 |
| Global usage count | 10 |
| Local definition count | 0 |
| Documentation references | 0 |
| Initialization flow references | 0 |
| Known engine namespace | yes |
| Default UI presence | yes |
| Event binding presence | no |
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
WindowGetShowing(windowName)
```

## Description

Observed querying runtime window state or metadata.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| windowName | Observed as a target window name. | Observed values: c_TEXT_ENTRY_WINDOW.."EntryBox", c_TEXT_ENTRY_WINDOW.."EntryBoxLanguageButton", c_TIDY_CHAT_OPTIONS |

## Returns

- Not confidently inferable from addon-api docs alone.

## Side Effects

- No side effect is confidently inferable from addon-api docs alone.

## Seen In

- TidyChat
- TidyRoll

## Examples

- TidyChat: TidyChat.OnEntryBoxUpdateShowing -> WindowGetShowing(c_TEXT_ENTRY_WINDOW.."EntryBox")
- TidyChat: TidyChat.OnLBU -> WindowGetShowing(c_TEXT_ENTRY_WINDOW.."EntryBox")
- TidyChat: TidyChat.ToggleOptions -> WindowGetShowing(c_TIDY_CHAT_OPTIONS)
- TidyChat: TidyChatFrames.Initialize -> WindowGetShowing(c_TEXT_ENTRY_WINDOW.."EntryBoxLanguageButton")
- TidyChat: TidyChatHooks.DestroyWindowGroupHook -> WindowGetShowing(c_TIDY_CHAT_OPTIONS)
- TidyChat: TidyChatHooks.OnEnterChatTextHook -> WindowGetShowing(c_TEXT_ENTRY_WINDOW.."EntryBox")

## Related APIs

- none

## Used With

- [LayoutEditor.RegisterWindow](window_LayoutEditor.RegisterWindow.md) (HIGH 100/100) - Window Function
- [WindowAddAnchor](window_WindowAddAnchor.md) (HIGH 100/100) - Window Function
- [WindowClearAnchors](window_WindowClearAnchors.md) (HIGH 100/100) - Window Function
- [WindowRegisterCoreEventHandler](window_WindowRegisterCoreEventHandler.md) (HIGH 100/100) - Window Function
- [WindowSetShowing](window_WindowSetShowing.md) (HIGH 100/100) - Window Function
- [DoesWindowExist](../../globals/functions/global_DoesWindowExist.md) (HIGH 71/100) - Global Function

## Triggered By

- [OnHidden](../../events/window_events/window_event_OnHidden.md) (HIGH 100/100) - Window Event
- [OnRButtonUp](../../events/window_events/window_event_OnRButtonUp.md) (HIGH 100/100) - Window Event
- [OnShown](../../events/window_events/window_event_OnShown.md) (HIGH 100/100) - Window Event
- [SystemData.Events.L_BUTTON_UP_PROCESSED](../../events/game_events/game_event_SystemData.Events.L_BUTTON_UP_PROCESSED.md) (HIGH 100/100) - Game Event

## Affects

- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Notes

- none
