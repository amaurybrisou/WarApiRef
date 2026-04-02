# EditBox

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
| Addons seen in | AdvancedRenownTrainer, Aura, BuffHead, Busted, EA_UiDebugTools, EA_UiModWindow, Enemy, Killer |
| Files seen in | `/workspace_addons/Aura/Source/AuraConfig.xml:1096`, `/workspace_addons/Aura/Source/AuraConfig.xml:1137`, `/workspace_addons/Aura/Source/AuraConfig.xml:116`, `/workspace_addons/Aura/Source/AuraConfig.xml:1202`, `/workspace_addons/Aura/Source/AuraConfig.xml:1243`, `/workspace_addons/Aura/Source/AuraConfig.xml:1347`, `/workspace_addons/Aura/Source/AuraConfig.xml:1471`, `/workspace_addons/Aura/Source/AuraConfig.xml:1559` |
| Namespaces detected | EditBox |
| Source kinds | xml_frames, xml_handlers |
| Example locations | AdvancedRenownTrainer: AdvancedRenownTrainingImportNameInputWindowNameInputBox, AdvancedRenownTrainer: AdvancedRenownTrainingImportWindowLinkInputBox, AdvancedRenownTrainer: AdvancedRenownTrainingImportWindowNameInputBox, AdvancedRenownTrainer: AdvancedRenownTrainingLinkWindowLinkBox, AdvancedRenownTrainer: AdvancedRenownTrainingPresetsWindowSaveNameInput, Aura: AuraConfigActivationAlertTextText |
| XML usage count | 212 |
| XML attribute usage count | 212 |
| Lua usage count | 6 |
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

Observed XML element type instantiated by 24 addons.

## Common Attributes

- name
- inherits
- maxchars
- input
- layer
- warnOnTextCropped
- font
- taborder
- handleinput
- scrolling
- maxChars
- background

## Common Handlers

- [OnTextChanged](../handlers/handler_OnTextChanged.md)
- [OnKeyEnter](../handlers/handler_OnKeyEnter.md)
- [OnKeyEscape](../handlers/handler_OnKeyEscape.md)
- [OnMouseOver](../handlers/handler_OnMouseOver.md)
- [OnKeyTab](../handlers/handler_OnKeyTab.md)
- [OnShown](../handlers/handler_OnShown.md)

## Common Handler Functions

- Enemy.ConfigurationWindow_OnChange
- Enemy.ConfigurationWindow_ShowTooltip
- Enemy.UnitFramesUI_EffectsIndicatorDialog_UpdateExample
- Killer.OnSettingsEditChanged
- MiracleGrow2.LayoutBarCChanged
- MiracleGrow2.LayoutProgDimChanged
- BuffHead.Setup.SelectColor.OnTintChanged
- MapMonster.PinTypeEditor.MouseOverDescription
- MapPin.TimeChanged
- MiracleGrow2.ConfigThrobCChanged
- ShiniesAutoUI.OnPriceChange
- ShiniesPostUI.OnPriceChange


## XML Event Bindings

| Event | Common Lua Bindings | Expected Callback | Args Confidence |
|-------|---------------------|-------------------|-----------------|
| [OnKeyEnter](../handlers/handler_OnKeyEnter.md) | MapPin.SendCommand, MapPin.SendText, ObjectInspector.InspectObject, DebugWindow.TextSend, DevPadWindow.ConfirmRename, DevPadWindow.CreateNewFile | `function(...)` | LOW |
| [OnKeyEscape](../handlers/handler_OnKeyEscape.md) | DebugWindow.TextClear, DevPadWindow.OnKeyEscape | `function(...)` | LOW |
| [OnKeyTab](../handlers/handler_OnKeyTab.md) | DebugWindow.OnKeyTab, DevPadWindow.OnKeyTab | `function(...)` | LOW |
| [OnMouseOver](../handlers/handler_OnMouseOver.md) | Enemy.ConfigurationWindow_ShowTooltip, MapMonster.PinTypeEditor.MouseOverDescription, Enemy.GroupsUI_EffectFilterDialog_OnAbilityIdsMouseOver, Enemy.UnitFramesUI_UnitFrameClickCastingDialog_OnActionConfig2CommandMouseOver, LPET.OnMouseOver, WbLeadHelperMessage.OnMouseOverLabelEditBox | `function()` | MEDIUM |
| [OnShown](../handlers/handler_OnShown.md) | DebugWindow.OnShowFocus, DevPadWindow.OnShown | `function()` | MEDIUM |
| [OnTextChanged](../handlers/handler_OnTextChanged.md) | Enemy.ConfigurationWindow_OnChange, Enemy.UnitFramesUI_EffectsIndicatorDialog_UpdateExample, Killer.OnSettingsEditChanged, MiracleGrow2.LayoutBarCChanged, MiracleGrow2.LayoutProgDimChanged, BuffHead.Setup.SelectColor.OnTintChanged | `function()` | MEDIUM |

