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
| Addons seen in | AdvancedPetAssist, AdvancedRenownTrainer, AggroMeter, Aura, TidyChat, TidyRoll, WSCT, WarBoard |
| Files seen in | `/workspace/data/raw/AdvancedPetAssist/APAGui.xml:0`, `/workspace/data/raw/AggroMeter/AggroMeter.xml:0`, `/workspace/data/raw/Aura/Source/Templates.xml:0`, `/workspace/data/raw/TidyChat/TidyChatTemplates.xml:0`, `/workspace/data/raw/TidyRoll/TidyRoll.xml:0`, `/workspace/data/raw/WarBoard/WarBoardOptions.xml:0`, `/workspace/data/raw/advancedrenowntrainer/AdvancedRenownTrainingWindow.xml:0`, `/workspace/data/raw/wsct/wsct_options/wsct_options.xml:0` |
| Namespaces detected | EA_Button_Tab |
| Source kinds | xml_attributes |
| Example locations | APAOptionsTabsAutoRecall, APAOptionsTabsControls, APAOptionsTabsFollowTarget, APAOptionsTabsGeneral, APAOptionsTabsHUD, APAOptionsTabsKiting |
| XML usage count | 19 |
| XML attribute usage count | 19 |
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

Observed engine XML template or inherited constant referenced by 8 addons.

## Seen In

- AdvancedPetAssist
- AdvancedRenownTrainer
- AggroMeter
- Aura
- TidyChat
- TidyRoll
- WSCT
- WarBoard

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
- AuraTabButtonTemplate
- TChatTabButton
- TRollTabButton
- TabButtonTemplate
- WSCTTabButton

## Related APIs

- [Button](../../xml/element_types/element_Button.md) (HIGH 100/100) - XML Element Type

## Used With

- none

## Triggered By

- none

## Affects

- none

## Notes

- none
