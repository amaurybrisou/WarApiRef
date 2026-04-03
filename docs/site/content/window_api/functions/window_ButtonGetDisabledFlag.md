# ButtonGetDisabledFlag

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
| Files seen in | `/workspace/data/raw/TidyChat/TidyChat.lua:2031`, `/workspace/data/raw/TidyChat/TidyChat.lua:2275`, `/workspace/data/raw/TidyChat/TidyChat.lua:2287`, `/workspace/data/raw/TidyRoll/CustomAutoRoll.lua:387`, `/workspace/data/raw/TidyRoll/TidyRollOptions.lua:828` |
| Namespaces detected | ButtonGetDisabledFlag |
| Source kinds | lua_calls |
| Example locations | TidyChat: TidyChat.Copy.CopyNext, TidyChat: TidyChat.Copy.CopyPrev, TidyChat: TidyChat.Options.OnCheckboxLBU, TidyRoll: TidyRoll.CustomAutoRoll.OnDeleteButton, TidyRoll: TidyRollOptions.OnCheckboxLBU |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 5 |
| Global usage count | 5 |
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
ButtonGetDisabledFlag(arg1)
```

## Description

Observed as a window function across 2 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a runtime window or control identifier. | Observed values: buttonName, c_TIDY_CHAT_COPY.."Next", c_TIDY_CHAT_COPY.."Prev" |

## Returns

- Not confidently inferable from addon-api docs alone.

## Side Effects

- No side effect is confidently inferable from addon-api docs alone.

## Seen In

- TidyChat
- TidyRoll

## Examples

- TidyChat: TidyChat.Copy.CopyNext -> ButtonGetDisabledFlag(c_TIDY_CHAT_COPY.."Next")
- TidyChat: TidyChat.Copy.CopyPrev -> ButtonGetDisabledFlag(c_TIDY_CHAT_COPY.."Prev")
- TidyChat: TidyChat.Options.OnCheckboxLBU -> ButtonGetDisabledFlag(checkboxName)
- TidyRoll: TidyRoll.CustomAutoRoll.OnDeleteButton -> ButtonGetDisabledFlag(buttonName)
- TidyRoll: TidyRollOptions.OnCheckboxLBU -> ButtonGetDisabledFlag(checkboxName)

## Related APIs

- none

## Used With

- [Button](../../xml/element_types/element_Button.md) (HIGH 100/100) - XML Element Type
- [ButtonGetPressedFlag](window_ButtonGetPressedFlag.md) (HIGH 100/100) - Window Function
- [ButtonSetPressedFlag](window_ButtonSetPressedFlag.md) (HIGH 100/100) - Window Function
- [OnLButtonUp](../../xml/handlers/handler_OnLButtonUp.md) (HIGH 100/100) - XML Event
- [SystemData.MouseOverWindow.name](../../systemdata/fields/systemdata_SystemData.MouseOverWindow.name.md) (HIGH 100/100) - SystemData Field

## Triggered By

- [OnLButtonUp](../../xml/handlers/handler_OnLButtonUp.md) (HIGH 100/100) - XML Event
- [OnLButtonUp](../../events/window_events/window_event_OnLButtonUp.md) (HIGH 100/100) - Window Event

## Affects

- [Button](../../xml/element_types/element_Button.md) (HIGH 100/100) - XML Element Type
- [SystemData.ActiveWindow.name](../../systemdata/fields/systemdata_SystemData.ActiveWindow.name.md) (HIGH 100/100) - SystemData Field
- [SystemData.MouseOverWindow.name](../../systemdata/fields/systemdata_SystemData.MouseOverWindow.name.md) (HIGH 100/100) - SystemData Field
- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Notes

- none
