# HorizontalResizeImage

- Category: XML Element Type
- Confidence level: HIGH
- Confidence score: 100/100

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 135

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, seen in 4 or more addons, used directly in xml handler attributes.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +30 Used directly in XML handler attributes: XML exposure suggests an engine-level contract.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | AdvancedPetAssist, AggroMeter, Aura, BuffHead, PartyCast, PotionBar, RoR_SoR, Shinies |
| Files seen in | `/workspace/data/raw/AdvancedPetAssist/APAGui.xml:0`, `/workspace/data/raw/AggroMeter/AggroMeter.xml:0`, `/workspace/data/raw/Aura/Source/AuraConfig.xml:0`, `/workspace/data/raw/BuffHead/Setup/SetupLayout.xml:0`, `/workspace/data/raw/PartyCast/PartyCast.xml:0`, `/workspace/data/raw/PotionBar/settings/Settings.xml:14`, `/workspace/data/raw/RoR_SoR/RoR_SoR.xml:0`, `/workspace/data/raw/Shinies/Source/Shinies.xml:0` |
| Namespaces detected | HorizontalResizeImage |
| Source kinds | xml_frames |
| Example locations | AdvancedPetAssist: APAFollowTargetHUDFill, AdvancedPetAssist: APAInstantOnlyHUDFill, AdvancedPetAssist: APAKitingHUDFill, AdvancedPetAssist: APAPetTargetHUDBg, AggroMeter: AggroMeterGrayWindowSeparatorMiddle, Aura: AuraConfigTabsSeparatorActivation |
| XML usage count | 19 |
| XML attribute usage count | 19 |
| Lua usage count | 0 |
| Global usage count | 0 |
| Local definition count | 0 |
| Documentation references | 0 |
| Initialization flow references | 0 |
| Known engine namespace | yes |
| Default UI presence | yes |
| Event binding presence | no |
| Observed in XML and Lua | no |
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

Observed XML element type instantiated by 9 addons.

## Common Attributes

- handleinput
- inherits
- layer
- name
- savesettings
- texture
- textureScale

## Common Inherits

- BuffHeadLayoutHorizontalResizeImage
- EA_BrownHorizontalRule
- EA_HorizontalResizeImage_DefaultTopFrame
- EA_HorizontalResizeImage_TabSeparatorMiddle

## Common Parent Elements

- [Windows](element_Windows.md) — 19× (HIGH)

## Common Structural Child Elements

- [Anchors](element_Anchors.md) — 11× (HIGH)
- [Sizes](element_Sizes.md) — 8× (HIGH)
- [TexCoords](element_TexCoords.md) — 8× (HIGH)
- [TintColor](element_TintColor.md) — 7× (HIGH)
- [Size](element_Size.md) — 4× (MEDIUM)

## Common Template Bases

- BuffHeadLayoutHorizontalResizeImage
- EA_BrownHorizontalRule
- EA_HorizontalResizeImage_DefaultTopFrame
- EA_HorizontalResizeImage_TabSeparatorMiddle

## Typical XML Structure

```xml
<HorizontalResizeImage handleinput="false" name="..." savesettings="false" texture="EA_TintableImage">
  <Sizes left="0" middle="30" right="0"/>
  <TexCoords>
    <Left x="0" y="0"/>
    <Middle x="0" y="0"/>
    <Right x="0" y="30"/>
  </TexCoords>
  <TintColor b="0" g="0" r="255"/>
</HorizontalResizeImage>
```

## Attribute Reference

| Attribute | Required | Usage % | Sample Values |
| --- | --- | --- | --- |
| `inherits` | optional | 57% | EA_HorizontalResizeImage_TabSeparatorMiddle, BuffHeadLayoutHorizontalResizeImage, EA_BrownHorizontalRule, EA_HorizontalResizeImage_DefaultTopFrame |
| `texture` | optional | 42% | EA_TintableImage, EA_Training_Specialization, EA_HUD_01 |
| `handleinput` | optional | 26% | false |
| `layer` | optional | 26% | background |
| `savesettings` | optional | 21% | false |
| `textureScale` | optional | 5% | 0.89 |
## Structural Sub-Elements

### [Anchors](element_Anchors.md)

Observed 11 times as an unnamed child.

### [Sizes](element_Sizes.md)

Observed 8 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `middle` | optional | 30, 200, 61, 50 |
| `left` | optional | 0, 7 |
| `right` | optional | 0 |
| `bottom` | optional | 0 |
| `top` | optional | 0 |
### [TexCoords](element_TexCoords.md)

Observed 8 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `x` | optional | 0, 133, 88, 32 |
| `y` | optional | 0, 163, 51, 32 |
### [TintColor](element_TintColor.md)

Observed 7 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `b` | **required** | 0, 110, 130, 50 |
| `g` | **required** | 0, 200, 185, 130 |
| `r` | **required** | 255, 0, 180, 200 |
| `a` | optional | 30, 255, 0.5, 200 |
### [Size](element_Size.md)

Observed 4 times as an unnamed child.

