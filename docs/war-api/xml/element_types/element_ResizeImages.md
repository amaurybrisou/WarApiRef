# ResizeImages

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
| Addons seen in | AdvancedPetAssist, Aura, BuffHead, Shinies, WSCT |
| Files seen in | `/workspace/data/raw/AdvancedPetAssist/APAGui.xml:0`, `/workspace/data/raw/Aura/Source/Templates.xml:0`, `/workspace/data/raw/BuffHead/Setup/SetupLayout.xml:0`, `/workspace/data/raw/Shinies/Source/ShiniesUITemplates.xml:0`, `/workspace/data/raw/wsct/wsct_options/wsct_options.xml:0` |
| Namespaces detected | ResizeImages |
| Source kinds | xml_frames |
| Example locations | AdvancedPetAssist: APA_ComboBoxButton, AdvancedPetAssist: APA_ComboBoxButtonWide, Aura: Aura_Button_DefaultMenuButton, Aura: Aura_Button_DefaultMenuButtonLarge, Aura: Aura_Button_DefaultMenuButtonTiny, Aura: Aura_Button_DefaultResizableTinyComboBoxSelected |
| XML usage count | 14 |
| XML attribute usage count | 14 |
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

Observed XML element type instantiated by 5 addons.

## Common Attributes

- none

## Common Inherits

- none

## Common Parent Elements

- [Button](element_Button.md) — 14× (HIGH)

## Common Structural Child Elements

- NormalHighlit — 14× (HIGH)
- Normal — 10× (HIGH)
- Pressed — 10× (HIGH)
- PressedHighlit — 9× (HIGH)
- Disabled — 8× (HIGH)

## Structural Sub-Elements

### NormalHighlit

Observed 14 times as an unnamed child.

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
### Normal

Observed 10 times as an unnamed child.

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
### Pressed

Observed 10 times as an unnamed child.

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
### PressedHighlit

Observed 9 times as an unnamed child.

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
### Disabled

Observed 8 times as an unnamed child.

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
## Seen In

- AdvancedPetAssist
- Aura
- BuffHead
- Shinies
- WSCT

## Examples

- AdvancedPetAssist: APA_ComboBoxButton -> ResizeImages in Button APA_ComboBoxButton
- AdvancedPetAssist: APA_ComboBoxButtonWide -> ResizeImages in Button APA_ComboBoxButtonWide
- Aura: Aura_Button_DefaultMenuButton -> ResizeImages in Button Aura_Button_DefaultMenuButton
- Aura: Aura_Button_DefaultMenuButtonLarge -> ResizeImages in Button Aura_Button_DefaultMenuButtonLarge
- Aura: Aura_Button_DefaultMenuButtonTiny -> ResizeImages in Button Aura_Button_DefaultMenuButtonTiny
- Aura: Aura_Button_DefaultResizableTinyComboBoxSelected -> ResizeImages in Button Aura_Button_DefaultResizableTinyComboBoxSelected

## Related APIs

- [Button](element_Button.md) (HIGH 100/100) - XML Element Type

## Used With

- none

## Triggered By

- none

## Affects

- none
