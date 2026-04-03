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
| Addons seen in | AdvancedRenownTrainer, Aura, BuffHead, DAoCBuff, Enemy, Killer, Pocket Palette, Shinies |
| Files seen in | `/workspace/data/raw/Aura/Source/AuraConfig.xml:0`, `/workspace/data/raw/Aura/Source/AuraShares.xml:0`, `/workspace/data/raw/Aura/Source/Templates.xml:0`, `/workspace/data/raw/BuffHead/Setup/General.xml:0`, `/workspace/data/raw/BuffHead/Setup/SetupAdvancedCompressionItem.xml:0`, `/workspace/data/raw/BuffHead/Setup/SetupAdvancedCompressionItemEffect.xml:0`, `/workspace/data/raw/BuffHead/Setup/SetupAdvancedContainersItemProperties.xml:0`, `/workspace/data/raw/BuffHead/Setup/SetupDisplay.xml:0` |
| Namespaces detected | EditBox |
| Source kinds | xml_frames, xml_handlers |
| Example locations | AdvancedRenownTrainer: AdvancedRenownTrainingImportNameInputWindowNameInputBox, AdvancedRenownTrainer: AdvancedRenownTrainingImportWindowLinkInputBox, AdvancedRenownTrainer: AdvancedRenownTrainingImportWindowNameInputBox, AdvancedRenownTrainer: AdvancedRenownTrainingLinkWindowLinkBox, AdvancedRenownTrainer: AdvancedRenownTrainingPresetsWindowSaveNameInput, Aura: AuraConfigActivationAlertTextText |
| XML usage count | 151 |
| XML attribute usage count | 151 |
| Lua usage count | 4 |
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

Observed XML element type instantiated by 13 addons.

## Common Attributes

- background
- font
- handleinput
- inherits
- input
- layer
- maxChars
- maxchars
- name
- scrolling
- taborder
- warnOnTextCropped

## Common Handlers

- [OnKeyEnter](../handlers/handler_OnKeyEnter.md)
- [OnKeyEscape](../handlers/handler_OnKeyEscape.md)
- [OnMouseOver](../handlers/handler_OnMouseOver.md)
- [OnTextChanged](../handlers/handler_OnTextChanged.md)

## Common Handler Functions

- Enemy.ConfigurationWindow_OnChange
- Enemy.ConfigurationWindow_ShowTooltip
- Enemy.UnitFramesUI_EffectsIndicatorDialog_UpdateExample
- Killer.OnSettingsEditChanged
- BuffHead.Setup.SelectColor.OnTintChanged
- ShiniesAutoUI.OnPriceChange
- ShiniesPostUI.OnPriceChange
- TexturedButtons.Setup.SelectColor.OnTintChanged
- TexturedButtons.Setup.Tint.OnTintChanged
- TurretRange.Setup.Display.OnTintChanged
- TurretRange.Setup.Distance.OnTintChanged
- AuraConfig.OnTextureOffsetXChanged


## XML Event Bindings

