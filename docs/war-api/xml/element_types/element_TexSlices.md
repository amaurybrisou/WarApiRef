# TexSlices

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
| Addons seen in | AggroMeter, AnywhereTrainer, BankArkel, BuffHead, CMap, EA_ThreePartBar, GuardList, GuardRange |
| Files seen in | `/workspace_addons/AggroMeter/AggroMeter.xml:39`, `/workspace_addons/AnywhereTrainer/source/AnywhereTrainer.xml:28`, `/workspace_addons/BankArkel/BankArkel.xml:124`, `/workspace_addons/BankArkel/BankArkel.xml:142`, `/workspace_addons/BankArkel/BankArkel.xml:160`, `/workspace_addons/BankArkel/BankArkel.xml:178`, `/workspace_addons/BankArkel/BankArkel.xml:196`, `/workspace_addons/BankArkel/BankArkel.xml:214` |
| Namespaces detected | TexSlices |
| Source kinds | xml_frames |
| Example locations | AggroMeter: Aggro_Button_Template, AnywhereTrainer: AnywhereTrainerTabTemplateInactiveImage, BankArkel: PackWinTab1, BankArkel: PackWinTab2, BankArkel: PackWinTab3, BankArkel: PackWinTab4 |
| XML usage count | 54 |
| XML attribute usage count | 54 |
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

Observed XML element type instantiated by 13 addons.

## Common Attributes

- none

## Common Inherits

- none

## Common Parent Elements

- [Button](element_Button.md)
- [FullResizeImage](element_FullResizeImage.md)
- [HorizontalResizeImage](element_HorizontalResizeImage.md)

## Seen In

- AggroMeter
- AnywhereTrainer
- BankArkel
- BuffHead
- CMap
- EA_ThreePartBar
- GuardList
- GuardRange
- Miracle Grow Remix
- PotionBar
- RVMOD_Manager
- RoR_SoR
- Shinies

## Examples

- AggroMeter: Aggro_Button_Template -> TexSlices in Button Aggro_Button_Template
- AnywhereTrainer: AnywhereTrainerTabTemplateInactiveImage -> TexSlices in Button AnywhereTrainerTabTemplateInactiveImage
- BankArkel: PackWinTab1 -> TexSlices in Button PackWinTab1
- BankArkel: PackWinTab2 -> TexSlices in Button PackWinTab2
- BankArkel: PackWinTab3 -> TexSlices in Button PackWinTab3
- BankArkel: PackWinTab4 -> TexSlices in Button PackWinTab4

## Related APIs

- none

## Used With

- none

## Triggered By

- none

## Affects

- none
