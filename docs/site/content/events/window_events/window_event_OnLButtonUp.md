# OnLButtonUp

- Category: Window Event
- Confidence level: HIGH
- Confidence score: 100/100

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 188

- Rationale: Promoted as HIGH confidence because seen in 4 or more addons, used directly in xml handler attributes, referenced by generated docs or reference files.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
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
| Addons seen in | Soloq, TidyChat, TidyRoll, minesweep |
| Files seen in | `/workspace/data/raw/Soloq/ui/Overview.xml:69`, `/workspace/data/raw/TidyChat/TidyChat.xml:65`, `/workspace/data/raw/TidyChat/TidyChat.xml:76`, `/workspace/data/raw/TidyChat/TidyChat.xml:87`, `/workspace/data/raw/TidyChat/TidyChat.xml:99`, `/workspace/data/raw/TidyChat/TidyChatCopy.xml:34`, `/workspace/data/raw/TidyChat/TidyChatCopy.xml:61`, `/workspace/data/raw/TidyChat/TidyChatCopy.xml:83` |
| Namespaces detected | OnLButtonUp |
| Source kinds | event_page, flows, lua_event_registration, xml_handlers |
| Example locations | Soloq: SoloqOverviewWindowCloseButton.OnLButtonUp, TidyChat: TChatCheckboxTemplate.OnLButtonUp, TidyChat: TChatTabButton.OnLButtonUp, TidyChat: TidyChatCopyClose.OnLButtonUp, TidyChat: TidyChatCopyNext.OnLButtonUp, TidyChat: TidyChatCopyPrev.OnLButtonUp |
| XML usage count | 22 |
| XML attribute usage count | 22 |
| Lua usage count | 4 |
| Global usage count | 4 |
| Local definition count | 0 |
| Documentation references | 3 |
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

Observed as an engine-supplied UI event hook used by 4 addons.

## Handler Pattern

Observed as an On* callback routed into a module-qualified Lua function.

## Payload

- Window callback arguments are not fully inferable from addon-api docs alone.

## Seen In

- Soloq
- TidyChat
- TidyRoll
- minesweep

## Registrars And Handlers

- FrameManager.OnLButtonUp
- Soloq.hideOverviewWindow
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
- minesweep.Close
- minesweep.LButtonUp

## Examples

- TidyRoll: TidyRollOptions.Initialize -> OnLButtonUp -> TidyRoll.CustomAutoRoll.ToggleOptions
- TidyRoll: TidyRollOptions.Initialize -> OnLButtonUp -> TidyRollOptions.OnApply
- TidyRoll: TidyRollOptions.Initialize -> OnLButtonUp -> TidyRollOptions.OnClose
- TidyRoll: TidyRollOptions.Initialize -> OnLButtonUp -> TidyRollOptions.Reset
- Soloq: SoloqOverviewWindowCloseButton -> SoloqOverviewWindowCloseButton.OnLButtonUp -> Soloq.hideOverviewWindow
- TidyChat: TChatCheckboxTemplate -> TChatCheckboxTemplate.OnLButtonUp -> TidyChat.Options.OnCheckboxLBU

## Related APIs

- [ButtonGetPressedFlag](../../window_api/functions/window_ButtonGetPressedFlag.md) (HIGH 100/100) - Window Function
- [ButtonSetPressedFlag](../../window_api/functions/window_ButtonSetPressedFlag.md) (HIGH 100/100) - Window Function
- [ComboBoxGetSelectedMenuItem](../../window_api/functions/window_ComboBoxGetSelectedMenuItem.md) (HIGH 100/100) - Window Function
- [ComboBoxSetSelectedMenuItem](../../window_api/functions/window_ComboBoxSetSelectedMenuItem.md) (HIGH 100/100) - Window Function
- [DynamicImageSetTextureSlice](../../window_api/functions/window_DynamicImageSetTextureSlice.md) (HIGH 100/100) - Window Function
- [TextEditBoxGetText](../../window_api/functions/window_TextEditBoxGetText.md) (HIGH 100/100) - Window Function
- [WindowGetId](../../window_api/functions/window_WindowGetId.md) (HIGH 100/100) - Window Function
- [ComboBoxGetSelectedText](../../window_api/functions/window_ComboBoxGetSelectedText.md) (HIGH 80/100) - Window Function
- [DestroyWindow](../../globals/functions/global_DestroyWindow.md) (HIGH 75/100) - Global Function

## Used With

- [Button](../../xml/element_types/element_Button.md) (HIGH 100/100) - XML Element Type
- [ButtonGetPressedFlag](../../window_api/functions/window_ButtonGetPressedFlag.md) (HIGH 100/100) - Window Function
- [ButtonSetPressedFlag](../../window_api/functions/window_ButtonSetPressedFlag.md) (HIGH 100/100) - Window Function
- [ComboBoxGetSelectedMenuItem](../../window_api/functions/window_ComboBoxGetSelectedMenuItem.md) (HIGH 100/100) - Window Function
- [ComboBoxSetSelectedMenuItem](../../window_api/functions/window_ComboBoxSetSelectedMenuItem.md) (HIGH 100/100) - Window Function
- [OnLButtonUp](../../xml/handlers/handler_OnLButtonUp.md) (HIGH 100/100) - XML Event

## Triggered By

- none

## Affects

- none

## Notes

- none
