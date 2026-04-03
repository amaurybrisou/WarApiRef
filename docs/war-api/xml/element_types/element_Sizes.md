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
| Addons seen in | AdvancedPetAssist, BuffHead, CCTV, CM_ClosetGoblin, CMap, CleanCastbar, Crusher, Dascore |
| Files seen in | Code/Core/Common.xml, Code/ScenarioInfo/ScenarioInfo.xml, Configuration/Config.xml, Configuration/HopperConfig.xml, Configuration/WCDBConfig.xml, Configuration/WCDPConfig.xml, Modules/UI/Shinies-UI-Auctions.xml, Modules/UI/Shinies-UI-Auto.xml |
| Namespaces detected | Sizes |
| Source kinds | xml_frames |
| Example locations | AdvancedPetAssist: APAFollowTargetHUDFill, AdvancedPetAssist: APAInstantOnlyHUDFill, AdvancedPetAssist: APAKitingHUDFill, AdvancedPetAssist: APAPetTargetHUDBg, BuffHead: BuffHeadLayoutHorizontalResizeImage, BuffHead: BuffHeadLayoutVerticalResizeImage |
| XML usage count | 81 |
| XML attribute usage count | 81 |
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

Sizes is a structural XML sub-element. It commonly appears under FullResizeImage and HorizontalResizeImage. It is typically used to organize structural children such as BottomRight, Middle, TopLeft.

## Common Attributes

- bottom
- left
- middle
- right
- top

## Common Inherits

- none

## Common Parent Elements

- [HorizontalResizeImage](element_HorizontalResizeImage.md) — 43× (HIGH)
- [FullResizeImage](element_FullResizeImage.md) — 30× (HIGH)
- [VerticalResizeImage](element_VerticalResizeImage.md) — 8× (MEDIUM)

## Common Structural Child Elements

- [BottomRight](element_BottomRight.md) — 30× (HIGH)
- [Middle](element_Middle.md) — 30× (HIGH)
- [TopLeft](element_TopLeft.md) — 30× (HIGH)

## Attribute Reference

| Attribute | Required | Usage % | Sample Values |
| --- | --- | --- | --- |
| `middle` | optional | 62% | 30, 200, 90, 50, ... |
| `left` | optional | 54% | 0, 7, 13, 15, ... |
| `right` | optional | 54% | 0, 7, 13, 15, ... |
| `bottom` | optional | 6% | 0, 13, 3 |
| `top` | optional | 6% | 0, 13, 3 |
## Structural Sub-Elements

### [BottomRight](element_BottomRight.md)

Observed 30 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `x` | **required** | 10, 137, 16, 350 |
| `y` | **required** | 9, 609, 15, 197 |
| `id` | optional | Action-Bar-Frame-Bottom-Right, Border-Bottom-Right |
### [Middle](element_Middle.md)

Observed 30 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `x` | **required** | 0, 424, 50, 78 |
| `y` | **required** | 0, 762, 23, 25 |
| `id` | optional | Order-VP-Bar-horiz, Dest-VP-Bar-horiz, Victor-Bar-horiz, Loser-Bar-horiz |
### [TopLeft](element_TopLeft.md)

Observed 30 times as an unnamed child.

| Attribute | Required | Sample Values |
| --- | --- | --- |
| `x` | **required** | 10, 36, 16, 256 |
| `y` | **required** | 10, 571, 15, 157 |
| `id` | optional | Action-Bar-Frame-Top-Left, Border-Top-Left |
## Recursive Hierarchy

- Root: [Sizes](element_Sizes.md)
- [BottomRight](element_BottomRight.md) (structural, 30×, HIGH)
- [Middle](element_Middle.md) (structural, 30×, HIGH)
- [TopLeft](element_TopLeft.md) (structural, 30×, HIGH)

## Seen In

- AdvancedPetAssist
- BuffHead
- CCTV
- CM_ClosetGoblin
- CMap
- CleanCastbar
- Crusher
- Dascore
- EA_ThreePartBar
- EA_UiDebugTools
- Enemy
- FlagCap
- Hopper
- Map
- MapMonster
- MoraleSet
- Motion
- PartyCast
- Pure
- Pure Careerbar
- RVAPI_ColorDialog
- RVMOD_PlayerStatus
- RVMOD_SquaredDistances
- ResHelp
- RoR_SoR
- SNT_CASTBAR
- SNT_INFO
- Sequencer
- Shinies
- SpamBayes
- TwisterSet
- VerticalMorale
- WTes
- Wikki's Cooldown Bar
- Wikki's Cooldown Pulse
- zMailMod

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
- [VerticalResizeImage](element_VerticalResizeImage.md) (HIGH 100/100) - XML Element Type
- [BottomRight](element_BottomRight.md) (MEDIUM 45/100) - XML Element Type
- [Middle](element_Middle.md) (MEDIUM 45/100) - XML Element Type
- [TopLeft](element_TopLeft.md) (MEDIUM 45/100) - XML Element Type

## Used With

- [TexCoords](element_TexCoords.md) (HIGH 100/100) - XML Element Type
- [TexSlices](element_TexSlices.md) (HIGH 100/100) - XML Element Type
