# TexCoords

- Category: XML Element Type
- Confidence level: HIGH
- Confidence score: 100/100

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 110

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, seen in 4 or more addons, used directly in xml handler attributes.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +30 Used directly in XML handler attributes: XML exposure suggests an engine-level contract.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | AdvancedPetAssist, AggroMeter, Aura, BuffHead, CM_ClosetGoblin, Enemy, GuardLine, MoraleCircle |
| Files seen in | `/workspace/data/raw/AdvancedPetAssist/APAGui.xml:0`, `/workspace/data/raw/AggroMeter/AggroMeter.xml:0`, `/workspace/data/raw/Aura/Source/Templates.xml:0`, `/workspace/data/raw/BuffHead/Setup/SetupLayout.xml:0`, `/workspace/data/raw/ClosetGoblin/ClosetGoblin.xml:0`, `/workspace/data/raw/Enemy/Code/Assist/Assist.xml:0`, `/workspace/data/raw/Enemy/Code/Core/Common.xml:0`, `/workspace/data/raw/Enemy/Code/Core/Icon.xml:0` |
| Namespaces detected | TexCoords |
| Source kinds | xml_frames |
| Example locations | AdvancedPetAssist: APAFollowTargetHUDFill, AdvancedPetAssist: APAInstantOnlyHUDFill, AdvancedPetAssist: APAKitingHUDFill, AdvancedPetAssist: APAPetTargetHUDBg, AggroMeter: Aggro_Button_TemplateIcon, Aura: AuraScreenFlashFrameImage |
| XML usage count | 59 |
| XML attribute usage count | 59 |
| Lua usage count | 0 |
| Global usage count | 0 |
| Local definition count | 0 |
| Documentation references | 0 |
| Initialization flow references | 0 |
| Known engine namespace | no |
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

TexCoords is a structural XML sub-element. It commonly appears under Button and CircleImage. It is typically used to organize structural children such as BottomCenter, BottomLeft, BottomRight.

## Common Attributes

- x
- y

## Common Inherits

- none

## Common Parent Elements

- [CircleImage](element_CircleImage.md) — 20× (HIGH)
- [DynamicImage](element_DynamicImage.md) — 14× (HIGH)
- [FullResizeImage](element_FullResizeImage.md) — 9× (MEDIUM)
- [HorizontalResizeImage](element_HorizontalResizeImage.md) — 8× (MEDIUM)
- [Button](element_Button.md) — 7× (MEDIUM)
- [VerticalResizeImage](element_VerticalResizeImage.md) — 1× (LOW)

## Common Structural Child Elements

- [BottomCenter](element_BottomCenter.md) — 9× (MEDIUM)
- [BottomLeft](element_BottomLeft.md) — 9× (MEDIUM)
- [BottomRight](element_BottomRight.md) — 9× (MEDIUM)
- [Middle](element_Middle.md) — 9× (MEDIUM)
- [MiddleCenter](element_MiddleCenter.md) — 9× (MEDIUM)
- [MiddleLeft](element_MiddleLeft.md) — 9× (MEDIUM)
- [MiddleRight](element_MiddleRight.md) — 9× (MEDIUM)
- [TopCenter](element_TopCenter.md) — 9× (MEDIUM)
- [TopLeft](element_TopLeft.md) — 9× (MEDIUM)
- [TopRight](element_TopRight.md) — 9× (MEDIUM)
- [Disabled](element_Disabled.md) — 7× (MEDIUM)
- [Left](element_Left.md) — 7× (MEDIUM)
- [Normal](element_Normal.md) — 7× (MEDIUM)
- [NormalHighlit](element_NormalHighlit.md) — 7× (MEDIUM)
- [Pressed](element_Pressed.md) — 7× (MEDIUM)
- [Right](element_Right.md) — 7× (MEDIUM)
- [PressedHighlit](element_PressedHighlit.md) — 1× (LOW)

## Attribute Reference

| Attribute | Required | Usage % | Sample Values |
| --- | --- | --- | --- |
| `x` | optional | 57% | 0, 133, 88, 32, ... |
| `y` | optional | 57% | 0, 163, 51, 32, ... |
## Structural Sub-Elements

### [BottomCenter](element_BottomCenter.md)

Observed 9 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `x` | optional | 272, 60, 356, 26 |
| `y` | optional | 197, 35, 684, 38 |
| `id` | optional | Border-Bottom-Center |
### [BottomLeft](element_BottomLeft.md)

Observed 9 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `x` | optional | 256, 0, 346 |
| `y` | optional | 197, 35, 684, 38 |
| `id` | optional | Border-Bottom-Left |
### [BottomRight](element_BottomRight.md)

Observed 9 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `x` | **required** | 16, 350, 2, 135 |
| `y` | **required** | 15, 197, 2, 35 |
| `id` | optional | Border-Bottom-Right |
### [Middle](element_Middle.md)

