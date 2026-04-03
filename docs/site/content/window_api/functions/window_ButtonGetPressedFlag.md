# ButtonGetPressedFlag

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
| Files seen in | `/workspace/data/raw/InfoScroller/libs/LibGUI.lua:755`, `/workspace/data/raw/InfoScroller/libs/LibGUI.lua:824`, `/workspace/data/raw/PartyCast/libs/LibGUI.lua:755`, `/workspace/data/raw/PartyCast/libs/LibGUI.lua:824`, `/workspace/data/raw/TidyChat/TidyChat.lua:1915`, `/workspace/data/raw/TidyChat/TidyChat.lua:1973`, `/workspace/data/raw/TidyChat/TidyChat.lua:2031`, `/workspace/data/raw/TidyRoll/TidyRollOptions.lua:746` |
| Namespaces detected | ButtonGetPressedFlag |
| Source kinds | lua_calls |
| Example locations | InfoScroller: LIBGUI_Checkbox:GetValue, InfoScroller: LIBGUI_Optionbutton:GetValue, PartyCast: LIBGUI_Checkbox:GetValue, PartyCast: LIBGUI_Optionbutton:GetValue, TidyChat: TidyChat.Options.OnApply, TidyChat: TidyChat.Options.OnCheckboxLBU |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 32 |
| Global usage count | 32 |
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

Observed as a window function across 4 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a function or method reference. | Observed values: TCHAT_LOGS_AUCTION_FILTER_CHECKBOX.."Button", TCHAT_LOGS_COPY_SHOWING_CHECKBOX.."Button", TCHAT_LOGS_LOOT_ROLL_FILTER_CHECKBOX.."Button" |

## Returns

- Not confidently inferable from addon-api docs alone.

## Side Effects

- No side effect is confidently inferable from addon-api docs alone.

## Seen In

- InfoScroller
- PartyCast
- TidyChat
- TidyRoll

## Examples

- InfoScroller: LIBGUI_Checkbox:GetValue -> ButtonGetPressedFlag(self.name)
- InfoScroller: LIBGUI_Optionbutton:GetValue -> ButtonGetPressedFlag(self.name)
- PartyCast: LIBGUI_Checkbox:GetValue -> ButtonGetPressedFlag(self.name)
- PartyCast: LIBGUI_Optionbutton:GetValue -> ButtonGetPressedFlag(self.name)
- TidyChat: TidyChat.Options.OnApply -> ButtonGetPressedFlag(TCHAT_TEXT_ENTRY_FREE_TEXT_ENTRY_CHECKBOX.."Button")
- TidyChat: TidyChat.Options.OnApply -> ButtonGetPressedFlag(TCHAT_TEXT_ENTRY_CHANNEL_SHOWING_CHECKBOX.."Button")

## Related APIs

- none

## Used With

- [Button](../../xml/element_types/element_Button.md) (HIGH 100/100) - XML Element Type
- [ButtonSetPressedFlag](window_ButtonSetPressedFlag.md) (HIGH 100/100) - Window Function
- [ComboBoxGetSelectedMenuItem](window_ComboBoxGetSelectedMenuItem.md) (HIGH 100/100) - Window Function
- [OnLButtonUp](../../events/window_events/window_event_OnLButtonUp.md) (HIGH 100/100) - Window Event
- [OnLButtonUp](../../xml/handlers/handler_OnLButtonUp.md) (HIGH 100/100) - XML Event

## Triggered By

- [OnLButtonUp](../../xml/handlers/handler_OnLButtonUp.md) (HIGH 100/100) - XML Event
- [OnLButtonUp](../../events/window_events/window_event_OnLButtonUp.md) (HIGH 100/100) - Window Event

## Affects

- [Button](../../xml/element_types/element_Button.md) (HIGH 100/100) - XML Element Type
- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Notes

- none
