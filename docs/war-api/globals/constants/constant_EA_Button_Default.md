# EA_Button_Default

- Category: Constant
- Confidence level: HIGH
- Confidence score: 100/100

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 170

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, seen in 4 or more addons, used directly in xml handler attributes.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +30 Used directly in XML handler attributes: XML exposure suggests an engine-level contract.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +10 Referenced from initialization flow: Lifecycle reconstruction references this symbol.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +25 Referenced by generated docs or reference files: The symbol is reinforced outside a single call page.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | AggroMeter, BuffHead, CM_ClosetGoblin, CMap, GuardList, GuardRange, MapMonster, Miracle Grow Remix |
| Files seen in | `/workspace/AggroMeter/AggroMeter.xml:39`, `/workspace/BuffHead/Setup/SetupLayout.xml:4`, `/workspace/ClosetGoblin/ClosetGoblin.xml:33`, `/workspace/GuardList/GuardList.xml:35`, `/workspace/GuardRange/GuardRange.xml:35`, `/workspace/MGRemix/MGRemix.xml:105`, `/workspace/MGRemix/MGRemix.xml:3`, `/workspace/MGRemix/MGRemix.xml:70` |
| Namespaces detected | EA_Button_Default |
| Source kinds | flows, xml_attributes |
| Example locations | Aggro_Button_Template, BuffHeadLayoutResizeButton, CMapWindowFilterMenuButton, CMapWindowOverheadCurrentEventsButton, ClosetGoblinDefaultButton, GuardList_Window0Shield |
| XML usage count | 34 |
| XML attribute usage count | 34 |
| Lua usage count | 0 |
| Global usage count | 0 |
| Local definition count | 0 |
| Documentation references | 1 |
| Initialization flow references | 3 |
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

Observed engine XML template or inherited constant referenced by 16 addons.

## Seen In

- AggroMeter
- BuffHead
- CM_ClosetGoblin
- CMap
- GuardList
- GuardRange
- MapMonster
- Miracle Grow Remix
- MiracleGrow
- MiracleGrowLight
- Pocket Palette
- PotionBar
- RoR_SoR
- Shinies
- Tortall_DPS
- nRarity

## Used By

- Aggro_Button_Template
- BuffHeadLayoutResizeButton
- CMapWindowFilterMenuButton
- CMapWindowOverheadCurrentEventsButton
- ClosetGoblinDefaultButton
- GuardList_Window0Shield
- GuardRange_Window0Shield
- ItemWindowSlotButton
- MapMonsterIconChooserWindow_RowTemplate
- MiracleGrow2Button
- MiracleGrow2HarvestRptButton
- MiracleGrow2Repeat
- MiracleGrowButton
- MiracleGrowLightButton
- PotionBarButtonTemplate
- RoR_SoR_RealmTemplateCLAIM_WINDOW1BUTTON
- RoR_SoR_RealmTemplateCLAIM_WINDOW1BUTTON2
- RoR_SoR_RealmTemplateCLAIM_WINDOW2BUTTON
- RoR_SoR_RealmTemplateCLAIM_WINDOW2BUTTON2
- Shinies_IconButton
- Shinies_IconButton_Overlay
- TortallDPSDetailTemplate
- TortallDPSGroupTemplate
- TortallDPSMeterAbilities
- TortallDPSMeterDamage
- TortallDPSMeterGroup
- TortallDPSMeterHealing
- TortallDPSMeterReset
- TortallDPSMeterScale
- TortallDPSMeterTargets
- TortallDPSMeterWhatHM
- TortallDPSMeterWhoHM
- TortallDPSToggle
- nRarityBorderButton

## Related APIs

- none

## Used With

- none

## Triggered By

- none

## Affects

- none

## Notes

- none
