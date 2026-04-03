# Normal

- Category: XML Element Type
- Confidence level: HIGH
- Confidence score: 98/100

## Confidence Assessment

- Level: HIGH

- Score: 98/100

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, used directly in xml handler attributes, seen in 2 to 3 addons.

## Evidence Signals

- +18 Seen in 2 to 3 addons: Cross-addon spread is present but limited.
- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +30 Used directly in XML handler attributes: XML exposure suggests an engine-level contract.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | PartyCast, TidyRoll |
| Files seen in | `/workspace/data/raw/PartyCast/PartyCast.xml:186`, `/workspace/data/raw/PartyCast/PartyCast.xml:327`, `/workspace/data/raw/PartyCast/PartyCast.xml:469`, `/workspace/data/raw/PartyCast/PartyCast.xml:631`, `/workspace/data/raw/PartyCast/PartyCast.xml:80`, `/workspace/data/raw/TidyRoll/TidyRoll.xml:103` |
| Namespaces detected | Normal |
| Source kinds | xml_frames |
| Example locations | PartyCast: PartyCastWindow_TemplateButton, PartyCast: PartyCastWindow_Template_LargeButton, PartyCast: PartyCastWindow_Template_PlainButton, PartyCast: PartyCastWindow_Template_SmallButton, PartyCast: PartyCastWindow_Template_SpecialButton, TidyRoll: TRollItemButton |
| XML usage count | 6 |
| XML attribute usage count | 6 |
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

Observed XML element type instantiated by 2 addons.

## Common Attributes

- id
- texture
- x
- y

## Common Inherits

- none

## Common Parent Elements

- [TexSlices](element_TexSlices.md) — 5× (HIGH)

## Attribute Reference

| Attribute | Required | Usage % | Sample Values |
| --- | --- | --- | --- |
| `id` | **required** | 83% | morale-white, ability-white, tactic-black |
| `texture` | optional | 16% | TidyRoll_SquareFrame |
| `x` | optional | 16% | 0 |
| `y` | optional | 16% | 0 |
## Seen In

- PartyCast
- TidyRoll

## Examples

- PartyCast: PartyCastWindow_TemplateButton -> Normal in Button PartyCastWindow_TemplateButton
- PartyCast: PartyCastWindow_Template_LargeButton -> Normal in Button PartyCastWindow_Template_LargeButton
- PartyCast: PartyCastWindow_Template_PlainButton -> Normal in Button PartyCastWindow_Template_PlainButton
- PartyCast: PartyCastWindow_Template_SmallButton -> Normal in Button PartyCastWindow_Template_SmallButton
- PartyCast: PartyCastWindow_Template_SpecialButton -> Normal in Button PartyCastWindow_Template_SpecialButton
- TidyRoll: TRollItemButton -> Normal in Button TRollItemButton

## Related APIs

- [TexSlices](element_TexSlices.md) (MEDIUM 45/100) - XML Element Type

## Used With

- none

## Triggered By

- none

## Affects

- none
