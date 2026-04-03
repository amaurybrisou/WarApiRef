# EA_Button_DefaultListSort

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
| Addons seen in | Aura, EA_OpenPartyWindow, Enemy, FozAuction, PieTracker, QuickTacticSwitch, TaxPayer, nLootLink |
| Files seen in | Code/CombatLog/CombatLogStatsWindow.xml, Code/ScenarioInfo/ScenarioInfo.xml, PieTracker.xml, Source/AuraSettings.xml, Source/AuraShares.xml, Source/OpenPartyWindowTabNearby.xml, Source/OpenPartyWindowTabWorld.xml, Source/Templates.xml |
| Namespaces detected | EA_Button_DefaultListSort |
| Source kinds | xml_attributes |
| Example locations | AuctionWindowSortButton, AuraSharesSortButton, AuraSortingHeaderTemplate, AuraWindowSortButton, EA_Template_OpenPartyNearbySortButton, EA_Window_OpenPartyWorldSocketDefInterestSortButton |
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

Engine-supplied XML constant or template class referenced by 8 addons.

## Seen In

- Aura
- EA_OpenPartyWindow
- Enemy
- FozAuction
- PieTracker
- QuickTacticSwitch
- TaxPayer
- nLootLink

## Used By

- AuctionWindowSortButton
- AuraSharesSortButton
- AuraSortingHeaderTemplate
- AuraWindowSortButton
- EA_Template_OpenPartyNearbySortButton
- EA_Window_OpenPartyWorldSocketDefInterestSortButton
- EA_Window_OpenPartyWorldSocketDefLocationSortButton
- EA_Window_OpenPartyWorldSocketDefNumPlayersSortButton
- EA_Window_OpenPartyWorldSocketDefPartyLeaderSortButton
- EnemyCombatLogStatsWindow_ListHeaderTemplateSortColumn1
- EnemyCombatLogStatsWindow_ListHeaderTemplateSortColumn2
- EnemyCombatLogStatsWindow_ListHeaderTemplateSortColumn3
- EnemyCombatLogStatsWindow_ListHeaderTemplateSortColumn4
- EnemyCombatLogStatsWindow_ListHeaderTemplateSortColumn5
- EnemyCombatLogStatsWindow_ListHeaderTemplateSortColumn6
- EnemyScenarioInfoDialog_PlayerStatsHeaderTemplateSortArchetype2
- EnemyScenarioInfoDialog_PlayerStatsHeaderTemplateSortLevel
- EnemyScenarioInfoDialog_PlayerStatsHeaderTemplateSortName
- EnemyScenarioInfoDialog_PlayerStatsHeaderTemplateSortValue1
- EnemyScenarioInfoDialog_PlayerStatsHeaderTemplateSortValue2
- PieTrackerSortButtonTemplate
- QuickTacticSwitchWindowSortButton
- TaxPayerSortButtonTemplate
- nLootLinkSortingHeaderTemplate

## Related APIs

- [Button](../../xml/element_types/element_Button.md) (HIGH 100/100) - XML Element Type

## Notes

- none
