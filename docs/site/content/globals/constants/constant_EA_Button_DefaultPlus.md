# EA_Button_DefaultPlus

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
| Addons seen in | AutoBand, Busted, CraftingWillard, Deathblow2, DuffTimer, EA_UiDebugTools, Miracle Grow Remix, ObjectInspector |
| Files seen in | AutoBandWindowConfig.xml, AutoBandWindowTemplate.xml, Busted.xml, CraftingWillard.xml, Deathblow2.xml, DuffTimerOptionsDefn.xml, ObjectInspector.xml, PartyAdWindow.xml |
| Namespaces detected | EA_Button_DefaultPlus |
| Source kinds | xml_attributes |
| Example locations | AbilityLinkTemplatePlus, AutoBandWindowConfigButtonPlusHealerRank, AutoBandWindowConfigButtonPlusMaxMdps, AutoBandWindowConfigButtonPlusMaxRdps, AutoBandWindowConfigButtonPlusRank, AutoBandWindowConfigButtonPlusTime |
| XML usage count | 24 |
| XML attribute usage count | 24 |
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

Engine-supplied XML constant or template class referenced by 13 addons.

## Seen In

- AutoBand
- Busted
- CraftingWillard
- Deathblow2
- DuffTimer
- EA_UiDebugTools
- Miracle Grow Remix
- ObjectInspector
- PartyAd
- TastyButtons
- WarBoard
- Warbuilder
- ZonePOP

## Used By

- AbilityLinkTemplatePlus
- AutoBandWindowConfigButtonPlusHealerRank
- AutoBandWindowConfigButtonPlusMaxMdps
- AutoBandWindowConfigButtonPlusMaxRdps
- AutoBandWindowConfigButtonPlusRank
- AutoBandWindowConfigButtonPlusTime
- AutoBandWindowTemplateButtonPlusDps
- AutoBandWindowTemplateButtonPlusHealers
- AutoBandWindowTemplateButtonPlusTanks
- BustedGUINextError
- CraftingWillardPlus
- DuffTimerOptionsAdvHeaderPlus
- MiracleGrow2LayoutProgTemplateBarPlus
- MiracleGrow2LayoutSettingsTemplateVisPlus
- NemesisButton
- ObjectInspectorDepthPlusButton
- PartyAdWindowDpsPlusButton
- PartyAdWindowHealerPlusButton
- PartyAdWindowTankPlusButton
- PlusMinusMaxButton
- PlusMinusVertMaxButton
- TastyButtonsButtonSelectWindowButtonListViewButtonAddSelectedButton_All
- WarbuilderMainWindowPlus
- ZonePOPWndAddButton

## Related APIs

- [Button](../../xml/element_types/element_Button.md) (HIGH 100/100) - XML Element Type

## Notes

- none
