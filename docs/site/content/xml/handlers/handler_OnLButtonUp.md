# OnLButtonUp

- Type: XML Handler
- Confidence level: HIGH
- Confidence score: 100/100

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 138

- Rationale: Promoted as HIGH confidence because seen in 4 or more addons, used directly in xml handler attributes, referenced by generated docs or reference files.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +30 Used directly in XML handler attributes: XML exposure suggests an engine-level contract.
- +18 Used in event registration or dispatch: Observed in event-driven engine hooks.
- +20 Observed in both XML and Lua paths: Cross-source linkage reinforces platform-level usage.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +25 Referenced by generated docs or reference files: The symbol is reinforced outside a single call page.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | Soloq, TidyChat, TidyRoll, minesweep |
| Files seen in | `/workspace/data/raw/Soloq/ui/Overview.xml:69`, `/workspace/data/raw/TidyChat/TidyChat.xml:65`, `/workspace/data/raw/TidyChat/TidyChat.xml:76`, `/workspace/data/raw/TidyChat/TidyChat.xml:87`, `/workspace/data/raw/TidyChat/TidyChat.xml:99`, `/workspace/data/raw/TidyChat/TidyChatCopy.xml:34`, `/workspace/data/raw/TidyChat/TidyChatCopy.xml:61`, `/workspace/data/raw/TidyChat/TidyChatCopy.xml:83` |
| Namespaces detected | OnLButtonUp |
| Source kinds | bindings, xml_handlers |
| Example locations | Soloq: SoloqOverviewWindowCloseButton.OnLButtonUp, TidyChat: TChatCheckboxTemplate.OnLButtonUp, TidyChat: TChatTabButton.OnLButtonUp, TidyChat: TidyChatCopyClose.OnLButtonUp, TidyChat: TidyChatCopyNext.OnLButtonUp, TidyChat: TidyChatCopyPrev.OnLButtonUp |
| XML usage count | 22 |
| XML attribute usage count | 22 |
| Lua usage count | 22 |
| Global usage count | 0 |
| Local definition count | 0 |
| Documentation references | 1 |
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

Observed as an XML handler hook bound by 4 addons through frame event handlers.

## Expected Lua Binding

```lua
function(...)
```

## Element Types

- Button
- ListBox
- Window

## Seen In

- Soloq
- TidyChat
- TidyRoll
- minesweep

## Examples

- Soloq: SoloqOverviewWindowCloseButton -> SoloqOverviewWindowCloseButton.OnLButtonUp -> Soloq.hideOverviewWindow
- TidyChat: TChatCheckboxTemplate -> TChatCheckboxTemplate.OnLButtonUp -> TidyChat.Options.OnCheckboxLBU
- TidyChat: TChatTabButton -> TChatTabButton.OnLButtonUp -> TidyChat.Options.OnTabLBU
- TidyChat: TidyChatCopyClose -> TidyChatCopyClose.OnLButtonUp -> TidyChat.Copy.OnClose
- TidyChat: TidyChatCopyNext -> TidyChatCopyNext.OnLButtonUp -> TidyChat.Copy.CopyNext
- TidyChat: TidyChatCopyPrev -> TidyChatCopyPrev.OnLButtonUp -> TidyChat.Copy.CopyPrev

## Related APIs

- [Button](../element_types/element_Button.md) (HIGH 100/100) - XML Element Type
- [ButtonGetPressedFlag](../../window_api/functions/window_ButtonGetPressedFlag.md) (HIGH 100/100) - Window Function
- [ButtonSetPressedFlag](../../window_api/functions/window_ButtonSetPressedFlag.md) (HIGH 100/100) - Window Function
- [ComboBoxGetSelectedMenuItem](../../window_api/functions/window_ComboBoxGetSelectedMenuItem.md) (HIGH 100/100) - Window Function
- [ComboBoxSetSelectedMenuItem](../../window_api/functions/window_ComboBoxSetSelectedMenuItem.md) (HIGH 100/100) - Window Function
- [DynamicImageSetTextureSlice](../../window_api/functions/window_DynamicImageSetTextureSlice.md) (HIGH 100/100) - Window Function
- [ListBox](../element_types/element_ListBox.md) (HIGH 100/100) - XML Element Type
- [TextEditBoxGetText](../../window_api/functions/window_TextEditBoxGetText.md) (HIGH 100/100) - Window Function
- [Window](../element_types/element_Window.md) (HIGH 100/100) - XML Element Type
- [WindowGetId](../../window_api/functions/window_WindowGetId.md) (HIGH 100/100) - Window Function
- [DestroyWindow](../../globals/functions/global_DestroyWindow.md) (HIGH 75/100) - Global Function

## Used With

- [Button](../element_types/element_Button.md) (HIGH 100/100) - XML Element Type
- [ButtonGetPressedFlag](../../window_api/functions/window_ButtonGetPressedFlag.md) (HIGH 100/100) - Window Function
- [ComboBoxGetSelectedMenuItem](../../window_api/functions/window_ComboBoxGetSelectedMenuItem.md) (HIGH 100/100) - Window Function
- [OnLButtonUp](../../events/window_events/window_event_OnLButtonUp.md) (HIGH 100/100) - Window Event
- [Window](../element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Triggered By

- none

## Affects

- none

## Notes

- Expected binding arguments remain uncertain because addon-api docs capture symbol linkage, not full handler signatures.
