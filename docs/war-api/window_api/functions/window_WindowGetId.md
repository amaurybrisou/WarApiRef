# WindowGetId

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
| Files seen in | `/workspace/data/raw/TidyChat/TidyChat.lua:1074`, `/workspace/data/raw/TidyChat/TidyChat.lua:2094`, `/workspace/data/raw/TidyChat/TidyChat.lua:2210`, `/workspace/data/raw/TidyRoll/CustomAutoRoll.lua:378`, `/workspace/data/raw/TidyRoll/CustomAutoRoll.lua:387`, `/workspace/data/raw/TidyRoll/TidyRollOptions.lua:847`, `/workspace/data/raw/TidyRoll/TidyRollOptions.lua:895` |
| Namespaces detected | WindowGetId |
| Source kinds | lua_calls |
| Example locations | TidyChat: TidyChat.Copy.OnCopyButton, TidyChat: TidyChat.Options.OnTabLBU, TidyChat: TidyChatFrames.OnToBottomButton, TidyRoll: TidyRoll.CustomAutoRoll.OnChoiceChange, TidyRoll: TidyRoll.CustomAutoRoll.OnDeleteButton, TidyRoll: TidyRollOptions.OnRadioLBU |
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
WindowGetId(windowName)
```

## Description

Observed querying runtime window state or metadata.

## Parameters

| Name | Role | Evidence |
| --- | --- | --- |
| windowName | Observed as a target window name. | Observed values: EA_Window_ContextMenu.activeWindow, SystemData.ActiveWindow.name, buttonName |

## Returns

- Not confidently inferable from addon-api docs alone.

## Side Effects

- No side effect is confidently inferable from addon-api docs alone.

## Seen In

- TidyChat
- TidyRoll

## Examples

- TidyChat: TidyChat.Copy.OnCopyButton -> WindowGetId(EA_Window_ContextMenu.activeWindow)
- TidyChat: TidyChat.Options.OnTabLBU -> WindowGetId(SystemData.ActiveWindow.name)
- TidyChat: TidyChatFrames.OnToBottomButton -> WindowGetId(EA_Window_ContextMenu.activeWindow)
- TidyRoll: TidyRoll.CustomAutoRoll.OnChoiceChange -> WindowGetId(SystemData.ActiveWindow.name)
- TidyRoll: TidyRoll.CustomAutoRoll.OnDeleteButton -> WindowGetId(buttonName)
- TidyRoll: TidyRollOptions.OnRadioLBU -> WindowGetId(radioName)

## Related APIs

- none

## Used With

- [OnLButtonUp](../../xml/handlers/handler_OnLButtonUp.md) (HIGH 100/100) - XML Event
- [SystemData.ActiveWindow.name](../../systemdata/fields/systemdata_SystemData.ActiveWindow.name.md) (HIGH 100/100) - SystemData Field

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
