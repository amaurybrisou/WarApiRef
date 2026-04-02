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


## Structural Sub-Elements

### [TextOffset](element_TextOffset.md)

- Observed in 20 parent frames
- Attributes: `x`, `y`
  - `x`: `4`, `0`, `8`, `7`
  - `y`: `2`, `3`, `4`

## Typical XML Structure

```xml
<EditBox name="..." font="font_chat_text" background="EA_FullResizeImage_TanBorder">
  <TextOffset x="4" y="2"/>
</EditBox>
```


## Attribute Reference

| Attribute | Role | Observed Values | Count |
|-----------|------|-----------------|-------|
| `name` | string | `AdvancedRenownTrainingImportNameInputWindowNameInputBox`, `AdvancedRenownTrainingImportWindowLinkInputBox`, `AdvancedRenownTrainingImportWindowNameInputBox`, `AdvancedRenownTrainingLinkWindowLinkBox`, … | 212 |
| `inherits` | frame-ref | `EA_EditBox_DefaultFrame`, `Aura_EditBox_DefaultFrame`, `EA_EditBox_DefaultFrame_Multiline`, `EA_EditBox_NoFrame`, … | 201 |
| `maxchars` | number | `76`, `20`, `5`, `3`, … | 132 |
| `input` | string | `nospaces`, `numbers` | 71 |
| `layer` | string | `secondary`, `default` | 54 |
| `warnOnTextCropped` | boolean | `false` | 37 |
| `font` | string | `font_chat_text`, `font_clear_small`, `font_clear_default`, `font_clear_small_bold`, … | 35 |
| `taborder` | number | `1`, `2`, `3`, `6`, … | 32 |
| `handleinput` | boolean | `true` | 16 |
| `scrolling` | string | `vert`, `none` | 16 |
| `maxChars` | number | `300`, `3`, `4`, `2`, … | 14 |
| `background` | string | `EA_FullResizeImage_TanBorder` | 11 |
| `id` | number | `1`, `5`, `3`, `4`, … | 11 |
| `align` | string | `rightcenter`, `center` | 7 |
| `scrollbar` | frame-ref | `EA_ScrollBar_DefaultVerticalChain`, `CopyScrollBar`, `$parentDevPadCodeScrollBar`, `$parentObjectScrollbar` | 7 |

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