## Common Inherits

- EA_EditBox_DefaultFrame
- Aura_EditBox_DefaultFrame
- EA_EditBox_DefaultFrame_Multiline
- IraConfigNumEdit
- RVAPI_ColorDialogEditBoxTemplate
- IraConfigNumSpin
- Shinies_GoldCoin_EditBox_DefaultFrame
- MapMonsterEditorWindowEditBoxCoord
- MapMonsterPinTypeEditorWindowEditBoxDefault
- Shinies_BrassCoin_EditBox_DefaultFrame
- Shinies_SilverCoin_EditBox_DefaultFrame
- EA_EditBox_NoFrame

## Common Parent Elements

- [Window](element_Window.md)

## Common Named Child Elements

- [VerticalScrollbar](element_VerticalScrollbar.md)
- [Label](element_Label.md)

## Common Structural Child Elements

- [TextOffset](element_TextOffset.md)

## Typical XML Structure

```xml
<EditBox name="..." font="font_chat_text" background="EA_FullResizeImage_TanBorder">
  <TextOffset x="4" y="2"/>
</EditBox>
```

## Attribute Reference

| Attribute | Required | Usage % | Sample Values |
| --- | --- | --- | --- |
| `inherits` | **required** | 94% | IraConfigNumSpin, EA_EditBox_DefaultFrame, IraConfigNumEdit, EA_EditBox_DefaultFrame_Multiline, ... |
| `maxchars` | optional | 62% | 4096, 10, 3, 5, ... |
| `input` | optional | 33% | numbers, nospaces |
| `layer` | optional | 25% | default, secondary |
| `warnOnTextCropped` | optional | 17% | false |
| `font` | optional | 16% | font_clear_small_bold, font_chat_text, font_clear_small, font_clear_medium, ... |
| `taborder` | optional | 15% | 3, 2, 1, 6, ... |
| `handleinput` | optional | 7% | true |
| `scrolling` | optional | 7% | none, vert |
| `maxChars` | optional | 6% | 3, 2, 5, 64, ... |
| `background` | optional | 5% | EA_FullResizeImage_TanBorder |
| `id` | optional | 5% | 2, 1, 4, 3, ... |
| `align` | optional | 3% | rightcenter, center |
| `scrollbar` | optional | 3% | $parentDevPadCodeScrollBar, EA_ScrollBar_DefaultVerticalChain, CopyScrollBar, $parentObjectScrollbar |
| `alpha` | optional | 2% | 0.97 |
| `history` | optional | 2% | 30, 10 |
| `textalign` | optional | 1% | leftcenter |
| `autoHideScrollBar` | optional | 0% | true |
| `maxchar` | optional | 0% | 79000 |
| `wordwrap` | optional | 0% | false |
| `linecount` | optional | 0% | 9 |
## Structural Sub-Elements

### [TextOffset](element_TextOffset.md)

Observed 20 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `x` | **required** |  |
| `y` | **required** |  |
## Lua API Usage (from Handlers)

API functions commonly called from event handler Lua functions on this element type:

| API Function | Call Count | From Events |
| --- | --- | --- |
| `TextEditBoxGetText` | 182 | OnKeyEnter, OnKeyEscape, OnTextChanged |
| `WindowSetAlpha` | 36 | OnTextChanged |
| `WindowGetId` | 16 | OnMouseOver, OnTextChanged |
| `WindowSetTintColor` | 9 | OnTextChanged |
| `WindowSetShowing` | 7 | OnKeyEnter, OnKeyEscape |
| `LabelSetText` | 5 | OnTextChanged |
| `LabelSetTextColor` | 3 | OnTextChanged |
| `TextEditBoxSetText` | 3 | OnKeyEnter, OnTextChanged |
| `WindowSetDimensions` | 3 | OnTextChanged |
| `WindowGetShowing` | 2 | OnKeyEscape, OnShown |
| `LabelGetText` | 1 | OnKeyEnter |
| `RegisterEventHandler` | 1 | OnTextChanged |
## Handler Callback Signatures

Expected callback argument patterns for event handlers on this element type:

