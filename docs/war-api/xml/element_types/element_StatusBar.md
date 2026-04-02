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
| Addons seen in | RoR_SoR, Targets |
| Files seen in | `/workspace_addons/RoR_SoR/RoR_SoR.xml:1138`, `/workspace_addons/RoR_SoR/RoR_SoR.xml:1156`, `/workspace_addons/RoR_SoR/RoR_SoR.xml:1337`, `/workspace_addons/RoR_SoR/RoR_SoR.xml:1352`, `/workspace_addons/targets/TargetsWindow.xml:42` |
| Namespaces detected | StatusBar |
| Source kinds | xml_frames |
| Example locations | RoR_SoR: RoR_SoR_FortTemplateLORDLEFTBAR, RoR_SoR: RoR_SoR_FortTemplateLORDRIGHTBAR, RoR_SoR: RoR_SoR_RealmTemplateVPDESTROBAR, RoR_SoR: RoR_SoR_RealmTemplateVPORDERBAR, Targets: TargetsUnitFrameHPBar |
| XML usage count | 5 |
| XML attribute usage count | 5 |
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

- name
- popable
- foreground
- handleinput
- reverseFill
- inherits
- interpolate
- layer

## Common Inherits

- UnitFrameFriendlyHealthStatusSmallBar

## Common Parent Elements

- [Window](element_Window.md)

## Attribute Reference

| Attribute | Required | Usage % | Sample Values |
| --- | --- | --- | --- |
| `popable` | optional | 50% | false |
| `foreground` | optional | 40% | T1_DefaultTintableBar, FORT_DefaultTintableBar |
| `handleinput` | optional | 40% | false |
| `reverseFill` | optional | 20% | true |
| `inherits` | optional | 10% | UnitFrameFriendlyHealthStatusSmallBar |
| `interpolate` | optional | 10% | true |
| `layer` | optional | 10% | background |
## Seen In

- RoR_SoR
- Targets

## Examples

- RoR_SoR: RoR_SoR_FortTemplateLORDLEFTBAR -> StatusBar RoR_SoR_FortTemplateLORDLEFTBAR
- RoR_SoR: RoR_SoR_FortTemplateLORDRIGHTBAR -> StatusBar RoR_SoR_FortTemplateLORDRIGHTBAR
- RoR_SoR: RoR_SoR_RealmTemplateVPDESTROBAR -> StatusBar RoR_SoR_RealmTemplateVPDESTROBAR
- RoR_SoR: RoR_SoR_RealmTemplateVPORDERBAR -> StatusBar RoR_SoR_RealmTemplateVPORDERBAR
- Targets: TargetsUnitFrameHPBar -> StatusBar TargetsUnitFrameHPBar

## Related APIs

- [LabelSetFont](../../window_api/functions/window_LabelSetFont.md) (HIGH 100/100) - Window Function
- [LabelSetTextAlign](../../window_api/functions/window_LabelSetTextAlign.md) (HIGH 100/100) - Window Function
- [WindowSetScale](../../window_api/functions/window_WindowSetScale.md) (HIGH 100/100) - Window Function

## Used With

- none

## Triggered By

- none

## Affects

- none
