# Window

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
| Addons seen in | Moth, TidyChat, TidyRoll |
| Files seen in | `/workspace/data/raw/Moth/Moth.xml:124`, `/workspace/data/raw/Moth/Moth.xml:153`, `/workspace/data/raw/Moth/Moth.xml:6`, `/workspace/data/raw/TidyChat/TidyChat.xml:138`, `/workspace/data/raw/TidyChat/TidyChat.xml:139`, `/workspace/data/raw/TidyChat/TidyChat.xml:140`, `/workspace/data/raw/TidyChat/TidyChat.xml:141`, `/workspace/data/raw/TidyChat/TidyChat.xml:151` |
| Namespaces detected | Window |
| Source kinds | xml_frames, xml_handlers |
| Example locations | Moth: Moth, Moth: MothCellTemplate, Moth: MothRowTemplate, TidyChat: TChatCheckboxTemplate, TidyChat: TChatTabLogsTemplate, TidyChat: TChatTabLogsTemplateAuctionFilterCheckbox |
| XML usage count | 48 |
| XML attribute usage count | 48 |
| Lua usage count | 8 |
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

Observed XML element type instantiated by 3 addons.

## Common Attributes

- name
- inherits
- layer
- movable
- handleinput
- savesettings
- popable
- draganddrop

## Common Handlers

- [OnHidden](../handlers/handler_OnHidden.md)
- [OnShown](../handlers/handler_OnShown.md)
- [OnLButtonUp](../handlers/handler_OnLButtonUp.md)
- [OnMouseWheel](../handlers/handler_OnMouseWheel.md)
- [OnInitialize](../handlers/handler_OnInitialize.md)
- [OnMButtonUp](../handlers/handler_OnMButtonUp.md)
- [OnRButtonUp](../handlers/handler_OnRButtonUp.md)
- [OnUpdate](../handlers/handler_OnUpdate.md)

## Common Handler Functions

- EA_Window_DefaultLabelToggleCircle.Initialize
- FrameManager.OnLButtonUp
- FrameManager.OnMouseWheel
- FrameManager.OnRButtonUp
- TidyChat.Copy.OnHidden
- TidyChat.Copy.OnMouseWheel
- TidyChat.Copy.OnShown
- TidyChat.LootRoll.OnHidden
- TidyChat.LootRoll.OnShown
- TidyChat.Options.OnCheckboxLBU
- TidyChat.Options.OnHidden
- TidyChat.Options.OnShown


## XML Event Bindings

| Event | Category | Common Lua Bindings | Expected Callback | Args Confidence |
|-------|----------|---------------------|-------------------|-----------------|
| [OnHidden](../handlers/handler_OnHidden.md) | lifecycle | TidyChat.Copy.OnHidden, TidyChat.LootRoll.OnHidden, TidyChat.Options.OnHidden, TidyRoll.CustomAutoRoll.OnHidden, TidyRoll.OnEsc, TidyRollOptions.OnHidden | `function()` | MEDIUM |
| [OnInitialize](../handlers/handler_OnInitialize.md) | lifecycle | EA_Window_DefaultLabelToggleCircle.Initialize | `function()` | MEDIUM |
| [OnLButtonUp](../handlers/handler_OnLButtonUp.md) | input | FrameManager.OnLButtonUp, TidyChat.Options.OnCheckboxLBU, TidyRollOptions.OnCheckboxLBU, TidyRollOptions.OnRadioLBU | `function(...)` | LOW |
| [OnMButtonUp](../handlers/handler_OnMButtonUp.md) | input | TidyRollFrame.OnMButtonUp | `flags, x, y` | MEDIUM |
| [OnMouseWheel](../handlers/handler_OnMouseWheel.md) | input | FrameManager.OnMouseWheel, TidyChat.Copy.OnMouseWheel | `function(delta)` | MEDIUM |
| [OnRButtonUp](../handlers/handler_OnRButtonUp.md) | input | FrameManager.OnRButtonUp | `function(...)` | LOW |
| [OnShown](../handlers/handler_OnShown.md) | lifecycle | TidyChat.Copy.OnShown, TidyChat.LootRoll.OnShown, TidyChat.Options.OnShown, TidyRoll.CustomAutoRoll.OnShown, TidyRollOptions.OnShown, WindowUtils.OnShown | `function()` | MEDIUM |
| [OnUpdate](../handlers/handler_OnUpdate.md) | lifecycle | TidyRoll.OnUpdate | `function(elapsed)` | MEDIUM |

### Per-Event Lua API Calls

