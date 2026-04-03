# TexSlices

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
| Addons seen in | AggroMeter, AnywhereTrainer, BankArkel, BuffHead, PartyCast, PotionBar, RoR_SoR, Shinies |
| Files seen in | `/workspace/data/raw/AggroMeter/AggroMeter.xml:0`, `/workspace/data/raw/AnywhereTrainer/source/AnywhereTrainer.xml:0`, `/workspace/data/raw/BankArkel/BankArkel.xml:0`, `/workspace/data/raw/BuffHead/Setup/SetupLayout.xml:0`, `/workspace/data/raw/PartyCast/PartyCast.xml:0`, `/workspace/data/raw/PotionBar/settings/Settings.xml:22`, `/workspace/data/raw/PotionBar/settings/Settings.xml:45`, `/workspace/data/raw/RoR_SoR/RoR_SoR.xml:0` |
| Namespaces detected | TexSlices |
| Source kinds | xml_frames |
| Example locations | AggroMeter: Aggro_Button_Template, AnywhereTrainer: AnywhereTrainerTabTemplateInactiveImage, BankArkel: PackWinTab1, BankArkel: PackWinTab2, BankArkel: PackWinTab3, BankArkel: PackWinTab4 |
| XML usage count | 27 |
| XML attribute usage count | 27 |
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

TexSlices is a structural XML sub-element. It commonly appears under Button and FullResizeImage. It is typically used to organize structural children such as BottomCenter, BottomLeft, BottomRight.

## Common Attributes

- none

## Common Inherits

- none

## Common Parent Elements

- [Button](element_Button.md) — 22× (HIGH)
- [FullResizeImage](element_FullResizeImage.md) — 5× (MEDIUM)

## Common Structural Child Elements

- [Normal](element_Normal.md) — 22× (HIGH)
- [NormalHighlit](element_NormalHighlit.md) — 16× (HIGH)
- [Pressed](element_Pressed.md) — 16× (HIGH)
- [PressedHighlit](element_PressedHighlit.md) — 16× (HIGH)
- [Disabled](element_Disabled.md) — 11× (HIGH)
- [DisabledPressed](element_DisabledPressed.md) — 11× (HIGH)
- [BottomCenter](element_BottomCenter.md) — 5× (MEDIUM)
- [BottomLeft](element_BottomLeft.md) — 5× (MEDIUM)
- [BottomRight](element_BottomRight.md) — 5× (MEDIUM)
- [MiddleCenter](element_MiddleCenter.md) — 5× (MEDIUM)
- [MiddleLeft](element_MiddleLeft.md) — 5× (MEDIUM)
- [MiddleRight](element_MiddleRight.md) — 5× (MEDIUM)
- [TopCenter](element_TopCenter.md) — 5× (MEDIUM)
- [TopLeft](element_TopLeft.md) — 5× (MEDIUM)
- [TopRight](element_TopRight.md) — 5× (MEDIUM)

## Structural Sub-Elements

### [Normal](element_Normal.md)

Observed 22 times as an unnamed child.

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

Observed 16 times as an unnamed child.

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

Observed 16 times as an unnamed child.

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
### [PressedHighlit](element_PressedHighlit.md)

Observed 16 times as an unnamed child.

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
### [Disabled](element_Disabled.md)

Observed 11 times as an unnamed child.

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
### [DisabledPressed](element_DisabledPressed.md)

Observed 11 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `id` | optional | morale-white, RightTabFrame, ability-white, tactic-black |
| `a` | optional | 255 |
| `b` | optional | 36 |
| `g` | optional | 57 |
| `r` | optional | 95 |
### [BottomCenter](element_BottomCenter.md)

Observed 5 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `x` | optional | 272, 60, 356, 26 |
| `y` | optional | 197, 35, 684, 38 |
| `id` | optional | Border-Bottom-Center |
### [BottomLeft](element_BottomLeft.md)

Observed 5 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `x` | optional | 256, 0, 346 |
| `y` | optional | 197, 35, 684, 38 |
| `id` | optional | Border-Bottom-Left |
### [BottomRight](element_BottomRight.md)

Observed 5 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `x` | **required** | 16, 350, 2, 135 |
| `y` | **required** | 15, 197, 2, 35 |
| `id` | optional | Border-Bottom-Right |
### [MiddleCenter](element_MiddleCenter.md)