### OnTextChanged

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `text` | wstring | current_text |
### OnTextChanged

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `text` | wstring | current_text |
### OnTextChanged

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `text` | wstring | current_text |
### OnTextChanged

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `text` | wstring | current_text |
### OnTextChanged

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `text` | wstring | current_text |
### OnTextChanged

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `text` | wstring | current_text |
### OnTextChanged

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `text` | wstring | current_text |
### OnTextChanged

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `text` | wstring | current_text |
### OnTextChanged

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `text` | wstring | current_text |
### OnTextChanged

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `text` | wstring | current_text |
### OnTextChanged

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `text` | wstring | current_text |
### OnTextChanged

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `text` | wstring | current_text |
### OnTextChanged

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `text` | wstring | current_text |
### OnTextChanged

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `text` | wstring | current_text |
### OnTextChanged

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `text` | wstring | current_text |
### OnTextChanged

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `text` | wstring | current_text |
### OnTextChanged

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `text` | wstring | current_text |
### OnTextChanged

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `text` | wstring | current_text |
### OnTextChanged

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `text` | wstring | current_text |
### OnTextChanged

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `text` | wstring | current_text |
### OnTextChanged

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `text` | wstring | current_text |
### OnTextChanged

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `text` | wstring | current_text |
### OnTextChanged

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `text` | wstring | current_text |
### OnTextChanged

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `text` | wstring | current_text |
### OnTextChanged

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `text` | wstring | current_text |
### OnTextChanged

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `text` | wstring | current_text |
### OnTextChanged

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `text` | wstring | current_text |
### OnTextChanged

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `text` | wstring | current_text |
### OnTextChanged

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `text` | wstring | current_text |
### OnTextChanged

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `text` | wstring | current_text |
### OnTextChanged

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `text` | wstring | current_text |
### OnTextChanged

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `text` | wstring | current_text |
### OnTextChanged

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `text` | wstring | current_text |
### OnTextChanged

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `text` | wstring | current_text |
### OnTextChanged

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `text` | wstring | current_text |
### OnTextChanged

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `text` | wstring | current_text |
### OnTextChanged

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `text` | wstring | current_text |
### OnTextChanged

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `text` | wstring | current_text |
### OnTextChanged

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `text` | wstring | current_text |
### OnTextChanged

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `text` | wstring | current_text |
### OnTextChanged

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `text` | wstring | current_text |
### OnTextChanged

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `text` | wstring | current_text |
### OnTextChanged

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `text` | wstring | current_text |
### OnTextChanged

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `text` | wstring | current_text |
### OnTextChanged

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `text` | wstring | current_text |
### OnTextChanged

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `text` | wstring | current_text |
### OnTextChanged

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `text` | wstring | current_text |
### OnTextChanged

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `text` | wstring | current_text |
### OnTextChanged

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `text` | wstring | current_text |
### OnTextChanged

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `text` | wstring | current_text |
### OnTextChanged

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `text` | wstring | current_text |
### OnTextChanged

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `text` | wstring | current_text |
### OnTextChanged

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `text` | wstring | current_text |
### OnTextChanged

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `text` | wstring | current_text |
### OnTextChanged

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `text` | wstring | current_text |
### OnTextChanged

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `text` | wstring | current_text |
### OnTextChanged

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `text` | wstring | current_text |
### OnTextChanged

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `text` | wstring | current_text |
### OnTextChanged

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `text` | wstring | current_text |
### OnTextChanged

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `text` | wstring | current_text |
### OnTextChanged

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `text` | wstring | current_text |
### OnTextChanged

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `text` | wstring | current_text |
### OnTextChanged

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `text` | wstring | current_text |
### OnTextChanged

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `text` | wstring | current_text |
### OnTextChanged

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `text` | wstring | current_text |
### OnTextChanged

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `text` | wstring | current_text |
### OnTextChanged

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `text` | wstring | current_text |
### OnTextChanged

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `text` | wstring | current_text |
### OnTextChanged

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `text` | wstring | current_text |
### OnTextChanged

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `text` | wstring | current_text |
### OnTextChanged

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `text` | wstring | current_text |
### OnTextChanged

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `text` | wstring | current_text |
### OnTextChanged

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `text` | wstring | current_text |
### OnTextChanged

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `text` | wstring | current_text |
### OnTextChanged

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `text` | wstring | current_text |
### OnTextChanged

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `text` | wstring | current_text |
### OnTextChanged

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `text` | wstring | current_text |
### OnTextChanged

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `text` | wstring | current_text |
### OnTextChanged

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `text` | wstring | current_text |
### OnTextChanged

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `text` | wstring | current_text |
### OnTextChanged

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `text` | wstring | current_text |
### OnTextChanged

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `text` | wstring | current_text |
### OnTextChanged

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `text` | wstring | current_text |
### OnTextChanged

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `text` | wstring | current_text |
### OnTextChanged

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `text` | wstring | current_text |
### OnTextChanged

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `text` | wstring | current_text |
### OnTextChanged

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `text` | wstring | current_text |
### OnTextChanged

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `text` | wstring | current_text |
### OnTextChanged

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `text` | wstring | current_text |
### OnTextChanged

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `text` | wstring | current_text |
### OnTextChanged

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `text` | wstring | current_text |
### OnTextChanged

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `text` | wstring | current_text |
### OnTextChanged

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `text` | wstring | current_text |
### OnTextChanged

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `text` | wstring | current_text |
### OnTextChanged

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `text` | wstring | current_text |
### OnTextChanged

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `text` | wstring | current_text |
### OnTextChanged

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `text` | wstring | current_text |
### OnTextChanged

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `text` | wstring | current_text |
### OnTextChanged

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `text` | wstring | current_text |
### OnTextChanged

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `text` | wstring | current_text |
### OnTextChanged

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `text` | wstring | current_text |
### OnTextChanged

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `text` | wstring | current_text |
### OnTextChanged

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `text` | wstring | current_text |
### OnTextChanged

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `text` | wstring | current_text |
### OnTextChanged

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `text` | wstring | current_text |
## Lua Functions Manipulating This Type

