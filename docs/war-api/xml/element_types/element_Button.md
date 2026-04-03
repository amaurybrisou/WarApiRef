# Button

- Category: XML Element Type
- Confidence level: HIGH
- Confidence score: 100/100

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 198

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, seen in 4 or more addons, used directly in xml handler attributes.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
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
| Addons seen in | PartyCast, Soloq, TidyChat, TidyRoll, minesweep |
| Files seen in | `/workspace/data/raw/PartyCast/PartyCast.xml:186`, `/workspace/data/raw/PartyCast/PartyCast.xml:327`, `/workspace/data/raw/PartyCast/PartyCast.xml:469`, `/workspace/data/raw/PartyCast/PartyCast.xml:631`, `/workspace/data/raw/PartyCast/PartyCast.xml:80`, `/workspace/data/raw/Soloq/ui/Overview.xml:65`, `/workspace/data/raw/TidyChat/TidyChat.xml:106`, `/workspace/data/raw/TidyChat/TidyChat.xml:114` |
| Namespaces detected | Button |
| Source kinds | xml_frames, xml_handlers |
| Example locations | PartyCast: PartyCastWindow_TemplateButton, PartyCast: PartyCastWindow_Template_LargeButton, PartyCast: PartyCastWindow_Template_PlainButton, PartyCast: PartyCastWindow_Template_SmallButton, PartyCast: PartyCastWindow_Template_SpecialButton, Soloq: SoloqOverviewWindowCloseButton |
| XML usage count | 33 |
| XML attribute usage count | 33 |
| Lua usage count | 3 |
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

Observed XML element type instantiated by 5 addons.

## Common Attributes

- name
- inherits
- layer
- handleinput
- drawchildrenfirst
- font
- backgroundtexture
- highlighttexture
- textureScale
- id
- textalign
- movable

## Common Handlers

- [OnLButtonUp](../handlers/handler_OnLButtonUp.md)
- [OnMouseOver](../handlers/handler_OnMouseOver.md)
- [OnRButtonUp](../handlers/handler_OnRButtonUp.md)

## Common Handler Functions

- TidyChat.Options.OnClose
- FrameManager.OnMouseOver
- Soloq.hideOverviewWindow
- TidyChat.Copy.CopyNext
- TidyChat.Copy.CopyPrev
- TidyChat.Copy.OnClose
- TidyChat.LootRoll.OnClose
- TidyChat.Options.OnApply
- TidyChat.Options.OnTabLBU
- TidyChat.Options.Reset
- TidyRoll.CustomAutoRoll.AddById
- TidyRoll.CustomAutoRoll.OnApply


## XML Event Bindings

| Event | Category | Common Lua Bindings | Expected Callback | Args Confidence |
|-------|----------|---------------------|-------------------|-----------------|
| [OnLButtonUp](../handlers/handler_OnLButtonUp.md) | input | TidyChat.Options.OnClose, Soloq.hideOverviewWindow, TidyChat.Copy.CopyNext, TidyChat.Copy.CopyPrev, TidyChat.Copy.OnClose, TidyChat.LootRoll.OnClose | `flags, x, y` | MEDIUM |
| [OnMouseOver](../handlers/handler_OnMouseOver.md) | input | FrameManager.OnMouseOver | `function()` | MEDIUM |
| [OnRButtonUp](../handlers/handler_OnRButtonUp.md) | input | minesweep.RButtonUp | `flags, x, y` | MEDIUM |

### Per-Event Lua API Calls

**OnLButtonUp** handlers call: `ButtonGetDisabledFlag`, `ButtonGetPressedFlag`, `ButtonSetPressedFlag`, `ComboBoxGetSelectedMenuItem`, `ComboBoxSetSelectedMenuItem`, `DestroyWindow`, `DynamicImageSetTextureSlice`, `LabelSetText`, `SliderBarGetCurrentPosition`, `SliderBarSetCurrentPosition`, `TextEditBoxGetText`, `TextEditBoxSetText`, `WindowGetId`, `WindowSetShowing`, `WindowSetTintColor`

**OnRButtonUp** handlers call: `DynamicImageSetTextureSlice`, `WindowGetId`, `WindowSetShowing`, `WindowSetTintColor`

## Common Inherits

- EA_Button_DefaultWindowClose
- EA_Button_Default
- EA_Button_DefaultResizeable
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
- [CircleImage](element_CircleImage.md)
- [FullResizeImage](element_FullResizeImage.md)
- [Label](element_Label.md)

## Common Structural Child Elements

- [TexSlices](element_TexSlices.md) — 5× (MEDIUM)

## Common Template Bases

- EA_Button_Default
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
| `inherits` | **required** | 93% | EA_Button_Default, EA_Button_DefaultWindowClose, TChatButton, TChatTabButton, ... |
| `layer` | optional | 30% | overlay, secondary, popup |
| `handleinput` | optional | 27% | false, true |
| `drawchildrenfirst` | optional | 21% | false, true |
| `font` | optional | 21% | font_chat_text, font_clear_default, font_clear_tiny |
| `backgroundtexture` | optional | 15% | EA_Training_Specialization |
| `highlighttexture` | optional | 15% | EA_Training_Specialization |
| `textureScale` | optional | 15% | 0.58, 0.7, 0.0, 0.9 |
| `id` | optional | 12% | 1, 2, 3, 4 |
| `textalign` | optional | 9% | center, bottomright |
| `movable` | optional | 3% | false |
| `popable` | optional | 3% | false |
| `savesettings` | optional | 3% | false |
| `texturescale` | optional | 3% | 0.85 |
## Structural Sub-Elements

