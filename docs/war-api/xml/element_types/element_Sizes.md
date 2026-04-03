# Sizes

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
| Addons seen in | AdvancedPetAssist, BuffHead, CM_ClosetGoblin, Enemy, PartyCast, RoR_SoR, Shinies |
| Files seen in | `/workspace/data/raw/AdvancedPetAssist/APAGui.xml:0`, `/workspace/data/raw/BuffHead/Setup/SetupLayout.xml:0`, `/workspace/data/raw/ClosetGoblin/ClosetGoblin.xml:0`, `/workspace/data/raw/Enemy/Code/Core/Common.xml:0`, `/workspace/data/raw/Enemy/Code/ScenarioInfo/ScenarioInfo.xml:0`, `/workspace/data/raw/PartyCast/PartyCast.xml:0`, `/workspace/data/raw/RoR_SoR/RoR_SoR.xml:0`, `/workspace/data/raw/Shinies/Modules/UI/Shinies-UI-Auctions.xml:0` |
| Namespaces detected | Sizes |
| Source kinds | xml_frames |
| Example locations | AdvancedPetAssist: APAFollowTargetHUDFill, AdvancedPetAssist: APAInstantOnlyHUDFill, AdvancedPetAssist: APAKitingHUDFill, AdvancedPetAssist: APAPetTargetHUDBg, BuffHead: BuffHeadLayoutHorizontalResizeImage, BuffHead: BuffHeadLayoutVerticalResizeImage |
| XML usage count | 23 |
| XML attribute usage count | 23 |
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

Observed XML element type instantiated by 7 addons.

## Common Attributes

- middle
- left
- right
- bottom
- top

## Common Inherits

- none

## Common Parent Elements

- [FullResizeImage](element_FullResizeImage.md) — 14× (HIGH)
- [HorizontalResizeImage](element_HorizontalResizeImage.md) — 8× (HIGH)
- [VerticalResizeImage](element_VerticalResizeImage.md) — 1× (LOW)

## Common Structural Child Elements

- Middle — 14× (HIGH)

## Attribute Reference

| Attribute | Required | Usage % | Sample Values |
| --- | --- | --- | --- |
| `middle` | optional | 39% | 30, 200, 61, 50 |
| `left` | optional | 34% | 0, 7 |
| `right` | optional | 34% | 0 |
| `bottom` | optional | 4% | 0 |
| `top` | optional | 4% | 0 |
## Structural Sub-Elements

### Middle

Observed 14 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `x` | **required** | 0, 424, 78, 2 |
| `y` | **required** | 0, 762, 25, 2 |
## Seen In

- AdvancedPetAssist
- BuffHead
- CM_ClosetGoblin
- Enemy
- PartyCast
- RoR_SoR
- Shinies

## Examples

- AdvancedPetAssist: APAFollowTargetHUDFill -> Sizes in HorizontalResizeImage APAFollowTargetHUDFill
- AdvancedPetAssist: APAInstantOnlyHUDFill -> Sizes in HorizontalResizeImage APAInstantOnlyHUDFill
- AdvancedPetAssist: APAKitingHUDFill -> Sizes in HorizontalResizeImage APAKitingHUDFill
- AdvancedPetAssist: APAPetTargetHUDBg -> Sizes in HorizontalResizeImage APAPetTargetHUDBg
- BuffHead: BuffHeadLayoutHorizontalResizeImage -> Sizes in HorizontalResizeImage BuffHeadLayoutHorizontalResizeImage
- BuffHead: BuffHeadLayoutVerticalResizeImage -> Sizes in VerticalResizeImage BuffHeadLayoutVerticalResizeImage

## Related APIs

- [FullResizeImage](element_FullResizeImage.md) (HIGH 100/100) - XML Element Type
- [HorizontalResizeImage](element_HorizontalResizeImage.md) (HIGH 100/100) - XML Element Type
- [VerticalResizeImage](element_VerticalResizeImage.md) (HIGH 90/100) - XML Element Type

## Used With

- none

## Triggered By

- none

## Affects

- none
