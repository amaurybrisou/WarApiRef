# ButtonGetPressedFlag

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
| Files seen in | `/workspace/data/raw/TidyChat/TidyChat.lua:1915`, `/workspace/data/raw/TidyChat/TidyChat.lua:1973`, `/workspace/data/raw/TidyChat/TidyChat.lua:2031`, `/workspace/data/raw/TidyRoll/TidyRollOptions.lua:746`, `/workspace/data/raw/TidyRoll/TidyRollOptions.lua:782`, `/workspace/data/raw/TidyRoll/TidyRollOptions.lua:828`, `/workspace/data/raw/TidyRoll/TidyRollOptions.lua:853`, `/workspace/data/raw/TidyRoll/TidyRollOptions.lua:865` |
| Namespaces detected | ButtonGetPressedFlag |
| Source kinds | lua_calls |
| Example locations | TidyChat: TidyChat.Options.OnApply, TidyChat: TidyChat.Options.OnCheckboxLBU, TidyChat: TidyChat.Options.UpdateDisabledFlags, TidyRoll: TidyRollOptions.OnApply, TidyRoll: TidyRollOptions.OnCheckboxLBU, TidyRoll: TidyRollOptions.RadioGetId |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 28 |
| Global usage count | 28 |
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
ButtonGetPressedFlag(arg1)
```

## Description

Observed as a window function across 2 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a function or method reference. | Observed values: TCHAT_LOGS_AUCTION_FILTER_CHECKBOX.."Button", TCHAT_LOGS_COPY_SHOWING_CHECKBOX.."Button", TCHAT_LOGS_LOOT_ROLL_FILTER_CHECKBOX.."Button" |

## Returns

- Not confidently inferable from addon-api docs alone.

## Side Effects

- No side effect is confidently inferable from addon-api docs alone.

## Seen In

- TidyChat
- TidyRoll

## Examples

- TidyChat: TidyChat.Options.OnApply -> ButtonGetPressedFlag(TCHAT_TEXT_ENTRY_FREE_TEXT_ENTRY_CHECKBOX.."Button")
- TidyChat: TidyChat.Options.OnApply -> ButtonGetPressedFlag(TCHAT_TEXT_ENTRY_CHANNEL_SHOWING_CHECKBOX.."Button")
- TidyChat: TidyChat.Options.OnApply -> ButtonGetPressedFlag(TCHAT_LOGS_AUCTION_FILTER_CHECKBOX.."Button")
- TidyChat: TidyChat.Options.OnApply -> ButtonGetPressedFlag(TCHAT_LOGS_LOOT_ROLL_FILTER_CHECKBOX.."Button")
- TidyChat: TidyChat.Options.OnApply -> ButtonGetPressedFlag(TCHAT_LOGS_COPY_SHOWING_CHECKBOX.."Button")
- TidyChat: TidyChat.Options.OnApply -> ButtonGetPressedFlag(TCHAT_MISC_SOCIAL_SHOWING_CHECKBOX.."Button")

## Related APIs

- none

## Used With

- [Button](../../xml/element_types/element_Button.md) (HIGH 100/100) - XML Element Type
- [ButtonGetDisabledFlag](window_ButtonGetDisabledFlag.md) (HIGH 100/100) - Window Function
- [ButtonSetPressedFlag](window_ButtonSetPressedFlag.md) (HIGH 100/100) - Window Function
- [ComboBoxGetSelectedMenuItem](window_ComboBoxGetSelectedMenuItem.md) (HIGH 100/100) - Window Function
- [ComboBoxSetSelectedMenuItem](window_ComboBoxSetSelectedMenuItem.md) (HIGH 100/100) - Window Function
- [OnLButtonUp](../../events/window_events/window_event_OnLButtonUp.md) (HIGH 100/100) - Window Event
- [OnLButtonUp](../../xml/handlers/handler_OnLButtonUp.md) (HIGH 100/100) - XML Handler
- [SystemData.MouseOverWindow.name](../../systemdata/fields/systemdata_SystemData.MouseOverWindow.name.md) (HIGH 100/100) - SystemData Field
- [TextEditBoxSetText](window_TextEditBoxSetText.md) (HIGH 100/100) - Window Function
- [ButtonSetDisabledFlag](window_ButtonSetDisabledFlag.md) (HIGH 90/100) - Window Function
- [TextEditBoxGetText](window_TextEditBoxGetText.md) (HIGH 90/100) - Window Function

## Triggered By

- [OnLButtonUp](../../xml/handlers/handler_OnLButtonUp.md) (HIGH 100/100) - XML Handler
- [OnLButtonUp](../../events/window_events/window_event_OnLButtonUp.md) (HIGH 100/100) - Window Event

## Affects

- [Button](../../xml/element_types/element_Button.md) (HIGH 100/100) - XML Element Type
- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Notes

- none
