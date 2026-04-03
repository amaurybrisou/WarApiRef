# ButtonSetPressedFlag

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
| Files seen in | `/workspace/data/raw/TidyChat/TidyChat.lua:1936`, `/workspace/data/raw/TidyChat/TidyChat.lua:2031`, `/workspace/data/raw/TidyRoll/TidyRollOptions.lua:746`, `/workspace/data/raw/TidyRoll/TidyRollOptions.lua:828`, `/workspace/data/raw/TidyRoll/TidyRollOptions.lua:853` |
| Namespaces detected | ButtonSetPressedFlag |
| Source kinds | lua_calls |
| Example locations | TidyChat: TidyChat.Options.OnCheckboxLBU, TidyChat: TidyChat.Options.Reset, TidyRoll: TidyRollOptions.OnCheckboxLBU, TidyRoll: TidyRollOptions.RadioSetId, TidyRoll: TidyRollOptions.Reset |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 21 |
| Global usage count | 21 |
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
ButtonSetPressedFlag(arg1, arg2)
```

## Description

Observed mutating button text or pressed state on existing controls.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a function or method reference. | Observed values: TCHAT_LOGS_AUCTION_FILTER_CHECKBOX.."Button", TCHAT_LOGS_COPY_SHOWING_CHECKBOX.."Button", TCHAT_LOGS_LOOT_ROLL_FILTER_CHECKBOX.."Button" |
| arg2 | Observed as a runtime window or control identifier. | Observed values: GetSetting("auto-greed"), GetSetting("auto-need-for-medallions"), GetSetting("career-icon-show") |

## Returns

- Not confidently inferable from addon-api docs alone.

## Side Effects

- Mutates engine-visible UI state or presentation.

## Seen In

- TidyChat
- TidyRoll

## Examples

- TidyChat: TidyChat.Options.OnCheckboxLBU -> ButtonSetPressedFlag(checkboxName, not ButtonGetPressedFlag(checkboxName))
- TidyChat: TidyChat.Options.Reset -> ButtonSetPressedFlag(TCHAT_TEXT_ENTRY_FREE_TEXT_ENTRY_CHECKBOX.."Button", Settings.textentry_free)
- TidyChat: TidyChat.Options.Reset -> ButtonSetPressedFlag(TCHAT_TEXT_ENTRY_CHANNEL_SHOWING_CHECKBOX.."Button", Settings.textentry_channel_showing)
- TidyChat: TidyChat.Options.Reset -> ButtonSetPressedFlag(TCHAT_LOGS_AUCTION_FILTER_CHECKBOX.."Button", Settings.auction_filter_showing)
- TidyChat: TidyChat.Options.Reset -> ButtonSetPressedFlag(TCHAT_LOGS_LOOT_ROLL_FILTER_CHECKBOX.."Button", Settings.advanced_loot_roll_showing)
- TidyChat: TidyChat.Options.Reset -> ButtonSetPressedFlag(TCHAT_LOGS_COPY_SHOWING_CHECKBOX.."Button", Settings.copy_showing)

## Related APIs

- none

## Used With

- [Button](../../xml/element_types/element_Button.md) (HIGH 100/100) - XML Element Type
- [ButtonGetDisabledFlag](window_ButtonGetDisabledFlag.md) (HIGH 100/100) - Window Function
- [ButtonGetPressedFlag](window_ButtonGetPressedFlag.md) (HIGH 100/100) - Window Function
- [ComboBoxSetSelectedMenuItem](window_ComboBoxSetSelectedMenuItem.md) (HIGH 100/100) - Window Function
- [OnLButtonUp](../../events/window_events/window_event_OnLButtonUp.md) (HIGH 100/100) - Window Event
- [OnLButtonUp](../../xml/handlers/handler_OnLButtonUp.md) (HIGH 100/100) - XML Event
- [SystemData.MouseOverWindow.name](../../systemdata/fields/systemdata_SystemData.MouseOverWindow.name.md) (HIGH 100/100) - SystemData Field
- [TextEditBoxSetText](window_TextEditBoxSetText.md) (HIGH 100/100) - Window Function
- [ButtonSetDisabledFlag](window_ButtonSetDisabledFlag.md) (HIGH 90/100) - Window Function

## Triggered By

- [OnLButtonUp](../../xml/handlers/handler_OnLButtonUp.md) (HIGH 100/100) - XML Event
- [OnLButtonUp](../../events/window_events/window_event_OnLButtonUp.md) (HIGH 100/100) - Window Event

## Affects

- [Button](../../xml/element_types/element_Button.md) (HIGH 100/100) - XML Element Type
- [SystemData.MouseOverWindow.name](../../systemdata/fields/systemdata_SystemData.MouseOverWindow.name.md) (HIGH 100/100) - SystemData Field
- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Notes

- none
