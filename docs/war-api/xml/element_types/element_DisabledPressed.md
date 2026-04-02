# DisabledPressed

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
| Addons seen in | AggroMeter, AnywhereTrainer, CMap, GuardList, GuardRange, PotionBar, RoR_SoR, Shinies |
| Files seen in | `/workspace_addons/AggroMeter/AggroMeter.xml:39`, `/workspace_addons/AnywhereTrainer/source/AnywhereTrainer.xml:28`, `/workspace_addons/GuardList/GuardList.xml:35`, `/workspace_addons/GuardRange/GuardRange.xml:35`, `/workspace_addons/PotionBar/settings/Settings.xml:22`, `/workspace_addons/PotionBar/settings/Settings.xml:45`, `/workspace_addons/RoR_SoR/RoR_SoR.xml:1598`, `/workspace_addons/RoR_SoR/RoR_SoR.xml:1612` |
| Namespaces detected | DisabledPressed |
| Source kinds | xml_frames |
| Example locations | AggroMeter: Aggro_Button_Template, AnywhereTrainer: AnywhereTrainerTabTemplateInactiveImage, CMap: Cmap_Template_OverheadMapZoomSliderInButton, CMap: Cmap_Template_OverheadMapZoomSliderOutButton, GuardList: GuardList_Window0Shield, GuardRange: GuardRange_Window0Shield |
| XML usage count | 14 |
| XML attribute usage count | 14 |
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

Observed XML element type instantiated by 8 addons.

## Common Attributes

- id
- a
- b
- g
- r

## Common Handlers

- none

## Common Inherits

- none

## Seen In

- AggroMeter
- AnywhereTrainer
- CMap
- GuardList
- GuardRange
- PotionBar
- RoR_SoR
- Shinies

## Examples

- AggroMeter: Aggro_Button_Template -> DisabledPressed in Button Aggro_Button_Template
- AnywhereTrainer: AnywhereTrainerTabTemplateInactiveImage -> DisabledPressed in Button AnywhereTrainerTabTemplateInactiveImage
- CMap: Cmap_Template_OverheadMapZoomSliderInButton -> DisabledPressed in Button Cmap_Template_OverheadMapZoomSliderInButton
- CMap: Cmap_Template_OverheadMapZoomSliderOutButton -> DisabledPressed in Button Cmap_Template_OverheadMapZoomSliderOutButton
- GuardList: GuardList_Window0Shield -> DisabledPressed in Button GuardList_Window0Shield
- GuardRange: GuardRange_Window0Shield -> DisabledPressed in Button GuardRange_Window0Shield

## Related APIs

- none

## Used With

- none

## Triggered By

- none

## Affects

- none