| Event | Category | Common Lua Bindings | Expected Callback | Args Confidence |
|-------|----------|---------------------|-------------------|-----------------|
| [OnMouseOver](../handlers/handler_OnMouseOver.md) | input | Enemy.ConfigurationWindow_ShowTooltip, Enemy.GroupsUI_EffectFilterDialog_OnAbilityIdsMouseOver, Enemy.UnitFramesUI_UnitFrameClickCastingDialog_OnActionConfig2CommandMouseOver | `function()` | MEDIUM |
| [OnTextChanged](../handlers/handler_OnTextChanged.md) | data | Enemy.ConfigurationWindow_OnChange, Enemy.UnitFramesUI_EffectsIndicatorDialog_UpdateExample, Killer.OnSettingsEditChanged, BuffHead.Setup.SelectColor.OnTintChanged, ShiniesAutoUI.OnPriceChange, ShiniesPostUI.OnPriceChange | `function()` | MEDIUM |
| [OnTextChanged](../handlers/handler_OnTextChanged.md) | data | Enemy.ConfigurationWindow_OnChange, Enemy.UnitFramesUI_EffectsIndicatorDialog_UpdateExample, Killer.OnSettingsEditChanged, BuffHead.Setup.SelectColor.OnTintChanged, ShiniesAutoUI.OnPriceChange, ShiniesPostUI.OnPriceChange, TexturedButtons.Setup.SelectColor.OnTintChanged, TexturedButtons.Setup.Tint.OnTintChanged, TurretRange.Setup.Display.OnTintChanged, TurretRange.Setup.Distance.OnTintChanged, AuraConfig.OnTextureOffsetXChanged, AuraConfig.OnTextureOffsetYChanged, AuraConfig.OnTimerOffsetXChanged, AuraConfig.OnTimerOffsetYChanged, BuffHead.Setup.AdvancedContainersItem.Properties.OnContainerPaddingXChanged, BuffHead.Setup.AdvancedContainersItem.Properties.OnContainerPaddingYChanged, BuffHead.Setup.AdvancedContainersItem.Properties.OnContainerSizeColumnsChanged, BuffHead.Setup.AdvancedContainersItem.Properties.OnContainerSizeRowsChanged, BuffHead.Setup.AdvancedContainersItem.Properties.OnIndicatorPaddingXChanged, BuffHead.Setup.AdvancedContainersItem.Properties.OnIndicatorPaddingYChanged, BuffHead.Setup.AdvancedContainersItem.Properties.OnMaximumThresholdChanged, BuffHead.Setup.AdvancedContainersItem.Properties.OnMinimumThresholdChanged, BuffHead.Setup.AdvancedContainersItem.Properties.OnOffsetXChanged, BuffHead.Setup.AdvancedContainersItem.Properties.OnOffsetYChanged, BuffHead.Setup.AdvancedContainersItem.Properties.OnPlacementFixedOffsetXChanged, BuffHead.Setup.AdvancedContainersItem.Properties.OnPlacementFixedOffsetYChanged, BuffHead.Setup.Display.OnOffsetXChanged, BuffHead.Setup.Display.OnOffsetYChanged, BuffHead.Setup.EffectCache.OnSearchChanged, BuffHead.Setup.Layout.Properties.OnCoreSizeSizeHeightChanged, BuffHead.Setup.Layout.Properties.OnCoreSizeSizeWidthChanged, BuffHead.Setup.Layout.Properties.OnOffsetXChanged, BuffHead.Setup.Layout.Properties.OnOffsetYChanged, BuffHead.Setup.Layout.Properties.OnSizeScaleChanged, BuffHead.Setup.Layout.Properties.OnSizeSizeHeightChanged, BuffHead.Setup.Layout.Properties.OnSizeSizeWidthChanged, BuffHead.Setup.PriorityEffectsItem.OnAbilityIdChanged, PP.UpdateDyeFilter, ShiniesConfigGeneral.OnTextChanged_UIScale, ShiniesPostUI.OnStackChange, ShiniesPostUI.OnStackSizeChange, TexturedButtons.Setup.Actionbar.OnPaddingXChanged, TexturedButtons.Setup.Actionbar.OnPaddingYChanged, TexturedButtons.Setup.Actionbar.OnSpacingXChanged, TexturedButtons.Setup.Actionbar.OnSpacingYChanged, TurretRange.Setup.Display.OnDistanceOffsetXChanged, TurretRange.Setup.Display.OnDistanceOffsetYChanged, TurretRange.Setup.Display.OnGraphicLimitChanged | `text` | LOW |
| [OnKeyEnter](../handlers/handler_OnKeyEnter.md) | custom | - | `` |  |
| [OnKeyEscape](../handlers/handler_OnKeyEscape.md) | custom | - | `` |  |
| [OnMouseOver](../handlers/handler_OnMouseOver.md) | input | Enemy.ConfigurationWindow_ShowTooltip, Enemy.GroupsUI_EffectFilterDialog_OnAbilityIdsMouseOver, Enemy.UnitFramesUI_UnitFrameClickCastingDialog_OnActionConfig2CommandMouseOver | `` |  |

### Per-Event Lua API Calls

**OnTextChanged** handlers call: `ButtonGetPressedFlag`, `ComboBoxGetSelectedMenuItem`, `RegisterEventHandler`, `SliderBarGetCurrentPosition`, `SliderBarSetCurrentPosition`, `TextEditBoxGetText`, `WindowSetTintColor`

**OnTextChanged** handlers call: `ButtonGetPressedFlag`, `ComboBoxGetSelectedMenuItem`, `RegisterEventHandler`, `SliderBarGetCurrentPosition`, `SliderBarSetCurrentPosition`, `TextEditBoxGetText`, `WindowSetTintColor`

## Common Inherits

- Aura_EditBox_DefaultFrame
- EA_EditBox_DefaultFrame
- EA_EditBox_DefaultFrame_Multiline
- Shinies_BrassCoin_EditBox_DefaultFrame
- Shinies_GoldCoin_EditBox_DefaultFrame
- Shinies_SilverCoin_EditBox_DefaultFrame

