# Middle

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
| Addons seen in | AdvancedPetAssist, BuffHead, CM_ClosetGoblin, CMap, EA_ThreePartBar, EA_UiDebugTools, Enemy, MapMonster |
| Files seen in | `/workspace_addons/AdvancedPetAssist/APAGui.xml:1389`, `/workspace_addons/AdvancedPetAssist/APAGui.xml:1430`, `/workspace_addons/AdvancedPetAssist/APAGui.xml:1471`, `/workspace_addons/AdvancedPetAssist/APAGui.xml:1509`, `/workspace_addons/BuffHead/Setup/SetupLayout.xml:67`, `/workspace_addons/BuffHead/Setup/SetupLayout.xml:95`, `/workspace_addons/ClosetGoblin/ClosetGoblin.xml:70`, `/workspace_addons/Enemy/Code/Core/Common.xml:64` |
| Namespaces detected | Middle |
| Source kinds | xml_frames |
| Example locations | AdvancedPetAssist: APAFollowTargetHUDFill, AdvancedPetAssist: APAInstantOnlyHUDFill, AdvancedPetAssist: APAKitingHUDFill, AdvancedPetAssist: APAPetTargetHUDBg, BuffHead: BuffHeadLayoutHorizontalResizeImage, BuffHead: BuffHeadLayoutVerticalResizeImage |
| XML usage count | 30 |
| XML attribute usage count | 30 |
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

Observed XML element type instantiated by 11 addons.

## Common Attributes

- x
- y
- id

## Common Handlers

- none

## Common Inherits

- none

## Seen In

- AdvancedPetAssist
- BuffHead
- CM_ClosetGoblin
- CMap
- EA_ThreePartBar
- EA_UiDebugTools
- Enemy
- MapMonster
- RVAPI_ColorDialog
- RoR_SoR
- Shinies

## Examples

- AdvancedPetAssist: APAFollowTargetHUDFill -> Middle in HorizontalResizeImage APAFollowTargetHUDFill
- AdvancedPetAssist: APAInstantOnlyHUDFill -> Middle in HorizontalResizeImage APAInstantOnlyHUDFill
- AdvancedPetAssist: APAKitingHUDFill -> Middle in HorizontalResizeImage APAKitingHUDFill
- AdvancedPetAssist: APAPetTargetHUDBg -> Middle in HorizontalResizeImage APAPetTargetHUDBg
- BuffHead: BuffHeadLayoutHorizontalResizeImage -> Middle in HorizontalResizeImage BuffHeadLayoutHorizontalResizeImage
- BuffHead: BuffHeadLayoutVerticalResizeImage -> Middle in VerticalResizeImage BuffHeadLayoutVerticalResizeImage

## Related APIs

- none

## Used With

- none

## Triggered By

- none

## Affects

- none
