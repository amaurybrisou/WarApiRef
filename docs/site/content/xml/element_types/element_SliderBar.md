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
| Addons seen in | Aura, BuffHead, DAoCBuff, Enemy, LibGroup, MoraleCircle, PotionBar, RoR_SoR |
| Files seen in | `/workspace/data/raw/Aura/Source/AuraColorPicker.xml:0`, `/workspace/data/raw/Aura/Source/AuraConfig.xml:0`, `/workspace/data/raw/Aura/Source/Templates.xml:0`, `/workspace/data/raw/BuffHead/Setup/General.xml:0`, `/workspace/data/raw/BuffHead/Setup/SetupAdvancedContainersItemProperties.xml:0`, `/workspace/data/raw/BuffHead/Setup/SetupContainer.xml:0`, `/workspace/data/raw/BuffHead/Setup/SetupDisplay.xml:0`, `/workspace/data/raw/BuffHead/Setup/SetupGeneral.xml:0` |
| Namespaces detected | SliderBar |
| Source kinds | xml_frames, xml_handlers |
| Example locations | Aura: AuraColorPickerAlpha, Aura: AuraColorPickerBlue, Aura: AuraColorPickerGreen, Aura: AuraColorPickerRed, Aura: AuraConfigGeneralTextureRotationSlider, Aura: AuraConfigGeneralTextureScaleSlider |
| XML usage count | 85 |
| XML attribute usage count | 85 |
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

Observed XML element type instantiated by 14 addons.

## Common Attributes

- name
- inherits
- numticks
- locktoticks
- scale
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
- MoraleCircle.OnSetCustomColor
- MoraleCircle.OnSetCustomColorEmpty
- MoraleCircle.OnSetCustomColorFill
- MoraleCircle.OnSetCustomColorFull
- TexturedButtons.Setup.SelectColor.OnSlideTint
- TexturedButtons.Setup.Tint.OnSlideTint
- TurretRange.Setup.Display.OnSlideTint
- TurretRange.Setup.Distance.OnSlideTint
- WSCT.OnSetCustomColor


## XML Event Bindings

| Event | Category | Common Lua Bindings | Expected Callback | Args Confidence |
|-------|----------|---------------------|-------------------|-----------------|
| [OnMouseOver](../handlers/handler_OnMouseOver.md) | input | WSCT.OnMouseOver | `function()` | MEDIUM |
| [OnSlide](../handlers/handler_OnSlide.md) | custom | AuraColorPicker.OnSlide, Enemy.UnitFramesUI_EffectsIndicatorDialog_UpdateExample, BuffHead.Setup.SelectColor.OnSlideTint, MoraleCircle.OnSetCustomColor, MoraleCircle.OnSetCustomColorEmpty, MoraleCircle.OnSetCustomColorFill | `function(...)` | LOW |

### Per-Event Lua API Calls

**OnMouseOver** handlers call: `WindowGetParent`

**OnSlide** handlers call: `ButtonGetPressedFlag`, `ComboBoxGetSelectedMenuItem`, `LabelSetText`, `LabelSetTextColor`, `SliderBarGetCurrentPosition`, `SliderBarSetCurrentPosition`, `TextEditBoxGetText`, `TextEditBoxSetText`, `WindowGetParent`, `WindowSetAlpha`, `WindowSetScale`, `WindowSetTintColor`

## Common Inherits

- EA_Default_SliderBar
- Aura_Default_SliderBar

## Common Parent Elements

- [Windows](element_Windows.md) — 83× (HIGH)
- [Window](element_Window.md) — 2× (LOW)

## Common Structural Child Elements

- [Size](element_Size.md) — 76× (HIGH)

## Common Template Bases

- Aura_Default_SliderBar
- EA_Default_SliderBar


> **Note**: This element type commonly acts as a template base.

## Typical XML Structure

```xml
<SliderBar inherits="Aura_Default_SliderBar" locktoticks="false" name="..." numticks="360">
  <Size/>
</SliderBar>
```

## Attribute Reference

| Attribute | Required | Usage % | Sample Values |
| --- | --- | --- | --- |
| `inherits` | **required** | 98% | EA_Default_SliderBar, Aura_Default_SliderBar |
| `numticks` | optional | 14% | 5, 360, 250, 200, ... |
| `locktoticks` | optional | 10% | false |
| `scale` | optional | 4% | 0.4 |
| `background` | optional | 1% | EA_BrownHorizontalRule |
| `sliderbutton` | optional | 1% | Aura_Default_SliderButton |
| `tickmark` | optional | 1% | Aura_Default_SliderTick |
## Structural Sub-Elements

### [Size](element_Size.md)

Observed 76 times as an unnamed child.

## Lua API Usage (from Handlers)

API functions commonly called from event handler Lua functions on this element type:

| API Function | Category | Call Count | From Events |
| --- | --- | --- | --- |
| `SliderBarGetCurrentPosition` | ui | 173 | OnSlide |
| `LabelSetText` | ui | 61 | OnSlide |
| `TextEditBoxSetText` | ui | 46 | OnSlide |
| `TextEditBoxGetText` | ui | 36 | OnSlide |
| `ComboBoxGetSelectedMenuItem` | ui | 28 | OnSlide |
| `WindowSetTintColor` | ui | 15 | OnSlide |
| `ButtonGetPressedFlag` | ui | 12 | OnSlide |
| `LabelSetTextColor` | ui | 6 | OnSlide |
| `SliderBarSetCurrentPosition` | ui | 6 | OnSlide |
| `WindowGetParent` | ui | 4 | OnMouseOver, OnSlide |
| `WindowSetAlpha` | ui | 1 | OnSlide |
| `WindowSetScale` | ui | 1 | OnSlide |
## Handler Callback Signatures

Expected callback argument patterns for event handlers on this element type:

### OnMouseOver

Confidence: HIGH

### OnSlide

Confidence: LOW

## Lua Functions Manipulating This Type

- MoraleCircle.OnSetCustomColorFill
- RoR_SoR.OnWindowOptionsSetOffset
- MoraleCircle.OnSetCustomColorEmpty
- WSCT.OnLButtonUpColorPicker
- WSCT.OnSetCustomColor
- MoraleCircle.OnSetCustomColor
- MoraleCircle.ColorChanger1
- MoraleCircle.ColorChanger3
- RoR_SoR.OnWindowOptionsSetOpacity
- RoR_SoR.OnWindowOptionsSetScale
- WSCT.ColorOnButtonUp
- Enemy.UnitFramesUI_EffectsIndicatorDialog_Open
- Enemy.UnitFramesUI_EffectsIndicatorDialog_UpdateExample
- MoraleCircle.ColorChanger2
- MoraleCircle.OnSetCustomColorFull
- MoraleCircle.ColorChanger4


## Binding Resolution

- Total handler declarations: 81
- Resolved to Lua functions: 81 (100%)

## Seen In

- Aura
- BuffHead
- DAoCBuff
- Enemy
- LibGroup
- MoraleCircle
- PotionBar
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

- [Size](element_Size.md) (HIGH 100/100) - XML Element Type
- [Window](element_Window.md) (HIGH 100/100) - XML Element Type
- [Windows](element_Windows.md) (HIGH 100/100) - XML Element Type

## Used With

- [OnSlide](../handlers/handler_OnSlide.md) (HIGH 100/100) - XML Event

## Triggered By

- [OnMouseOver](../handlers/handler_OnMouseOver.md) (HIGH 100/100) - XML Event
- [OnSlide](../handlers/handler_OnSlide.md) (HIGH 100/100) - XML Event

## Affects

- none
