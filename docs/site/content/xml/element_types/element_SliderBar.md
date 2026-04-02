# SliderBar

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
| Addons seen in | Aura, BuffHead, Enemy, LibGroup, MapMonster, MapPin, MoraleCircle, PotionBar |
| Files seen in | `/workspace_addons/Aura/Source/AuraColorPicker.xml:26`, `/workspace_addons/Aura/Source/AuraColorPicker.xml:58`, `/workspace_addons/Aura/Source/AuraColorPicker.xml:78`, `/workspace_addons/Aura/Source/AuraColorPicker.xml:98`, `/workspace_addons/Aura/Source/AuraConfig.xml:1670`, `/workspace_addons/Aura/Source/AuraConfig.xml:246`, `/workspace_addons/Aura/Source/AuraConfig.xml:301`, `/workspace_addons/Aura/Source/Templates.xml:317` |
| Namespaces detected | SliderBar |
| Source kinds | xml_frames, xml_handlers |
| Example locations | Aura: AuraColorPickerAlpha, Aura: AuraColorPickerBlue, Aura: AuraColorPickerGreen, Aura: AuraColorPickerRed, Aura: AuraConfigGeneralTextureRotationSlider, Aura: AuraConfigGeneralTextureScaleSlider |
| XML usage count | 91 |
| XML attribute usage count | 91 |
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

Observed XML element type instantiated by 17 addons.

## Common Attributes

- name
- inherits
- numticks
- handleinput
- scale
- locktoticks
- autoresize
- handlinput
- numTicks
- background
- sliderbutton
- tickmark

## Common Handlers

- [OnSlide](../handlers/handler_OnSlide.md)
- [OnMouseOver](../handlers/handler_OnMouseOver.md)

## Common Handler Functions

- AuraColorPicker.OnSlide
- Enemy.UnitFramesUI_EffectsIndicatorDialog_UpdateExample
- BuffHead.Setup.SelectColor.OnSlideTint
- MapMonster.PinTypeEditor.OnSetCustomColor
- MoraleCircle.OnSetCustomColor
- MoraleCircle.OnSetCustomColorEmpty
- MoraleCircle.OnSetCustomColorFill
- MoraleCircle.OnSetCustomColorFull
- TexturedButtons.Setup.SelectColor.OnSlideTint
- TexturedButtons.Setup.Tint.OnSlideTint
- TurretRange.Setup.Display.OnSlideTint
- TurretRange.Setup.Distance.OnSlideTint


## XML Event Bindings

| Event | Category | Common Lua Bindings | Expected Callback | Args Confidence |
|-------|----------|---------------------|-------------------|-----------------|
| [OnMouseOver](../handlers/handler_OnMouseOver.md) | input | WSCT.OnMouseOver | `function()` | MEDIUM |
| [OnSlide](../handlers/handler_OnSlide.md) | custom | AuraColorPicker.OnSlide, Enemy.UnitFramesUI_EffectsIndicatorDialog_UpdateExample, BuffHead.Setup.SelectColor.OnSlideTint, MapMonster.PinTypeEditor.OnSetCustomColor, MoraleCircle.OnSetCustomColor, MoraleCircle.OnSetCustomColorEmpty | `function(...)` | LOW |

### Per-Event Lua API Calls

**OnMouseOver** handlers call: `WindowGetParent`

**OnSlide** handlers call: `ButtonGetPressedFlag`, `ComboBoxGetSelectedMenuItem`, `LabelSetText`, `LabelSetTextColor`, `SliderBarGetCurrentPosition`, `SliderBarSetCurrentPosition`, `TextEditBoxGetText`, `TextEditBoxSetText`, `WindowGetParent`, `WindowGetScale`, `WindowSetAlpha`, `WindowSetScale`, `WindowSetTintColor`

## Common Inherits

- EA_Default_SliderBar
- Aura_Default_SliderBar
- RVMOD_ManagerSliderTemplate
- RVAPI_ColorDialogSliderTemplate

## Common Parent Elements

- [Window](element_Window.md) — 41× (HIGH)

## Common Named Child Elements

- [Label](element_Label.md) — 2× (LOW)

## Common Template Bases

- Aura_Default_SliderBar
- EA_Default_SliderBar
- RVAPI_ColorDialogSliderTemplate
- RVMOD_ManagerSliderTemplate


> **Note**: This element type commonly acts as a template base.

## Attribute Reference