Observed 9 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `x` | **required** | 0, 424, 78, 2 |
| `y` | **required** | 0, 762, 25, 2 |
### [MiddleCenter](element_MiddleCenter.md)

Observed 9 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `x` | optional | 272, 60, 356, 32 |
| `y` | optional | 172, 18, 680, 32 |
| `id` | optional | Border-Middle-Center |
### [MiddleLeft](element_MiddleLeft.md)

Observed 9 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `x` | optional | 256, 0, 346 |
| `y` | optional | 172, 18, 680, 25 |
| `id` | optional | Border-Middle-Left |
### [MiddleRight](element_MiddleRight.md)

Observed 9 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `x` | optional | 350, 135, 346, 39 |
| `y` | optional | 172, 18, 680, 25 |
| `id` | optional | Border-Middle-Right |
### [TopCenter](element_TopCenter.md)

Observed 9 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `x` | optional | 272, 60, 356, 26 |
| `y` | optional | 157, 0, 669 |
| `id` | optional | Border-Top-Center |
### [TopLeft](element_TopLeft.md)

Observed 9 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `x` | **required** | 16, 256, 2, 0 |
| `y` | **required** | 15, 157, 2, 0 |
| `id` | optional | Border-Top-Left |
### [TopRight](element_TopRight.md)

Observed 9 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `x` | optional | 350, 135, 400, 39 |
| `y` | optional | 157, 0, 669 |
| `id` | optional | Border-Top-Right |
### [Disabled](element_Disabled.md)

Observed 7 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `b` | optional | 92, 36, 102 |
| `g` | optional | 92, 57, 204 |
| `r` | optional | 92, 95, 255 |
| `a` | optional | 255 |
| `x` | optional | 27, 92, 0 |
| `y` | optional | 56, 44, 0 |
| `id` | optional | morale-white, RightTabFrame, ability-white, tactic-black |
| `def` | optional | EA_HorizontalResizeImage_DefaultComboBox, EA_Button_ListSortNormal |
| `texture` | optional | EA_SquareFrame, ShiniesIconBorderNormal, TidyRoll_SquareFrame |
### [Left](element_Left.md)

Observed 7 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `x` | **required** | 0, 72, 346 |
| `y` | **required** | 0, 241, 655 |
### [Normal](element_Normal.md)

Observed 7 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `x` | optional | 0, 92 |
| `y` | optional | 28, 44, 0 |
| `b` | optional | 255, 73, 102 |
| `g` | optional | 255, 175, 204 |
| `id` | optional | morale-yellow, RightTabFrame, LayoutCorner-TopLeft, LayoutCorner-TopRight |
| `r` | optional | 255 |
| `a` | optional | 255 |
| `texture` | optional | bpKtxt, EA_SquareFrame, ShiniesIconBorderNormal, TidyRoll_SquareFrame |
| `def` | optional | EA_HorizontalResizeImage_DefaultComboBox, EA_AnimatedImage_DefaultChoiceOverlay, BuffHeadLayoutVerticalButtonNormal, BuffHeadLayoutHorizontalButtonNormal |
### [NormalHighlit](element_NormalHighlit.md)

Observed 7 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `b` | optional | 63, 0 |
| `g` | optional | 213, 85 |
| `id` | optional | morale-white, RightTabFrame-Rollover, LayoutCorner-TopLeft-ROLLOVER, LayoutCorner-TopRight-ROLLOVER |
| `r` | optional | 250, 255 |
| `a` | optional | 255 |
| `def` | optional | EA_HorizontalResizeImage_DefaultComboBox, EA_AnimatedImage_DefaultChoiceOverlay, EA_FullResizeImage_RedTransparent, BuffHeadLayoutVerticalButtonHighlight |
| `x` | optional | 27, 105, 0 |
| `y` | optional | 28, 44, 0 |
| `texture` | optional | EA_SquareFrame_Highlight, ShiniesIconBorderHighlight, TidyRoll_SquareFrame_Highlight |
### [Pressed](element_Pressed.md)

Observed 7 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `b` | optional | 63, 36, 0 |
| `g` | optional | 213, 57, 85 |
| `id` | optional | morale-white, RightTabFrame-Rollover, LayoutCorner-TopLeft-ROLLOVER, LayoutCorner-TopRight-ROLLOVER |
| `r` | optional | 250, 95, 255 |
| `a` | optional | 255 |
| `x` | optional | 0, 120 |
| `y` | optional | 56, 44, 0 |
| `def` | optional | EA_HorizontalResizeImage_DefaultComboBox, BuffHeadLayoutVerticalButtonPressed, BuffHeadLayoutHorizontalButtonPressed, EA_Button_ListSortPressed |
| `texture` | optional | EA_SquareFrame_Pressed, ShiniesIconBorderHighlight, TidyRoll_SquareFrame |
### [Right](element_Right.md)

