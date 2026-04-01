# EA_Window_TabSeparatorRightSide

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
| Addons seen in | AdvancedRenownTrainer, AggroMeter, Aura, AutoBand, LoyalPet, Queue Queuer, RVMOD_Manager, WSCT |
| Files seen in | `/workspace/AggroMeter/AggroMeter.xml:121`, `/workspace/Aura/Source/AuraConfig.xml:420`, `/workspace/Autoband/AutoBandWindow.xml:132`, `/workspace/LoyalPet/gui/lpet_gui.xml:219`, `/workspace/QueueQueuer/QueueQueuer_GUI.xml:264`, `/workspace/RVMOD_Manager/RVMOD_ManagerWindow.xml:174`, `/workspace/WarBoard/WarBoardOptions.xml:185`, `/workspace/advancedrenowntrainer/AdvancedRenownTrainingWindow.xml:66` |
| Namespaces detected | EA_Window_TabSeparatorRightSide |
| Source kinds | xml_attributes |
| Example locations | AdvancedRenownTrainingWindowTabsTabSeparatorRight, AggroMeterGrayWindowTabSeparatorRight, AuraConfigTabsSeparatorRight, AutoBandWindowTabsSeparatorRight, LPETOptionsTabTabSeparatorRight, QueueQueuer_GUI_TabsSeparatorRight |
| XML usage count | 10 |
| XML attribute usage count | 10 |
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

Observed engine XML template or inherited constant referenced by 10 addons.

## Seen In

- AdvancedRenownTrainer
- AggroMeter
- Aura
- AutoBand
- LoyalPet
- Queue Queuer
- RVMOD_Manager
- WSCT
- WarBoard
- wbLeadHelper

## Used By

- AdvancedRenownTrainingWindowTabsTabSeparatorRight
- AggroMeterGrayWindowTabSeparatorRight
- AuraConfigTabsSeparatorRight
- AutoBandWindowTabsSeparatorRight
- LPETOptionsTabTabSeparatorRight
- QueueQueuer_GUI_TabsSeparatorRight
- RVMOD_ManagerWindowTabSeparatorRight
- WSCTOptionsTabTabSeparatorRight
- WarBoardOptionsTabsSeparatorRight
- wbLeadHelperConfigWindowTabsSeparatorRight

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
