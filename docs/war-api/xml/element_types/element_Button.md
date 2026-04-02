# Button

- Category: XML Element Type
- Confidence level: HIGH
- Confidence score: 100/100

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 186

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, used directly in xml handler attributes, matches a known engine namespace.

## Evidence Signals

- +18 Seen in 2 to 3 addons: Cross-addon spread is present but limited.
- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +30 Used directly in XML handler attributes: XML exposure suggests an engine-level contract.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +18 Used in event registration or dispatch: Observed in event-driven engine hooks.
- +20 Observed in both XML and Lua paths: Cross-source linkage reinforces platform-level usage.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +25 Referenced by generated docs or reference files: The symbol is reinforced outside a single call page.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | TidyChat, TidyRoll |
| Files seen in | `/workspace/data/raw/TidyChat/TidyChat.xml:106`, `/workspace/data/raw/TidyChat/TidyChat.xml:114`, `/workspace/data/raw/TidyChat/TidyChat.xml:122`, `/workspace/data/raw/TidyChat/TidyChat.xml:130`, `/workspace/data/raw/TidyChat/TidyChat.xml:63`, `/workspace/data/raw/TidyChat/TidyChat.xml:69`, `/workspace/data/raw/TidyChat/TidyChat.xml:80`, `/workspace/data/raw/TidyChat/TidyChat.xml:92` |
| Namespaces detected | Button |
| Source kinds | xml_frames, xml_handlers |
| Example locations | TidyChat: TChatButton, TidyChat: TChatCheckboxTemplateButton, TidyChat: TChatTabButton, TidyChat: TidyChatCopyClose, TidyChat: TidyChatCopyNext, TidyChat: TidyChatCopyPrev |
| XML usage count | 25 |
| XML attribute usage count | 25 |
| Lua usage count | 2 |
| Global usage count | 0 |
| Local definition count | 0 |
| Documentation references | 1 |
| Initialization flow references | 0 |
| Known engine namespace | yes |
| Default UI presence | yes |
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

Observed XML element type instantiated by 2 addons.

## Common Attributes

- name
- inherits
- font
- id
- layer
- handleinput
- textalign
- drawchildrenfirst
- movable
- popable
- savesettings

## Common Handlers

- [OnLButtonUp](../handlers/handler_OnLButtonUp.md)
- [OnMouseOver](../handlers/handler_OnMouseOver.md)

## Common Handler Functions

- TidyChat.Options.OnClose
- FrameManager.OnMouseOver
- TidyChat.Copy.CopyNext
- TidyChat.Copy.CopyPrev
- TidyChat.Copy.OnClose
- TidyChat.LootRoll.OnClose
- TidyChat.Options.OnApply
- TidyChat.Options.OnTabLBU
- TidyChat.Options.Reset
- TidyRoll.CustomAutoRoll.AddById
- TidyRoll.CustomAutoRoll.OnApply
- TidyRoll.CustomAutoRoll.OnClose


## XML Event Bindings

| Event | Category | Common Lua Bindings | Expected Callback | Args Confidence |
|-------|----------|---------------------|-------------------|-----------------|
| [OnLButtonUp](../handlers/handler_OnLButtonUp.md) | input | TidyChat.Options.OnClose, TidyChat.Copy.CopyNext, TidyChat.Copy.CopyPrev, TidyChat.Copy.OnClose, TidyChat.LootRoll.OnClose, TidyChat.Options.OnApply | `flags, x, y` | MEDIUM |
| [OnMouseOver](../handlers/handler_OnMouseOver.md) | input | FrameManager.OnMouseOver | `function()` | MEDIUM |

### Per-Event Lua API Calls

**OnLButtonUp** handlers call: `ButtonGetDisabledFlag`, `ButtonGetPressedFlag`, `ButtonSetPressedFlag`, `ComboBoxGetSelectedMenuItem`, `ComboBoxSetSelectedMenuItem`, `SliderBarGetCurrentPosition`, `SliderBarSetCurrentPosition`, `TextEditBoxGetText`, `TextEditBoxSetText`, `WindowGetId`, `WindowSetShowing`

## Common Inherits

- EA_Button_DefaultResizeable
- EA_Button_DefaultWindowClose
- TChatTabButton
- TChatButton
- TRollButton
- EA_Button_DefaultCheckBox
- EA_Button_Tab
- EA_Button_DefaultToggleCircle
- TRollItemButton

## Common Parent Elements

- [Window](element_Window.md)

## Common Named Child Elements

- [DynamicImage](element_DynamicImage.md)

## Common Structural Child Elements

- [Disabled](element_Disabled.md)
- [Normal](element_Normal.md)
- [NormalHighlit](element_NormalHighlit.md)
- [Pressed](element_Pressed.md)

## Common Template Bases

- EA_Button_DefaultCheckBox
- EA_Button_DefaultResizeable
- EA_Button_DefaultToggleCircle
- EA_Button_DefaultWindowClose
- EA_Button_Tab
- TChatButton
- TChatTabButton
- TRollButton
- TRollItemButton


> **Note**: This element type commonly acts as a template base.

## Attribute Reference

