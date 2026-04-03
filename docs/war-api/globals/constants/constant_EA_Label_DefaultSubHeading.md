# EA_Label_DefaultSubHeading

- Category: Constant
- Confidence level: HIGH
- Confidence score: 100/100

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 135

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, seen in 4 or more addons, used directly in xml handler attributes.

## Evidence Signals

- +30 Seen in 4 or more addons: Cross-addon spread is strong.
- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +30 Used directly in XML handler attributes: XML exposure suggests an engine-level contract.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | CMap, EA_OpenPartyWindow, EA_UiModWindow, Minmap, NerfedButtons, PeaceOut, PotionBar, RvRStats |
| Files seen in | CMap.xml, PeaceOut.xml, SOROptions.xml, Source/OpenPartyWindow.xml, Source/OpenPartyWindowCommon.xml, Source/OpenPartyWindowTabManage.xml, Source/OpenPartyWindowTabNearby.xml, Source/OpenPartyWindowTabWorld.xml |
| Namespaces detected | EA_Label_DefaultSubHeading |
| Source kinds | xml_attributes |
| Example locations | $parentRightPaneTitle, CMapWindowAreaNameText, CMapWindowDigitinf, CharacterWindowRvRStatsHeader, CharacterWindowRvRStatsTitle, EA_Template_OpenPartyGroupLineLocationTypeText |
| XML usage count | 21 |
| XML attribute usage count | 21 |
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

Engine-supplied XML constant or template class referenced by 10 addons.

## Seen In

- CMap
- EA_OpenPartyWindow
- EA_UiModWindow
- Minmap
- NerfedButtons
- PeaceOut
- PotionBar
- RvRStats
- RvRStatsTab
- SOR

## Used By

- $parentRightPaneTitle
- CMapWindowAreaNameText
- CMapWindowDigitinf
- CharacterWindowRvRStatsHeader
- CharacterWindowRvRStatsTitle
- EA_Template_OpenPartyGroupLineLocationTypeText
- EA_Template_OpenPartyWorldGroupLineInterestText
- EA_Tooltip_OpenPartyPartyTypeText
- EA_Tooltip_OpenPartyWorldPartyTypeText
- EA_Window_OpenPartyFlyOutHeader
- EA_Window_OpenPartyManageSocketDefLootHeaderText
- EA_Window_OpenPartyManageSocketDefWarbandHeaderText
- EA_Window_OpenPartyNearbySocketDefNoOpenPartiesHeader
- EA_Window_OpenPartyWorldSocketDefNoResultsHeader
- MinmapNameText
- NBSetup_SaveCreateFromInstructions
- NBSetup_SaveInstructions
- PeaceOutTitle
- SOR.Options.Title
- SOR.Version.Title
- UiModVersionMismatchWindowCurrentVersionText

## Related APIs

- [Label](../../xml/element_types/element_Label.md) (HIGH 100/100) - XML Element Type

## Notes

- none
