# WindowSetHandleInput

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
| Files seen in | `/workspace/data/raw/TidyChat/TidyChat.lua:1089`, `/workspace/data/raw/TidyChat/TidyChat.lua:239`, `/workspace/data/raw/TidyChat/TidyChat.lua:329`, `/workspace/data/raw/TidyChat/TidyChat.lua:344`, `/workspace/data/raw/TidyRoll/TidyRollOptions.lua:136` |
| Namespaces detected | WindowSetHandleInput |
| Source kinds | lua_calls |
| Example locations | TidyChat: TidyChat.OnEntryBoxUpdateShowing, TidyChat: TidyChatCore.SetWindowGroup, TidyChat: TidyChatCore.SetWindowTabsHandleInput, TidyChat: TidyChatCore.SetWindowTextLog, TidyRoll: TidyRollOptions.Initialize |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 8 |
| Global usage count | 8 |
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
WindowSetHandleInput(windowName, arg2)
```

## Description

Observed mutating runtime window state or presentation.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| windowName | Observed as a target window name. | Observed values: c_TEXT_ENTRY_WINDOW, c_TIDY_ROLL_OPTIONS.."Background", c_TIDY_ROLL_OPTIONS.."Frame" |
| arg2 | Observed as a runtime window or control identifier. | Observed values: chatwindow_tabs_handle_input~=false, false, not chatwindow_click_through |

## Returns

- Not confidently inferable from addon-api docs alone.

## Side Effects

- Mutates engine-visible UI state or presentation.

## Seen In

- TidyChat
- TidyRoll

## Examples

- TidyChat: TidyChat.OnEntryBoxUpdateShowing -> WindowSetHandleInput(c_TEXT_ENTRY_WINDOW, textEntryShowing)
- TidyChat: TidyChatCore.SetWindowGroup -> WindowSetHandleInput(wndGroupName, not chatwindow_click_through)
- TidyChat: TidyChatCore.SetWindowTabsHandleInput -> WindowSetHandleInput(wndGroupName.."TabWindow", chatwindow_tabs_handle_input~=false)
- TidyChat: TidyChatCore.SetWindowTextLog -> WindowSetHandleInput(scrollbarName, scrollbar_position~=c_SCROLLBAR_POSITION_HIDDEN)
- TidyRoll: TidyRollOptions.Initialize -> WindowSetHandleInput(c_TIDY_ROLL_OPTIONS.."Background", false)
- TidyRoll: TidyRollOptions.Initialize -> WindowSetHandleInput(c_TIDY_ROLL_OPTIONS.."Frame", false)

## Related APIs

- none

## Used With

- [WindowAddAnchor](window_WindowAddAnchor.md) (HIGH 100/100) - Window Function
- [WindowClearAnchors](window_WindowClearAnchors.md) (HIGH 100/100) - Window Function
- [WindowSetLayer](window_WindowSetLayer.md) (HIGH 100/100) - Window Function
- [WindowSetShowing](window_WindowSetShowing.md) (HIGH 100/100) - Window Function
- [WindowSetAlpha](window_WindowSetAlpha.md) (HIGH 90/100) - Window Function
- [WindowStopAlphaAnimation](window_WindowStopAlphaAnimation.md) (HIGH 90/100) - Window Function
- [WindowGetOffsetFromParent](window_WindowGetOffsetFromParent.md) (HIGH 80/100) - Window Function

## Triggered By

- [OnHidden](../../events/window_events/window_event_OnHidden.md) (HIGH 100/100) - Window Event
- [OnShown](../../events/window_events/window_event_OnShown.md) (HIGH 100/100) - Window Event

## Affects

- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Notes

- none