Observed 5 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `x` | optional | 272, 60, 356, 32 |
| `y` | optional | 172, 18, 680, 32 |
| `id` | optional | Border-Middle-Center |
### [MiddleLeft](element_MiddleLeft.md)

Observed 5 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `x` | optional | 256, 0, 346 |
| `y` | optional | 172, 18, 680, 25 |
| `id` | optional | Border-Middle-Left |
### [MiddleRight](element_MiddleRight.md)

Observed 5 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `x` | optional | 350, 135, 346, 39 |
| `y` | optional | 172, 18, 680, 25 |
| `id` | optional | Border-Middle-Right |
### [TopCenter](element_TopCenter.md)

Observed 5 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `x` | optional | 272, 60, 356, 26 |
| `y` | optional | 157, 0, 669 |
| `id` | optional | Border-Top-Center |
### [TopLeft](element_TopLeft.md)

Observed 5 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `x` | **required** | 16, 256, 2, 0 |
| `y` | **required** | 15, 157, 2, 0 |
| `id` | optional | Border-Top-Left |
### [TopRight](element_TopRight.md)

Observed 5 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `x` | optional | 350, 135, 400, 39 |
| `y` | optional | 157, 0, 669 |
| `id` | optional | Border-Top-Right |
## Recursive Hierarchy

- Root: [TexSlices](element_TexSlices.md)
- [BottomCenter](element_BottomCenter.md) (structural, 5×, MEDIUM)
- [BottomLeft](element_BottomLeft.md) (structural, 5×, MEDIUM)
- [BottomRight](element_BottomRight.md) (structural, 5×, MEDIUM)
- [Disabled](element_Disabled.md) (structural, 11×, HIGH)
- [DisabledPressed](element_DisabledPressed.md) (structural, 11×, HIGH)
- [MiddleCenter](element_MiddleCenter.md) (structural, 5×, MEDIUM)
- [MiddleLeft](element_MiddleLeft.md) (structural, 5×, MEDIUM)
- [MiddleRight](element_MiddleRight.md) (structural, 5×, MEDIUM)
- [Normal](element_Normal.md) (structural, 22×, HIGH)
- [NormalHighlit](element_NormalHighlit.md) (structural, 16×, HIGH)
- [Pressed](element_Pressed.md) (structural, 16×, HIGH)
- [PressedHighlit](element_PressedHighlit.md) (structural, 16×, HIGH)
- [TopCenter](element_TopCenter.md) (structural, 5×, MEDIUM)
- [TopLeft](element_TopLeft.md) (structural, 5×, MEDIUM)
- [TopRight](element_TopRight.md) (structural, 5×, MEDIUM)

## Seen In

- AggroMeter
- AnywhereTrainer
- BankArkel
- BuffHead
- PartyCast
- PotionBar
- RoR_SoR
- Shinies

## Examples

- AggroMeter: Aggro_Button_Template -> TexSlices in Button Aggro_Button_Template
- AnywhereTrainer: AnywhereTrainerTabTemplateInactiveImage -> TexSlices in Button AnywhereTrainerTabTemplateInactiveImage
- BankArkel: PackWinTab1 -> TexSlices in Button PackWinTab1
- BankArkel: PackWinTab2 -> TexSlices in Button PackWinTab2
- BankArkel: PackWinTab3 -> TexSlices in Button PackWinTab3
- BankArkel: PackWinTab4 -> TexSlices in Button PackWinTab4

## Related APIs

- [Button](element_Button.md) (HIGH 100/100) - XML Element Type
- [FullResizeImage](element_FullResizeImage.md) (HIGH 100/100) - XML Element Type
- [BottomCenter](element_BottomCenter.md) (MEDIUM 45/100) - XML Element Type
- [BottomLeft](element_BottomLeft.md) (MEDIUM 45/100) - XML Element Type
- [BottomRight](element_BottomRight.md) (MEDIUM 45/100) - XML Element Type
- [Disabled](element_Disabled.md) (MEDIUM 45/100) - XML Element Type
- [DisabledPressed](element_DisabledPressed.md) (MEDIUM 45/100) - XML Element Type
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

## Used With

- [OverlayTexCoords](element_OverlayTexCoords.md) (HIGH 100/100) - XML Element Type
- [ResizeImages](element_ResizeImages.md) (HIGH 100/100) - XML Element Type
- [TexCoords](element_TexCoords.md) (HIGH 100/100) - XML Element Type
- [TextColors](element_TextColors.md) (HIGH 100/100) - XML Element Type

## Triggered By

- none

## Affects

- none
