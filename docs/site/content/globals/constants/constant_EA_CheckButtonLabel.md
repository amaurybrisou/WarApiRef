# EA_CheckButtonLabel

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
| Addons seen in | ArmorGraphicToggle, AutoBand, DetauntHelper, EA_UiDebugTools, HealGrid, NerfedButtons, PartyAd, Phantom |
| Files seen in | AGTSettingsWindow.xml, AutoBandWindowConfig.xml, AutoBandWindowTemplate.xml, AutoBandWindowTools.xml, Gui/HealGridGuiTemplates.xml, PartyAdWindow.xml, Phantom.xml, RememberIt.xml |
| Namespaces detected | EA_CheckButtonLabel |
| Source kinds | xml_attributes |
| Example locations | AGTSettingsWindowCloakLabel, AGTSettingsWindowDescLabel, AGTSettingsWindowHelmLabel, AGTSettingsWindowHeraldryLabel, AutoBandWindowConfigAutoKickTimeoutLabel, AutoBandWindowConfigDefaultTemplateLabel |
| XML usage count | 77 |
| XML attribute usage count | 77 |
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

- ArmorGraphicToggle
- AutoBand
- DetauntHelper
- EA_UiDebugTools
- HealGrid
- NerfedButtons
- PartyAd
- Phantom
- RememberIt
- nLootLink

## Used By

- AGTSettingsWindowCloakLabel
- AGTSettingsWindowDescLabel
- AGTSettingsWindowHelmLabel
- AGTSettingsWindowHeraldryLabel
- AutoBandWindowConfigAutoKickTimeoutLabel
- AutoBandWindowConfigDefaultTemplateLabel
- AutoBandWindowConfigMaxMdpsValueLabel
- AutoBandWindowConfigMaxRdpsValueLabel
- AutoBandWindowConfigMinHealerRankValueLabel
- AutoBandWindowConfigMinRankValueLabel
- AutoBandWindowTemplateDefaultLabel
- AutoBandWindowTemplateDpsValueLabel
- AutoBandWindowTemplateHealersValueLabel
- AutoBandWindowTemplateRightClickTemplateMenuLabel
- AutoBandWindowTemplateTanksValueLabel
- AutoBandWindowToolsAutoFormSearchLabel
- AutoBandWindowToolsAutoPartyNoteLabel
- AutoBandWindowToolsAutoSearchLabel
- AutoBandWindowToolsDiscordReqLabel
- AutoBandWindowToolsKickOfflineLabel
- AutoBandWindowToolsKickRankLabel
- AutoBandWindowToolsKickZoneLabel
- AutoBandWindowToolsNotifyBuffsLabel
- AutoBandWindowToolsPrintRoleLabel
- AutoBandWindowToolsRightClickOrganizeLabel
- AutoBandWindowToolsSearchRoleChan1Label
- AutoBandWindowToolsSearchRoleChan5Label
- AutoBandWindowToolsSearchRoleChanT4Label
- DTC_AbilityRowTemplateLabel
- DTC_BarSettingsTemplateBarFontsFontColorLabel
- DTC_BarSettingsTemplateBarFontsFontStyleLabel
- DTC_BarSettingsTemplateBarSizeHeightLabel
- DTC_BarSettingsTemplateBarSizeScaleLabel
- DTC_BarSettingsTemplateBarSizeWidthLabel
- DTC_BarSettingsTemplateMiscellaneousBQEditBoxLabel
- DTC_BarSettingsTemplateMiscellaneousDelayEditBoxLabel
- DTC_BarSettingsTemplateMiscellaneousMREditBoxLabel
- DTC_BarSettingsTemplateMouseClearOnLabel
- DTC_BarSettingsTemplateMouseMonitorOnLabel
- DTC_BarSettingsTemplateMouseTargetOnLabel
- DTC_BarSettingsTemplateSoundNotificationsOnDetauntLabel
- DTC_BarSettingsTemplateSoundNotificationsOnDetauntReadyLabel
- EA_LabelCheckButtonSmallCopyLabel
- HGG_LabelCheckButtonLabel
- NBSBCoreWindowCriteriaAddon1Label
- NBSBCoreWindowCriteriaSequenceLabel
- NBSB_ChecksRowTemplateLabel
- NBSB_ParamRowTemplateLabel
- PartyAdWindowDpsValue
- PartyAdWindowHealerValue
- PartyAdWindowTankValue
- PhantomSettingsHideBarLockLabel
- PhantomSettingsHideGroupLabel
- PhantomSettingsHideMainAssistLabel
- PhantomSettingsHideMapAreaLabel
- PhantomSettingsHideMapCityLabel
- PhantomSettingsHideMapFrameLabel
- PhantomSettingsHideMapMailLabel
- PhantomSettingsHideMapPinsLabel
- PhantomSettingsHideMapRallyLabel
- PhantomSettingsHideMapScenLabel
- PhantomSettingsHideMapWorldLabel
- PhantomSettingsHideMapZoomLabel
- PhantomSettingsHidePetLabel
- PhantomSettingsHidePlayerLabel
- PhantomSettingsHideSocialLabel
- PhantomSettingsHideTargetLabel
- PhantomSettingsInstructions2Label
- PhantomSettingsInstructionsLabel
- RememberItSettingsDeathListLabel
- RememberItSettingsKillListLabel
- RememberItSettingsOnDeathLabel
- RememberItSettingsOnKillLabel
- RememberItSettingsOnRankUpLabel
- RememberItSettingsOnRenownRankUpLabel
- RememberItSettingsOnScenarioEndLabel
- nLootLinkOptionsCheckRarityConfirmLabel

## Related APIs

- [Label](../../xml/element_types/element_Label.md) (HIGH 100/100) - XML Element Type

## Notes

- none
