# ComboBoxGetSelectedMenuItem

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
| Files seen in | `/workspace/data/raw/InfoScroller/libs/LibGUI.lua:1107`, `/workspace/data/raw/PartyCast/libs/LibGUI.lua:1107`, `/workspace/data/raw/TidyChat/TidyChat.lua:1973`, `/workspace/data/raw/TidyChat/TidyChat.lua:2117`, `/workspace/data/raw/TidyRoll/CustomAutoRoll.lua:378`, `/workspace/data/raw/TidyRoll/TidyRollOptions.lua:782` |
| Namespaces detected | ComboBoxGetSelectedMenuItem |
| Source kinds | lua_calls |
| Example locations | InfoScroller: LIBGUI_Combobox:SelectedIndex, PartyCast: LIBGUI_Combobox:SelectedIndex, TidyChat: TidyChat.Options.OnApply, TidyChat: TidyChat.Options.UpdateGroupTabs, TidyRoll: TidyRoll.CustomAutoRoll.OnChoiceChange, TidyRoll: TidyRollOptions.OnApply |
| XML usage count | 0 |
| XML attribute usage count | 0 |
| Lua usage count | 12 |
| Global usage count | 12 |
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
ComboBoxGetSelectedMenuItem(arg1)
```

## Description

Observed as a window function across 4 addons.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| arg1 | Observed as a runtime window or control identifier. | Observed values: TCHAT_TEXT_ENTRY_ANCHOR_POINT_COMBO, TCHAT_TEXT_ENTRY_RELATIVE_TO_COMBO, TCHAT_WINDOWS_SELECT_WINDOW_COMBO |

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

- InfoScroller: LIBGUI_Combobox:SelectedIndex -> ComboBoxGetSelectedMenuItem(self.name)
- PartyCast: LIBGUI_Combobox:SelectedIndex -> ComboBoxGetSelectedMenuItem(self.name)
- TidyChat: TidyChat.Options.OnApply -> ComboBoxGetSelectedMenuItem(TCHAT_TEXT_ENTRY_RELATIVE_TO_COMBO)
- TidyChat: TidyChat.Options.OnApply -> ComboBoxGetSelectedMenuItem(TCHAT_TEXT_ENTRY_ANCHOR_POINT_COMBO)
- TidyChat: TidyChat.Options.UpdateGroupTabs -> ComboBoxGetSelectedMenuItem(TCHAT_WINDOWS_SELECT_WINDOW_COMBO)
- TidyRoll: TidyRoll.CustomAutoRoll.OnChoiceChange -> ComboBoxGetSelectedMenuItem(comboName)

## Related APIs

- none

## Used With

- [Button](../../xml/element_types/element_Button.md) (HIGH 100/100) - XML Element Type
- [ButtonGetPressedFlag](window_ButtonGetPressedFlag.md) (HIGH 100/100) - Window Function
- [OnLButtonUp](../../events/window_events/window_event_OnLButtonUp.md) (HIGH 100/100) - Window Event
- [OnLButtonUp](../../xml/handlers/handler_OnLButtonUp.md) (HIGH 100/100) - XML Event

## Triggered By

- [OnLButtonUp](../../xml/handlers/handler_OnLButtonUp.md) (HIGH 100/100) - XML Event
- [OnLButtonUp](../../events/window_events/window_event_OnLButtonUp.md) (HIGH 100/100) - Window Event
- [OnSelChanged](../../xml/handlers/handler_OnSelChanged.md) (HIGH 100/100) - XML Event
- [OnSelChanged](../../events/window_events/window_event_OnSelChanged.md) (HIGH 100/100) - Window Event

## Affects

- [Button](../../xml/element_types/element_Button.md) (HIGH 100/100) - XML Element Type
- [ComboBox](../../xml/element_types/element_ComboBox.md) (HIGH 100/100) - XML Element Type
- [SystemData.ActiveWindow.name](../../systemdata/fields/systemdata_SystemData.ActiveWindow.name.md) (HIGH 100/100) - SystemData Field
- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Notes

- none
