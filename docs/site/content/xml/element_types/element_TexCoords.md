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

Observed XML element type instantiated by 14 addons.

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

- Middle — 9× (MEDIUM)
- Disabled — 7× (MEDIUM)
- Left — 7× (MEDIUM)
- Normal — 7× (MEDIUM)
- NormalHighlit — 7× (MEDIUM)
- Pressed — 7× (MEDIUM)
- Right — 7× (MEDIUM)
- PressedHighlit — 1× (LOW)

## Attribute Reference

| Attribute | Required | Usage % | Sample Values |
| --- | --- | --- | --- |
| `x` | optional | 57% | 0, 133, 88, 32, ... |
| `y` | optional | 57% | 0, 163, 51, 32, ... |
## Structural Sub-Elements

### Middle

Observed 9 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `x` | **required** | 0, 424, 78, 2 |
| `y` | **required** | 0, 762, 25, 2 |
### Disabled

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
### Left

Observed 7 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `x` | **required** | 0, 72, 346 |
| `y` | **required** | 0, 241, 655 |
### Normal

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
### NormalHighlit

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
### Pressed

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
### Right

Observed 7 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `x` | **required** | 0, 119, 396 |
| `y` | **required** | 30, 24, 241, 655 |
### PressedHighlit

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

## Used With

- [Size](element_Size.md) (HIGH 100/100) - XML Element Type
- [TintColor](element_TintColor.md) (HIGH 100/100) - XML Element Type
- [Windows](element_Windows.md) (HIGH 100/100) - XML Element Type

## Triggered By

- none

## Affects

- none
