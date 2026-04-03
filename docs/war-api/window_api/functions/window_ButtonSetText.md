# ButtonSetText

- Category: Window Function
- Confidence level: HIGH
- Confidence score: 100/100
- Seen in: 4 addons

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 135

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, seen in 4 or more addons, matches a known engine namespace.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +10 Argument pattern is consistent: Observed argument positions remain stable.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | InfoScroller, PartyCast, TidyChat, TidyRoll |
| Files seen in | `/workspace/data/raw/InfoScroller/libs/LibGUI.lua:533`, `/workspace/data/raw/PartyCast/libs/LibGUI.lua:533`, `/workspace/data/raw/TidyChat/TidyChat.lua:2201`, `/workspace/data/raw/TidyChat/TidyChat.lua:980`, `/workspace/data/raw/TidyRoll/CustomAutoRoll.lua:145`, `/workspace/data/raw/TidyRoll/TidyRollOptions.lua:136` |
| Namespaces detected | ButtonSetText |
| Source kinds | lua_calls |
| Example locations | InfoScroller: LIBGUI_Button:SetText, PartyCast: LIBGUI_Button:SetText, TidyChat: TidyChat.Copy.InitializeCopyWindow, TidyChat: TidyChatFrames.InitializeChannelMenuTidyChannelButtons, TidyRoll: TidyRoll.CustomAutoRoll.Initialize, TidyRoll: TidyRollOptions.Initialize |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 17 |
| Global usage count | 17 |
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
ButtonSetText(windowName, text)
```

## Description

Observed mutating button text or pressed state on existing controls.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| windowName | Observed as a target control name. | Observed values: c_AUTO_ROLL_ADD_BY_ID_BUTTON, c_AUTO_ROLL_APPLY_BUTTON, c_CHANNEL_MENU.."AdviceButton" |
| text | Observed as a text or wstring payload. | Observed values: L "/"..GetChatString(StringTables.Chat.CHAT_CHANNEL_NAME_ADVICE), L "/"..GetChatString(StringTables.Chat.CHAT_CHANNEL_NAME_ALLIANCE), L "/"..GetChatString(StringTables.Chat.CHAT_CHANNEL_NAME_SCENARIO) |

## Returns

- Not confidently inferable from addon-api docs alone.

## Side Effects

- Mutates engine-visible UI state or presentation.

## Seen In

- InfoScroller
- PartyCast
- TidyChat
- TidyRoll

## Examples

- InfoScroller: LIBGUI_Button:SetText -> ButtonSetText(self.name, text and L "true" or L "false")
- InfoScroller: LIBGUI_Button:SetText -> ButtonSetText(self.name, towstring(text))
- PartyCast: LIBGUI_Button:SetText -> ButtonSetText(self.name, text and L "true" or L "false")
- PartyCast: LIBGUI_Button:SetText -> ButtonSetText(self.name, towstring(text))
- TidyChat: TidyChat.Copy.InitializeCopyWindow -> ButtonSetText(c_TIDY_CHAT_COPY.."Prev", L "Prev "..entriesPerPage)
- TidyChat: TidyChat.Copy.InitializeCopyWindow -> ButtonSetText(c_TIDY_CHAT_COPY.."Next", L "Next "..entriesPerPage)

## Related APIs

- none

## Used With

- [ComboBoxAddMenuItem](window_ComboBoxAddMenuItem.md) (HIGH 100/100) - Window Function
- [ComboBoxClearMenuItems](window_ComboBoxClearMenuItems.md) (HIGH 100/100) - Window Function
- [LabelSetFont](window_LabelSetFont.md) (HIGH 100/100) - Window Function
- [LabelSetText](window_LabelSetText.md) (HIGH 100/100) - Window Function
- [CreateWindow](../../globals/functions/global_CreateWindow.md) (HIGH 75/100) - Global Function
- [CreateWindowFromTemplate](../../globals/functions/global_CreateWindowFromTemplate.md) (HIGH 75/100) - Global Function
- [towstring](../../globals/functions/global_towstring.md) (HIGH 75/100) - Global Function

## Triggered By

- none

## Affects

- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Notes

- none
