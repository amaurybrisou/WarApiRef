# OverlayTexCoords

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
| Addons seen in | AdvancedPetAssist, Aura, Shinies, WSCT |
| Files seen in | `/workspace/data/raw/AdvancedPetAssist/APAGui.xml:0`, `/workspace/data/raw/Aura/Source/Templates.xml:0`, `/workspace/data/raw/Shinies/Source/ShiniesUITemplates.xml:0`, `/workspace/data/raw/wsct/wsct_options/wsct_options.xml:0` |
| Namespaces detected | OverlayTexCoords |
| Source kinds | xml_frames |
| Example locations | AdvancedPetAssist: APA_ComboBoxButton, AdvancedPetAssist: APA_ComboBoxButtonWide, Aura: Aura_Button_DefaultResizableTinyComboBoxSelected, Aura: Aura_ComboBox_DefaultResizableComboBoxSelected, Aura: Aura_ComboBox_DefaultResizableComboBoxSelectedLarge, Shinies: Shinies_ComboBox_DefaultResizableComboBoxSelectedLarge |
| XML usage count | 7 |
| XML attribute usage count | 7 |
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

OverlayTexCoords is a structural XML sub-element. It commonly appears under Button. It is typically used to organize structural children such as Disabled, Normal, NormalHighlit.

## Common Attributes

- none

## Common Inherits

- none

## Common Parent Elements

- [Button](element_Button.md) — 7× (HIGH)

## Common Structural Child Elements

- [Disabled](element_Disabled.md) — 7× (HIGH)
- [Normal](element_Normal.md) — 7× (HIGH)
- [NormalHighlit](element_NormalHighlit.md) — 7× (HIGH)
- [Pressed](element_Pressed.md) — 7× (HIGH)
- [PressedHighlit](element_PressedHighlit.md) — 7× (HIGH)

## Structural Sub-Elements

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
### [PressedHighlit](element_PressedHighlit.md)

Observed 7 times as an unnamed child.

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

- Root: [OverlayTexCoords](element_OverlayTexCoords.md)
- [Disabled](element_Disabled.md) (structural, 7×, HIGH)
- [Normal](element_Normal.md) (structural, 7×, HIGH)
- [NormalHighlit](element_NormalHighlit.md) (structural, 7×, HIGH)
- [Pressed](element_Pressed.md) (structural, 7×, HIGH)
- [PressedHighlit](element_PressedHighlit.md) (structural, 7×, HIGH)

## Seen In

- AdvancedPetAssist
- Aura
- Shinies
- WSCT

## Examples

- AdvancedPetAssist: APA_ComboBoxButton -> OverlayTexCoords in Button APA_ComboBoxButton
- AdvancedPetAssist: APA_ComboBoxButtonWide -> OverlayTexCoords in Button APA_ComboBoxButtonWide
- Aura: Aura_Button_DefaultResizableTinyComboBoxSelected -> OverlayTexCoords in Button Aura_Button_DefaultResizableTinyComboBoxSelected
- Aura: Aura_ComboBox_DefaultResizableComboBoxSelected -> OverlayTexCoords in Button Aura_ComboBox_DefaultResizableComboBoxSelected
- Aura: Aura_ComboBox_DefaultResizableComboBoxSelectedLarge -> OverlayTexCoords in Button Aura_ComboBox_DefaultResizableComboBoxSelectedLarge
- Shinies: Shinies_ComboBox_DefaultResizableComboBoxSelectedLarge -> OverlayTexCoords in Button Shinies_ComboBox_DefaultResizableComboBoxSelectedLarge

## Related APIs

- [Button](element_Button.md) (HIGH 100/100) - XML Element Type
- [Disabled](element_Disabled.md) (MEDIUM 45/100) - XML Element Type
- [Normal](element_Normal.md) (MEDIUM 45/100) - XML Element Type
- [NormalHighlit](element_NormalHighlit.md) (MEDIUM 45/100) - XML Element Type
- [Pressed](element_Pressed.md) (MEDIUM 45/100) - XML Element Type
- [PressedHighlit](element_PressedHighlit.md) (MEDIUM 45/100) - XML Element Type

## Used With

- [ResizeImages](element_ResizeImages.md) (HIGH 100/100) - XML Element Type
- [TexCoords](element_TexCoords.md) (HIGH 100/100) - XML Element Type
- [TexSlices](element_TexSlices.md) (HIGH 100/100) - XML Element Type

## Triggered By

- none

## Affects

- none