| Attribute | Required | Usage % | Sample Values |
| --- | --- | --- | --- |
| `inherits` | **required** | 96% | EA_Button_DefaultWindowClose, TChatButton, TChatTabButton, EA_Button_DefaultResizeable, ... |
| `font` | optional | 24% | font_clear_small_bold, font_clear_default |
| `id` | optional | 16% | 1, 2, 3, 4 |
| `layer` | optional | 16% | secondary, popup |
| `handleinput` | optional | 12% | false |
| `textalign` | optional | 8% | center |
| `drawchildrenfirst` | optional | 4% | true |
| `movable` | optional | 4% | false |
| `popable` | optional | 4% | false |
| `savesettings` | optional | 4% | false |
## Lua API Usage (from Handlers)

API functions commonly called from event handler Lua functions on this element type:

| API Function | Category | Call Count | From Events |
| --- | --- | --- | --- |
| `ButtonGetPressedFlag` | ui | 7 | OnLButtonUp |
| `ButtonSetPressedFlag` | ui | 7 | OnLButtonUp |
| `WindowSetShowing` | ui | 6 | OnLButtonUp |
| `ButtonGetDisabledFlag` | ui | 3 | OnLButtonUp |
| `WindowGetId` | ui | 3 | OnLButtonUp |
| `ComboBoxGetSelectedMenuItem` | ui | 2 | OnLButtonUp |
| `ComboBoxSetSelectedMenuItem` | ui | 2 | OnLButtonUp |
| `SliderBarGetCurrentPosition` | ui | 2 | OnLButtonUp |
| `SliderBarSetCurrentPosition` | ui | 2 | OnLButtonUp |
| `TextEditBoxGetText` | ui | 2 | OnLButtonUp |
| `TextEditBoxSetText` | ui | 2 | OnLButtonUp |
## Handler Callback Signatures

Expected callback argument patterns for event handlers on this element type:

### OnLButtonUp

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `flags` | number | modifier_flags |
| 1 | `x` | number | mouse_x |
| 2 | `y` | number | mouse_y |
### OnMouseOver

Confidence: LOW

## Lua Functions Manipulating This Type

- TidyRollOptions.Initialize


## Binding Resolution

- Total handler declarations: 15
- Resolved to Lua functions: 14 (93%)

## Seen In

- TidyChat
- TidyRoll

## Examples

- TidyChat: TChatButton -> Button TChatButton
- TidyChat: TChatCheckboxTemplateButton -> Button TChatCheckboxTemplateButton
- TidyChat: TChatTabButton -> Button TChatTabButton
- TidyChat: TidyChatCopyClose -> Button TidyChatCopyClose
- TidyChat: TidyChatCopyNext -> Button TidyChatCopyNext
- TidyChat: TidyChatCopyPrev -> Button TidyChatCopyPrev

## Related APIs

- [ButtonGetDisabledFlag](../../window_api/functions/window_ButtonGetDisabledFlag.md) (HIGH 100/100) - Window Function
- [ButtonGetPressedFlag](../../window_api/functions/window_ButtonGetPressedFlag.md) (HIGH 100/100) - Window Function
- [ButtonSetPressedFlag](../../window_api/functions/window_ButtonSetPressedFlag.md) (HIGH 100/100) - Window Function
- [ComboBoxGetSelectedMenuItem](../../window_api/functions/window_ComboBoxGetSelectedMenuItem.md) (HIGH 100/100) - Window Function
- [ComboBoxSetSelectedMenuItem](../../window_api/functions/window_ComboBoxSetSelectedMenuItem.md) (HIGH 100/100) - Window Function
- [TextEditBoxSetText](../../window_api/functions/window_TextEditBoxSetText.md) (HIGH 100/100) - Window Function
- [WindowGetId](../../window_api/functions/window_WindowGetId.md) (HIGH 100/100) - Window Function
- [TextEditBoxGetText](../../window_api/functions/window_TextEditBoxGetText.md) (HIGH 90/100) - Window Function

## Used With

- [ButtonGetDisabledFlag](../../window_api/functions/window_ButtonGetDisabledFlag.md) (HIGH 100/100) - Window Function
- [ButtonGetPressedFlag](../../window_api/functions/window_ButtonGetPressedFlag.md) (HIGH 100/100) - Window Function
- [ButtonSetPressedFlag](../../window_api/functions/window_ButtonSetPressedFlag.md) (HIGH 100/100) - Window Function
- [ComboBoxGetSelectedMenuItem](../../window_api/functions/window_ComboBoxGetSelectedMenuItem.md) (HIGH 100/100) - Window Function
- [ComboBoxSetSelectedMenuItem](../../window_api/functions/window_ComboBoxSetSelectedMenuItem.md) (HIGH 100/100) - Window Function
- [OnLButtonUp](../../events/window_events/window_event_OnLButtonUp.md) (HIGH 100/100) - Window Event
- [OnLButtonUp](../handlers/handler_OnLButtonUp.md) (HIGH 100/100) - XML Handler
- [TextEditBoxSetText](../../window_api/functions/window_TextEditBoxSetText.md) (HIGH 100/100) - Window Function
- [TextEditBoxGetText](../../window_api/functions/window_TextEditBoxGetText.md) (HIGH 90/100) - Window Function

## Triggered By

- none

## Affects

- none
