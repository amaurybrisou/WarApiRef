# EA_ComboBox_DefaultResizable_Fixed

- Category: Constant
- Confidence level: HIGH
- Confidence score: 100/100

## Confidence Assessment

- Level: HIGH

- Final score: 100/100

- Raw weighted score: 158

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, used directly in xml handler attributes, matches a known engine namespace.

## Evidence Signals

- +18 Seen in 2 to 3 addons: Cross-addon spread is present but limited.
- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +30 Used directly in XML handler attributes: XML exposure suggests an engine-level contract.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.
- +10 Referenced from initialization flow: Lifecycle reconstruction references this symbol.
- +15 Role is consistent across addons: The same symbol serves the same kind of job across addons.
- +25 Referenced by generated docs or reference files: The symbol is reinforced outside a single call page.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | LoyalPet, WSCT |
| Files seen in | `/workspace/LoyalPet/gui/lpet_gui.xml:112`, `/workspace/LoyalPet/gui/lpet_gui.xml:2560`, `/workspace/LoyalPet/gui/lpet_gui.xml:533`, `/workspace/LoyalPet/gui/lpet_gui.xml:561`, `/workspace/wsct/wsct_options/wsct_options.xml:1032`, `/workspace/wsct/wsct_options/wsct_options.xml:210`, `/workspace/wsct/wsct_options/wsct_options.xml:382`, `/workspace/wsct/wsct_options/wsct_options.xml:647` |
| Namespaces detected | EA_ComboBox_DefaultResizable_Fixed |
| Source kinds | flows, xml_attributes |
| Example locations | LPETComboBoxTemplateCombo, LPETOptionsPetAttackCombo, LPETOptionsPetFollowCombo, LPETOptionsProfilesCombo, WSCTComboBoxTemplateCombo, WSCTOptionsEventWindowCombo |
| XML usage count | 8 |
| XML attribute usage count | 8 |
| Lua usage count | 0 |
| Global usage count | 0 |
| Local definition count | 0 |
| Documentation references | 1 |
| Initialization flow references | 3 |
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

Observed engine XML template or inherited constant referenced by 2 addons.

## Seen In

- LoyalPet
- WSCT

## Used By

- LPETComboBoxTemplateCombo
- LPETOptionsPetAttackCombo
- LPETOptionsPetFollowCombo
- LPETOptionsProfilesCombo
- WSCTComboBoxTemplateCombo
- WSCTOptionsEventWindowCombo
- WSCTOptionsFrameWindowCombo
- WSCTOptionsProfileWindowCombo

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
