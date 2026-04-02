# WindowRegisterCoreEventHandler

- Category: Window Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 2 addons

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 141

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, matches a known engine namespace, called globally with no local definition.

## Evidence Signals

- +18 Seen in 2 to 3 addons: Cross-addon spread is present but limited.
- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +18 Used in event registration or dispatch: Observed in event-driven engine hooks.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +10 Argument pattern is consistent: Observed argument positions remain stable.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | TidyChat, TidyRoll |
| Files seen in | `/workspace/data/raw/TidyChat/TidyChat.lua:930`, `/workspace/data/raw/TidyRoll/TidyRollOptions.lua:136` |
| Namespaces detected | WindowRegisterCoreEventHandler |
| Source kinds | lua_calls |
| Example locations | TidyChat: TidyChatFrames.Initialize, TidyRoll: TidyRollOptions.Initialize |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 9 |
| Global usage count | 9 |
| Local definition count | 0 |
| Documentation references | 0 |
| Initialization flow references | 0 |
| Known engine namespace | yes |
| Default UI presence | yes |
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
WindowRegisterCoreEventHandler(windowName, windowEvent, handlerName)
```

## Description

Observed binding On* window events directly to a named window.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| windowName | Observed as a target window name. | Observed values: c_TEXT_ENTRY_WINDOW.."ChannelButton", c_TEXT_ENTRY_WINDOW.."EntryBox", c_TEXT_ENTRY_WINDOW.."EntryBoxTextInput" |
| windowEvent | Observed as an On* window event string. | Observed values: "OnHidden", "OnLButtonDown", "OnLButtonUp" |
| handlerName | Observed as a Lua handler reference. | Observed values: "TidyChat.OnChannelButtonMButtonUp", "TidyChat.OnEntryBoxTextInputLBD", "TidyChat.OnEntryBoxUpdateShowing" |

## Returns

- Not confidently inferable from addon-api docs alone.

## Side Effects

- Registers engine or library callbacks for later dispatch.

## Seen In

- TidyChat
- TidyRoll

## Examples

- TidyChat: TidyChatFrames.Initialize -> WindowRegisterCoreEventHandler(c_TEXT_ENTRY_WINDOW.."ChannelButton", "OnRButtonUp", "TidyChat.ToggleOptions")
- TidyChat: TidyChatFrames.Initialize -> WindowRegisterCoreEventHandler(c_TEXT_ENTRY_WINDOW.."ChannelButton", "OnMButtonUp", "TidyChat.OnChannelButtonMButtonUp")
- TidyChat: TidyChatFrames.Initialize -> WindowRegisterCoreEventHandler(c_TEXT_ENTRY_WINDOW.."EntryBox", "OnShown", "TidyChat.OnEntryBoxUpdateShowing")
- TidyChat: TidyChatFrames.Initialize -> WindowRegisterCoreEventHandler(c_TEXT_ENTRY_WINDOW.."EntryBox", "OnHidden", "TidyChat.OnEntryBoxUpdateShowing")
- TidyChat: TidyChatFrames.Initialize -> WindowRegisterCoreEventHandler(c_TEXT_ENTRY_WINDOW.."EntryBoxTextInput", "OnLButtonDown", "TidyChat.OnEntryBoxTextInputLBD")
- TidyRoll: TidyRollOptions.Initialize -> WindowRegisterCoreEventHandler(c_TROLL_CLOSE, "OnLButtonUp", "TidyRollOptions.OnClose")

## Related APIs

- none

## Used With

- [LayoutEditor.RegisterWindow](window_LayoutEditor.RegisterWindow.md) (HIGH 100/100) - Window Function
- [WindowAddAnchor](window_WindowAddAnchor.md) (HIGH 100/100) - Window Function
- [WindowClearAnchors](window_WindowClearAnchors.md) (HIGH 100/100) - Window Function
- [WindowGetShowing](window_WindowGetShowing.md) (HIGH 100/100) - Window Function
- [DoesWindowExist](../../globals/functions/global_DoesWindowExist.md) (HIGH 71/100) - Global Function

## Triggered By

- none

## Affects

- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Notes

- none
