# WindowAddAnchor

- Category: Window Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 3 addons

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
| Addons seen in | Moth, TidyChat, TidyRoll |
| Files seen in | `/workspace/data/raw/Moth/Moth.lua:205`, `/workspace/data/raw/Moth/Moth.lua:578`, `/workspace/data/raw/TidyChat/TidyChat.lua:239`, `/workspace/data/raw/TidyChat/TidyChat.lua:344`, `/workspace/data/raw/TidyChat/TidyChat.lua:404`, `/workspace/data/raw/TidyChat/TidyChat.lua:930`, `/workspace/data/raw/TidyChat/TidyChat.lua:980`, `/workspace/data/raw/TidyRoll/TidyRollOptions.lua:136` |
| Namespaces detected | WindowAddAnchor |
| Source kinds | lua_calls |
| Example locations | Moth: Moth.Anchor, Moth: Moth.HealthBar, TidyChat: TidyChatCore.SetTextEntry, TidyChat: TidyChatCore.SetWindowGroup, TidyChat: TidyChatCore.SetWindowTextLog, TidyChat: TidyChatFrames.Initialize |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 90 |
| Global usage count | 90 |
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
WindowAddAnchor(windowName, point, relativeTo, relativePoint, offsetX, offsetY)
```

## Description

Observed positioning windows relative to other runtime UI elements.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| windowName | Observed as the window being positioned. | Observed values: "Moth", "MothHealthBar", c_CHANNEL_MENU.."AdviceButton" |
| point | Observed as an anchor point string. | Observed values: "bottom", "bottomleft", "bottomright" |
| relativeTo | Observed as a reference window name. | Observed values: "CursorWindow", "MothBackground", "Root" |
| relativePoint | Observed as a reference anchor point string. | Observed values: "bottom", "bottomleft", "bottomright" |
| offsetX | Observed as a numeric horizontal offset. | Observed values: -10, -2, -24 |
| offsetY | Observed as a numeric vertical offset. | Observed values: -10, -2, -29 |

## Returns

- Not confidently inferable from addon-api docs alone.

## Side Effects

- Mutates runtime window layout state.

## Seen In

- Moth
- TidyChat
- TidyRoll

## Examples

- Moth: Moth.Anchor -> WindowAddAnchor("Moth", "bottomleft", "CursorWindow", "topleft", 0, 0)
- Moth: Moth.HealthBar -> WindowAddAnchor("MothHealthBar", "bottomleft", "MothBackground", "topleft", 0, yOffset)
- TidyChat: TidyChatCore.SetTextEntry -> WindowAddAnchor(c_TEXT_ENTRY_WINDOW.."EntryBox", "topleft", c_TEXT_ENTRY_WINDOW, "topleft", EntryBoxAnchorXOffset, 0)
- TidyChat: TidyChatCore.SetTextEntry -> WindowAddAnchor(c_TEXT_ENTRY_WINDOW.."EntryBox", "bottomright", c_TEXT_ENTRY_WINDOW, "bottomright", 0, 0)
- TidyChat: TidyChatCore.SetTextEntry -> WindowAddAnchor(c_TEXT_ENTRY_WINDOW, TEAnchorPoint.."left", TEAnchorRelativeWindow, TEAnchorRelativePoint.."left", 0, 0)
- TidyChat: TidyChatCore.SetTextEntry -> WindowAddAnchor(c_TEXT_ENTRY_WINDOW, TEAnchorPoint.."right", TEAnchorRelativeWindow, TEAnchorRelativePoint.."right", 0, 0)

## Related APIs

- none

## Used With

- [ButtonSetText](window_ButtonSetText.md) (HIGH 100/100) - Window Function
- [LabelSetFont](window_LabelSetFont.md) (HIGH 100/100) - Window Function
- [LabelSetText](window_LabelSetText.md) (HIGH 100/100) - Window Function
- [LayoutEditor.RegisterWindow](window_LayoutEditor.RegisterWindow.md) (HIGH 100/100) - Window Function
- [SystemData.ChatLogFilters.ADVICE](../../systemdata/fields/systemdata_SystemData.ChatLogFilters.ADVICE.md) (HIGH 100/100) - SystemData Field
- [SystemData.ChatLogFilters.ALLIANCE](../../systemdata/fields/systemdata_SystemData.ChatLogFilters.ALLIANCE.md) (HIGH 100/100) - SystemData Field
- [SystemData.ChatLogFilters.SCENARIO](../../systemdata/fields/systemdata_SystemData.ChatLogFilters.SCENARIO.md) (HIGH 100/100) - SystemData Field
- [WindowClearAnchors](window_WindowClearAnchors.md) (HIGH 100/100) - Window Function
- [WindowGetShowing](window_WindowGetShowing.md) (HIGH 100/100) - Window Function
- [WindowRegisterCoreEventHandler](window_WindowRegisterCoreEventHandler.md) (HIGH 100/100) - Window Function
- [WindowSetHandleInput](window_WindowSetHandleInput.md) (HIGH 100/100) - Window Function
- [WindowSetLayer](window_WindowSetLayer.md) (HIGH 100/100) - Window Function
- [WindowSetShowing](window_WindowSetShowing.md) (HIGH 100/100) - Window Function
- [ComboBoxAddMenuItem](window_ComboBoxAddMenuItem.md) (HIGH 90/100) - Window Function
- [ComboBoxClearMenuItems](window_ComboBoxClearMenuItems.md) (HIGH 90/100) - Window Function
- [WindowSetAlpha](window_WindowSetAlpha.md) (HIGH 90/100) - Window Function
- [WindowGetOffsetFromParent](window_WindowGetOffsetFromParent.md) (HIGH 80/100) - Window Function
- [DoesWindowExist](../../globals/functions/global_DoesWindowExist.md) (HIGH 71/100) - Global Function

## Triggered By

- none

## Affects

- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Notes

- none