**OnHidden** handlers call: `TextEditBoxSetText`, `WindowUtils.OnHidden`

**OnInitialize** handlers call: `ButtonSetText`, `CreateWindow`, `LabelSetText`

**OnLButtonUp** handlers call: `ButtonGetDisabledFlag`, `ButtonGetPressedFlag`, `ButtonSetDisabledFlag`, `ButtonSetPressedFlag`, `WindowGetId`

**OnShown** handlers call: `WindowUtils.OnShown`

## Common Inherits

- TChatCheckboxTemplate
- EA_Window_DefaultContextMenuFrame
- EA_TitleBar_Default
- EA_Window_DefaultBackgroundFrame
- EA_Window_DefaultButtonBottomFrame
- EA_Window_DefaultFrame
- EA_Window_DefaultFrameStatusBar
- TChatTabLogsTemplate
- TChatTabMiscTemplate
- TChatTabTextEntryTemplate
- TChatTabWindowsTemplate
- TRollOverlay

## Common Parent Elements

- [Window](element_Window.md)

## Common Named Child Elements

- [Label](element_Label.md)
- [Button](element_Button.md)
- [FullResizeImage](element_FullResizeImage.md)
- [Window](element_Window.md)
- [ComboBox](element_ComboBox.md)
- [DynamicImage](element_DynamicImage.md)
- [EditBox](element_EditBox.md)
- [ListBox](element_ListBox.md)

## Common Template Bases

- EA_TitleBar_Default
- EA_Window_DefaultBackgroundFrame
- EA_Window_DefaultButtonBottomFrame
- EA_Window_DefaultContextMenuFrame
- EA_Window_DefaultFrame
- EA_Window_DefaultFrameStatusBar
- TChatCheckboxTemplate
- TChatTabLogsTemplate
- TChatTabMiscTemplate
- TChatTabTextEntryTemplate
- TChatTabWindowsTemplate
- TRollOverlay


> **Note**: This element type commonly acts as a template base.

## Attribute Reference

| Attribute | Required | Usage % | Sample Values |
| --- | --- | --- | --- |
| `inherits` | optional | 52% | EA_Window_DefaultContextMenuFrame, TChatTabWindowsTemplate, TChatTabTextEntryTemplate, TChatTabLogsTemplate, ... |
| `layer` | optional | 37% | default, overlay, background, popup, ... |
| `movable` | optional | 33% | true, false |
| `handleinput` | optional | 25% | false, true |
| `savesettings` | optional | 22% | false, true |
| `popable` | optional | 16% | false, true |
| `draganddrop` | optional | 2% | true |
## Lua API Usage (from Handlers)

API functions commonly called from event handler Lua functions on this element type:

| API Function | Category | Call Count | From Events |
| --- | --- | --- | --- |
| `LabelSetText` | ui | 6 | OnInitialize |
| `WindowUtils.OnHidden` | ui | 6 | OnHidden |
| `WindowUtils.OnShown` | ui | 6 | OnShown |
| `ButtonGetPressedFlag` | ui | 4 | OnLButtonUp |
| `TextEditBoxSetText` | ui | 3 | OnHidden |
| `ButtonGetDisabledFlag` | ui | 2 | OnLButtonUp |
| `ButtonSetDisabledFlag` | ui | 2 | OnLButtonUp |
| `ButtonSetPressedFlag` | ui | 2 | OnLButtonUp |
| `ButtonSetText` | ui | 2 | OnInitialize |
| `CreateWindow` | ui | 1 | OnInitialize |
| `WindowGetId` | ui | 1 | OnLButtonUp |
## Handler Callback Signatures

Expected callback argument patterns for event handlers on this element type:

### OnHidden

Confidence: HIGH

### OnInitialize

Confidence: HIGH

### OnLButtonUp

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `flags` | number | modifier_flags |
| 1 | `x` | number | mouse_x |
| 2 | `y` | number | mouse_y |
### OnMButtonUp

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `flags` | number | modifier_flags |
| 1 | `x` | number | mouse_x |
| 2 | `y` | number | mouse_y |
### OnMouseWheel

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `x` | number | mouse_x |
| 1 | `y` | number | mouse_y |
| 2 | `delta` | number | wheel_delta |
### OnRButtonUp

Confidence: LOW

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `flags` | number | modifier_flags |
| 1 | `x` | number | mouse_x |
| 2 | `y` | number | mouse_y |
### OnShown

Confidence: HIGH

### OnUpdate

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `elapsed` | number | time_delta |
## Lua Functions Manipulating This Type