| Attribute | Required | Usage % | Sample Values |
| --- | --- | --- | --- |
| `inherits` | optional | 68% | EA_Default_SliderBar, Aura_Default_SliderBar, RVMOD_ManagerSliderTemplate, RVAPI_ColorDialogSliderTemplate |
| `numticks` | optional | 6% | 200, 250, 11, 5, ... |
| `handleinput` | optional | 5% | true |
| `scale` | optional | 3% | 0.4 |
| `locktoticks` | optional | 2% | false |
| `autoresize` | optional | 1% | true |
| `handlinput` | optional | 1% | true |
| `numTicks` | optional | 1% | 11 |
| `background` | optional | 0% | EA_BrownHorizontalRule |
| `sliderbutton` | optional | 0% | Aura_Default_SliderButton |
| `tickmark` | optional | 0% | Aura_Default_SliderTick |
## Lua API Usage (from Handlers)

API functions commonly called from event handler Lua functions on this element type:

| API Function | Category | Call Count | From Events |
| --- | --- | --- | --- |
| `SliderBarGetCurrentPosition` | ui | 150 | OnSlide |
| `LabelSetText` | ui | 66 | OnSlide |
| `TextEditBoxSetText` | ui | 46 | OnSlide |
| `TextEditBoxGetText` | ui | 36 | OnSlide |
| `ComboBoxGetSelectedMenuItem` | ui | 28 | OnSlide |
| `WindowSetTintColor` | ui | 18 | OnSlide |
| `ButtonGetPressedFlag` | ui | 12 | OnSlide |
| `SliderBarSetCurrentPosition` | ui | 9 | OnSlide |
| `LabelSetTextColor` | ui | 6 | OnSlide |
| `WindowSetAlpha` | ui | 6 | OnSlide |
| `WindowGetParent` | ui | 4 | OnMouseOver, OnSlide |
| `WindowSetScale` | ui | 2 | OnSlide |
| `WindowGetScale` | ui | 1 | OnSlide |
## Handler Callback Signatures

Expected callback argument patterns for event handlers on this element type:

### OnMouseOver

Confidence: HIGH

### OnSlide

Confidence: LOW

## Lua Functions Manipulating This Type

- MoraleCircle.MoraleCircle.OnSetCustomColorFill
- MoraleCircle.MoraleCircle.OnSetCustomColorFull
- Enemy.Enemy.UnitFramesUI_EffectsIndicatorDialog_UpdateExample
- WSCT.WSCT.OnSetCustomColor
- WSCT.WSCT.ColorOnButtonUp
- MoraleCircle.MoraleCircle.ColorChanger1
- MoraleCircle.MoraleCircle.OnSetCustomColor
- RoR_SoR.RoR_SoR.OnWindowOptionsSetOpacity
- MoraleCircle.MoraleCircle.ColorChanger2
- MoraleCircle.MoraleCircle.ColorChanger4
- RoR_SoR.RoR_SoR.OnWindowOptionsSetOffset
- RoR_SoR.RoR_SoR.OnWindowOptionsSetScale
- MoraleCircle.MoraleCircle.OnSetCustomColorEmpty
- Enemy.Enemy.UnitFramesUI_EffectsIndicatorDialog_Open
- MoraleCircle.MoraleCircle.ColorChanger3
- WSCT.WSCT.OnLButtonUpColorPicker


## Binding Resolution

- Total handler declarations: 84
- Resolved to Lua functions: 84 (100%)

## Seen In

- Aura
- BuffHead
- Enemy
- LibGroup
- MapMonster
- MapPin
- MoraleCircle
- PotionBar
- RVAPI_ColorDialog
- RVMOD_Manager
- RoR_SoR
- Shinies
- TexturedButtons
- TidyChat
- TurretRange
- WSCT
- WarBoard

## Examples

- Aura: AuraColorPickerAlpha -> SliderBar AuraColorPickerAlpha
- Aura: AuraColorPickerBlue -> SliderBar AuraColorPickerBlue
- Aura: AuraColorPickerGreen -> SliderBar AuraColorPickerGreen
- Aura: AuraColorPickerRed -> SliderBar AuraColorPickerRed
- Aura: AuraConfigGeneralTextureRotationSlider -> SliderBar AuraConfigGeneralTextureRotationSlider
- Aura: AuraConfigGeneralTextureScaleSlider -> SliderBar AuraConfigGeneralTextureScaleSlider

## Related APIs

- none

## Used With

- [OnSlide](../../events/window_events/window_event_OnSlide.md) (HIGH 100/100) - Window Event
- [OnSlide](../handlers/handler_OnSlide.md) (HIGH 100/100) - XML Handler

## Triggered By

- none

## Affects

- none
