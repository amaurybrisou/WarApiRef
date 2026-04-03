# EA_Button_DefaultIconFrame_Large

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
| Addons seen in | Aura, Enemy |
| Files seen in | `/workspace/data/raw/Aura/Source/AuraConfig.xml:0`, `/workspace/data/raw/Aura/Source/AuraTexture.xml:0`, `/workspace/data/raw/Enemy/Code/Core/ChooseIconDialog.xml:0` |
| Namespaces detected | EA_Button_DefaultIconFrame_Large |
| Source kinds | xml_attributes |
| Example locations | AuraConfigGeneralIconButton, AuraIconButton, EnemyChooseIconDialogList_IconButton |
| XML usage count | 3 |
| XML attribute usage count | 3 |
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

Observed engine XML template or inherited constant referenced by 2 addons.

## Seen In

- Aura
- Enemy

## Used By

- AuraConfigGeneralIconButton
- AuraIconButton
- EnemyChooseIconDialogList_IconButton

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
