# StatusBar

- Category: XML Element Type
- Confidence level: HIGH
- Confidence score: 100/100

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 123

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, used directly in xml handler attributes, matches a known engine namespace.

## Evidence Signals

- +18 Seen in 2 to 3 addons: Cross-addon spread is present but limited.
- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +30 Used directly in XML handler attributes: XML exposure suggests an engine-level contract.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | PartyCast, RoR_SoR |
| Files seen in | `/workspace/data/raw/PartyCast/PartyCast.xml:0`, `/workspace/data/raw/RoR_SoR/RoR_SoR.xml:0` |
| Namespaces detected | StatusBar |
| Source kinds | xml_frames |
| Example locations | PartyCast: PartyCastWindow_Template_LargeTimerBar, PartyCast: PartyCastWindow_Template_PlainTimerBar, PartyCast: PartyCastWindow_Template_SmallTimerBar, PartyCast: PartyCastWindow_Template_SpecialTimerBar, PartyCast: PartyFrameStatusBarBar, RoR_SoR: RoR_SoR_FortTemplateLORDLEFTBAR |
| XML usage count | 9 |
| XML attribute usage count | 9 |
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

Observed XML element type instantiated by 2 addons.

## Common Attributes

- handleinput
- name
- popable
- foreground
- inherits
- reverseFill
- background

## Common Inherits

- EA_StatusBar_DefaultTintable
- RRQTomeStatusBar

## Common Parent Elements

- [Windows](element_Windows.md) — 9× (HIGH)

## Common Template Bases

- EA_StatusBar_DefaultTintable
- RRQTomeStatusBar


> **Note**: This element type commonly acts as a template base.

## Attribute Reference

| Attribute | Required | Usage % | Sample Values |
| --- | --- | --- | --- |
| `handleinput` | **required** | 100% | false |
| `popable` | **required** | 100% | false |
| `foreground` | optional | 55% | LOL_BAR, FORT_DefaultTintableBar, T1_DefaultTintableBar |
| `inherits` | optional | 44% | EA_StatusBar_DefaultTintable, RRQTomeStatusBar |
| `reverseFill` | optional | 22% | true |
| `background` | optional | 11% | LOL_BAR |
## Seen In

- PartyCast
- RoR_SoR

## Examples

- PartyCast: PartyCastWindow_Template_LargeTimerBar -> StatusBar PartyCastWindow_Template_LargeTimerBar
- PartyCast: PartyCastWindow_Template_PlainTimerBar -> StatusBar PartyCastWindow_Template_PlainTimerBar
- PartyCast: PartyCastWindow_Template_SmallTimerBar -> StatusBar PartyCastWindow_Template_SmallTimerBar
- PartyCast: PartyCastWindow_Template_SpecialTimerBar -> StatusBar PartyCastWindow_Template_SpecialTimerBar
- PartyCast: PartyFrameStatusBarBar -> StatusBar PartyFrameStatusBarBar
- RoR_SoR: RoR_SoR_FortTemplateLORDLEFTBAR -> StatusBar RoR_SoR_FortTemplateLORDLEFTBAR

## Related APIs

- [Windows](element_Windows.md) (HIGH 100/100) - XML Element Type

## Used With

- none

## Triggered By

- none

## Affects

- none
