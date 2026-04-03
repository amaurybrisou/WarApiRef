# TextEditBoxSetText

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
| Files seen in | `/workspace/data/raw/TidyChat/TidyChat.lua:2249`, `/workspace/data/raw/TidyRoll/CustomAutoRoll.lua:329`, `/workspace/data/raw/TidyRoll/CustomAutoRoll.lua:434`, `/workspace/data/raw/TidyRoll/TidyRollOptions.lua:746` |
| Namespaces detected | TextEditBoxSetText |
| Source kinds | lua_calls |
| Example locations | TidyChat: TidyChat.Copy.OnHidden, TidyRoll: TidyRoll.CustomAutoRoll.AddById, TidyRoll: TidyRoll.CustomAutoRoll.Reset, TidyRoll: TidyRollOptions.Reset |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 7 |
| Global usage count | 7 |
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
TextEditBoxSetText(windowName, text)
```

## Description

Observed reading from or writing to edit-box controls.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| windowName | Observed as a target control name. | Observed values: c_AUTO_ROLL_ADD_BY_ID_ID_EDITBOX, c_AUTO_ROLL_ADD_BY_ID_NAME_EDITBOX, c_TIDY_CHAT_COPY.."Log" |
| text | Observed as a text or wstring payload. | Observed values: L "", towstring(GetSetting("button-number")), towstring(GetSetting("button-offset")) |

## Returns

- Not confidently inferable from addon-api docs alone.

## Side Effects

- Mutates engine-visible UI state or presentation.

## Seen In

- TidyChat
- TidyRoll

## Examples

- TidyChat: TidyChat.Copy.OnHidden -> TextEditBoxSetText(c_TIDY_CHAT_COPY.."Log", L "")
- TidyRoll: TidyRoll.CustomAutoRoll.AddById -> TextEditBoxSetText(c_AUTO_ROLL_ADD_BY_ID_ID_EDITBOX, L "")
- TidyRoll: TidyRoll.CustomAutoRoll.AddById -> TextEditBoxSetText(c_AUTO_ROLL_ADD_BY_ID_NAME_EDITBOX, L "")
- TidyRoll: TidyRoll.CustomAutoRoll.Reset -> TextEditBoxSetText(c_AUTO_ROLL_ADD_BY_ID_ID_EDITBOX, L "")
- TidyRoll: TidyRoll.CustomAutoRoll.Reset -> TextEditBoxSetText(c_AUTO_ROLL_ADD_BY_ID_NAME_EDITBOX, L "")
- TidyRoll: TidyRollOptions.Reset -> TextEditBoxSetText(c_TROLL_BNUM_TBOX, towstring(GetSetting("button-number")))

## Related APIs

- none

## Used With

- [Button](../../xml/element_types/element_Button.md) (HIGH 100/100) - XML Element Type
- [ButtonGetPressedFlag](window_ButtonGetPressedFlag.md) (HIGH 100/100) - Window Function
- [ButtonSetPressedFlag](window_ButtonSetPressedFlag.md) (HIGH 100/100) - Window Function
- [ComboBoxSetSelectedMenuItem](window_ComboBoxSetSelectedMenuItem.md) (HIGH 100/100) - Window Function
- [OnLButtonUp](../../events/window_events/window_event_OnLButtonUp.md) (HIGH 100/100) - Window Event
- [OnLButtonUp](../../xml/handlers/handler_OnLButtonUp.md) (HIGH 100/100) - XML Event
- [ButtonSetDisabledFlag](window_ButtonSetDisabledFlag.md) (HIGH 90/100) - Window Function
- [TextEditBoxGetText](window_TextEditBoxGetText.md) (HIGH 90/100) - Window Function

## Triggered By

- [OnHidden](../../xml/handlers/handler_OnHidden.md) (HIGH 100/100) - XML Event
- [OnHidden](../../events/window_events/window_event_OnHidden.md) (HIGH 100/100) - Window Event
- [OnLButtonUp](../../xml/handlers/handler_OnLButtonUp.md) (HIGH 100/100) - XML Event
- [OnLButtonUp](../../events/window_events/window_event_OnLButtonUp.md) (HIGH 100/100) - Window Event

## Affects

- [Button](../../xml/element_types/element_Button.md) (HIGH 100/100) - XML Element Type
- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Notes

- none
