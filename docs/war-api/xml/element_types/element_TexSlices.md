# TexSlices

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
| Addons seen in | PartyCast |
| Files seen in | `/workspace/data/raw/PartyCast/PartyCast.xml:186`, `/workspace/data/raw/PartyCast/PartyCast.xml:327`, `/workspace/data/raw/PartyCast/PartyCast.xml:469`, `/workspace/data/raw/PartyCast/PartyCast.xml:631`, `/workspace/data/raw/PartyCast/PartyCast.xml:80` |
| Namespaces detected | TexSlices |
| Source kinds | xml_frames |
| Example locations | PartyCast: PartyCastWindow_TemplateButton, PartyCast: PartyCastWindow_Template_LargeButton, PartyCast: PartyCastWindow_Template_PlainButton, PartyCast: PartyCastWindow_Template_SmallButton, PartyCast: PartyCastWindow_Template_SpecialButton |
| XML usage count | 5 |
| XML attribute usage count | 5 |
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

Observed XML element type instantiated by 1 addons.

## Common Attributes

- none

## Common Inherits

- none

## Common Parent Elements

- [Button](element_Button.md) — 5× (HIGH)

## Common Structural Child Elements

- [Disabled](element_Disabled.md) — 5× (HIGH)
- [DisabledPressed](element_DisabledPressed.md) — 5× (HIGH)
- [Normal](element_Normal.md) — 5× (HIGH)
- [NormalHighlit](element_NormalHighlit.md) — 5× (HIGH)
- [Pressed](element_Pressed.md) — 5× (HIGH)
- [PressedHighlit](element_PressedHighlit.md) — 5× (HIGH)

## Structural Sub-Elements

### [Disabled](element_Disabled.md)

Observed 5 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `id` | **required** | morale-white, ability-white, tactic-black |
| `texture` | optional | TidyRoll_SquareFrame |
| `x` | optional | 0 |
| `y` | optional | 0 |
### [DisabledPressed](element_DisabledPressed.md)

Observed 5 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `id` | **required** | morale-white, ability-white, tactic-black |
### [Normal](element_Normal.md)

Observed 5 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `id` | **required** | morale-white, ability-white, tactic-black |
| `texture` | optional | TidyRoll_SquareFrame |
| `x` | optional | 0 |
| `y` | optional | 0 |
### [NormalHighlit](element_NormalHighlit.md)

Observed 5 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `id` | **required** | morale-white, ability-white, tactic-black |
| `texture` | optional | TidyRoll_SquareFrame_Highlight |
| `x` | optional | 0 |
| `y` | optional | 0 |
### [Pressed](element_Pressed.md)

Observed 5 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `id` | **required** | morale-white, ability-white, tactic-black |
| `texture` | optional | TidyRoll_SquareFrame |
| `x` | optional | 0 |
| `y` | optional | 0 |
### [PressedHighlit](element_PressedHighlit.md)

Observed 5 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `id` | **required** | morale-white, ability-white, tactic-black |
## Seen In

- PartyCast

## Examples

- PartyCast: PartyCastWindow_TemplateButton -> TexSlices in Button PartyCastWindow_TemplateButton
- PartyCast: PartyCastWindow_Template_LargeButton -> TexSlices in Button PartyCastWindow_Template_LargeButton
- PartyCast: PartyCastWindow_Template_PlainButton -> TexSlices in Button PartyCastWindow_Template_PlainButton
- PartyCast: PartyCastWindow_Template_SmallButton -> TexSlices in Button PartyCastWindow_Template_SmallButton
- PartyCast: PartyCastWindow_Template_SpecialButton -> TexSlices in Button PartyCastWindow_Template_SpecialButton

## Related APIs

- [Button](element_Button.md) (HIGH 100/100) - XML Element Type
- [Disabled](element_Disabled.md) (HIGH 98/100) - XML Element Type
- [Normal](element_Normal.md) (HIGH 98/100) - XML Element Type
- [NormalHighlit](element_NormalHighlit.md) (HIGH 98/100) - XML Element Type
- [Pressed](element_Pressed.md) (HIGH 98/100) - XML Element Type
- [DisabledPressed](element_DisabledPressed.md) (MEDIUM 45/100) - XML Element Type
- [PressedHighlit](element_PressedHighlit.md) (MEDIUM 45/100) - XML Element Type

## Used With

- none

## Triggered By

- none

## Affects

- none
