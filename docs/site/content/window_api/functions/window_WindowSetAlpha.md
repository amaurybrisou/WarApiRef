# WindowSetAlpha

- Category: Window Function
- Confidence level: HIGH
- Confidence score: 90/100
- Seen in: 1 addons

## Confidence Assessment

- Level: HIGH

- Score: 90/100

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, matches a known engine namespace, called globally with no local definition.

## Evidence Signals

- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +10 Argument pattern is consistent: Observed argument positions remain stable.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | TidyChat |
| Files seen in | `/workspace/data/raw/TidyChat/TidyChat.lua:329`, `/workspace/data/raw/TidyChat/TidyChat.lua:344`, `/workspace/data/raw/TidyChat/TidyChat.lua:404` |
| Namespaces detected | WindowSetAlpha |
| Source kinds | lua_calls |
| Example locations | TidyChat: TidyChatCore.SetTextEntry, TidyChat: TidyChatCore.SetWindowTabsHandleInput, TidyChat: TidyChatCore.SetWindowTextLog |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 5 |
| Global usage count | 5 |
| Local definition count | 0 |
| Documentation references | 0 |
| Initialization flow references | 0 |
| Known engine namespace | yes |
| Default UI presence | yes |
| Event binding presence | no |
| Observed in XML and Lua | no |
| Consistent role | no |
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
WindowSetAlpha(windowName, arg2)
```

## Description

Observed mutating runtime window state or presentation.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| windowName | Observed as a target window name. | Observed values: c_TEXT_ENTRY_WINDOW.."ChannelButton", c_TEXT_ENTRY_WINDOW.."EntryBoxBG", logDisplayName.."ToBottomButton" |
| arg2 | Observed as a runtime window or control identifier. | Observed values: (chatwindow_tabs_handle_input~=false and 1)or 0, (scrollbar_position==c_SCROLLBAR_POSITION_HIDDEN and 0)or scrollbar_alpha, textentry_background_alpha |

## Returns

- Not confidently inferable from addon-api docs alone.

## Side Effects

- Mutates engine-visible UI state or presentation.

## Seen In

- TidyChat

## Examples

- TidyChat: TidyChatCore.SetTextEntry -> WindowSetAlpha(c_TEXT_ENTRY_WINDOW.."ChannelButton", textentry_channel_alpha)
- TidyChat: TidyChatCore.SetTextEntry -> WindowSetAlpha(c_TEXT_ENTRY_WINDOW.."EntryBoxBG", textentry_background_alpha)
- TidyChat: TidyChatCore.SetWindowTabsHandleInput -> WindowSetAlpha(wndGroupName.."TabWindow", (chatwindow_tabs_handle_input~=false and 1)or 0)
- TidyChat: TidyChatCore.SetWindowTextLog -> WindowSetAlpha(scrollbarName, (scrollbar_position==c_SCROLLBAR_POSITION_HIDDEN and 0)or scrollbar_alpha)
- TidyChat: TidyChatCore.SetWindowTextLog -> WindowSetAlpha(logDisplayName.."ToBottomButton", (scrollbar_position==c_SCROLLBAR_POSITION_HIDDEN and 0)or scrollbar_alpha)

## Related APIs

- none

## Used With

- [WindowAddAnchor](window_WindowAddAnchor.md) (HIGH 100/100) - Window Function
- [WindowClearAnchors](window_WindowClearAnchors.md) (HIGH 100/100) - Window Function
- [WindowSetHandleInput](window_WindowSetHandleInput.md) (HIGH 100/100) - Window Function
- [WindowSetShowing](window_WindowSetShowing.md) (HIGH 100/100) - Window Function
- [WindowStopAlphaAnimation](window_WindowStopAlphaAnimation.md) (HIGH 90/100) - Window Function

## Triggered By

- none

## Affects

- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Notes

- Only one addon surfaced this symbol in the current corpus.
