# EA_Label_DefaultText_Small

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
| Addons seen in | AutoBand, EA_OpenPartyWindow, FastFriends, LoyalPet, MegaphonePlusPlus, Pocket Palette, WSCT, nLootLink |
| Files seen in | AutoBandWindowConfig.xml, FastFriends_ConfigUI.xml, MegaphonePlusPlus.xml, PocketPalette.xml, Source/OpenPartyWindow.xml, Source/OpenPartyWindowCommon.xml, Source/OpenPartyWindowTabNearby.xml, Source/OpenPartyWindowTabWorld.xml |
| Namespaces detected | EA_Label_DefaultText_Small |
| Source kinds | xml_attributes |
| Example locations | AutoBandWindowConfigBackfillExperimentalLabel, AutoBandWindowConfigCommonRaceNamesLabel, AutoBandWindowConfigVersionLabel, DyeListBoxRowTemplateCount, DyeListBoxRowTemplateName, DyeWindowSelectedDyeCount |
| XML usage count | 60 |
| XML attribute usage count | 60 |
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

Engine-supplied XML constant or template class referenced by 8 addons.

## Seen In

- AutoBand
- EA_OpenPartyWindow
- FastFriends
- LoyalPet
- MegaphonePlusPlus
- Pocket Palette
- WSCT
- nLootLink

## Used By

- AutoBandWindowConfigBackfillExperimentalLabel
- AutoBandWindowConfigCommonRaceNamesLabel
- AutoBandWindowConfigVersionLabel
- DyeListBoxRowTemplateCount
- DyeListBoxRowTemplateName
- DyeWindowSelectedDyeCount
- DyeWindowSelectedDyeName
- EA_Template_OpenPartyGroupLineLeaderNameText
- EA_Template_OpenPartyWorldGroupLineLeaderNameText
- EA_Template_OpenParty_Label_Small
- EA_Tooltip_OpenPartyLocationText
- EA_Tooltip_OpenPartyLocationTypeText
- EA_Tooltip_OpenPartyNoteText
- EA_Tooltip_OpenPartyObjectiveText
- EA_Tooltip_OpenPartyWbText
- EA_Tooltip_OpenPartyWorldLocationText
- EA_Tooltip_OpenPartyWorldLocationTypeText
- EA_Tooltip_OpenPartyWorldNoteText
- EA_Tooltip_OpenPartyWorldSpecificInterest1
- EA_Tooltip_OpenPartyWorldSpecificInterest2
- EA_Tooltip_OpenPartyWorldSpecificInterest3
- EA_Tooltip_OpenPartyWorldSpecificInterest4
- EA_Tooltip_OpenPartyWorldWbText
- EA_Window_OpenPartyFlyOutLeader1RatioText
- EA_Window_OpenPartyFlyOutLeader1Text
- EA_Window_OpenPartyFlyOutLeader2RatioText
- EA_Window_OpenPartyFlyOutLeader2Text
- EA_Window_OpenPartyFlyOutLeader3RatioText
- EA_Window_OpenPartyFlyOutLeader3Text
- FastFriendsConfigMainCharOverrideDefaultLabel
- FastFriendsConfigMainCharOverrideOffLabel
- FastFriendsConfigMainCharOverrideOnLabel
- FastFriendsConfigMainListFriendsLabel
- FastFriendsConfigMainListIgnoreLabel
- FastFriendsConfigMainModeAllLabel
- FastFriendsConfigMainModeLevelLabel
- FastFriendsConfigMainModeOptInLabel
- ItemWindowGuide
- LPETComboBoxTemplateLabel
- MegaphoneMainFontTitleLabel
- MegaphoneMainHighlightLeaderLabel
- MegaphoneMainMaxLengthLabel
- MegaphoneMainSFXTitleLabel
- MegaphoneMainShowLeaderLabel
- PPMainIntroText
- PPMainSaveSettingsLabel
- WSCTComboBoxTemplateLabel
- WSCTOptionsColorPickerWindowCustomColorBlueText
- WSCTOptionsColorPickerWindowCustomColorGreenText
- WSCTOptionsColorPickerWindowCustomColorRedText
- WSCTSliderTemplateLabel
- WSCTSliderTemplateValue
- nLootLinkGUICount
- nLootLinkGUIVersion
- nLootLinkOptionsLabelVersion
- nLootLinkRowTemplateItemName
- nLootLinkRowTemplateItemRank
- nLootLinkRowTemplateItemRenown
- nLootLinkRowTemplateItemSubType
- nLootLinkRowTemplateItemType

## Related APIs

- [Label](../../xml/element_types/element_Label.md) (HIGH 100/100) - XML Element Type

## Notes

- none
