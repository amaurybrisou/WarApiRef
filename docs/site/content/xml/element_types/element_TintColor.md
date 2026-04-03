# TintColor

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
| Addons seen in | AdvancedPetAssist, AggroMeter, Aura, AutoMark, BuffHead, GuardLine, Killer, LibGroup |
| Files seen in | `/workspace/data/raw/AdvancedPetAssist/APAGui.xml:0`, `/workspace/data/raw/AggroMeter/AggroMeter.xml:0`, `/workspace/data/raw/Aura/Source/AuraTooltip.xml:0`, `/workspace/data/raw/AutoMark/Source/AutoMark.xml:0`, `/workspace/data/raw/BuffHead/Setup/General.xml:0`, `/workspace/data/raw/BuffHead/Setup/SetupAdvancedContainersItem.xml:0`, `/workspace/data/raw/BuffHead/Setup/SetupContainer.xml:0`, `/workspace/data/raw/BuffHead/Setup/SetupLayout.xml:0` |
| Namespaces detected | TintColor |
| Source kinds | xml_frames |
| Example locations | AdvancedPetAssist: APAFollowTargetHUDFill, AdvancedPetAssist: APAInstantOnlyHUDFill, AdvancedPetAssist: APAKitingHUDFill, AdvancedPetAssist: APAPetTargetHUDBg, AggroMeter: AggroMeterWindow_AggroWindow1BorderCheck, AggroMeter: AggroMeterWindow_AggroWindow1Seperator1 |
| XML usage count | 83 |
| XML attribute usage count | 83 |
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

TintColor is a XML UI element. It commonly appears under CircleImage and DynamicImage.

## Common Attributes

- a
- b
- g
- r

## Common Inherits

- none

## Common Parent Elements

- [FullResizeImage](element_FullResizeImage.md) — 28× (HIGH)
- [DynamicImage](element_DynamicImage.md) — 25× (HIGH)
- [Label](element_Label.md) — 14× (HIGH)
- [HorizontalResizeImage](element_HorizontalResizeImage.md) — 7× (MEDIUM)
- [CircleImage](element_CircleImage.md) — 6× (MEDIUM)
- [VerticalResizeImage](element_VerticalResizeImage.md) — 3× (MEDIUM)

## Attribute Reference

| Attribute | Required | Usage % | Sample Values |
| --- | --- | --- | --- |
| `b` | **required** | 100% | 0, 110, 130, 50, ... |
| `g` | **required** | 100% | 0, 200, 185, 130, ... |
| `r` | **required** | 100% | 255, 0, 180, 200, ... |
| `a` | optional | 40% | 30, 255, 0.5, 200, ... |
## Seen In

- AdvancedPetAssist
- AggroMeter
- Aura
- AutoMark
- BuffHead
- GuardLine
- Killer
- LibGroup
- PartyCast
- Pocket Palette
- RoR_SoR
- Shinies
- TexturedButtons
- TurretRange

## Examples

- AdvancedPetAssist: APAFollowTargetHUDFill -> TintColor in HorizontalResizeImage APAFollowTargetHUDFill
- AdvancedPetAssist: APAInstantOnlyHUDFill -> TintColor in HorizontalResizeImage APAInstantOnlyHUDFill
- AdvancedPetAssist: APAKitingHUDFill -> TintColor in HorizontalResizeImage APAKitingHUDFill
- AdvancedPetAssist: APAPetTargetHUDBg -> TintColor in HorizontalResizeImage APAPetTargetHUDBg
- AggroMeter: AggroMeterWindow_AggroWindow1BorderCheck -> TintColor in FullResizeImage AggroMeterWindow_AggroWindow1BorderCheck
- AggroMeter: AggroMeterWindow_AggroWindow1Seperator1 -> TintColor in FullResizeImage AggroMeterWindow_AggroWindow1Seperator1

## Related APIs

- [CircleImage](element_CircleImage.md) (HIGH 100/100) - XML Element Type
- [DynamicImage](element_DynamicImage.md) (HIGH 100/100) - XML Element Type
- [FullResizeImage](element_FullResizeImage.md) (HIGH 100/100) - XML Element Type
- [HorizontalResizeImage](element_HorizontalResizeImage.md) (HIGH 100/100) - XML Element Type
- [Label](element_Label.md) (HIGH 100/100) - XML Element Type
- [VerticalResizeImage](element_VerticalResizeImage.md) (HIGH 90/100) - XML Element Type

## Used With

- none

## Triggered By

- none

## Affects

- none
