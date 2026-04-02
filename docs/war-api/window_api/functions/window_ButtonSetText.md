# ButtonSetText

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
| Files seen in | `/workspace/data/raw/TidyChat/TidyChat.lua:2201`, `/workspace/data/raw/TidyChat/TidyChat.lua:980`, `/workspace/data/raw/TidyRoll/CustomAutoRoll.lua:145`, `/workspace/data/raw/TidyRoll/TidyRollOptions.lua:136` |
| Namespaces detected | ButtonSetText |
| Source kinds | lua_calls |
| Example locations | TidyChat: TidyChat.Copy.InitializeCopyWindow, TidyChat: TidyChatFrames.InitializeChannelMenuTidyChannelButtons, TidyRoll: TidyRoll.CustomAutoRoll.Initialize, TidyRoll: TidyRollOptions.Initialize |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 13 |
| Global usage count | 13 |
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

- TidyChat
- TidyRoll

## Examples

- TidyChat: TidyChat.Copy.InitializeCopyWindow -> ButtonSetText(c_TIDY_CHAT_COPY.."Prev", L "Prev "..entriesPerPage)
- TidyChat: TidyChat.Copy.InitializeCopyWindow -> ButtonSetText(c_TIDY_CHAT_COPY.."Next", L "Next "..entriesPerPage)
- TidyChat: TidyChatFrames.InitializeChannelMenuTidyChannelButtons -> ButtonSetText(c_CHANNEL_MENU.."AllianceButton", L "/"..GetChatString(StringTables.Chat.CHAT_CHANNEL_NAME_ALLIANCE))
- TidyChat: TidyChatFrames.InitializeChannelMenuTidyChannelButtons -> ButtonSetText(c_CHANNEL_MENU.."AdviceButton", L "/"..GetChatString(StringTables.Chat.CHAT_CHANNEL_NAME_ADVICE))
- TidyChat: TidyChatFrames.InitializeChannelMenuTidyChannelButtons -> ButtonSetText(c_CHANNEL_MENU.."ScenarioButton", L "/"..GetChatString(StringTables.Chat.CHAT_CHANNEL_NAME_SCENARIO))
- TidyRoll: TidyRoll.CustomAutoRoll.Initialize -> ButtonSetText(c_AUTO_ROLL_APPLY_BUTTON, L "Ok")

## Related APIs

- none

## Used With

- [LabelSetFont](window_LabelSetFont.md) (HIGH 100/100) - Window Function
- [LabelSetText](window_LabelSetText.md) (HIGH 100/100) - Window Function
- [SystemData.ChatLogFilters.ADVICE](../../systemdata/fields/systemdata_SystemData.ChatLogFilters.ADVICE.md) (HIGH 100/100) - SystemData Field
- [SystemData.ChatLogFilters.ALLIANCE](../../systemdata/fields/systemdata_SystemData.ChatLogFilters.ALLIANCE.md) (HIGH 100/100) - SystemData Field
- [SystemData.ChatLogFilters.SCENARIO](../../systemdata/fields/systemdata_SystemData.ChatLogFilters.SCENARIO.md) (HIGH 100/100) - SystemData Field
- [WindowAddAnchor](window_WindowAddAnchor.md) (HIGH 100/100) - Window Function
- [WindowClearAnchors](window_WindowClearAnchors.md) (HIGH 100/100) - Window Function
- [ComboBoxAddMenuItem](window_ComboBoxAddMenuItem.md) (HIGH 90/100) - Window Function
- [ComboBoxClearMenuItems](window_ComboBoxClearMenuItems.md) (HIGH 90/100) - Window Function

## Triggered By

- none

## Affects

- [SystemData.ChatLogFilters.ADVICE](../../systemdata/fields/systemdata_SystemData.ChatLogFilters.ADVICE.md) (HIGH 100/100) - SystemData Field
- [SystemData.ChatLogFilters.ALLIANCE](../../systemdata/fields/systemdata_SystemData.ChatLogFilters.ALLIANCE.md) (HIGH 100/100) - SystemData Field
- [SystemData.ChatLogFilters.SCENARIO](../../systemdata/fields/systemdata_SystemData.ChatLogFilters.SCENARIO.md) (HIGH 100/100) - SystemData Field
- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Notes

- none
