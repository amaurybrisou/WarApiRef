# OverlayTexSlices

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
| Addons seen in | SpamBayes |
| Namespaces detected | OverlayTexSlices |
| Source kinds | xml_frames |
| Example locations | SpamBayes: SpamBayesTemplateScrollBarDownArrowButton, SpamBayes: SpamBayesTemplateScrollBarUpArrowButton |
| XML usage count | 2 |
| XML attribute usage count | 2 |
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

OverlayTexSlices is a structural XML sub-element. It commonly appears under Button. It is typically used to organize structural children such as Normal, NormalHighlit, Pressed.

## Common Attributes

- none

## Common Inherits

- none

## Common Parent Elements

- [Button](element_Button.md) — 2× (HIGH)

## Common Structural Child Elements

- [Normal](element_Normal.md) — 2× (HIGH)
- [NormalHighlit](element_NormalHighlit.md) — 2× (HIGH)
- [Pressed](element_Pressed.md) — 2× (HIGH)

## Structural Sub-Elements

### [Normal](element_Normal.md)

Observed 2 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `id` | optional | morale-yellow, RightTabFrame, Tactics-Button, SquareButton |
| `b` | optional | 255, 0, 102, 73 |
| `g` | optional | 255, 0, 204, 175 |
| `r` | optional | 255, 155, 222, 226 |
| `x` | optional | 0, 92, 172, 494 |
| `y` | optional | 28, 44, 0, 341 |
| `a` | optional | 255, 1 |
| `texture` | optional | bpKtxt, EA_SquareFrame, CrusherIconBorderNormal, EA_RoundFrame |
| `def` | optional | EA_HorizontalResizeImage_DefaultComboBox, EA_AnimatedImage_DefaultChoiceOverlay, BuffHeadLayoutVerticalButtonNormal, BuffHeadLayoutHorizontalButtonNormal |
### [NormalHighlit](element_NormalHighlit.md)

Observed 2 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `id` | optional | morale-white, RightTabFrame-Rollover, Tactics-Button-Rollover, SquareButton-Rollover |
| `b` | optional | 63, 36, 0, 50 |
| `g` | optional | 213, 57, 85, 192 |
| `r` | optional | 250, 95, 255, 222 |
| `a` | optional | 255, 1 |
| `x` | optional | 27, 105, 0, 201 |
| `y` | optional | 28, 44, 0, 341 |
| `def` | optional | EA_HorizontalResizeImage_DefaultComboBox, EA_AnimatedImage_DefaultChoiceOverlay, EA_FullResizeImage_RedTransparent, EA_Button_GuildRosterRowHighlight |
| `texture` | optional | EA_SquareFrame_Highlight, CrusherIconBorderHighlight, EA_RoundFrame_Pressed, PinBG |
### [Pressed](element_Pressed.md)

Observed 2 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `id` | optional | morale-white, RightTabFrame-Rollover, Tactics-Button-Depressed, SquareButton-Depressed |
| `b` | optional | 63, 36, 102, 235 |
| `g` | optional | 213, 57, 204, 235 |
| `r` | optional | 250, 95, 255, 235 |
| `a` | optional | 255, 1 |
| `x` | optional | 0, 120, 172, 494 |
| `y` | optional | 56, 44, 0, 370 |
| `def` | optional | EA_HorizontalResizeImage_DefaultComboBox, EA_Button_GuildRosterRowPressed, BuffHeadLayoutVerticalButtonPressed, BuffHeadLayoutHorizontalButtonPressed |
| `texture` | optional | EA_SquareFrame_Pressed, CrusherIconBorderHighlight, EA_RoundFrame_Highlight, PinBG |
## Recursive Hierarchy

- Root: [OverlayTexSlices](element_OverlayTexSlices.md)
- [Normal](element_Normal.md) (structural, 2×, HIGH)
- [NormalHighlit](element_NormalHighlit.md) (structural, 2×, HIGH)
- [Pressed](element_Pressed.md) (structural, 2×, HIGH)

## Seen In

- SpamBayes

## Examples

- SpamBayes: SpamBayesTemplateScrollBarDownArrowButton -> OverlayTexSlices in Button SpamBayesTemplateScrollBarDownArrowButton
- SpamBayes: SpamBayesTemplateScrollBarUpArrowButton -> OverlayTexSlices in Button SpamBayesTemplateScrollBarUpArrowButton

## Related APIs

- [Button](element_Button.md) (HIGH 100/100) - XML Element Type
- [Normal](element_Normal.md) (MEDIUM 45/100) - XML Element Type
- [NormalHighlit](element_NormalHighlit.md) (MEDIUM 45/100) - XML Element Type
- [Pressed](element_Pressed.md) (MEDIUM 45/100) - XML Element Type
