# EA_Window_TabSeparatorLeftSide

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
| Addons seen in | AdvancedRenownTrainer, AggroMeter, Aura, AutoBand, DPSMeter, Deathblow, Deathblow2, EA_OpenPartyWindow |
| Files seen in | AdvancedRenownTrainingWindow.xml, AggroMeter.xml, AutoBandWindow.xml, DPSMeterWindow.xml, Deathblow.xml, Deathblow2.xml, GDesOptions.xml, GesOptions.xml |
| Namespaces detected | EA_Window_TabSeparatorLeftSide |
| Source kinds | xml_attributes |
| Example locations | $parentTabsTabSeparatorLeft, AdvancedRenownTrainingWindowTabsTabSeparatorLeft, AggroMeterGrayWindowTabSeparatorLeft, AuctionWindowTabsSeparatorLeft, AuraConfigTabsSeparatorLeft, AutoBandWindowTabsSeparatorLeft |
| XML usage count | 26 |
| XML attribute usage count | 26 |
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

Engine-supplied XML constant or template class referenced by 25 addons.

## Seen In

- AdvancedRenownTrainer
- AggroMeter
- Aura
- AutoBand
- DPSMeter
- Deathblow
- Deathblow2
- EA_OpenPartyWindow
- FozAuction
- GDes
- Ges
- HealGrid
- Kwestor
- LoyalPet
- Queue Queuer
- RVMOD_Manager
- RaidMeter
- SocialWindow 2.0
- TastyButtons
- TomeTracker
- WSCT
- WarBoard
- Warbuilder
- XpStatus+G
- wbLeadHelper

## Used By

- $parentTabsTabSeparatorLeft
- AdvancedRenownTrainingWindowTabsTabSeparatorLeft
- AggroMeterGrayWindowTabSeparatorLeft
- AuctionWindowTabsSeparatorLeft
- AuraConfigTabsSeparatorLeft
- AutoBandWindowTabsSeparatorLeft
- DPSMeterWindowTabsTabSeperatorLeft
- DeathblowWin1WindowTabSeparatorLeft
- EA_Window_OpenPartyTabSeparatorLeft
- GDesOptionsTabsSeparatorLeft
- GesOptionsTabsSeparatorLeft
- HGG_TabSpellTrackWindowTemplateTabSeparatorLeft
- KwestorGui_KwestorTabSeparatorLeft
- LPETOptionsTabTabSeparatorLeft
- QueueQueuer_GUI_TabsSeparatorLeft
- RVMOD_ManagerWindowTabSeparatorLeft
- RaidMeterWindowTabSeparatorLeft
- SocialWindowBuddyListTabSeparatorLeft
- SocialWindowTabSeparatorLeft
- TastyButtonsButtonSelectWindowTabsSeparatorLeft
- TastyButtonsOptionsWindowTabsSeparatorLeft
- TomeTracker_JournalWindowTabTabSeparatorLeft
- WSCTOptionsTabTabSeparatorLeft
- WarBoardOptionsTabsSeparatorLeft
- WarbuilderMainWindowTabSeparatorLeft
- wbLeadHelperConfigWindowTabsSeparatorLeft

## Related APIs

- [Window](../../xml/element_types/element_Window.md) (HIGH 100/100) - XML Element Type

## Notes

- none