## Recursive Hierarchy

- Root: [HorizontalResizeImage](element_HorizontalResizeImage.md)
- [Anchors](element_Anchors.md) (structural, 11×, HIGH)
  - [AbsPoint](element_AbsPoint.md) (structural, 3×, MEDIUM)
  - [Anchor](element_Anchor.md) (structural, 3889×, HIGH)
    - [AbsPoint](element_AbsPoint.md) (structural, 3500×, HIGH)
    - [Anchor](element_Anchor.md) (structural, 22×, HIGH)
      - (cycle)
- [Size](element_Size.md) (structural, 4×, MEDIUM)
  - [AbsPoint](element_AbsPoint.md) (structural, 2634×, HIGH)
- [Sizes](element_Sizes.md) (structural, 8×, HIGH)
  - [BottomRight](element_BottomRight.md) (structural, 14×, HIGH)
  - [Middle](element_Middle.md) (structural, 14×, HIGH)
  - [TopLeft](element_TopLeft.md) (structural, 14×, HIGH)
- [TexCoords](element_TexCoords.md) (structural, 8×, HIGH)
  - [BottomCenter](element_BottomCenter.md) (structural, 9×, MEDIUM)
  - [BottomLeft](element_BottomLeft.md) (structural, 9×, MEDIUM)
  - [BottomRight](element_BottomRight.md) (structural, 9×, MEDIUM)
  - [Disabled](element_Disabled.md) (structural, 7×, MEDIUM)
  - [Left](element_Left.md) (structural, 7×, MEDIUM)
  - [Middle](element_Middle.md) (structural, 9×, MEDIUM)
  - [MiddleCenter](element_MiddleCenter.md) (structural, 9×, MEDIUM)
  - [MiddleLeft](element_MiddleLeft.md) (structural, 9×, MEDIUM)
  - [MiddleRight](element_MiddleRight.md) (structural, 9×, MEDIUM)
  - [Normal](element_Normal.md) (structural, 7×, MEDIUM)
  - [NormalHighlit](element_NormalHighlit.md) (structural, 7×, MEDIUM)
  - [Pressed](element_Pressed.md) (structural, 7×, MEDIUM)
  - [PressedHighlit](element_PressedHighlit.md) (structural, 1×, LOW)
  - [Right](element_Right.md) (structural, 7×, MEDIUM)
  - [TopCenter](element_TopCenter.md) (structural, 9×, MEDIUM)
  - [TopLeft](element_TopLeft.md) (structural, 9×, MEDIUM)
  - [TopRight](element_TopRight.md) (structural, 9×, MEDIUM)
- [TintColor](element_TintColor.md) (structural, 7×, HIGH)

## Lua Functions Manipulating This Type

- APAGui.UpdateKitingHUD
- APAGui.UpdateInstantOnlyHUD
- APAGui.UpdatePetTargetHUD
- PotionBarSettings.OnShown
- APAGui.UpdateFollowTargetHUD

## Seen In

- AdvancedPetAssist
- AggroMeter
- Aura
- BuffHead
- PartyCast
- PotionBar
- RoR_SoR
- Shinies
- WSCT

## Examples

- AdvancedPetAssist: APAFollowTargetHUDFill -> HorizontalResizeImage APAFollowTargetHUDFill
- AdvancedPetAssist: APAInstantOnlyHUDFill -> HorizontalResizeImage APAInstantOnlyHUDFill
- AdvancedPetAssist: APAKitingHUDFill -> HorizontalResizeImage APAKitingHUDFill
- AdvancedPetAssist: APAPetTargetHUDBg -> HorizontalResizeImage APAPetTargetHUDBg
- AggroMeter: AggroMeterGrayWindowSeparatorMiddle -> HorizontalResizeImage AggroMeterGrayWindowSeparatorMiddle
- Aura: AuraConfigTabsSeparatorActivation -> HorizontalResizeImage AuraConfigTabsSeparatorActivation

## Related APIs

- [EA_BrownHorizontalRule](../../globals/constants/constant_EA_BrownHorizontalRule.md) (HIGH 100/100) - Constant
- [EA_HorizontalResizeImage_TabSeparatorMiddle](../../globals/constants/constant_EA_HorizontalResizeImage_TabSeparatorMiddle.md) (HIGH 100/100) - Constant
- [Size](element_Size.md) (HIGH 100/100) - XML Element Type
- [Sizes](element_Sizes.md) (HIGH 100/100) - XML Element Type
- [TexCoords](element_TexCoords.md) (HIGH 100/100) - XML Element Type
- [TintColor](element_TintColor.md) (HIGH 100/100) - XML Element Type
- [Windows](element_Windows.md) (HIGH 100/100) - XML Element Type
- [EA_HorizontalResizeImage_DefaultTopFrame](../../globals/constants/constant_EA_HorizontalResizeImage_DefaultTopFrame.md) (HIGH 90/100) - Constant
- [Anchors](element_Anchors.md) (MEDIUM 55/100) - XML Element Type

## Used With

- none

## Triggered By

- none

## Affects

- none
