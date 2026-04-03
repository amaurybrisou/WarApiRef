# EA_Window_HelpDescriptionTemplate

- Category: Constant
- Confidence level: HIGH
- Confidence score: 90/100

## Confidence Assessment

- Level: HIGH

- Score: 90/100

- Rationale: Promoted as HIGH confidence because matches default ui or extracted base ui surface, used directly in xml handler attributes, matches a known engine namespace.

## Evidence Signals

- +35 Matches default UI or extracted base UI surface: Symbol aligns with known default-interface namespaces.
- +30 Used directly in XML handler attributes: XML exposure suggests an engine-level contract.
- +25 Matches a known engine namespace: Namespace shape matches WAR engine APIs.

## Evidence Summary

| Evidence | Value |
| --- | --- |
| Addons seen in | Calling |
| Files seen in | Calling.xml, CalljoinGUI.xml |
| Namespaces detected | EA_Window_HelpDescriptionTemplate |
| Source kinds | xml_attributes |
| Example locations | CallingSetupBindButtonText, CallingSetupBindSelectText, CallingSetupBindTargetText, CallingSetupGroupInviteLabel, CallingSetupGroupInviteOnRequestLabel, CallingSetupMacroInstructions |
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

Engine-supplied XML constant or template class referenced by 1 addons.

## Seen In

- Calling

## Used By

- CallingSetupBindButtonText
- CallingSetupBindSelectText
- CallingSetupBindTargetText
- CallingSetupGroupInviteLabel
- CallingSetupGroupInviteOnRequestLabel
- CallingSetupMacroInstructions
- CallingSetupPriosText
- CallingSetupTutorialDescription
- CallingSetupTutorialTutText
- CalljoinGUIWindowInstructions

## Related APIs

- [Label](../../xml/element_types/element_Label.md) (HIGH 100/100) - XML Element Type

## Notes

- none
