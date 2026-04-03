# ButtonSetPressedFlag

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
| Files seen in | `/workspace/data/raw/InfoScroller/libs/LibGUI.lua:760`, `/workspace/data/raw/InfoScroller/libs/LibGUI.lua:829`, `/workspace/data/raw/PartyCast/libs/LibGUI.lua:760`, `/workspace/data/raw/PartyCast/libs/LibGUI.lua:829`, `/workspace/data/raw/TidyChat/TidyChat.lua:1936`, `/workspace/data/raw/TidyChat/TidyChat.lua:2031`, `/workspace/data/raw/TidyRoll/TidyRollOptions.lua:746`, `/workspace/data/raw/TidyRoll/TidyRollOptions.lua:828` |
| Namespaces detected | ButtonSetPressedFlag |
| Source kinds | lua_calls |
| Example locations | InfoScroller: LIBGUI_Checkbox:SetValue, InfoScroller: LIBGUI_Optionbutton:SetValue, PartyCast: LIBGUI_Checkbox:SetValue, PartyCast: LIBGUI_Optionbutton:SetValue, TidyChat: TidyChat.Options.OnCheckboxLBU, TidyChat: TidyChat.Options.Reset |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 25 |
| Global usage count | 25 |
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

- InfoScroller
- PartyCast
- TidyChat
- TidyRoll

## Examples

- InfoScroller: LIBGUI_Checkbox:SetValue -> ButtonSetPressedFlag(self.name, checkState)
- InfoScroller: LIBGUI_Optionbutton:SetValue -> ButtonSetPressedFlag(self.name, checkState)
- PartyCast: LIBGUI_Checkbox:SetValue -> ButtonSetPressedFlag(self.name, checkState)
- PartyCast: LIBGUI_Optionbutton:SetValue -> ButtonSetPressedFlag(self.name, checkState)
- TidyChat: TidyChat.Options.OnCheckboxLBU -> ButtonSetPressedFlag(checkboxName, not ButtonGetPressedFlag(checkboxName))
- TidyChat: TidyChat.Options.Reset -> ButtonSetPressedFlag(TCHAT_TEXT_ENTRY_FREE_TEXT_ENTRY_CHECKBOX.."Button", Settings.textentry_free)

## Related APIs

- none

## Used With

- [ButtonGetPressedFlag](window_ButtonGetPressedFlag.md) (HIGH 100/100) - Window Function
- [ComboBoxSetSelectedMenuItem](window_ComboBoxSetSelectedMenuItem.md) (HIGH 100/100) - Window Function
- [OnLButtonUp](../../events/window_events/window_event_OnLButtonUp.md) (HIGH 100/100) - Window Event

## Triggered By

- [OnLButtonUp](../../xml/handlers/handler_OnLButtonUp.md) (HIGH 100/100) - XML Event
- [OnLButtonUp](../../events/window_events/window_event_OnLButtonUp.md) (HIGH 100/100) - Window Event

## Affects

- [Button](../../xml/element_types/element_Button.md) (HIGH 100/100) - XML Element Type
- [SystemData.MouseOverWindow.name](../../systemdata/fields/systemdata_SystemData.MouseOverWindow.name.md) (HIGH 100/100) - SystemData Field
- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Notes

- none
