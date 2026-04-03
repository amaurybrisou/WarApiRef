# EA_Settings_SectionTitle

- Category: Constant
- Confidence level: HIGH
- Confidence score: 100/100

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 123

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, used directly in xml handler attributes, matches a known engine namespace.

## Evidence Signals

- +18 Seen in 2 to 3 addons: Cross-addon spread is present but limited.
- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +30 Used directly in XML handler attributes: XML exposure suggests an engine-level contract.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | Aura, DAoCBuff, PotionBar |
| Files seen in | `/workspace/data/raw/Aura/Source/AuraConfig.xml:0`, `/workspace/data/raw/DAoCBuff/Source/DAoCBuffMsgWindow.xml:0`, `/workspace/data/raw/DAoCBuff/Source/DAoCBuffSettingsTabs.xml:1008`, `/workspace/data/raw/DAoCBuff/Source/DAoCBuffSettingsTabs.xml:1043`, `/workspace/data/raw/DAoCBuff/Source/DAoCBuffSettingsTabs.xml:1409`, `/workspace/data/raw/DAoCBuff/Source/DAoCBuffSettingsTabs.xml:1606`, `/workspace/data/raw/DAoCBuff/Source/DAoCBuffSettingsTabs.xml:1693`, `/workspace/data/raw/DAoCBuff/Source/DAoCBuffSettingsTabs.xml:2166` |
| Namespaces detected | EA_Settings_SectionTitle |
| Source kinds | xml_attributes |
| Example locations | AuraConfigTriggerAbilityDragAndDropNote, DAoCBuffFrameSettingsTab_ScrollChild_FilterLabel, DAoCBuffFrameSettingsTab_ScrollChild_FrameSettingsLabel, DAoCBuffFrameSettingsTab_ScrollChild_LayoutLabel, DAoCBuffFrameSettingsTab_ScrollChild_SortLabel, DAoCBuffGeneralSettingsTab_ScrollChild_GeneralSettingsLabel |
| XML usage count | 15 |
| XML attribute usage count | 15 |
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

Observed engine XML template or inherited constant referenced by 3 addons.

## Seen In

- Aura
- DAoCBuff
- PotionBar

## Used By

- AuraConfigTriggerAbilityDragAndDropNote
- DAoCBuffFrameSettingsTab_ScrollChild_FilterLabel
- DAoCBuffFrameSettingsTab_ScrollChild_FrameSettingsLabel
- DAoCBuffFrameSettingsTab_ScrollChild_LayoutLabel
- DAoCBuffFrameSettingsTab_ScrollChild_SortLabel
- DAoCBuffGeneralSettingsTab_ScrollChild_GeneralSettingsLabel
- DAoCBuffGeneralSettingsTab_ScrollChild_GlobalFrameSettingsLabel
- DAoCBuffListManagerTab_ScrollChild_ListSettingsLabel
- DAoCBuffMessageWindowScrollWindowScrollChildTitleLabel
- DAoCBuff_Settings_FilterFrame_ScrollChild_ActionLabel
- DAoCBuff_Settings_FilterFrame_ScrollChild_ConditionLabel
- DAoCBuff_Settings_FilterFrame_ScrollChild_EnableFilterCheckBoxText
- DAoCBuff_Settings_FilterFrame_ScrollChild_EnvironmentLabel
- DAoCBuff_Settings_FilterFrame_ScrollChild_MiscLabel
- PotionBarAboutVersionInfo

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