- Moth.Clear
- Moth.UpdateLevel
- TidyChat.LootRoll.OnRollLinkLButtonUp
- Moth.UpdateTarget
- Moth.Initialize
- Moth.UpdateHealthBar
- Moth.HideBorders
- Moth.HealthBar
- TidyRollOptions.Initialize
- Moth.Anchor


## Binding Resolution

- Total handler declarations: 22
- Resolved to Lua functions: 21 (95%)

## Seen In

- Moth
- TidyChat
- TidyRoll

## Examples

- Moth: Moth -> Window Moth
- Moth: MothCellTemplate -> Window MothCellTemplate
- Moth: MothRowTemplate -> Window MothRowTemplate
- TidyChat: TChatCheckboxTemplate -> Window TChatCheckboxTemplate
- TidyChat: TChatTabLogsTemplate -> Window TChatTabLogsTemplate
- TidyChat: TChatTabLogsTemplateAuctionFilterCheckbox -> Window TChatTabLogsTemplateAuctionFilterCheckbox

## Related APIs

- [ButtonGetDisabledFlag](../../window_api/functions/window_ButtonGetDisabledFlag.md) (HIGH 100/100) - Window Function
- [ButtonGetPressedFlag](../../window_api/functions/window_ButtonGetPressedFlag.md) (HIGH 100/100) - Window Function
- [ButtonSetPressedFlag](../../window_api/functions/window_ButtonSetPressedFlag.md) (HIGH 100/100) - Window Function
- [ButtonSetText](../../window_api/functions/window_ButtonSetText.md) (HIGH 100/100) - Window Function
- [ComboBoxGetSelectedMenuItem](../../window_api/functions/window_ComboBoxGetSelectedMenuItem.md) (HIGH 100/100) - Window Function
- [ComboBoxSetSelectedMenuItem](../../window_api/functions/window_ComboBoxSetSelectedMenuItem.md) (HIGH 100/100) - Window Function
- [LabelSetFont](../../window_api/functions/window_LabelSetFont.md) (HIGH 100/100) - Window Function
- [LabelSetText](../../window_api/functions/window_LabelSetText.md) (HIGH 100/100) - Window Function
- [LayoutEditor.RegisterWindow](../../window_api/functions/window_LayoutEditor.RegisterWindow.md) (HIGH 100/100) - Window Function
- [TextEditBoxSetText](../../window_api/functions/window_TextEditBoxSetText.md) (HIGH 100/100) - Window Function
- [WindowAddAnchor](../../window_api/functions/window_WindowAddAnchor.md) (HIGH 100/100) - Window Function
- [WindowClearAnchors](../../window_api/functions/window_WindowClearAnchors.md) (HIGH 100/100) - Window Function
- [WindowGetDimensions](../../window_api/functions/window_WindowGetDimensions.md) (HIGH 100/100) - Window Function
- [WindowGetId](../../window_api/functions/window_WindowGetId.md) (HIGH 100/100) - Window Function
- [WindowGetShowing](../../window_api/functions/window_WindowGetShowing.md) (HIGH 100/100) - Window Function
- [WindowRegisterCoreEventHandler](../../window_api/functions/window_WindowRegisterCoreEventHandler.md) (HIGH 100/100) - Window Function
- [WindowSetHandleInput](../../window_api/functions/window_WindowSetHandleInput.md) (HIGH 100/100) - Window Function
- [WindowSetId](../../window_api/functions/window_WindowSetId.md) (HIGH 100/100) - Window Function
- [WindowSetLayer](../../window_api/functions/window_WindowSetLayer.md) (HIGH 100/100) - Window Function
- [WindowSetShowing](../../window_api/functions/window_WindowSetShowing.md) (HIGH 100/100) - Window Function
- [WindowRegisterEventHandler](../../window_api/functions/window_WindowRegisterEventHandler.md) (HIGH 98/100) - Window Function
- [WindowSetDimensions](../../window_api/functions/window_WindowSetDimensions.md) (HIGH 98/100) - Window Function
- [WindowSetTintColor](../../window_api/functions/window_WindowSetTintColor.md) (HIGH 98/100) - Window Function
- [ButtonSetDisabledFlag](../../window_api/functions/window_ButtonSetDisabledFlag.md) (HIGH 90/100) - Window Function
- [ButtonSetTextColor](../../window_api/functions/window_ButtonSetTextColor.md) (HIGH 90/100) - Window Function
- [ComboBoxAddMenuItem](../../window_api/functions/window_ComboBoxAddMenuItem.md) (HIGH 90/100) - Window Function
- [ComboBoxClearMenuItems](../../window_api/functions/window_ComboBoxClearMenuItems.md) (HIGH 90/100) - Window Function
- [DynamicImageSetTexture](../../window_api/functions/window_DynamicImageSetTexture.md) (HIGH 90/100) - Window Function
- [LabelSetTextColor](../../window_api/functions/window_LabelSetTextColor.md) (HIGH 90/100) - Window Function
- [TextEditBoxGetText](../../window_api/functions/window_TextEditBoxGetText.md) (HIGH 90/100) - Window Function
- [WindowSetAlpha](../../window_api/functions/window_WindowSetAlpha.md) (HIGH 90/100) - Window Function
- [WindowSetParent](../../window_api/functions/window_WindowSetParent.md) (HIGH 90/100) - Window Function
- [WindowStopAlphaAnimation](../../window_api/functions/window_WindowStopAlphaAnimation.md) (HIGH 90/100) - Window Function
- [WindowUnregisterCoreEventHandler](../../window_api/functions/window_WindowUnregisterCoreEventHandler.md) (HIGH 90/100) - Window Function
- [ComboBoxGetSelectedText](../../window_api/functions/window_ComboBoxGetSelectedText.md) (HIGH 80/100) - Window Function
- [ComboBoxSetDisabledFlag](../../window_api/functions/window_ComboBoxSetDisabledFlag.md) (HIGH 80/100) - Window Function
- [DynamicImageSetTextureScale](../../window_api/functions/window_DynamicImageSetTextureScale.md) (HIGH 80/100) - Window Function
- [DynamicImageSetTextureSlice](../../window_api/functions/window_DynamicImageSetTextureSlice.md) (HIGH 80/100) - Window Function
- [LabelGetTextDimensions](../../window_api/functions/window_LabelGetTextDimensions.md) (HIGH 80/100) - Window Function
- [LayoutEditor.UnregisterWindow](../../window_api/functions/window_LayoutEditor.UnregisterWindow.md) (HIGH 80/100) - Window Function
- [StatusBarSetBackgroundTint](../../window_api/functions/window_StatusBarSetBackgroundTint.md) (HIGH 80/100) - Window Function
- [StatusBarSetCurrentValue](../../window_api/functions/window_StatusBarSetCurrentValue.md) (HIGH 80/100) - Window Function
- [StatusBarSetForegroundTint](../../window_api/functions/window_StatusBarSetForegroundTint.md) (HIGH 80/100) - Window Function
- [StatusBarSetMaximumValue](../../window_api/functions/window_StatusBarSetMaximumValue.md) (HIGH 80/100) - Window Function
- [WindowAssignFocus](../../window_api/functions/window_WindowAssignFocus.md) (HIGH 80/100) - Window Function
- [WindowGetAnchor](../../window_api/functions/window_WindowGetAnchor.md) (HIGH 80/100) - Window Function
- [WindowGetOffsetFromParent](../../window_api/functions/window_WindowGetOffsetFromParent.md) (HIGH 80/100) - Window Function
- [DoesWindowExist](../../globals/functions/global_DoesWindowExist.md) (HIGH 71/100) - Global Function

