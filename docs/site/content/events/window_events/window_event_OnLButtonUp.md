# OnLButtonUp

- Category: Window Event
- Confidence level: HIGH
- Confidence score: 100/100

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 176

- Rationale: Promoted as HIGH confidence because used directly in xml handler attributes, referenced by generated docs or reference files, called globally with no local definition.

## Evidence Signals

- +18 Seen in 2 to 3 addons: Cross-addon spread is present but limited.
- +30 Used directly in XML handler attributes: XML exposure suggests an engine-level contract.
- +20 Called globally with no local definition: No addon-local definition was observed in the generated corpus.
- +18 Used in event registration or dispatch: Observed in event-driven engine hooks.
- +10 Referenced from initialization flow: Lifecycle reconstruction references this symbol.
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
| Source kinds | event_page, examples, flows, lua_event_registration, xml_handlers |
| Example locations | TidyChat: TChatCheckboxTemplate.OnLButtonUp, TidyChat: TChatTabButton.OnLButtonUp, TidyChat: TidyChatCopyClose.OnLButtonUp, TidyChat: TidyChatCopyNext.OnLButtonUp, TidyChat: TidyChatCopyPrev.OnLButtonUp, TidyChat: TidyChatLootRollClose.OnLButtonUp |
| XML usage count | 19 |
| XML attribute usage count | 19 |
| Lua usage count | 4 |
| Global usage count | 4 |
| Local definition count | 0 |
| Documentation references | 4 |
| Initialization flow references | 1 |
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

Observed as an engine-supplied UI event hook used by 2 addons.

## Handler Pattern

Observed as an On* callback routed into a module-qualified Lua function.

## Payload

- Window callback arguments are not fully inferable from addon-api docs alone.

## Seen In

- TidyChat
- TidyRoll

## Registrars And Handlers

- FrameManager.OnLButtonUp
- TidyChat.Copy.CopyNext
- TidyChat.Copy.CopyPrev
- TidyChat.Copy.OnClose
- TidyChat.LootRoll.OnClose
- TidyChat.Options.OnApply
- TidyChat.Options.OnCheckboxLBU
- TidyChat.Options.OnClose
- TidyChat.Options.OnTabLBU
- TidyChat.Options.Reset
- TidyRoll.CustomAutoRoll.AddById
- TidyRoll.CustomAutoRoll.OnApply
- TidyRoll.CustomAutoRoll.OnClose
- TidyRoll.CustomAutoRoll.OnDeleteButton
- TidyRoll.CustomAutoRoll.OnListLbuttonUp
- TidyRoll.CustomAutoRoll.ToggleOptions
- TidyRollOptions.OnApply
- TidyRollOptions.OnCheckboxLBU
- TidyRollOptions.OnClose
- TidyRollOptions.OnRadioLBU
- TidyRollOptions.OnTabLBU
- TidyRollOptions.Reset
- WindowRegisterCoreEventHandler
- core

## Examples

- TidyRoll: TidyRollOptions.Initialize -> OnLButtonUp -> TidyRoll.CustomAutoRoll.ToggleOptions
- TidyRoll: TidyRollOptions.Initialize -> OnLButtonUp -> TidyRollOptions.OnApply
- TidyRoll: TidyRollOptions.Initialize -> OnLButtonUp -> TidyRollOptions.OnClose
- TidyRoll: TidyRollOptions.Initialize -> OnLButtonUp -> TidyRollOptions.Reset
- TidyChat: TChatCheckboxTemplate -> TChatCheckboxTemplate.OnLButtonUp -> TidyChat.Options.OnCheckboxLBU
- TidyChat: TChatTabButton -> TChatTabButton.OnLButtonUp -> TidyChat.Options.OnTabLBU

## Related APIs

- [ButtonGetDisabledFlag](../../window_api/functions/window_ButtonGetDisabledFlag.md) (HIGH 100/100) - Window Function
- [ButtonGetPressedFlag](../../window_api/functions/window_ButtonGetPressedFlag.md) (HIGH 100/100) - Window Function
- [ButtonSetPressedFlag](../../window_api/functions/window_ButtonSetPressedFlag.md) (HIGH 100/100) - Window Function
- [ComboBoxGetSelectedMenuItem](../../window_api/functions/window_ComboBoxGetSelectedMenuItem.md) (HIGH 100/100) - Window Function
- [ComboBoxSetSelectedMenuItem](../../window_api/functions/window_ComboBoxSetSelectedMenuItem.md) (HIGH 100/100) - Window Function
- [TextEditBoxSetText](../../window_api/functions/window_TextEditBoxSetText.md) (HIGH 100/100) - Window Function
- [WindowGetId](../../window_api/functions/window_WindowGetId.md) (HIGH 100/100) - Window Function
- [ButtonSetDisabledFlag](../../window_api/functions/window_ButtonSetDisabledFlag.md) (HIGH 90/100) - Window Function
- [TextEditBoxGetText](../../window_api/functions/window_TextEditBoxGetText.md) (HIGH 90/100) - Window Function
- [ComboBoxGetSelectedText](../../window_api/functions/window_ComboBoxGetSelectedText.md) (HIGH 80/100) - Window Function

## Used With

- [Button](../../xml/element_types/element_Button.md) (HIGH 100/100) - XML Element Type
- [ButtonGetPressedFlag](../../window_api/functions/window_ButtonGetPressedFlag.md) (HIGH 100/100) - Window Function
- [ButtonSetPressedFlag](../../window_api/functions/window_ButtonSetPressedFlag.md) (HIGH 100/100) - Window Function
- [ComboBoxGetSelectedMenuItem](../../window_api/functions/window_ComboBoxGetSelectedMenuItem.md) (HIGH 100/100) - Window Function
- [ComboBoxSetSelectedMenuItem](../../window_api/functions/window_ComboBoxSetSelectedMenuItem.md) (HIGH 100/100) - Window Function
- [OnLButtonUp](../../xml/handlers/handler_OnLButtonUp.md) (HIGH 100/100) - XML Handler
- [SystemData.MouseOverWindow.name](../../systemdata/fields/systemdata_SystemData.MouseOverWindow.name.md) (HIGH 100/100) - SystemData Field
- [TextEditBoxSetText](../../window_api/functions/window_TextEditBoxSetText.md) (HIGH 100/100) - Window Function
- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Triggered By

- none

## Affects

- none

## Notes

- none
