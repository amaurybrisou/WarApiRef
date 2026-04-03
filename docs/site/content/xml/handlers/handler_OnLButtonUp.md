# OnLButtonUp

- Type: XML Handler
- Confidence level: HIGH
- Confidence score: 100/100

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 146

- Rationale: Promoted as HIGH confidence because used directly in xml handler attributes, referenced by generated docs or reference files, observed in both xml and lua paths.

## Evidence Signals

- +18 Seen in 2 to 3 addons: Cross-addon spread is present but limited.
- +30 Used directly in XML handler attributes: XML exposure suggests an engine-level contract.
- +18 Used in event registration or dispatch: Observed in event-driven engine hooks.
- +20 Observed in both XML and Lua paths: Cross-source linkage reinforces platform-level usage.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +25 Referenced by generated docs or reference files: The symbol is reinforced outside a single call page.
- +20 Reinforced across multiple generated source types: Evidence comes from several independent addon-api source types.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | TidyChat, TidyRoll |
| Files seen in | `/workspace/data/raw/TidyChat/TidyChat.xml:65`, `/workspace/data/raw/TidyChat/TidyChat.xml:76`, `/workspace/data/raw/TidyChat/TidyChat.xml:87`, `/workspace/data/raw/TidyChat/TidyChat.xml:99`, `/workspace/data/raw/TidyChat/TidyChatCopy.xml:34`, `/workspace/data/raw/TidyChat/TidyChatCopy.xml:61`, `/workspace/data/raw/TidyChat/TidyChatCopy.xml:83`, `/workspace/data/raw/TidyChat/TidyChatLootRoll.xml:60` |
| Namespaces detected | OnLButtonUp |
| Source kinds | bindings, examples, xml_handlers |
| Example locations | TidyChat: TChatCheckboxTemplate.OnLButtonUp, TidyChat: TChatTabButton.OnLButtonUp, TidyChat: TidyChatCopyClose.OnLButtonUp, TidyChat: TidyChatCopyNext.OnLButtonUp, TidyChat: TidyChatCopyPrev.OnLButtonUp, TidyChat: TidyChatLootRollClose.OnLButtonUp |
| XML usage count | 19 |
| XML attribute usage count | 19 |
| Lua usage count | 19 |
| Global usage count | 0 |
| Local definition count | 0 |
| Documentation references | 2 |
| Initialization flow references | 0 |
| Known engine namespace | no |
| Default UI presence | no |
| Event binding presence | yes |
| Observed in XML and Lua | yes |
| Consistent role | yes |
| Consistent arguments | no |
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

## Description

Observed as an XML handler hook bound by 2 addons through frame event handlers.

## Expected Lua Binding

```lua
function(...)
```

## Element Types

- Button
- ListBox
- Window

## Seen In

- TidyChat
- TidyRoll

## Examples

- TidyChat: TChatCheckboxTemplate -> TChatCheckboxTemplate.OnLButtonUp -> TidyChat.Options.OnCheckboxLBU
- TidyChat: TChatTabButton -> TChatTabButton.OnLButtonUp -> TidyChat.Options.OnTabLBU
- TidyChat: TidyChatCopyClose -> TidyChatCopyClose.OnLButtonUp -> TidyChat.Copy.OnClose
- TidyChat: TidyChatCopyNext -> TidyChatCopyNext.OnLButtonUp -> TidyChat.Copy.CopyNext
- TidyChat: TidyChatCopyPrev -> TidyChatCopyPrev.OnLButtonUp -> TidyChat.Copy.CopyPrev
- TidyChat: TidyChatLootRollClose -> TidyChatLootRollClose.OnLButtonUp -> TidyChat.LootRoll.OnClose

## Related APIs

- [Button](../element_types/element_Button.md) (HIGH 100/100) - XML Element Type
- [ButtonGetDisabledFlag](../../window_api/functions/window_ButtonGetDisabledFlag.md) (HIGH 100/100) - Window Function
- [ButtonGetPressedFlag](../../window_api/functions/window_ButtonGetPressedFlag.md) (HIGH 100/100) - Window Function
- [ButtonSetPressedFlag](../../window_api/functions/window_ButtonSetPressedFlag.md) (HIGH 100/100) - Window Function
- [ComboBoxGetSelectedMenuItem](../../window_api/functions/window_ComboBoxGetSelectedMenuItem.md) (HIGH 100/100) - Window Function
- [ComboBoxSetSelectedMenuItem](../../window_api/functions/window_ComboBoxSetSelectedMenuItem.md) (HIGH 100/100) - Window Function
- [ListBox](../element_types/element_ListBox.md) (HIGH 100/100) - XML Element Type
- [TextEditBoxSetText](../../window_api/functions/window_TextEditBoxSetText.md) (HIGH 100/100) - Window Function
- [Window](../element_types/element_Window.md) (HIGH 100/100) - XML Element Type
- [WindowGetId](../../window_api/functions/window_WindowGetId.md) (HIGH 100/100) - Window Function
- [ButtonSetDisabledFlag](../../window_api/functions/window_ButtonSetDisabledFlag.md) (HIGH 90/100) - Window Function
- [TextEditBoxGetText](../../window_api/functions/window_TextEditBoxGetText.md) (HIGH 90/100) - Window Function

## Used With

- [Button](../element_types/element_Button.md) (HIGH 100/100) - XML Element Type
- [ButtonGetDisabledFlag](../../window_api/functions/window_ButtonGetDisabledFlag.md) (HIGH 100/100) - Window Function
- [ButtonGetPressedFlag](../../window_api/functions/window_ButtonGetPressedFlag.md) (HIGH 100/100) - Window Function
- [ButtonSetPressedFlag](../../window_api/functions/window_ButtonSetPressedFlag.md) (HIGH 100/100) - Window Function
- [ComboBoxGetSelectedMenuItem](../../window_api/functions/window_ComboBoxGetSelectedMenuItem.md) (HIGH 100/100) - Window Function
- [ComboBoxSetSelectedMenuItem](../../window_api/functions/window_ComboBoxSetSelectedMenuItem.md) (HIGH 100/100) - Window Function
- [OnLButtonUp](../../events/window_events/window_event_OnLButtonUp.md) (HIGH 100/100) - Window Event
- [TextEditBoxSetText](../../window_api/functions/window_TextEditBoxSetText.md) (HIGH 100/100) - Window Function
- [Window](../element_types/element_Window.md) (HIGH 100/100) - XML Element Type
- [WindowGetId](../../window_api/functions/window_WindowGetId.md) (HIGH 100/100) - Window Function
- [TextEditBoxGetText](../../window_api/functions/window_TextEditBoxGetText.md) (HIGH 90/100) - Window Function

## Triggered By

- none

## Affects

- none

## Notes

- Expected binding arguments remain uncertain because addon-api docs capture symbol linkage, not full handler signatures.
