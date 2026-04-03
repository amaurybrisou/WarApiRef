# AnimatedImages

- Category: XML Element Type
- Confidence level: MEDIUM
- Confidence score: 45/100

## Confidence Assessment

- Level: MEDIUM

- Score: 45/100

- Rationale: Promoted as MEDIUM confidence because matches default ui or extracted base ui surface, used directly in xml handler attributes.

## Evidence Signals

- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +30 Used directly in XML handler attributes: XML exposure suggests an engine-level contract.
- -20 Only one weak usage site: Evidence is too shallow to trust as platform API.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | AdvancedRenownTrainer |
| Files seen in | `/workspace/data/raw/advancedrenowntrainer/AdvancedRenownTrainingWindow.xml:0` |
| Namespaces detected | AnimatedImages |
| Source kinds | xml_frames |
| Example locations | AdvancedRenownTrainer: AdvancedRenownTrainingWindowPurchaseButton |
| XML usage count | 1 |
| XML attribute usage count | 1 |
| Lua usage count | 0 |
| Global usage count | 0 |
| Local definition count | 0 |
| Documentation references | 0 |
| Initialization flow references | 0 |
| Known engine namespace | no |
| Default UI presence | yes |
| Event binding presence | no |
| Observed in XML and Lua | no |
| Consistent role | no |
| Consistent arguments | no |
| Consistent returns | no |
| Slash command presence | no |
| Weak usage only | yes |
| Project-specific name | no |
| Placeholder or computed name | no |
| Conflicting signatures | no |
| Conflicting roles | no |
| Wrapper likely | no |
| Never outside local graph | no |
| Local helper only | no |

## Description

AnimatedImages is a structural XML sub-element. It commonly appears under Button. It is typically used to organize structural children such as Normal, NormalHighlit.

## Common Attributes

- none

## Common Inherits

- none

## Common Parent Elements

- [Button](element_Button.md) — 1× (HIGH)

## Common Structural Child Elements

- [Normal](element_Normal.md) — 1× (HIGH)
- [NormalHighlit](element_NormalHighlit.md) — 1× (HIGH)

## Structural Sub-Elements

### [Normal](element_Normal.md)

Observed 1 times as an unnamed child.

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

Observed 1 times as an unnamed child.

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
## Recursive Hierarchy

- Root: [AnimatedImages](element_AnimatedImages.md)
- [Normal](element_Normal.md) (structural, 1×, HIGH)
- [NormalHighlit](element_NormalHighlit.md) (structural, 1×, HIGH)

## Seen In

- AdvancedRenownTrainer

## Examples

- AdvancedRenownTrainer: AdvancedRenownTrainingWindowPurchaseButton -> AnimatedImages in Button AdvancedRenownTrainingWindowPurchaseButton

## Related APIs

- [Button](element_Button.md) (HIGH 100/100) - XML Element Type
- [Normal](element_Normal.md) (MEDIUM 45/100) - XML Element Type
- [NormalHighlit](element_NormalHighlit.md) (MEDIUM 45/100) - XML Element Type

## Used With

- none

## Triggered By

- none

## Affects

- none
