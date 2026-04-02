# ComboBoxSetSelectedMenuItem

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
| Files seen in | `/workspace/data/raw/TidyChat/TidyChat.lua:1936`, `/workspace/data/raw/TidyRoll/TidyRollOptions.lua:746` |
| Namespaces detected | ComboBoxSetSelectedMenuItem |
| Source kinds | lua_calls |
| Example locations | TidyChat: TidyChat.Options.Reset, TidyRoll: TidyRollOptions.Reset |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 9 |
| Global usage count | 9 |
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
ComboBoxSetSelectedMenuItem(arg1, arg2)
```

## Description

Observed as a window function across 2 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a runtime window or control identifier. | Observed values: TCHAT_TEXT_ENTRY_ANCHOR_POINT_COMBO, TCHAT_TEXT_ENTRY_RELATIVE_TO_COMBO, c_TROLL_DIRECTION_COMBO |
| arg2 | Observed as a runtime window or control identifier. | Observed values: GetSetting("managment-bind-greed"), GetSetting("managment-bind-need"), GetSetting("managment-bind-pass") |

## Returns

- Not confidently inferable from addon-api docs alone.

## Side Effects

- Mutates engine-visible UI state or presentation.

## Seen In

- TidyChat
- TidyRoll

## Examples

- TidyChat: TidyChat.Options.Reset -> ComboBoxSetSelectedMenuItem(TCHAT_TEXT_ENTRY_RELATIVE_TO_COMBO, Settings.textentry_relative_to)
- TidyChat: TidyChat.Options.Reset -> ComboBoxSetSelectedMenuItem(TCHAT_TEXT_ENTRY_ANCHOR_POINT_COMBO, Settings.textentry_anchor_point)
- TidyRoll: TidyRollOptions.Reset -> ComboBoxSetSelectedMenuItem(c_TROLL_DIRECTION_COMBO, directionComboIndex[GetSetting("button-growth-direction")])
- TidyRoll: TidyRollOptions.Reset -> ComboBoxSetSelectedMenuItem(c_TROLL_ONESC_COMBO, GetSetting("managment-onesc-rollchoice")+2)
- TidyRoll: TidyRollOptions.Reset -> ComboBoxSetSelectedMenuItem(c_TROLL_NEED_COMBO, GetSetting("managment-bind-need"))
- TidyRoll: TidyRollOptions.Reset -> ComboBoxSetSelectedMenuItem(c_TROLL_GREED_COMBO, GetSetting("managment-bind-greed"))

## Related APIs

- none

## Used With

- [Button](../../xml/element_types/element_Button.md) (HIGH 100/100) - XML Element Type
- [ButtonGetPressedFlag](window_ButtonGetPressedFlag.md) (HIGH 100/100) - Window Function
- [ButtonSetPressedFlag](window_ButtonSetPressedFlag.md) (HIGH 100/100) - Window Function
- [OnLButtonUp](../../events/window_events/window_event_OnLButtonUp.md) (HIGH 100/100) - Window Event
- [OnLButtonUp](../../xml/handlers/handler_OnLButtonUp.md) (HIGH 100/100) - XML Handler
- [TextEditBoxSetText](window_TextEditBoxSetText.md) (HIGH 100/100) - Window Function
- [ButtonSetDisabledFlag](window_ButtonSetDisabledFlag.md) (HIGH 90/100) - Window Function

## Triggered By

- [OnLButtonUp](../../xml/handlers/handler_OnLButtonUp.md) (HIGH 100/100) - XML Handler
- [OnLButtonUp](../../events/window_events/window_event_OnLButtonUp.md) (HIGH 100/100) - Window Event

## Affects

- [Button](../../xml/element_types/element_Button.md) (HIGH 100/100) - XML Element Type
- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Notes

- none
