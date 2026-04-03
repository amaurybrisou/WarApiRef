# EA_FullResizeImage_DefaultFrame

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
| Addons seen in | Aura, DAoCBuff, bigger_MacroWindow |
| Files seen in | `/workspace/data/raw/Aura/Source/Templates.xml:0`, `/workspace/data/raw/DAoCBuff/Source/DAoCBuffSettingsTabs.xml:1033`, `/workspace/data/raw/DAoCBuff/Source/DAoCBuffSettingsTabs.xml:1399`, `/workspace/data/raw/DAoCBuff/Source/DAoCBuffSettingsTabs.xml:1596`, `/workspace/data/raw/DAoCBuff/Source/DAoCBuffSettingsTabs.xml:1683`, `/workspace/data/raw/DAoCBuff/Source/DAoCBuffSettingsTabs.xml:2485`, `/workspace/data/raw/DAoCBuff/Source/DAoCBuffSettingsTabs.xml:2604`, `/workspace/data/raw/DAoCBuff/Source/DAoCBuffSettingsTabs.xml:285` |
| Namespaces detected | EA_FullResizeImage_DefaultFrame |
| Source kinds | xml_attributes |
| Example locations | AuraWindowBackground, DAoCBuffFrameSettingsTab_ScrollChildFilterFrame, DAoCBuffFrameSettingsTab_ScrollChildFrameSettingsFrame, DAoCBuffFrameSettingsTab_ScrollChildLayoutFrame, DAoCBuffFrameSettingsTab_ScrollChildSortFrame, DAoCBuffListManagerTab_ScrollChildAddDeleteFrame |
| XML usage count | 12 |
| XML attribute usage count | 12 |
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
- bigger_MacroWindow

## Used By

- AuraWindowBackground
- DAoCBuffFrameSettingsTab_ScrollChildFilterFrame
- DAoCBuffFrameSettingsTab_ScrollChildFrameSettingsFrame
- DAoCBuffFrameSettingsTab_ScrollChildLayoutFrame
- DAoCBuffFrameSettingsTab_ScrollChildSortFrame
- DAoCBuffListManagerTab_ScrollChildAddDeleteFrame
- DAoCBuffListManagerTab_ScrollChildManagerFrame
- DAoCBuff_Settings_FilterFrame_ScrollChildActionFrame
- DAoCBuff_Settings_FilterFrame_ScrollChildConditionFrame
- DAoCBuff_Settings_FilterFrame_ScrollChildEnvironmentFrame
- DAoCBuff_Settings_FilterFrame_ScrollChildMiscFrame
- EA_Window_MacroMacrosBackground

## Related APIs

- [FullResizeImage](../../xml/element_types/element_FullResizeImage.md) (HIGH 100/100) - XML Element Type

## Used With

- none

## Triggered By

- none

## Affects

- none

## Notes

- none