## Used With

- [OnHidden](../handlers/handler_OnHidden.md) (HIGH 100/100) - XML Event
- [OnLButtonUp](../handlers/handler_OnLButtonUp.md) (HIGH 100/100) - XML Event
- [OnMouseWheel](../handlers/handler_OnMouseWheel.md) (HIGH 100/100) - XML Event
- [OnShown](../handlers/handler_OnShown.md) (HIGH 100/100) - XML Event

## Triggered By

- [OnHidden](../handlers/handler_OnHidden.md) (HIGH 100/100) - XML Event
- [OnLButtonUp](../handlers/handler_OnLButtonUp.md) (HIGH 100/100) - XML Event
- [OnMouseWheel](../handlers/handler_OnMouseWheel.md) (HIGH 100/100) - XML Event
- [OnShown](../handlers/handler_OnShown.md) (HIGH 100/100) - XML Event
- [OnInitialize](../handlers/handler_OnInitialize.md) (HIGH 73/100) - XML Event
- [OnMButtonUp](../handlers/handler_OnMButtonUp.md) (HIGH 73/100) - XML Event
- [OnRButtonUp](../handlers/handler_OnRButtonUp.md) (HIGH 73/100) - XML Event
- [OnUpdate](../handlers/handler_OnUpdate.md) (HIGH 73/100) - XML Event

## Affects

- none