## Common Parent Elements

- [Windows](element_Windows.md) — 151× (HIGH)

## Common Structural Child Elements

- [Size](element_Size.md) — 127× (HIGH)
- [Anchors](element_Anchors.md) — 126× (HIGH)
- [EventHandlers](element_EventHandlers.md) — 102× (HIGH)
- [TextOffset](element_TextOffset.md) — 15× (HIGH)

## Common Template Bases

- Aura_EditBox_DefaultFrame
- EA_EditBox_DefaultFrame
- EA_EditBox_DefaultFrame_Multiline
- Shinies_BrassCoin_EditBox_DefaultFrame
- Shinies_GoldCoin_EditBox_DefaultFrame
- Shinies_SilverCoin_EditBox_DefaultFrame


> **Note**: This element type commonly acts as a template base.

## Typical XML Structure

```xml
<EditBox background="EA_FullResizeImage_TanBorder" font="font_clear_medium" name="..." taborder="1">
  <TextOffset x="4" y="2"/>
  <Size/>
</EditBox>
```

## Attribute Reference

| Attribute | Required | Usage % | Sample Values |
| --- | --- | --- | --- |
| `inherits` | **required** | 92% | EA_EditBox_DefaultFrame, Aura_EditBox_DefaultFrame, EA_EditBox_DefaultFrame_Multiline, Shinies_GoldCoin_EditBox_DefaultFrame, ... |
| `maxchars` | optional | 72% | 20, 76, 2048, 5, ... |
| `input` | optional | 41% | nospaces, numbers |
| `layer` | optional | 35% | secondary, default |
| `warnOnTextCropped` | optional | 24% | false |
| `font` | optional | 15% | font_chat_text, font_clear_small, font_clear_small_bold, font_clear_medium, ... |
| `taborder` | optional | 13% | 1, 2, 3 |
| `handleinput` | optional | 7% | true |
| `background` | optional | 7% | EA_FullResizeImage_TanBorder |
| `scrolling` | optional | 7% | vert, none |
| `maxChars` | optional | 5% | 5, 2, 3, 64 |
| `align` | optional | 4% | rightcenter, center |
| `history` | optional | 3% | 0 |
| `scrollbar` | optional | 1% | EA_ScrollBar_DefaultVerticalChain |
| `linecount` | optional | 0% | 9 |
## Structural Sub-Elements

### [Size](element_Size.md)

Observed 127 times as an unnamed child.

### [Anchors](element_Anchors.md)

Observed 126 times as an unnamed child.

### [EventHandlers](element_EventHandlers.md)

Observed 102 times as an unnamed child.

### [TextOffset](element_TextOffset.md)

Observed 15 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `x` | **required** | 5, 4, 10, 0 |
| `y` | **required** | 5, 2, 3, 0 |
## Recursive Hierarchy

- Root: [EditBox](element_EditBox.md)
- [Anchors](element_Anchors.md) (structural, 126×, HIGH)
  - [AbsPoint](element_AbsPoint.md) (structural, 3×, MEDIUM)
  - [Anchor](element_Anchor.md) (structural, 3889×, HIGH)
    - [AbsPoint](element_AbsPoint.md) (structural, 3500×, HIGH)
    - [Anchor](element_Anchor.md) (structural, 22×, HIGH)
      - (cycle)
- [EventHandlers](element_EventHandlers.md) (structural, 102×, HIGH)
  - [EventHandler](element_EventHandler.md) (structural, 1707×, HIGH)
- [Size](element_Size.md) (structural, 127×, HIGH)
  - [AbsPoint](element_AbsPoint.md) (structural, 2634×, HIGH)
- [TextOffset](element_TextOffset.md) (structural, 15×, HIGH)

## Lua API Usage (from Handlers)

API functions commonly called from event handler Lua functions on this element type:

| API Function | Category | Call Count | From Events |
| --- | --- | --- | --- |
| `TextEditBoxGetText` | ui | 138 | OnTextChanged |
| `ComboBoxGetSelectedMenuItem` | ui | 51 | OnTextChanged |
| `SliderBarSetCurrentPosition` | ui | 46 | OnTextChanged |
| `SliderBarGetCurrentPosition` | ui | 28 | OnTextChanged |
| `ButtonGetPressedFlag` | ui | 21 | OnTextChanged |
| `WindowSetTintColor` | ui | 6 | OnTextChanged |
| `RegisterEventHandler` | event | 1 | OnTextChanged |
## Handler Callback Signatures