Observed 7 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `x` | **required** | 0, 119, 396 |
| `y` | **required** | 30, 24, 241, 655 |
### [PressedHighlit](element_PressedHighlit.md)

Observed 1 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `b` | optional | 63, 36, 0 |
| `g` | optional | 213, 57, 85 |
| `id` | optional | morale-white, RightTabFrame-Rollover, LayoutCorner-TopLeft-ROLLOVER, LayoutCorner-TopRight-ROLLOVER |
| `r` | optional | 250, 95, 255 |
| `a` | optional | 255 |
| `def` | optional | EA_HorizontalResizeImage_DefaultComboBox, BuffHeadLayoutVerticalButtonPressed, BuffHeadLayoutHorizontalButtonPressed |
| `x` | optional | 0, 120 |
| `y` | optional | 56, 44 |
## Recursive Hierarchy

- Root: [TexCoords](element_TexCoords.md)
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

## Seen In

- AdvancedPetAssist
- AggroMeter
- Aura
- BuffHead
- CM_ClosetGoblin
- Enemy
- GuardLine
- MoraleCircle
- PartyCast
- Pocket Palette
- PotionBar
- RoR_SoR
- Shinies
- TidyRoll

## Examples

- AdvancedPetAssist: APAFollowTargetHUDFill -> TexCoords in HorizontalResizeImage APAFollowTargetHUDFill
- AdvancedPetAssist: APAInstantOnlyHUDFill -> TexCoords in HorizontalResizeImage APAInstantOnlyHUDFill
- AdvancedPetAssist: APAKitingHUDFill -> TexCoords in HorizontalResizeImage APAKitingHUDFill
- AdvancedPetAssist: APAPetTargetHUDBg -> TexCoords in HorizontalResizeImage APAPetTargetHUDBg
- AggroMeter: Aggro_Button_TemplateIcon -> TexCoords in CircleImage Aggro_Button_TemplateIcon
- Aura: AuraScreenFlashFrameImage -> TexCoords in DynamicImage AuraScreenFlashFrameImage

## Related APIs

- [Button](element_Button.md) (HIGH 100/100) - XML Element Type
- [CircleImage](element_CircleImage.md) (HIGH 100/100) - XML Element Type
- [DynamicImage](element_DynamicImage.md) (HIGH 100/100) - XML Element Type
- [FullResizeImage](element_FullResizeImage.md) (HIGH 100/100) - XML Element Type
- [HorizontalResizeImage](element_HorizontalResizeImage.md) (HIGH 100/100) - XML Element Type
- [VerticalResizeImage](element_VerticalResizeImage.md) (HIGH 90/100) - XML Element Type
- [BottomCenter](element_BottomCenter.md) (MEDIUM 45/100) - XML Element Type
- [BottomLeft](element_BottomLeft.md) (MEDIUM 45/100) - XML Element Type
- [BottomRight](element_BottomRight.md) (MEDIUM 45/100) - XML Element Type
- [Disabled](element_Disabled.md) (MEDIUM 45/100) - XML Element Type
- [Middle](element_Middle.md) (MEDIUM 45/100) - XML Element Type
- [MiddleCenter](element_MiddleCenter.md) (MEDIUM 45/100) - XML Element Type
- [MiddleLeft](element_MiddleLeft.md) (MEDIUM 45/100) - XML Element Type
- [MiddleRight](element_MiddleRight.md) (MEDIUM 45/100) - XML Element Type
- [Normal](element_Normal.md) (MEDIUM 45/100) - XML Element Type
- [NormalHighlit](element_NormalHighlit.md) (MEDIUM 45/100) - XML Element Type
- [Pressed](element_Pressed.md) (MEDIUM 45/100) - XML Element Type
- [PressedHighlit](element_PressedHighlit.md) (MEDIUM 45/100) - XML Element Type
- [TopCenter](element_TopCenter.md) (MEDIUM 45/100) - XML Element Type
- [TopLeft](element_TopLeft.md) (MEDIUM 45/100) - XML Element Type
- [TopRight](element_TopRight.md) (MEDIUM 45/100) - XML Element Type
- [Left](element_Left.md) (MEDIUM 30/100) - XML Element Type
- [Right](element_Right.md) (MEDIUM 30/100) - XML Element Type

## Used With

- [OverlayTexCoords](element_OverlayTexCoords.md) (HIGH 100/100) - XML Element Type
- [ResizeImages](element_ResizeImages.md) (HIGH 100/100) - XML Element Type
- [Sizes](element_Sizes.md) (HIGH 100/100) - XML Element Type
- [TexSlices](element_TexSlices.md) (HIGH 100/100) - XML Element Type

## Triggered By

- none

## Affects

- none