- LibSlash.LibSlash.Initialize
- Enemy.Enemy.GroupsUI_EffectFilterDialog_OnDurationTypeSelChanged

## Seen In

- AdvancedRenownTrainer
- Aura
- BuffHead
- Busted
- EA_UiDebugTools
- EA_UiModWindow
- Enemy
- Killer
- LoyalPet
- MapMonster
- MapPin
- Miracle Grow Remix
- ObjectInspector
- Pocket Palette
- RVAPI_ColorDialog
- RVMOD_Manager
- RandomMount
- Shinies
- TexturedButtons
- TidyChat
- TidyRoll
- TurretRange
- bigger_MacroWindow
- wbLeadHelper

## Examples

- AdvancedRenownTrainer: AdvancedRenownTrainingImportNameInputWindowNameInputBox -> EditBox AdvancedRenownTrainingImportNameInputWindowNameInputBox
- AdvancedRenownTrainer: AdvancedRenownTrainingImportWindowLinkInputBox -> EditBox AdvancedRenownTrainingImportWindowLinkInputBox
- AdvancedRenownTrainer: AdvancedRenownTrainingImportWindowNameInputBox -> EditBox AdvancedRenownTrainingImportWindowNameInputBox
- AdvancedRenownTrainer: AdvancedRenownTrainingLinkWindowLinkBox -> EditBox AdvancedRenownTrainingLinkWindowLinkBox
- AdvancedRenownTrainer: AdvancedRenownTrainingPresetsWindowSaveNameInput -> EditBox AdvancedRenownTrainingPresetsWindowSaveNameInput
- Aura: AuraConfigActivationAlertTextText -> EditBox AuraConfigActivationAlertTextText

## Related APIs

- [DebugWindow.OnShowFocus](../../globals/functions/global_DebugWindow.OnShowFocus.md) (HIGH 100/100) - Global Function
- [TextEditBoxSetTextColor](../../window_api/functions/window_TextEditBoxSetTextColor.md) (HIGH 100/100) - Window Function
- [TextEditBoxInsertText](../../window_api/functions/window_TextEditBoxInsertText.md) (HIGH 80/100) - Window Function

## Used With

- [OnKeyEnter](../../events/window_events/window_event_OnKeyEnter.md) (HIGH 100/100) - Window Event
- [OnKeyEnter](../handlers/handler_OnKeyEnter.md) (HIGH 100/100) - XML Handler
- [OnKeyEscape](../../events/window_events/window_event_OnKeyEscape.md) (HIGH 100/100) - Window Event
- [OnKeyEscape](../handlers/handler_OnKeyEscape.md) (HIGH 100/100) - XML Handler
- [OnMouseOver](../../events/window_events/window_event_OnMouseOver.md) (HIGH 100/100) - Window Event
- [OnMouseOver](../handlers/handler_OnMouseOver.md) (HIGH 100/100) - XML Handler
- [OnShown](../handlers/handler_OnShown.md) (HIGH 100/100) - XML Handler
- [OnTextChanged](../../events/window_events/window_event_OnTextChanged.md) (HIGH 100/100) - Window Event
- [OnTextChanged](../handlers/handler_OnTextChanged.md) (HIGH 100/100) - XML Handler
- [OnKeyTab](../handlers/handler_OnKeyTab.md) (HIGH 93/100) - XML Handler
- [OnKeyTab](../../events/window_events/window_event_OnKeyTab.md) (HIGH 73/100) - Window Event

## Triggered By

- none

## Affects

- none