Expected callback argument patterns for event handlers on this element type:

### OnKeyEnter

Confidence: LOW

### OnKeyEscape

Confidence: LOW

### OnMouseOver

Confidence: HIGH

### OnTextChanged

Confidence: MEDIUM

| Position | Name | Type | Role |
| --- | --- | --- | --- |
| 0 | `text` | wstring | current_text |
## Lua Functions Manipulating This Type

- PP.UpdateDyeFilter
- Enemy.GroupsUI_EffectFilterDialog_Open
- Enemy.GroupsUI_EffectFilterDialog_Ok
- Enemy.UnitFramesUI_UnitFramePartDialog_UpdateExample
- Enemy.UnitFramesUI_UnitFramePartDialog_Open
- Enemy.UnitFramesUI_UnitFrameClickCastingDialog_Open
- Enemy.UnitFramesUI_UnitFrameClickCastingDialog_Ok
- Enemy.GroupsUI_EffectFilterDialog_OnDurationTypeSelChanged
- Enemy.IntercomUI_ChooseChannelDialog_OnOkButton
- Enemy.UnitFramesUI_EffectsIndicatorDialog_Open
- Enemy.UnitFramesUI_EffectsIndicatorDialog_UpdateExample
- Enemy.UnitFramesUI_EffectsIndicatorDialog_OnLScaleCheckBoxChanged


## Binding Resolution

- Total handler declarations: 139
- Resolved to Lua functions: 95 (68%)

## Seen In

- AdvancedRenownTrainer
- Aura
- BuffHead
- DAoCBuff
- Enemy
- Killer
- Pocket Palette
- Shinies
- TexturedButtons
- TidyChat
- TidyRoll
- TurretRange
- bigger_MacroWindow

## Examples

- AdvancedRenownTrainer: AdvancedRenownTrainingImportNameInputWindowNameInputBox -> EditBox AdvancedRenownTrainingImportNameInputWindowNameInputBox
- AdvancedRenownTrainer: AdvancedRenownTrainingImportWindowLinkInputBox -> EditBox AdvancedRenownTrainingImportWindowLinkInputBox
- AdvancedRenownTrainer: AdvancedRenownTrainingImportWindowNameInputBox -> EditBox AdvancedRenownTrainingImportWindowNameInputBox
- AdvancedRenownTrainer: AdvancedRenownTrainingLinkWindowLinkBox -> EditBox AdvancedRenownTrainingLinkWindowLinkBox
- AdvancedRenownTrainer: AdvancedRenownTrainingPresetsWindowSaveNameInput -> EditBox AdvancedRenownTrainingPresetsWindowSaveNameInput
- Aura: AuraConfigActivationAlertTextText -> EditBox AuraConfigActivationAlertTextText

## Related APIs

- [EA_EditBox_DefaultFrame](../../globals/constants/constant_EA_EditBox_DefaultFrame.md) (HIGH 100/100) - Constant
- [EA_EditBox_DefaultFrame_Multiline](../../globals/constants/constant_EA_EditBox_DefaultFrame_Multiline.md) (HIGH 100/100) - Constant
- [Size](element_Size.md) (HIGH 100/100) - XML Element Type
- [TextOffset](element_TextOffset.md) (HIGH 100/100) - XML Element Type
- [Windows](element_Windows.md) (HIGH 100/100) - XML Element Type
- [Anchors](element_Anchors.md) (MEDIUM 55/100) - XML Element Type
- [EventHandlers](element_EventHandlers.md) (MEDIUM 55/100) - XML Element Type

## Used With

- [OnKeyEscape](../handlers/handler_OnKeyEscape.md) (HIGH 100/100) - XML Event
- [OnMouseOver](../handlers/handler_OnMouseOver.md) (HIGH 100/100) - XML Event
- [OnTextChanged](../handlers/handler_OnTextChanged.md) (HIGH 100/100) - XML Event
- [OnKeyEnter](../handlers/handler_OnKeyEnter.md) (HIGH 93/100) - XML Event

## Triggered By

- [OnKeyEscape](../handlers/handler_OnKeyEscape.md) (HIGH 100/100) - XML Event
- [OnMouseOver](../handlers/handler_OnMouseOver.md) (HIGH 100/100) - XML Event
- [OnTextChanged](../handlers/handler_OnTextChanged.md) (HIGH 100/100) - XML Event
- [OnKeyEnter](../handlers/handler_OnKeyEnter.md) (HIGH 93/100) - XML Event

## Affects

- none
