# EA_Button_Tab

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
| Addons seen in | AdvancedPetAssist, AdvancedRenownTrainer, AggroMeter, Atlas, Aura, AutoBand, DPSMeter, Deathblow |
| Files seen in | APAGui.xml, AdvancedRenownTrainingWindow.xml, AggroMeter.xml, AutoBandWindow.xml, DPSMeterWindow.xml, Deathblow.xml, Deathblow2.xml, Emojii.xml |
| Namespaces detected | EA_Button_Tab |
| Source kinds | xml_attributes |
| Example locations | APAOptionsTabsAutoRecall, APAOptionsTabsControls, APAOptionsTabsFollowTarget, APAOptionsTabsGeneral, APAOptionsTabsHUD, APAOptionsTabsKiting |
| XML usage count | 55 |
| XML attribute usage count | 55 |
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

Engine-supplied XML constant or template class referenced by 30 addons.

## Seen In

- AdvancedPetAssist
- AdvancedRenownTrainer
- AggroMeter
- Atlas
- Aura
- AutoBand
- DPSMeter
- Deathblow
- Deathblow2
- EA_OpenPartyWindow
- Emojii
- FozAuction
- GDes
- Ges
- HealGrid
- Kwestor
- LoyalPet
- Queue Queuer
- RVMOD_Manager
- RaidMeter
- RandomMount
- SocialWindow 2.0
- TastyButtons
- TidyChat
- TidyRoll
- TomeTracker
- WSCT
- WarBoard
- Warbuilder
- wbLeadHelper

## Used By

- APAOptionsTabsAutoRecall
- APAOptionsTabsControls
- APAOptionsTabsFollowTarget
- APAOptionsTabsGeneral
- APAOptionsTabsHUD
- APAOptionsTabsKiting
- APAOptionsTabsLos
- AdvancedRenownTrainingWindowTabsActiveTab
- AdvancedRenownTrainingWindowTabsAdvancedTab
- AdvancedRenownTrainingWindowTabsBasicTab
- AdvancedRenownTrainingWindowTabsDefTab
- AggroMeterGrayWindowBlackTab
- AggroMeterGrayWindowListTab
- AggroMeterGrayWindowWhiteTab
- AuctionWindowTabButton
- AuraTabButtonTemplate
- AutoBandWindowTabsConfig
- AutoBandWindowTabsTemplate
- AutoBandWindowTabsTools
- DPSMeterWindowTabsAbilityTabButton
- DPSMeterWindowTabsOverallTabButton
- DeathblowWin1WindowClassTab
- DeathblowWin1WindowProfileTab
- DeathblowWin1WindowSessionTab
- DeathblowWin1WindowSummTab
- DeathblowWin1WindowTotalTab
- EA_Button_OpenPartyTab
- EmojiiChooseIconDialogTab1
- EmojiiChooseIconDialogTab2
- EmojiiChooseIconDialogTab3
- GDesTabButton_Template
- GesTabButtonTemplate
- HGG_TabSpellTrackTabTemplate
- KwestorGui_TabTemplate
- LPETTabButton
- RVMOD_ManagerWindowTabGeneral
- RVMOD_ManagerWindowTabSettings
- RaidMeterWindowBossTab
- RaidMeterWindowTotalTab
- RandomMountWindowTabBlacklist
- RandomMountWindowTabMounts
- RandomMountWindowTabSettings
- SocialWindowBuddyListTabTemplate
- SocialWindowTabTemplate
- TAtlasTab
- TChatTabButton
- TRollTabButton
- TabButton
- TabButtonTemplate
- TastyButtonsOptionsWindowTabTemplate
- TomeTracker_JournalTabButton
- WSCTTabButton
- WarbuilderTabTemplate
- wbLeadHelperConfigWindowTabsConfig
- wbLeadHelperConfigWindowTabsMessages

## Related APIs

- [Button](../../xml/element_types/element_Button.md) (HIGH 100/100) - XML Element Type

## Notes

- none