### [TexSlices](element_TexSlices.md)

Observed 5 times as an unnamed child.

## Lua API Usage (from Handlers)

API functions commonly called from event handler Lua functions on this element type:

| API Function | Category | Call Count | From Events |
| --- | --- | --- | --- |
| `WindowSetShowing` | ui | 9 | OnLButtonUp, OnRButtonUp |
| `ButtonGetPressedFlag` | ui | 7 | OnLButtonUp |
| `ButtonSetPressedFlag` | ui | 7 | OnLButtonUp |
| `WindowGetId` | ui | 5 | OnLButtonUp, OnRButtonUp |
| `WindowSetTintColor` | ui | 4 | OnLButtonUp, OnRButtonUp |
| `ButtonGetDisabledFlag` | ui | 3 | OnLButtonUp |
| `ComboBoxGetSelectedMenuItem` | ui | 2 | OnLButtonUp |
| `ComboBoxSetSelectedMenuItem` | ui | 2 | OnLButtonUp |
| `DynamicImageSetTextureSlice` | ui | 2 | OnLButtonUp, OnRButtonUp |
| `SliderBarGetCurrentPosition` | ui | 2 | OnLButtonUp |
| `SliderBarSetCurrentPosition` | ui | 2 | OnLButtonUp |
| `TextEditBoxGetText` | ui | 2 | OnLButtonUp |
| `TextEditBoxSetText` | ui | 2 | OnLButtonUp |
| `DestroyWindow` | ui | 1 | OnLButtonUp |
| `LabelSetText` | ui | 1 | OnLButtonUp |
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

### OnRButtonUp

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `flags` | number | modifier_flags |
| 1 | `x` | number | mouse_x |
| 2 | `y` | number | mouse_y |
## Lua Functions Manipulating This Type

- TidyRollOptions.Initialize


## Binding Resolution

- Total handler declarations: 19
- Resolved to Lua functions: 18 (94%)

## Seen In

- PartyCast
- Soloq
- TidyChat
- TidyRoll
- minesweep

## Examples

- PartyCast: PartyCastWindow_TemplateButton -> Button PartyCastWindow_TemplateButton
- PartyCast: PartyCastWindow_Template_LargeButton -> Button PartyCastWindow_Template_LargeButton
- PartyCast: PartyCastWindow_Template_PlainButton -> Button PartyCastWindow_Template_PlainButton
- PartyCast: PartyCastWindow_Template_SmallButton -> Button PartyCastWindow_Template_SmallButton
- PartyCast: PartyCastWindow_Template_SpecialButton -> Button PartyCastWindow_Template_SpecialButton
- Soloq: SoloqOverviewWindowCloseButton -> Button SoloqOverviewWindowCloseButton

## Related APIs

- [ButtonGetPressedFlag](../../window_api/functions/window_ButtonGetPressedFlag.md) (HIGH 100/100) - Window Function
- [ButtonSetPressedFlag](../../window_api/functions/window_ButtonSetPressedFlag.md) (HIGH 100/100) - Window Function
- [ComboBoxGetSelectedMenuItem](../../window_api/functions/window_ComboBoxGetSelectedMenuItem.md) (HIGH 100/100) - Window Function
- [ComboBoxSetSelectedMenuItem](../../window_api/functions/window_ComboBoxSetSelectedMenuItem.md) (HIGH 100/100) - Window Function
- [DynamicImage](element_DynamicImage.md) (HIGH 100/100) - XML Element Type
- [DynamicImageSetTextureSlice](../../window_api/functions/window_DynamicImageSetTextureSlice.md) (HIGH 100/100) - Window Function
- [FullResizeImage](element_FullResizeImage.md) (HIGH 100/100) - XML Element Type
- [Label](element_Label.md) (HIGH 100/100) - XML Element Type
- [TextEditBoxGetText](../../window_api/functions/window_TextEditBoxGetText.md) (HIGH 100/100) - Window Function
- [Window](element_Window.md) (HIGH 100/100) - XML Element Type
- [WindowGetId](../../window_api/functions/window_WindowGetId.md) (HIGH 100/100) - Window Function
- [DestroyWindow](../../globals/functions/global_DestroyWindow.md) (HIGH 75/100) - Global Function
- [CircleImage](element_CircleImage.md) (MEDIUM 45/100) - XML Element Type
- [TexSlices](element_TexSlices.md) (MEDIUM 45/100) - XML Element Type

## Used With

- [ButtonGetPressedFlag](../../window_api/functions/window_ButtonGetPressedFlag.md) (HIGH 100/100) - Window Function
- [ComboBoxGetSelectedMenuItem](../../window_api/functions/window_ComboBoxGetSelectedMenuItem.md) (HIGH 100/100) - Window Function
- [OnLButtonUp](../../events/window_events/window_event_OnLButtonUp.md) (HIGH 100/100) - Window Event
- [OnLButtonUp](../handlers/handler_OnLButtonUp.md) (HIGH 100/100) - XML Event
- [Window](element_Window.md) (HIGH 100/100) - XML Element Type

## Triggered By

- [OnLButtonUp](../handlers/handler_OnLButtonUp.md) (HIGH 100/100) - XML Event
- [OnRButtonUp](../handlers/handler_OnRButtonUp.md) (HIGH 100/100) - XML Event
- [OnMouseOver](../handlers/handler_OnMouseOver.md) (HIGH 73/100) - XML Event

## Affects

- none
