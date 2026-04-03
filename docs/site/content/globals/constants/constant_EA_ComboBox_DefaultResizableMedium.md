# EA_ComboBox_DefaultResizableMedium

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
| Addons seen in | BankArkel, CCTV, DeepSleep, RVMOD_SquaredDistances, Statdoll Remix, Warbuilder |
| Files seen in | BankArkel.xml, CCTV.xml, RVMOD_SquaredDistances.xml, Settings.xml, Source/Warbuilder.xml, StatdollOptions.xml |
| Namespaces detected | EA_ComboBox_DefaultResizableMedium |
| Source kinds | xml_attributes |
| Example locations | BankArkelBackpackCombo, CCTVSettingsCombobox, CCTVSettingsCombobox2, CCTVSettingsCombobox3, CCTVSettingsCombobox4, DeepSleep_Templates_CheckBoxComboBox |
| XML usage count | 9 |
| XML attribute usage count | 9 |
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

Engine-supplied XML constant or template class referenced by 6 addons.

## Seen In

- BankArkel
- CCTV
- DeepSleep
- RVMOD_SquaredDistances
- Statdoll Remix
- Warbuilder

## Used By

- BankArkelBackpackCombo
- CCTVSettingsCombobox
- CCTVSettingsCombobox2
- CCTVSettingsCombobox3
- CCTVSettingsCombobox4
- DeepSleep_Templates_CheckBoxComboBox
- RVMOD_SquaredDistancesSettingsWindowComboBoxAnchorPoints
- StatdollOptionsComboboxTemplateCombobox
- WarbuilderMainWindowCombo

## Related APIs

- [ComboBox](../../xml/element_types/element_ComboBox.md) (HIGH 100/100) - XML Element Type

## Notes

- none
