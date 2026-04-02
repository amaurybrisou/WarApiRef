# Left

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
| Addons seen in | AdvancedPetAssist, EA_ThreePartBar, EA_UiDebugTools, RVAPI_ColorDialog, RoR_SoR |
| Files seen in | `/workspace_addons/AdvancedPetAssist/APAGui.xml:1389`, `/workspace_addons/AdvancedPetAssist/APAGui.xml:1430`, `/workspace_addons/AdvancedPetAssist/APAGui.xml:1471`, `/workspace_addons/AdvancedPetAssist/APAGui.xml:1509`, `/workspace_addons/RVAPI_ColorDialog/RVAPI_ColorDialog.xml:70`, `/workspace_addons/RoR_SoR/RoR_SoR.xml:226`, `/workspace_addons/RoR_SoR/RoR_SoR.xml:238`, `/workspace_addons/ea_uidebugtools/Source/DebugWindowVerticalScroll.xml:123` |
| Namespaces detected | Left |
| Source kinds | xml_frames |
| Example locations | AdvancedPetAssist: APAFollowTargetHUDFill, AdvancedPetAssist: APAInstantOnlyHUDFill, AdvancedPetAssist: APAKitingHUDFill, AdvancedPetAssist: APAPetTargetHUDBg, EA_ThreePartBar: RewardPoolFilledBarLoser, EA_ThreePartBar: VictoryPointsFilledBarDestruction |
| XML usage count | 10 |
| XML attribute usage count | 10 |
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

- x
- y
- id

## Common Inherits

- none

## Common Parent Elements

- [HorizontalResizeImage](element_HorizontalResizeImage.md) — 10× (HIGH)

## Attribute Reference

| Attribute | Required | Usage % | Sample Values |
| --- | --- | --- | --- |
| `x` | optional | 80% |  |
| `y` | optional | 80% |  |
| `id` | optional | 20% |  |
## Seen In

- AdvancedPetAssist
- EA_ThreePartBar
- EA_UiDebugTools
- RVAPI_ColorDialog
- RoR_SoR

## Examples

- AdvancedPetAssist: APAFollowTargetHUDFill -> Left in HorizontalResizeImage APAFollowTargetHUDFill
- AdvancedPetAssist: APAInstantOnlyHUDFill -> Left in HorizontalResizeImage APAInstantOnlyHUDFill
- AdvancedPetAssist: APAKitingHUDFill -> Left in HorizontalResizeImage APAKitingHUDFill
- AdvancedPetAssist: APAPetTargetHUDBg -> Left in HorizontalResizeImage APAPetTargetHUDBg
- EA_ThreePartBar: RewardPoolFilledBarLoser -> Left in HorizontalResizeImage RewardPoolFilledBarLoser
- EA_ThreePartBar: VictoryPointsFilledBarDestruction -> Left in HorizontalResizeImage VictoryPointsFilledBarDestruction

## Related APIs

- none

## Used With

- none

## Triggered By

- none

## Affects

- none
